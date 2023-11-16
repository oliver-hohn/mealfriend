package graph

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/oliver-hohn/mealfriend/models"
)

func SaveRecipe(ctx context.Context, driver neo4j.DriverWithContext, r *models.Recipe) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(
			ctx,
			`create (r:Recipe {name: $name, code: $code, source: $source, ingredients: $ingredients, cookTime: $cookTime, imageURL: $imageURL}) return r`,
			map[string]interface{}{
				"name":        r.Name,
				"code":        r.Code,
				"source":      r.Source.String(),
				"ingredients": r.Ingredients,
				"cookTime":    r.CookTimeStr(),
				"imageURL":    r.ImageURL(),
			},
		)
		if err != nil {
			return nil, fmt.Errorf("unable to create recipe: %w", err)
		}

		for _, t := range r.Tags {
			_, err = tx.Run(
				ctx,
				`match (r:Recipe {code: $code})
				match (t:Tag {value: $tag})
				create (r)-[:tagged_as]->(t)`,
				map[string]interface{}{"code": r.Code, "tag": t},
			)

			if err != nil {
				return nil, fmt.Errorf("unable to link tags: %w", err)
			}
		}

		return nil, nil
	})

	return err
}

func FindRecipes(ctx context.Context, driver neo4j.DriverWithContext, t models.Tag, count int, excludedRecipeCodes []string) ([]*models.Recipe, error) {
	var query string
	params := map[string]interface{}{"excludedRecipes": excludedRecipeCodes, "count": count}

	if t == models.UNSPECIFIED {
		query = `MATCH (r:Recipe)
		WHERE NOT r.code IN $excludedRecipes
		WITH r, rand() AS rnd
		ORDER BY rnd
		RETURN r
		LIMIT $count`
	} else {
		query = `CALL {
			MATCH (r:Recipe)-[:tagged_as]->(t:Tag)
			WHERE t.value = $tag AND NOT r.code IN $excludedRecipes
			RETURN r as r
		}
		WITH r, rand() as rnd
		ORDER BY rnd
		RETURN r
		LIMIT $count`

		params["tag"] = t
	}

	res, err := neo4j.ExecuteQuery[*neo4j.EagerResult](
		ctx,
		driver,
		query,
		params,
		neo4j.EagerResultTransformer,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch recipe, due to: %w", err)
	}

	ret := make([]*models.Recipe, len(res.Records))
	for i, r := range res.Records {
		recipe, err := parseResult(r)
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
	sourceURL, err := url.Parse(source)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s as a URL: %w", source, err)
	}

	cookTimeStr, err := neo4j.GetProperty[string](node, "cookTime")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch cook time from %w", err)
	}
	cookTime, err := time.ParseDuration(cookTimeStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse: %s, back into a duration: %w", cookTimeStr, err)
	}

	rawImageURL, err := neo4j.GetProperty[string](node, "imageURL")
	if err != nil {
		return nil, fmt.Errorf("unable to fetch image URL from %w", err)
	}
	imageURL, err := url.Parse(rawImageURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s as a URL: %w", imageURL, err)
	}

	return &models.Recipe{Code: code, Name: name, Source: sourceURL, CookTime: cookTime, Image: imageURL}, nil
}
