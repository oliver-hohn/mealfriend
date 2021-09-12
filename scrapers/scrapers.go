package scrapers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
	pbmodels "github.com/oliver-hohn/mealfriend/protos/models"
)

type ScraperClient struct {
	httpClient *http.Client
}

func NewScraperClient() *ScraperClient {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	return &ScraperClient{
		httpClient: retryClient.StandardClient(),
	}
}

func (s *ScraperClient) Run(u *url.URL) (*pbmodels.Recipe, error) {
	switch u.Hostname() {
	case CAFE_DELITES_HOST:
		return s.scrapeFromCafeDelites(u)
	case TASTY_HOST:
		return s.scrapeFromTasty(u)
	default:
		return nil, fmt.Errorf("unsupported host: %s", u.Hostname())
	}
}
