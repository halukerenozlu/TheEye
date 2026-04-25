package main

import (
	"testing"
	"time"
)

func TestLoadConfigDefaultsInterval(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("COLLECTOR_INTERVAL", "")
	t.Setenv("USGS_FEED_URL", "")

	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("loadConfig returned error: %v", err)
	}

	if cfg.Interval != defaultInterval {
		t.Fatalf("Interval = %s, want %s", cfg.Interval, defaultInterval)
	}
}

func TestLoadConfigParsesInterval(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("COLLECTOR_INTERVAL", "30s")

	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("loadConfig returned error: %v", err)
	}

	if cfg.Interval != 30*time.Second {
		t.Fatalf("Interval = %s, want 30s", cfg.Interval)
	}
}

func TestLoadConfigRejectsMissingDatabaseURL(t *testing.T) {
	t.Setenv("DATABASE_URL", "")

	if _, err := loadConfig(); err == nil {
		t.Fatal("loadConfig error = nil, want non-nil")
	}
}
