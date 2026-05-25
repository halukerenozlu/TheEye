package collector

import (
	"context"

	"theeye/services/collector/models"
)

type Source interface {
	Fetch(ctx context.Context) ([]models.NormalizedEvent, error)
	Name() string
}
