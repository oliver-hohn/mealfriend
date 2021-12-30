package scrapers

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/iancoleman/strcase"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers/utils"
)

const TASTY_HOST = "tasty.co"

func (s *ScraperClient) scrapeFromTasty(u *url.URL) (*models.Recipe, error) {
	res, err := s.httpClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("unable to fetch %s: %w", u.String(), err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 status code received: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML in response: %w", err)
	}

	recipe := models.Recipe{}

	// Ensure only one node for the name is found
	nameSelection := doc.Find(".recipe-name")
	if len(nameSelection.Nodes) != 1 {
		return nil, fmt.Errorf(
			"unexpected recipe format: %d node(s) for the name have been found, expected 1", len(nameSelection.Nodes),
		)
	}
	nameSelection.Each(func(i int, s *goquery.Selection) {
		recipe.Name = s.Text()
		recipe.Code = strcase.ToCamel(recipe.Name)
	})

	doc.Find(".ingredient").Each(func(i int, s *goquery.Selection) {
		recipe.Ingredients = append(recipe.Ingredients, utils.NewIngredient(s.Text()))
	})

	return &recipe, nil
}
