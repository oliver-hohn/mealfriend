package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers"
	"gorm.io/gorm"
)

var inputURL = flag.String("input_url", "", "Where to scrape")
var shouldStore = flag.Bool("store", false, "Whether the fetched recipe should be stored")

func main() {
	flag.Parse()

	if inputURL == nil || *inputURL == "" {
		log.Fatal("missing input_url")
	}

	u, err := url.Parse(*inputURL)
	if err != nil {
		log.Fatalf("invalid input_url: %v", err)
	}

	conn, err := database.CreateConn(database.DatabaseConfig{
		Host:     envs.MustGetEnv("PGHOST"),
		Port:     envs.MustGetIntEnv("PGPORT"),
		Database: envs.MustGetEnv("PGDATABASE"),
		Username: envs.MustGetEnv("PGUSER"),
		Password: envs.MustGetEnv("PGPASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}

	r, err := scrapers.Scrape(u)
	if err != nil {
		log.Fatalf("unable to scrape %s: %v", *inputURL, err)
	}

	prettyPrintRecipe(r)

	if *shouldStore {
		if err := store(conn, r); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("stored %s\n", r.Name)
		}
	}
}

func prettyPrintRecipe(r *models.Recipe) {
	fmt.Printf("Title: %s\n", r.Name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type"})

	for _, i := range r.Ingredients {
		// TODO: Add Ingredient Type
		table.Append([]string{
			i.Name, "",
		})
	}
	table.Render()
}

func store(db *gorm.DB, r *models.Recipe) error {
	if err := db.Create(r).Error; err != nil {
		return fmt.Errorf("unable to write the recipe to the DB: %w", err)
	}

	return nil
}
