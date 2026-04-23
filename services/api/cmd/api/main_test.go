package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

	lastListFilters listEventsFilters
}

func (f *fakeEventsReader) ListEvents(_ context.Context, filters listEventsFilters) ([]Event, error) {
	f.lastListFilters = filters

	if f.err != nil {
		return nil, f.err
	}

	items := make([]Event, 0, len(f.items))
	for _, item := range f.items {
		if filters.Type != "" && item.Type != filters.Type {
			continue
		}

		startedAt, err := time.Parse(time.RFC3339, item.StartedAt)
		if err != nil {
			return nil, err
		}

		if filters.StartedAfter != nil && startedAt.Before(*filters.StartedAfter) {
			continue
		}

		if filters.StartedBefore != nil && startedAt.After(*filters.StartedBefore) {
			continue
		}

		items = append(items, item)
	}

	return items, nil
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
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
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
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
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
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
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
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
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

func TestEventsListFilteringByType(t *testing.T) {
	reader := &fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:1",
				Type:      "earthquake",
				Title:     "EQ 1",
				Status:    "confirmed",
				Severity:  1,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:21Z",
			},
			{
				ID:        "fire:1",
				Type:      "wildfire",
				Title:     "WF 1",
				Status:    "active",
				Severity:  3,
				StartedAt: "2023-11-14T22:13:25Z",
				UpdatedAt: "2023-11-14T22:13:26Z",
			},
		},
	}
	r := newRouterWithEventsReader(testConfig(), reader)
	req := httptest.NewRequest(http.MethodGet, "/v1/events?type=earthquake", nil)
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
	if got.Items[0].Type != "earthquake" {
		t.Fatalf("items[0].type = %q, want %q", got.Items[0].Type, "earthquake")
	}
	if reader.lastListFilters.Type != "earthquake" {
		t.Fatalf("lastListFilters.Type = %q, want %q", reader.lastListFilters.Type, "earthquake")
	}
}

func TestEventsListFilteringByStartedAtRange(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:1",
				Type:      "earthquake",
				Title:     "Old",
				Status:    "confirmed",
				Severity:  1,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:20Z",
			},
			{
				ID:        "usgs:2",
				Type:      "earthquake",
				Title:     "InRange",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:30Z",
				UpdatedAt: "2023-11-14T22:13:30Z",
			},
			{
				ID:        "usgs:3",
				Type:      "earthquake",
				Title:     "New",
				Status:    "confirmed",
				Severity:  3,
				StartedAt: "2023-11-14T22:13:40Z",
				UpdatedAt: "2023-11-14T22:13:40Z",
			},
		},
	})
	req := httptest.NewRequest(
		http.MethodGet,
		"/v1/events?started_after=2023-11-14T22:13:25Z&started_before=2023-11-14T22:13:35Z",
		nil,
	)
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
	if got.Items[0].ID != "usgs:2" {
		t.Fatalf("items[0].id = %q, want %q", got.Items[0].ID, "usgs:2")
	}
}

func TestEventsListInvalidQueryParameter(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?severity=2", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	if got.Error != "bad_request" {
		t.Fatalf("error code = %q, want %q", got.Error, "bad_request")
	}
	if got.Message != "invalid query parameter: severity" {
		t.Fatalf("message = %q, want %q", got.Message, "invalid query parameter: severity")
	}
}

func TestEventsListInvalidStartedAfterFormat(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?started_after=not-a-time", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	if got.Error != "bad_request" {
		t.Fatalf("error code = %q, want %q", got.Error, "bad_request")
	}
	if got.Message != "invalid query parameter: started_after must be RFC3339" {
		t.Fatalf("message = %q, want %q", got.Message, "invalid query parameter: started_after must be RFC3339")
	}
}

func TestEventsListInvalidStartedAtRange(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(
		http.MethodGet,
		"/v1/events?started_after=2023-11-14T22:13:35Z&started_before=2023-11-14T22:13:25Z",
		nil,
	)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var got errorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal error response: %v", err)
	}

	if got.Error != "bad_request" {
		t.Fatalf("error code = %q, want %q", got.Error, "bad_request")
	}
	if got.Message != "invalid query parameter: started_after must be before or equal to started_before" {
		t.Fatalf("message = %q, want %q", got.Message, "invalid query parameter: started_after must be before or equal to started_before")
	}
}
