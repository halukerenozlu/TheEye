# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

Future changes will be documented here.

## [0.1.0] - Initial Working MVP

### Added

- Initial map-first global signal platform documentation baseline.
- Go API service skeleton.
- `GET /v1/healthz`, `GET /v1/readyz`, and `GET /v1/meta`.
- Typed Event response model baseline.
- Consistent JSON error shape baseline.
- USGS earthquakes as the first real source.
- USGS fetch client and source payload decoding.
- Deterministic normalization for USGS records.
- Duplicate-safe persistence using `source_name + source_event_id` idempotency.
- Real stored event data through `GET /v1/events`.
- Real event detail data through `GET /v1/events/{id}`.
- Filter, sort, and pagination support for `GET /v1/events`.
- Initial Command Center Lite dashboard shell with top bar, feed panel, center map area, and detail panel.
- Signal Feed wired to real `/v1/events` API data.
- Map rendering in the center panel using MapLibre GL with geometry-based event markers.
- Feed, map, and detail panel selection sync.
- Backend-owned `severity_level`.
- Normalized `category`.
- Simple client-side polling for MVP freshness.
- Collector service in Docker Compose with periodic USGS ingest.

### Changed

- TheEye planning model migrated from Phase/Sprint planning to Version Milestones.
- Product framing clarified as a map-first global signal platform rather than only a natural disaster dashboard.
- Event selection synchronized across feed, map markers, and Event Intelligence detail panel.
- Loading, empty, and error states refined for clearer operator-facing feedback.
- Feed/detail visual consistency improved while preserving restrained dark design language.
- Category/filter control visual inconsistencies cleaned up where applicable.
- Type/category separation documented for future multi-source behavior.
- Polling and freshness visibility improved in the dashboard flow.
- Level 1 / Level 2 / Level 3 visual distinction improved across feed, map, and detail views.

### Fixed

- Map viewport sizing and camera behavior improved so basemap and markers render reliably in the active map area.
- Local dashboard/API development readiness improved for geometry-aware event rendering and stable integration flow.
- Temporary dashboard debug logging and transient local verification artifacts were cleaned up.

### Documentation

- Completed baseline scope consolidated under `docs/VERSION_PLAN.md`.
- Phase/Sprint planning references replaced with Version Milestone / Work Item / Implementation Slice terminology.
- Roadmap rewritten as a high-level version milestone roadmap.
- Architecture, API, and DB docs added for the current baseline.
