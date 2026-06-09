# mss-boot Community Bot

Official-channel community assistant for the `mss-boot-io` open-source
ecosystem.

The project is intentionally separate from the backend and admin console. It
keeps community automation, GitHub triage, and chat-platform integration in one
small service with explicit safety boundaries.

## Direction

- GitHub Discussions and Issues are the source of truth for reusable knowledge,
  bug reports, feature requests, RFCs, and public decisions.
- QQ official bot is the first chat automation target because QQ provides an
  official bot platform for group, channel, and direct-message scenarios.
- WeChat has official bot and intelligent-customer-service capabilities through
  the WeChat Conversation Open Platform. For this project, the practical WeChat
  targets are official account, mini program, H5, and Open API entry points.
- Personal WeChat groups still do not have an official group bot API. The
  current WeChat group path remains manual AI-assisted guidance plus links to
  official WeChat bot entry points.
- Enterprise WeChat is optional. This project must remain useful without it.

## Current Capabilities

- `GET /healthz`: liveness probe.
- `GET /readyz`: readiness signal. It reports `degraded` until QQ credentials
  are configured.
- `POST /webhooks/github`: GitHub webhook intake with optional
  `X-Hub-Signature-256` verification.
- `POST /webhooks/qq`: placeholder endpoint for the official QQ bot adapter.
- `POST /policy/evaluate`: local safety policy evaluation for community
  messages.

## Safety Boundaries

The bot can help with public project work:

- mss-boot usage, deployment, documentation, CI, PR review, and contribution
  guidance.
- Turning actionable community feedback into GitHub Issues or Discussions.
- Summarizing public GitHub activity for community channels.

The bot must not accept chat-room instructions to operate the maintainer's
computer:

- no shutdown, command execution, file reading, screenshots, local clicks, or
  credential handling;
- no QR-code login, friend requests, private messages, or chat export;
- no bypassing prompts, policies, GitHub review, or maintainer authorization.

## Local Run

```bash
go test ./...
MSS_BOT_ADDR=:8080 go run ./cmd/mss-boot-community-bot
```

```bash
curl http://localhost:8080/healthz
curl -X POST http://localhost:8080/policy/evaluate \
  -H 'Content-Type: application/json' \
  -d '{"text":"mss-boot 怎么配置 RBAC 权限?"}'
```

## Configuration

| Environment variable | Required | Description |
| --- | --- | --- |
| `MSS_BOT_ADDR` | No | HTTP listen address. Defaults to `:8080`. |
| `MSS_BOT_GITHUB_WEBHOOK_SECRET` | No | Enables GitHub webhook signature verification. |
| `MSS_BOT_QQ_APP_ID` | For QQ | Official QQ bot app ID. |
| `MSS_BOT_QQ_APP_SECRET` | For QQ | Official QQ bot app secret. |

## Roadmap

1. Add QQ official bot signature verification and event parsing.
2. Add a WeChat Conversation Open Platform adapter for official account, H5, or
   Open API entry points.
3. Implement `/help`, `/docs`, `/issue`, `/discussion`, and `/latest` commands.
4. Add GitHub Discussions and Issues search backed by public repository data.
5. Add response citations and uncertainty handling.
6. Add rate limits, audit logs with redaction, and maintainer handoff.
7. Publish deployment manifests after the first real QQ sandbox or WeChat entry
   verification.

## License

MIT
