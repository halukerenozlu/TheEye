# VISION.md

# TheEye Vision

## What TheEye Is

TheEye is a map-first global signal platform designed to help users observe meaningful world activity through a unified event-driven interface.

Its long-term direction is not to become a simple natural disaster dashboard, but a layered world monitoring product that helps users understand critical signals, system stress, and global instability with geographic context.

In the near term, TheEye is built by collecting, normalizing, and presenting reliable multi-source events through one coherent monitoring experience.

---

## What TheEye Is Not

TheEye is not:

- a generic news website
- a social feed or discussion platform
- a raw data wall that shows everything equally
- a consumer travel assistant
- a massive all-in-one intelligence platform on day one
- a prematurely overengineered analytics system
- a project that uses long-term ambition to justify uncontrolled scope expansion

---

## North Star vs Current Product Reality

TheEye has a long-term north star and a near-term delivery reality.

### Long-Term North Star

TheEye can evolve toward a map-first intelligence platform focused on:

- human systems
- global stability
- critical infrastructure signals
- geopolitical and operational disruption
- high-value event correlation across domains

Examples of future signal families may include:

- flights
- internet outages
- cyber incidents
- market or infrastructure disruptions
- security or instability signals

### Current Product Reality

The current product should remain grounded in reliable event ingestion, normalization, and map-based monitoring.

For the near-term roadmap, TheEye should first become a trustworthy multi-source world signal platform before attempting country intelligence, profiling, or cross-signal reasoning.

Long-term vision should guide direction, but it must not override phase discipline.

---

## Why It Exists

The modern web is full of fragmented monitoring surfaces.

Important signals are scattered across:

- separate APIs
- public feeds
- dashboards
- trackers
- reference sites
- news and alert surfaces

TheEye exists to reduce that fragmentation.

Its goal is to collect, normalize, and present meaningful signals through one readable interface so users can understand what is happening, where it is happening, and why it matters.

---

## Core Product Principles

### 1. Fast Initial Signal

A user should be able to open the product and quickly see meaningful activity.

### 2. Map-First Experience

Geographic context is primary, not secondary.

### 3. Signal Over Noise

TheEye should not show everything equally.
It should favor meaningful signals, usable prioritization, and readable context over raw volume.

### 4. Event Normalization

Different sources should converge into one understandable event model wherever practical.

### 5. Practical Real-Time

The MVP does not require perfect real-time.
Near real-time is acceptable if the system is reliable, refreshable, and operationally trustworthy.

### 6. Controlled Engineering

The project should grow through disciplined, reviewable steps rather than chaotic feature expansion.

### 7. Backend-Led Contract Stability

Frontend implementation should follow the latest accepted backend contract so integration drift remains low.

### 8. Reliability Before Breadth

Before adding new signal families, TheEye should prove that ingestion, normalization, storage, and display are dependable.

---

## Product Center of Gravity

TheEye should be built around three persistent questions:

- **What is happening?**
- **Where is it happening?**
- **Why does it matter?**

The product surface should support that structure:

- **Map (center):** where the signal is happening
- **Feed (left):** what is happening and in what order
- **Detail (right):** why the selected signal matters

This right-side panel may evolve over time, but it should do so gradually:

1. selected event detail
2. selected event + nearby or related active context
3. limited country or regional context
4. much later, richer intelligence or profiling layers

TheEye should not jump directly from event detail into a full country intelligence product too early.

---

## Severity and Importance Philosophy

TheEye should not assume that every source can share one identical severity rule.

Near-term implementation should prefer **source-native severity**:

- earthquakes may use magnitude-based severity
- fires may use area or containment-based severity
- other sources may require different native logic

A broader cross-source importance model may exist later, but it should not be forced too early.

The product should first support:

- stable source-native severity
- readable prioritization
- source-aware interpretation

Cross-category importance and intelligence scoring belong to a later maturity stage.

---

## MVP and Near-Term Direction

The MVP is not intended to be everything.

The first useful versions should:

- run locally in a stable way
- expose a clear backend API
- ingest real sources reliably
- normalize those sources into a common event shape
- show those events through a usable dashboard flow
- refresh data in a predictable operational way
- keep the user experience restrained and readable

Near-term expansion should remain disciplined.

For the next phases, TheEye should prioritize:

- multi-source event reliability
- source/category clarity
- stable ingestion pipelines
- trustworthy map + feed + detail behavior
- minimal but strong UX improvements

Near-term scope should not jump into:

- country intelligence profiling
- cross-signal reasoning engines
- speculative AI analysis
- broad economic or cyber modeling
- heavy archival analytics

---

## Natural Events vs Human Systems

Natural and physical events remain valid and useful in TheEye.

However, over the long term, they are best understood as one signal family within a broader world monitoring vision.

In the near term, natural events may remain the primary operational focus because they are more structured, accessible, and practical for building a reliable multi-source platform.

Human systems should enter later, carefully, and only when:

- reliable source access exists
- normalization rules are clear
- ingestion is repeatable
- the product can absorb the new complexity without losing clarity

---

## Long-Term Direction

Over time, TheEye can evolve toward:

- multiple normalized sources
- stronger source/category filtering
- better geospatial querying
- richer event details
- saved filters and views
- alerting and notification layers
- source confidence support
- limited regional context
- eventually, carefully scoped human-system signals

But long-term ambition must not break short-term focus.

---

## Working Standard

Every feature should clearly improve at least one of these areas:

- better signal visibility
- better signal prioritization
- better event normalization
- better map-based exploration
- better operational reliability
- better user understanding

If a change does not improve one of these areas, it should be questioned before implementation.

---

## Vision Governance

This document defines long-term product direction.

It is a guiding north star, not a license for uncontrolled implementation.

Active phase and sprint documents remain the operational source of truth for what is built now.

TheEye should protect vision **and** discipline at the same time.
