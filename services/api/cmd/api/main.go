package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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

func main() {
	cfg := loadConfig()

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

	addr := ":" + cfg.Port
	log.Printf("api listening on %s", addr)
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
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
