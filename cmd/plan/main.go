package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/helpers"
	"github.com/oliver-hohn/mealfriend/models"
)

var poultry = flag.Int("poultry", -1, "TODO")
var fish = flag.Int("fish", -1, "TODO")

var count = flag.Int("count", 5, "number of recipes to plan")

func main() {
	flag.Parse()

	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(envs.MustGetEnv("NEO4J_URI"), neo4j.BasicAuth(envs.MustGetEnv("NEO4J_USERNAME"), envs.MustGetEnv("NEO4J_PASSWORD"), ""))
	if err != nil {
		log.Fatalf("unable to initialize neo4j driver: %v", err)
	}
	defer driver.Close(ctx)

	if err := driver.VerifyConnectivity(ctx); err != nil {
		log.Fatalf("invalid neo4j connection: %v", err)
	}

	recipes := []*models.Recipe{}

	if *poultry > 0 {
		res, err := fetchRecipesByTag(ctx, driver, models.POULTRY, *poultry, collectCodes(recipes))
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, res...)
	}

	if *fish > 0 {
		res, err := fetchRecipesByTag(ctx, driver, models.FISH, *fish, collectCodes(recipes))
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, res...)
	}

	if len(recipes) < *count {
		res, err := fetchRecipes(ctx, driver, *count-len(recipes), collectCodes(recipes))
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, res...)
	}

	for _, r := range recipes {
		helpers.PrettyPrintRecipe(r)
		fmt.Println()
	}
}

func fetchRecipesByTag(ctx context.Context, driver neo4j.DriverWithContext, t models.Tag, count int, excludedRecipes []string) ([]*models.Recipe, error) {
	ret := make([]*models.Recipe, count)

	res, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		`MATCH (r:Recipe)-[:tagged_as]->(t:Tag)
		WHERE t.value = $tag AND NOT r.code IN $excludedRecipes
		RETURN r`,
		map[string]interface{}{"tag": t, "excludedRecipes": excludedRecipes},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch recipe for tag: %s, due to: %w", t, err)
	}

	for i := 0; i < len(ret); i++ {
		j := rand.Intn(len(res.Records))

		recipe, err := parseResult(res.Records[j])
		if err != nil {
			return nil, fmt.Errorf("unable to parse result (i: %d): %w", i, err)
		}

		ret[i] = recipe
	}

	return ret, nil
}

func fetchRecipes(ctx context.Context, driver neo4j.DriverWithContext, count int, excludedRecipes []string) ([]*models.Recipe, error) {
	ret := make([]*models.Recipe, count)

	res, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		`MATCH (r:Recipe)
		WHERE NOT r.code IN $excludedRecipes
		RETURN r`,
		map[string]interface{}{"excludedRecipes": excludedRecipes},
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch recipe due to: %w", err)
	}

	for i := 0; i < len(ret); i++ {
		j := rand.Intn(len(res.Records))

		recipe, err := parseResult(res.Records[j])
		if err != nil {
			return nil, fmt.Errorf("unable to parse result (i: %d): %w", i, err)
		}

		ret[i] = recipe
	}

	return ret, nil
}

func parseResult(r *db.Record) (*models.Recipe, error) {
	node, _, err := neo4j.GetRecordValue[neo4j.Node](r, "r")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch node: %w", err)
	}

	code, err := neo4j.GetProperty[string](node, "code")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch code from %w", err)
	}
	name, err := neo4j.GetProperty[string](node, "name")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch name from %w", err)
	}
	source, err := neo4j.GetProperty[string](node, "source")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch source from %w", err)
	}

	sourceUrl, err := url.Parse(source)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s as a URL: %w", source, err)
	}

	return &models.Recipe{Code: code, Name: name, Source: sourceUrl}, nil
}

func collectCodes(recipes []*models.Recipe) []string {
	ret := []string{}
	for _, r := range recipes {
		ret = append(ret, r.Code)
	}

	return ret
}
