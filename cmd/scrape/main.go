package main

import (
	"flag"
	"fmt"
	"log"
	"main/models"
	"main/scrapers"
	"net/url"
	"os"

	"github.com/olekukonko/tablewriter"
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

func prettyPrintRecipe(r *models.Recipe) {
	fmt.Printf("Title: %s\n", r.Name)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Quantity"})

	for _, i := range r.Ingredients {
		table.Append([]string{
			i.Name, i.Type.GetName(), fmt.Sprintf("%s %s", i.Quantity.Amount, i.Quantity.Unit),
		})
	}
	table.Render()
}
