# PHASE_06.md

# Phase 6 - Signal Expansion Foundation

## Status

**Active**

Current sprint:

- Sprint 1 - Reliability and Contract Foundation (Completed)

Current active step:

- None (Sprint 1 completed)

---

## Goal

Phase 6 moves TheEye from a working single-source dashboard MVP toward a reliable multi-signal foundation.

The primary goal is to stabilize the current single-source pipeline and dashboard flow first, then prepare the product for one additional real data source without broadening scope into a larger platform redesign.

---

## Why This Phase Is Needed

Phase 5 proved that the dashboard MVP works end to end.

Phase 6 is needed because the product should not jump from a fragile single-source slice directly into broader signal expansion. The current ingestion, API, and UI flow must become more reliable and more explicit before a second real source is added.

This phase also clarifies the next stable contract decisions so frontend behavior does not depend on source-specific raw fields.

---

## Core Direction

- keep the current single-source pipeline and dashboard flow stable first
- move severity normalization to the backend
- decide the category model before multi-signal expansion
- use simple client-side polling for MVP freshness
- add only one second real source in this phase
- treat mixed feed + mixed map as the overall direction for the phase

---

## Contract Decisions For Phase 6

### Severity

Severity normalization is a backend responsibility in this phase.

Accepted direction:

- API responses should return `severity_level` as `1 | 2 | 3`
- frontend should not derive severity level from raw magnitude or other source-specific fields
- severity display logic should consume the normalized backend field directly

### Category

Category will be introduced as a separate normalized field rather than overloading `type`.

Accepted direction:

- `type` should continue to represent the specific event type
- `category` should represent the broader grouping used across multiple signals
- category/type expectations must be documented before the second source is integrated

This keeps multi-signal filtering and feed/map grouping more stable as additional sources are added.

### Refresh Model

Accepted direction:

- simple client-side polling is sufficient for MVP freshness in Phase 6
- SSE is intentionally deferred to Phase 7 or later

---

## Step 1 Completion - Severity backend normalization and API contract decision

Status: Completed

Delivered:

- backend severity normalization implemented for the current source flow
- API responses now include `severity_level`
- legacy `severity` was kept temporarily for compatibility
- minimal backend tests were updated for the accepted contract change

## Step 2 Completion - Category field decision and docs/spec alignment

Status: Completed

Delivered:

- `category` was added to the backend contract
- the distinction between `type` and `category` was kept explicit
- current USGS events were aligned to `category = natural_disaster`
- relevant backend tests were updated

## Step 3 Completion - Simple client-side polling

Status: Completed

Delivered:

- frontend polling was added
- a collector service was added to compose
- collector now starts automatically and runs periodic ingest
- live USGS data now flows into the DB
- `/v1/events` now returns current records with `category` and `severity_level`

## Step 4 Completion - Level 2 / Level 3 visual distinction cleanup

Status: Completed

Delivered:

- polling and freshness visibility are now clear in the dashboard flow
- Level 1 / Level 2 / Level 3 visual distinction is now clearer across feed, map, and detail views
- the restrained dark layout and existing interaction model were preserved

## Step 5 Completion - Verification and sprint closure docs sync

Status: Completed

Delivered:

- USGS live flow is now reliable for the current single-source scope
- collector runs automatically in compose with periodic ingest
- polling and freshness visibility are in place
- severity/category contract baseline is stabilized for frontend use
- Level 1 / Level 2 / Level 3 visual separation is now clearer

---

## Phase Scope

In scope:

- reliability improvements for the current single-source ingestion to API to dashboard flow
- backend normalization of `severity_level`
- explicit category model decision and doc/spec alignment
- simple client-side polling for dashboard freshness
- minimal UI reliability/polish to make Level 2 and Level 3 states more clearly distinguishable
- preparation for exactly one additional real data source
- mixed feed + mixed map as the product direction for accepted Phase 6 work

Out of scope:

- clustering
- hover-heavy interaction patterns
- large UI redesign
- auth
- notifications
- timeline views
- SSE implementation in this phase
- adding more than one new real source

---

## Exit Criteria

Phase 6 is complete when:

- the single-source MVP flow is more reliable across ingestion, API, and dashboard usage
- API severity normalization is backend-owned and stable for frontend consumption
- category handling is explicitly decided and documented
- dashboard freshness uses acceptable MVP polling without introducing SSE complexity
- Level 2 and Level 3 visual distinction is clear without redesigning the dashboard
- one second real source can be integrated on the accepted normalized foundation
- mixed feed + mixed map behavior is achievable within the accepted contract direction
- scope remains controlled without UX or platform creep

---

## Version Note

Target version candidate for this phase may be:

- `v0.2.0`

This is a planning marker only.
No tag or release work is part of this task, and sprint progress is not a version milestone by itself.

---

## Next Sprint

- Sprint 2 - To be defined
