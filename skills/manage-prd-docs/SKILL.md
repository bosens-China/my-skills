---
name: manage-prd-docs
description: "Discuss, materialize, and maintain product requirements as phased PRD and todo documents under docs/. Use when the user asks to 讨论需求、讨论产品、讨论 PRD、更新需求、落盘, finalize a PRD, plan product iterations, or synchronize confirmed requirement changes. Trigger during discussion, but write only after requirements are mature and the user explicitly authorizes the initial landing."
---

# Manage PRD Docs

围绕产品需求进行讨论，在需求成型后按迭代路径落盘，并持续同步已确认的变更。

## 区分讨论与落盘

- 用户提到讨论需求、产品、PRD、更新需求或落盘时，进入本工作流。
- 触发技能不等于允许写文件。需求未成型时只继续讨论，不创建或修改文档。
- 判断需求是否成型：目标、范围、核心需求和验收标准已经明确，且不存在会改变迭代路径的关键未决项。
- 首次落盘必须等待用户明确说“落盘”“写入”“定稿”或同等指令。即使用户要求落盘，需求仍未成型时也先指出缺口。
- 已有本技能维护的文档后，直接同步用户已经确认的变更；探索性想法和未定案内容不写入。无法判断是否确认时先询问。

## 准备文档

1. 定位仓库根目录，默认使用根目录下的 `docs/`。
2. 读取已有 `docs/index.md`、阶段目录、需求文档和 Todo，增量修改并保留无关的人工内容。
3. 使用当前本地日期，格式为 `YYYY-MM-DD`。
4. 把 PRD 视为当前有效需求，不在正文保留已经失效的旧需求；Todo 按下文规则保留变更历史。

## 拆分产品阶段

- 按用户价值、依赖关系、风险和最小可用闭环规划迭代路径。
- 动态决定阶段数量；允许一个或多个阶段，不把用户列出的每个需求点机械地变成独立阶段。
- 把基础能力和最小闭环放在前面，把增强体验、扩展能力和规模化工作放在后续阶段。
- 使用连续编号和稳定的英文 slug 命名目录，例如 `phase-01-mvp/`、`phase-02-core-flow/`。
- 每个阶段至少包含 `requirements.md` 和 `todo.md`。

默认结构：

```text
docs/
├── index.md
├── phase-01-mvp/
│   ├── requirements.md
│   └── todo.md
└── phase-02-core-flow/
    ├── requirements.md
    └── todo.md
```

## 编写最小 PRD

每个需求文件使用最小通用结构：

```markdown
# 第一阶段：MVP

- 创建日期：YYYY-MM-DD
- 更新日期：YYYY-MM-DD

## 背景与目标

## 需求范围

## 功能需求

## 验收标准
```

- 新建时让创建日期和更新日期相同。
- 后续只更新实际发生内容变化的需求文件及其更新日期，保留创建日期。
- 写清可验证的需求和验收标准，不写 `TBD`、猜测或尚未确认的方案。

## 维护 Todo

- 使用 `- [ ]` 表示待完成任务，使用 `- [x]` 表示已完成任务。
- 让任务对应当前阶段的已确认需求，并按合理实施顺序排列。
- 保留已完成任务，不从文件中删除。
- 需求变更、任务取消或任务内容需要修改时，不直接改写或删除原任务；给原任务添加删除线、变更日期和简短原因，再新增替代任务。
- 已完成任务后来失效时，保留 `[x]` 状态并添加删除线。

```markdown
- [ ] 待完成任务
- [x] 已完成任务
- [ ] ~~原任务~~（YYYY-MM-DD：需求变更）
- [ ] 变更后的新任务
- [x] ~~原实现方式~~（YYYY-MM-DD：方案已调整）
```

不要拆分 `todo.md`。

## 维护总索引

让 `docs/index.md` 按产品迭代顺序仅维护阶段链接列表：

```markdown
- [第一阶段：MVP](./phase-01-mvp/requirements.md)
- [第二阶段：核心流程](./phase-02-core-flow/requirements.md)
```

新增、删除、改名或调整阶段顺序时同步更新索引，并确认每个链接指向存在的文件。

## 审计并拆分长文档

每次首次落盘或同步变更后，使用 `$file-line-audit` 在仓库根目录分析需求文件：

- 阈值设为 `400`。
- include 范围设为 `docs/**/requirements*.md`。
- 不扫描或拆分 `todo.md`。

把 400 行作为推荐拆分线，不在固定行号机械截断。达到阈值时：

1. 按业务模块或完整主题识别自然边界。
2. 保留 `requirements.md` 作为阶段概览和子文档索引。
3. 在同一阶段目录创建 `requirements-<topic>.md`，不增加更深目录。
4. 在 `requirements.md` 中链接所有拆分文件，并让 `docs/index.md` 继续指向该阶段的 `requirements.md`。
5. 再次使用 `$file-line-audit` 确认拆分结果。

## 完成检查

- 只写入已确认且成型的需求。
- 阶段顺序符合产品迭代路径。
- 每个阶段都包含 `requirements.md` 和 `todo.md`。
- 创建日期、更新日期和 Todo 状态正确。
- `docs/index.md` 顺序正确且没有失效链接。
- PRD、Todo 和索引反映同一份当前需求。
- 已完成 `$file-line-audit` 检查，并按需拆分长需求文档。
