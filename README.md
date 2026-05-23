# File Line Audit

[简体中文](./README.zh-CN.md)

An Agent Skill for auditing source files that exceed a line-count threshold in any repository. The agent passes scan parameters directly on the command line instead of relying on a local config file.

## Installation

Install and add this skill to your agent:

```bash
npx skills add bosens-China/file-line-audit
```

## Usage

Run the packaged binary from the repository root and pass the scan scope explicitly:

```bash
line-audit \
  --threshold 400 \
  --include "src/**/*.{ts,tsx,js,jsx,vue}" \
  --include "apps/**/*.{ts,tsx,js,jsx,vue}" \
  --exclude "dist/" \
  --exclude "build/"
```

You can also pass a JSON object:

```bash
line-audit --json '{
  "threshold": 400,
  "include": ["src/**/*.go", "internal/**/*.go"],
  "exclude": ["vendor/"]
}'
```

### Parameters

- `--include` / `-i`: required glob pattern; repeat for multiple patterns
- `--exclude` / `-e`: optional extra ignore rule in `.gitignore` syntax
- `--threshold` / `-t`: minimum line count to report, default `400`
- `--json` / `-j`: optional JSON object with `threshold`, `include`, and `exclude`

Repository `.gitignore` rules are always applied before any extra exclude patterns.

## License

MIT
