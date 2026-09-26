# Arivon OS — Docker

## Installation

```bash
arivon docker install
```

Installs Docker Engine, CLI, containerd, and Compose plugin from Docker's
official repository.

## Management

```bash
arivon docker status
arivon docker ps
arivon docker images
arivon docker volumes
arivon docker networks
```

## Security Audit

```bash
arivon docker audit
```

Detects:
- Privileged containers
- Host networking
- Host PID namespace
- Dangerous capabilities
- Docker socket mounts
- Host filesystem mounts
- Exposed ports
- Containers running as root
- Missing resource limits

## Cleanup

```bash
arivon docker cleanup
```

Removes unused containers, images, volumes, and networks.
Requires confirmation before deletion.
