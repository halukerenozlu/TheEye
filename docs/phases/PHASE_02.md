# Phase 2 — Backend Foundation

## Purpose

This phase establishes the first working backend foundation for TheEye.

The goal is not real ingestion yet.
The goal is a stable API skeleton that reflects the product direction and can support future data work.

---

## Target Outcome

By the end of Phase 2, the project should have:

- a runnable API service
- basic health and metadata endpoints
- structured config baseline
- graceful shutdown
- first events list placeholder
- first event detail placeholder
- an initial typed Event response draft

This phase is about building the backend shell correctly.

---

## Current Status

### Sprint 1 — Service Skeleton

Progress so far:

#### Step 1 — Minimal API skeleton

Status: Completed

Delivered:

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`

#### Step 2 — Structured config + graceful shutdown

Status: Completed

Delivered:

- `PORT`
- `APP_NAME`
- `APP_ENV`
- `APP_VERSION`
- `http.Server`
- graceful shutdown

#### Step 3 — `GET /v1/events` placeholder

Status: Completed

Delivered:

- returns:
  - `{ "items": [], "next_cursor": "" }`

#### Step 4 — Minimal typed Event response model draft

Status: Completed

Delivered:

- introduced minimal typed `Event` model with:
  - `id`
  - `type`
  - `title`
  - `status`
  - `severity`
  - `started_at`
  - `updated_at`

#### Step 5 — `GET /v1/events/{id}` placeholder

Status: Completed

Delivered:

- route exists
- currently returns JSON `404` placeholder

Example:

```json
{
  "error": "event_not_found",
  "message": "event not found"
}
```

#### Step 6 — Response and error shape cleanup

Status: Planned

Planned focus:

- standardize small JSON response helpers if useful
- keep changes minimal
- avoid premature abstraction
- prepare for future API consistency

---

## What Phase 2 Explicitly Does Not Do

Not in scope yet:

- real database-backed event lookup
- Redis-backed caching/streaming
- collector-based ingestion
- PostGIS bbox querying
- SSE event streaming
- real filtering/pagination logic
- frontend/backend production integration

---

## Why This Phase Matters

Phase 2 gives the project:

- a stable backend entry point
- an inspectable API contract
- a place to attach ingestion later
- a foundation for CI and future tests

Without this phase, later data and dashboard work would drift.

---

## Exit Criteria

Phase 2 is considered complete when:

- API skeleton is stable
- all base placeholder endpoints work
- config and shutdown behavior are solid
- response shapes are defined enough to support the first ingestion work
- the team can safely move into the first real data pipeline

---

## Likely Next Direction After Phase 2

After finishing the remaining cleanup in this phase, the next phase should move toward:

- first ingestion pipeline
- source normalization
- storage-backed event flow
- first real data appearing through `/v1/events`
