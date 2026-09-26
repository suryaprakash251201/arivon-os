# Arivon OS

> **Arivon OS — Intelligent Infrastructure.**

Lightweight, security-focused, terminal-first Linux distribution
optimized for servers, SSH administration, Docker containers,
networking, storage, self-hosting, and home labs.

## Status

**Phase 8 — Release Engineering** (complete)

- [x] Repository structure
- [x] Go CLI (18 commands)
- [x] Build system (mkosi)
- [x] Debian base configuration
- [x] systemd service units
- [x] Security defaults (sshd, nftables)
- [x] Installation profiles
- [x] CI/CD pipeline
- [x] Test framework (14 tests)
- [x] Documentation
- [x] Release engineering
- [ ] Bootable ISO (requires Linux host)
- [ ] VM testing (requires QEMU)

## Quick Start

```bash
# Build
make build

# Test
make test

# Build ISO (Linux host with mkosi)
make iso

# Build release
make release VERSION=0.1.0
```

## Architecture

See [docs/architecture.md](docs/architecture.md).

## CLI Commands

```bash
arivon info           # OS identity
arivon status         # System status
arivon health         # Health check
arivon doctor         # Diagnose problems
arivon update         # Manage updates
arivon logs           # View logs
arivon services       # Manage services
arivon security       # Security audit
arivon firewall       # Manage firewall
arivon ssh            # Manage SSH
arivon docker         # Manage Docker
arivon network        # Network management
arivon storage        # Storage management
arivon backup         # Backup management
arivon tailscale      # Tailscale (optional)
arivon benchmark      # System benchmark
arivon monitoring     # Monitoring
```

## Documentation

- [Installation](docs/installation.md)
- [Getting Started](docs/getting-started.md)
- [Building](docs/building.md)
- [Development](docs/development.md)
- [Security](docs/security.md)
- [Docker](docs/docker.md)
- [Networking](docs/networking.md)
- [Storage](docs/storage.md)
- [Backups](docs/backups.md)
- [Monitoring](docs/monitoring.md)
- [Troubleshooting](docs/troubleshooting.md)
- [Upgrades](docs/upgrades.md)
- [Recovery](docs/recovery.md)

## License

MIT
