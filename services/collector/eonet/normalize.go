package eonet

import (
	"fmt"

	"theeye/services/collector/models"
)

func NormalizeEvents(events []Event) []models.NormalizedEvent {
	out := make([]models.NormalizedEvent, 0, len(events))
	for _, event := range events {
		normalized, ok := NormalizeEvent(event)
		if !ok {
			continue
		}

		out = append(out, normalized)
	}
	return out
}

func NormalizeEvent(event Event) (models.NormalizedEvent, bool) {
	if len(event.Geometries) == 0 {
		return models.NormalizedEvent{}, false
	}

	geometry := event.Geometries[0]
	if len(geometry.Coordinates) < 2 {
		return models.NormalizedEvent{}, false
	}

	longitude := geometry.Coordinates[0]
	latitude := geometry.Coordinates[1]
	eventType := ""
	if len(event.Categories) > 0 {
		eventType = event.Categories[0].ID
	}

	return models.NormalizedEvent{
		ID:        fmt.Sprintf("eonet:%s", event.ID),
		Source:    "eonet",
		Category:  "natural_disaster",
		Type:      eventType,
		Title:     event.Title,
		Status:    "open",
		Severity:  1,
		StartedAt: geometry.Date,
		UpdatedAt: geometry.Date,
		Longitude: &longitude,
		Latitude:  &latitude,
	}, true
}
