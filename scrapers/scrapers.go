package scrapers

import (
	"fmt"
	"main/models"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
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

func (s *ScraperClient) Run(u *url.URL) (*models.Recipe, error) {
	switch u.Hostname() {
	case CAFE_DELITES_HOST:
		return s.scrapeFromCafeDelites(u)
	default:
		return nil, fmt.Errorf("unsupported host: %s", u.Hostname())
	}
}
