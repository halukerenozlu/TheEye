# PHASE_03.md

# Phase 3 - First Ingestion Pipeline

## Status

**Active**

Current sprint:

- Sprint 1 - First Source Ingestion

Current active step:

- Step 3 - Normalize source records

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
