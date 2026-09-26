"""
Arivon OS Docker integration tests.

Tests Docker installation and basic functionality.
"""

import pytest
import os
import time
import paramiko

SSH_USER = os.environ.get("ARIVON_SSH_USER", "admin")
SSH_PORT = int(os.environ.get("ARIVON_SSH_PORT", "2222"))


def run_ssh_command(cmd):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect("127.0.0.1", port=SSH_PORT, username=SSH_USER,
                   key_filename=os.path.expanduser("~/.ssh/id_rsa"))
    stdin, stdout, stderr = client.exec_command(cmd)
    output = stdout.read().decode()
    client.close()
    return output


def test_docker_installed():
    output = run_ssh_command("which docker")
    assert "/usr/bin/docker" in output


def test_docker_running():
    output = run_ssh_command("systemctl is-active docker")
    assert "active" in output


def test_docker_ps():
    output = run_ssh_command("docker ps")
    assert "CONTAINER" in output


def test_docker_images():
    output = run_ssh_command("docker images")
    assert "REPOSITORY" in output


def test_docker_compose():
    output = run_ssh_command("docker compose version")
    assert "version" in output.lower()


def test_docker_audit():
    output = run_ssh_command("arivon docker audit")
    assert "DOCKER" in output or "No running containers" in output
