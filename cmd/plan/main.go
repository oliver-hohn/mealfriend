package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/graph"
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
		res, err := graph.FindRecipes(ctx, driver, models.POULTRY, *poultry, collectCodes(recipes))
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, res...)
	}

	if *fish > 0 {
		res, err := graph.FindRecipes(ctx, driver, models.POULTRY, *fish, collectCodes(recipes))
		if err != nil {
			log.Fatal(err)
		}

		recipes = append(recipes, res...)
	}

	if len(recipes) < *count {
		res, err := graph.FindRecipes(ctx, driver, models.UNSPECIFIED, *count-len(recipes), collectCodes(recipes))
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

func collectCodes(recipes []*models.Recipe) []string {
	ret := []string{}
	for _, r := range recipes {
		ret = append(ret, r.Code)
	}

	return ret
}
