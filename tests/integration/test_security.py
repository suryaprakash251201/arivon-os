"""
Arivon OS security integration tests.

Tests SSH, firewall, and security audit functionality.
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


def test_ssh_hardening():
    output = run_ssh_command("cat /etc/ssh/sshd_config.d/arivon.conf")
    assert "PermitRootLogin no" in output
    assert "PasswordAuthentication no" in output


def test_firewall_rules():
    output = run_ssh_command("nft list ruleset")
    assert "table inet arivon" in output


def test_security_audit():
    output = run_ssh_command("arivon security audit")
    assert "PASS" in output or "WARNING" in output


def test_no_root_login():
    output = run_ssh_command("grep -c 'PermitRootLogin no' /etc/ssh/sshd_config.d/arivon.conf")
    assert "1" in output


def test_ssh_banner():
    output = run_ssh_command("cat /etc/arivon/ssh_banner")
    assert "ARIVON" in output
