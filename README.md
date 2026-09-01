# my-skills

这个是我目前工作沉淀的相关 skills，希望也能对你有所启迪。

## 技能列表

| 技能                                                         | 作用                                                             |
| ------------------------------------------------------------ | ---------------------------------------------------------------- |
| [apply-yliu-product-development-practices](./skills/apply-yliu-product-development-practices/) | 应用 Yliu 的产品开发方法论，覆盖项目、架构、实现、数据、界面、测试和运行。 |
| [chinese-article-writing](./skills/chinese-article-writing/) | 规范中文技术文章、博客和技术分享的结构、语气、术语、标点与排版。 |
| [dogecloud-oss](./skills/dogecloud-oss/)                     | 编写多吉云 OSS SDK、文件上传、图床和文件管理接入代码。            |
| [file-line-audit](./skills/file-line-audit/)                 | 扫描仓库中的源码文件，找出达到指定行数阈值的超长文件。           |
| [git-step-commit](./skills/git-step-commit/)                 | 分析 Git 更改，先给出分批提交计划，再按用户确认的批次提交。      |
| [manage-prd-docs](./skills/manage-prd-docs/)                 | 维护现行产品决策，并用临时工作目录推进和收口进行中需求。         |
| [publish-npm-packages](./skills/publish-npm-packages/)       | 配置符合当前 npm OIDC 政策的单包、组织包与 monorepo 安全发布流程。 |
| [spug-sms](./skills/spug-sms/)                               | 编写 Spug 短信发送、状态查询、发送记录与错误处理接入代码。        |

## 使用方式

通过 Skills CLI 安装，并按提示选择需要使用的平台：

```sh
npx skills add bosens-China/my-skills
```

也可以作为 Codex marketplace 安装：

```sh
codex plugin marketplace add bosens-China/my-skills
codex plugin add yliu-skills@yliu-marketplace
```

Claude Code 可以从同一仓库安装：

```sh
claude plugin marketplace add bosens-China/my-skills
claude plugin install yliu-skills@yliu-marketplace
```

Cursor 团队可以在 Dashboard 的 Plugins 页面把这个仓库导入 Team Marketplace。Antigravity 当前直接使用仓库中的通用 `skills/`。

## 协议

[MIT](./LICENSE)
