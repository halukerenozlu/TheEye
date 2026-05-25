package models

type NormalizedEvent struct {
	ID        string   `db:"id"`
	Source    string   `db:"source_name"`
	Category  string   `db:"category"`
	Type      string   `db:"type"`
	Title     string   `db:"title"`
	Status    string   `db:"status"`
	Severity  int      `db:"severity"`
	StartedAt string   `db:"started_at"`
	UpdatedAt string   `db:"updated_at"`
	Longitude *float64 `db:"longitude"`
	Latitude  *float64 `db:"latitude"`
}
