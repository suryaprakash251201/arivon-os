#!/usr/bin/env bash
set -euo pipefail

VERSION="${1:-0.1.0-dev}"
ARCH="${2:-amd64}"
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BUILD_DIR="$ROOT_DIR/build"
OUTPUT_DIR="$ROOT_DIR/images/output"
DEB_DIR="$BUILD_DIR/deb"

echo "Building arivon-cli $VERSION ($ARCH)..."

mkdir -p "$OUTPUT_DIR"
rm -rf "$DEB_DIR"
mkdir -p "$DEB_DIR/DEBIAN"
mkdir -p "$DEB_DIR/usr/bin"
mkdir -p "$DEB_DIR/etc/arivon"
mkdir -p "$DEB_DIR/etc/systemd/system"
mkdir -p "$DEB_DIR/etc/ssh/sshd_config.d"
mkdir -p "$DEB_DIR/etc/nftables.d"
mkdir -p "$DEB_DIR/etc/systemd/journald.conf.d"
mkdir -p "$DEB_DIR/etc/systemd/network"
mkdir -p "$DEB_DIR/usr/lib/arivon"

cp "$BUILD_DIR/arivon" "$DEB_DIR/usr/bin/arivon"
chmod 755 "$DEB_DIR/usr/bin/arivon"

cp "$BUILD_DIR/arivon-installer" "$DEB_DIR/usr/bin/arivon-installer"
chmod 755 "$DEB_DIR/usr/bin/arivon-installer"

cp "$ROOT_DIR/configs/arivon/config.toml" "$DEB_DIR/etc/arivon/config.toml"
cp "$ROOT_DIR/configs/sshd/arivon.conf" "$DEB_DIR/etc/ssh/sshd_config.d/arivon.conf"
cp "$ROOT_DIR/configs/ssh/banner" "$DEB_DIR/etc/arivon/ssh_banner"
cp "$ROOT_DIR/configs/nftables/arivon.nft" "$DEB_DIR/etc/nftables.d/arivon.nft"
cp "$ROOT_DIR/configs/journald/arivon.conf" "$DEB_DIR/etc/systemd/journald.conf.d/arivon.conf"
cp "$ROOT_DIR/configs/network/arivon.network" "$DEB_DIR/etc/systemd/network/arivon.network"

cp "$ROOT_DIR/services/arivon-firstboot.service" "$DEB_DIR/etc/systemd/system/"
cp "$ROOT_DIR/services/arivon-health.service" "$DEB_DIR/etc/systemd/system/"
cp "$ROOT_DIR/services/arivon-health.timer" "$DEB_DIR/etc/systemd/system/"
cp "$ROOT_DIR/services/arivon-security-audit.service" "$DEB_DIR/etc/systemd/system/"
cp "$ROOT_DIR/services/arivon-security-audit.timer" "$DEB_DIR/etc/systemd/system/"

cat > "$DEB_DIR/DEBIAN/control" <<EOF
Package: arivon-cli
Version: $VERSION
Architecture: $ARCH
Maintainer: Arivon OS Team
Depends: libc6
Section: admin
Priority: optional
Description: Arivon OS system administration CLI
 Intelligent Infrastructure.
EOF

cat > "$DEB_DIR/DEBIAN/postinst" <<'EOF'
#!/bin/bash
set -e
systemctl daemon-reload
systemctl enable arivon-firstboot.service
systemctl enable arivon-health.timer
systemctl enable arivon-security-audit.timer
echo "Arivon OS CLI installed."
echo "Run: arivon status"
EOF
chmod 755 "$DEB_DIR/DEBIAN/postinst"

dpkg-deb --build --root-owner-group "$DEB_DIR" "$OUTPUT_DIR/arivon-cli_${VERSION}_${ARCH}.deb"

echo "  -> $OUTPUT_DIR/arivon-cli_${VERSION}_${ARCH}.deb"
