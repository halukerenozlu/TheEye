package usgs

import "testing"

func TestNormalizeFeatureDeterministicIDMapping(t *testing.T) {
	feature := Feature{
		ID: "abc123",
		Properties: FeatureProperties{
			Title:   "M 2.5 - Test",
			Status:  "reviewed",
			Time:    1700000000000,
			Updated: 1700000001000,
		},
	}

	first := NormalizeFeature(feature)
	second := NormalizeFeature(feature)

	if first.ID != "usgs:abc123" {
		t.Fatalf("id = %q, want %q", first.ID, "usgs:abc123")
	}
	if first.ID != second.ID {
		t.Fatalf("deterministic id failed: first=%q second=%q", first.ID, second.ID)
	}
}

func TestNormalizeFeatureTypeAndTitleMapping(t *testing.T) {
	feature := Feature{
		ID: "id-1",
		Properties: FeatureProperties{
			Title:   "M 1.2 - Title From Source",
			Status:  "automatic",
			Time:    1700000000000,
			Updated: 1700000000000,
		},
	}

	got := NormalizeFeature(feature)
	if got.Type != "earthquake" {
		t.Fatalf("type = %q, want %q", got.Type, "earthquake")
	}
	if got.Title != "M 1.2 - Title From Source" {
		t.Fatalf("title = %q, want %q", got.Title, "M 1.2 - Title From Source")
	}
}

func TestNormalizeFeatureStatusMapping(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "reviewed", in: "reviewed", want: "confirmed"},
		{name: "automatic", in: "automatic", want: "preliminary"},
		{name: "deleted", in: "deleted", want: "removed"},
		{name: "unknown", in: "something_else", want: "unknown"},
		{name: "empty", in: "", want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feature := Feature{
				ID: "id-status",
				Properties: FeatureProperties{
					Status:  tt.in,
					Time:    1700000000000,
					Updated: 1700000000000,
				},
			}

			got := NormalizeFeature(feature)
			if got.Status != tt.want {
				t.Fatalf("status = %q, want %q", got.Status, tt.want)
			}
		})
	}
}

func TestNormalizeFeatureTimeConversion(t *testing.T) {
	feature := Feature{
		ID: "id-time",
		Properties: FeatureProperties{
			Time:    1700000000000,
			Updated: 1700000001000,
		},
	}

	got := NormalizeFeature(feature)
	if got.StartedAt != "2023-11-14T22:13:20Z" {
		t.Fatalf("started_at = %q, want %q", got.StartedAt, "2023-11-14T22:13:20Z")
	}
	if got.UpdatedAt != "2023-11-14T22:13:21Z" {
		t.Fatalf("updated_at = %q, want %q", got.UpdatedAt, "2023-11-14T22:13:21Z")
	}
}

func TestNormalizeFeatureSeverityMapping(t *testing.T) {
	tests := []struct {
		name string
		mag  *float64
		want int
	}{
		{name: "nil", mag: nil, want: 0},
		{name: "low", mag: ptrFloat64(1.9), want: 1},
		{name: "moderate", mag: ptrFloat64(2.0), want: 2},
		{name: "strong", mag: ptrFloat64(4.0), want: 3},
		{name: "major", mag: ptrFloat64(6.0), want: 4},
		{name: "great", mag: ptrFloat64(7.0), want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feature := Feature{
				ID: "id-severity",
				Properties: FeatureProperties{
					Mag:     tt.mag,
					Time:    1700000000000,
					Updated: 1700000000000,
				},
			}

			got := NormalizeFeature(feature)
			if got.Severity != tt.want {
				t.Fatalf("severity = %d, want %d", got.Severity, tt.want)
			}
		})
	}
}

func ptrFloat64(v float64) *float64 {
	return &v
}
