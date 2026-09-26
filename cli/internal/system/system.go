package system

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var osReleaseCache map[string]string

func OSRelease() map[string]string {
	if osReleaseCache != nil {
		return osReleaseCache
	}
	osReleaseCache = make(map[string]string)
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return osReleaseCache
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			osReleaseCache[strings.TrimSpace(parts[0])] = strings.Trim(strings.TrimSpace(parts[1]), "\"")
		}
	}
	return osReleaseCache
}

func DebianVersion() string {
	rel := OSRelease()
	if v, ok := rel["VERSION"]; ok {
		return v
	}
	if v, ok := rel["VERSION_ID"]; ok {
		return v
	}
	return "unknown"
}

func KernelVersion() string {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return "unknown"
	}
	fields := strings.Fields(string(data))
	if len(fields) >= 3 {
		return fields[2]
	}
	return "unknown"
}

func Architecture() string {
	return runtime.GOARCH
}

func Uptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}
	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return "unknown"
	}
	seconds, _ := strconv.ParseFloat(fields[0], 64)
	d := time.Duration(seconds) * time.Second
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd %dh", days, hours)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

type MemoryInfo struct {
	Total     uint64
	Available uint64
	SwapTotal uint64
	SwapFree  uint64
}

func Memory() MemoryInfo {
	var mem MemoryInfo
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		return mem
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		val, _ := strconv.ParseUint(fields[1], 10, 64)
		val *= 1024
		switch fields[0] {
		case "MemTotal:":
			mem.Total = val
		case "MemAvailable:":
			mem.Available = val
		case "SwapTotal:":
			mem.SwapTotal = val
		case "SwapFree:":
			mem.SwapFree = val
		}
	}
	return mem
}

func LoadAverage() string {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return "unknown"
	}
	fields := strings.Fields(string(data))
	if len(fields) >= 3 {
		return fields[0] + " " + fields[1] + " " + fields[2]
	}
	return "unknown"
}

func CPUCount() int {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return runtime.NumCPU()
	}
	defer f.Close()
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "processor") {
			count++
		}
	}
	if count == 0 {
		return runtime.NumCPU()
	}
	return count
}

func IPAddress() string {
	out, err := exec.Command("ip", "-4", "addr", "show").Output()
	if err != nil {
		return "unknown"
	}
	re := regexp.MustCompile(`inet (\d+\.\d+\.\d+\.\d+)`)
	matches := re.FindAllStringSubmatch(string(out), -1)
	for _, m := range matches {
		if m[1] != "127.0.0.1" {
			return m[1]
		}
	}
	return "unknown"
}

func ServiceStatus(name string) string {
	out, err := exec.Command("systemctl", "is-active", name).Output()
	if err != nil {
		return "inactive"
	}
	status := strings.TrimSpace(string(out))
	if status == "active" {
		return "Running"
	}
	return "Stopped"
}

func DockerStatus() string {
	out, err := exec.Command("systemctl", "is-active", "docker").Output()
	if err != nil {
		return "Not installed"
	}
	status := strings.TrimSpace(string(out))
	if status == "active" {
		return "Running"
	}
	return "Stopped"
}

func FirewallStatus() string {
	out, err := exec.Command("nft", "list", "ruleset").Output()
	if err != nil {
		return "Inactive"
	}
	if len(out) > 0 {
		return "Active"
	}
	return "Inactive"
}

func FailedServices() []string {
	var failed []string
	out, err := exec.Command("systemctl", "--failed", "--no-legend", "--no-pager").Output()
	if err != nil {
		return failed
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			failed = append(failed, line)
		}
	}
	return failed
}

func PendingUpdates() int {
	out, err := exec.Command("apt", "list", "--upgradable").Output()
	if err != nil {
		return 0
	}
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "upgradable") {
			count++
		}
	}
	return count
}

func DiskUsage() string {
	out, err := exec.Command("df", "-h", "/").Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) >= 2 {
		fields := strings.Fields(lines[1])
		if len(fields) >= 5 {
			return fields[2] + " / " + fields[1]
		}
	}
	return "unknown"
}
