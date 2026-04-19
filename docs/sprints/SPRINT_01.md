# Sprint 1 — Backend Service Skeleton

## Phase

Phase 2 — Backend Foundation

## Goal

Build the first stable, minimal backend API skeleton for TheEye without introducing real data ingestion yet.

This sprint focuses on correctness, shape, and local development stability.

---

## Scope

In scope:

- API bootstrap
- basic health/meta routes
- structured config
- graceful shutdown
- placeholder events list endpoint
- placeholder event detail endpoint
- initial typed Event response model

Out of scope:

- database-backed reads
- Redis integration
- collector logic
- SSE
- real filtering and pagination
- real event ingestion
- frontend integration work

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

Status: Planned

Planned focus:

- standardize small JSON response helpers if useful
- keep changes minimal
- avoid premature abstraction
- prepare for future API consistency

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- the API boots cleanly
- health/meta/events routes behave predictably
- placeholder list/detail contracts exist
- config/shutdown behavior is stable
- the code remains small and reviewable

---

## Notes

This sprint is intentionally foundation-heavy.
The aim is not feature richness.
The aim is to create a backend shell that later ingestion work can attach to safely.
