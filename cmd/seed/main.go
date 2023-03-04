package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"

	"github.com/oliver-hohn/mealfriend/database"
	"github.com/oliver-hohn/mealfriend/envs"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers"
	"gorm.io/gorm"
)

var seedFilePath = flag.String("seed_file", "", "path to seed file")
var dryRun = flag.Bool("dryrun", false, "set to true to not store the parsed recipes")

func main() {
	flag.Parse()

	if seedFilePath == nil || *seedFilePath == "" {
		log.Fatal("missing seed_file")
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

	recipes, err := parseSeedFile(*seedFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if !*dryRun {
		clearRecipes(conn)
		storeRecipes(conn, recipes)
	}
}

func clearRecipes(db *gorm.DB) {
	db.Unscoped().Where("1=1").Delete(&models.Ingredient{})
	db.Unscoped().Where("1=1").Delete(&models.Recipe{})
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

func storeRecipes(db *gorm.DB, recipes []*models.Recipe) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, r := range recipes {
			if err := db.Create(r).Error; err != nil {
				return fmt.Errorf("unable to store %s: %w", r.Code, err)
			}

			fmt.Printf("stored %s\n", r.Code)
		}

		return nil
	})
}
