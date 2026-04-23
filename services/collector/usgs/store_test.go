package usgs

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
)

type execCall struct {
	query string
	args  []any
}

type fakeExecer struct {
	calls  []execCall
	failAt int
	err    error
}

func (f *fakeExecer) ExecContext(_ context.Context, query string, args ...any) (sql.Result, error) {
	f.calls = append(f.calls, execCall{query: query, args: args})

	if f.err != nil && len(f.calls) == f.failAt {
		return fakeResult(0), f.err
	}

	return fakeResult(1), nil
}

type fakeResult int64

func (r fakeResult) LastInsertId() (int64, error) {
	return int64(r), nil
}

func (r fakeResult) RowsAffected() (int64, error) {
	return int64(r), nil
}

func TestEnsureSchemaSuccess(t *testing.T) {
	exec := &fakeExecer{}
	store := NewStore(exec)

	if err := store.EnsureSchema(context.Background()); err != nil {
		t.Fatalf("EnsureSchema returned error: %v", err)
	}

	if len(exec.calls) != 1 {
		t.Fatalf("exec call count = %d, want 1", len(exec.calls))
	}
	if !strings.Contains(exec.calls[0].query, "CREATE TABLE IF NOT EXISTS ingested_events") {
		t.Fatalf("schema query mismatch: %q", exec.calls[0].query)
	}
}

func TestUpsertNormalizedEventsSuccess(t *testing.T) {
	exec := &fakeExecer{}
	store := NewStore(exec)
	events := []NormalizedEvent{
		{
			ID:        "usgs:abc123",
			Type:      "earthquake",
			Title:     "M 3.4 - Test",
			Status:    "confirmed",
			Severity:  2,
			StartedAt: "2023-11-14T22:13:20Z",
			UpdatedAt: "2023-11-14T22:13:21Z",
		},
	}

	written, err := store.UpsertNormalizedEvents(context.Background(), events)
	if err != nil {
		t.Fatalf("UpsertNormalizedEvents returned error: %v", err)
	}

	if written != 1 {
		t.Fatalf("written = %d, want 1", written)
	}
	if len(exec.calls) != 1 {
		t.Fatalf("exec call count = %d, want 1", len(exec.calls))
	}
	if !strings.Contains(exec.calls[0].query, "ON CONFLICT (source_name, source_event_id)") {
		t.Fatalf("upsert query missing conflict clause: %q", exec.calls[0].query)
	}

	args := exec.calls[0].args
	if args[0] != "usgs" {
		t.Fatalf("source_name arg = %v, want usgs", args[0])
	}
	if args[1] != "abc123" {
		t.Fatalf("source_event_id arg = %v, want abc123", args[1])
	}
}

func TestUpsertNormalizedEventsDuplicateSafeInBatch(t *testing.T) {
	exec := &fakeExecer{}
	store := NewStore(exec)
	dupe := NormalizedEvent{
		ID:        "usgs:dup-1",
		Type:      "earthquake",
		Title:     "duplicate",
		Status:    "confirmed",
		Severity:  1,
		StartedAt: "2023-11-14T22:13:20Z",
		UpdatedAt: "2023-11-14T22:13:21Z",
	}

	written, err := store.UpsertNormalizedEvents(context.Background(), []NormalizedEvent{dupe, dupe})
	if err != nil {
		t.Fatalf("UpsertNormalizedEvents returned error: %v", err)
	}

	if written != 1 {
		t.Fatalf("written = %d, want 1", written)
	}
	if len(exec.calls) != 1 {
		t.Fatalf("exec call count = %d, want 1", len(exec.calls))
	}
}

func TestUpsertNormalizedEventsInvalidID(t *testing.T) {
	exec := &fakeExecer{}
	store := NewStore(exec)
	events := []NormalizedEvent{
		{
			ID:        "bad-id",
			Type:      "earthquake",
			Title:     "M 1.0 - Bad",
			Status:    "unknown",
			Severity:  1,
			StartedAt: "2023-11-14T22:13:20Z",
			UpdatedAt: "2023-11-14T22:13:21Z",
		},
	}

	_, err := store.UpsertNormalizedEvents(context.Background(), events)
	if err == nil {
		t.Fatal("UpsertNormalizedEvents error = nil, want non-nil")
	}
	if !strings.Contains(err.Error(), "must start with") {
		t.Fatalf("error = %q, want contains %q", err.Error(), "must start with")
	}
	if len(exec.calls) != 0 {
		t.Fatalf("exec call count = %d, want 0", len(exec.calls))
	}
}

func TestUpsertNormalizedEventsDBError(t *testing.T) {
	exec := &fakeExecer{
		failAt: 1,
		err:    errors.New("db unavailable"),
	}
	store := NewStore(exec)
	events := []NormalizedEvent{
		{
			ID:        "usgs:db-1",
			Type:      "earthquake",
			Title:     "M 1.1 - DB",
			Status:    "preliminary",
			Severity:  1,
			StartedAt: "2023-11-14T22:13:20Z",
			UpdatedAt: "2023-11-14T22:13:21Z",
		},
	}

	_, err := store.UpsertNormalizedEvents(context.Background(), events)
	if err == nil {
		t.Fatal("UpsertNormalizedEvents error = nil, want non-nil")
	}
	if !strings.Contains(err.Error(), "upsert normalized event") {
		t.Fatalf("error = %q, want contains %q", err.Error(), "upsert normalized event")
	}
}
