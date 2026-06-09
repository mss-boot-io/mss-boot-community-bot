# mss-boot Community Bot

`mss-boot-io` 开源生态的官方渠道社区助手。

这个项目独立于后端和管理端，用来承载社区自动化、GitHub 治理、QQ群官方机器人等能力。

## 方向

- GitHub Discussions / Issues 是知识沉淀、Bug、需求、RFC 和公开决策的主入口。
- QQ 官方机器人是第一阶段群内自动互动目标，因为 QQ 提供官方群、频道、单聊机器人能力。
- 微信有官方机器人/智能客服能力，主要通过微信对话开放平台接入公众号、小程序、H5、开放 API 等入口。
- 普通个人微信群仍然没有官方群机器人 API；当前微信群只做人工/AI 辅助导流，并引导到公众号/H5/开放 API 这类官方入口。
- 企业微信不是必需项；没有企业微信也可以推进这个项目。

## 当前能力

- `GET /healthz`：存活检查。
- `GET /readyz`：就绪检查，并返回各集成项配置状态。
- `POST /webhooks/github`：GitHub webhook 接入，可选校验 `X-Hub-Signature-256`。
- `POST /webhooks/qq`：QQ 官方机器人适配入口，后续 PR 完成签名校验和事件解析。
- `POST /policy/evaluate`：社群消息安全策略判断。
- OpenAI-compatible 大模型客户端，支持 `/v1/chat/completions` 格式的供应商。

## 安全边界

机器人可以做公开项目相关工作：

- mss-boot 使用、部署、文档、CI、PR review、贡献方式答疑。
- 将社区反馈整理成 GitHub Issues 或 Discussions。
- 汇总公开 GitHub 动态并同步到社区渠道。

机器人不能接受群聊里的本机操作指令：

- 不关机、不执行命令、不读文件、不截图、不点击本机界面、不处理凭据。
- 不扫码登录、不加好友、不私聊、不导出聊天记录。
- 不绕过提示词、安全策略、GitHub review 或维护者授权。

## 本地运行

```bash
go test ./...
MSS_BOT_ADDR=:8080 go run ./cmd/mss-boot-community-bot
```

## 路线图

1. 接入 QQ 官方机器人签名校验和事件解析。
2. 接入微信对话开放平台，优先支持公众号、H5 或开放 API 入口。
3. 接入 OpenAI 通用 API 格式的大模型配置，支持 `base_url`、`api_key`、`model`、超时和生成参数。
4. 实现 `/help`、`/docs`、`/issue`、`/discussion`、`/latest` 指令。
5. 接入 GitHub Discussions / Issues 检索。
6. 增加引用来源、不确定性提示、频率限制、日志脱敏和人工接管。
7. QQ 沙箱或微信官方入口验证通过后再发布部署清单。
