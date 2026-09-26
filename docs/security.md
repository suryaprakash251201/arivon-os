# Arivon OS — Security

## Security Audit

```bash
arivon security audit
```

Checks:
- SSH configuration (root login, password auth)
- Firewall status
- Available updates
- Failed services
- Listening sockets

## Firewall

```bash
arivon firewall status
arivon firewall enable
arivon firewall disable
arivon firewall list
arivon firewall allow 80/tcp
arivon firewall remove 80/tcp
```

Default policy: deny all inbound, allow loopback, established/related, SSH.

## SSH

```bash
arivon ssh status
arivon ssh audit
arivon ssh harden
arivon ssh keys
```

Hardening applies `/etc/ssh/sshd_config.d/arivon.conf` with:
- Root login disabled
- Password authentication disabled
- Public key authentication only
- Connection limits
- Session limits

## Updates

```bash
arivon update check
arivon update security
arivon update system
arivon update history
```

Relies on Debian security infrastructure. No custom update mechanism.
