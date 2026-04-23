package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
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

func developmentTestConfig() config {
	return config{
		Port:        "8080",
		AppName:     "theeye-api",
		Env:         "development",
		Version:     "test",
		DatabaseURL: "",
	}
}

type fakeEventsReader struct {
	items []Event
	err   error

	eventByID map[string]Event
	detailErr error

	lastListQuery listEventsQuery
}

func (f *fakeEventsReader) ListEvents(_ context.Context, query listEventsQuery) (eventsListResponse, error) {
	f.lastListQuery = query

	if f.err != nil {
		return eventsListResponse{}, f.err
	}

	filtered := make([]Event, 0, len(f.items))
	for _, item := range f.items {
		if query.Type != "" && item.Type != query.Type {
			continue
		}

		startedAt, err := time.Parse(time.RFC3339, item.StartedAt)
		if err != nil {
			return eventsListResponse{}, err
		}

		if query.StartedAfter != nil && startedAt.Before(*query.StartedAfter) {
			continue
		}

		if query.StartedBefore != nil && startedAt.After(*query.StartedBefore) {
			continue
		}

		filtered = append(filtered, item)
	}

	sort.Slice(filtered, func(i, j int) bool {
		leftUpdatedAt, err := time.Parse(time.RFC3339, filtered[i].UpdatedAt)
		if err != nil {
			return false
		}
		rightUpdatedAt, err := time.Parse(time.RFC3339, filtered[j].UpdatedAt)
		if err != nil {
			return false
		}

		if leftUpdatedAt.Equal(rightUpdatedAt) {
			return filtered[i].ID < filtered[j].ID
		}

		if query.Sort == listEventsSortUpdatedAtAsc {
			return leftUpdatedAt.Before(rightUpdatedAt)
		}

		return leftUpdatedAt.After(rightUpdatedAt)
	})

	if query.Limit == nil {
		return eventsListResponse{
			Items:      filtered,
			NextCursor: "",
		}, nil
	}

	if query.Cursor >= len(filtered) {
		return eventsListResponse{
			Items:      []Event{},
			NextCursor: "",
		}, nil
	}

	end := query.Cursor + *query.Limit
	nextCursor := ""
	if end < len(filtered) {
		nextCursor = strconv.Itoa(end)
	} else {
		end = len(filtered)
	}

	return eventsListResponse{
		Items:      filtered[query.Cursor:end],
		NextCursor: nextCursor,
	}, nil
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
	if reader.lastListQuery.Type != "earthquake" {
		t.Fatalf("lastListQuery.Type = %q, want %q", reader.lastListQuery.Type, "earthquake")
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

func TestEventsListSortingUpdatedAtAsc(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:2",
				Type:      "earthquake",
				Title:     "Newer",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:30Z",
				UpdatedAt: "2023-11-14T22:13:30Z",
			},
			{
				ID:        "usgs:1",
				Type:      "earthquake",
				Title:     "Older",
				Status:    "confirmed",
				Severity:  1,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:20Z",
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?sort=updated_at_asc", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got eventsListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal events list response: %v", err)
	}

	if len(got.Items) != 2 {
		t.Fatalf("items length = %d, want 2", len(got.Items))
	}
	if got.Items[0].ID != "usgs:1" {
		t.Fatalf("items[0].id = %q, want %q", got.Items[0].ID, "usgs:1")
	}
	if got.Items[1].ID != "usgs:2" {
		t.Fatalf("items[1].id = %q, want %q", got.Items[1].ID, "usgs:2")
	}
}

func TestEventsListInvalidSortValue(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?sort=severity_desc", nil)
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
	if got.Message != "invalid query parameter: sort must be one of updated_at_desc,updated_at_asc" {
		t.Fatalf("message = %q, want %q", got.Message, "invalid query parameter: sort must be one of updated_at_desc,updated_at_asc")
	}
}

func TestEventsListPaginationLimitAndNextCursor(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:3",
				Type:      "earthquake",
				Title:     "Third",
				Status:    "confirmed",
				Severity:  3,
				StartedAt: "2023-11-14T22:13:40Z",
				UpdatedAt: "2023-11-14T22:13:40Z",
			},
			{
				ID:        "usgs:2",
				Type:      "earthquake",
				Title:     "Second",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:30Z",
				UpdatedAt: "2023-11-14T22:13:30Z",
			},
			{
				ID:        "usgs:1",
				Type:      "earthquake",
				Title:     "First",
				Status:    "confirmed",
				Severity:  1,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:20Z",
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?limit=2", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got eventsListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal events list response: %v", err)
	}

	if len(got.Items) != 2 {
		t.Fatalf("items length = %d, want 2", len(got.Items))
	}
	if got.NextCursor != "2" {
		t.Fatalf("next_cursor = %q, want %q", got.NextCursor, "2")
	}
}

func TestEventsListPaginationWithCursor(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{
		items: []Event{
			{
				ID:        "usgs:3",
				Type:      "earthquake",
				Title:     "Third",
				Status:    "confirmed",
				Severity:  3,
				StartedAt: "2023-11-14T22:13:40Z",
				UpdatedAt: "2023-11-14T22:13:40Z",
			},
			{
				ID:        "usgs:2",
				Type:      "earthquake",
				Title:     "Second",
				Status:    "confirmed",
				Severity:  2,
				StartedAt: "2023-11-14T22:13:30Z",
				UpdatedAt: "2023-11-14T22:13:30Z",
			},
			{
				ID:        "usgs:1",
				Type:      "earthquake",
				Title:     "First",
				Status:    "confirmed",
				Severity:  1,
				StartedAt: "2023-11-14T22:13:20Z",
				UpdatedAt: "2023-11-14T22:13:20Z",
			},
		},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?limit=2&cursor=2", nil)
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
	if got.Items[0].ID != "usgs:1" {
		t.Fatalf("items[0].id = %q, want %q", got.Items[0].ID, "usgs:1")
	}
	if got.NextCursor != "" {
		t.Fatalf("next_cursor = %q, want empty string", got.NextCursor)
	}
}

func TestEventsListInvalidCursorWithoutLimit(t *testing.T) {
	r := newRouterWithEventsReader(testConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(http.MethodGet, "/v1/events?cursor=2", nil)
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
	if got.Message != "invalid query parameter: cursor requires limit" {
		t.Fatalf("message = %q, want %q", got.Message, "invalid query parameter: cursor requires limit")
	}
}

func TestEventsListIncludesGeometryWhenAvailable(t *testing.T) {
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
				Geometry: &EventGeometry{
					Longitude: -117.5,
					Latitude:  35.7,
				},
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
	if got.Items[0].Geometry == nil {
		t.Fatal("items[0].geometry = nil, want non-nil")
	}
	if got.Items[0].Geometry.Longitude != -117.5 {
		t.Fatalf("items[0].geometry.longitude = %v, want %v", got.Items[0].Geometry.Longitude, -117.5)
	}
	if got.Items[0].Geometry.Latitude != 35.7 {
		t.Fatalf("items[0].geometry.latitude = %v, want %v", got.Items[0].Geometry.Latitude, 35.7)
	}
}

func TestEventsDetailIncludesGeometryWhenAvailable(t *testing.T) {
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
				Geometry: &EventGeometry{
					Longitude: -117.5,
					Latitude:  35.7,
				},
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

	if got.Geometry == nil {
		t.Fatal("geometry = nil, want non-nil")
	}
	if got.Geometry.Longitude != -117.5 {
		t.Fatalf("geometry.longitude = %v, want %v", got.Geometry.Longitude, -117.5)
	}
	if got.Geometry.Latitude != 35.7 {
		t.Fatalf("geometry.latitude = %v, want %v", got.Geometry.Latitude, 35.7)
	}
}

func TestLocalDevelopmentCORSAllowsDashboardOrigin(t *testing.T) {
	r := newRouterWithEventsReader(developmentTestConfig(), &fakeEventsReader{
		items: []Event{},
	})
	req := httptest.NewRequest(http.MethodGet, "/v1/events", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
		t.Fatalf("Access-Control-Allow-Origin = %q, want %q", got, "http://localhost:3000")
	}
}

func TestLocalDevelopmentCORSPreflight(t *testing.T) {
	r := newRouterWithEventsReader(developmentTestConfig(), &fakeEventsReader{})
	req := httptest.NewRequest(http.MethodOptions, "/v1/events", nil)
	req.Header.Set("Origin", "http://localhost:3000")
	req.Header.Set("Access-Control-Request-Method", "GET")
	req.Header.Set("Access-Control-Request-Headers", "Content-Type")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
		t.Fatalf("Access-Control-Allow-Origin = %q, want %q", got, "http://localhost:3000")
	}
	if got := rec.Header().Get("Access-Control-Allow-Methods"); got != "GET, OPTIONS" {
		t.Fatalf("Access-Control-Allow-Methods = %q, want %q", got, "GET, OPTIONS")
	}
}
