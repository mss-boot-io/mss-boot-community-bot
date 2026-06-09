# Contributing to mss-boot Community Bot

## Issues

Use the issue templates and include the channel, event payload shape,
reproduction steps, expected behavior, actual behavior, and logs when reporting
bugs. Do not report vulnerabilities in public issues.

## Pull Requests

- Use a Conventional Commits title, for example `feat(qq): parse group messages`.
- Keep changes scoped and include tests for behavior changes.
- Run `go test ./...` before requesting review.
- Update README, docs, `aigc/prompts/`, or the organization docs project when
  behavior, configuration, release policy, or AI workflow changes.
- For chat adapters, document platform terms, signature verification, rate
  limits, privacy impact, and rollback behavior.

AI-assisted changes are welcome when the generated output is reviewed and
verified by the contributor.
