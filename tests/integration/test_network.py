"""
Arivon OS network integration tests.

Tests network configuration and diagnostics.
"""

import pytest
import os
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


def test_network_status():
    output = run_ssh_command("arivon network status")
    assert "NETWORK STATUS" in output


def test_ip_address():
    output = run_ssh_command("arivon status")
    assert "IP" in output or "Address" in output


def test_dns_configured():
    output = run_ssh_command("cat /etc/resolv.conf")
    assert "nameserver" in output


def test_network_interfaces():
    output = run_ssh_command("arivon network interfaces")
    assert "lo" in output or "eth" in output


def test_network_test():
    output = run_ssh_command("arivon network test")
    assert "DIAGNOSTICS" in output
