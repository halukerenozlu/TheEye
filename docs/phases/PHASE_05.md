# PHASE_05.md

# Phase 5 - First Dashboard

## Status

**Active**

Current sprint:

- Sprint 1 - Command Center Lite Dashboard

Current active step:

- Step 1 - Dashboard shell layout

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

