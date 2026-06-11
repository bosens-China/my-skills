# 我的 Agent Skills

这是我的个人 Agent Skills 仓库，用来收集日常开发、写作和工作中常用的技能。

## 技能列表

| 技能 | 作用 |
| --- | --- |
| [chinese-article-writing](./skills/chinese-article-writing/) | 规范中文技术文章、博客和技术分享的结构、语气、术语、标点与排版。 |
| [file-line-audit](./skills/file-line-audit/) | 扫描仓库中的源码文件，找出达到指定行数阈值的超长文件。 |
| [git-step-commit](./skills/git-step-commit/) | 分析 Git 更改，先给出分批提交计划，再按用户确认的批次提交。 |

## 使用方式

克隆仓库：

```bash
git clone https://github.com/bosens-China/my-skills.git
```

按需复制 `skills/` 下的某个技能目录到你的 Agent 客户端技能目录中，例如复制 `skills/git-step-commit`。

如果客户端支持直接从仓库加载 skills，也可以把本仓库的 `skills/` 目录作为技能来源。

## 开发

当前 `file-line-audit` 的源码在 `packages/file-line-audit/`，打包产物输出到 `skills/file-line-audit/scripts/`。

运行测试：

```bash
pnpm test
```

重新构建 `file-line-audit` 二进制：

```bash
pnpm run build:file-line-audit
```

## 协议

[MIT](./LICENSE)
