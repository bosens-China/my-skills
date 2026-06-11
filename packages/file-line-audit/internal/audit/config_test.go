package audit

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseConfigRequiresInclude(t *testing.T) {
	t.Parallel()

	if _, err := ParseConfig(nil); err == nil {
		t.Fatal("expected include required error")
	}
}

func TestParseConfigReadsFlags(t *testing.T) {
	t.Parallel()

	config, err := ParseConfig([]string{
		"--threshold", "128",
		"--include", "src/**/*.go",
		"--include", "apps/**/*.{ts,tsx}",
		"--exclude", "dist/",
		"--exclude", "*.snap",
	})
	if err != nil {
		t.Fatalf("ParseConfig returned error: %v", err)
	}

	if config.Threshold != 128 {
		t.Fatalf("expected threshold 128, got %d", config.Threshold)
	}

	wantInclude := []string{"src/**/*.go", "apps/**/*.{ts,tsx}"}
	if !reflect.DeepEqual(config.Include, wantInclude) {
		t.Fatalf("unexpected include patterns: %#v", config.Include)
	}

	wantExclude := []string{"dist/", "*.snap"}
	if !reflect.DeepEqual(config.Exclude, wantExclude) {
		t.Fatalf("unexpected exclude patterns: %#v", config.Exclude)
	}
}

func TestParseConfigReadsJSON(t *testing.T) {
	t.Parallel()

	config, err := ParseConfig([]string{
		"--json", `{"threshold":256,"include":["internal/**/*.go"],"exclude":["vendor/"]}`,
	})
	if err != nil {
		t.Fatalf("ParseConfig returned error: %v", err)
	}

	if config.Threshold != 256 {
		t.Fatalf("expected threshold 256, got %d", config.Threshold)
	}

	wantInclude := []string{"internal/**/*.go"}
	if !reflect.DeepEqual(config.Include, wantInclude) {
		t.Fatalf("unexpected include patterns: %#v", config.Include)
	}

	wantExclude := []string{"vendor/"}
	if !reflect.DeepEqual(config.Exclude, wantExclude) {
		t.Fatalf("unexpected exclude patterns: %#v", config.Exclude)
	}
}

func TestParseConfigRejectsInvalidIncludePattern(t *testing.T) {
	t.Parallel()

	if _, err := ParseConfig([]string{
		"--include", "src/**[",
	}); err == nil {
		t.Fatal("expected invalid include pattern error")
	}
}

func TestParseConfigShowsHelp(t *testing.T) {
	t.Parallel()

	_, err := ParseConfig([]string{"--help"})
	if !errors.Is(err, errHelp) {
		t.Fatalf("expected help error, got %v", err)
	}
}
