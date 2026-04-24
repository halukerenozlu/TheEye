# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial Command Center Lite dashboard shell with top bar, feed panel, center map area, and detail panel.
- Map rendering in the center panel using MapLibre GL with geometry-based event markers.

### Changed

- Signal Feed wired to real `/v1/events` API data with backend-supported filter and sort behavior.
- Event selection synchronized across feed, map markers, and Event Intelligence detail panel.
- Loading, empty, and error states refined for clearer operator-facing feedback.
- Feed/detail visual consistency improved while preserving restrained dark design language.
- Category/filter control visual inconsistencies cleaned up where applicable.
- Final Sprint 1 verification and documentation sync completed for the accepted Phase 5 dashboard scope.
- Phase 5 dashboard milestone documented as completed, with Phase 6 indicated as the next phase.

### Fixed

- Map viewport sizing and camera behavior improved so basemap and markers render reliably in the active map area.
- Local dashboard/API development readiness improved for geometry-aware event rendering and stable integration flow.
- Temporary dashboard debug logging and transient local verification artifacts were cleaned up.
