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

	r := utils.NewRecipe(recipeName, ingredients, u)

	cookTime, err := s.extractCookTime(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to extract cook time: %w", err)
	}
	r.CookTime = cookTime

	imageURL, err := s.extractImageURL(recipeName)
	if err != nil {
		return nil, fmt.Errorf("unable to extract image URL: %w", err)
	}
	r.Image = imageURL

	return r, nil
}

func (s *TastyScraper) extractCookTime(doc *goquery.Document) (time.Duration, error) {
	// Tasty does not provide a "total/cook time", so instead calculate it from the times
	// in the instructions.
	totalCookTime := time.Duration(0)

	var err error
	doc.Find(".prep-steps > li").EachWithBreak(func(i int, s *goquery.Selection) bool {
		text := strings.TrimSpace(s.Text())
		if len(text) == 0 {
			return true
		}

		d, durErr := utils.NewDuration(text)
		if durErr != nil {
			// Ignore NoDurationErrors, as they are returned when the instruction step does not
			// have a duration
			if _, ok := durErr.(*utils.NoDurationError); ok {
				return true
			}

			err = fmt.Errorf("unable to parse duration in instruction: \"%s\"; due to: %w", text, err)
			return false
		}

		totalCookTime += d

		return true
	})

	return totalCookTime, err
}

func (s *TastyScraper) extractImageURL(recipeName string) (*url.URL, error) {
	// Query tasty to find the image for the recipe in the search response (it is not present on the recipe
	// page)
	searchURL := fmt.Sprintf("https://%s/search?q=%s", TASTY_HOST, url.QueryEscape(recipeName))

	res, err := s.httpClient.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch %s: %w", searchURL, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 status code received: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML in response: %w", err)
	}

	image := doc.Find(".feed__items > li img").First()
	if image == nil {
		return nil, fmt.Errorf("no image found")
	}

	src, exists := image.Attr("src")
	if !exists {
		return nil, errors.New("no src attribute found on image")
	}

	u, err := url.Parse(src)
	if err != nil {
		return nil, fmt.Errorf("unable to parse image src: %s, into a URL: %w", src, err)
	}

	// Clear the query params to get a non-resized image URL
	u.RawQuery = ""

	return u, nil
}
