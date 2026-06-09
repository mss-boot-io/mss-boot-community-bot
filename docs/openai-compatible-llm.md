# OpenAI-Compatible LLM Adapter

The community bot uses an OpenAI-compatible Chat Completions adapter instead of
binding to one vendor.

## Configuration

```bash
MSS_BOT_LLM_BASE_URL=https://api.openai.com/v1
MSS_BOT_LLM_API_KEY=...
MSS_BOT_LLM_MODEL=...
MSS_BOT_LLM_TIMEOUT=30s
MSS_BOT_LLM_TEMPERATURE=0.2
MSS_BOT_LLM_MAX_TOKENS=1024
```

`MSS_BOT_LLM_BASE_URL` should point at the provider's OpenAI-compatible `/v1`
base URL. The bot sends requests to:

```text
POST {MSS_BOT_LLM_BASE_URL}/chat/completions
Authorization: Bearer {MSS_BOT_LLM_API_KEY}
```

## Scope

The adapter is only a transport layer. Before a message reaches the model,
community-specific policy must still:

- reject local-computer operation instructions;
- attach mss-boot project context from GitHub/docs;
- require citations for factual answers;
- avoid secrets, private chat logs, and account operations;
- escalate uncertain project commitments to maintainers.

## Provider Notes

Any provider can be used if it implements the OpenAI-compatible Chat
Completions shape with `model`, `messages`, `temperature`, `max_tokens`, and
`choices[].message.content`.
