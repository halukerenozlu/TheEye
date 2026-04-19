# WORKFLOW.md

## Purpose

This document defines the standard development workflow for TheEye.

TheEye is developed through a role-separated multi-agent process:

- Human vision and product direction
- ChatGPT for planning, architecture guidance, prompt generation, and review interpretation
- Codex for scoped implementation
- Claude Code for review

The goal is controlled, reviewable progress instead of freeform code generation.

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

- product vision
- priorities
- tradeoff decisions
- approval or rejection of work
- deciding what gets committed and tagged

### ChatGPT

Responsible for:

- roadmap planning
- phase / sprint / step definition
- Codex implementation prompt generation
- Claude Code review prompt generation
- interpreting review results
- commit and tag suggestions
- keeping the work aligned with project direction

### Codex

Responsible for:

- implementing the requested step
- making minimal scoped code changes
- preserving project structure and conventions
- avoiding scope creep

### Claude Code

Responsible for:

- reviewing Codex changes
- checking scope correctness
- checking doc alignment
- identifying risks, regressions, and unnecessary complexity
- suggesting only minimal fixes when needed

---

## Standard Delivery Loop

### Step 1 — Define the target

Before implementation, identify:

- current phase
- current sprint
- current step
- expected outcome
- scope boundaries

### Step 2 — Generate Codex prompt

ChatGPT prepares a focused implementation prompt containing:

- files to read first
- exact goal
- constraints
- out-of-scope items
- expected deliverable

### Step 3 — Implement with Codex

Codex performs only the requested work.

### Step 4 — Generate Claude review prompt

ChatGPT prepares a review prompt focused on:

- scope compliance
- architectural alignment
- correctness
- minimalism
- risk detection

### Step 5 — Review with Claude Code

Claude reviews the uncommitted changes and returns one of:

- Accept
- Accept with minimal patch
- Rework needed
- Reject

### Step 6 — Interpret the review

Claude output is evaluated and one of the following happens:

- accept as is
- apply a minimal patch
- request a corrected implementation
- reject and redo the step

### Step 7 — Commit

After approval, commit the finished unit.

### Step 8 — Tag

Create a version tag only when a meaningful milestone is complete.

---

## Scope Rules

### Allowed

- work only on the current step
- make small supporting changes directly required by the step
- update docs if the step clearly changes project understanding

### Not Allowed

- unrelated refactors
- speculative optimization
- broad architectural changes without planning
- hidden feature expansion
- unnecessary dependencies

---

## Review Philosophy

Claude Code is a reviewer, not a second primary implementer.

Priority order:

1. Is the implementation within scope?
2. Does it align with AGENTS.md and project docs?
3. Is it correct and safe?
4. Is it minimal?
5. Are fixes actually required?

---

## Commit Discipline

A commit should usually represent one of these:

- one completed step
- one approved patch after review
- one bounded docs update set

Avoid mixing unrelated concerns unless it is a deliberate bootstrap commit.

---

## Tag Discipline

Tags are for milestones, not routine progress.

Examples:

- infra and repo governance baseline complete
- backend foundation milestone complete
- first ingestion pipeline complete
- first user-facing dashboard milestone complete

---

## Current Active Pattern

The current practical workflow for TheEye is:

1. Human defines direction
2. ChatGPT defines the next step
3. Codex implements
4. Claude Code reviews
5. ChatGPT interprets the review
6. Human approves and commits

---

## Source of Truth Order

If documents conflict, follow this priority:

1. AGENTS.md
2. WORKFLOW.md
3. VERSIONING.md
4. current phase document under `docs/phases/`
5. current sprint document under `docs/sprints/`
6. implementation details in code

Implementation must follow the documented plan, not invent a new one.
