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
