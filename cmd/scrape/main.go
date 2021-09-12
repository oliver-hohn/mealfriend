package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	gcp "cloud.google.com/go/storage"
	"github.com/olekukonko/tablewriter"
	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
	"github.com/oliver-hohn/mealfriend/scrapers"
	"github.com/oliver-hohn/mealfriend/storage"
	"google.golang.org/protobuf/proto"
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

	c := scrapers.NewScraperClient()
	r, err := c.Run(u)
	if err != nil {
		log.Fatalf("unable to scrape %s: %v", *inputURL, err)
	}

	prettyPrintRecipe(r)

	if *shouldStore {
		if err := store(r); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("stored %s\n", r.Name)
		}
	}
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

func store(r *pbmodels.Recipe) error {
	ctx := context.Background()
	client, err := gcp.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("unable to create GCP client: %w", err)
	}

	b, err := proto.Marshal(r)
	if err != nil {
		return fmt.Errorf("unable to serialize Recipe proto: %w", err)
	}

	key := fmt.Sprintf("recipes/%s.pb", r.Code)
	if err := storage.Save(ctx, client, "mealfriend-datastore", key, b); err != nil {
		return fmt.Errorf("unable to write to GCP: %w", err)
	}

	return nil
}
