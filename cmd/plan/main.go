package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/models"
	"gorm.io/gorm"
)

var count = flag.Int("count", -1, "number of recipes")

func main() {
	flag.Parse()

	if *count < 1 {
		log.Fatal("provide a >0 value for count")
	}

	rand.Seed(time.Now().UnixNano())

	db, err := database.CreateConn(database.DatabaseConfig{
		Host:     envs.MustGetEnv("PGHOST"),
		Port:     envs.MustGetIntEnv("PGPORT"),
		Database: envs.MustGetEnv("PGDATABASE"),
		Username: envs.MustGetEnv("PGUSER"),
		Password: envs.MustGetEnv("PGPASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}

	recipes, err := getRecipes(db)
	if err != nil {
		log.Fatal(err)
	}

	menu := generateMenu(recipes, *count)
	prettyPrintMenu(menu)
}

func getRecipes(db *gorm.DB) ([]*models.Recipe, error) {
	var recipes []*models.Recipe
	if err := db.Preload("Ingredients").Find(&recipes).Error; err != nil {
		return nil, fmt.Errorf("unable to fetch recipes: %w", err)
	}

	return recipes, nil
}

func generateMenu(recipes []*models.Recipe, count int) []*models.Recipe {
	// shuffle recipes to add variety
	rand.Shuffle(len(recipes), func(a, b int) { recipes[a], recipes[b] = recipes[b], recipes[a] })

	menu := []*models.Recipe{}

	for i := 0; i < count; i++ {
		menu = append(menu, recipes[0])

		// remove already selected recipie
		recipes = recipes[1:]
	}

	return menu
}

func prettyPrintMenu(m []*models.Recipe) {
	data := [][]string{}

	for _, recipe := range m {
		data = append(data, []string{recipe.Name})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Menu"})
	table.SetRowSeparator("-")

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
