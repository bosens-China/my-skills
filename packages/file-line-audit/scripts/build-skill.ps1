Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$packageRoot = Split-Path -Parent $PSScriptRoot
$repoRoot = Split-Path -Parent (Split-Path -Parent $packageRoot)
$outputDir = Join-Path $repoRoot "skills\file-line-audit\scripts"

New-Item -ItemType Directory -Force -Path $outputDir | Out-Null

$targets = @(
  @{ GOOS = "windows"; GOARCH = "amd64"; Output = "line-audit-windows-amd64.exe" },
  @{ GOOS = "linux"; GOARCH = "amd64"; Output = "line-audit-linux-amd64" },
  @{ GOOS = "linux"; GOARCH = "arm64"; Output = "line-audit-linux-arm64" },
  @{ GOOS = "darwin"; GOARCH = "amd64"; Output = "line-audit-darwin-amd64" },
  @{ GOOS = "darwin"; GOARCH = "arm64"; Output = "line-audit-darwin-arm64" }
)

foreach ($target in $targets) {
  $env:GOOS = $target.GOOS
  $env:GOARCH = $target.GOARCH
  $outputPath = Join-Path $outputDir $target.Output
  Push-Location $packageRoot
  & go build -o $outputPath ./cmd/line-audit
  Pop-Location
}

Remove-Item Env:GOOS -ErrorAction SilentlyContinue
Remove-Item Env:GOARCH -ErrorAction SilentlyContinue
