# CLAUDE.md — Working Agreement for Claude Code

Read **AGENTS.md** first. If there is any conflict, **AGENTS.md wins**.

## Quick Commands (local)

- Install JS deps (workspace): `pnpm -w install`
- Dashboard dev: `pnpm --filter dashboard dev`
- API dev (later): `make dev` or `go run ./cmd/api` (to be added)
- Full stack: `docker compose up` (from `infra/` or repo root, depending on setup)

## Do / Don't

- Do NOT break `docker compose up` local flow.
- Keep changes small and reviewable.
- Avoid introducing new frameworks.
- Always add timeouts + retries (bounded) for outbound HTTP calls in collectors.
- Prefer SSE (Server-Sent Events (Sunucu Taraflı Olay Akışı)) for MVP realtime.

## Output expectations

When you propose changes, include:

1. Assumptions (3–7 bullets)
2. Exact files changed
3. Exact commands to verify
