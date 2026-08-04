# 产品 PRD

## Git 分步提交偏好

- 当前决策：`git-step-commit` 支持全局和当前项目的持久偏好。提交模式为 `default`、`review` 或 `direct`；提交语言为 `auto` 或指定语言。按“本轮明确指令 → 项目配置 → 全局配置 → Skill 内置行为”的顺序按字段求值。
- 为什么：使用 Git 原生的全局和仓库配置，可在 Windows、macOS 和 Linux 上共用同一工作流，不引入 Python、Bash、Go 或额外配置文件。
- 边界 / 非目标：本轮明确指令可以临时覆盖持久偏好，但不修改配置。`direct` 只省略确认，不绕过凭据、冲突、失败验证或归属不明等安全停止条件。项目偏好只存于 Git 本地配置，不写入可提交的 `skills/` 目录。
- 权威入口：[Git Step Commit Skill](../skills/git-step-commit/SKILL.md)。
