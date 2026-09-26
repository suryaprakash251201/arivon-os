# Arivon OS — Storage

## Status

```bash
arivon storage status
arivon storage disks
arivon storage mounts
arivon storage health
arivon storage smart
```

## Supported Filesystems

- ext4
- XFS
- Btrfs
- swap

## SMART

```bash
arivon storage smart
```

Shows SMART health for all disks (requires `smartmontools`).

## Mounts

```bash
arivon storage mounts
```

Shows all mounted filesystems with type and options.
