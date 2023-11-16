package scrapers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers/utils"
)

const CAFE_DELITES_HOST = "cafedelites.com"

type CafeDelitesScraper struct {
	httpClient *http.Client
}

func NewCafeDelitesScraper(httpClient *http.Client) *CafeDelitesScraper {
	return &CafeDelitesScraper{httpClient: httpClient}
}

func (s *CafeDelitesScraper) Run(u *url.URL) (*models.Recipe, error) {
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
	nameSelection := doc.Find(".wprm-recipe-name")
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
	doc.Find(".wprm-recipe-ingredients > .wprm-recipe-ingredient").Each(func(i int, s *goquery.Selection) {
		ingredients = append(ingredients, s.Text())
	})

	r := utils.NewRecipe(recipeName, ingredients, u)

	cookTime, err := s.extractCookTime(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to parse cook time: %w", err)
	}
	r.CookTime = cookTime

	imageURL, err := s.extractImageURL(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to extract image from recipe page: %w", err)
	}
	r.Image = imageURL

	return r, nil
}

func (s *CafeDelitesScraper) extractCookTime(doc *goquery.Document) (time.Duration, error) {
	cookTimes := doc.Find(".wprm-recipe-total-time-container > .wprm-recipe-time > .wprm-recipe-details").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	return utils.NewDuration(strings.Join(cookTimes, " "))
}

func (s *CafeDelitesScraper) extractImageURL(doc *goquery.Document) (*url.URL, error) {
	// In all seen cases, the first img in the HTML that is centered and large best
	// represents the recipe.
	var image *goquery.Selection
	doc.Find("img").EachWithBreak(func(i int, s *goquery.Selection) bool {
		classes, exists := s.Attr("class")
		if exists && strings.Contains(classes, "aligncenter") && strings.Contains(classes, "size-full") {
			// Break out the loop early once we find the first selection that matches
			image = s
			return false
		}

		return true
	})

	if image == nil {
		return nil, fmt.Errorf("no image found")
	}

	src, exists := image.Attr("data-lazy-src")
	if !exists {
		return nil, errors.New("no src attribute found on image")
	}

	u, err := url.Parse(src)
	if err != nil {
		return nil, fmt.Errorf("unable to parse image src: %s, into a URL: %w", src, err)
	}

	return u, nil
}
