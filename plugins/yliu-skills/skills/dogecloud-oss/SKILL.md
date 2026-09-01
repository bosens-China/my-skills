---
name: dogecloud-oss
description: Provide implementation guidance and official documentation routes for integrating DogeCloud OSS while writing, modifying, or debugging code. Use for Web/JavaScript, Android, iOS, WeChat Mini Programs, uni-app, PHP, Python, Node.js, Java, and Go integrations involving SDK setup, temporary credentials, S3-compatible uploads, buckets, file management, remote fetches, or usage statistics APIs.
---

# 多吉云 OSS

在编写、修改或排查代码时，依据多吉云官方文档快速选择正确的 SDK、上传方式和 API，并完成最小、安全、可维护的 OSS 接入。优先复用项目现有框架、HTTP 客户端和 S3 SDK。

## 选择接入路径

1. 先检查项目语言、框架、依赖和现有网络/上传代码，再确认本次需要上传、临时密钥、Bucket、文件管理、远程抓取还是统计能力。
2. 浏览器、Android、iOS 客户端直传：读取 [references/client-upload.md](references/client-upload.md)，由业务服务端签发临时密钥。
3. 微信小程序或 uni-app：读取同一客户端参考，使用平台 `uploadFile` 流程，不照搬浏览器 S3 SDK。
4. 服务端上传或文件管理：读取 [references/server-sdks.md](references/server-sdks.md)，选择项目语言对应的官方 SDK；小文件也可使用简单上传 API。
5. 无合适 SDK、需要 Bucket/文件/统计接口或排查签名：读取 [references/api-reference.md](references/api-reference.md)。
6. 构建图床时同时读取 [references/image-hosting.md](references/image-hosting.md)。

只读取当前任务需要的 reference。实现前打开其中链接的官方页面，核对字段、限制和示例；不要凭记忆补造参数。

## 编码流程

1. 找到当前任务对应的 reference，只打开其中直接相关的官方页面。
2. 按项目现有技术栈选择 SDK 或 REST API，不替换用户已选框架，不新增无必要依赖。
3. 从官方页面核对安装方式、初始化参数、请求字段、响应字段、限制和错误处理。
4. 直接修改或生成可运行代码，并沿用仓库的配置、类型、错误处理和测试风格。
5. 客户端上传时，由服务端签发最小权限临时密钥；服务端管理时，使用对应语言 SDK 或 API。
6. 为鉴权、key/scope 生成、签名、上传或分页逻辑运行一个最小验证。
7. 在交付中指出使用的官方页面、需要配置的环境变量，以及仍需用户提供的 Bucket 或访问域名。

## 安全边界

- 永远不要把永久密钥写入客户端、仓库、日志、示例响应或错误信息。
- 临时密钥最长有效期为 2 小时；按响应过期时间使用，不自行延长或长期缓存。
- 不硬编码 `s3Bucket` 或 `s3Endpoint`；它们可能变化，使用临时密钥接口的当前响应。
- 客户端默认只授予单一对象 key；确需批量上传才授予用户隔离的窄前缀。不要使用 `*`。
- 先完成业务身份认证、权限检查、限流和 key 归属校验，再签发临时密钥。
- 在服务端校验文件大小、扩展名/MIME 白名单和生成的 key；不要信任客户端文件名或 MIME。
- 列表、复制、删除和下载默认放在服务端。只有用户明确需要客户端管理时，才添加最少的 `allowActions`。
- 对生命周期删除等不可逆操作保留显式确认和错误处理。
- 不关闭 TLS 证书验证，不把官方示例中的演示性占位密钥当作可运行配置。

## 输出要求

- 优先直接给出或完成接入代码，并附上实际使用的官方文档链接。
- 明确列出需要用户配置的环境变量、Bucket、CDN/自定义域名和 CORS/平台合法域名。
- 不声称已创建 Bucket、域名或凭证，除非实际工具结果证明已完成。
- 文档与项目依赖版本冲突时，先查该依赖当前官方文档，再适配多吉云返回的 S3 兼容参数。
