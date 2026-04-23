package usgs

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type execer interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type Store struct {
	db execer
}

func NewStore(db execer) *Store {
	return &Store{db: db}
}

func (s *Store) EnsureSchema(ctx context.Context) error {
	const ddl = `
CREATE TABLE IF NOT EXISTS ingested_events (
  source_name TEXT NOT NULL,
  source_event_id TEXT NOT NULL,
  id TEXT NOT NULL,
  type TEXT NOT NULL,
  title TEXT NOT NULL,
  status TEXT NOT NULL,
  severity INTEGER NOT NULL,
  started_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  ingested_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (source_name, source_event_id)
);`

	if _, err := s.db.ExecContext(ctx, ddl); err != nil {
		return fmt.Errorf("ensure ingested_events schema: %w", err)
	}

	return nil
}

func (s *Store) UpsertNormalizedEvents(ctx context.Context, events []NormalizedEvent) (int, error) {
	if len(events) == 0 {
		return 0, nil
	}

	unique := deduplicateNormalizedEvents(events)
	const upsert = `
INSERT INTO ingested_events (
  source_name,
  source_event_id,
  id,
  type,
  title,
  status,
  severity,
  started_at,
  updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (source_name, source_event_id) DO UPDATE SET
  id = EXCLUDED.id,
  type = EXCLUDED.type,
  title = EXCLUDED.title,
  status = EXCLUDED.status,
  severity = EXCLUDED.severity,
  started_at = EXCLUDED.started_at,
  updated_at = EXCLUDED.updated_at,
  ingested_at = NOW();`

	for _, event := range unique {
		sourceEventID, err := sourceEventIDFromNormalizedID(event.ID)
		if err != nil {
			return 0, err
		}

		startedAt, err := parseRFC3339(event.StartedAt)
		if err != nil {
			return 0, fmt.Errorf("parse started_at for %q: %w", event.ID, err)
		}

		updatedAt, err := parseRFC3339(event.UpdatedAt)
		if err != nil {
			return 0, fmt.Errorf("parse updated_at for %q: %w", event.ID, err)
		}

		if _, err := s.db.ExecContext(
			ctx,
			upsert,
			"usgs",
			sourceEventID,
			event.ID,
			event.Type,
			event.Title,
			event.Status,
			event.Severity,
			startedAt,
			updatedAt,
		); err != nil {
			return 0, fmt.Errorf("upsert normalized event %q: %w", event.ID, err)
		}
	}

	return len(unique), nil
}

func deduplicateNormalizedEvents(events []NormalizedEvent) []NormalizedEvent {
	unique := make([]NormalizedEvent, 0, len(events))
	seen := make(map[string]struct{}, len(events))

	for _, event := range events {
		if _, ok := seen[event.ID]; ok {
			continue
		}

		seen[event.ID] = struct{}{}
		unique = append(unique, event)
	}

	return unique
}

func sourceEventIDFromNormalizedID(normalizedID string) (string, error) {
	const prefix = "usgs:"
	if !strings.HasPrefix(normalizedID, prefix) {
		return "", fmt.Errorf("normalized id %q must start with %q", normalizedID, prefix)
	}

	sourceID := strings.TrimPrefix(normalizedID, prefix)
	if sourceID == "" {
		return "", fmt.Errorf("normalized id %q has empty source id", normalizedID)
	}

	return sourceID, nil
}

func parseRFC3339(v string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}
