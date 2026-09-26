# Arivon OS — Development

## Repository Structure

```
arivon-os/
├── build/           # Build scripts
├── installer/       # TUI installer (Go)
├── cli/             # arivon CLI (Go)
├── packages/        # .deb packaging
├── configs/         # Default configs
├── services/        # systemd units
├── security/        # Security defaults
├── scripts/         # Helper scripts
├── tests/           # Test framework
├── docs/            # Documentation
├── images/          # mkosi image configs
├── profiles/        # Installation profiles
└── .github/         # CI workflows
```

## Development Workflow

1. Create a feature branch
2. Implement the smallest working component
3. Test it
4. Open a PR

## Coding Standards

- Go: standard formatting, meaningful names, no unnecessary comments
- Bash: `shellcheck` clean, `set -euo pipefail`
- Config: TOML with safe defaults
- Commits: conventional commits (`feat:`, `fix:`, `docs:`)

## Testing

```bash
# Lint
golangci-lint run
shellcheck build/*.sh scripts/*.sh

# Unit tests
go test ./...

# Build verification
./build/build.sh
```
