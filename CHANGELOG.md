# Changelog

All notable changes to Arivon OS will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Go CLI with 18 commands (info, status, health, doctor, update, logs,
  services, security, firewall, ssh, docker, network, storage, backup,
  tailscale, benchmark, monitoring)
- TUI installer with disk detection and partitioning
- Build system with mkosi (minimal, server, docker, network, full profiles)
- Cloud image support
- systemd service units (health, security audit, first boot)
- Security defaults (sshd hardening, nftables)
- CI/CD pipeline (GitHub Actions)
- Release pipeline
- Test framework (Go unit tests, bats shell tests, pytest integration tests)
- Documentation (installation, getting started, building, development,
  security, docker, networking, storage, backups, monitoring,
  troubleshooting, upgrades, recovery)
- ShellCheck validation
- Pre-commit hooks
- ARM64 cross-compilation support
