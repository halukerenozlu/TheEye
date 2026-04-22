# AGENTS.md — TheEye Repository Rules

> Read this file **first** before making any code changes.
> This repository is designed for controlled work across humans, Codex, Gemini, Claude Code, and ChatGPT.

## 0) Project Summary

**TheEye** is a near real-time, map-first world monitoring platform.

Primary signal categories include:

- earthquakes
- wildfires
- storms
- flights
- traffic incidents

Core UX:

- a primary map view
- event layers
- a live feed panel
- time and type filters
- one normalized event model across sources

---

## 1) Product Goals (Must-Haves)

- **Fast initial signal**: meaningful events should appear quickly.
- **Map-first**: geographic exploration is primary.
- **Normalized event model**: different sources converge into one `Event` schema.
- **Near real-time is enough for MVP**: reliability matters more than theoretical perfection.
- **Operational safety**: caching, rate limiting, retries, and observability-ready behavior.
- **Stable local development**: Docker-based startup must remain dependable.

---

## 2) Non-Goals (Avoid)

- big-bang rewrites
- premature microservices fragmentation
- heavy deployment refactors early
- unofficial or unsafe scraping
- scope creep hidden behind "helpful" improvements
- frontend work that invents backend contracts

---

## 3) Tech Stack (Locked for MVP)

### Frontend

- Next.js (TypeScript)
- Tailwind CSS
- shadcn/ui
- MapLibre GL JS
- TanStack Query
- Zustand

### Backend

- Go
- minimal HTTP router
- SSE (Server-Sent Events) for MVP realtime
- PostgreSQL + PostGIS
- Redis

### Dev / Infra

- Docker + Docker Compose
- GitHub Actions for CI later

---

## 4) Repository Layout (Target)

```text
apps/
  dashboard/      # Frontend app

services/
  api/            # Go API service
  collector/      # Go ingestion workers/connectors

shared/
  schema/         # Shared event contracts and generated types

infra/
  docker-compose.yml
```

Agents must not invent a fundamentally different layout without explicit approval.

---

## 5) Source-of-Truth Data Model: `Event`

All sources must normalize into the same `Event` structure.

### Required fields

- `id`
- `type`
- `title`
- `status`
- `severity`
- `started_at`
- `updated_at`
- `geometry`
- `source`

### Recommended fields

- `confidence`
- `ended_at`
- `location`
- `tags`
- `metrics`
- `raw`

Database uniqueness should enforce idempotent source ingestion using:

- `UNIQUE (source_name, source_event_id)`

---

## 6) API Contract (MVP)

### Health

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`

### Events

- `GET /v1/events`
- `GET /v1/events/{id}`
- `GET /v1/events/changes`

### Realtime

- `GET /v1/stream/events` via SSE

Once introduced, endpoints should remain stable.
Breaking changes should prefer new versioned routes rather than silent mutation.

---

## 7) Local Development Rules (Do Not Break)

The baseline local flow must continue to work with:

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

This flow should remain compatible with the services required for the MVP, including:

- PostgreSQL / PostGIS
- Redis
- API
- optional collector

Frontend should remain runnable with a single command such as:

```bash
pnpm --filter dashboard dev
```

Do not break Docker startup, service wiring, or the local-first workflow.

---

## 8) Engineering Standards

### Go

- explicit error handling
- no panics in request paths
- context-aware IO
- structured logging preferred
- timeouts on outbound HTTP calls
- bounded retries with backoff in collectors

### TypeScript / Frontend

- strict TypeScript
- avoid `any` unless justified
- keep map rendering performant
- debounce viewport-driven queries
- handle loading, empty, and error states explicitly

### Dependencies

- add new dependencies only when they clearly reduce complexity
- avoid framework churn

---

## 9) Performance and Reliability Baseline

- use PostGIS indexes for geographic queries
- cache hot paths in Redis where appropriate
- rate limit inbound and outbound traffic where needed
- keep collector writes idempotent
- avoid unbounded world-scale queries when bbox filtering exists

---

## 10) Security and Secrets

- never commit secrets
- keep `.env.example` current
- validate all query parameters
- keep local and shared config explicit

---

## 11) Testing and Verification Expectations

Minimum expectation for meaningful changes:

- code compiles
- local run steps are clear
- impacted routes or UI flows are verified
- Docker flow remains healthy
- frontend/backend contract is checked if the boundary changed

---

## 12) Multi-Agent Role Separation

### Human

Responsible for:

- product direction
- priorities
- approval or rejection
- final tradeoff decisions

### ChatGPT

Responsible for:

- architecture framing
- roadmap / phase / sprint / step definition
- Codex prompt generation
- Gemini prompt generation
- Claude Code review prompt generation
- review interpretation
- commit and tag suggestions
- documentation alignment guidance

### Codex

Primary implementation agent.

Responsible for:

- backend implementation by default
- repo-wide implementation work
- focused code changes
- document sync at the end of accepted work
- keeping scope tight

Codex must not silently expand the sprint.

### Gemini

Primarily responsible for:

- frontend direction when design is still undefined
- frontend implementation
- backend-aware integration checks before frontend coding
- UI structure, component flow, and UX shaping within scope

Gemini must not invent backend fields, response shapes, or new product scope.

### Claude Code

Selective review agent.

Responsible for:

- final review on risky, milestone, or cross-cutting work
- checking scope correctness
- checking contract drift
- identifying regressions and unnecessary complexity
- suggesting minimal fixes only when necessary

Claude Code is **not** the primary implementer in this project.

---

## 13) Backend-First Integration Protocol

When work touches the frontend/backend boundary, follow this order:

1. ChatGPT defines the exact step and boundaries.
2. Codex implements the backend or contract-changing work first.
3. Gemini reads the latest backend diff, docs, and contract.
4. Gemini reports integration risks or frontend impact before frontend coding begins.
5. Codex applies any required backend patch.
6. Gemini implements the frontend against the finalized backend behavior.
7. Claude Code reviews the integrated result only when the change is risky, milestone-level, or cross-cutting.
8. Codex syncs the docs last.

This is the default integration path for TheEye.

---

## 14) Step-Based Delivery Rule

All implementation work must map to:

- Phase
- Sprint
- Step

No coding work should begin unless the current step is explicit.

Allowed:

- work only on the active step
- minimal supporting changes required by that step
- doc updates needed to reflect accepted behavior

Not allowed:

- unrelated refactors
- speculative optimization
- hidden feature expansion
- changing contracts casually
- parallel redesign of product direction

---

## 15) Review Decision Categories

Review results should be interpreted as one of:

- Accept
- Accept with minimal patch
- Rework needed
- Reject

Required fixes and optional suggestions should be separated clearly.

---

## 16) Documentation Sync Rule

Documentation should not drift from accepted implementation.

Preferred rule:

- planning docs are clarified before work
- code changes are implemented
- review is completed
- **Codex updates the final documents last** to reflect the accepted state

Do not leave backend, Docker, contract, or sprint status changes undocumented.

---

## 17) Version Tag Discipline

Use milestone-based tags in the format:

- `vMAJOR.MINOR.PATCH`

Rules:

- commits are for normal progress
- tags are for meaningful milestones
- sprint progress alone usually does not justify a tag
- phase completion is a strong candidate for a tag
- docs should be synced before tagging

---

## 18) Source of Truth Order

If documents conflict, follow this order:

1. `AGENTS.md`
2. `WORKFLOW.md`
3. `VERSIONING.md`
4. current phase document
5. current sprint document
6. code details

The repository documentation wins over ad-hoc tool output.
