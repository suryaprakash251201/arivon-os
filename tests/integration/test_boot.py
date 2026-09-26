"""
Arivon OS VM integration tests.

Requires: QEMU, pytest, libvirt (or direct QEMU invocation).
These tests boot the OS image in a VM and verify basic functionality.
"""

import pytest
import subprocess
import time
import os
import paramiko

IMAGE_PATH = os.environ.get("ARIVON_IMAGE", "images/output/arivon-os-minimal.raw")
SSH_USER = os.environ.get("ARIVON_SSH_USER", "admin")
SSH_PORT = int(os.environ.get("ARIVON_SSH_PORT", "2222"))
VM_MEMORY = os.environ.get("ARIVON_VM_MEMORY", "512")


@pytest.fixture(scope="module")
def vm():
    """Boot Arivon OS in QEMU and yield SSH connection info."""
    if not os.path.exists(IMAGE_PATH):
        pytest.skip(f"Image not found: {IMAGE_PATH}")

    qemu_cmd = [
        "qemu-system-x86_64",
        "-m", VM_MEMORY,
        "-drive", f"file={IMAGE_PATH},format=raw",
        "-netdev", f"user,id=net0,hostfwd=tcp::{SSH_PORT}-:22",
        "-device", "e1000,netdev=net0",
        "-display", "none",
        "-daemonize",
    ]

    subprocess.run(qemu_cmd, check=True)

    # Wait for SSH to become available
    for _ in range(60):
        try:
            client = paramiko.SSHClient()
            client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
            client.connect("127.0.0.1", port=SSH_PORT, username=SSH_USER,
                           key_filename=os.path.expanduser("~/.ssh/id_rsa"),
                           timeout=5)
            client.close()
            break
        except Exception:
            time.sleep(2)
    else:
        pytest.fail("VM did not become ready in time")

    yield {"host": "127.0.0.1", "port": SSH_PORT, "user": SSH_USER}

    # Cleanup
    subprocess.run(["pkill", "-f", "qemu-system-x86_64"], check=False)


def run_ssh_command(vm, cmd):
    """Run a command over SSH and return stdout."""
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(vm["host"], port=vm["port"], username=vm["user"],
                   key_filename=os.path.expanduser("~/.ssh/id_rsa"))
    stdin, stdout, stderr = client.exec_command(cmd)
    output = stdout.read().decode()
    client.close()
    return output


def test_boot_ssh_accessible(vm):
    """VM boots and SSH is accessible."""
    output = run_ssh_command(vm, "echo hello")
    assert "hello" in output


def test_arivon_binary_exists(vm):
    """arivon CLI is installed."""
    output = run_ssh_command(vm, "which arivon")
    assert "/usr/bin/arivon" in output


def test_arivon_info(vm):
    """arivon info returns version information."""
    output = run_ssh_command(vm, "arivon info")
    assert "Version" in output
    assert "Kernel" in output


def test_arivon_status(vm):
    """arivon status returns system status."""
    output = run_ssh_command(vm, "arivon status")
    assert "Overall" in output


def test_ssh_running(vm):
    """SSH service is active."""
    output = run_ssh_command(vm, "systemctl is-active ssh")
    assert "active" in output


def test_systemd_running(vm):
    """systemd is PID 1."""
    output = run_ssh_command(vm, "ps -p 1 -o comm=")
    assert "systemd" in output
