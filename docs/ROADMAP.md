# ROADMAP.md

## Vision

TheEye aims to become a monitoring and intelligence-style platform that collects, organizes, and presents meaningful signals, events, and sources through a map-first, event-driven product.

The early goal is not maximum scale.
The early goal is a clean, credible, end-to-end working product slice.

---

## Current Planning Status

- **Current active phase:** Phase 2 — Backend Foundation
- **Current active sprint:** Sprint 1 — Backend Service Skeleton
- **Current active step:** Step 6 — Response and error shape cleanup

Phase 1 is treated as effectively complete based on the current documentation set.
Phase 2 remains active.

---

## Phase 0 — Foundation

**Goal:**

- repository structure
- infra baseline
- local development workflow
- scripts and docs hygiene

**Target version:**

- `v0.1.0`

---

## Phase 1 — Product Definition and Domain Model

**Status:** Complete in substance

**Goal:**

- define product scope
- define user scenarios
- define MVP boundaries
- define domain entities and database draft

**Target version:**

- `v0.2.0`

---

## Phase 2 — Backend Foundation

**Status:** Active

**Goal:**

- application skeleton
- configuration and health checks
- database integration baseline
- migration base
- service boundaries
- stable local Docker-backed flow for backend development

**Target version:**

- `v0.3.0`

---

## Phase 3 — First Ingestion Pipeline

**Status:** Upcoming

**Goal:**

- ingest data from one source
- normalize it
- store it
- avoid duplicates
- establish the first usable data pipeline

**Target version:**

- `v0.4.0`

---

## Phase 4 — API Layer

**Status:** Upcoming

**Goal:**

- event listing
- event detail
- filtering
- sorting
- pagination

**Target version:**

- `v0.5.0`

---

## Phase 5 — First Dashboard

**Status:** Upcoming

**Goal:**

- map-first dashboard shell
- list and feed views
- filters
- event detail UI
- stable first user-facing interface

**Target version:**

- `v0.6.0`

---

## Phase 6 — Better UX and Product Depth

**Status:** Upcoming

**Goal:**

- timeline and visualization
- richer filtering
- saved views
- better usability

**Target version:**

- `v0.7.0`

---

## Phase 7 — Alerts and Management

**Status:** Upcoming

**Goal:**

- alert rules
- tracked topics or regions
- basic admin capabilities

**Target version:**

- `v0.8.0`

---

## Phase 8 — Stabilization and Release Prep

**Status:** Upcoming

**Goal:**

- performance checks
- cleanup
- hardening
- demo readiness
- release preparation

**Target version:**

- `v0.9.0` to `v1.0.0`

---

## Progress Visibility Rule

Completed phases should remain visible for historical clarity, but should not stay mixed into the active work mentally.

Use:

- active phase / sprint / step for current work
- completed sections for historical reference
- version tags only at meaningful checkpoints
