package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/oliver-hohn/mealfriend/helpers"
	"github.com/oliver-hohn/mealfriend/scrapers"
)

var inputURL = flag.String("input_url", "", "Where to scrape")

func main() {
	flag.Parse()

	if inputURL == nil || *inputURL == "" {
		log.Fatal("missing input_url")
	}

	u, err := url.Parse(*inputURL)
	if err != nil {
		log.Fatalf("invalid input_url: %v", err)
	}

	r, err := scrapers.Scrape(u)
	if err != nil {
		log.Fatalf("unable to scrape %s: %v", *inputURL, err)
	}

	helpers.PrettyPrintRecipe(r)
}
