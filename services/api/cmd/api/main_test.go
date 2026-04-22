package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testConfig() config {
	return config{
		Port:    "8080",
		AppName: "theeye-api",
		Env:     "test",
		Version: "test",
	}
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

func TestEventsDetailPlaceholderErrorShape(t *testing.T) {
	r := newRouter(testConfig())
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
	r := newRouter(testConfig())
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
