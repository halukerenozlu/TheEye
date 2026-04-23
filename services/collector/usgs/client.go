package usgs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultFeedURL = "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	feedURL    string
	httpClient *http.Client
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	ID         string            `json:"id"`
	Properties FeatureProperties `json:"properties"`
	Geometry   FeatureGeometry   `json:"geometry"`
}

type FeatureProperties struct {
	Time    int64    `json:"time"`
	Updated int64    `json:"updated"`
	Mag     *float64 `json:"mag"`
	Status  string   `json:"status"`
	Title   string   `json:"title"`
}

type FeatureGeometry struct {
	Type        string    `json:"type"`
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

func (c *Client) Fetch(ctx context.Context) (FeatureCollection, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.feedURL, nil)
	if err != nil {
		return FeatureCollection{}, fmt.Errorf("build usgs request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return FeatureCollection{}, fmt.Errorf("fetch usgs feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 256))
		return FeatureCollection{}, fmt.Errorf(
			"fetch usgs feed: unexpected status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var payload FeatureCollection
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return FeatureCollection{}, fmt.Errorf("decode usgs feed: %w", err)
	}

	if payload.Type != "FeatureCollection" {
		return FeatureCollection{}, fmt.Errorf("decode usgs feed: unexpected type %q", payload.Type)
	}

	return payload, nil
}
