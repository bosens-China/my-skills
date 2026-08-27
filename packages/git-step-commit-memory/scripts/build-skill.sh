#!/usr/bin/env bash
set -euo pipefail

package_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
repo_root="$(cd "$package_root/../.." && pwd)"
output_dir="$repo_root/skills/git-step-commit/scripts"

mkdir -p "$output_dir"

build_target() {
  local goos="$1"
  local goarch="$2"
  local output="$3"
  CGO_ENABLED=0 GOOS="$goos" GOARCH="$goarch" go build -trimpath -buildvcs=false -o "$output_dir/$output" ./cmd/commit-memory
}

cd "$package_root"

build_target windows amd64 commit-memory-windows-amd64.exe
build_target linux amd64 commit-memory-linux-amd64
build_target linux arm64 commit-memory-linux-arm64
build_target darwin amd64 commit-memory-darwin-amd64
build_target darwin arm64 commit-memory-darwin-arm64
