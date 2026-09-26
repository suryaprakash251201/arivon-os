#!/usr/bin/env bash
set -euo pipefail

# Arivon OS development environment setup
# Run this on a fresh Ubuntu/Debian machine to set up the dev environment.

echo "Arivon OS Development Setup"
echo ""

# Install Go
if ! command -v go &>/dev/null; then
    echo "Installing Go..."
    GO_VERSION="1.26.1"
    wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -O /tmp/go.tar.gz
    sudo tar -C /usr/local -xzf /tmp/go.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
fi

# Install build tools
echo "Installing build tools..."
sudo apt-get update
sudo apt-get install -y \
    git \
    shellcheck \
    golangci-lint \
    mkosi \
    qemu-utils \
    qemu-system-x86 \
    dpkg-dev \
    debhelper \
    bats \
    python3-pip

# Install Python test dependencies
pip3 install pytest paramiko

echo ""
echo "Development environment ready."
echo "Run: cd cli && go build -o ../build/arivon ."
