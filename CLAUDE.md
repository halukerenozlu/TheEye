# CLAUDE.md - Working Agreement for Claude Code

Read **AGENTS.md** first. If there is any conflict, **AGENTS.md wins**.

## Role in This Project

Claude Code is used **selectively**, not as the primary development agent.

Use Claude Code mainly for:

- risky changes
- cross-cutting changes
- milestone-level validation
- final integrated review before an important commit or tag
- finding regressions, scope drift, or contract mismatch

Claude Code should behave like a reviewer, not a second primary implementer.

---

## Review Priorities

Check these in order:

1. Is the work inside the requested version milestone, work item, and implementation slice?
2. Does it align with `AGENTS.md`, `docs/VERSION_PLAN.md`, and the relevant contract docs?
3. Does it preserve the Docker and local development flow?
4. Does it preserve the backend/frontend contract?
5. Is the change minimal and reviewable?
6. Are there required fixes, or only optional improvements?

---

## What to Watch Carefully

- silent API shape changes
- frontend assumptions that do not match backend behavior
- changes that break `docker compose -f ./infra/docker-compose.yml up --build`
- unnecessary abstraction
- hidden scope expansion
- local run commands that no longer match the docs

---

## Patch Philosophy

If a patch is needed:

- keep it minimal
- prefer correcting the specific issue
- do not redesign the implementation
- do not create a new roadmap inside the review
- separate required fixes from optional ideas

---

## Output Expectations

When you produce a review, include:

1. Assumptions
2. Review scope
3. Findings grouped as:
   - Required fixes
   - Optional suggestions
4. Exact files involved
5. Exact commands to verify
6. Final verdict:
   - Accept
   - Accept with minimal patch
   - Rework needed
   - Reject

---

## Quick Commands

### Frontend dev

```bash
pnpm --filter dashboard dev
```

### Full stack / local infra

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

### Stop full stack

```bash
docker compose -f ./infra/docker-compose.yml down
```

Use current repo-specific commands if they exist, but do not break the baseline local Docker flow.
