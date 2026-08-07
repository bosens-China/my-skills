---
name: spug-sms
description: Guide and implement Spug Push SMS integrations in application code. Use when Codex needs to add template SMS sending, dynamic template variables, SMS delivery-status queries, recent SMS send-history queries, configuration, error handling, or tests for push.spug.cc in an existing project.
---

# Spug 短信接入

帮助 Codex 在现有项目中编写 Spug 短信接入代码。先检查项目使用的语言、HTTP 客户端、配置方式和测试习惯，再读取 `references/api.md`；优先复用现有依赖和代码模式。

## 接入流程

1. 确认用户要接入短信发送、短信状态查询还是短信历史查询；仅实现实际需要的能力，不扩展邮件等其他通道。
2. 找到项目现有的 HTTP 客户端、配置读取、错误处理和测试方式，不另建无必要的抽象或依赖。
3. 将模板编码或受信任的模板 URL 与查询 `token` 放入环境变量或项目现有密钥系统，不写入源码、示例、日志或提交内容；不要接收业务请求传入的任意 URL。
4. 默认将发送函数设计为接收并始终发送 `to`、`name`、`code`。只有项目维护了“模板 → 所需变量”配置时，才根据模板元数据选择字段。
5. 正确区分“接口受理成功”和“通道处理成功”：发送接口返回 `code: 200` 后保存 `request_id`；需要最终结果时再调用状态查询接口。
6. 使用项目现有测试工具模拟 HTTP 响应，至少覆盖成功响应和接口错误；测试不得真实发送短信。

## 实现约束

- 发送接口同时支持 `POST` 和 `GET`。默认使用 `POST https://push.spug.cc/sms/{template_code}` 并发送 JSON；仅在项目现有约定或调用环境明确需要时使用 GET 查询参数形式。
- GET 参数必须交给 HTTP 客户端编码，不要手工拼接；手机号、验证码等敏感动态值会进入 URL，因此不得记录完整 URL。
- 在发送前校验：`to` 中每个手机号均为 11 位数字，`name` 为 1–10 位中文、英文字母或数字，`code` 为 4–6 位英文字母或数字。
- 多个手机号使用英文逗号分隔，单次最多 5 个。除非项目已有批处理约定，否则不要额外实现队列、重试器或 SDK 封装层。
- 模板编码无需另行登录即可代表对应用户，必须按凭据保护。无论使用 GET 还是 POST，日志中都不得输出含模板编码的完整请求 URL。
- 查询状态与历史使用独立的 `token`，不要误用模板编码。
- 对待发送、发送中或审核中状态使用间隔查询，避免紧密轮询，并沿用项目已有的任务机制。
- 查询历史时固定传入 `type: "sms"`，只处理最近 30 个自然日内的短信记录。

## 输出代码

直接修改用户项目或给出符合其语言和框架习惯的最小代码。说明需要配置的环境变量、函数入参、返回值与错误行为。除非用户明确要求，不要执行真实发送请求。
