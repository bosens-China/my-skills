package audit

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

const DefaultThreshold = 400

type Config struct {
	Threshold int      `json:"threshold"`
	Include   []string `json:"include"`
	Exclude   []string `json:"exclude"`
}

func ParseConfig(args []string) (Config, error) {
	config := Config{
		Threshold: DefaultThreshold,
	}

	for index := 0; index < len(args); index++ {
		arg := args[index]

		switch arg {
		case "--help", "-h":
			return Config{}, errHelp
		case "--json", "-j":
			if index+1 >= len(args) {
				return Config{}, errors.New("--json requires a JSON object")
			}

			if err := json.Unmarshal([]byte(args[index+1]), &config); err != nil {
				return Config{}, fmt.Errorf("parse --json: %w", err)
			}

			index++
		case "--threshold", "-t":
			if index+1 >= len(args) {
				return Config{}, errors.New("--threshold requires a positive integer")
			}

			threshold, err := strconv.Atoi(args[index+1])
			if err != nil || threshold <= 0 {
				return Config{}, errors.New("--threshold requires a positive integer")
			}

			config.Threshold = threshold
			index++
		case "--include", "-i":
			if index+1 >= len(args) {
				return Config{}, errors.New("--include requires a glob pattern")
			}

			config.Include = append(config.Include, args[index+1])
			index++
		case "--exclude", "-e":
			if index+1 >= len(args) {
				return Config{}, errors.New("--exclude requires a glob pattern")
			}

			config.Exclude = append(config.Exclude, args[index+1])
			index++
		default:
			return Config{}, fmt.Errorf("unknown argument %q", arg)
		}
	}

	config.Include = normalizePatterns(config.Include)
	config.Exclude = normalizePatterns(config.Exclude)

	if len(config.Include) == 0 {
		return Config{}, errors.New("at least one --include pattern is required")
	}

	for _, pattern := range config.Include {
		if !doublestar.ValidatePattern(pattern) {
			return Config{}, fmt.Errorf("invalid include pattern %q", pattern)
		}
	}

	for _, pattern := range config.Exclude {
		if !doublestar.ValidatePattern(pattern) {
			return Config{}, fmt.Errorf("invalid exclude pattern %q", pattern)
		}
	}

	return config, nil
}

var errHelp = errors.New("help requested")

func HelpText() string {
	return strings.TrimSpace(`
File Line Audit

Usage:
  line-audit --include <glob> [--include <glob> ...] [options]

Options:
  -i, --include <glob>    Include glob pattern (required, repeatable)
  -e, --exclude <glob>    Extra exclude glob in .gitignore syntax (repeatable)
  -t, --threshold <n>     Minimum line count to report (default: 400)
  -j, --json <object>     Full config as JSON: {"threshold":400,"include":[],"exclude":[]}
  -h, --help              Show this help message

Examples:
  line-audit --include "src/**/*.{ts,tsx}" --include "apps/**/*.{ts,tsx}" --threshold 400
  line-audit --json '{"threshold":400,"include":["src/**/*.go"],"exclude":["dist/"]}'
`)
}

func normalizePatterns(patterns []string) []string {
	normalized := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		trimmed := strings.TrimSpace(pattern)
		if trimmed == "" {
			continue
		}

		normalized = append(normalized, filepath.ToSlash(trimmed))
	}

	return normalized
}
