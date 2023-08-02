package scrapers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/oliver-hohn/mealfriend/models"
)

type Scraper interface {
	Run(u *url.URL) (*models.Recipe, error)
}

func Scrape(u *url.URL) (*models.Recipe, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	var s Scraper

	host := strings.TrimPrefix(u.Hostname(), "www.")

	switch host {
	case CAFE_DELITES_HOST:
		s = NewCafeDelitesScraper(retryClient.HTTPClient)
	case DELISH_HOST:
		s = NewDelishScraper(retryClient.HTTPClient)
	case TASTY_HOST:
		s = NewTastyScraper(retryClient.HTTPClient)
	default:
		return nil, fmt.Errorf("unsupported host: %s", host)
	}

	return s.Run(u)
}
