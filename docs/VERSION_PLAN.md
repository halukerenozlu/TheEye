# TheEye Version Plan

## Current Version

- Current completed baseline: v0.1.0 - Initial Working MVP
- Active planning model: Version Milestones
- Phase/Sprint planning is no longer used.

---

## Versioning Rules

- `vMAJOR.MINOR.PATCH`
- `MINOR`: meaningful product milestone
- `PATCH`: bug fix, doc sync, small stabilization
- Tags are milestone checkpoints, not routine progress markers.

### MAJOR

Increase `MAJOR` for major product direction changes, public release maturity shifts, or deliberately breaking architectural changes.

### MINOR

Increase `MINOR` when a meaningful product milestone is completed and intentionally checkpointed.

### PATCH

Increase `PATCH` for small fixes, documentation sync, review follow-up, and stabilization that does not change milestone scope.

### Tag Discipline

Annotated tags are preferred when tagging a milestone:

```bash
git tag -a v0.1.0 -m "v0.1.0: initial working MVP"
git push origin v0.1.0
```

Before tagging:

- implementation should be reviewed at the right risk level
- local Docker-based development flow should still work
- frontend/backend integration should be stable if the boundary changed
- key docs should be synced
- root `VERSION` should match the intended Git tag
- generated UI version files should be synced with `pnpm version:sync`
- the intended tag should pass `pnpm version:check vMAJOR.MINOR.PATCH`
- the working tree should be clean

---

## v0.1.0 - Initial Working MVP

Status: Completed

v0.1.0 is the completed initial working baseline for TheEye. It consolidates the completed planning, backend, ingestion, API, dashboard, and reliability work that previously lived across Phase 1 through Phase 6 Sprint 1.

This version proves the first useful end-to-end loop:

```text
USGS source -> fetch -> decode -> normalize -> persist -> API -> dashboard feed/map/detail
```

### Product and Documentation Foundation

Completed scope:

- map-first global signal platform vision clarified
- long-term direction framed beyond a natural disaster dashboard
- natural and physical events established as the first practical signal family
- future human systems, global stability, and critical world signals kept as controlled long-term expansion paths
- MVP boundaries documented around ingestion, normalization, storage, API, and dashboard behavior
- multi-agent responsibilities documented for Human, ChatGPT, Codex, Gemini, and Claude Code
- backend-first integration discipline established to prevent frontend/backend contract drift
- local Docker-based development flow preserved as a core constraint
- normalized Event-oriented product model established as the center of the system

### Backend Foundation

Completed scope:

- Go API service skeleton created
- minimal HTTP service structure established
- structured configuration baseline added
- graceful shutdown behavior added
- `GET /v1/healthz` added
- `GET /v1/readyz` added
- `GET /v1/meta` added
- initial `GET /v1/events` placeholder added with stable response shape
- initial `GET /v1/events/{id}` placeholder added
- typed Event response model started
- consistent JSON error shape baseline created
- router-level JSON errors added for not found and method-not-allowed cases
- backend tests added for skeleton, response, and error behavior

### First Ingestion Pipeline

Completed scope:

- USGS earthquakes selected as the first real source
- USGS GeoJSON `FeatureCollection` source contract documented
- USGS fetch client added
- source payload decode implemented
- non-200 upstream handling added
- malformed JSON handling added
- deterministic normalization added
- normalized event ids use the `usgs:{source_id}` baseline
- type/title/status mapping added
- UTC time conversion added
- deterministic severity mapping added
- duplicate-safe persistence added
- batch-level deduplication before writes added
- database-level conflict-safe upsert behavior added
- `source_name + source_event_id` idempotency logic established
- minimal schema creation path added for the current baseline
- current USGS records are stored as normalized events
- ingestion tests added around fetch, decode, normalization, persistence, duplicate safety, invalid input, and DB error handling

### API Layer

Completed scope:

- `/v1/events` now returns real stored event data
- `/v1/events/{id}` now returns real stored event detail data
- unknown event ids keep the consistent JSON `404` shape
- events list response shape remains:
  - `items`
  - `next_cursor`
- optional filtering added for:
  - `type`
  - `started_after`
  - `started_before`
- invalid query handling returns consistent JSON `400` errors
- sorting support added for:
  - `updated_at_desc`
  - `updated_at_asc`
- pagination support added with:
  - `limit`
  - `cursor`
  - empty string `next_cursor` when no next page exists
- invalid cursor usage returns consistent JSON `400` errors
- backend-owned `severity_level` added
- normalized `category` added
- `type` and `category` distinction documented
- current USGS events aligned to `category = natural_disaster`
- API tests added/updated for list, detail, filter, sort, pagination, category, and severity behavior

### First Dashboard

Completed scope:

- dashboard shell created
- top bar added
- left Signal Feed panel added
- center map area added
- right Event Intelligence detail panel added
- restrained dark visual hierarchy established
- Signal Feed wired to real `/v1/events` API data
- backend-supported filters wired into the dashboard
- MapLibre map rendering added
- event geometry used to render markers
- map viewport adjusts to visible event bounds
- events without geometry are handled safely
- feed item selection added
- marker selection added
- right-side detail panel shows selected real event data
- feed/map/detail selection sync added
- map focus/ease behavior added for selected events where available
- loading, empty, and error states improved
- selected-state clarity improved
- feed/detail visual consistency improved
- temporary debug logging and local verification artifacts cleaned up
- frontend build and runtime dashboard behavior were verified for the accepted baseline

### Reliability and Contract Foundation

Completed scope:

- backend-owned severity normalization established
- `severity_level` exposed as the frontend-facing severity level baseline
- legacy `severity` kept for compatibility in the current baseline
- normalized `category` field introduced
- category/type distinction made explicit for future multi-source work
- simple client-side polling added for MVP freshness
- collector service added to Docker Compose flow
- collector starts automatically in compose
- periodic USGS ingest works
- live USGS data flows into the database
- `/v1/events` returns current records with `category` and `severity_level`
- polling and freshness visibility improved in the dashboard
- Level 1 / Level 2 / Level 3 visual distinction improved across feed, map, and detail views
- USGS live flow stabilized for the current single-source scope

### Completion Criteria

v0.1.0 is complete because:

- the product has a clear map-first global signal platform direction
- the local Docker-backed stack has a working API and collector baseline
- one real source, USGS earthquakes, is ingested end to end
- source data is decoded, normalized, stored, and exposed through the API
- duplicate-safe persistence exists through `source_name + source_event_id`
- `/v1/events` and `/v1/events/{id}` return real stored records
- list filtering, sorting, and pagination exist for the MVP API layer
- the dashboard renders real event data in feed, map, and detail surfaces
- MapLibre markers render from event geometry
- selection state is synchronized across feed, map, and detail
- freshness polling works for the MVP baseline
- severity/category contract decisions are explicit enough for the next source
- the baseline remains scoped, reviewable, and ready for stabilization or multi-source expansion

---

## v0.1.x - Stabilization and Documentation Cleanup

Purpose:

- small corrections after the version milestone migration
- docs alignment
- README clarity
- local run verification
- small bug fixes / UI polish

Potential work items:

- verify all docs point to the Version Milestone model
- confirm local Docker startup commands still match implementation reality
- tighten API examples and local smoke-test notes
- fix small dashboard polish issues discovered during smoke testing
- avoid adding new sources unless the work is explicitly moved into v0.2.0

---

## v0.2.0 - Multi-source Event Foundation

Purpose:

- add a second real event source
- establish normalized event source architecture
- prepare the foundation for mixed feed/map behavior

Main items:

- choose a second natural/physical source
- define a source adapter boundary
- make the normalized Event model source-independent
- clarify category/type/source metadata behavior
- keep `/v1/events` multi-source compatible
- support mixed source behavior in feed and map
- move forward without breaking the existing USGS flow

Guardrails:

- do not redesign the whole ingestion system
- do not introduce human-system signals before the natural/physical multi-source foundation is stable
- do not let frontend invent new backend fields

---

## v0.3.0 - Context and Enrichment Foundation

Purpose:

- prepare the foundation for event detail and context enrichment
- evaluate MCP or similar context integration approaches

Main items:

- document the purpose of MCP usage for TheEye
- define the enrichment boundary
- plan context fields for the event detail panel
- avoid mixing external context sources directly into core ingestion
- preserve current event API stability

Guardrails:

- enrichment should not become a hidden second ingestion pipeline
- context should be clearly separated from normalized source-of-truth event data
- speculative AI summaries should not enter the product without explicit approval

---

## v0.4.0 - Map-first Mixed Feed Experience

Purpose:

- present multi-source events through a more readable map/feed workflow

Main items:

- source/category/type filters
- mixed feed improvements
- marker behavior improvements
- selected event detail improvements
- freshness visibility improvements
- empty/loading/error state polish

Guardrails:

- preserve map-first interaction
- avoid heavy dashboard clutter
- avoid broad UI redesign unless explicitly scoped

---

## v0.5.0 - Reliability and Signal Quality

Purpose:

- establish the foundation for source reliability, freshness, duplicate handling, and signal quality

Main items:

- source confidence model draft
- duplicate/near-duplicate handling
- event freshness rules
- severity/priority model refinement
- ingestion observability/logging baseline

Guardrails:

- keep reliability improvements tied to real source behavior
- avoid overbuilding a generalized scoring system before source behavior is understood
- preserve existing API and dashboard behavior unless a change is intentionally versioned
