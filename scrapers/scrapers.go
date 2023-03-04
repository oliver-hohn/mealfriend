package scrapers

import (
	"fmt"
	"net/url"

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
	switch u.Hostname() {
	case CAFE_DELITES_HOST:
		s = NewCafeDelitesScraper(retryClient.HTTPClient)
	case TASTY_HOST:
		s = NewTastyScraper(retryClient.HTTPClient)
	default:
		return nil, fmt.Errorf("unsupported host: %s", u.Hostname())
	}

	return s.Run(u)
}
