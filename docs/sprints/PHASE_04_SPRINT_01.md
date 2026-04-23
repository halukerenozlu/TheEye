# PHASE_04_SPRINT_01.md

# Sprint 1 - Events API Foundations

## Phase

Phase 4 - API Layer

## Status

**Active**

Current active step:

- Step 4 - Sorting support

---

## Goal

Deliver the first stable API-layer slice on top of stored event data, starting with real event detail behavior.

---

## Scope

### In scope

- real event detail endpoint implementation first
- list endpoint behavior kept stable while detail path matures
- filtering, sorting, and pagination in later sprint steps
- minimal backend tests for endpoint behavior and response shape

### Out of scope

- ingestion pipeline redesign
- scheduler or worker orchestration changes
- Redis and SSE expansion
- multi-source abstraction redesign
- unrelated refactors

---

## Step Breakdown

### Step 1 - Real event detail endpoint

Status: Completed

Delivered:

- `GET /v1/events/{id}` now returns real stored event data when found
- unknown id keeps the consistent JSON `404` shape
- existing `GET /v1/events` behavior remains unchanged
- minimal backend tests were added/updated for found and not-found behavior

Not added in Step 1:

- filtering/sorting/pagination changes
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model redesign

### Step 2 - Events list query baseline cleanup

Status: Completed

Delivered:

- `/v1/events` list behavior remains stable against real stored records
- no contract drift is introduced while preparing list query improvements
- baseline query parsing/validation added for `GET /v1/events`
- default route behavior remains unchanged when no query params are provided

Not added in Step 2:

- sorting changes
- pagination changes
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model redesign

### Step 3 - Filtering support

Status: Completed

Delivered:

- agreed MVP filters are implemented with explicit parameter handling
- invalid query parameters are handled predictably
- optional filtering added for `type`, `started_after`, `started_before`
- invalid query handling now returns consistent JSON `400` errors
- minimal backend tests added/updated for filtering and invalid query cases

Not added in Step 3:

- sorting changes
- pagination changes
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model redesign

### Step 4 - Sorting support

Delivered when:

- agreed sort behavior is implemented with stable defaults
- unsupported sort values are handled clearly

### Step 5 - Pagination support

Delivered when:

- practical MVP pagination behavior is implemented
- response contract remains stable for frontend integration

### Step 6 - Tests and docs sync

Delivered when:

- minimal backend tests cover detail/list/filter/sort/pagination behavior
- sprint and phase docs reflect accepted implementation state

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- real event detail endpoint is stable
- list/filter/sort/pagination behavior for the sprint scope is implemented
- API response contracts remain consistent
- code and docs are reviewable and aligned
