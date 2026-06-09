# QQ Official Bot Setup Notes

## Current Configuration

The QQ bot has been created in the official QQ bot console.

- Bot display name: `mssboot助手`
- Introduction: `mss-boot 开源项目社区问答助手，提供文档检索、Issue/Discussion 引导与贡献协作支持。`
- Message-list command configured: `/help`
- Command description: `查看帮助`

## Important Platform Limitation

The current QQ bot console shows that QQ group configuration is not supported
for this bot:

```text
暂不支持群相关配置，敬请期待
```

The sandbox page also warns that AIGC bots are not currently supported in social
group scenarios or full public release to all users. Treat this bot as a
message-list / one-to-one bot until the platform supports the required group
scenario or a separate non-AIGC bot type is created.

## Remaining Setup

- Configure an HTTPS event callback URL after this service is deployed to a
  public HTTPS domain.
- Alternatively implement the official WebSocket Gateway adapter for development
  and private testing.
- Do not store AppSecret or token values in the repository or AI memory.
- Add message URL allowlist entries before sending replies that contain links.
- Upload a privacy policy before public release.
- Complete the official self-test report before release review.

## Recommended Next Technical Step

Implement the QQ adapter in this order:

1. Access token retrieval.
2. WebSocket Gateway connection for development.
3. Event parsing for message-list command events.
4. `/help` response with no external links first.
5. OpenAI-compatible LLM response path only after policy filtering and GitHub
   context retrieval are in place.
