# REST API 路由

需要具体字段、频率限制和响应结构时，打开对应官方页面；不要仅凭本索引实现。

## 通用规则

- 基础域名与响应错误：[API 介绍](https://docs.dogecloud.com/oss/api-introduction)
- 鉴权签名：[AccessToken 验证机制](https://docs.dogecloud.com/oss/api-access-token)
- 数据编码：[数据格式](https://docs.dogecloud.com/oss/api-format)
- 临时凭证：[获取临时密钥](https://docs.dogecloud.com/oss/api-tmp-token)

服务端鉴权核心为 `Authorization: TOKEN <AccessKey>:<hex-hmac-sha1>`。签名原文是包含 QueryString 的 API path、换行、原始请求 Body；Body 字节必须与实际发送内容完全一致。优先复用官方服务端 SDK，减少编码顺序和 JSON 序列化差异导致的签名错误。

所有响应都检查业务 `code`，不要只看 HTTP 状态。处理未授权、签名失败、参数错误、频率过高、无权限、格式不支持和资源不存在。

## Bucket

- [获取列表](https://docs.dogecloud.com/oss/api-bucket-list)
- [创建](https://docs.dogecloud.com/oss/api-bucket-create)
- [删除](https://docs.dogecloud.com/oss/api-bucket-delete)

创建/删除属于管理操作，只有用户明确要求且参数已确认时执行。

## 上传与远程抓取

- [简单上传小文件](https://docs.dogecloud.com/oss/api-upload-put)：`POST /oss/upload/put.json`
- [异步抓取远程资源](https://docs.dogecloud.com/oss/api-fetch)：`POST /oss/fetch.json`
- [查询抓取任务](https://docs.dogecloud.com/oss/api-fetch-query)

简单上传的文件内容是请求 Body；签名必须覆盖同一原始 Body。远程抓取 URL 属于不可信输入时，先做业务授权和 SSRF 风险控制。

## 文件管理

- [列表](https://docs.dogecloud.com/oss/api-file-list)
- [移动/重命名](https://docs.dogecloud.com/oss/api-file-move)
- [删除](https://docs.dogecloud.com/oss/api-file-delete)
- [生命周期](https://docs.dogecloud.com/oss/api-file-lifecycle)
- [信息](https://docs.dogecloud.com/oss/api-file-info)
- [设置 MIME](https://docs.dogecloud.com/oss/api-file-mime)
- [复制](https://docs.dogecloud.com/oss/api-file-copy)
- [预取](https://docs.dogecloud.com/oss/api-file-prefetch)

列表按官方 `continue` 游标分页，不把游标当页码。生命周期和删除不可恢复；保留显式确认。MIME 参数按数据格式文档做 URLSafeBase64 编码。

## 统计

- [请求次数](https://docs.dogecloud.com/oss/api-stat-count)
- [存储量](https://docs.dogecloud.com/oss/api-stat-storage)
- [流量](https://docs.dogecloud.com/oss/api-stat-traffic)
- [带宽](https://docs.dogecloud.com/oss/api-stat-bandwidth)

严格使用页面要求的时间范围、粒度和单位；展示前标注单位与时区。
