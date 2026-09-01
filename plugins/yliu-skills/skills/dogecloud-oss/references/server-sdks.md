# 服务端 SDK

优先选与项目语言一致的多吉云服务端 SDK，并复用项目现有 HTTP/S3 依赖。永久密钥只从服务端环境变量或密钥系统读取。

## 官方语言文档

- [PHP](https://docs.dogecloud.com/oss/sdk-full-php)
- [Python](https://docs.dogecloud.com/oss/sdk-full-python)
- [Node.js](https://docs.dogecloud.com/oss/sdk-full-nodejs)
- [Java](https://docs.dogecloud.com/oss/sdk-full-java)
- [Go](https://docs.dogecloud.com/oss/sdk-full-go)

## 任务选择

- 客户端直传：服务端 SDK 调用 `/auth/tmp_token.json`，返回最小授权的临时凭证和动态 Bucket 信息。
- 小文件服务端上传：可使用 SDK 的 S3 实例或 [简单上传 API](https://docs.dogecloud.com/oss/api-upload-put)。
- 大文件、进度或分片：使用语言对应的 S3 SDK 能力，不把整个文件读入内存。
- Bucket、文件管理、远程抓取或统计：优先用 SDK 已封装的方法；没有封装再按 [api-reference.md](api-reference.md) 调 REST API。

## 签发临时密钥

1. 校验业务用户和上传意图。
2. 服务端生成目标 key。
3. 请求体使用 `channel: OSS_UPLOAD`，`scopes` 只包含目标 Bucket 与精确 key，或确有批量需求时的窄前缀。
4. 返回 `Credentials`、`ExpiredAt`、匹配 Bucket 的 `s3Bucket`/`s3Endpoint` 和目标 key。
5. 不透传完整上游错误、永久密钥或签名材料。

`OSS_FULL`/`OSS_CUSTOM` 只用于受控的服务端管理流程，不交给普通客户端。
