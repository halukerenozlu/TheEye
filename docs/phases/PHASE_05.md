# PHASE_05.md

# Phase 5 - First Dashboard

## Status

**Active**

Current sprint:

- Sprint 1 - Command Center Lite Dashboard

Current active step:

- Step 5 - Basic visual polish for loading/empty/error states

---

## Purpose

Phase 5 delivers the first usable, map-first dashboard on top of the accepted Phase 4 API behavior.

The target is a restrained "Command Center Lite" UI that is practical, calm, and reviewable for MVP.

---

## Dashboard Direction

- dark, calm visual baseline
- top bar
- left event/feed panel
- center map
- right event detail panel
- click-first interaction over hover-heavy behavior

---

## Backend Baseline for Phase 5

Available API behavior for dashboard integration:

- `GET /v1/healthz`
- `GET /v1/readyz`
- `GET /v1/meta`
- `GET /v1/events`
- `GET /v1/events/{id}`
- filtering, sorting, and pagination support
- map-ready geometry exposure when present in stored records

---

## Step 1 Completion - Dashboard Shell Layout

Status: Completed

Delivered:

- dashboard shell layout created
- top bar, left feed panel, center map area, and right detail panel scaffolded
- restrained dark visual hierarchy established
- loading/empty placeholders added where relevant

Not added in Step 1:

- real API data wiring
- map library integration
- advanced interaction logic

## Step 2 Completion - Event Feed + Basic Filters Wiring

Status: Completed

Delivered:

- left Signal Feed panel wired to real `/v1/events` data
- basic backend-supported filters wired
- loading / empty / error handling integrated into the existing dashboard shell
- design language and shell layout preserved
- right detail panel remains mostly scaffolded
- map area remains placeholder for the next step

Not added in Step 2:

- full map rendering
- full right-panel detail wiring
- backend contract expansion beyond the already accepted readiness patch

## Step 3 Completion - Map Marker Rendering From Current API Data

Status: Completed

Delivered:

- MapLibre GL integrated into the center map area
- dark basemap visibly renders in the dashboard map panel
- markers render for events with geometry
- map viewport adjusts to visible event bounds
- events without geometry are handled safely without UI breakage
- restrained dark map styling preserved

Not added in Step 3:

- full right-panel detail wiring
- clustering or heavy map effects
- country interaction logic
- backend contract redesign

## Step 4 Completion - Event Selection And Right-Side Detail Panel

Status: Completed

Delivered:

- feed item click selects the event
- marker click selects the event
- right-side Event Intelligence panel shows real selected-event data
- selected state is synchronized across feed, map, and detail panel
- map can focus/ease toward the selected event
- clear/reset selection behavior exists where implemented

Not added in Step 4:

- country interaction logic
- clustering
- backend contract redesign

---

## Phase Scope

In scope:

- first real dashboard shell and event-centered interaction
- real API data wiring to feed, map, and detail views
- loading, empty, and error state clarity for the MVP dashboard
- restrained visual system suitable for longer use
- small backend follow-up support only if strictly needed for frontend integration

Out of scope:

- webcam-style live media walls
- dense command-center clutter or AI insight overlays
- broad frontend architecture rewrites
- heavy backend redesign or new platform services
- country-first interaction as a hard requirement

---

## Working Mode

- Gemini-heavy for frontend implementation in this phase
- Codex only for small backend support patches when strictly necessary

---

## Exit Criteria

Phase 5 is complete when:

- a usable map-first dashboard is available with the agreed shell layout
- event feed, map markers, and detail panel work together on real API data
- core interactions are stable and reviewable for MVP
- loading, empty, and error states are intentionally handled
- scope remains restrained without product-direction drift

---

## Next Phase

- Phase 6 - Better UX and Product Depth
