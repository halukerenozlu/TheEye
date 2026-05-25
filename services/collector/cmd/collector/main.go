package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	collector "theeye/services/collector"
	"theeye/services/collector/eonet"
	"theeye/services/collector/store"
	"theeye/services/collector/usgs"
)

const defaultInterval = 5 * time.Minute

type config struct {
	DatabaseURL  string
	USGSFeedURL  string
	EONETFeedURL string
	Interval     time.Duration
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() (config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return config{}, fmt.Errorf("DATABASE_URL is required")
	}

	interval := defaultInterval
	if rawInterval := os.Getenv("COLLECTOR_INTERVAL"); rawInterval != "" {
		parsedInterval, err := time.ParseDuration(rawInterval)
		if err != nil {
			return config{}, fmt.Errorf("COLLECTOR_INTERVAL must be a valid duration: %w", err)
		}
		if parsedInterval <= 0 {
			return config{}, fmt.Errorf("COLLECTOR_INTERVAL must be greater than zero")
		}
		interval = parsedInterval
	}

	return config{
		DatabaseURL:  databaseURL,
		USGSFeedURL:  os.Getenv("USGS_FEED_URL"),
		EONETFeedURL: os.Getenv("EONET_FEED_URL"),
		Interval:     interval,
	}, nil
}

func run(ctx context.Context, cfg config) error {
	log.Printf("collector starting with interval %s", cfg.Interval)

	runAndLog(ctx, cfg)

	ticker := time.NewTicker(cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Print("collector shutdown complete")
			return nil
		case <-ticker.C:
			runAndLog(ctx, cfg)
		}
	}
}

func runAndLog(ctx context.Context, cfg config) {
	if err := runOnce(ctx, cfg); err != nil {
		log.Printf("collector ingest failed: %v", err)
	}
}

func runOnce(parent context.Context, cfg config) error {
	ctx, cancel := context.WithTimeout(parent, 30*time.Second)
	defer cancel()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	eventStore := store.New(db)
	if err := eventStore.EnsureSchema(ctx); err != nil {
		return err
	}

	sources := []collector.Source{
		usgs.NewSource(cfg.USGSFeedURL, &http.Client{Timeout: 10 * time.Second}),
		eonet.NewSource(cfg.EONETFeedURL, &http.Client{Timeout: 10 * time.Second}),
	}

	for _, src := range sources {
		events, err := src.Fetch(ctx)
		if err != nil {
			return fmt.Errorf("%s ingest fetch: %w", src.Name(), err)
		}

		written, err := eventStore.UpsertNormalizedEvents(ctx, events)
		if err != nil {
			return fmt.Errorf("%s ingest store: %w", src.Name(), err)
		}

		log.Printf("ingested %d/%d %s events", written, len(events), src.Name())
	}

	return nil
}
