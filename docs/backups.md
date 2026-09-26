# Arivon OS — Backups

## Overview

Backups use `restic` with S3-compatible storage support.

## Commands

```bash
arivon backup status
arivon backup create
arivon backup restore <snapshot>
```

## Configuration

Set the repository via environment variable:

```bash
export RESTIC_REPOSITORY=s3:https://backup.example.com/arivon
export RESTIC_PASSWORD=your-password
export AWS_ACCESS_KEY_ID=your-key
export AWS_SECRET_ACCESS_KEY=your-secret
```

Default repository: `/var/backups/arivon`

## Scheduling

Backups can be scheduled via systemd timers:

```bash
systemctl enable restic-backup.timer
systemctl start restic-backup.timer
```

## Safety

- Destructive operations require confirmation
- Backups are never automatically deleted unless retention is configured
- Restore requires explicit confirmation
