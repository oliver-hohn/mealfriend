package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/helpers"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers"
)

var seedFilePath = flag.String("seed_file", "", "path to seed file")

func main() {
	flag.Parse()

	if seedFilePath == nil || *seedFilePath == "" {
		log.Fatal("missing seed_file")
	}

	recipes, err := parseSeedFile(*seedFilePath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(envs.MustGetEnv("NEO4J_URI"), neo4j.BasicAuth(envs.MustGetEnv("NEO4J_USERNAME"), envs.MustGetEnv("NEO4J_PASSWORD"), ""))
	if err != nil {
		log.Fatalf("unable to initialize neo4j driver: %v", err)
	}
	defer driver.Close(ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		log.Fatalf("invalid neo4j connection: %v", err)
	}

	if err := clearGraph(ctx, driver); err != nil {
		log.Fatalf("unable to clear graph: %v", err)
	}

	if err := seedTags(ctx, driver); err != nil {
		log.Fatalf("unable to seed tags: %v", err)
	}

	for _, r := range recipes {
		if err := saveRecipe(ctx, driver, r); err != nil {
			log.Fatalf("unable to save %s: %v", r.Code, err)
		}
	}

	for _, r := range recipes {
		helpers.PrettyPrintRecipe(r)
		fmt.Println()
	}
}

func parseSeedFile(p string) ([]*models.Recipe, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("unable to resolve working directory: %w", err)
	}

	b, err := os.ReadFile(path.Join(wd, p))
	if err != nil {
		return nil, fmt.Errorf("unable to read %s: %w", p, err)
	}

	recipes := []*models.Recipe{}

	r := csv.NewReader(bytes.NewReader(b))
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("unable to parse row: %w", err)
		}

		recipe, err := parseSeedRow(row)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func parseSeedRow(row []string) (*models.Recipe, error) {
	u, err := url.Parse(row[0])
	if err != nil {
		return nil, fmt.Errorf("invalid url %v: %w", row[0], err)
	}

	recipe, err := scrapers.Scrape(u)
	if err != nil {
		return nil, fmt.Errorf("unable to scrape %s: %w", u.String(), err)
	}

	fmt.Printf("parsed %s\n", u.String())

	return recipe, nil
}

func clearGraph(ctx context.Context, driver neo4j.DriverWithContext) error {
	_, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		`match (n) detach delete n`,
		map[string]interface{}{},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("unable to create recipe: %w", err)
	}

	return nil
}

func seedTags(ctx context.Context, driver neo4j.DriverWithContext) error {
	var query bytes.Buffer
	params := map[string]interface{}{}
	for i, t := range models.AvailableTags {
		paramKey := fmt.Sprintf("value%d", i)
		query.WriteString(fmt.Sprintf("create (t%d:Tag {value: $%s})\n", i, paramKey))
		params[paramKey] = t
	}

	_, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		query.String(),
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("unable to create recipe: %w", err)
	}

	return nil
}

func saveRecipe(ctx context.Context, driver neo4j.DriverWithContext, r *models.Recipe) error {
	_, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		`create (r:Recipe {name: $name, code: $code, source: $source}) return r`,
		map[string]interface{}{"name": r.Name, "code": r.Code, "source": r.Source.String()},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return fmt.Errorf("unable to create recipe: %w", err)
	}

	for _, t := range r.Tags {
		fmt.Printf("creating tag: %s\n", t)
		_, err = neo4j.ExecuteQuery[*neo4j.EagerResult](
			ctx,
			driver,
			`match (r:Recipe {code: $code})
			 match (t:Tag {value: $tag})
			 create (r)-[:tagged_as]->(t)
			`,
			map[string]interface{}{"code": r.Code, "tag": t},
			neo4j.EagerResultTransformer,
		)

		if err != nil {
			return fmt.Errorf("unable to link tags: %w", err)
		}
	}

	return nil
}
