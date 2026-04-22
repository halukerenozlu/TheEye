# PHASE_02.md

# Phase 2 — Backend Foundation

## Status

**Completed**

Current final sprint:

- Sprint 1 — Backend Service Skeleton (completed)

Current final step:

- Step 6 — Response and error shape cleanup (completed)

Next phase:

- Phase 3 — First Ingestion Pipeline
---

## Purpose

This phase establishes the first working backend foundation for TheEye.

The goal is not real ingestion yet.
The goal is a stable API skeleton that reflects the product direction, preserves the Docker-based local workflow, and prepares the project for later integration and data work.

---

## Target Outcome

By the end of Phase 2, the project should have:

- a runnable API service
- basic health and metadata endpoints
- structured config baseline
- graceful shutdown
- first events list placeholder
- first event detail placeholder
- an initial typed `Event` response draft
- a stable local backend foundation that later ingestion work can build on safely

---

## Current Status

### Sprint 1 — Service Skeleton

#### Completed steps

##### Step 1 — Minimal API skeleton

Delivered:

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`

##### Step 2 — Structured config + graceful shutdown

Delivered:

- `PORT`
- `APP_NAME`
- `APP_ENV`
- `APP_VERSION`
- `http.Server`
- graceful shutdown

##### Step 3 — `GET /v1/events` placeholder

Delivered:

- placeholder list route
- response foundation:
  - `{ "items": [], "next_cursor": "" }`

##### Step 4 — Minimal typed Event response model draft

Delivered:

- introduced minimal typed `Event` model with:
  - `id`
  - `type`
  - `title`
  - `status`
  - `severity`
  - `started_at`
  - `updated_at`

##### Step 5 — `GET /v1/events/{id}` placeholder

Delivered:

- detail route exists
- currently returns JSON `404` placeholder

Example:

```json
{
  "error": "event_not_found",
  "message": "event not found"
}
```

#### Step 6 completion

##### Step 6 — Response and error shape cleanup

Status: Completed

Delivered:

- standardized minimal JSON error writing through a small reusable helper
- kept list/detail placeholder behavior stable and made error responses more consistent
- added router-level consistent JSON errors for `404` and `405`
- added minimal backend tests covering Step 6 response/error behavior
- no scope expansion beyond Step 6

---

## Integration Note

Frontend work should follow the stabilized backend contract from this phase.

If contract issues are discovered from the frontend side:

- Gemini identifies the integration problem
- Codex applies the minimal backend correction
- frontend implementation proceeds only after the contract is accepted

---

## What Phase 2 Explicitly Does Not Do

Not in scope yet:

- real database-backed event lookup
- Redis-backed caching or stream fan-out
- collector-based ingestion
- PostGIS bbox querying
- SSE event streaming
- real filtering and pagination logic
- production deployment work
- large-scale frontend implementation

---

## Why This Phase Matters

Phase 2 gives the project:

- a stable backend entry point
- an inspectable API contract
- a place to attach ingestion later
- a safer integration target for future frontend work
- a foundation for CI and future tests
- a backend that still respects the local Docker workflow

Without this phase, later data and dashboard work would drift.

---

## Exit Criteria

Phase 2 is considered complete when:

- API skeleton is stable
- all base placeholder endpoints work
- config and shutdown behavior are solid
- response shapes are defined enough to support the first ingestion work
- local Docker-backed development remains healthy
- the team can safely move into the first real data pipeline

---

## Next Phase

Phase 3 should now proceed with:

- first ingestion pipeline
- source normalization
- storage-backed event flow
- first real data appearing through `/v1/events`


