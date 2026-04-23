# PHASE_04.md

# Phase 4 - API Layer

## Status

**Active**

Current sprint:

- Sprint 1 - Events API Foundations

Current active step:

- Step 6 - Tests and docs sync

---

## Step 1 Completion - Real Event Detail Endpoint

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

## Step 2 Completion - Events List Query Baseline Cleanup

Status: Completed

Delivered:

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

## Step 3 Completion - Filtering Support

Status: Completed

Delivered:

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

## Step 4 Completion - Sorting Support

Status: Completed

Delivered:

- optional sorting added for `GET /v1/events`
- supported sort values are limited to `updated_at_desc` and `updated_at_asc`
- invalid sort values return consistent JSON `400` errors

Not added in Step 4:

- detail-route behavior changes
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model redesign

## Step 5 Completion - Pagination Support

Status: Completed

Delivered:

- optional `limit` and `cursor` parameters added
- pagination keeps the stable response shape with `items` and `next_cursor`
- empty string `next_cursor` is still used when there is no next page
- invalid cursor usage returns consistent JSON `400` errors
- minimal backend tests were added/updated for sorting and pagination behavior

Not added in Step 5:

- detail-route behavior changes
- Redis logic
- SSE logic
- scheduler logic
- multi-source abstraction
- Event model redesign

---

## Purpose

Phase 4 turns the first stored ingestion output into stable API behavior.

The goal is to keep the API contract predictable while adding practical event read capabilities in small, reviewable steps.

---

## Sprint 1 Direction

Recommended implementation order:

1. real event detail endpoint first
2. filtering, sorting, and pagination later in the sprint

---

## Phase Scope

In scope:

- event detail read path on existing API route shape
- event list query improvements in later steps
- minimal, focused backend tests for API behavior
- contract consistency for frontend integration

Out of scope:

- ingestion redesign
- multi-source architecture changes
- Redis and SSE expansion for this phase baseline
- large frontend implementation work
- broad platform refactors

---

## Exit Criteria

Phase 4 is complete when:

- event detail endpoint returns real stored records reliably
- event listing supports agreed API-layer behavior for this phase
- filtering, sorting, and pagination behavior for MVP is implemented and stable
- API response contracts remain consistent and reviewable
- local Docker-backed workflow remains healthy

---

## Next Phase

- Phase 5 - First Dashboard
