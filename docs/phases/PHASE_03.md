# PHASE_03.md

# Phase 3 - First Ingestion Pipeline

## Status

**Completed**

Current sprint:

- Sprint 1 - First Source Ingestion

Current active step:

- Sprint 1 completed; no active step in Phase 3

---

## Purpose

Phase 3 establishes the first real ingestion pipeline.

The phase goal is one small, credible end-to-end flow that can fetch, normalize, store, and expose real source data.

---

## Recommended First Source

- USGS earthquakes feed
- source contract format: USGS GeoJSON `FeatureCollection`

---

## Step 1 Contract and Normalization Target

USGS source fields to ingest per feature:

- `id`
- `properties.time`
- `properties.updated`
- `properties.mag`
- `properties.status`
- `properties.title`
- `geometry.coordinates`

Normalization target for this first slice:

- `id` <- `usgs:{id}`
- `type` <- `earthquake`
- `title` <- `properties.title`
- `status` <- `properties.status` (fallback `unknown`)
- `severity` <- magnitude-to-severity mapping (implementation detail, not fixed in this doc)
- `started_at` <- `properties.time`
- `updated_at` <- `properties.updated`

Idempotency target for writes:

- `source_name = usgs`
- `source_event_id = <feature.id>`

Step 1 guardrail:

- keep the current Event direction; do not expand API response fields in Step 1

---

## Step 2 Completion - Source Fetch Client

Status: Completed

Delivered:

- minimal USGS fetch client added
- FeatureCollection-style source payload decoding added
- non-200 upstream handling added
- malformed JSON handling added
- minimal backend tests added

Not added in Step 2:

- normalization into final Event model
- database writes
- API contract changes
- scheduler logic

---

## Step 3 Completion - Normalize Source Records

Status: Completed

Delivered:

- deterministic normalization layer added
- id baseline fixed as `usgs:{source_id}`
- type/title mapping added
- status mapping added
- UTC RFC3339 time conversion added
- deterministic severity mapping added
- minimal backend tests added

Not added in Step 3:

- database writes
- API behavior changes
- Redis/SSE logic
- scheduler logic
- Event model expansion

---

## Step 4 Completion - Store Normalized Events

Status: Completed

Delivered:

- minimal persistence layer added for normalized USGS records
- practical duplicate-safe behavior added
- batch-level deduplication before writes added
- DB-level upsert / conflict-safe write behavior added using `(source_name, source_event_id)`
- minimal schema creation path added only because migration infrastructure is not yet present
- minimal backend tests added for schema creation, successful upsert, duplicate-safe behavior, invalid input, and DB error handling

Not added in Step 4:

- `/v1/events` behavior changes
- DB read/query path
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model expansion

---

## Step 5 Completion - Expose Real Data Through `/v1/events`

Status: Completed

Delivered:

- `GET /v1/events` now reads stored real data
- response shape remains stable with `items` and `next_cursor`
- stable empty response remains when no stored data is available
- minimal backend tests were added/updated for real-data and empty-data behavior

Not added in Step 5:

- `/v1/events/{id}` behavior changes
- filtering/sorting/pagination expansion
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction

---

## Phase Scope

In scope:

- one real source only
- fetch client
- normalization
- persistence
- duplicate-safe writes
- `/v1/events` real-data exposure through existing route
- minimal backend tests
- minimal doc sync

Out of scope:

- second source
- SSE
- Redis optimization
- advanced filtering and sorting
- broad frontend work
- large schema redesign

---

## Exit Criteria

Phase 3 is complete when:

- one real source ingests end-to-end
- normalized records are stored
- duplicate-safe write behavior exists
- `/v1/events` returns real stored events
- implementation remains scoped and reviewable

---

## Next Phase

- Phase 4 - API Layer
