# Arivon OS — CLI Reference

## Global Flags

All commands support `--json` for machine-readable output where applicable.

## Commands

### arivon info

Show Arivon OS identity information.

```bash
arivon info
```

Output:
- Version
- Upstream (Debian version)
- Kernel
- Architecture

### arivon status

Show system status.

```bash
arivon status
arivon status --json
```

Output:
- System info (version, kernel, arch, uptime)
- Resources (CPU, RAM, swap, disk, load)
- Services (SSH, Docker, firewall)
- Security (updates, reboot)
- Overall health

### arivon health

Run system health check.

```bash
arivon health
```

Checks: CPU, memory, disk, network, DNS, SSH, Docker, firewall, updates.

### arivon doctor

Diagnose common server problems.

```bash
arivon doctor
```

Detects: broken DNS, no route, failed services, Docker failure,
full filesystem, SSH errors, firewall issues, clock sync problems.

### arivon update

Manage system updates.

```bash
arivon update check
arivon update security
arivon update system
arivon update history
```

### arivon logs

View system logs.

```bash
arivon logs
arivon logs ssh
arivon logs docker
arivon logs kernel
arivon logs boot
```

### arivon services

Manage system services.

```bash
arivon services list
arivon services status <service>
arivon services failed
```

### arivon security

Security management.

```bash
arivon security audit
```

### arivon firewall

Manage nftables firewall.

```bash
arivon firewall status
arivon firewall enable
arivon firewall disable
arivon firewall list
arivon firewall allow 80/tcp
arivon firewall remove 80/tcp
```

### arivon ssh

Manage SSH configuration.

```bash
arivon ssh status
arivon ssh audit
arivon ssh harden
arivon ssh keys
```

### arivon docker

Manage Docker.

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

### arivon network

Network management.

```bash
arivon network status
arivon network interfaces
arivon network routes
arivon network dns
arivon network test
```

### arivon storage

Storage management.

```bash
arivon storage status
arivon storage disks
arivon storage mounts
arivon storage health
arivon storage smart
```

### arivon backup

Backup management.

```bash
arivon backup status
arivon backup create
arivon backup restore <snapshot>
```

### arivon tailscale

Manage Tailscale (optional).

```bash
arivon tailscale install
arivon tailscale status
arivon tailscale connect
arivon tailscale disconnect
```

### arivon benchmark

Run system benchmark.

```bash
arivon benchmark
```

### arivon monitoring

Monitoring management.

```bash
arivon monitoring status
arivon monitoring metrics
```
