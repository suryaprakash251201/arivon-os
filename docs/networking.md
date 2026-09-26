# Arivon OS — Networking

## Status

```bash
arivon network status
arivon network interfaces
arivon network routes
arivon network dns
```

## Diagnostics

```bash
arivon network test
```

Runs:
- Gateway ping
- DNS lookup
- Socket inspection

## Configuration

Networking uses `systemd-networkd`. Configuration files:

```
/etc/systemd/network/*.network
/etc/systemd/network/*.netdev
```

## Supported

- DHCP
- Static IPv4
- IPv6
- DNS
- Routing
- VLAN
- Bridges
