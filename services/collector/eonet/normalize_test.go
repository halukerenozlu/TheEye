package eonet

import "testing"

func TestNormalizeEventsMapsNormalEvent(t *testing.T) {
	got := NormalizeEvents([]Event{
		{
			ID:    "EONET_123",
			Title: "Wildfire Test",
			Categories: []Category{
				{ID: "wildfires"},
			},
			Geometries: []Geometry{
				{
					Date:        "2026-05-24T12:00:00Z",
					Coordinates: []float64{-121.5, 38.25},
				},
			},
		},
	})

	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}

	event := got[0]
	if event.ID != "eonet:EONET_123" {
		t.Fatalf("id = %q, want %q", event.ID, "eonet:EONET_123")
	}
	if event.Source != "eonet" {
		t.Fatalf("source = %q, want %q", event.Source, "eonet")
	}
	if event.Title != "Wildfire Test" {
		t.Fatalf("title = %q, want %q", event.Title, "Wildfire Test")
	}
	if event.Type != "wildfires" {
		t.Fatalf("type = %q, want %q", event.Type, "wildfires")
	}
	if event.StartedAt != "2026-05-24T12:00:00Z" {
		t.Fatalf("started_at = %q, want %q", event.StartedAt, "2026-05-24T12:00:00Z")
	}
	if event.UpdatedAt != "2026-05-24T12:00:00Z" {
		t.Fatalf("updated_at = %q, want %q", event.UpdatedAt, "2026-05-24T12:00:00Z")
	}
	if event.Longitude == nil || *event.Longitude != -121.5 {
		t.Fatalf("longitude = %v, want -121.5", event.Longitude)
	}
	if event.Latitude == nil || *event.Latitude != 38.25 {
		t.Fatalf("latitude = %v, want 38.25", event.Latitude)
	}
}

func TestNormalizeEventsSkipsEmptyGeometries(t *testing.T) {
	got := NormalizeEvents([]Event{
		{
			ID:         "EONET_EMPTY",
			Title:      "No Geometry",
			Categories: []Category{{ID: "volcanoes"}},
		},
	})

	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
}

func TestNormalizeEventsTakesFirstGeometry(t *testing.T) {
	got := NormalizeEvents([]Event{
		{
			ID:         "EONET_MULTI",
			Title:      "Multiple Geometry",
			Categories: []Category{{ID: "volcanoes"}},
			Geometries: []Geometry{
				{
					Date:        "2026-05-23T01:00:00Z",
					Coordinates: []float64{10, 20},
				},
				{
					Date:        "2026-05-24T01:00:00Z",
					Coordinates: []float64{30, 40},
				},
			},
		},
	})

	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].StartedAt != "2026-05-23T01:00:00Z" {
		t.Fatalf("started_at = %q, want first geometry date", got[0].StartedAt)
	}
	if got[0].Longitude == nil || *got[0].Longitude != 10 {
		t.Fatalf("longitude = %v, want first geometry longitude 10", got[0].Longitude)
	}
	if got[0].Latitude == nil || *got[0].Latitude != 20 {
		t.Fatalf("latitude = %v, want first geometry latitude 20", got[0].Latitude)
	}
}
