package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Manage SSH configuration",
}

var sshStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show SSH status",
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := exec.Command("systemctl", "is-active", "ssh").Output()
		status := strings.TrimSpace(string(out))
		if status == "active" {
			fmt.Println("SSH: Running")
		} else {
			fmt.Println("SSH: Stopped")
		}

		out, _ = exec.Command("ss", "-tlnp").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, line := range lines {
			if strings.Contains(line, ":22 ") {
				fmt.Println("Port:  22")
				break
			}
		}

		if sshDropinExists() {
			fmt.Println("Config: /etc/ssh/sshd_config.d/arivon.conf")
			fmt.Printf("  Root login: %s\n", sshDropinValue("PermitRootLogin"))
			fmt.Printf("  Password auth: %s\n", sshDropinValue("PasswordAuthentication"))
		}
	},
}

var sshAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit SSH security",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SSH AUDIT")
		fmt.Println()

		issues := 0

		if !sshDropinExists() {
			fmt.Println("  WARNING  No Arivon SSH drop-in config found")
			issues++
		} else {
			if sshDropinValue("PermitRootLogin") != "no" {
				fmt.Println("  WARNING  Root login not disabled")
				issues++
			} else {
				fmt.Println("  PASS     Root login disabled")
			}

			if sshDropinValue("PasswordAuthentication") != "no" {
				fmt.Println("  WARNING  Password authentication not disabled")
				issues++
			} else {
				fmt.Println("  PASS     Password authentication disabled")
			}
		}

		out, _ := exec.Command("sshd", "-T").Output()
		config := string(out)
		if strings.Contains(config, "permitrootlogin yes") {
			fmt.Println("  CRITICAL sshd reports root login allowed")
			issues++
		}

		fmt.Println()
		if issues == 0 {
			fmt.Println("SSH configuration is secure.")
		} else {
			fmt.Printf("%d issue(s) found.\n", issues)
		}
	},
}

var sshHardenCmd = &cobra.Command{
	Use:   "harden",
	Short: "Apply SSH hardening",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Applying SSH hardening...")

		if err := os.MkdirAll("/etc/ssh/sshd_config.d", 0755); err != nil {
			fmt.Println("Failed to create sshd_config.d directory.")
			os.Exit(1)
		}

		dropin := `PermitRootLogin no
PasswordAuthentication no
PubkeyAuthentication yes
AuthenticationMethods publickey
MaxAuthTries 3
MaxSessions 5
LoginGraceTime 30
ClientAliveInterval 300
ClientAliveCountMax 2
X11Forwarding no
AllowTcpForwarding no
PermitTunnel no
`
		if err := os.WriteFile("/etc/ssh/sshd_config.d/arivon.conf", []byte(dropin), 0644); err != nil {
			fmt.Println("Failed to write SSH config.")
			os.Exit(1)
		}

		out, err := exec.Command("sshd", "-t").CombinedOutput()
		if err != nil {
			fmt.Printf("SSH config test failed: %s\n", string(out))
			os.Exit(1)
		}

		fmt.Println("SSH hardening applied.")
		fmt.Println("Run: systemctl reload ssh")
	},
}

var sshKeysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List SSH authorized keys",
	Run: func(cmd *cobra.Command, args []string) {
		home := os.Getenv("HOME")
		if home == "" {
			home = "/root"
		}
		authorizedKeys := home + "/.ssh/authorized_keys"
		out, err := exec.Command("cat", authorizedKeys).Output()
		if err != nil {
			fmt.Println("No authorized_keys found.")
			return
		}
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		fmt.Printf("Authorized keys (%d):\n", len(lines))
		for i, line := range lines {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				fmt.Printf("  %d. %s\n", i+1, fields[2])
			} else if len(fields) > 0 {
				fmt.Printf("  %d. %s\n", i+1, fields[0])
			}
		}
	},
}

func init() {
	sshCmd.AddCommand(sshStatusCmd)
	sshCmd.AddCommand(sshAuditCmd)
	sshCmd.AddCommand(sshHardenCmd)
	sshCmd.AddCommand(sshKeysCmd)
	rootCmd.AddCommand(sshCmd)
}
