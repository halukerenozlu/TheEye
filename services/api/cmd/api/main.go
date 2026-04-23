package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lib/pq"
)

type config struct {
	Port        string
	AppName     string
	Env         string
	Version     string
	DatabaseURL string
}

type statusResponse struct {
	Status string `json:"status"`
}

type metaResponse struct {
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Severity  int    `json:"severity"`
	StartedAt string `json:"started_at"`
	UpdatedAt string `json:"updated_at"`
}

type eventsListResponse struct {
	Items      []Event `json:"items"`
	NextCursor string  `json:"next_cursor"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var errEventNotFound = errors.New("event not found")

type eventsReader interface {
	ListEvents(ctx context.Context, filters listEventsFilters) ([]Event, error)
	GetEventByID(ctx context.Context, id string) (Event, error)
}

type emptyEventsReader struct{}

func (emptyEventsReader) ListEvents(_ context.Context, _ listEventsFilters) ([]Event, error) {
	return []Event{}, nil
}

func (emptyEventsReader) GetEventByID(_ context.Context, _ string) (Event, error) {
	return Event{}, errEventNotFound
}

type postgresEventsReader struct {
	db *sql.DB
}

type listEventsFilters struct {
	Type          string
	StartedAfter  *time.Time
	StartedBefore *time.Time
}

var supportedEventsListQueryParams = map[string]struct{}{
	"type":           {},
	"started_after":  {},
	"started_before": {},
}

func newPostgresEventsReader(db *sql.DB) *postgresEventsReader {
	return &postgresEventsReader{db: db}
}

func (r *postgresEventsReader) ListEvents(ctx context.Context, filters listEventsFilters) ([]Event, error) {
	q := `
SELECT id, type, title, status, severity, started_at, updated_at
FROM ingested_events
`
	args := make([]any, 0, 3)
	where := make([]string, 0, 3)
	nextPlaceholder := 1

	if filters.Type != "" {
		where = append(where, "type = $"+strconv.Itoa(nextPlaceholder))
		args = append(args, filters.Type)
		nextPlaceholder++
	}

	if filters.StartedAfter != nil {
		where = append(where, "started_at >= $"+strconv.Itoa(nextPlaceholder))
		args = append(args, *filters.StartedAfter)
		nextPlaceholder++
	}

	if filters.StartedBefore != nil {
		where = append(where, "started_at <= $"+strconv.Itoa(nextPlaceholder))
		args = append(args, *filters.StartedBefore)
		nextPlaceholder++
	}

	if len(where) > 0 {
		q += "WHERE " + strings.Join(where, " AND ") + "\n"
	}

	q += "ORDER BY updated_at DESC, id ASC;"

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "42P01" {
			return []Event{}, nil
		}
		return nil, fmt.Errorf("query ingested events: %w", err)
	}
	defer rows.Close()

	events := make([]Event, 0)
	for rows.Next() {
		var (
			ev        Event
			startedAt time.Time
			updatedAt time.Time
		)

		if err := rows.Scan(
			&ev.ID,
			&ev.Type,
			&ev.Title,
			&ev.Status,
			&ev.Severity,
			&startedAt,
			&updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan ingested event: %w", err)
		}

		ev.StartedAt = startedAt.UTC().Format(time.RFC3339)
		ev.UpdatedAt = updatedAt.UTC().Format(time.RFC3339)
		events = append(events, ev)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate ingested events: %w", err)
	}

	return events, nil
}

func (r *postgresEventsReader) GetEventByID(ctx context.Context, id string) (Event, error) {
	const q = `
SELECT id, type, title, status, severity, started_at, updated_at
FROM ingested_events
WHERE id = $1
LIMIT 1;
`

	var (
		ev        Event
		startedAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&ev.ID,
		&ev.Type,
		&ev.Title,
		&ev.Status,
		&ev.Severity,
		&startedAt,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, errEventNotFound
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "42P01" {
			return Event{}, errEventNotFound
		}

		return Event{}, fmt.Errorf("query ingested event detail: %w", err)
	}

	ev.StartedAt = startedAt.UTC().Format(time.RFC3339)
	ev.UpdatedAt = updatedAt.UTC().Format(time.RFC3339)

	return ev, nil
}

func main() {
	cfg := loadConfig()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {
	eventsReader := eventsReader(emptyEventsReader{})
	var db *sql.DB
	if cfg.DatabaseURL != "" {
		var err error
		db, err = openDB(cfg.DatabaseURL)
		if err != nil {
			return err
		}
		defer db.Close()

		eventsReader = newPostgresEventsReader(db)
	}

	r := newRouterWithEventsReader(cfg, eventsReader)

	addr := ":" + cfg.Port
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		log.Printf("api listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		log.Print("shutdown signal received")
	case err := <-errCh:
		return err
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	if err := <-errCh; err != nil {
		return err
	}

	log.Print("api shutdown complete")
	return nil
}

func newRouter(cfg config) chi.Router {
	return newRouterWithEventsReader(cfg, emptyEventsReader{})
}

func newRouterWithEventsReader(cfg config, reader eventsReader) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		writeError(w, http.StatusNotFound, "route_not_found", "route not found")
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, _ *http.Request) {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
	})

	// Health endpoints
	r.Get("/v1/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, statusResponse{Status: "healthy"})
	})

	// Placeholder ready check (later: DB + Redis ping).
	r.Get("/v1/readyz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, statusResponse{Status: "ready"})
	})

	r.Get("/v1/meta", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, metaResponse{
			Name:        cfg.AppName,
			Environment: cfg.Env,
			Version:     cfg.Version,
		})
	})

	r.Get("/v1/events", func(w http.ResponseWriter, req *http.Request) {
		filters, err := parseListEventsFilters(req.URL.Query())
		if err != nil {
			writeError(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		items, err := reader.ListEvents(req.Context(), filters)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "events_read_failed", "failed to read events")
			return
		}

		writeJSON(w, http.StatusOK, eventsListResponse{
			Items:      items,
			NextCursor: "",
		})
	})

	r.Get("/v1/events/{id}", func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")

		event, err := reader.GetEventByID(req.Context(), id)
		if err != nil {
			if errors.Is(err, errEventNotFound) {
				writeError(w, http.StatusNotFound, "event_not_found", "event not found")
				return
			}

			writeError(w, http.StatusInternalServerError, "event_read_failed", "failed to read event")
			return
		}

		writeJSON(w, http.StatusOK, event)
	})

	return r
}

func loadConfig() config {
	return config{
		Port:        getenv("PORT", "8080"),
		AppName:     getenv("APP_NAME", "theeye-api"),
		Env:         getenv("APP_ENV", "development"),
		Version:     getenv("APP_VERSION", "dev"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}

func openDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorResponse{
		Error:   code,
		Message: message,
	})
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseListEventsFilters(values url.Values) (listEventsFilters, error) {
	var filters listEventsFilters

	for key := range values {
		if _, ok := supportedEventsListQueryParams[key]; !ok {
			return listEventsFilters{}, fmt.Errorf("invalid query parameter: %s", key)
		}
	}

	if rawType := strings.TrimSpace(values.Get("type")); rawType != "" {
		filters.Type = rawType
	}

	if rawStartedAfter := strings.TrimSpace(values.Get("started_after")); rawStartedAfter != "" {
		startedAfter, err := time.Parse(time.RFC3339, rawStartedAfter)
		if err != nil {
			return listEventsFilters{}, errors.New("invalid query parameter: started_after must be RFC3339")
		}
		startedAfter = startedAfter.UTC()
		filters.StartedAfter = &startedAfter
	}

	if rawStartedBefore := strings.TrimSpace(values.Get("started_before")); rawStartedBefore != "" {
		startedBefore, err := time.Parse(time.RFC3339, rawStartedBefore)
		if err != nil {
			return listEventsFilters{}, errors.New("invalid query parameter: started_before must be RFC3339")
		}
		startedBefore = startedBefore.UTC()
		filters.StartedBefore = &startedBefore
	}

	if filters.StartedAfter != nil && filters.StartedBefore != nil && filters.StartedAfter.After(*filters.StartedBefore) {
		return listEventsFilters{}, errors.New("invalid query parameter: started_after must be before or equal to started_before")
	}

	return filters, nil
}
