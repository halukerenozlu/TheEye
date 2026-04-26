# TheEye Roadmap

## Direction

TheEye is a map-first global signal platform that collects, normalizes, and presents meaningful world activity through one event-driven product.

The current baseline proves a natural/physical signal workflow through USGS earthquakes. Future versions can broaden toward additional natural sources first, then carefully evaluated context enrichment and, later, human-system or global-stability signals.

The roadmap is now organized by **Version Milestones**. Detailed tasks belong in `docs/VERSION_PLAN.md`; this file stays high-level.

---

## v0.1 - Initial Working MVP

Status: Completed

v0.1 establishes the first useful end-to-end baseline:

- documented map-first global signal vision
- Go API service skeleton
- health, readiness, and metadata endpoints
- normalized Event response model baseline
- consistent JSON error shape
- USGS earthquake ingestion
- deterministic normalization
- duplicate-safe persistence
- idempotency through `source_name + source_event_id`
- real stored event list and detail endpoints
- filter, sort, and pagination support
- first dashboard shell
- Signal Feed connected to `/v1/events`
- MapLibre rendering and geometry-based markers
- feed/map/detail selection sync
- polling and freshness visibility
- backend-owned `severity_level`
- normalized `category`
- Level 1 / Level 2 / Level 3 visual distinction

---

## v0.2 - Multi-source Event Foundation

Status: Next

v0.2 should add the second real source and make the ingestion/event model foundation source-independent enough to support a mixed feed and mixed map without breaking USGS.

Primary direction:

- second natural/physical source
- source adapter boundary
- multi-source-compatible Event behavior
- clear category/type/source metadata
- mixed feed/map baseline

---

## v0.3 - Context and Enrichment Foundation

Status: Planned

v0.3 should define how event details can gain useful context without confusing enrichment with core ingestion.

Primary direction:

- MCP or similar context integration evaluation
- enrichment boundary
- event detail context planning
- API stability while context work matures

---

## v0.4 - Map-first Mixed Feed Experience

Status: Planned

v0.4 should improve the user-facing experience for mixed-source monitoring.

Primary direction:

- source/category/type filters
- mixed feed readability
- marker behavior improvements
- selected event detail improvements
- freshness visibility
- loading/empty/error polish

---

## v0.5 - Reliability and Signal Quality

Status: Planned

v0.5 should strengthen trust in source behavior, event freshness, duplicate handling, and prioritization.

Primary direction:

- source confidence model draft
- duplicate and near-duplicate handling
- event freshness rules
- severity/priority refinement
- ingestion observability/logging baseline

---

## Roadmap Governance

Long-term vision lives in `docs/VISION.md`.

Active milestone planning and detailed work items live in `docs/VERSION_PLAN.md`.

The roadmap should not become a task tracker. It should explain sequence, intent, and product direction at a version milestone level.
