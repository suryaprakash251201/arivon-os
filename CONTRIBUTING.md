# Contributing to Arivon OS

## Getting Started

1. Fork the repository
2. Clone your fork
3. Run `scripts/setup-dev.sh` to set up the development environment
4. Create a feature branch

## Development Workflow

1. Create a feature branch: `git checkout -b feature/my-feature`
2. Make your changes
3. Run tests: `make test`
4. Run linting: `make lint`
5. Commit with conventional commits: `feat:`, `fix:`, `docs:`, `chore:`
6. Open a pull request

## Coding Standards

### Go

- Standard Go formatting (`gofmt`)
- Meaningful variable and function names
- No unnecessary comments
- Tests for all new functionality

### Shell

- `set -euo pipefail` at the top
- ShellCheck clean
- Meaningful variable names
- Comments only where they provide technical context

### Configuration

- TOML for Arivon config files
- YAML for profiles
- systemd drop-ins for service configuration

## Testing

```bash
# Unit tests
cd cli && go test ./... -v

# Shell tests
bats tests/

# All tests
make test
```

## Commit Messages

Use conventional commits:

- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation
- `test:` tests
- `chore:` maintenance
- `refactor:` code refactoring

## Pull Requests

- Small, focused commits
- Clear description of changes
- Tests pass
- Documentation updated
