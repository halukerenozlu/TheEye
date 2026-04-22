# WORKFLOW.md

## Purpose

This document defines the standard development workflow for TheEye.

TheEye is developed through a controlled, role-separated process using:

- Human direction
- ChatGPT for planning and prompt generation
- Codex for primary implementation
- Gemini for frontend direction and integration-aware UI work
- Claude Code for selective review

The goal is disciplined, reviewable progress rather than freeform code generation.

---

## Core Principle

Every implementation task must belong to a defined structure:

- Phase
- Sprint
- Step

No coding work should begin unless the current step is explicitly identified.

---

## Roles

### Human

Responsible for:

- product direction
- priorities
- tradeoff decisions
- approval or rejection
- final commit and tag decisions

### ChatGPT

Responsible for:

- roadmap, phase, sprint, and step definition
- architecture framing
- Codex prompt generation
- Gemini prompt generation
- Claude Code review prompt generation
- review interpretation
- commit and tag suggestions
- keeping docs and decisions aligned

### Codex

Primary implementation agent.

Responsible for:

- backend work by default
- contract-shaping implementation
- scoped code changes
- final documentation sync after accepted work

### Gemini

Frontend and integration-focused agent.

Responsible for:

- frontend discovery when design is undefined
- frontend implementation
- reviewing the latest backend contract from a frontend perspective
- identifying integration mismatch before frontend coding begins

### Claude Code

Selective reviewer.

Responsible for:

- milestone-level review
- risky or cross-cutting review
- scope validation
- contract drift detection
- regression and complexity detection

---

## Standard Delivery Loop

### Step 1 — Define the target

Before implementation, identify:

- current phase
- current sprint
- current step
- expected outcome
- scope boundaries

### Step 2 — Decide whether frontend discovery is needed

If the task involves frontend work and the UI direction is still unclear:

- ChatGPT prepares a Gemini discovery prompt
- Gemini explores the scope and proposes realistic directions
- no file edits happen yet unless explicitly requested

If UI direction is already clear, skip to the next step.

### Step 3 — Generate Codex prompt

ChatGPT prepares a focused Codex implementation prompt containing:

- files to read first
- exact goal
- constraints
- out-of-scope items
- expected deliverable

### Step 4 — Implement backend or contract-changing work with Codex

Codex performs the implementation needed for the current step.

This is the default first implementation pass when the backend/frontend boundary may be affected.

### Step 5 — Gemini integration review before frontend coding

If frontend work depends on the new backend behavior:

- Gemini reads the latest backend diff and relevant docs
- Gemini identifies contract friction, missing fields, naming issues, or frontend risks
- Gemini reports the issues before frontend implementation starts

This is an integration check, not a full final review.

### Step 6 — Patch backend if needed

If Gemini finds a legitimate integration issue:

- ChatGPT interprets it
- Codex applies the smallest necessary backend patch
- backend behavior is stabilized before frontend coding

### Step 7 — Implement frontend with Gemini

Once the backend contract is stable:

- Gemini implements the frontend slice
- frontend should follow the documented API shape
- loading, empty, and error states should be handled explicitly

### Step 8 — Selective Claude review

Claude Code reviews the integrated result when the change is:

- risky
- cross-cutting
- milestone-level
- close to a tag-worthy checkpoint

Claude is not required for every trivial change, but should be used when the risk justifies it.

### Step 9 — Final documentation sync with Codex

After the accepted implementation and review flow:

- Codex updates the final docs
- docs should reflect backend behavior, Docker flow, sprint status, and milestone state
- documentation sync happens last to reduce drift

### Step 10 — Commit

After approval, commit the finished unit.

### Step 11 — Tag

Create a version tag only when a meaningful milestone or completed phase justifies it.

---

## Scope Rules

### Allowed

- work only on the current step
- make small supporting changes required by the step
- update docs when accepted work changes project understanding or project state

### Not Allowed

- unrelated refactors
- speculative optimization
- undocumented backend contract invention
- hidden feature expansion
- large architecture changes without planning

---

## Integration Rule

If the frontend/backend boundary changed, the work is not considered complete until:

- the backend contract is stable
- the frontend follows the latest accepted contract
- the local flow still works
- the related docs are synced

This rule prevents silent integration drift.

---

## Review Philosophy

Claude Code is a reviewer, not a second primary implementer.

Gemini is not the final backend authority.

Codex is not allowed to silently expand scope.

ChatGPT coordinates the flow and keeps the work aligned with the documented plan.

---

## Commit Discipline

A commit should usually represent one of these:

- one completed step
- one approved patch after review
- one bounded docs sync set
- one integrated slice that still remains clearly reviewable

Avoid mixing unrelated concerns unless it is a deliberate bootstrap change.

---

## Tag Discipline

Tags are for milestones, not routine progress.

Strong candidates for tags include:

- completed foundation milestone
- completed backend-foundation milestone
- first ingestion milestone
- first useful dashboard milestone
- completed phase

---

## Source of Truth Order

If documents conflict, follow this priority:

1. `AGENTS.md`
2. `WORKFLOW.md`
3. `VERSIONING.md`
4. current phase document
5. current sprint document
6. implementation details in code

Implementation must follow the documented plan, not invent a new one.
