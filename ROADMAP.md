# Arivon OS Roadmap

> **Goal:** make running a Linux server dramatically simpler —
> predictable, stable, transparent, recoverable, lightweight, secure.

## Shipped

### v0.1 — Bootable base ✅

- Debian trixie base, systemd-boot, GPT disk images via mkosi
- Profiles: minimal, server, docker, network, full, cloud
- SSH, networking, APT out of the box
- `arivon` CLI: 19 commands (status, health, doctor, update, logs,
  services, security, firewall, ssh, docker, network, storage,
  backup, tailscale, benchmark, monitoring, …)
- TUI installer skeleton, `.deb` packaging
- CI (lint, tests, Gosec, amd64+arm64), ISO workflow, release pipeline
- Docs, wiki, 14 unit tests

### Security baseline ✅

- nftables deny-by-default, SSH hardening drop-in, `Banner`
- `arivon security audit` + daily timer, Docker audit
- No telemetry, no accounts, no hidden connections

## Next

### v0.2 — Installer that installs

- [ ] Wire `runInstall()` into the TUI flow
- [ ] Editable text inputs (hostname, user, SSH port, timezone, locale)
- [ ] Manual partitioning + DATA partition + swap
- [ ] Destructive-disk confirmation gate (plan §8 requirement)
- [ ] NVMe device naming (`/dev/nvme0n1p1`) + empty-disk guard
- [ ] Apply user/hostname/SSH keys, `sshd -t` validation, rollback

### v0.3 — Config-driven system

- [ ] `config.toml` actually drives SSH / firewall / DNS
- [ ] Add `security.conf`, `docker.conf`, `network.conf`, `backup.conf`
- [ ] Document every key; safe defaults; atomic writes + backups
- [ ] `arivon ssh harden` renders from config (incl. `Banner` line)

### v0.4 — Quality and consistency

- [ ] `--json` on health, doctor, security audit, monitoring
- [ ] `RunE` error handling pass; no silent `apt`/`docker` failures
- [ ] Fix disk-full thresholds, `update history`, uname arch mapping
- [ ] bats coverage per command; `Makefile` test/lint path fixes
- [ ] `docs/ssh.md`, `docs/firewall.md`

### v0.5 — Hardware and cloud validation

- [ ] First bare-metal install + boot validation
- [ ] Boot the built image in QEMU in CI (smoke test)
- [ ] cloud-init image test on a cloud VM
- [ ] ARM64 image build (`GOARCH` works; images pending)

## Later (v1.0+)

- ARM64 images, Secure Boot (signed UKI + auto-enrollment)
- Prometheus/Grafana/node_exporter modules
- Cockpit / Portainer optional modules
- BIOS/legacy boot path (grub) alongside systemd-boot
- A/B updates with rollback, automated installer tests
- True ISO9660 output (today: bootable `.raw` disk images)

## How to track

- This file (version-controlled roadmap)
- Wiki [[Roadmap]] (narrative version)
- GitHub Project board (task-level tracking)
- `CHANGELOG.md` (what actually shipped)
