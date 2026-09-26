package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose common server problems",
	Run: func(cmd *cobra.Command, args []string) {
		runDoctor()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func runDoctor() {
	fmt.Println("ARIVON DOCTOR")
	fmt.Println()

	problems := 0

	if system.IPAddress() == "unknown" {
		reportProblem("No IP address", "Network interface not configured", "Check network configuration", "Low")
		problems++
	}

	if system.ServiceStatus("ssh") != "Running" {
		reportProblem("SSH not running", "SSH service is inactive", "Run: systemctl start ssh", "High")
		problems++
	}

	if system.DockerStatus() == "Stopped" {
		reportProblem("Docker not running", "Docker service is inactive", "Run: systemctl start docker", "Medium")
		problems++
	}

	if system.FirewallStatus() == "Inactive" {
		reportProblem("Firewall inactive", "nftables rules not loaded", "Run: arivon firewall enable", "Medium")
		problems++
	}

	out, _ := exec.Command("systemctl", "--failed", "--no-legend", "--no-pager").Output()
	failed := strings.TrimSpace(string(out))
	if failed != "" {
		reportProblem("Failed services", failed, "Run: systemctl --failed", "High")
		problems++
	}

	out, _ = exec.Command("df", "-h", "/").Output()
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) >= 2 {
		fields := strings.Fields(lines[1])
		if len(fields) >= 5 {
			pct := strings.TrimSuffix(fields[4], "%")
			if pct == "90" || pct == "95" || pct == "100" {
				reportProblem("Disk almost full", fields[4]+" used", "Free up disk space", "High")
				problems++
			}
		}
	}

	if problems == 0 {
		fmt.Println("No problems detected.")
	} else {
		fmt.Printf("%d problem(s) detected.\n", problems)
	}
}

func reportProblem(problem, cause, fix, risk string) {
	fmt.Printf("Problem:   %s\n", problem)
	fmt.Printf("Cause:     %s\n", cause)
	fmt.Printf("Fix:       %s\n", fix)
	fmt.Printf("Risk:      %s\n", risk)
	fmt.Println()
}
