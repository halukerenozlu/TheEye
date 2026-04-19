# SECURITY.md

## Security Policy

TheEye is an actively developing project and may change rapidly before `v1.0.0`.

Security reports are welcome and should be handled responsibly.

---

## Supported Versions

At this stage, support is best-effort and focused on the latest development state.

| Version            | Supported |
| ------------------ | --------- |
| `main`             | yes       |
| older pre-1.0 tags | no        |

---

## Reporting a Vulnerability

Please do **not** open a public GitHub issue for security-sensitive reports.

Instead, use one of these private paths when available:

- GitHub Security Advisories / private vulnerability reporting
- direct private contact with the maintainer

If neither is available yet, wait until a private reporting channel is configured rather than disclosing sensitive details publicly.

---

## What to Include

A useful report should include, when possible:

- a clear description of the issue
- affected component(s)
- steps to reproduce
- impact assessment
- proof of concept, if safe to share
- suggested mitigation, if known

---

## Response Expectations

The project will try to:

- acknowledge the report
- assess severity and impact
- validate the issue
- prepare a fix when appropriate
- disclose responsibly after a fix is available, if needed

Because the project is early-stage, response times may vary.

---

## Scope Notes

Security-sensitive areas may later include:

- API input validation
- external source ingestion
- secrets and environment configuration
- rate limiting and abuse controls
- dependency vulnerabilities

Contributors should avoid committing secrets and should follow secure local development practices.
