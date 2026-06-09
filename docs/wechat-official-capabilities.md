# WeChat Official Bot Capabilities

## Summary

WeChat does have official bot-like capabilities. The important distinction is
the entry point:

- WeChat Conversation Open Platform provides intelligent customer service and
  conversation bot capabilities.
- The official entry points are official account, mini program, H5, and Open
  API. The documentation also mentions WeChat customer service and enterprise
  WeChat channels.
- A personal WeChat group is not the same as these entry points. The current
  official documents do not provide a personal group bot API that lets a bot
  join a normal WeChat group and read/reply to group messages.

## What mss-boot Should Support

1. GitHub remains the durable source of truth for public project knowledge.
2. QQ official bot remains the first group-automation target.
3. WeChat Conversation Open Platform should be added as the official WeChat
   intake path:
   - official account for 1:1 public project FAQ;
   - H5 page or QR code for lightweight support entry;
   - Open API adapter for mss-boot docs and GitHub search;
   - mini program only if there is a concrete product reason.
4. Existing personal WeChat groups should continue as community channels with
   maintainer-authorized manual/AI-assisted replies and links to official
   entry points.

## Official References

- WeChat Conversation Open Platform introduction:
  https://developers.weixin.qq.com/doc/aispeech/platform/INTRODUCTION.html
- WeChat robot reply configuration:
  https://developers.weixin.qq.com/doc/aispeech/platform/dialog/skill-reply.html
- WeChat official account application binding:
  https://developers.weixin.qq.com/doc/aispeech/platform/application/official_account.html
- WeChat customer service application binding:
  https://developers.weixin.qq.com/doc/aispeech/platform/application/wxkefu.html
- Enterprise WeChat message push, formerly group robot:
  https://developer.work.weixin.qq.com/document/path/91770
