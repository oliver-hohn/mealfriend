package scrapers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers/utils"
)

const TASTY_HOST = "tasty.co"

type TastyScraper struct {
	httpClient *http.Client
}

func NewTastyScraper(httpClient *http.Client) *TastyScraper {
	return &TastyScraper{httpClient: httpClient}
}

func (s *TastyScraper) Run(u *url.URL) (*models.Recipe, error) {
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

	// Ensure only one node for the name is found
	nameSelection := doc.Find(".recipe-name")
	if len(nameSelection.Nodes) != 1 {
		return nil, fmt.Errorf(
			"unexpected recipe format: %d node(s) for the name have been found, expected 1", len(nameSelection.Nodes),
		)
	}
	var recipeName string
	nameSelection.Each(func(i int, s *goquery.Selection) {
		recipeName = s.Text()
	})

	ingredients := []string{}
	doc.Find(".ingredient").Each(func(i int, s *goquery.Selection) {
		ingredients = append(ingredients, s.Text())
	})

	return utils.NewRecipe(recipeName, ingredients, u), nil
}
