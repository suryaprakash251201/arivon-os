# Arivon OS — Monitoring

## Base Monitoring

```bash
arivon monitoring status
arivon monitoring metrics
```

Lightweight built-in monitoring for CPU, RAM, disk, load, network, and services.

## Optional Integrations

Not installed by default:

- Prometheus + node_exporter
- Grafana
- Netdata

Install via APT or as Docker containers.

## Health Checks

```bash
arivon health
arivon doctor
```

Health checks run automatically every 5 minutes via `arivon-health.timer`.

Security audits run daily via `arivon-security-audit.timer`.
