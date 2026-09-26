#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
VERSION="${1:-}"

if [ -z "$VERSION" ]; then
    echo "Usage: ./build/release.sh <version>"
    echo "Example: ./build/release.sh 0.1.0"
    exit 1
fi

if [[ ! "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version format. Use semantic versioning: X.Y.Z"
    exit 1
fi

echo "Arivon OS Release Build"
echo "  Version: $VERSION"
echo ""

export ARIVON_VERSION="$VERSION"

echo "[1/6] Building CLI..."
cd "$ROOT_DIR/cli"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags "-X github.com/arivon/arivon-os/cli/cmd.arivonVersion=$VERSION" \
  -o "$ROOT_DIR/build/arivon" .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
  -ldflags "-X github.com/arivon/arivon-os/cli/cmd.arivonVersion=$VERSION" \
  -o "$ROOT_DIR/build/arivon-arm64" .
echo "  -> Done"

echo "[2/6] Building installer..."
cd "$ROOT_DIR/installer"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -o "$ROOT_DIR/build/arivon-installer" .
echo "  -> Done"

echo "[3/6] Building .deb packages..."
cd "$ROOT_DIR/packages"
./build-deb.sh "$VERSION" amd64
./build-deb.sh "$VERSION" arm64
echo "  -> Done"

echo "[4/6] Building ISO images..."
cd "$ROOT_DIR/images"
sudo mkosi build --profile minimal
sudo mkosi build --profile server
sudo mkosi build --profile docker
echo "  -> Done"

echo "[5/6] Converting to qcow2..."
cd "$ROOT_DIR/images/output"
for img in */; do
    name=$(basename "$img")
    if [ -f "$img/$name.raw" ]; then
        qemu-img convert -f raw -O qcow2 "$img/$name.raw" "$img/$name.qcow2"
        echo "  -> $name.qcow2"
    fi
done

echo "[6/6] Generating checksums..."
cd "$ROOT_DIR/images/output"
sha256sum *.deb *.qcow2 */*.qcow2 2>/dev/null > SHA256SUMS || true
echo "  -> SHA256SUMS"

echo ""
echo "Release artifacts in: $ROOT_DIR/images/output"
echo "  - arivon-cli_${VERSION}_amd64.deb"
echo "  - arivon-cli_${VERSION}_arm64.deb"
echo "  - arivon-os-minimal/ (ISO + raw)"
echo "  - arivon-os-server/ (ISO + raw)"
echo "  - arivon-os-docker/ (ISO + raw)"
echo "  - SHA256SUMS"
echo ""
echo "Tag and push:"
echo "  git tag v$VERSION"
echo "  git push origin v$VERSION"
