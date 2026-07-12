---
name: spug-sms
description: Guide and implement Spug Push SMS integrations in application code. Use when Codex needs to add template SMS sending, dynamic template variables, delivery-status queries, recent send-history queries, configuration, error handling, or tests for push.spug.cc in an existing project.
---

# Spug 短信接入

帮助 Codex 在现有项目中编写 Spug 短信接入代码。先检查项目使用的语言、HTTP 客户端、配置方式和测试习惯，再读取 `references/api.md`；优先复用现有依赖和代码模式。

## 接入流程

1. 确认用户要接入发送、状态查询还是历史查询；仅实现实际需要的能力。
2. 找到项目现有的 HTTP 客户端、配置读取、错误处理和测试方式，不另建无必要的抽象或依赖。
3. 将模板编码和查询 `token` 放入环境变量或项目现有密钥系统，不写入源码、示例、日志或提交内容。
4. 将发送函数设计为接收手机号和动态变量集合。变量名由模板决定，不要固定为 `name`、`code`、`number` 等字段。
5. 正确区分“接口受理成功”和“短信已送达”：发送接口返回 `code: 200` 后保存 `request_id`；需要送达结果时再调用状态查询接口。
6. 使用项目现有测试工具模拟 HTTP 响应，至少覆盖成功响应和接口错误；测试不得真实发送短信。

## 实现约束

- 发送地址为 `POST https://push.spug.cc/sms/{template_code}`，请求体包含 `to` 和该模板要求的动态变量。
- 多个手机号使用英文逗号分隔。除非项目已有批处理约定，否则不要额外实现队列、重试器或 SDK 封装层。
- 模板编码无需另行登录即可代表对应用户，必须按凭据保护。日志中不得输出完整请求 URL，因为 URL 包含模板编码。
- 查询状态与历史使用独立的 `token`，不要误用模板编码。
- 状态可能延迟 1–3 分钟；如需轮询，间隔不得短于 1 分钟，并沿用项目已有的任务机制。
- 历史接口只支持最近 30 天的 `sms` 和 `mail`；不要为它声称支持语音等类型。

## 输出代码

直接修改用户项目或给出符合其语言和框架习惯的最小代码。说明需要配置的环境变量、函数入参、返回值与错误行为。除非用户明确要求，不要执行真实发送请求。
