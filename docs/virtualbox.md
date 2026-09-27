# Run Arivon OS in VirtualBox

## What you need

- VirtualBox 7.x on Windows (Extension Pack **not** required)
- The Arivon disk image (`arivon-os-minimal.raw`, ~1.1 GB)
- ~4 GB free disk, 2 GB free RAM

> The image is a bootable GPT disk (`.raw`), not an ISO — you attach
> it as a **hard disk**, not as a CD. (`dd`-able to USB for bare metal.)

## 1. Download the image

```powershell
# latest successful ISO Build run (minimal profile)
gh run download --repo suryaprakash251201/arivon-os `
  --pattern "arivon-os-minimal-image" --dir .\arivon-vm
```

Or download `arivon-os-minimal-image.zip` from the
[Actions](https://github.com/suryaprakash251201/arivon-os/actions/workflows/iso.yml)
page (GitHub login required) and unzip it. You need the `.raw` file.

## 2. Convert raw → VDI

VirtualBox cannot boot `.raw` directly. Convert once:

```powershell
& "C:\Program Files\Oracle\VirtualBox\VBoxManage.exe" convertfromraw `
  .\arivon-vm\arivon-os-minimal.raw .\arivon-vm\arivon-os-minimal.vdi --format VDI
```

## 3. Create the VM

One-shot setup (adjust RAM/CPUs to taste):

```powershell
$VBM = "C:\Program Files\Oracle\VirtualBox\VBoxManage.exe"
& $VBM createvm --name "ArivonOS" --ostype "Debian_64" --register
& $VBM modifyvm "ArivonOS" --memory 2048 --cpus 2 --firmware efi `
  --nic1 nat --natpf1 "ssh,tcp,,2222,,22" --audio none --usb off
& $VBM storagectl "ArivonOS" --name "SATA" --add sata --controller IntelAhci
& $VBM storageattach "ArivonOS" --storagectl "SATA" --port 0 --device 0 `
  --type hdd --medium "C:\path\to\arivon-vm\arivon-os-minimal.vdi"
```

Or use the GUI: New → Type *Linux*, Version *Debian (64-bit)* →
use existing disk (the `.vdi`) → Settings → System →
**☑ Enable EFI (special OSes only)** → Network → Advanced →
Port Forwarding → `ssh | TCP | 127.0.0.1 | 2222 | | 22`.

> **EFI is mandatory.** The image uses systemd-boot on a GPT disk and
> will not boot in legacy BIOS mode. Leave Secure Boot off (default).

## 4. First boot — create your user

v0.1 images ship **no login user** (Debian locks `root` by default),
so create one via the VM console on first boot:

1. Start the VM **with GUI** (not headless) the first time.
2. At the `systemd-boot` menu, hold **Space** if it flashes past,
   select the entry, press **`e`**, append `init=/bin/bash`, press Enter.
3. At the root shell:
   ```bash
   mount -o remount,rw /
   useradd -m -s /bin/bash -G sudo admin
   passwd admin
   exec /sbin/init
   ```
4. Log in on the console as `admin`.

## 5. SSH in and verify

```powershell
ssh -p 2222 admin@127.0.0.1
```

```bash
systemctl is-active ssh
ip -4 addr show
cat /etc/os-release
```

## 6. Install the Arivon CLI (optional, v0.1)

The CLI is not preinstalled in v0.1 images. Install the `.deb`
from CI artifacts or a release, then:

```bash
sudo dpkg -i arivon-cli_*.deb
arivon status
arivon health
arivon security audit
```

## Troubleshooting

| Symptom | Fix |
|---|---|
| `FATAL: No bootable medium` / EFI shell | EFI not enabled → `modifyvm --firmware efi` |
| Boots to `login:` but no credentials work | Do step 4 (create user via `init=/bin/bash`) |
| `ssh: connect to host … Connection refused` | VM still booting; check port-forward rule; use console |
| Black screen after kernel messages | Normal on first boot (initramfs + journal); wait 30 s |
| Want a fresh start | Snapshots: `VBoxManage snapshot "ArivonOS" take clean` |
