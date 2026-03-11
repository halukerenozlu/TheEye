# AGENTS.md — The EYE Repository Rules

> Read this file **first** before making any code changes.
> This repo is designed to be worked on by multiple coding agents (Claude Code, Gemini, ChatGPT Codex) and humans.

## 0) Project Summary

**World Pulse** is a real-time(ish) world monitoring dashboard:

- disasters (earthquakes, wildfires, storms…)
- live flights
- traffic (incidents/flow)
- optional context (reports/news)

Core UX: a single map view with layers + a live feed panel + time filters.

## 1) Product Goals (Must-Haves)

- **Fast initial signal**: map shows meaningful events within seconds.
- **Map-first**: viewport/bbox querying is the primary access pattern.
- **Normalized data model**: different sources → one `Event` schema.
- **Near real-time** is acceptable for MVP (30–120s latency).
- **Operational safety**: rate limiting, caching, reliable ingestion, observability hooks.

## 2) Non-Goals (Avoid)

- Big-bang rewrite or framework churn.
- Premature microservices split beyond: `api` + `collector` (+ dashboard).
- Heavy deployment refactors early. Keep local dev flow stable.
- Scraping sites without explicit permission. Prefer official APIs/feeds.

## 3) Tech Stack (Locked for MVP)

### Frontend

- Next.js (TypeScript)
- TailwindCSS + shadcn/ui
- MapLibre GL JS (WebGL map rendering)
- TanStack Query (data fetching/cache)
- Zustand (UI state: active layers, filters)

### Backend (API + Aggregation)

- Go (HTTP API)
- chi (router) or gin (router) — keep it minimal
- SSE (Server-Sent Events) for realtime MVP (WebSocket later if needed)
- PostgreSQL + PostGIS (geo queries)
- Redis (cache + rate limiting + stream fan-out)

### Dev/Infra

- Docker + Docker Compose (local parity)
- GitHub Actions (CI: lint/test/build)

## 4) Repository Layout (Target)

- apps/
  - dashboard/ # Next.js app
- services/
  - api/ # Go API (REST + SSE)
  - collector/ # Go ingestion workers/connectors
- infra/
  - docker-compose.yml
- shared/
  - schema/ # Event model + OpenAPI (or generated types)

Agents should NOT invent a different layout unless explicitly required.

## 5) Source-of-Truth Data Model: Event

All sources MUST normalize into the same `Event` structure.

Required fields:

- id (ULID/UUID)
- type (earthquake, wildfire, storm, flight, traffic_incident, …)
- title
- status (active|resolved|cancelled|unknown)
- severity (0..100)
- started_at, updated_at (ISO-8601)
- geometry (GeoJSON)
- source: { name, event_id, url, fetched_at }

Recommended fields:

- confidence (0..1)
- ended_at
- location (country_code/admin/place)
- tags (string[])
- metrics (type-specific numbers)
- raw (JSONB) (optional but useful for audits/debug)

DB uniqueness MUST be enforced by:

- UNIQUE (source_name, source_event_id)
  so collectors can UPSERT idempotently.

## 6) API Contract (MVP)

### Health

- GET /v1/healthz
- GET /v1/readyz
- GET /v1/meta

### Events

- GET /v1/events?bbox=...&since=...&types=...&severity_gte=...&status=...&limit=...&cursor=...
- GET /v1/events/{id}
- GET /v1/events/changes?since=...

### Realtime

- GET /v1/stream/events (SSE), supports filters similar to /events

Agents must keep endpoints stable once introduced.
If you need a breaking change, add v2 endpoints rather than breaking v1.

## 7) Local Development Rules (Do Not Break)

- `docker compose up` MUST start:
  - postgres (with PostGIS)
  - redis
  - api
  - (optional) collector
- Dashboard must run with a single command (`npm run dev` or `pnpm dev`)
- Avoid requiring paid keys for the base demo. Use public feeds where possible.

## 8) Coding Standards

### Go

- Keep packages small and cohesive.
- Prefer explicit error handling; no panics in request paths.
- Use context.Context for IO and request lifecycles.
- Structured logging (JSON) preferred.
- Add timeouts on outbound HTTP calls.
- Implement retries with backoff for ingestion connectors (bounded).

### TypeScript / Frontend

- Strict TypeScript. No `any` without justification.
- Keep map rendering performant:
  - clustering for points
  - avoid re-render loops
  - query only on viewport change (debounce)

### Dependencies

- Add new deps only if they materially reduce complexity.
- Avoid “framework hopping”.

## 9) Performance & Reliability (Minimum Bar)

- bbox queries must use PostGIS + GIST indexes.
- cache hot paths in Redis (short TTL is fine).
- rate limit outbound source calls and inbound API calls.
- collectors must be safe under restarts (idempotent upsert).
- avoid fetching “worldwide flights” indiscriminately; always filter by bbox.

## 10) Security & Secrets

- Never commit API keys.
- Use `.env` locally; provide `.env.example`.
- Validate and sanitize query params (bbox, since, types, limits).

## 11) Testing & CI Expectations

- Go: unit tests for connectors normalization + API handlers.
- Frontend: basic component test optional; critical flows E2E later.
- CI (GitHub Actions) should run:
  - Go fmt + go test
  - TS typecheck + lint
  - build (best-effort)

## 12) Git Workflow

- Branches: `main` (protected), feature branches per unit of work.
- Commits: Conventional Commits recommended:
  - feat:, fix:, chore:, refactor:, docs:
- PRs must include:
  - what changed
  - how to test locally
  - screenshots/gifs for UI changes when relevant

## 13) Multi-Agent Coordination Protocol

When an agent works on a task, it must:

1. State assumptions and scope in 3–7 bullets.
2. Make small, reviewable changes (prefer incremental PRs).
3. Update docs/config if behavior changes.
4. Avoid silent breaking changes.
5. Output exact commands to run for verification.

If multiple agents are used concurrently:

- designate ONE “integrator” (human or agent) to merge changes.
- avoid overlapping edits in the same files.
- define interfaces first (Event schema + OpenAPI), then implement.

## 14) Definition of Done (DoD)

A task is done when:

- code compiles
- relevant tests pass
- local run steps are documented
- no breaking changes to local dev flow
- endpoints and Event schema remain consistent

## 15) MVP Milestones (Suggested)

M1: docker-compose + DB schema + /healthz + /events (empty OK)
M2: USGS earthquakes connector → events appear on map
M3: SSE stream + dashboard live updates
M4: add EONET/GDACS + basic filtering + caching
M5: traffic/flights with bbox-only queries and strict rate limiting
