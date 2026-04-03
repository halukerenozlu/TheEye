package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	Port    string
	AppName string
	Env     string
	Version string
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

func main() {
	cfg := loadConfig()
	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

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

	r.Get("/v1/events", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, eventsListResponse{
			Items:      []Event{},
			NextCursor: "",
		})
	})

	r.Get("/v1/events/{id}", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusNotFound, errorResponse{
			Error:   "event_not_found",
			Message: "event not found",
		})
	})

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

func loadConfig() config {
	return config{
		Port:    getenv("PORT", "8080"),
		AppName: getenv("APP_NAME", "theeye-api"),
		Env:     getenv("APP_ENV", "development"),
		Version: getenv("APP_VERSION", "dev"),
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
