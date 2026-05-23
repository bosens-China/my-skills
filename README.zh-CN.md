# File Line Audit

[English](./README.md)

一个专门为 Agent Skills 设计的工具，用于审计仓库中超过指定行数阈值的源代码文件。Agent 通过命令行参数直接传入扫描范围，不再依赖默认配置文件。

## 安装

将此技能添加至你的 Agent：

```bash
npx skills add bosens-China/file-line-audit
```

## 使用方式

在仓库根目录运行打包好的二进制，并显式传入扫描参数：

```bash
line-audit \
  --threshold 400 \
  --include "src/**/*.{ts,tsx,js,jsx,vue}" \
  --include "apps/**/*.{ts,tsx,js,jsx,vue}" \
  --exclude "dist/" \
  --exclude "build/"
```

也可以直接传入 JSON 对象：

```bash
line-audit --json '{
  "threshold": 400,
  "include": ["src/**/*.go", "internal/**/*.go"],
  "exclude": ["vendor/"]
}'
```

### 参数说明

- `--include` / `-i`：必填，glob 匹配模式；可重复传入多个
- `--exclude` / `-e`：可选，额外排除规则，语法与 `.gitignore` 一致
- `--threshold` / `-t`：输出阈值，默认 `400`
- `--json` / `-j`：可选，传入包含 `threshold`、`include`、`exclude` 的 JSON 对象

仓库中的 `.gitignore` 规则始终生效，额外 `--exclude` 规则会在其后继续过滤。

## 开源协议

MIT
