# TheEye

TheEye is a **map-first global signal platform** built as a structured monorepo.

It exists to collect, normalize, and present meaningful world signals through one coherent event-driven experience. The current working baseline is **v0.1.0 - Initial Working MVP**: a local Docker-backed system with a Go API, USGS earthquake ingestion, stored normalized events, and a first map/feed/detail dashboard.

The long-term product direction is broader than a natural disaster dashboard. Natural and physical events are the first practical signal family, but TheEye can evolve toward human systems, global stability, critical infrastructure, and other high-value world signals as reliable source boundaries emerge.

---

## Purpose

TheEye reduces fragmentation across world-signal sources by giving users one readable place to understand:

- what is happening
- where it is happening
- why it matters

Near-term delivery stays focused on reliable multi-source event monitoring: ingestion, normalization, storage, API stability, and a map-first dashboard flow.

Guiding principles:

- fast initial signal
- map-first exploration
- normalized event model
- controlled, reviewable engineering
- stable local development with Docker-based infrastructure

---

## Current Working Model

TheEye uses a controlled multi-agent workflow:

- **ChatGPT** defines scope, prepares prompts, clarifies architecture, and keeps planning aligned.
- **Codex** is the primary implementation agent, especially for backend and repo-wide document sync.
- **Gemini** is used mainly for frontend direction, UI structure, and backend-aware integration work.
- **Claude Code** is used selectively for final review, risky changes, and milestone-level validation.

Planning now follows **Version Milestones**, not Phase/Sprint documents.

Current documented progress snapshot:

- `docs/VERSION_PLAN.md`
- `CHANGELOG.md`

---

## Backend-First Integration Rule

Frontend work must follow the latest stabilized backend contract.

When a change affects the frontend/backend boundary, use this order:

1. Define the target version milestone, work item, implementation slice, and constraints.
2. Codex implements or updates the backend slice first.
3. Gemini reads the latest backend contract and reviews frontend integration impact.
4. Codex applies any required backend patch if Gemini finds contract or usability issues.
5. Gemini implements the frontend against the latest backend behavior.
6. Claude Code performs selective final review for milestone, risky, or cross-cutting changes.
7. Codex performs final documentation sync if behavior or project state changed.

This rule exists to reduce silent contract drift.

---

## Repository Structure

```text
TheEye/
|- apps/
|  |- dashboard/                  # Frontend application
|- services/
|  |- api/                        # Go API service
|  |- collector/                  # Ingestion workers/connectors
|- infra/
|  |- docker-compose.yml          # Local infrastructure and service orchestration
|- docs/
|  |- VISION.md                   # Long-term product north star
|  |- ROADMAP.md                  # High-level version milestone roadmap
|  |- VERSION_PLAN.md             # Active milestone plan and version rules
|  |- ARCHITECTURE.md             # Current technical architecture
|  |- API.md                      # Current API contract
|  |- DB.md                       # Current persistence baseline and planned DB direction
|- scripts/                       # Helper scripts
|- README.md
|- AGENTS.md
|- GEMINI.md
|- CLAUDE.md
|- CHANGELOG.md
```

---

## Document Map

### `docs/VISION.md`

Use for long-term product meaning, MVP boundaries, and the broader map-first global signal direction.

### `docs/VERSION_PLAN.md`

Use for active version milestone planning, version rules, completed v0.1.0 scope, and planned v0.2-v0.5 direction.

### `docs/ROADMAP.md`

Use for high-level roadmap orientation. It should stay concise and should not become a detailed task list.

### `docs/ARCHITECTURE.md`

Use for current system structure, service responsibilities, and ingestion-to-dashboard flow.

### `docs/API.md`

Use for current backend API contract and planned-but-not-active API notes.

### `docs/DB.md`

Use for current event persistence behavior, idempotency rules, and planned database direction.

### `AGENTS.md`

Use for source-of-truth engineering rules, stack and contract rules, local dev constraints, and multi-agent coordination policy.

---

## Source of Truth Order

If documents conflict, follow this priority:

1. `AGENTS.md`
2. `docs/VISION.md`
3. `docs/VERSION_PLAN.md`
4. `docs/ROADMAP.md`
5. `docs/ARCHITECTURE.md`
6. `docs/API.md`
7. `docs/DB.md`
8. `README.md`
9. code

Implementation must follow the documented plan, not invent a new one.

---

## Local Development

The local development flow must remain stable.

### Start infrastructure and local stack

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

This is the baseline entry point for local development and should continue to support:

- PostgreSQL / PostGIS
- Redis
- API service
- collector service

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

1. confirm the target version milestone, work item, and implementation slice
2. clarify constraints with ChatGPT
3. implement backend or contract-changing work with Codex
4. let Gemini inspect the latest backend contract for frontend impact
5. patch backend if needed
6. implement frontend with Gemini only after the contract is stable
7. use Claude Code selectively for final or risky review
8. let Codex sync docs last when state changed
9. commit a clean, scoped unit
10. create a tag only after a meaningful milestone

---

## Scope Discipline

To avoid repo drift:

- do not mix unrelated work in one change set
- do not let agents silently expand the active work item
- do not treat tool output as a project decision until reflected in docs
- do not let frontend invent backend fields or endpoints
- do not break the Docker-based local workflow

---

## Tagging Philosophy

Tags should be milestone-based, not commit-based.

Prefer creating tags when:

- a meaningful version milestone is complete
- docs are synced
- the working tree is clean
- the result is reviewed and intentionally checkpointed

---

## Final Note

TheEye should be built in a way that remains understandable, traceable, and stable even when multiple AI tools are involved.

When in doubt, the documented project direction wins.
