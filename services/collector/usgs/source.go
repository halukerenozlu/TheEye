package usgs

import (
	"context"
	"net/http"

	"theeye/services/collector/models"
)

type Source struct {
	client *Client
}

func NewSource(feedURL string, httpClient *http.Client) *Source {
	return &Source{
		client: NewClient(feedURL, httpClient),
	}
}

func (s *Source) Fetch(ctx context.Context) ([]models.NormalizedEvent, error) {
	feed, err := s.client.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return NormalizeFeatures(feed.Features), nil
}

func (s *Source) Name() string {
	return "usgs"
}
