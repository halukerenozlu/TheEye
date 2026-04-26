# CONTRIBUTING.md

## Contributing to TheEye

Thanks for your interest in contributing to TheEye.

TheEye is built with a controlled, review-oriented workflow. Contributions should prioritize clarity, small scope, and alignment with the documented project direction.

---

## Before You Start

Please read these files first:

1. `AGENTS.md`
2. `docs/VISION.md`
3. `docs/VERSION_PLAN.md`
4. `docs/ROADMAP.md`
5. relevant contract docs under `docs/`

If documents conflict, follow the source-of-truth order defined in `AGENTS.md`.

---

## Development Principles

Contributions should be:

- small and reviewable
- aligned with the active version milestone, work item, and implementation slice
- consistent with the repository structure
- minimal in scope
- documented when behavior changes

Avoid:

- unrelated refactors
- speculative architecture
- hidden feature expansion
- unnecessary dependency additions

---

## Workflow

Typical contribution flow:

1. identify the current version milestone, work item, and implementation slice
2. make the smallest correct change
3. verify locally
4. keep changes reviewable
5. document behavior changes if needed
6. open a pull request with clear testing notes

---

## Branching

Use feature branches for work units.

Examples:

- `feat/api-events-placeholder`
- `fix/meta-endpoint-shape`
- `docs/version-plan-update`

Do not work directly on `main` unless repository policy explicitly allows it.

---

## Commit Style

Conventional Commits are preferred.

Examples:

- `feat(api): add placeholder event detail endpoint`
- `fix(api): return structured not found response`
- `docs: update version plan progress`

Recommended prefixes:

- `feat:`
- `fix:`
- `docs:`
- `chore:`
- `refactor:`
- `test:`

---

## Local Development

Current backend service:

```bash
cd services/api
go run ./cmd/api
```

Typical verification examples:

```bash
curl http://localhost:8080/v1/healthz
curl http://localhost:8080/v1/readyz
curl http://localhost:8080/v1/meta
curl http://localhost:8080/v1/events
curl http://localhost:8080/v1/events/some-id
```

If using Docker, ensure the local dev flow remains stable and does not regress.

---

## Pull Requests

A pull request should include:

- what changed
- why it changed
- how to test it locally
- screenshots or sample responses if relevant
- any known limitations

Keep PRs focused. Smaller PRs are strongly preferred over large mixed changes.

---

## Review Expectations

Review should prioritize:

1. scope correctness
2. alignment with docs
3. correctness and safety
4. minimalism
5. maintainability

Optional improvements should be clearly separated from required fixes.

---

## Documentation Updates

Update docs when the change affects:

- workflow
- architecture direction
- API behavior
- repository structure
- setup or verification steps

Do not create unnecessary documentation churn for tiny internal edits.

---

## Questions

If direction is unclear, resolve product and architecture questions before expanding implementation scope.
