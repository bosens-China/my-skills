import { cp, mkdir, readFile, readdir, rm, writeFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const skillsRoot = path.join(root, "skills");
const pluginName = "yliu-skills";
const pluginRoot = path.join(root, "plugins", pluginName);

function readArgument(name) {
  const index = process.argv.indexOf(name);
  return index === -1 ? undefined : process.argv[index + 1];
}

const packageJson = JSON.parse(await readFile(path.join(root, "package.json"), "utf8"));
const version = readArgument("--version") ?? packageJson.version;

if (!/^\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?(?:\+[0-9A-Za-z.-]+)?$/.test(version)) {
  throw new Error(`版本必须是 semver，当前值为 ${version}`);
}

const skillEntries = (await readdir(skillsRoot, { withFileTypes: true }))
  .filter((entry) => entry.isDirectory() && !entry.name.startsWith("."))
  .sort((left, right) => left.name.localeCompare(right.name));

if (skillEntries.length === 0) {
  throw new Error("skills/ 中没有可发布的技能");
}

for (const entry of skillEntries) {
  await readFile(path.join(skillsRoot, entry.name, "SKILL.md"), "utf8");
}

async function render(templatePath, destinationPath) {
  const template = await readFile(path.join(root, templatePath), "utf8");
  const content = template.replaceAll("{{VERSION}}", version);
  await mkdir(path.dirname(destinationPath), { recursive: true });
  await writeFile(destinationPath, content);
}

function keepPortableSkillFile(source) {
  const relative = path.relative(skillsRoot, source).split(path.sep);
  return !(
    relative.length === 3 &&
    relative[1] === "agents" &&
    relative[2] === "openai.yaml"
  );
}

async function copySkills(destination) {
  await cp(skillsRoot, destination, {
    recursive: true,
    filter: keepPortableSkillFile,
  });
}

await rm(pluginRoot, { recursive: true, force: true });

await render(
  "platforms/codex/marketplace.json",
  path.join(root, ".agents", "plugins", "marketplace.json"),
);
await render(
  "platforms/cursor/marketplace.json",
  path.join(root, ".cursor-plugin", "marketplace.json"),
);
await render(
  "platforms/claude-code/marketplace.json",
  path.join(root, ".claude-plugin", "marketplace.json"),
);

await render(
  "platforms/codex/plugin.json",
  path.join(pluginRoot, ".codex-plugin", "plugin.json"),
);
await render(
  "platforms/cursor/plugin.json",
  path.join(pluginRoot, ".cursor-plugin", "plugin.json"),
);
await render(
  "platforms/claude-code/plugin.json",
  path.join(pluginRoot, ".claude-plugin", "plugin.json"),
);

await copySkills(path.join(pluginRoot, "skills"));
await cp(path.join(root, "LICENSE"), path.join(pluginRoot, "LICENSE"));

console.log(`已生成 ${skillEntries.length} 个技能的单仓库多平台插件，版本号 ${version}`);
