# ROADMAP.md

## Vision

TheEye aims to become a map-first global signal platform that collects, organizes, and presents meaningful signals through an event-driven product.

Near-term roadmap focus:

- reliable multi-source event ingestion and normalization
- stable map + feed + detail monitoring flow
- disciplined delivery through active phase/sprint scope

Long-term north star:

- gradual expansion from natural/physical events toward broader global signals
- eventual inclusion of carefully scoped human-system and global-stability signals

Long-term direction should guide prioritization, but active implementation scope remains phase/sprint-driven.

---

## Current Planning Status

- **Current active phase:** Phase 6 - Signal Expansion Foundation
- **Current active sprint:** Sprint 1 - Reliability and Contract Foundation (Completed)
- **Current active step:** None (Sprint completed)

For current implementation boundaries, follow:

- `docs/phases/PHASE_06.md`
- `docs/sprints/PHASE_06_SPRINT_01.md`

---

## Phase 0 — Foundation

**Status:** Completed

**Goal:**

- repository structure
- infra baseline
- local development workflow
- scripts and docs hygiene

---

## Phase 1 — Product Definition and Domain Model

**Status:** Completed

**Goal:**

- define product scope
- define user scenarios
- define MVP boundaries
- define domain entities and database draft


---

## Phase 2 — Backend Foundation

**Status:** Completed

**Goal:**

- application skeleton
- configuration and health checks
- database integration baseline
- migration base
- service boundaries
- stable local Docker-backed flow for backend development


---

## Phase 3 — First Ingestion Pipeline

**Status:** Completed

**Goal:**

- ingest data from one source
- normalize it
- store it
- avoid duplicates
- establish the first usable data pipeline


---

## Phase 4 — API Layer

**Status:** Completed

**Goal:**

- event listing
- event detail
- filtering
- sorting
- pagination


---

## Phase 5 — First Dashboard

**Status:** Completed

**Goal:**

- map-first dashboard shell
- list and feed views
- filters
- event detail UI
- stable first user-facing interface


---

## Phase 6 - Signal Expansion Foundation

**Status:** Active

**Goal:**

- stabilize the current single-source monitoring flow
- establish backend contract reliability (`severity_level`, `category`)
- run predictable polling + refresh behavior
- keep collector-backed ingest reliable in local Docker flow
- prepare a clean base for one next real source without scope drift


---

## Phase 7 — Alerts and Management

**Status:** Upcoming

**Goal:**

- alert rules
- tracked topics or regions
- basic admin capabilities


---

## Phase 8 — Stabilization and Release Prep

**Status:** Upcoming

**Goal:**

- performance checks
- cleanup
- hardening
- demo readiness
- release preparation

---

## Progress Visibility Rule

Completed phases should remain visible for historical clarity, but should not stay mixed into the active work mentally.

Use:

- active phase / sprint / step for current work
- completed sections for historical reference
- version tags only at meaningful checkpoints

Vision/scope note:

- `VISION.md` is the north star for long-term direction.
- near-term build scope is controlled by active phase/sprint documents.

