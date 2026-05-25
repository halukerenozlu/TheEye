# CLAUDE.md - Working Agreement for Claude Code

> Read **AGENTS.md** first. If there is any conflict, **AGENTS.md wins**.

---

## Role in This Project

Claude Code is used **selectively**, not as the primary development agent.

Use Claude Code mainly for:

- risky changes
- cross-cutting changes
- milestone-level validation
- final integrated review before an important commit or tag
- finding regressions, scope drift, or contract mismatch
  Claude Code should behave like a **reviewer**, not a second primary implementer.

---

## Behavioral Guidelines (Karpathy Principles)

These govern how Claude Code thinks and acts — regardless of task type.

> **Tradeoff:** These bias toward caution over speed. For trivial tasks, use judgment.

### 1. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

- State assumptions explicitly. If uncertain, **ask**.
- If multiple interpretations exist, present them — don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, **stop**. Name what's confusing. Ask.

### 2. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If 200 lines could be 50, rewrite it.
  > Ask: _"Would a senior engineer say this is overcomplicated?"_ If yes, simplify.

### 3. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:

- Don't "improve" adjacent code, comments, or formatting.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, **mention it — don't delete it**.
  When your changes create orphans:

- Remove imports/variables/functions that **YOUR** changes made unused.
- Don't remove pre-existing dead code unless asked.
  > **The test:** Every changed line should trace directly to the user's request.

### 4. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:

| Instead of...    | Transform to...                                       |
| ---------------- | ----------------------------------------------------- |
| "Add validation" | "Write tests for invalid inputs, then make them pass" |
| "Fix the bug"    | "Write a test that reproduces it, then make it pass"  |
| "Refactor X"     | "Ensure tests pass before and after"                  |

For multi-step tasks, state a brief plan upfront:

```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
```

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

- silent API (Application Programming Interface / Uygulama Programlama Arayüzü) shape changes
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

## Output Format

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

---

## Skill Discovery

- Read `AGENTS.md` before using project skills.
- Check project skills under `.agents/skills`.
- Treat `.agents/skills` files as usable skill instructions.
- If a capability seems missing, inspect `.agents/skills` before deciding.
