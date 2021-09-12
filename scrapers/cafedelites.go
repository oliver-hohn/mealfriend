package scrapers

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
	"github.com/oliver-hohn/mealfriend/scrapers/utils"
)

const CAFE_DELITES_HOST = "cafedelites.com"

func (s *ScraperClient) scrapeFromCafeDelites(u *url.URL) (*pbmodels.Recipe, error) {
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

	recipe := pbmodels.Recipe{}

	// Ensure only one node for the name is found
	nameSelection := doc.Find(".wprm-recipe-name")
	if len(nameSelection.Nodes) != 1 {
		return nil, fmt.Errorf(
			"unexpected recipe format: %d node(s) for the name have been found, expected 1", len(nameSelection.Nodes),
		)
	}
	nameSelection.Each(func(i int, s *goquery.Selection) {
		recipe.Name = s.Text()
	})

	doc.Find(".wprm-recipe-ingredients > .wprm-recipe-ingredient").Each(func(i int, s *goquery.Selection) {
		recipe.Ingredients = append(recipe.Ingredients, utils.NewIngredient(s.Text()))
	})

	return &recipe, nil
}
