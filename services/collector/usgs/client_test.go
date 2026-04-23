package usgs

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchSuccessDecode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"type":"FeatureCollection",
			"features":[
				{
					"id":"abc123",
					"properties":{
						"time":1700000000000,
						"updated":1700000000500,
						"mag":3.4,
						"status":"reviewed",
						"title":"M 3.4 - Test"
					},
					"geometry":{
						"type":"Point",
						"coordinates":[-117.5,35.7,10.1]
					}
				}
			]
		}`))
	}))
	defer srv.Close()

	client := NewClient(srv.URL, &http.Client{Timeout: 2 * time.Second})
	got, err := client.Fetch(context.Background())
	if err != nil {
		t.Fatalf("Fetch returned error: %v", err)
	}

	if got.Type != "FeatureCollection" {
		t.Fatalf("type = %q, want %q", got.Type, "FeatureCollection")
	}
	if len(got.Features) != 1 {
		t.Fatalf("features length = %d, want 1", len(got.Features))
	}

	f := got.Features[0]
	if f.ID != "abc123" {
		t.Fatalf("id = %q, want %q", f.ID, "abc123")
	}
	if f.Properties.Time != 1700000000000 {
		t.Fatalf("properties.time = %d, want %d", f.Properties.Time, int64(1700000000000))
	}
	if f.Properties.Updated != 1700000000500 {
		t.Fatalf("properties.updated = %d, want %d", f.Properties.Updated, int64(1700000000500))
	}
	if f.Properties.Mag == nil || *f.Properties.Mag != 3.4 {
		t.Fatalf("properties.mag = %v, want %v", f.Properties.Mag, 3.4)
	}
	if f.Properties.Status != "reviewed" {
		t.Fatalf("properties.status = %q, want %q", f.Properties.Status, "reviewed")
	}
	if f.Properties.Title != "M 3.4 - Test" {
		t.Fatalf("properties.title = %q, want %q", f.Properties.Title, "M 3.4 - Test")
	}
	if len(f.Geometry.Coordinates) != 3 {
		t.Fatalf("geometry.coordinates length = %d, want 3", len(f.Geometry.Coordinates))
	}
}

func TestFetchNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "upstream down", http.StatusBadGateway)
	}))
	defer srv.Close()

	client := NewClient(srv.URL, &http.Client{Timeout: 2 * time.Second})
	_, err := client.Fetch(context.Background())
	if err == nil {
		t.Fatal("Fetch error = nil, want non-nil")
	}

	if !strings.Contains(err.Error(), "unexpected status 502") {
		t.Fatalf("error = %q, want contains %q", err.Error(), "unexpected status 502")
	}
}

func TestFetchMalformedJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"type":"FeatureCollection","features":[`))
	}))
	defer srv.Close()

	client := NewClient(srv.URL, &http.Client{Timeout: 2 * time.Second})
	_, err := client.Fetch(context.Background())
	if err == nil {
		t.Fatal("Fetch error = nil, want non-nil")
	}

	if !strings.Contains(err.Error(), "decode usgs feed") {
		t.Fatalf("error = %q, want contains %q", err.Error(), "decode usgs feed")
	}
}
