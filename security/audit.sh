#!/usr/bin/env bash
set -euo pipefail

PASS=0
WARNING=0
CRITICAL=0

report() {
    local level="$1"
    local message="$2"
    echo "  $level  $message"
    case "$level" in
        PASS) ((PASS++)) ;;
        WARNING) ((WARNING++)) ;;
        CRITICAL) ((CRITICAL++)) ;;
    esac
}

echo "ARIVON OS SECURITY AUDIT"
echo ""

echo "[SSH]"
if [ -f /etc/ssh/sshd_config.d/arivon.conf ]; then
    if grep -q "^PermitRootLogin no" /etc/ssh/sshd_config.d/arivon.conf; then
        report "PASS" "Root login disabled"
    else
        report "WARNING" "Root login not explicitly disabled"
    fi

    if grep -q "^PasswordAuthentication no" /etc/ssh/sshd_config.d/arivon.conf; then
        report "PASS" "Password authentication disabled"
    else
        report "WARNING" "Password authentication not explicitly disabled"
    fi

    if grep -q "^MaxAuthTries" /etc/ssh/sshd_config.d/arivon.conf; then
        report "PASS" "Max auth tries configured"
    else
        report "WARNING" "Max auth tries not configured"
    fi
else
    report "WARNING" "No Arivon SSH drop-in config found"
fi
echo ""

echo "[Firewall]"
if systemctl is-active --quiet nftables 2>/dev/null; then
    report "PASS" "nftables service active"
elif nft list ruleset 2>/dev/null | grep -q "table inet"; then
    report "PASS" "nftables rules loaded"
else
    report "WARNING" "Firewall not active"
fi
echo ""

echo "[Updates]"
UPDATES=$(apt list --upgradable 2>/dev/null | grep -c "upgradable" || true)
if [ "$UPDATES" -eq 0 ]; then
    report "PASS" "System up to date"
else
    report "WARNING" "$UPDATES updates available"
fi
echo ""

echo "[Services]"
FAILED=$(systemctl --failed --no-legend --no-pager 2>/dev/null | wc -l)
if [ "$FAILED" -eq 0 ]; then
    report "PASS" "No failed services"
else
    report "CRITICAL" "$FAILED failed services"
fi
echo ""

echo "[Network]"
LISTENING=$(ss -tlnp 2>/dev/null | tail -n +2 | wc -l)
echo "  INFO  $LISTENING listening sockets"
echo ""

echo "[Disk]"
ROOT_USAGE=$(df / | tail -1 | awk '{print $5}' | tr -d '%')
if [ "$ROOT_USAGE" -gt 90 ]; then
    report "CRITICAL" "Root filesystem ${ROOT_USAGE}% full"
elif [ "$ROOT_USAGE" -gt 80 ]; then
    report "WARNING" "Root filesystem ${ROOT_USAGE}% full"
else
    report "PASS" "Root filesystem ${ROOT_USAGE}% used"
fi
echo ""

echo "[Permissions]"
WORLD_WRITABLE=$(find /etc /usr/bin /usr/sbin -type f -perm -0002 2>/dev/null | wc -l)
if [ "$WORLD_WRITABLE" -eq 0 ]; then
    report "PASS" "No world-writable files in system directories"
else
    report "WARNING" "$WORLD_WRITABLE world-writable files found"
fi
echo ""

echo "RESULT: $PASS passed, $WARNING warnings, $CRITICAL critical"
