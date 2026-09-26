# Arivon OS — Installation

## Quick Start

1. Download the latest ISO from GitHub Releases
2. Write to USB: `dd if=arivon-os-0.1.0.iso of=/dev/sdX bs=4M status=progress`
3. Boot from USB
4. Run the installer: `arivon install`
5. Reboot
6. SSH in: `ssh admin@<ip>`

## Manual Installation (Advanced)

The installer supports:
- UEFI and BIOS boot
- Automatic disk detection
- Manual partitioning (ext4, XFS, Btrfs)
- Hostname, user, timezone, locale configuration
- SSH key setup
- Optional Docker installation
- Optional security profile

## Profiles

| Profile   | Use Case                          |
|-----------|-----------------------------------|
| minimal   | Bare server, 512 MB RAM target   |
| server    | General purpose server            |
| docker    | Docker host                       |
| network   | Network diagnostics / firewall    |
| full      | Everything included               |

## Post-Install

```bash
arivon status
arivon health
arivon security audit
```
