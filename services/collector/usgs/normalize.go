package usgs

import (
	"fmt"
	"strings"
	"time"
)

type NormalizedEvent struct {
	ID        string
	Type      string
	Title     string
	Status    string
	Severity  int
	StartedAt string
	UpdatedAt string
	Longitude *float64
	Latitude  *float64
}

func NormalizeFeatures(features []Feature) []NormalizedEvent {
	out := make([]NormalizedEvent, 0, len(features))
	for _, f := range features {
		out = append(out, NormalizeFeature(f))
	}
	return out
}

func NormalizeFeature(feature Feature) NormalizedEvent {
	longitude, latitude := normalizeCoordinates(feature.Geometry.Coordinates)

	return NormalizedEvent{
		ID:        fmt.Sprintf("usgs:%s", feature.ID),
		Type:      "earthquake",
		Title:     feature.Properties.Title,
		Status:    mapStatus(feature.Properties.Status),
		Severity:  mapSeverity(feature.Properties.Mag),
		StartedAt: formatUnixMilli(feature.Properties.Time),
		UpdatedAt: formatUnixMilli(feature.Properties.Updated),
		Longitude: longitude,
		Latitude:  latitude,
	}
}

func normalizeCoordinates(coords []float64) (*float64, *float64) {
	if len(coords) < 2 {
		return nil, nil
	}

	lon := coords[0]
	lat := coords[1]

	return &lon, &lat
}

func mapStatus(sourceStatus string) string {
	switch strings.ToLower(strings.TrimSpace(sourceStatus)) {
	case "reviewed":
		return "confirmed"
	case "automatic":
		return "preliminary"
	case "deleted":
		return "removed"
	default:
		return "unknown"
	}
}

func mapSeverity(mag *float64) int {
	if mag == nil {
		return 0
	}

	switch {
	case *mag >= 7.0:
		return 5
	case *mag >= 6.0:
		return 4
	case *mag >= 4.0:
		return 3
	case *mag >= 2.0:
		return 2
	default:
		return 1
	}
}

func formatUnixMilli(ms int64) string {
	if ms <= 0 {
		return ""
	}

	return time.UnixMilli(ms).UTC().Format(time.RFC3339)
}
