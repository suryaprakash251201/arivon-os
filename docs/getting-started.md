# Arivon OS — Getting Started

## First Login

After installation, log in via console or SSH:

```bash
ssh admin@<ip-address>
```

## Essential Commands

```bash
arion status          # system overview
arivon info           # OS identity
arivon health         # health check
arivon doctor         # diagnose problems
arivon update check   # check for updates
arivon security audit # security assessment
```

## JSON Output

All commands support `--json` for machine-readable output:

```bash
arivon status --json
```

## Configuration

Edit `/etc/arivon/config.toml` to change system settings.

## Documentation

See the `docs/` directory for full documentation on SSH, Docker,
networking, firewall, storage, backups, and monitoring.
