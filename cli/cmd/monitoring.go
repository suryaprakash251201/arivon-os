package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var monitoringCmd = &cobra.Command{
	Use:   "monitoring",
	Short: "Monitoring management",
}

var monitoringStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show monitoring status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MONITORING STATUS")
		fmt.Println()

		services := []string{"prometheus", "grafana-server", "netdata", "node_exporter"}
		for _, svc := range services {
			out, _ := exec.Command("systemctl", "is-active", svc).Output()
			status := strings.TrimSpace(string(out))
			if status == "active" {
				fmt.Printf("  %s: Running\n", svc)
			} else {
				fmt.Printf("  %s: Not installed\n", svc)
			}
		}
	},
}

var monitoringMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show current metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CURRENT METRICS")
		fmt.Println()

		out, _ := exec.Command("cat", "/proc/stat").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) > 0 {
			fields := strings.Fields(lines[0])
			if len(fields) >= 5 {
				user := fields[1]
				system_ := fields[3]
				idle := fields[4]
				fmt.Printf("CPU:   user=%s system=%s idle=%s\n", user, system_, idle)
			}
		}

		out, _ = exec.Command("cat", "/proc/meminfo").Output()
		for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			if strings.HasPrefix(line, "MemTotal:") || strings.HasPrefix(line, "MemAvailable:") {
				fmt.Printf("RAM:   %s\n", strings.TrimSpace(line))
			}
		}

		out, _ = exec.Command("cat", "/proc/loadavg").Output()
		fmt.Printf("Load:  %s\n", strings.TrimSpace(string(out)))

		out, _ = exec.Command("df", "-h", "/").Output()
		dfLines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(dfLines) >= 2 {
			fields := strings.Fields(dfLines[1])
			if len(fields) >= 5 {
				fmt.Printf("Disk:  %s / %s (%s)\n", fields[2], fields[1], fields[4])
			}
		}
	},
}

func init() {
	monitoringCmd.AddCommand(monitoringStatusCmd)
	monitoringCmd.AddCommand(monitoringMetricsCmd)
	rootCmd.AddCommand(monitoringCmd)
}
