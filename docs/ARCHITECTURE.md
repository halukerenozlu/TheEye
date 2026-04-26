# TheEye Architecture

## Current Baseline

TheEye is a structured monorepo with a map-first dashboard, a Go API, a collector service, and Docker-backed local infrastructure.

The current v0.1.0 baseline proves one real source end to end:

```text
USGS earthquakes -> collector/API ingestion code -> normalization -> persistence -> API -> dashboard
```

---

## Repository Areas

### `apps/dashboard`

The frontend application.

Responsibilities:

- render the map-first dashboard
- display the Signal Feed
- render event markers with MapLibre
- show selected event detail
- keep feed, map, and detail selection in sync
- poll the API for MVP freshness
- handle loading, empty, and error states clearly

The dashboard must not invent backend fields or response shapes.

### `services/api`

The Go API service.

Responsibilities:

- expose health, readiness, and metadata endpoints
- expose event list and detail endpoints
- validate query parameters
- return consistent JSON errors
- read stored normalized events
- own backend contract decisions such as `category` and `severity_level`

### `services/collector`

The ingestion worker service.

Responsibilities:

- run periodic source ingest
- fetch USGS earthquake data in the current baseline
- decode source payloads
- normalize source records into TheEye event shape
- write records idempotently
- preserve the existing API/dashboard flow while ingestion evolves

### `infra/docker-compose.yml`

Local service orchestration.

Responsibilities:

- start PostgreSQL / PostGIS
- start Redis
- start API service
- start collector service
- preserve the baseline local command:

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

---

## Data Flow

### 1. Ingestion

The current collector fetches USGS earthquake data as a GeoJSON `FeatureCollection`.

### 2. Normalization

Source-specific records are transformed into TheEye's normalized event direction.

Current USGS baseline:

- `id` uses `usgs:{source_id}`
- `type` is `earthquake`
- `category` is `natural_disaster`
- `title` comes from the source title
- `status` comes from the source status or fallback behavior
- `severity` and `severity_level` are derived by backend normalization
- `started_at` and `updated_at` are converted to API-safe timestamps
- `geometry` is used by the dashboard for marker rendering

### 3. Persistence

Normalized events are persisted with duplicate-safe behavior.

The idempotency baseline is:

- `source_name`
- `source_event_id`

Together these prevent repeated source fetches from creating duplicate events.

### 4. API

The API exposes stored normalized events through:

- `GET /v1/events`
- `GET /v1/events/{id}`

The list endpoint supports filtering, sorting, and pagination for the current MVP baseline.

### 5. Dashboard

The dashboard reads `/v1/events`, renders feed items and map markers, and keeps selected event state synchronized across the feed, map, and detail panel.

---

## PostgreSQL / PostGIS

PostgreSQL is the current persistence foundation.

PostGIS is part of the intended local infrastructure and should support future geographic querying, indexes, bbox filtering, and map-first event exploration. The current baseline uses event geometry for display; deeper geospatial querying should be introduced through scoped version milestone work.

---

## Redis

Redis is part of the local infrastructure baseline.

Planned roles may include:

- hot-path caching
- freshness coordination
- future fan-out or stream support
- operational buffering where useful

Redis should not be expanded without a clear work item.

---

## Context Enrichment Boundary

MCP or similar context enrichment may be evaluated in a later version milestone.

Context enrichment should remain separate from core ingestion:

- core ingestion owns source facts and normalized events
- enrichment may add supplemental context for detail views
- enrichment should not silently mutate source-of-truth event records
- external context sources should not be mixed into the ingestion path without an explicit boundary

This keeps the event model reliable while leaving room for richer context later.
