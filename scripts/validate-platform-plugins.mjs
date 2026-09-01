import { access, readFile, readdir } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const pluginName = "yliu-skills";
const pluginRoot = path.join(root, "plugins", pluginName);

async function readJson(relativePath) {
  const content = await readFile(path.join(root, relativePath), "utf8");
  if (content.includes("{{VERSION}}")) {
    throw new Error(`${relativePath} 仍然包含未替换的版本占位符`);
  }
  return JSON.parse(content);
}

async function listSkills(directory) {
  const entries = await readdir(directory, { withFileTypes: true });
  const names = [];
  for (const entry of entries) {
    if (!entry.isDirectory() || entry.name.startsWith(".")) continue;
    await access(path.join(directory, entry.name, "SKILL.md"));
    names.push(entry.name);
  }
  return names.sort();
}

function assert(condition, message) {
  if (!condition) throw new Error(message);
}

const sourceSkills = await listSkills(path.join(root, "skills"));
const generatedSkills = await listSkills(path.join(pluginRoot, "skills"));
assert(
  JSON.stringify(generatedSkills) === JSON.stringify(sourceSkills),
  "共享插件的技能列表与 skills/ 不一致",
);

const codexManifest = await readJson("plugins/yliu-skills/.codex-plugin/plugin.json");
const codexMarketplace = await readJson(".agents/plugins/marketplace.json");
assert(codexManifest.name === pluginName, "Codex 插件名称不一致");
assert(codexMarketplace.plugins[0]?.name === pluginName, "Codex marketplace 未引用插件");
assert(
  codexMarketplace.plugins[0]?.source?.path === "./plugins/yliu-skills",
  "Codex marketplace 插件路径不正确",
);

const cursorManifest = await readJson("plugins/yliu-skills/.cursor-plugin/plugin.json");
const cursorMarketplace = await readJson(".cursor-plugin/marketplace.json");
assert(cursorManifest.name === pluginName, "Cursor 插件名称不一致");
assert(cursorManifest.skills === "./skills/", "Cursor 技能路径不正确");
assert(cursorMarketplace.plugins[0]?.source === "plugins/yliu-skills", "Cursor marketplace 插件路径不正确");

const claudeManifest = await readJson("plugins/yliu-skills/.claude-plugin/plugin.json");
const claudeMarketplace = await readJson(".claude-plugin/marketplace.json");
assert(claudeManifest.name === pluginName, "Claude Code 插件名称不一致");
assert(claudeMarketplace.plugins[0]?.source === "./plugins/yliu-skills", "Claude Code marketplace 插件路径不正确");

const generatedVersions = [
  codexManifest.version,
  cursorManifest.version,
  cursorMarketplace.metadata?.version,
  cursorMarketplace.plugins[0]?.version,
  claudeManifest.version,
  claudeMarketplace.version,
  claudeMarketplace.plugins[0]?.version,
];
assert(generatedVersions.every((version) => version === generatedVersions[0]), "各平台生成版本不一致");

for (const skill of sourceSkills) {
  try {
    await access(path.join(pluginRoot, "skills", skill, "agents", "openai.yaml"));
    throw new Error("共享插件不应包含 Codex 专属的 agents/openai.yaml");
  } catch (error) {
    if (error.code !== "ENOENT") throw error;
  }
}

console.log(`单仓库多平台插件校验通过，共 ${sourceSkills.length} 个技能`);
