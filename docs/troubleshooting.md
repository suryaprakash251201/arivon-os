# Arivon OS — Troubleshooting

## Boot Issues

### System does not boot
- Check UEFI/BIOS settings
- Verify boot mode matches installation mode
- Try recovery boot entry

### Emergency mode
- Boot with `single` kernel parameter
- Run `arivon doctor` to diagnose

## SSH Issues

### Cannot connect
```bash
arivon ssh status
systemctl status ssh
ss -tlnp | grep :22
```

### Locked out
- Use console access
- Check `/etc/ssh/sshd_config.d/arivon.conf`
- Run `sshd -t` to validate config

## Network Issues

### No IP address
```bash
arivon network status
arivon network test
systemctl status systemd-networkd
```

### DNS not working
```bash
arivon network dns
cat /etc/resolv.conf
```

## Docker Issues

### Docker not starting
```bash
arivon docker status
systemctl status docker
journalctl -u docker
```

### Container won't start
```bash
docker logs <container>
docker inspect <container>
```

## Recovery

### Reset firewall
```bash
nft flush ruleset
nft -f /etc/nftables.d/arivon.nft
```

### Reset SSH config
```bash
cp /etc/ssh/sshd_config.d/arivon.conf /etc/ssh/sshd_config.d/arivon.conf.bak
systemctl reload ssh
```

### Reinstall CLI
```bash
apt install --reinstall arivon-cli
```
