# ADR 0001: Community Bot Boundaries

## Status

Accepted

## Context

The mss-boot community currently uses GitHub, QQ groups, and WeChat groups.
Desktop automation is useful for short maintainer-authorized interactions, but
it is not a safe or durable bot foundation for an open-source project.

The maintainer does not currently have enterprise WeChat. The first official
WeChat path therefore needs to work without enterprise WeChat and should use
WeChat Conversation Open Platform entry points such as official account, H5, or
Open API when possible.

## Decision

- Build the community bot as a separate GitHub repository.
- Treat GitHub Discussions and Issues as the durable source of truth.
- Prioritize the official QQ bot platform for group automation.
- Treat WeChat Conversation Open Platform as the official WeChat bot path for
  official account, mini program, H5, and Open API entry points.
- Keep personal WeChat groups on manual or AI-assisted guidance only because
  they are not the same thing as official bot entry points.
- Keep enterprise WeChat as an optional future adapter, not a requirement.
- Reject chat-room instructions that ask the bot to operate the maintainer's
  local computer or accounts.

## Consequences

- The first release can be small and auditable.
- We avoid reverse protocols, client hooks, and fragile screenshot automation.
- WeChat users can get support through official account, H5, Open API, GitHub
  links, and human/AI-assisted group replies, but not full automatic personal
  group bot behavior.
- QQ can become the first real-time community automation channel after sandbox
  verification and platform review.
