# PHASE_03_SPRINT_01.md

# Sprint 1 - First Source Ingestion

## Phase

Phase 3 - First Ingestion Pipeline

## Status

**Completed**

Current active step:

- none (Sprint 1 completed)

---

## Goal

Deliver the first end-to-end ingestion slice using one real source (USGS earthquakes).

---

## Recommended Source

- USGS earthquakes feed (`FeatureCollection`)

---

## Scope

### In scope

- source contract definition for USGS feed
- normalization target for current Event direction
- source fetch client
- minimal persistence flow
- duplicate-safe writes
- wiring `/v1/events` to stored real data
- minimal backend tests
- minimal docs sync

### Out of scope

- second source
- SSE
- Redis enhancements
- advanced API filtering
- frontend implementation
- major Event model redesign

---

## Step Breakdown

### Step 1 - Source contract and normalization target

Delivered when:

- first source and feed format are fixed
- ingest fields are fixed
- normalization target mapping is fixed
- idempotency keys are defined
- scope guardrails are explicit

Step 1 mapping baseline:

- `id` <- `usgs:{id}`
- `type` <- `earthquake`
- `title` <- `properties.title`
- `status` <- `properties.status` (fallback `unknown`)
- `severity` <- magnitude-to-severity mapping (left to implementation)
- `started_at` <- `properties.time`
- `updated_at` <- `properties.updated`

Step 1 idempotency baseline:

- `source_name = usgs`
- `source_event_id = <feature.id>`

### Step 2 - Source fetch client

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

### Step 3 - Normalize source records

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

### Step 4 - Store normalized events

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

### Step 5 - Expose real data through `/v1/events`

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

### Step 6 - Tests and docs sync

Status: Completed

Delivered:

- minimal backend tests pass for ingestion slice
- sprint docs reflect accepted implementation state

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- one real source works end-to-end
- `/v1/events` returns real stored data
- duplicate-safe behavior is verified
- scope remains narrow and reviewable
