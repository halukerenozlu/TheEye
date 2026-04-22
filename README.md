# TheEye

TheEye is a **map-first, near real-time world monitoring platform** built as a structured monorepo.

The repository is designed to support disciplined product planning, scoped implementation, and controlled multi-agent collaboration without losing architectural clarity or breaking the local development flow.

---

## Purpose

TheEye exists to reduce fragmentation across world-signal sources by collecting, normalizing, and presenting them through one coherent event-driven experience.

Early signal categories include:

- earthquakes
- wildfires
- storms
- flights
- traffic incidents

The guiding principle is simple:

- **fast initial signal**
- **map-first exploration**
- **normalized event model**
- **controlled, reviewable engineering**
- **stable local development with Docker-based infrastructure**

---

## Current Working Model

TheEye now follows a **hybrid multi-agent workflow**:

- **ChatGPT** defines scope, prepares prompts, clarifies architecture, and keeps docs aligned.
- **Codex** is the primary implementation agent, especially for backend and repo-wide document sync.
- **Gemini** is used mainly for frontend direction, UI structure, and backend-aware integration work.
- **Claude Code** is used more selectively for final review, risky changes, and milestone-level validation.

This is intentional. The project should lean more heavily on **Codex + Gemini** and use **Claude Code** sparingly.

---

## Backend-First Integration Rule

Frontend work must follow the latest stabilized backend contract.

When a change affects the frontend/backend boundary, use this order:

1. ChatGPT defines the current phase, sprint, step, and constraints.
2. Codex implements or updates the backend slice first.
3. Gemini reads the latest backend contract and reviews it for frontend integration impact.
4. Codex applies any required backend patch if Gemini finds contract or usability issues.
5. Gemini implements the frontend against the latest backend behavior.
6. Claude Code performs a selective final review for milestone, risky, or cross-cutting changes.
7. Codex performs the final documentation sync if behavior or project state changed.

This rule exists to reduce silent contract drift.

---

## Repository Structure

> Keep the meaning of these areas stable even if the exact folder tree evolves.

```text
TheEye/
├─ apps/
│  └─ dashboard/          # Frontend application
├─ services/
│  ├─ api/                # Go API service
│  └─ collector/          # Ingestion workers/connectors
├─ shared/
│  └─ schema/             # Shared event schema, contracts, generated types
├─ infra/
│  └─ docker-compose.yml  # Local infrastructure and service orchestration
├─ scripts/               # Helper scripts
├─ README.md
├─ AGENTS.md
├─ WORKFLOW.md
├─ VERSIONING.md
├─ VISION.md
├─ GEMINI.md
├─ CLAUDE.md
├─ ROADMAP.md
├─ PHASE_01.md
├─ PHASE_02.md
└─ SPRINT_01.md
```

---

## Document Map

### `README.md`
Use for:

- project overview
- repository orientation
- local setup basics
- collaboration rules
- current progress snapshot

### `VISION.md`
Use for:

- product meaning
- MVP boundaries
- long-term direction
- what TheEye is and is not

### `ROADMAP.md`
Use for:

- phase sequence
- milestone intent
- long-term build order

### `PHASE_01.md`, `PHASE_02.md`, ...
Use for:

- phase-specific goals
- scope boundaries
- exit criteria
- current status

### `SPRINT_01.md`, `SPRINT_02.md`, ...
Use for:

- sprint goal
- step breakdown
- active and completed work inside the sprint

### `AGENTS.md`
Use for:

- source of truth engineering rules
- stack and contract rules
- local dev constraints
- multi-agent coordination policy

### `WORKFLOW.md`
Use for:

- exact delivery sequence
- who does what and when
- review and docs update order

### `VERSIONING.md`
Use for:

- version and tag rules
- milestone release discipline

### `GEMINI.md`
Use for:

- Gemini-specific working agreement
- frontend and integration workflow rules

### `CLAUDE.md`
Use for:

- Claude Code review agreement
- selective review expectations
- output format for review results

---

## Source of Truth Order

If documents conflict, follow this priority:

1. `AGENTS.md`
2. `WORKFLOW.md`
3. `VERSIONING.md`
4. current phase document
5. current sprint document
6. `README.md`
7. implementation details in code

Implementation must follow the documented plan, not invent a new one.

---

## Current Documented Progress Snapshot

### Completed or effectively closed work

- **Phase 1 — Product Definition and Domain Model** is documented as complete in substance and can be treated as closed unless new product-direction changes appear.
- **Phase 2 / Sprint 1 / Steps 1-5** are documented as completed:
  - `GET /v1/healthz`
  - `GET /v1/readyz`
  - `GET /v1/meta`
  - structured config + graceful shutdown
  - placeholder `GET /v1/events`
  - minimal typed `Event` draft
  - placeholder `GET /v1/events/{id}`

### Active work

- **Phase 2 — Backend Foundation**
- **Sprint 1 — Backend Service Skeleton**
- **Step 6 — Response and error shape cleanup**

### Important note

Completed work should be moved out of the active focus mentally, but **not deleted from project history**. Keep completed work visible through a small status snapshot or archive section rather than mixing it into active tasks.

---

## Local Development

The local development flow must remain stable.

### Start infrastructure and local stack

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

This should remain the baseline entry point for local development and should continue to support backend-related services such as:

- PostgreSQL / PostGIS
- Redis
- API service
- optional collector service

### Stop local stack

```bash
docker compose -f ./infra/docker-compose.yml down
```

### Frontend development

```bash
pnpm --filter dashboard dev
```

If repo-level shortcuts exist later through `Makefile` or scripts, keep them consistent with the same local flow.

---

## Development Workflow Summary

A clean working loop for this repository is:

1. confirm active phase, sprint, and step
2. clarify constraints with ChatGPT
3. implement backend or contract-changing work with Codex
4. let Gemini inspect the latest backend contract for frontend impact
5. patch backend if needed
6. implement frontend with Gemini only after the contract is stable
7. use Claude Code selectively for final or risky review
8. let Codex sync the docs last when state changed
9. commit a clean, scoped unit
10. create a tag only after a meaningful milestone or completed phase

---

## Scope Discipline

To avoid repo drift:

- do not mix unrelated work in one change set
- do not let agents silently expand the sprint scope
- do not treat tool output as a project decision until reflected in docs
- do not let frontend invent backend fields or endpoints
- do not break the Docker-based local workflow

---

## Tagging Philosophy

Tags should be milestone-based, not commit-based.

Prefer creating tags when:

- a phase is complete
- a major milestone is approved
- the docs are synced
- the working tree is clean
- the result is reviewed and intentionally checkpointed

---

## Final Note

TheEye should not only be built correctly; it should be built in a way that remains understandable, traceable, and stable even when multiple AI tools are involved.

When in doubt, the documented project direction wins.
