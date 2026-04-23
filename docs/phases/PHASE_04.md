# PHASE_04.md

# Phase 4 - API Layer

## Status

**Active**

Current sprint:

- Sprint 1 - Events API Foundations

Current active step:

- Step 1 - Real event detail endpoint

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
