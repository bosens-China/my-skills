# 仓库维护约定

这个仓库是所有技能的唯一来源。日常只维护 `skills/` 和 `platforms/`，不要直接修改 CI 生成的插件目录。

## 我们怎样支持不同平台

- `skills/` 保存尽量通用的 Agent Skills，Codex 是目前的主要参考环境。
- `platforms/` 只保存 Codex、Cursor 和 Claude Code 需要的清单模板，不复制技能正文。
- `scripts/build-platform-plugins.mjs` 在仓库根部生成各平台的 marketplace，并生成共享的 `plugins/yliu-skills/`。
- `.agents/`、`.cursor-plugin/`、`.claude-plugin/` 和 `plugins/yliu-skills/` 都是生成结果，不要手工修改。
- Antigravity 直接使用根部的 `skills/`。公开规范没有稳定说明完整 manifest，因此不要自行猜测字段。

新增或修改技能时，优先保持 `SKILL.md` 的正文与平台无关。Codex 专属的展示信息可以继续放在技能目录的 `agents/openai.yaml` 中；共享插件副本会忽略这个文件。

## 本地检查

```sh
npm run plugins:build
npm run plugins:check
```

生成后的目录结构如下：

```text
.agents/plugins/marketplace.json
.cursor-plugin/marketplace.json
.claude-plugin/marketplace.json
plugins/yliu-skills/
```

## 发布方式

Pull Request 会构建并校验插件。代码进入 `main` 后，GitHub Actions 使用当前仓库自带的 `GITHUB_TOKEN` 重新生成上述目录并提交回来，不需要额外的仓库变量或 token。

CI 生成的插件版本由 `package.json` 版本和源提交 SHA 组成。生成目录中的人工修改会在下次构建时丢失，需要修改内容时请改 `skills/`、`platforms/` 或构建脚本。
