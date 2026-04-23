# PHASE_03_SPRINT_01.md

# Sprint 1 - First Source Ingestion

## Phase

Phase 3 - First Ingestion Pipeline

## Status

**Active**

Current active step:

- Step 1 - Source contract and ingestion shape definition

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

Delivered when:

- source fetch succeeds reliably
- basic failure handling exists

### Step 3 - Normalize source records

Delivered when:

- source records map deterministically to the target shape

### Step 4 - Store normalized events

Delivered when:

- normalized records persist successfully
- repeated runs remain duplicate-safe

### Step 5 - Expose real data through `/v1/events`

Delivered when:

- existing route returns stored real records
- response shape stays stable

### Step 6 - Tests and docs sync

Delivered when:

- minimal backend tests pass for ingestion slice
- docs reflect accepted implementation state

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- one real source works end-to-end
- `/v1/events` returns real stored data
- duplicate-safe behavior is verified
- scope remains narrow and reviewable
