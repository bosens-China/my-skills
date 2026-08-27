package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestLifecycle(t *testing.T) {
	repo := filepath.Join(t.TempDir(), "repo")
	if err := os.Mkdir(repo, 0o755); err != nil {
		t.Fatal(err)
	}
	if output, err := exec.Command("git", "-C", repo, "init", "--quiet").CombinedOutput(); err != nil {
		t.Fatalf("git init: %v\n%s", err, output)
	}
	stateDir := filepath.Join(t.TempDir(), "state")

	status := mustRun(t, []string{"status", "--repo", repo}, stateDir)
	if status["enabled"] != false || status["learned_mode"] != nil {
		t.Fatalf("unexpected initial status: %v", status)
	}
	mustRun(t, []string{"enable"}, stateDir)
	for count := 1; count <= threshold; count++ {
		result := mustRun(t, []string{"record-direct", "--repo", repo}, stateDir)
		if result["recommended_successes"] != count {
			t.Fatalf("record %d: %v", count, result)
		}
	}
	status = mustRun(t, []string{"status", "--repo", repo}, stateDir)
	if status["learned_mode"] != "direct" {
		t.Fatalf("expected direct: %v", status)
	}
	result := mustRun(t, []string{"record-direct", "--repo", repo}, stateDir)
	if result["recorded"] != false || result["recommended_successes"] != threshold {
		t.Fatalf("record after learning should be a no-op: %v", result)
	}
	mustRun(t, []string{"disable"}, stateDir)
	status = mustRun(t, []string{"status", "--repo", repo}, stateDir)
	if status["enabled"] != false || status["learned_mode"] != "direct" {
		t.Fatalf("disable should preserve memory: %v", status)
	}
	mustRun(t, []string{"forget", "--repo", repo}, stateDir)
	status = mustRun(t, []string{"status", "--repo", repo}, stateDir)
	if status["recommended_successes"] != 0 || status["learned_mode"] != nil {
		t.Fatalf("forget should clear memory: %v", status)
	}
}

func mustRun(t *testing.T, args []string, stateDir string) map[string]any {
	t.Helper()
	result, err := run(args, stateDir)
	if err != nil {
		t.Fatal(err)
	}
	return result.(map[string]any)
}
