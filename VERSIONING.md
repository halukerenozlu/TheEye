# VERSIONING.md

## Purpose
This document defines how versioning, tags, and milestone releases are handled in TheEye project.

## Versioning Model
The project uses a lightweight form of SemVer (Semantic Versioning / Anlamsal Sürümleme):

`vMAJOR.MINOR.PATCH`

Example:
- `v0.1.0`
- `v0.2.0`
- `v0.2.1`

## Meaning of Each Part

### MAJOR
Increase MAJOR when there is a major product direction change, breaking architecture shift, or a public release that significantly changes the system.

Examples:
- complete rewrite of backend structure
- moving from MVP architecture to production-grade modular architecture
- major API breaking changes

### MINOR
Increase MINOR when a meaningful new milestone is completed.

Examples:
- infrastructure setup completed
- domain model completed
- first ingestion pipeline completed
- first dashboard completed

### PATCH
Increase PATCH for small fixes, review patches, cleanup, and non-structural improvements.

Examples:
- bug fix after review
- typo and doc fixes
- minor refactor without changing scope
- small test fixes

## Initial Release Roadmap
Planned early versions:

- `v0.1.0` → repository structure, infra baseline, local development workflow
- `v0.2.0` → product scope, domain model, database draft
- `v0.3.0` → backend foundation and service skeleton
- `v0.4.0` → first ingestion pipeline
- `v0.5.0` → API list/detail endpoints
- `v0.6.0` → first dashboard UI
- `v0.7.0` → filters, sorting, timeline, better UX
- `v0.8.0` → alert logic, saved views, admin basics
- `v0.9.0` → stabilization and release preparation
- `v1.0.0` → first stable showcase release

## Commit vs Tag Rules

### Commits
Create commits for meaningful completed work units.

Examples:
- one completed step
- one scoped review fix
- one finished doc update set

### Tags
Create tags only when a milestone is complete.

Tags should not be used for every commit.
Tags should represent a clear checkpoint in product progress.

## Tagging Convention
Annotated tags are preferred.

Example:
```bash
git tag -a v0.1.0 -m "v0.1.0: establish infra and local workflow baseline"
git push origin v0.1.0

Release Discipline

Before creating a tag:

related sprint or milestone must be completed
Codex implementation must be reviewed by Claude Code
final decision must be approved manually
working tree should be clean
key docs should be updated if needed
Notes

During early development, versions remain under 0.x.y.
This means the project is still evolving and some internal changes may happen rapidly.


---

# 2) `WORKFLOW.md`

```md
# WORKFLOW.md

## Purpose
This document defines the development workflow for TheEye.

The workflow is based on a role-separated collaboration model:
- Human vision and product direction
- Codex for implementation
- Claude Code for review
- ChatGPT for planning, prompt generation, architecture guidance, and review interpretation

## Core Principle
Every change must belong to a defined structure:

- Phase
- Sprint
- Step

No implementation should start without being mapped to a step.

---

## Roles

### Human
Responsible for:
- product vision
- priorities
- approval
- final decisions
- deciding whether to accept, patch, or reject changes

### ChatGPT
Responsible for:
- roadmap planning
- phase and sprint definition
- step breakdown
- Codex implementation prompt generation
- Claude Code review prompt generation
- interpreting review results
- commit and tag suggestions

### Codex
Responsible for:
- implementing the requested step
- making scoped code changes
- editing files needed for the task
- avoiding scope creep

### Claude Code
Responsible for:
- reviewing Codex changes
- checking scope correctness
- identifying risks and inconsistencies
- suggesting minimal fixes only when necessary

---

## Standard Delivery Loop

### Step 1 — Planning
ChatGPT defines:
- current phase
- current sprint
- current step
- expected output
- scope boundaries

### Step 2 — Codex Prompt
ChatGPT creates a Codex prompt that includes:
- files to read first
- exact scope
- constraints
- what not to change
- expected deliverables

### Step 3 — Implementation
Codex implements only the requested step.

### Step 4 — Review Prompt
ChatGPT creates a Claude Code review prompt focused on:
- scope compliance
- architecture consistency
- doc alignment
- minimalism of changes
- risks and possible regressions

### Step 5 — Review Output
Claude Code review output is brought back for interpretation.

### Step 6 — Decision
One of these decisions is made:
- accept as is
- apply a minimal patch
- request a corrected implementation
- reject and redo the step

### Step 7 — Commit
After approval, commit the finished unit.

### Step 8 — Tag
After a meaningful milestone, create a version tag.

---

## Scope Rules

### Allowed
- work only on the current step
- minimal fixes directly required by the step
- small supporting updates in docs if clearly necessary

### Not Allowed
- unrelated refactors
- speculative optimization
- architecture rewrites without planning
- changing naming conventions without reason
- hidden scope expansion

---

## Review Philosophy
Claude Code should behave like a reviewer, not a second implementer.

Priority order:
1. Is the work within the requested step?
2. Does it align with docs and architecture?
3. Are there correctness or safety issues?
4. Are minimal fixes needed?

---

## Commit Discipline
A commit should usually represent one of the following:
- one completed step
- one approved patch after review
- one well-bounded documentation update set

Avoid mixing multiple unrelated concerns in one commit.


---

# 2) `WORKFLOW.md`

```md
# WORKFLOW.md

## Purpose
This document defines the development workflow for TheEye.

The workflow is based on a role-separated collaboration model:
- Human vision and product direction
- Codex for implementation
- Claude Code for review
- ChatGPT for planning, prompt generation, architecture guidance, and review interpretation

## Core Principle
Every change must belong to a defined structure:

- Phase
- Sprint
- Step

No implementation should start without being mapped to a step.

---

## Roles

### Human
Responsible for:
- product vision
- priorities
- approval
- final decisions
- deciding whether to accept, patch, or reject changes

### ChatGPT
Responsible for:
- roadmap planning
- phase and sprint definition
- step breakdown
- Codex implementation prompt generation
- Claude Code review prompt generation
- interpreting review results
- commit and tag suggestions

### Codex
Responsible for:
- implementing the requested step
- making scoped code changes
- editing files needed for the task
- avoiding scope creep

### Claude Code
Responsible for:
- reviewing Codex changes
- checking scope correctness
- identifying risks and inconsistencies
- suggesting minimal fixes only when necessary

---

## Standard Delivery Loop

### Step 1 — Planning
ChatGPT defines:
- current phase
- current sprint
- current step
- expected output
- scope boundaries

### Step 2 — Codex Prompt
ChatGPT creates a Codex prompt that includes:
- files to read first
- exact scope
- constraints
- what not to change
- expected deliverables

### Step 3 — Implementation
Codex implements only the requested step.

### Step 4 — Review Prompt
ChatGPT creates a Claude Code review prompt focused on:
- scope compliance
- architecture consistency
- doc alignment
- minimalism of changes
- risks and possible regressions

### Step 5 — Review Output
Claude Code review output is brought back for interpretation.

### Step 6 — Decision
One of these decisions is made:
- accept as is
- apply a minimal patch
- request a corrected implementation
- reject and redo the step

### Step 7 — Commit
After approval, commit the finished unit.

### Step 8 — Tag
After a meaningful milestone, create a version tag.

---

## Scope Rules

### Allowed
- work only on the current step
- minimal fixes directly required by the step
- small supporting updates in docs if clearly necessary

### Not Allowed
- unrelated refactors
- speculative optimization
- architecture rewrites without planning
- changing naming conventions without reason
- hidden scope expansion

---

## Review Philosophy
Claude Code should behave like a reviewer, not a second implementer.

Priority order:
1. Is the work within the requested step?
2. Does it align with docs and architecture?
3. Are there correctness or safety issues?
4. Are minimal fixes needed?

---

## Commit Discipline
A commit should usually represent one of the following:
- one completed step
- one approved patch after review
- one well-bounded documentation update set

Avoid mixing multiple unrelated concerns in one commit.

---

## Tag Discipline
Create tags for milestone completion, not for routine progress.

Examples:
- after infra baseline is accepted
- after domain model and database design are accepted
- after first end-to-end slice works

---

## Source of Truth
When there is any conflict:
1. `AGENTS.md`
2. current phase/sprint document
3. roadmap documents
4. implementation details

Implementation must follow the documented plan, not invent a new one.