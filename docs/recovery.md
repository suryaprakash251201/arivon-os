# Arivon OS — Recovery

## Boot Recovery

1. Boot from Arivon OS ISO
2. Select "Recovery" boot entry
3. Mount root filesystem:
   ```bash
   mount /dev/sda1 /mnt
   chroot /mnt
   ```
4. Fix the issue
5. Reboot

## Filesystem Recovery

```bash
fsck /dev/sda1
```

## Network Recovery

```bash
ip link set eth0 up
dhclient eth0
```

## Docker Recovery

```bash
systemctl stop docker
rm -rf /var/lib/docker
systemctl start docker
```

## Full System Recovery

1. Boot from ISO
2. Run `arivon-installer`
3. Select "Recover" option
4. Restore from backup:
   ```bash
   arivon backup restore <snapshot>
   ```

## Emergency Console

If the system is unbootable:
- Use ISO recovery mode
- Mount and chroot into the installed system
- Fix configuration or restore from backup
