#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
BUILD_DIR="$ROOT_DIR/build"
OUTPUT_DIR="$ROOT_DIR/images/output"
VERSION="${ARIVON_VERSION:-0.1.0-dev}"
ARCH="${TARGET_ARCH:-amd64}"

echo "Arivon OS Build System"
echo "  Version: $VERSION"
echo "  Arch:    $ARCH"
echo ""

mkdir -p "$OUTPUT_DIR"

echo "[1/5] Building arivon CLI..."
cd "$ROOT_DIR/cli"
CGO_ENABLED=0 GOOS=linux GOARCH="$ARCH" go build \
  -ldflags "-X github.com/arivon/arivon-os/cli/cmd.arivonVersion=$VERSION" \
  -o "$BUILD_DIR/arivon" .
echo "  -> $BUILD_DIR/arivon"

echo "[2/5] Building arivon installer..."
cd "$ROOT_DIR/installer"
CGO_ENABLED=0 GOOS=linux GOARCH="$ARCH" go build \
  -o "$BUILD_DIR/arivon-installer" .
echo "  -> $BUILD_DIR/arivon-installer"

echo "[3/5] Building .deb package..."
cd "$ROOT_DIR/packages"
./build-deb.sh "$VERSION" "$ARCH"
echo "  -> $OUTPUT_DIR/arivon-cli_${VERSION}_${ARCH}.deb"

echo "[4/5] Building OS image with mkosi..."
cd "$ROOT_DIR/images"
mkosi --version
echo "  (mkosi build requires Linux host with mkosi installed)"
echo "  Run: sudo mkosi build --profile minimal"

echo "[5/5] Build complete."
echo "  Artifacts in: $OUTPUT_DIR"
