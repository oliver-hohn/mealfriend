package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/olekukonko/tablewriter"
	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
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

	c := scrapers.NewScraperClient()
	r, err := c.Run(u)
	if err != nil {
		log.Fatalf("unable to scrape %s: %v", *inputURL, err)
	}

	prettyPrintRecipe(r)
}

func prettyPrintRecipe(r *pbmodels.Recipe) {
	fmt.Printf("Title: %s\n", r.Name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type"})

	for _, i := range r.Ingredients {
		table.Append([]string{
			i.Name, i.GetType().String(),
		})
	}
	table.Render()
}
