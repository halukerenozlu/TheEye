# PHASE_06_SPRINT_01.md

# Sprint 1 - Reliability and Contract Foundation

## Phase

Phase 6 - Signal Expansion Foundation

## Status

**Completed**

Current active step:

- None (Sprint completed)

---

## Sprint Goal

Make the current single-source dashboard flow more reliable, lock the next API contract decisions, and prepare the smallest acceptable foundation for one additional real signal source.

---

## Scope

### In scope

- backend-owned severity normalization for dashboard use
- category field decision and spec alignment
- simple client-side polling for MVP freshness
- minimal UI polish limited to clearer Level 2 / Level 3 distinction
- verification and docs sync for the accepted sprint scope

### Out of scope

- second-source implementation beyond planning-ready foundation
- clustering
- hover-heavy map interaction
- broad dashboard redesign
- auth, notifications, or timeline work
- SSE implementation

---

## Step Breakdown

### Step 1 - Severity backend normalization and API contract decision

Status: Completed

Delivered:

- backend severity normalization implemented
- API responses now expose `severity_level`
- legacy `severity` kept temporarily for compatibility
- minimal backend tests updated

### Step 2 - Category field decision and docs/spec alignment

Status: Completed

Delivered:

- `category` was added to the backend contract
- the distinction between `type` and `category` was kept explicit
- current USGS events were aligned to `category = natural_disaster`
- relevant backend tests were updated

### Step 3 - Simple client-side polling

Status: Completed

Delivered:

- frontend polling was added
- a collector service was added to compose
- collector now starts automatically and runs periodic ingest
- live USGS data now flows into the DB
- `/v1/events` now returns current records with `category` and `severity_level`

### Step 4 - Level 2 / Level 3 visual distinction cleanup

Status: Completed

Delivered:

- polling and freshness visibility are now clear in the dashboard flow
- Level 1 / Level 2 / Level 3 visual distinction is now clearer across feed, map, and detail views
- the restrained dark layout and existing interaction model were preserved

### Step 5 - Verification and docs sync

Status: Completed

Delivered:

- USGS live flow is now reliable for the current single-source scope
- collector runs automatically in compose with periodic ingest
- polling and freshness visibility are in place
- severity/category contract baseline is stabilized for frontend use
- Level 1 / Level 2 / Level 3 visual separation is now clearer

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- severity normalization ownership is clearly backend-side
- the category decision is explicit and documented
- MVP freshness direction is settled on simple polling
- Level 2 and Level 3 presentation is clearer without redesign
- docs are aligned and implementation-ready for the next Phase 6 work
