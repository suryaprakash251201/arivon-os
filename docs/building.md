# Arivon OS — Building From Source

## Prerequisites

- Linux host (Ubuntu 22.04+ or Debian 12+)
- Go 1.22+
- mkosi
- qemu-utils
- dpkg-deb

## Build Steps

```bash
git clone https://github.com/arivon/arivon-os.git
cd arivon-os

# Build everything
./build/build.sh

# Or build individually:
cd cli && go build -o ../build/arivon .
cd ../packages && ./build-deb.sh 0.1.0 amd64
cd ../images && sudo mkosi build --profile minimal
```

## Output

Artifacts are placed in `images/output/`:
- `arivon-cli_<version>_<arch>.deb`
- `arivon-os-<profile>/` (disk image directory)

## Testing

```bash
# Unit tests
cd cli && go test ./...

# Shell tests
bats tests/

# VM integration tests (requires QEMU)
pytest tests/integration/
```
