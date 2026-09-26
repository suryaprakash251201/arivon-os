# Arivon OS — Upgrades

## System Updates

```bash
arivon update check
arivon update security
arivon update system
```

## Arivon CLI Updates

```bash
apt update
apt install --only-upgrade arivon-cli
```

## Debian Upgrades

Standard Debian upgrade procedure:

```bash
apt update
apt upgrade
apt dist-upgrade
```

## Kernel Updates

```bash
apt update
apt install linux-image-amd64
reboot
```

## Backup Before Upgrades

```bash
arivon backup create
```

## Rollback

If an upgrade fails:

```bash
arivon doctor
arivon backup restore <snapshot>
```
