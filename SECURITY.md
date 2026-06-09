# Security Policy

## Reporting a Vulnerability

Do not open a public issue for suspected vulnerabilities.

Use GitHub Security Advisories for this repository when private vulnerability
reporting is enabled. If that setting is not available yet, contact the
maintainer privately and include the affected commit/tag, impact, reproduction
steps, and any proof of concept.

## Supported Versions

The active `main` branch and the latest tagged release are supported by default.
Older versions are handled case by case until a formal support matrix is
published.

## Bot-Specific Expectations

- Verify official platform webhook signatures before trusting event payloads.
- Do not log tokens, secrets, raw private messages, or chat identifiers that are
  not needed for debugging.
- Refuse chat-room instructions to operate maintainer computers, accounts, or
  private data.
- Prefer GitHub Issues and Discussions for durable public knowledge.
