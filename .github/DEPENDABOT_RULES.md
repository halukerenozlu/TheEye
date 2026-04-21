# DEPENDABOT_RULES.md

## Purpose

This file defines a simple decision policy for reviewing and merging Dependabot pull requests in TheEye.

The goal is to reduce noise, avoid unnecessary hesitation, and keep dependency updates controlled.

---

## Core Principle

Not every Dependabot PR should be treated the same.

Use this classification:

- GitHub Actions updates
- minor / patch dependency updates
- major dependency updates

Each category has a different merge policy.

---

## Category 1 — GitHub Actions PRs

Examples:

- `actions/checkout`
- `actions/setup-go`
- `actions/setup-node`

### Default rule

These can usually be merged directly if:

- the PR only changes workflow files under `.github/workflows/`
- the change is small and isolated
- the PR is marked ready to merge
- there are no merge conflicts
- no unrelated files were changed

### Do not merge immediately if

- the PR changes multiple unrelated workflow behaviors
- required checks are failing
- the workflow change affects deployment or production secrets handling
- the change is not limited to a straightforward version bump

### Practical policy

If it is a simple version bump in CI workflow files only, merge it.

---

## Category 2 — Minor / Patch Dependency PRs

Examples:

- grouped dashboard dependency updates
- grouped Go module updates
- small non-breaking package bumps

### Default rule

These should be merged only after local verification.

### Dashboard verification

Run from `apps/dashboard`:

```bash
pnpm install
pnpm build
```

If available, also run:

```bash
pnpm lint
pnpm test
```

### API verification

Run from `services/api`:

```bash
go test ./...
go build ./...
```

### Merge rule

If the update is minor or patch only, and local verification passes, merge it.

### Do not merge if

- build fails
- tests fail
- lint fails in a meaningful way
- the PR unexpectedly changes source code or config outside dependency files
- there are compatibility concerns that require manual code changes

---

## Category 3 — Major Dependency PRs

Examples:

- `typescript 5 -> 6`
- `react 18 -> 19`
- `next 15 -> 16`
- `eslint 9 -> 10`

### Default rule

Do not merge major updates automatically.

### Review requirements

Before merging a major dependency PR:

- inspect the release impact
- verify whether config changes are needed
- run local build and tests
- check whether code changes are required
- merge only after deliberate review

### Practical policy

Major updates should be reviewed manually and treated as planned maintenance, not routine dependency cleanup.

---

## Merge Decision Table

### Merge directly

Use for:

- simple GitHub Actions version bumps
- isolated workflow-only PRs

### Test then merge

Use for:

- grouped minor / patch npm updates
- grouped minor / patch Go module updates

### Hold for manual review

Use for:

- major dependency updates
- framework version jumps
- toolchain jumps
- suspicious or unexpectedly broad PRs

---

## Current Project-Specific Guidance

### Safe to merge quickly

- `actions/checkout`
- `actions/setup-go`
- `actions/setup-node`

If they are isolated workflow-only PRs.

### Merge after local verification

- grouped dashboard `minor/patch` PRs
- grouped API `minor/patch` PRs

### Do not auto-merge

- major TypeScript updates
- major React / Next.js updates
- major ESLint updates
- anything that may require code or config adaptation

---

## Grouped Dependabot Policy

The repository may use grouped Dependabot updates.

### Expected behavior

- minor / patch dashboard updates may arrive as one grouped PR
- major updates may remain separate
- GitHub Actions updates may be grouped or remain isolated depending on configuration

### Interpretation rule

A grouped PR should be judged by the highest-risk dependency in the group.

If a grouped PR contains only minor/patch changes, use the normal local verification flow.
If it includes a major update, treat it as manual-review required.

---

## Red Flags

Stop and inspect more carefully if:

- the PR touches unexpected files
- source code changes appear in addition to dependency file changes
- the lockfile diff is unusually large
- local install/build fails
- the app no longer starts or builds
- the PR description and changed files do not match

---

## Default Maintainer Rule

Use this simple rule of thumb:

- **Actions PRs:** merge if isolated and clean
- **Minor/Patch dependency PRs:** test locally, then merge
- **Major dependency PRs:** do not auto-merge

---

## Final Note

Dependabot should reduce maintenance burden, not create blind automation.

TheEye should accept safe routine updates quickly, but should never normalize unreviewed major dependency changes.
