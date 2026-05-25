package eonet

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultFeedURL = "https://eonet.gsfc.nasa.gov/api/v3/events?status=open&limit=100"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	feedURL    string
	httpClient *http.Client
}

type EventCollection struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Categories []Category `json:"categories"`
	Geometries []Geometry `json:"geometry"`
}

type Category struct {
	ID string `json:"id"`
}

type Geometry struct {
	Date        string    `json:"date"`
	Coordinates []float64 `json:"coordinates"`
}

func NewClient(feedURL string, httpClient *http.Client) *Client {
	if feedURL == "" {
		feedURL = DefaultFeedURL
	}

	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	return &Client{
		feedURL:    feedURL,
		httpClient: httpClient,
	}
}

func (c *Client) Fetch(ctx context.Context) (EventCollection, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.feedURL, nil)
	if err != nil {
		return EventCollection{}, fmt.Errorf("build eonet request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return EventCollection{}, fmt.Errorf("fetch eonet feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256))
		return EventCollection{}, fmt.Errorf(
			"fetch eonet feed: unexpected status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var payload EventCollection
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return EventCollection{}, fmt.Errorf("decode eonet feed: %w", err)
	}

	return payload, nil
}
