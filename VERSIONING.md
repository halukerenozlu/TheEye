# VERSIONING.md

## Purpose

This document defines how versioning, tags, and milestone releases are handled in TheEye.

The project uses a lightweight form of **SemVer (Semantic Versioning / Anlamsal Sürümleme)**:

`vMAJOR.MINOR.PATCH`

Examples:

- `v0.1.0`
- `v0.2.0`
- `v0.2.1`

---

## Meaning of Each Part

### MAJOR

Increase **MAJOR** when there is a major product direction change, a breaking architectural shift, or a public release that significantly changes the system.

Examples:

- complete rewrite of service structure
- major contract break between versions
- transition from MVP architecture to a deliberately different production architecture

### MINOR

Increase **MINOR** when a meaningful milestone or phase is completed.

Examples:

- infra and local workflow baseline completed
- product definition and domain model completed
- backend foundation completed
- first ingestion pipeline completed
- first dashboard milestone completed

### PATCH

Increase **PATCH** for small fixes and post-review improvements that do not change milestone scope.

Examples:

- bug fix after review
- doc sync after accepted implementation
- small cleanup without structural change
- narrow regression fix

---

## Tagging Philosophy

Tags are milestone checkpoints, not routine progress markers.

Preferred tagging unit:

- **completed phase**
- **approved milestone**
- **clean checkpoint worth revisiting later**

A sprint can finish without requiring a tag.
A normal commit should usually not create a tag by itself.

---

## Initial Release Roadmap

Planned early versions:

- `v0.1.0` → repository structure, infra baseline, local development workflow
- `v0.2.0` → product scope, domain model, documentation baseline
- `v0.3.0` → backend foundation and service skeleton
- `v0.4.0` → first ingestion pipeline
- `v0.5.0` → API list/detail milestone
- `v0.6.0` → first dashboard milestone
- `v0.7.0` → stronger UX, filters, timeline, better interaction
- `v0.8.0` → alerts, saved views, admin basics
- `v0.9.0` → stabilization and release preparation
- `v1.0.0` → first stable showcase release

---

## Commit vs Tag Rules

### Commits

Create commits for:

- one completed step
- one accepted patch
- one bounded docs update set
- one clear integrated slice

### Tags

Create tags only when a milestone is complete and intentionally checkpointed.

Tags should not be used for:

- every commit
- every small fix
- every sprint by default

---

## Release Discipline

Before creating a tag:

- the relevant phase or milestone should be complete
- the implementation should be reviewed at the right risk level
- frontend/backend integration should be stable if the boundary changed
- local Docker-based development flow should still work
- the working tree should be clean
- key docs should be synced to the accepted state

Preferred final docs sync should happen before tagging.

---

## Tagging Convention

Annotated tags are preferred.

Example:

```bash
git tag -a v0.3.0 -m "v0.3.0: complete backend foundation milestone"
git push origin v0.3.0
```

---

## Notes

During early development, versions should remain under `0.x.y`.

This signals that the project is still evolving and that some internal details may change rapidly, but version tags should still represent meaningful, deliberate checkpoints.
