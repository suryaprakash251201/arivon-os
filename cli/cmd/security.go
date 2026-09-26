package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var securityCmd = &cobra.Command{
	Use:   "security",
	Short: "Security management",
}

var securityAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Run security audit",
	Run: func(cmd *cobra.Command, args []string) {
		runSecurityAudit()
	},
}

func init() {
	securityCmd.AddCommand(securityAuditCmd)
	rootCmd.AddCommand(securityCmd)
}

func runSecurityAudit() {
	fmt.Println("ARIVON OS SECURITY AUDIT")
	fmt.Println()

	passed := 0
	warnings := 0
	critical := 0

	fmt.Println("[SSH]")
	if sshDropinExists() && sshDropinValue("PermitRootLogin") == "no" {
		fmt.Println("  PASS     Root login disabled")
		passed++
	} else {
		fmt.Println("  WARNING  Root login not explicitly disabled")
		warnings++
	}

	if sshDropinExists() && sshDropinValue("PasswordAuthentication") == "no" {
		fmt.Println("  PASS     Password authentication disabled")
		passed++
	} else {
		fmt.Println("  WARNING  Password authentication not explicitly disabled")
		warnings++
	}
	fmt.Println()

	fmt.Println("[Firewall]")
	if system.FirewallStatus() == "Active" {
		fmt.Println("  PASS     Firewall active")
		passed++
	} else {
		fmt.Println("  WARNING  Firewall not active")
		warnings++
	}
	fmt.Println()

	fmt.Println("[Updates]")
	updates := system.PendingUpdates()
	if updates == 0 {
		fmt.Println("  PASS     System up to date")
		passed++
	} else {
		fmt.Printf("  WARNING  %d updates available\n", updates)
		warnings++
	}
	fmt.Println()

	fmt.Println("[Services]")
	failed := system.FailedServices()
	if len(failed) == 0 {
		fmt.Println("  PASS     No failed services")
		passed++
	} else {
		fmt.Printf("  CRITICAL %d failed services\n", len(failed))
		critical++
	}
	fmt.Println()

	fmt.Println("[Network]")
	out, _ := exec.Command("ss", "-tlnp").Output()
	listening := strings.TrimSpace(string(out))
	if listening != "" {
		lines := strings.Split(listening, "\n")
		fmt.Printf("  INFO     %d listening sockets\n", len(lines)-1)
	}
	fmt.Println()

	fmt.Printf("RESULT: %d passed, %d warnings, %d critical\n", passed, warnings, critical)

	if critical > 0 {
		exitCode = 1
	}
}

var exitCode int

func sshDropinExists() bool {
	_, err := exec.Command("test", "-f", "/etc/ssh/sshd_config.d/arivon.conf").Output()
	return err == nil
}

func sshDropinValue(key string) string {
	out, _ := exec.Command("grep", "-E", "^"+key, "/etc/ssh/sshd_config.d/arivon.conf").Output()
	fields := strings.Fields(string(out))
	if len(fields) >= 2 {
		return fields[1]
	}
	return ""
}
