# Arivon OS — Architecture

## Overview

Arivon OS is a Debian Stable-based, minimal server operating system.
The Arivon layer sits on top of Debian and provides: installer, CLI,
security defaults, server profiles, Docker integration, health tooling,
and build infrastructure.

## Layers

```
+------------------------------------------+
|              Arivon Layer                 |
|  arivon CLI  |  Installer  |  Configs     |
|  Security   |  Profiles   |  Build       |
+------------------------------------------+
|              Debian Layer                 |
|  APT  |  Kernel  |  systemd  |  OpenSSH    |
|  nftables  |  glibc  |  Packages             |
+------------------------------------------+
|              Hardware                     |
+------------------------------------------+
```

## Design Principles

- Minimal — install only what is required
- Secure by default — reduced attack surface
- Container-first — Docker is a first-class feature
- SSH-first — remote administration is the primary interface
- Modular — optional features do not increase base complexity
- Reproducible — buildable from source
- Upstream-friendly — no unnecessary forks

## Component Map

| Component     | Path              | Technology        |
|---------------|-------------------|-------------------|
| CLI           | cli/              | Go + cobra        |
| Installer     | installer/        | Go + bubbletea    |
| Build system  | build/            | Bash + mkosi      |
| Images        | images/           | mkosi configs     |
| Configs       | configs/          | TOML + drop-ins   |
| Services      | services/         | systemd units     |
| Security      | security/         | Audit scripts     |
| Profiles      | profiles/         | YAML              |
| Packaging     | packages/         | dpkg-deb          |
| Tests         | tests/            | Go test + bats    |
| CI            | .github/workflows | GitHub Actions   |

## Configuration

All Arivon configuration lives under `/etc/arivon/`:

```
/etc/arivon/config.toml    # main config
/etc/arivon/security.conf  # security settings
/etc/arivon/docker.conf    # Docker settings
/etc/arivon/network.conf   # network settings
/etc/arivon/backup.conf    # backup settings
```

SSH hardening uses `/etc/ssh/sshd_config.d/arivon.conf` (drop-in).
Firewall uses `/etc/nftables/arivon.nft`.

## Version Identity

Four version strings are kept separate:
- Arivon OS version (e.g. 0.1.0)
- Debian upstream version (e.g. 12.5)
- Kernel version (e.g. 6.1.0)
- Arivon package version (e.g. 0.1.0)
