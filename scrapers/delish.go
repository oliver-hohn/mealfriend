package scrapers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oliver-hohn/mealfriend/models"
	"github.com/oliver-hohn/mealfriend/scrapers/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const DELISH_HOST = "delish.com"

type DelishScraper struct {
	httpClient *http.Client
}

func NewDelishScraper(httpClient *http.Client) *DelishScraper {
	return &DelishScraper{httpClient: httpClient}
}

func (s *DelishScraper) Run(u *url.URL) (*models.Recipe, error) {
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

	// Extract recipe name from the URL, as there is no easy CSS selector to scrape
	path := strings.TrimSuffix(u.Path, "/")
	pathSegments := strings.Split(path, "/")
	if len(pathSegments) == 0 {
		return nil, fmt.Errorf(
			"unable to extract title from URL path: %s", u.Path,
		)
	}
	titleSegment := pathSegments[len(pathSegments)-1]
	title := strings.ReplaceAll(titleSegment, "-", " ")
	title = strings.TrimSuffix(title, "recipe")

	recipeName := cases.Title(language.English).String(title)

	ingredients := []string{}
	doc.Find(".ingredient-lists > li").Each(func(i int, s *goquery.Selection) {
		ingredients = append(ingredients, s.Text())
	})

	r := utils.NewRecipe(recipeName, ingredients, u)

	cookTime, err := s.extractCookTime(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to extract cook time: %w", err)
	}
	r.CookTime = cookTime

	return r, nil
}

func (s *DelishScraper) extractCookTime(doc *goquery.Document) (time.Duration, error) {
	var totalTimeNode *goquery.Selection
	doc.Find("dt").Each(func(i int, s *goquery.Selection) {
		if strings.TrimSpace(strings.ToLower(s.Text())) == "total time:" {
			totalTimeNode = s
		}
	})
	if totalTimeNode == nil {
		return -1, fmt.Errorf("unable to find total time HTML element")
	}

	cookTimes := totalTimeNode.Siblings().Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	return utils.NewDuration(strings.Join(cookTimes, " "))
}
