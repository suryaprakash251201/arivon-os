# Arivon OS

> **Arivon OS — Intelligent Infrastructure.**
> அறிவு (Arivu) — knowledge, intelligence, wisdom.

[![CI](https://github.com/suryaprakash251201/arivon-os/actions/workflows/ci.yml/badge.svg)](https://github.com/suryaprakash251201/arivon-os/actions/workflows/ci.yml)
[![ISO Build](https://github.com/suryaprakash251201/arivon-os/actions/workflows/iso.yml/badge.svg)](https://github.com/suryaprakash251201/arivon-os/actions/workflows/iso.yml)
[![Release](https://github.com/suryaprakash251201/arivon-os/actions/workflows/release.yml/badge.svg)](https://github.com/suryaprakash251201/arivon-os/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Debian](https://img.shields.io/badge/Upstream-Debian%20trixie-red.svg)](https://www.debian.org/)
[![Go](https://img.shields.io/badge/CLI-Go%201.26-blue.svg)](https://go.dev/)

A lightweight, security-focused, **terminal-first** Linux distribution
based on **Debian Stable** — built for servers, SSH administration,
Docker containers, networking, storage, self-hosting, and home labs.

No desktop. No telemetry. No mandatory accounts. Just a boring,
predictable server that is easy to operate.

---

## Why Arivon OS?

| Plain Debian server | Arivon OS |
|---|---|
| Hand-tuned SSH, firewall, journal | Secure defaults out of the box |
| `systemctl` + `journalctl` + `nft` + … | One CLI: `arivon status`, `arivon health`, `arivon doctor` |
| Manual Docker repo setup | `arivon docker install` |
| Ad-hoc hardening guides | `arivon security audit`, `arivon ssh harden` |
| DIY image builds | Reproducible mkosi images + CI releases |

Upstream packages are never forked — Arivon is a thin,
well-isolated layer on top of Debian.

## Features

- **Minimal** — boots on 1 CPU / 512 MB RAM; profiles from minimal to full
- **Secure by default** — nftables deny-inbound, key-only SSH, daily audits
- **Container-first** — Docker Engine + Compose as first-class citizens
- **SSH-first** — everything manageable over a single SSH session
- **Observable** — health checks, doctor diagnostics, JSON output for scripts
- **Reproducible** — `git clone` + `make iso` = bootable image

## Quick start

```bash
# Build the CLI locally
make build && ./build/arivon status

# Run tests
make test

# Build a bootable disk image (Linux host + mkosi)
make iso

# Flash it (bootable GPT image, dd-able like an ISO)
sudo dd if=images/output/arivon-os-minimal.raw of=/dev/sdX bs=4M status=progress conv=fsync
```

Or grab a prebuilt image from
[Releases](https://github.com/suryaprakash251201/arivon-os/releases)
— see the [[Installation]] wiki page.

## The `arivon` CLI

```bash
arivon status            # system overview (add --json for scripts)
arivon health            # PASS/WARNING per subsystem
arivon doctor            # problem → cause → fix → risk
arivon security audit    # SSH, firewall, updates, services, sockets
arivon ssh harden        # key-only SSH with validation + rollback
arivon firewall allow 443/tcp
arivon docker audit      # privileged / socket mounts / root containers
arivon backup create     # restic backup (S3-compatible)
```

Full reference: [docs/cli-reference.md](docs/cli-reference.md) ·
[Wiki CLI Reference](https://github.com/suryaprakash251201/arivon-os/wiki/CLI-Reference).

## Documentation

| Guide | Path |
|---|---|
| Installation | [docs/installation.md](docs/installation.md) |
| Getting started | [docs/getting-started.md](docs/getting-started.md) |
| Architecture | [docs/architecture.md](docs/architecture.md) |
| Building from source | [docs/building.md](docs/building.md) |
| Development | [docs/development.md](docs/development.md) |
| Security | [docs/security.md](docs/security.md) |
| Docker | [docs/docker.md](docs/docker.md) |
| Networking | [docs/networking.md](docs/networking.md) |
| Storage | [docs/storage.md](docs/storage.md) |
| Backups | [docs/backups.md](docs/backups.md) |
| Monitoring | [docs/monitoring.md](docs/monitoring.md) |
| Troubleshooting | [docs/troubleshooting.md](docs/troubleshooting.md) |
| Upgrades | [docs/upgrades.md](docs/upgrades.md) |
| Recovery | [docs/recovery.md](docs/recovery.md) |

Plus the [GitHub Wiki](https://github.com/suryaprakash251201/arivon-os/wiki)
(Home, FAQ, narrative guides).

## Roadmap

See [ROADMAP.md](ROADMAP.md): v0.2 real installer → v0.3 config-driven
system → v0.4 quality pass → v0.5 hardware validation → v1.0.

## Repository layout

```text
cli/          arivon CLI (Go + cobra)
installer/    TUI installer (Go + bubbletea)
images/       mkosi.conf + mkosi.profiles/ (minimal…full, cloud)
configs/      sshd / nftables / journald / network drop-ins
services/     systemd units + timers (health, audit, first boot)
profiles/     install-profile manifests
packages/     .deb build script
tests/        bats + pytest integration tests
docs/         user documentation
```

## Contributing

Small focused PRs, conventional commits, `make lint && make test`
green. Start with [CONTRIBUTING.md](CONTRIBUTING.md) and the
[Development wiki page](https://github.com/suryaprakash251201/arivon-os/wiki/Development).

## License

[MIT](LICENSE) — free for personal, homelab, and commercial use.
