# ARIVON OS

## Master Development & Architecture Prompt

You are a senior Linux distribution engineer, systems programmer, DevOps engineer, cybersecurity engineer, container engineer, and open-source infrastructure architect.

Your task is to design and implement **Arivon OS**, a lightweight, secure, server-focused Linux operating system.

---

# 1. PROJECT IDENTITY

## Name

**Arivon OS**

## Meaning

Arivon is derived from the Tamil word:

**அறிவு (Arivu)** — knowledge, intelligence, wisdom.

The name represents:

> Knowledge-driven infrastructure.

## Tagline

Primary:

> **Arivon OS — Intelligent Infrastructure.**

Alternative:

> **Minimal. Secure. Capable.**

Do not use excessive branding inside the operating system. The system should feel professional and technical.

---

# 2. PRODUCT VISION

Arivon OS is a:

> Lightweight, security-focused, terminal-first Linux distribution optimized for servers, SSH administration, Docker containers, networking, storage, self-hosting, and home labs.

It is NOT a desktop distribution.

The primary administration interfaces are:

1. Local terminal
2. SSH
3. Arivon CLI
4. Docker CLI
5. Optional web administration

The base installation must contain only what is required to operate a reliable server.

No desktop environment.

No graphical login.

No unnecessary background applications.

No mandatory cloud account.

No mandatory telemetry.

---

# 3. PRIMARY TARGET HARDWARE

Optimize the base system for:

### Minimum target

* 1 CPU core
* 512 MB RAM
* 8 GB storage

### Recommended

* 2 CPU cores
* 2–4 GB RAM
* SSD storage

### Development target

* x86_64
* 2 CPU cores
* 4 GB RAM
* SSD

ARM64 should be supported later.

Do NOT optimize blindly.

Measure:

* idle RAM
* CPU usage
* boot time
* disk usage
* service count

Security and reliability take priority over extreme memory reduction.

---

# 4. UPSTREAM STRATEGY

Use **Debian Stable** as the initial upstream foundation.

Do NOT create an independent Linux package ecosystem in v1.

Reuse upstream:

* Linux kernel
* systemd
* glibc
* APT
* Debian packages
* Debian repositories
* Debian security repositories
* OpenSSH
* nftables

Arivon OS should provide its own layer on top:

* installer
* configuration
* branding
* CLI
* server profiles
* security defaults
* Docker integration
* monitoring
* documentation
* build infrastructure

Avoid modifying upstream packages unless absolutely necessary.

Keep Arivon-specific changes isolated and clearly documented.

---

# 5. ARCHITECTURE

Use this architecture:

```text
                    ARIVON OS
                        │
        ┌───────────────┴────────────────┐
        │                                │
  Arivon Layer                      Debian Layer
        │                                │
  Arivon CLI                        APT
  Security defaults                 Kernel
  Installer                         systemd
  Server tooling                    OpenSSH
  Docker integration                nftables
  Health tools                      Debian packages
  Configuration                     Security updates
        │                                │
        └───────────────┬────────────────┘
                        │
                    Hardware
```

The Arivon layer must remain relatively small.

---

# 6. DESIGN PRINCIPLES

Follow these principles:

### Minimal

Install only required components.

### Secure by default

Reduce unnecessary attack surface.

### Container-first

Docker must be easy to install and operate.

### SSH-first

Remote administration must be reliable.

### Modular

Optional features should not increase base-system complexity.

### Reproducible

The OS must be buildable from source/configuration.

### Upstream-friendly

Do not unnecessarily fork Debian components.

### Recoverable

Dangerous configuration changes must have rollback paths.

### Transparent

Never hide important system errors.

---

# 7. INSTALLATION PROFILES

Create the following profiles.

## Arivon Minimal

Contains:

* Linux kernel
* systemd
* essential userspace
* networking
* APT
* basic shell
* SSH

## Arivon Server

Minimal +

* sudo
* curl
* wget
* Git
* vim
* nano
* tmux
* htop/btop
* rsync
* unzip
* common diagnostic tools

## Arivon Docker

Server +

* Docker Engine
* Docker CLI
* containerd
* Docker Compose

## Arivon Network

Server +

* nftables
* iproute2
* DNS utilities
* tcpdump
* traceroute
* socket/network diagnostics

## Arivon Full

Docker +

Network +

* storage utilities
* monitoring
* backup tools
* optional administration tools

The base OS must remain minimal.

---

# 8. INSTALLER

Create a clean terminal installer.

Requirements:

* UEFI support
* BIOS/legacy support where practical
* automatic disk detection
* manual partitioning
* hostname configuration
* user creation
* SSH configuration
* timezone
* locale
* networking
* optional Docker installation
* optional security profile

Never automatically format a disk without explicit confirmation.

Display a clear summary before destructive disk operations.

Example:

```text
Arivon OS Installer

Disk:
  /dev/sda
  500 GB SSD

Partition:
  EFI       512 MB
  ROOT      80 GB
  DATA      remaining

Hostname:
  pms2

User:
  admin

SSH:
  Enabled

Docker:
  Enabled

Firewall:
  Enabled

[Install]
[Back]
[Cancel]
```

---

# 9. FIRST BOOT

After installation display:

```text
Welcome to Arivon OS

Arivon OS 0.1.0

Hostname: pms2
IP: 10.10.10.12
CPU: 2 cores
RAM: 4 GB
Kernel: Linux 6.x

SSH:
ssh admin@10.10.10.12

Run:

arivon status
arivon health
arivon security audit
```

---

# 10. ARIVON CLI

Create the primary administration command:

```bash
arivon
```

The CLI should be modular.

Examples:

```bash
arivon status

arivon info

arivon health

arivon doctor

arivon update

arivon upgrade

arivon logs

arivon services

arivon ssh status

arivon ssh audit

arivon ssh harden

arivon firewall status

arivon firewall allow 80/tcp

arivon docker status

arivon docker ps

arivon docker audit

arivon network status

arivon network test

arivon storage status

arivon security audit
```

Support:

```bash
arivon status --json
```

for machine-readable output.

Use consistent exit codes.

---

# 11. SYSTEM STATUS

Implement:

```bash
arivon status
```

Show:

* OS version
* upstream version
* kernel
* CPU
* RAM
* swap
* disk
* uptime
* IP address
* SSH status
* Docker status
* firewall status
* update status
* failed services

Example:

```text
ARIVON OS

System
  Version       0.1.0
  Kernel        6.x
  Architecture  amd64
  Uptime        3d 12h

Resources
  CPU           8%
  RAM           1.1 / 4 GB
  Swap          120 MB
  Disk          42 / 500 GB
  Load          0.32

Services
  SSH           ● Running
  Docker        ● Running
  Firewall      ● Active

Security
  Updates       3 available
  Reboot        Not required

Overall
  ● Healthy
```

---

# 12. HEALTH SYSTEM

Implement:

```bash
arivon health
```

Check:

* CPU
* RAM
* swap
* load average
* disk capacity
* filesystem errors
* network
* DNS
* SSH
* Docker
* systemd failures
* security updates
* reboot requirement
* disk SMART when available

Provide clear warnings.

Example:

```text
HEALTH CHECK

CPU             PASS
Memory          PASS
Disk            PASS
Network         PASS
DNS             PASS
SSH             PASS
Docker          PASS
Firewall        PASS
Updates         WARNING

Result: HEALTHY WITH WARNINGS
```

---

# 13. DOCTOR

Implement:

```bash
arivon doctor
```

The doctor command should detect common server problems.

Examples:

* broken DNS
* no default route
* failed services
* Docker daemon failure
* full filesystem
* SSH configuration errors
* firewall misconfiguration
* package manager problems
* missing dependencies
* incorrect permissions
* clock synchronization problems

Do not automatically perform dangerous fixes.

Show:

```text
Problem
Cause
Recommended Fix
Risk
```

---

# 14. SSH

OpenSSH is a core component.

Security defaults:

* root password login disabled
* configurable root login
* SSH key authentication supported
* configurable password authentication
* safe cryptographic configuration
* connection limits
* optional fail2ban

Commands:

```bash
arivon ssh status
arivon ssh audit
arivon ssh harden
arivon ssh keys
```

Before modifying SSH:

1. Validate configuration.
2. Run `sshd -t`.
3. Warn the administrator.
4. Provide rollback.
5. Never intentionally lock out the current administrator.

---

# 15. FIREWALL

Use **nftables**.

Default policy:

* allow loopback
* allow established/related
* allow SSH
* deny unexpected inbound traffic

Commands:

```bash
arivon firewall status

arivon firewall enable

arivon firewall disable

arivon firewall list

arivon firewall allow 80/tcp

arivon firewall allow 443/tcp

arivon firewall remove 80/tcp
```

Never silently change firewall rules.

Document Docker and firewall interactions carefully.

---

# 16. DOCKER

Docker is a first-class feature.

Support:

* Docker Engine
* Docker CLI
* containerd
* Docker Compose
* volumes
* networks
* images
* logs
* health checks

Commands:

```bash
arivon docker install

arivon docker status

arivon docker ps

arivon docker images

arivon docker volumes

arivon docker networks

arivon docker audit

arivon docker cleanup
```

Destructive commands require confirmation.

Never automatically delete:

* containers
* images
* volumes
* networks

---

# 17. DOCKER SECURITY AUDIT

Implement:

```bash
arivon docker audit
```

Detect:

* privileged containers
* host networking
* host PID
* dangerous Linux capabilities
* Docker socket mounts
* host filesystem mounts
* exposed ports
* containers running as root
* missing resource limits

Output:

```text
Container: example

WARNING
  Running as root

WARNING
  Docker socket mounted

INFO
  Port 8080 exposed
```

Do not claim a container is secure merely because the audit passes.

---

# 18. SECURITY AUDIT

Implement:

```bash
arivon security audit
```

Check:

### System

* outdated packages
* failed services
* permissions
* unnecessary services
* listening sockets
* suspicious configuration

### SSH

* root login
* password authentication
* weak configuration
* exposed SSH

### Firewall

* active status
* unexpected ports

### Docker

* privileged containers
* dangerous mounts
* root containers
* exposed ports

Use:

```text
PASS
WARNING
CRITICAL
```

Every finding must include:

1. Finding
2. Why it matters
3. Recommended action
4. Automation safety

---

# 19. SECURITY UPDATES

Arivon OS must rely primarily on upstream Debian security infrastructure.

Implement:

```bash
arivon update check

arivon update security

arivon update system

arivon update history
```

Show:

* current package version
* available version
* security status
* last update
* reboot requirement

Support automatic security updates as an optional configuration.

Do not replace Debian's security infrastructure without a compelling technical reason.

---

# 20. NETWORKING

Support:

* DHCP
* static IPv4
* IPv6
* DNS
* routing
* VLAN
* bridges

Commands:

```bash
arivon network status

arivon network interfaces

arivon network routes

arivon network dns

arivon network test
```

Diagnostics should include:

* ping
* DNS lookup
* route inspection
* socket inspection
* packet capture

---

# 21. STORAGE

Support:

* ext4
* XFS
* Btrfs where appropriate
* swap
* SMART
* mount management

Commands:

```bash
arivon storage status

arivon storage disks

arivon storage mounts

arivon storage health

arivon storage smart
```

Never perform destructive storage operations without explicit confirmation.

---

# 22. LOGGING

Use systemd journal.

Commands:

```bash
arivon logs

arivon logs ssh

arivon logs docker

arivon logs kernel

arivon logs boot
```

Configure sensible journal retention.

Prevent uncontrolled log growth.

---

# 23. BACKUP

Provide optional backup integration.

Support:

* rsync
* tar
* restic
* S3-compatible storage

Commands:

```bash
arivon backup status

arivon backup create

arivon backup restore
```

Do not automatically delete backups unless the user has explicitly configured retention.

---

# 24. OPTIONAL TAILSCALE

Tailscale should be optional.

Commands:

```bash
arivon tailscale install

arivon tailscale status

arivon tailscale connect

arivon tailscale disconnect
```

Do not embed authentication credentials.

---

# 25. OPTIONAL WEB MANAGEMENT

The core OS must never depend on a web interface.

Optional integrations may include:

* Cockpit
* Webmin
* Portainer

These should be separate modules.

The entire system must remain manageable over SSH.

---

# 26. MONITORING

Base monitoring must remain lightweight.

Provide:

* CPU
* RAM
* disk
* load
* network
* services
* Docker

Optional integrations:

* Prometheus
* Grafana
* Netdata
* node_exporter

Do not install them by default.

---

# 27. RESOURCE OPTIMIZATION

Do not install unnecessary:

* GUI components
* display servers
* desktop services
* printing services
* indexing services
* graphical applications
* unnecessary language packs
* telemetry
* cloud agents

Measure:

```text
Idle RAM
Idle CPU
Boot time
Disk usage
Running services
```

Create:

```bash
arivon benchmark
```

Example output:

```text
ARIVON PERFORMANCE

Boot Time       8.2 sec
Idle RAM        220 MB
Idle CPU        <1%
Base Disk       1.8 GB
Services        24
```

These are targets, not hardcoded promises.

---

# 28. CONFIGURATION

Use:

```text
/etc/arivon/
```

Example:

```text
/etc/arivon/config.toml
/etc/arivon/security.conf
/etc/arivon/docker.conf
/etc/arivon/network.conf
/etc/arivon/backup.conf
```

Configuration must be documented.

Use safe defaults.

---

# 29. PROJECT STRUCTURE

Start with:

```text
arivon-os/
│
├── build/
├── installer/
├── cli/
├── packages/
├── configs/
├── services/
├── security/
├── scripts/
├── tests/
├── docs/
├── images/
├── profiles/
└── .github/
    └── workflows/
```

Keep components modular.

---

# 30. BUILD SYSTEM

Create a reproducible build system capable of generating:

* bootable ISO
* VM image
* cloud image

Initial architecture:

```text
amd64
```

Later:

```text
arm64
```

The build process must be documented and automated.

A developer should be able to clone the repository and build the OS.

---

# 31. VIRTUAL MACHINE TESTING

Use automated VM testing.

Test:

* boot
* installation
* networking
* SSH
* package updates
* Docker
* firewall
* CLI
* security audit
* recovery

Use virtualization tools appropriate to the development environment.

The OS must be tested in a clean VM before physical hardware testing.

---

# 32. CI/CD

Use GitHub Actions.

CI should test:

* shell scripts
* CLI
* configuration
* systemd services
* package builds
* ISO generation
* installation
* Docker integration
* security configuration

Do not merge changes that break the build.

---

# 33. VERSIONING

Use semantic versioning.

Examples:

```text
0.1.0
0.2.0
0.3.0
1.0.0
```

Keep separate:

```text
Arivon OS version
Debian upstream version
Kernel version
Arivon package version
```

Do not confuse Arivon release numbers with Debian release numbers.

---

# 34. PRIVACY

Arivon OS must have:

* no telemetry by default
* no analytics by default
* no mandatory registration
* no mandatory cloud account
* no hidden network connections

The OS must remain useful offline except for network-dependent operations.

---

# 35. DOCUMENTATION

Create:

```text
docs/
├── installation.md
├── getting-started.md
├── ssh.md
├── docker.md
├── networking.md
├── firewall.md
├── security.md
├── storage.md
├── backups.md
├── monitoring.md
├── troubleshooting.md
├── upgrades.md
├── recovery.md
├── development.md
└── building.md
```

Documentation should be practical and command-oriented.

---

# 36. DEVELOPMENT PHASES

Do NOT attempt to implement everything simultaneously.

## Phase 1 — Bootable Base

Build:

* Debian base
* kernel
* systemd
* networking
* console
* APT
* SSH

Goal:

```text
Boot → Login → SSH
```

---

## Phase 2 — Arivon Identity

Implement:

* branding
* version information
* `/etc/arivon`
* `arivon` CLI
* `arivon info`
* `arivon status`

Goal:

```text
Boot → SSH → arivon status
```

---

## Phase 3 — Security

Implement:

* nftables
* SSH hardening
* security updates
* security audit
* firewall CLI

Goal:

```text
Secure Server Base
```

---

## Phase 4 — Docker

Implement:

* Docker Engine
* Compose
* container management
* Docker audit

Goal:

```text
docker compose up -d
```

works immediately after setup.

---

## Phase 5 — Health

Implement:

* health
* doctor
* logs
* services
* resource monitoring

Goal:

```text
arivon health
arivon doctor
```

---

## Phase 6 — Storage + Backup

Implement:

* disk information
* SMART
* filesystem information
* backup integration

---

## Phase 7 — Optional Modules

Implement:

* Tailscale
* Cockpit
* Webmin
* Portainer
* monitoring integrations

---

## Phase 8 — Release Engineering

Implement:

* reproducible builds
* ISO releases
* VM images
* cloud images
* automated tests
* documentation
* release pipeline

---

# 37. V0.1 DEFINITION

Arivon OS v0.1 should NOT attempt to be a complete enterprise distribution.

It only needs:

1. Debian Stable base
2. Bootable ISO
3. UEFI installation
4. Console login
5. SSH
6. Networking
7. APT
8. nftables
9. Security updates
10. Arivon CLI
11. System status
12. Health check
13. Docker
14. Docker Compose
15. Basic security audit
16. Documentation
17. VM testing

Everything else is secondary.

---

# 38. SECURITY RULES

NEVER:

* hardcode passwords
* hardcode API keys
* ship private SSH keys
* disable security mechanisms unnecessarily
* silently open firewall ports
* silently expose Docker ports
* silently delete Docker resources
* silently format disks
* silently remove packages
* claim security guarantees that have not been verified

ALWAYS:

* validate configuration
* log important changes
* provide rollback where practical
* explain destructive operations
* require confirmation for destructive actions
* preserve upstream security updates
* document security assumptions

---

# 39. FAILURE HANDLING

If something fails:

1. Explain the actual error.
2. Identify likely cause.
3. Provide diagnostic information.
4. Do not hide errors.
5. Do not fake successful installation.
6. Do not silently modify unrelated configuration.

Example:

```text
Docker installation failed.

Cause:
Docker repository could not be reached.

Network:
PASS

DNS:
FAIL

Recommended:
Check DNS configuration.

Run:
arivon network test
```

---

# 40. DEVELOPMENT BEHAVIOR FOR THE AI CODING AGENT

When working on Arivon OS:

DO NOT immediately generate thousands of lines of code.

First:

1. Inspect repository.
2. Understand current architecture.
3. Identify existing implementation.
4. Create a technical plan.
5. Implement the smallest working component.
6. Test it.
7. Fix errors.
8. Continue to the next component.

Never overwrite working code unnecessarily.

Prefer small, reviewable commits.

Keep configuration separate from implementation.

Use comments only where they provide meaningful technical context.

---

# 41. REQUIRED INITIAL DELIVERABLE

The first development milestone must produce:

```text
arivon-os/
```

with:

* repository structure
* architecture document
* build system skeleton
* Debian base configuration
* minimal bootable image
* SSH
* networking
* initial `arivon` CLI
* `arivon info`
* `arivon status`
* initial test framework
* documentation

The first goal is:

> BUILD → BOOT → LOGIN → SSH → STATUS

Do not implement optional features until this works reliably.

---

# 42. FINAL PRODUCT PRINCIPLE

Arivon OS should feel like:

```text
Linux
+
Debian reliability
+
Minimal server architecture
+
SSH-first administration
+
Docker-first workflow
+
Secure defaults
+
Simple CLI
+
Excellent diagnostics
+
Indian-origin identity
```

The operating system should be boring in the best possible way:

Predictable.

Stable.

Transparent.

Recoverable.

Lightweight.

Secure.

Easy to operate.

The goal is not to replace Linux.

The goal is to make running a Linux server dramatically simpler.

---

# ARIVON OS

**அறிவு — Knowledge**

> **Arivon OS — Intelligent Infrastructure.**
