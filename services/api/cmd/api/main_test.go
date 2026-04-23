package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testConfig() config {
	return config{
		Port:        "8080",
		AppName:     "theeye-api",
		Env:         "test",
		Version:     "test",
		DatabaseURL: "",
	}
}

type fakeEventsReader struct {
	items []Event
	err   error

	eventByID map[string]Event
	detailErr error
}

func (f fakeEventsReader) ListEvents(_ context.Context) ([]Event, error) {
	return f.items, f.err
}

func (f fakeEventsReader) GetEventByID(_ context.Context, id string) (Event, error) {
	if f.detailErr != nil {
		return Event{}, f.detailErr
	}

	ev, ok := f.eventByID[id]
	if !ok {
		return Event{}, errEventNotFound
	}

	return ev, nil
}

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()

	writeError(rec, http.StatusBadRequest, "bad_request", "bad request")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	want := errorResponse{
		Error:   "bad_request",
		Message: "bad request",
	}
	if got != want {
		t.Fatalf("body = %+v, want %+v", got, want)
	}
}

func TestEventsDetailNotFoundErrorShape(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), fakeEventsReader{})
	req := httptest.NewRequest(http.MethodGet, "/v1/events/abc", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	want := errorResponse{
		Error:   "event_not_found",
		Message: "event not found",
	}
	if got != want {
		t.Fatalf("body = %+v, want %+v", got, want)
	}
}

func TestEventsDetailReturnsStoredData(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), fakeEventsReader{
		eventByID: map[string]Event{
			"usgs:abc123": {
				ID:        "usgs:abc123",
				Type:      "earthquake",
				Title:     "M 3.4 - Test",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:21Z",
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events/usgs:abc123", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got Event
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal event detail response: %v", err)
	}

	if got.ID != "usgs:abc123" {
		t.Fatalf("id = %q, want %q", got.ID, "usgs:abc123")
	}
	if got.Type != "earthquake" {
		t.Fatalf("type = %q, want %q", got.Type, "earthquake")
	}
}

func TestRouterNotFoundErrorShape(t *testing.T) {
	r := newRouter(testConfig())
	req := httptest.NewRequest(http.MethodGet, "/v1/unknown", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	want := errorResponse{
		Error:   "route_not_found",
		Message: "route not found",
	}
	if got != want {
		t.Fatalf("body = %+v, want %+v", got, want)
	}
}

func TestRouterMethodNotAllowedErrorShape(t *testing.T) {
	r := newRouter(testConfig())
	req := httptest.NewRequest(http.MethodPost, "/v1/events", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	want := errorResponse{
		Error:   "method_not_allowed",
		Message: "method not allowed",
	}
	if got != want {
		t.Fatalf("body = %+v, want %+v", got, want)
	}
}

func TestEventsListShapeRemainsStable(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), fakeEventsReader{
		items: []Event{},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got eventsListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal events list response: %v", err)
	}

	if len(got.Items) != 0 {
		t.Fatalf("items length = %d, want 0", len(got.Items))
	}
	if got.NextCursor != "" {
		t.Fatalf("next_cursor = %q, want empty string", got.NextCursor)
	}
}

func TestEventsListReturnsStoredData(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:abc123",
				Type:      "earthquake",
				Title:     "M 3.4 - Test",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:21Z",
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got eventsListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal events list response: %v", err)
	}

	if len(got.Items) != 1 {
		t.Fatalf("items length = %d, want 1", len(got.Items))
	}
	if got.Items[0].ID != "usgs:abc123" {
		t.Fatalf("items[0].id = %q, want %q", got.Items[0].ID, "usgs:abc123")
	}
	if got.NextCursor != "" {
		t.Fatalf("next_cursor = %q, want empty string", got.NextCursor)
	}
}
