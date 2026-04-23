# PHASE_05_SPRINT_01.md

# Sprint 1 - Command Center Lite Dashboard

## Phase

Phase 5 - First Dashboard

## Status

**Active**

Current active step:

- Step 3 - Map marker rendering from current API data

---

## Goal

Deliver the first usable dashboard slice with map-first interaction using current real API data.

---

## Scope

### In scope

- restrained Command Center Lite shell (top bar + left feed + center map + right detail)
- event-centered interaction using existing API routes
- map marker rendering from available geometry
- selection flow between feed, map, and detail panel
- basic visual polish for loading, empty, and error states
- minimal tests and docs sync for accepted sprint behavior

### Out of scope

- backend contract invention beyond tiny follow-up support if strictly necessary
- broad frontend architecture refactor
- heavy command-center visual density
- webcam/live news wall modules
- AI insight overlays
- mandatory country-focused interaction

---

## Step Breakdown

### Step 1 - Dashboard shell layout

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

### Step 2 - Event feed + basic filters wiring

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

### Step 3 - Map marker rendering from current API data

Delivered when:

- markers render from available event geometry
- events without geometry are handled safely without UI breakage

### Step 4 - Event selection and right-side detail panel

Delivered when:

- selecting an event from feed/map drives detail panel updates
- map can focus/center on selected event when geometry is available

### Step 5 - Basic visual polish for loading/empty/error states

Delivered when:

- loading, empty, and error states are clear and calm
- limited category color usage remains readable on dark theme

### Step 6 - Tests and docs sync

Delivered when:

- minimal relevant tests are updated for accepted frontend/backend integration behavior
- sprint and phase docs reflect accepted implementation state

---

## Sprint Exit Criteria

Sprint 1 is complete when:

- shell, feed, map markers, and detail panel work as one flow
- real API data drives the core dashboard interactions
- click-first interaction is clear and stable
- MVP scope remains restrained and reviewable
