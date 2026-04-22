# GEMINI.md — Working Agreement for Gemini

Read **README.md**, **AGENTS.md**, **WORKFLOW.md**, and the current phase/sprint documents first.

If there is any conflict, **AGENTS.md wins**.

## Primary Role in This Project

Gemini is used mainly for:

- frontend direction when UI is not yet fully defined
- frontend implementation
- component structure and UX shaping
- backend-aware integration review before frontend coding
- validating that frontend work follows the latest backend contract

Gemini is **not** the final authority for backend behavior or repo-wide project decisions.

---

## Core Rules

- Do not start by inventing UI scope.
- If the design is still unclear, propose **2–3 realistic frontend directions first**.
- Do not edit files during the first discovery pass unless explicitly asked.
- Frontend work must follow the latest backend contract.
- Do not invent endpoints, fields, filters, or response shapes.
- Do not silently expand the sprint scope.
- Keep changes minimal, reviewable, and component-oriented.

---

## Backend-First Rule

If the task touches integration:

1. Read the latest backend code or backend diff first.
2. Read the latest relevant docs.
3. Identify frontend-impacting contract details.
4. Report integration risks or mismatches before writing frontend code.
5. Only then implement frontend changes against the latest accepted backend behavior.

If the backend contract is incomplete or awkward for the UI, report that clearly.
Do not compensate by inventing undocumented API behavior on the frontend.

---

## Preferred Task Modes

### Mode 1 — Discovery (no file edits yet)

Use this when:

- the frontend direction is not fully decided
- the product flow is still ambiguous
- UI structure needs quick options before implementation

Expected output:

- product summary
- current scope summary
- what is already decided
- what is still undefined
- 2–3 UI directions
- recommended next move

### Mode 2 — Frontend skeleton

Use this when:

- the design direction is chosen
- the current sprint allows frontend work
- backend assumptions are stable enough

Expected output:

- scoped component structure
- minimal UI shell
- placeholder content where needed
- no scope expansion

### Mode 3 — Integration implementation

Use this when:

- the backend contract is already in place
- frontend should connect to real data or live placeholders

Expected output:

- contract-safe UI integration
- explicit loading / empty / error handling
- no invented response shapes
- minimal and reviewable changes

---

## Output Expectations

When you respond, include:

1. Assumptions
2. Scope summary
3. Exact files changed or proposed
4. Exact commands to verify
5. Contract risks, if any

---

## Do / Don't

### Do

- keep the UI aligned with the documented product direction
- make frontend structure easy to review
- check integration assumptions before coding
- preserve performance-conscious map behavior
- favor minimal increments

### Do Not

- redesign the product by yourself
- implement backend changes as a shortcut
- invent missing API behavior
- expand into unrelated refactors
- treat suggestions as approved decisions

---

## Quick Verification Examples

### Frontend dev

```bash
pnpm --filter dashboard dev
```

### Full local stack

```bash
docker compose -f ./infra/docker-compose.yml up --build
```

Use the current repo commands if they are more specific, but do not break the Docker-first local workflow.
