# AGENTS.md - TheEye Repository Rules

> Read this file **first** before making any code changes.
> These rules are model-agnostic. They apply to every contributor, human or agent,
> regardless of which tool or model is used.

## 0) Project Summary

**TheEye** is a near real-time, map-first global signal platform.

Natural and physical events are the current first signal family, but the long-term direction is broader: human systems, global stability, critical infrastructure, and other critical world signals may enter later through controlled version milestones.

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
- SSE planned later, not active in the current baseline
- PostgreSQL + PostGIS
- Redis

### Dev / Infra

- Docker + Docker Compose
- pnpm workspaces
- GitHub Actions

---

## 4) Repository Layout

```text
apps/
  dashboard/      # Frontend app

services/
  api/            # Go API service
  collector/      # Go ingestion workers/connectors

shared/
  schema/         # Shared event contracts and generated types (planned, not present yet)

infra/
  docker-compose.yml
  .env.example

scripts/          # Local helper scripts (PowerShell + Node)

docs/
  VISION.md
  ROADMAP.md
  VERSION_PLAN.md
  ARCHITECTURE.md
  API.md
  DB.md

.agents/skills/   # Vendored agent skills, pinned in skills-lock.json
.claude/          # Claude Code permissions and slash commands
.github/          # CI workflows and Dependabot config
```

Do not invent a fundamentally different layout without explicit approval.

---

## 5) Source-of-Truth Data Model: `Event`

All sources must normalize into the same `Event` direction.

### Required fields

- `id`
- `type`
- `category`
- `title`
- `status`
- `severity`
- `severity_level`
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

## 6) API Contract (Current Baseline)

### Health

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`

### Events

- `GET /v1/events`
- `GET /v1/events/{id}`

### Planned Later

- `GET /v1/events/changes`
- `GET /v1/stream/events` via SSE
- alert or saved-view endpoints

Once introduced, endpoints should remain stable. Breaking changes should prefer new versioned routes rather than silent mutation.

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
- collector

Frontend should remain runnable with a single command such as:

```bash
pnpm --filter dashboard dev
```

Compose values are parameterized with defaults, so the stack starts without a `.env`. Copy `infra/.env.example` to `infra/.env` only when you need to override something.

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

- use PostGIS indexes for geographic queries where applicable
- cache hot paths in Redis where appropriate
- rate limit inbound and outbound traffic where needed
- keep collector writes idempotent
- avoid unbounded world-scale queries when bbox filtering exists

---

## 10) Security and Secrets

- never commit secrets
- keep `infra/.env.example` current
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

## 12) Planning Model

TheEye uses **Version Milestones**, not Phase/Sprint planning.

All implementation work must map to:

- Version Milestone
- Work Item
- Implementation Slice

No coding work should begin unless the current version milestone and work item are explicit.

Allowed:

- work only on the accepted work item
- make minimal supporting changes required by that slice
- update docs needed to reflect accepted behavior

Not allowed:

- unrelated refactors
- speculative optimization
- hidden feature expansion
- changing contracts casually
- parallel redesign of product direction

---

## 13) Working Principles

Adapted from [multica-ai/andrej-karpathy-skills](https://github.com/multica-ai/andrej-karpathy-skills),
derived from Andrej Karpathy's observations on LLM coding pitfalls.

> **Tradeoff:** these bias toward caution over speed. For trivial tasks, use judgment.

### 13.1 Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

- State assumptions explicitly. If uncertain, **ask**.
- If multiple interpretations exist, present them — don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, **stop**. Name what's confusing. Ask.

### 13.2 Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If 200 lines could be 50, rewrite it.

> Ask: _"Would a senior engineer say this is overcomplicated?"_ If yes, simplify.

### 13.3 Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:

- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, **mention it — don't delete it**.

When your changes create orphans:

- Remove imports/variables/functions that **your** changes made unused.
- Don't remove pre-existing dead code unless asked.

> **The test:** every changed line should trace directly to the request.

### 13.4 Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:

| Instead of...    | Transform to...                                      |
| ---------------- | ---------------------------------------------------- |
| "Add validation" | "Write tests for invalid inputs, then make them pass" |
| "Fix the bug"    | "Write a test that reproduces it, then make it pass"  |
| "Refactor X"     | "Ensure tests pass before and after"                  |

For multi-step tasks, state a brief plan upfront:

```text
1. [Step] -> verify: [check]
2. [Step] -> verify: [check]
3. [Step] -> verify: [check]
```

---

## 14) Implementation and Testing Responsibility

Whoever implements a scoped change is also responsible for the minimum necessary tests for that change.

Rules:

- Backend changes come with the relevant backend tests, run before handing off.
- Frontend changes come with the relevant frontend tests, run when practical.
- Keep tests minimal, relevant, and scoped to the current implementation slice.
- A reviewer judges implementation and test adequacy; a reviewer is not the primary test author.
- The human maintainer performs final smoke testing, approval, commit, and tag decisions.

This rule keeps implementation ownership and test ownership aligned and prevents review passes from turning into implementation passes.

---

## 15) Backend-First Integration Protocol

When work touches the frontend/backend boundary, follow this order:

1. Define the exact version milestone, work item, implementation slice, and boundaries.
2. Implement the backend or contract-changing work first.
3. Read the latest backend diff, docs, and contract before writing any frontend code.
4. Report integration risks or frontend impact before frontend coding begins.
5. Apply any required backend patch.
6. Implement the frontend against the finalized backend behavior.
7. Review the integrated result when the change is risky, milestone-level, or cross-cutting.
8. Sync the docs last.

Frontend work must never invent backend fields, response shapes, endpoints, or product scope.

This is the default integration path for TheEye.

---

## 16) Review Decision Categories

Review results should be interpreted as one of:

- Accept
- Accept with minimal patch
- Rework needed
- Reject

Required fixes and optional suggestions should be separated clearly.

---

## 17) Documentation Sync Rule

Documentation should not drift from accepted implementation.

Preferred rule:

- planning docs are clarified before work
- code changes are implemented
- review is completed
- **documents are updated last** to reflect the accepted state

Do not leave backend, Docker, contract, or milestone status changes undocumented.

---

## 18) Version Tag Discipline

Use milestone-based tags in the format:

- `vMAJOR.MINOR.PATCH`

Rules:

- commits are for normal progress
- tags are for meaningful version milestones
- routine work item progress usually does not justify a tag
- docs should be synced before tagging
- version and tag rules live in `docs/VERSION_PLAN.md`

---

## 19) Source of Truth Order

`docs/VISION.md` defines long-term direction, but active implementation scope is controlled by `docs/VERSION_PLAN.md` and the accepted work item.

If documents conflict, follow this order:

1. `AGENTS.md`
2. `docs/VISION.md`
3. `docs/VERSION_PLAN.md`
4. `docs/ROADMAP.md`
5. `docs/ARCHITECTURE.md`
6. `docs/API.md`
7. `docs/DB.md`
8. `README.md`
9. code

The repository documentation wins over ad-hoc tool output.

---

## 20) Branch and Commit Conventions

### Branch names

Every branch uses a type prefix:

```text
<prefix>/<short-kebab-description>
```

Allowed prefixes: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`.

```text
feat/api-events-cursor-pagination
fix/collector-eonet-timeout
docs/version-plan-update
chore/bump-go-chi
```

Do not commit directly to `master` for anything beyond trivial maintenance.

### Commit messages

Conventional Commits for the header, enforced locally by commitlint:

```text
<type>(<scope>): <subject>
```

Types: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`.
Scopes: `api`, `collector`, `dashboard`, `infra`, `ci`, `deps`, `docs`.

Non-trivial commits add a body. The body groups changes under **`Added`**, **`Changed`**, **`Removed`**, or **`Fixed`** — only these four. Use only the headings that apply, in that order. Every item starts with a dash and a single space:

```text
fix(ci): run dashboard job with pnpm from workspace root

Changed
- dashboard job runs from the workspace root and targets the app with --filter
- pnpm version pinned through the packageManager field

Removed
- package-manager detection ladder that never matched

Fixed
- lockfile was ignored because the install step fell through to npm
```

The four headings mirror the ones used in `CHANGELOG.md`, so a milestone's changelog entry can be assembled from its commits.

Rules:

- header stays under ~72 characters, imperative mood, no trailing period
- one blank line between header and body, and between each section
- items are `- ` (dash, one space), not `*` or `-` without a space
- do not restate the header in the body
- trivial commits (typo, formatting, lockfile refresh) may omit the body

Exempt: commits generated by Dependabot and squash-merge commits created on GitHub. Those are produced server-side and cannot pass through local hooks.

---

## 21) Quick Commands

### Frontend

```bash
pnpm --filter dashboard dev
pnpm --filter dashboard lint
pnpm --filter dashboard typecheck
pnpm --filter dashboard test
pnpm --filter dashboard build
```

### Backend

```bash
cd services/api       && go vet ./... && go test ./... && go build ./...
cd services/collector && go vet ./... && go test ./... && go build ./...
gofmt -l services/api services/collector
```

### Full stack

```bash
docker compose -f ./infra/docker-compose.yml up --build
docker compose -f ./infra/docker-compose.yml down
```

### Version

```bash
pnpm version:sync
pnpm version:check vMAJOR.MINOR.PATCH
```

---

## 22) Skills

Project skills live under `.agents/skills/`, vendored from upstream sources and pinned by hash in `skills-lock.json`.

- Check `.agents/skills/` before concluding that a capability is missing.
- Treat the files there as usable skill instructions.
- Do not edit vendored skills in place; they are replaced wholesale on update.

Currently vendored:

| Skill | Use for |
| --- | --- |
| `tdd` | test-first implementation |
| `systematic-debugging` | root-cause work instead of symptom patching |
| `verification-before-completion` | proving a change actually works |
| `webapp-testing` | Playwright-driven frontend verification |
| `next-best-practices` | Next.js App Router patterns |
| `vercel-composition-patterns` | React component composition |
| `typescript-advanced-types` | non-trivial typing |
| `supabase-postgres-best-practices` | schema, index, and query strategy |
| `improve-codebase-architecture` | structural refactors |
| `grill-with-docs` | ADR and context capture |
| `to-prd` | turning an idea into a scoped work item |
| `triage` | deciding what is in scope |
| `zoom-out` | stepping back when a change is sprawling |
