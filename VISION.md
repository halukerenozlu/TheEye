# VISION.md

# TheEye Vision

## What TheEye Is

TheEye is a near real-time world monitoring platform built around a map-first interface and a unified event model.

Its purpose is to help users observe meaningful signals from multiple sources in one place without needing to manually jump between separate tools, feeds, and websites.

Early signal categories may include:

- earthquakes
- wildfires
- storms
- flights
- traffic incidents

Optional context may later include:

- reports
- alerts
- news references
- source confidence layers

---

## Why It Exists

The modern web is full of fragmented monitoring surfaces.
Important signals are often scattered across:

- separate APIs
- public feeds
- dashboards
- trackers
- news sources

TheEye exists to reduce that fragmentation.

Instead of forcing the user to monitor many disconnected systems, TheEye aims to collect, normalize, and present those signals through one coherent event-driven experience.

---

## Core Product Principles

### 1. Fast Initial Signal

A user should be able to open the product and quickly see meaningful activity.

### 2. Map-First Experience

Geographic context is primary, not secondary.

### 3. Event Normalization

Different sources should converge into one understandable event model.

### 4. Practical Real-Time

The MVP does not need perfect real-time.
Near real-time is acceptable if the system is reliable and useful.

### 5. Controlled Engineering

The project should grow through disciplined, reviewable steps rather than chaotic feature expansion.

### 6. Backend-Led Contract Stability

For the MVP, frontend implementation should follow the latest accepted backend contract so integration drift stays low.

---

## MVP Direction

The MVP is not intended to be everything.

The first useful version should:

- run locally in a stable way
- expose a clear backend API
- ingest at least one real source
- normalize that source into a common event shape
- show those events in a usable dashboard flow

---

## What TheEye Is Not

TheEye is not:

- a generic news website
- a social feed
- a massive all-in-one intelligence platform on day one
- a prematurely overengineered data platform
- a project that sacrifices product clarity for architecture complexity

---

## Long-Term Direction

Over time, TheEye can evolve toward:

- multiple normalized sources
- better geospatial querying
- saved filters and views
- alerts
- source confidence scoring
- richer event details
- stronger operational visibility

But long-term ambition should not break short-term focus.

---

## Working Standard

Every feature should support one of these:

- better signal visibility
- better event normalization
- better map-based exploration
- better operational reliability
- better user understanding

If a change does not improve one of these areas, it should be questioned before implementation.
