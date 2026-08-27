package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const threshold = 3

type app struct{ stateDir string }
type settings struct {
	Enabled bool `json:"enabled"`
}
type projectState struct {
	RecommendedSuccesses int `json:"recommended_successes"`
}
type project struct{ name, id, stateFile string }

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: commit-memory <enable|disable|status|record-direct|forget> [--repo PATH]")
		return
	}
	dir, err := defaultStateDir()
	if err == nil {
		var result any
		result, err = run(os.Args[1:], dir)
		if err == nil {
			err = json.NewEncoder(os.Stdout).Encode(result)
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "commit-memory failed: %v\n", err)
		os.Exit(2)
	}
}

func run(args []string, stateDir string) (any, error) {
	if len(args) == 0 {
		return nil, errors.New("missing command")
	}
	tool := app{stateDir: stateDir}
	switch args[0] {
	case "enable", "disable":
		if len(args) != 1 {
			return nil, fmt.Errorf("%s does not accept arguments", args[0])
		}
		return tool.setEnabled(args[0] == "enable")
	case "status", "record-direct", "forget":
		repo, err := repoArg(args[1:])
		if err != nil {
			return nil, err
		}
		switch args[0] {
		case "status":
			return tool.status(repo)
		case "record-direct":
			return tool.record(repo)
		default:
			return tool.forget(repo)
		}
	default:
		return nil, fmt.Errorf("unknown command %q", args[0])
	}
}

func repoArg(args []string) (string, error) {
	if len(args) == 0 {
		return ".", nil
	}
	if len(args) == 2 && args[0] == "--repo" && args[1] != "" {
		return args[1], nil
	}
	return "", errors.New("expected --repo PATH")
}

func defaultStateDir() (string, error) {
	if dir := os.Getenv("GIT_STEP_COMMIT_STATE_DIR"); dir != "" {
		return filepath.Abs(dir)
	}
	if dir := os.Getenv("XDG_STATE_HOME"); dir != "" {
		return filepath.Abs(filepath.Join(dir, "git-step-commit"))
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "state", "git-step-commit"), nil
}

func (tool app) setEnabled(enabled bool) (map[string]any, error) {
	if err := writeJSON(filepath.Join(tool.stateDir, "settings.json"), settings{Enabled: enabled}); err != nil {
		return nil, err
	}
	return map[string]any{"enabled": enabled, "state_directory": tool.stateDir}, nil
}

func (tool app) status(repo string) (map[string]any, error) {
	p, err := tool.project(repo)
	if err != nil {
		return nil, err
	}
	enabled, err := tool.enabled()
	if err != nil {
		return nil, err
	}
	state, err := readProject(p.stateFile)
	if err != nil {
		return nil, err
	}
	return projectResult(p, enabled, false, state.RecommendedSuccesses), nil
}

func (tool app) record(repo string) (map[string]any, error) {
	p, err := tool.project(repo)
	if err != nil {
		return nil, err
	}
	enabled, err := tool.enabled()
	if err != nil || !enabled {
		return projectResult(p, false, false, 0), err
	}
	state, err := readProject(p.stateFile)
	if err != nil {
		return nil, err
	}
	if state.RecommendedSuccesses == threshold {
		return projectResult(p, true, false, state.RecommendedSuccesses), nil
	}
	state.RecommendedSuccesses++
	if err := writeJSON(p.stateFile, state); err != nil {
		return nil, err
	}
	return projectResult(p, true, true, state.RecommendedSuccesses), nil
}

func (tool app) forget(repo string) (map[string]any, error) {
	p, err := tool.project(repo)
	if err != nil {
		return nil, err
	}
	err = os.Remove(p.stateFile)
	forgotten := err == nil
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return map[string]any{"forgotten": forgotten, "project": p.name, "project_id": p.id, "state_file": p.stateFile}, nil
}

func (tool app) enabled() (bool, error) {
	var value settings
	exists, err := readJSON(filepath.Join(tool.stateDir, "settings.json"), &value)
	return exists && value.Enabled, err
}

func (tool app) project(repo string) (project, error) {
	root, err := git(repo, "rev-parse", "--show-toplevel")
	if err != nil {
		return project{}, err
	}
	root, err = filepath.Abs(root)
	if err != nil {
		return project{}, err
	}
	commonDir, err := git(root, "rev-parse", "--git-common-dir")
	if err != nil {
		return project{}, err
	}
	if !filepath.IsAbs(commonDir) {
		commonDir = filepath.Join(root, commonDir)
	}
	commonDir, err = filepath.Abs(commonDir)
	if err != nil {
		return project{}, err
	}
	sum := sha256.Sum256([]byte(commonDir))
	id := hex.EncodeToString(sum[:])
	return project{filepath.Base(root), id, filepath.Join(tool.stateDir, "repos", id+".json")}, nil
}

func projectResult(p project, enabled bool, recorded bool, count int) map[string]any {
	var mode any
	if count == threshold {
		mode = "direct"
	}
	return map[string]any{
		"enabled": enabled, "recorded": recorded, "project": p.name, "project_id": p.id,
		"state_file": p.stateFile, "recommended_successes": count, "learned_mode": mode, "threshold": threshold,
	}
}

func readProject(path string) (projectState, error) {
	var state projectState
	_, err := readJSON(path, &state)
	if err == nil && (state.RecommendedSuccesses < 0 || state.RecommendedSuccesses > threshold) {
		err = fmt.Errorf("invalid state in %s", path)
	}
	return state, err
}

func readJSON(path string, target any) (bool, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(data, target)
}

func writeJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// ponytail: concurrent writes may lose one preference update; add locking only if concurrent commits become real.
	if err := os.WriteFile(path, append(data, '\n'), 0o600); err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}

func git(repo string, args ...string) (string, error) {
	output, err := exec.Command("git", append([]string{"-C", repo}, args...)...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), strings.TrimSpace(string(output)))
	}
	return strings.TrimSpace(string(output)), nil
}
