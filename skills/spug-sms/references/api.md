# Spug Push API 参考

编写接入代码时以本文件的请求结构为准，并使用项目已有的 HTTP 客户端。模板编码与查询 `token` 是两种不同凭据。

## 发送短信

`POST https://push.spug.cc/sms/{template_code}`

发送 JSON 数据，其中必须包含 `to`，并附带该模板配置的动态变量。不要在代码中固定变量字段。多个手机号使用英文逗号分隔。受理成功响应示例：

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
| `2` | 已送达 |
| `3` | 发送失败，查看 `reason` |

查询接口支持最近 30 天的短信、邮件和语音记录。状态更新可能延迟 1–3 分钟。`400` 通常表示缺少 token；`404` 表示记录不存在、已过期或请求 ID 错误。

## 查询历史

`POST https://push.spug.cc/request/history`

| 字段 | 必填 | 说明 |
| --- | --- | --- |
| `token` | 是 | 用户 ID/token |
| `type` | 否 | 仅支持 `sms` 或 `mail` |
| `start_date` | 否 | 格式为 `YYYY-MM-DD`，默认 30 天前 |
| `end_date` | 否 | 格式为 `YYYY-MM-DD`，默认今天 |
| `page` | 否 | 默认 `1` |
| `page_size` | 否 | 默认 `20`，范围 1–100 |

历史接口仅支持最近 30 天内的短信和邮件，不包含语音、机器人或公众号消息。
