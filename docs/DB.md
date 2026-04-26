# TheEye Database

## Current Baseline

TheEye currently persists normalized events from the first real source: USGS earthquakes.

The v0.1.0 baseline is intentionally small:

- store normalized event records
- avoid duplicate writes from repeated source fetches
- expose stored events through the API
- keep room for future PostGIS-backed geographic queries

---

## Event Persistence Model

The current persistence model stores source-derived events after normalization.

Core Event direction:

- `id`
- `type`
- `category`
- `title`
- `status`
- `severity`
- `severity_level`
- `started_at`
- `updated_at`
- `geometry`
- `source`

Recommended future fields may include:

- `confidence`
- `ended_at`
- `location`
- `tags`
- `metrics`
- `raw`

These should be added only through scoped version milestone work.

---

## Idempotency

Duplicate-safe ingestion is based on the source identity pair:

```text
source_name + source_event_id
```

The intended database uniqueness rule is:

```sql
UNIQUE (source_name, source_event_id)
```

This allows the collector to fetch the same upstream source records repeatedly without creating duplicate TheEye events.

Current behavior includes:

- batch-level deduplication before writes
- conflict-safe write behavior at the database layer
- deterministic event ids for USGS records using `usgs:{source_id}`

---

## USGS Source Persistence

USGS earthquakes are the current source-of-truth ingestion baseline.

Current USGS behavior:

- source name: `usgs`
- source event id: USGS feature id
- normalized id: `usgs:{source_id}`
- type: `earthquake`
- category: `natural_disaster`
- timestamps normalized from USGS source values
- geometry persisted for API/dashboard use
- repeated records update or preserve the existing stored event rather than creating duplicates

---

## PostGIS Direction

PostGIS is part of the local infrastructure and future map-first database direction.

Planned uses:

- geographic indexes
- bbox filtering
- map viewport queries
- proximity or region-based event lookups
- future mixed-source map performance improvements

Current baseline:

- event geometry exists for dashboard marker rendering
- deeper geospatial query behavior is planned, not yet the active documented API baseline

---

## Redis Direction

Redis is included in local infrastructure for future reliability and performance work.

Potential planned uses:

- caching hot API paths
- freshness or collector coordination
- future stream/fan-out support

Current baseline:

- Redis is not the source of truth for events
- event persistence remains database-backed

---

## Migration and Schema Status

Current baseline:

- a minimal schema creation path exists because full migration infrastructure is not yet the primary documented workflow
- the schema supports storing normalized USGS events and enforcing duplicate-safe behavior
- tests cover schema creation, successful upsert, duplicate-safe behavior, invalid input, and database error handling

Planned:

- formal migration workflow
- explicit schema documentation as the Event model matures
- PostGIS indexes for geographic queries
- stricter source adapter persistence boundaries for multi-source ingestion

When uncertain, implementation work should preserve current baseline behavior and introduce schema changes only through a clear version milestone work item.
