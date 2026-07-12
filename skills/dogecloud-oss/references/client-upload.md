# 客户端直传

客户端只使用服务端签发的临时凭证。响应至少包含 `credentials`（`accessKeyId`、`secretAccessKey`、`sessionToken`）、`s3Bucket`、`s3Endpoint`、授权 key/prefix 和过期时间。Bucket 与 Endpoint 必须取当前响应。

## 平台路由

- Web/JavaScript：[Web JavaScript S3](https://docs.dogecloud.com/oss/sdk-s3upload-js)。使用页面指定的 S3 客户端配置、上传和进度回调。优先复用项目已安装的 AWS S3 客户端。
- Android：[Android S3](https://docs.dogecloud.com/oss/sdk-s3upload-android)。按页面初始化 S3/TransferUtility，并保持 key 在服务端授权范围内。
- iOS：[iOS S3](https://docs.dogecloud.com/oss/sdk-s3upload-ios)。按页面初始化 AWSS3/TransferUtility；列表、下载、复制、删除默认不要放客户端。
- 微信小程序：[微信小程序上传](https://docs.dogecloud.com/oss/suggestion-wxapp-upload)。浏览器 SDK 不适用，按文档使用 `wx.uploadFile` 和签名字段，并配置上传合法域名。
- uni-app：[uni-app 上传](https://docs.dogecloud.com/oss/suggestion-uniapp-upload)。按文档使用 `uni.uploadFile`；不同运行端能力不一致时分别验证。

## 实现检查

- 服务端先生成最终 key，再按单 key 或用户隔离前缀签发 `OSS_UPLOAD`。
- 不把 `*` scope 或永久密钥发给客户端。
- 不从配置硬编码底层 Bucket/Endpoint。
- 上传 key 必须匹配服务端返回的授权范围，否则应预期 403。
- 需要进度、取消、重试或分片时，使用所选平台 SDK 的原生能力，不自造传输协议。
- 只有明确需要客户端列表、下载、复制、删除时，才按 Android/iOS 文档增加对应 `allowActions`；默认由服务端执行。
- 微信小程序的上传域名由 `s3Bucket` 和 `s3Endpoint` 组合，按官方页面示例加入合法域名。

临时凭证签发细节见 [服务端获取上传用的临时密钥](https://docs.dogecloud.com/oss/manual-tmp-token)。
