# PHASE_01.md

# Phase 1 — Product Definition and Domain Model

## Status

**Completed in substance**

This phase can be treated as closed unless new product-direction decisions require reopening it.

---

## Purpose

This phase defines what TheEye is, what it is not, and what the MVP must support before deeper backend and data integration work.

This phase exists to reduce ambiguity before the system grows.

---

## Target Outcome

By the end of Phase 1, the project should have:

- a clear product vision
- clear MVP boundaries
- clear non-goals
- a normalized Event-oriented domain direction
- an implementation roadmap that does not drift

---

## Main Questions This Phase Answers

- What exactly is TheEye?
- Who is it for?
- What is the MVP?
- What is explicitly out of scope?
- What is the core entity model?
- What is the first practical product slice?

---

## Working Product Definition

TheEye is a near real-time monitoring platform that aggregates multiple world-signal sources into a unified event-centric system.

Early examples include:

- earthquakes
- wildfires
- storms
- flights
- traffic incidents
- optional contextual reports or news later

The primary UX is:

- map-first exploration
- live feed panel
- time-based filtering
- source normalization into a common event shape

---

## Must-Haves for MVP

- fast initial signal
- map-first access pattern
- normalized event model
- near real-time behavior
- local-first development flow
- reliable ingestion direction
- safe operational baseline

---

## Non-Goals in Early Phase

- big-bang architecture rewrite
- premature microservices fragmentation
- heavy deployment refactors
- scraping unofficial or unsafe sources
- overdesigned domain abstractions before the first working slice

---

## Domain Direction

All sources should normalize into a common `Event` model.

Early core fields:

- `id`
- `type`
- `title`
- `status`
- `severity`
- `started_at`
- `updated_at`

Expected future additions:

- `geometry`
- `source`
- `location`
- `tags`
- `confidence`
- `metrics`
- `raw`

---

## Delivered Phase Outputs

This phase is reflected through the current documentation baseline, especially:

- `VISION.md`
- `ROADMAP.md`
- `AGENTS.md`
- phase and sprint planning documents
- event-driven product framing

---

## Exit Criteria

Phase 1 is considered complete when:

- vision is written and stable
- MVP boundaries are documented
- event-driven domain direction is agreed
- roadmap phases are defined
- implementation can proceed without product ambiguity

These conditions are satisfied in the current documentation set.

---

## Notes

Phase 1 is documentation and alignment work.
Its value is strategic clarity, not feature count.
