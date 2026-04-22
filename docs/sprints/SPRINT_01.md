# SPRINT_01.md

# Sprint 1 — Backend Service Skeleton

## Phase

Phase 2 — Backend Foundation

## Status

**Active**

Current remaining work:

- none in Sprint 1 (Steps 1-6 completed)

---

## Goal

Build the first stable, minimal backend API skeleton for TheEye without introducing real data ingestion yet.

This sprint focuses on correctness, shape, and local development stability.

---

## Scope

### In scope

- API bootstrap
- basic health/meta routes
- structured config
- graceful shutdown
- placeholder events list endpoint
- placeholder event detail endpoint
- initial typed `Event` response model
- response and error consistency cleanup

### Out of scope

- database-backed reads
- Redis integration
- collector logic
- SSE
- real filtering and pagination
- real event ingestion
- broad frontend implementation

---

## Steps

### Step 1 — Minimal API skeleton

Status: Completed

Delivered:

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`

---

### Step 2 — Structured config + graceful shutdown

Status: Completed

Delivered:

- environment-based config
- local defaults
- `http.Server`
- graceful shutdown flow

---

### Step 3 — `GET /v1/events` placeholder

Status: Completed

Delivered:

- placeholder list route
- empty response foundation:
  - `{ "items": [], "next_cursor": "" }`

---

### Step 4 — Minimal typed Event response model draft

Status: Completed

Delivered:

- minimal typed `Event` model
- typed empty events list response

---

### Step 5 — `GET /v1/events/{id}` placeholder

Status: Completed

Delivered:

- detail route exists
- placeholder JSON `404` response

---

### Step 6 — Response and error shape cleanup

Status: Completed

Delivered:

- standardized minimal JSON error writing through a small reusable helper
- kept `GET /v1/events` response shape stable:
  - `{ "items": [], "next_cursor": "" }`
- kept `GET /v1/events/{id}` placeholder behavior and aligned it to the consistent error shape
- added router-level consistent JSON errors for `404` and `405`
- added minimal backend tests for Step 6 behavior:
  - JSON error helper behavior
  - `GET /v1/events/{id}` placeholder error shape consistency
  - router-level `404` and `405` JSON error shape
  - `GET /v1/events` shape stability
- no scope expansion beyond Step 6

---

## Integration Handoff Rule

This sprint is backend-first.

If frontend-facing contract questions appear during or after Step 6:

1. Gemini reviews the latest backend behavior from an integration perspective.
2. Codex applies any necessary backend patch.
3. Frontend work begins only after the contract is stable.

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- the API boots cleanly
- health/meta/events routes behave predictably
- placeholder list/detail contracts exist
- config/shutdown behavior is stable
- response and error shapes are sufficiently consistent
- the code remains small and reviewable
- the local Docker flow still works

---

## Notes

This sprint is intentionally foundation-heavy.
The aim is not feature richness.
The aim is to create a backend shell that later ingestion work and frontend integration can attach to safely.

