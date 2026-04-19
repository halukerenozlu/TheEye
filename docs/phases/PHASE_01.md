# PHASE_01.md

# Phase 1 — Product Definition and Domain Model

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

TheEye is a near real-time monitoring dashboard that aggregates multiple world-signal sources into a unified event-centric system.

Early examples include:

- earthquakes
- wildfires
- storms
- flights
- traffic incidents
- optional contextual reports/news later

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
- overdesigned domain abstractions before first working slice

---

## Domain Direction

All sources should normalize into a common `Event` model.

Early response/model fields discussed so far:

- id
- type
- title
- status
- severity
- started_at
- updated_at

This is a minimal draft, not the final schema.

Expected future additions:

- geometry
- source metadata
- location
- tags
- confidence
- metrics
- raw payload for debugging/auditing

---

## Exit Criteria

Phase 1 is considered complete when:

- vision is written and stable
- MVP boundaries are documented
- event-driven domain direction is agreed
- roadmap phases are defined
- implementation can proceed without product ambiguity

---

## Notes

Phase 1 is primarily documentation and alignment work.
Its value is strategic clarity, not feature count.
