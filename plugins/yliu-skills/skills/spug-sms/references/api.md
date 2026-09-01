# Spug 短信 API 参考

编写短信接入代码时以本文件的请求结构为准，并使用项目已有的 HTTP 客户端。模板编码与查询用的账户开发者 Token 是两种不同凭据。本 skill 不维护邮件等其他通道。

## 模板与参数策略

模板编码或模板 URL 可以动态配置，但必须来自服务端配置或密钥系统。优先只配置 `{template_code}` 并由代码拼入固定的 `https://push.spug.cc/sms/` 基础地址；如果项目维护完整 URL，也要将它作为受信任配置，不要接受业务请求传入的任意 URL。

不同模板可能引用不同变量。项目没有模板元数据服务时，统一发送以下四个字段，避免切换模板后缺少参数：

| 字段 | 约束 |
| --- | --- |
| `to` | 接收手机号；每个号码均为 11 位数字 |
| `name` | 名称，1–10 位中文、英文字母或数字 |
| `code` | 验证码，4–6 位英文字母或数字 |
| `number` | 有效时长，支持数字或中文数字 |

只有项目能够维护“模板 → 所需变量”的映射时，才按模板元数据构造参数；映射至少记录模板编码和必需字段，并在发送前验证字段齐全。不要通过解析短信模板文本猜测变量。

## 发送短信

推荐方式：

```text
POST https://push.spug.cc/sms/{template_code}
Content-Type: application/json
```

默认发送包含 `to`、`name`、`code`、`number` 的 JSON 数据。若项目维护了可靠的模板变量元数据，则按元数据发送该模板需要的字段。多个手机号使用英文逗号分隔，单次最多 5 个。

接口也支持 GET：

```text
GET https://push.spug.cc/sms/{template_code}?to=...&name=...&code=...&number=...
```

GET 与 POST 均可使用，但新接入默认选择 POST。GET 只用于项目现有约定或调用环境明确需要的场景；使用 HTTP 客户端的查询参数功能进行编码，不要手工拼接 URL，也不要记录包含手机号、验证码或模板编码的完整 URL。

受理成功响应示例：

```json
{"code": 200, "msg": "请求成功", "request_id": "dwezjVRDoe5jgLkR"}
```

`code: 200` 仅表示请求已受理，不表示短信已送达。

## 查询状态

`POST https://push.spug.cc/request/query`

JSON 必填字段：`token`、`request_id`。

响应中的 `data` 数组使用以下状态值：

| 状态值 | 含义 |
| --- | --- |
| `0` | 待发送 |
| `1` | 发送中 |
| `2` | 成功 |
| `3` | 发送失败，查看 `reason` |
| `4` | 审核中，等待平台审核，不要立即重复提交 |

同一 `request_id` 可能包含多个手机号的明细，应逐条处理。对状态 `0`、`1`、`4` 使用间隔查询，避免紧密轮询。`400` 通常表示缺少 `token` 或 `request_id`；`403` 表示 Token 无效或账户已暂停；`404` 表示记录不存在或已过期；`429` 表示 IP 限流。

## 查询历史

`POST https://push.spug.cc/request/history`

| 字段 | 必填 | 说明 |
| --- | --- | --- |
| `token` | 是 | 账户开发者 Token |
| `type` | 是 | 本 skill 固定传 `sms` |
| `start_date` | 否 | 格式为 `YYYY-MM-DD` |
| `end_date` | 否 | 格式为 `YYYY-MM-DD`，默认今天 |
| `page` | 否 | 默认 `1` |
| `page_size` | 否 | 默认 `20`，范围 1–100 |

不传日期时，默认查询包含今天在内的最近 30 个自然日。开始和结束日期都必须位于可查询范围内，结束日期不能晚于今天，开始日期不能晚于结束日期。响应 `data` 包含 `records`、`total`、`page`、`page_size`、`start_date` 和 `end_date`；短信记录包含 `request_id`、脱敏 `target`、`fee`、`status`、`status_alias`、`reason` 和 `created_at`。
