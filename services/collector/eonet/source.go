package eonet

import (
	"context"
	"encoding/json"
	"log"
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

	logRawEvents(feed.Events)

	return NormalizeEvents(feed.Events), nil
}

func (s *Source) Name() string {
	return "eonet"
}

func logRawEvents(events []Event) {
	limit := 3
	if len(events) < limit {
		limit = len(events)
	}

	raw, err := json.Marshal(events[:limit])
	if err != nil {
		log.Printf("eonet raw events: marshal failed: %v", err)
		return
	}

	log.Printf("eonet raw events first %d/%d: %s", limit, len(events), raw)
}
