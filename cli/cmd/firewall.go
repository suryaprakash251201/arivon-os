package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Manage nftables firewall",
}

var firewallStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show firewall status",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("nft", "list", "ruleset").Output()
		if err != nil {
			fmt.Println("Firewall: Inactive")
			return
		}
		if len(out) > 0 {
			fmt.Println("Firewall: Active")
			fmt.Println()
			fmt.Println(string(out))
		} else {
			fmt.Println("Firewall: Inactive")
		}
	},
}

var firewallEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable firewall",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enabling firewall...")
		c := exec.Command("nft", "-f", "/etc/nftables.d/arivon.nft")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to enable firewall.")
			os.Exit(1)
		}
		fmt.Println("Firewall enabled.")
	},
}

var firewallDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable firewall",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Disabling firewall...")
		c := exec.Command("nft", "flush", "ruleset")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to disable firewall.")
			os.Exit(1)
		}
		fmt.Println("Firewall disabled.")
	},
}

var firewallListCmd = &cobra.Command{
	Use:   "list",
	Short: "List firewall rules",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("nft", "list", "ruleset").Output()
		if err != nil {
			fmt.Println("No firewall rules loaded.")
			return
		}
		fmt.Println(string(out))
	},
}

var firewallAllowCmd = &cobra.Command{
	Use:   "allow <port>/<protocol>",
	Short: "Allow incoming traffic",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rule := args[0]
		if !strings.Contains(rule, "/") {
			fmt.Println("Invalid rule format. Use: port/protocol (e.g. 80/tcp)")
			os.Exit(1)
		}
		parts := strings.SplitN(rule, "/", 2)
		port := parts[0]
		proto := parts[1]

		fmt.Printf("Allowing %s/%s...\n", port, proto)
		c := exec.Command("nft", "add", "rule", "inet", "arivon", "input", proto, "dport", port, "accept")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to add rule.")
			os.Exit(1)
		}
		fmt.Println("Rule added.")
	},
}

var firewallRemoveCmd = &cobra.Command{
	Use:   "remove <port>/<protocol>",
	Short: "Remove firewall rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rule := args[0]
		if !strings.Contains(rule, "/") {
			fmt.Println("Invalid rule format. Use: port/protocol (e.g. 80/tcp)")
			os.Exit(1)
		}
		parts := strings.SplitN(rule, "/", 2)
		port := parts[0]
		proto := parts[1]

		fmt.Printf("Removing %s/%s...\n", port, proto)
		c := exec.Command("nft", "delete", "rule", "inet", "arivon", "input", proto, "dport", port, "accept")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to remove rule.")
			os.Exit(1)
		}
		fmt.Println("Rule removed.")
	},
}

func init() {
	firewallCmd.AddCommand(firewallStatusCmd)
	firewallCmd.AddCommand(firewallEnableCmd)
	firewallCmd.AddCommand(firewallDisableCmd)
	firewallCmd.AddCommand(firewallListCmd)
	firewallCmd.AddCommand(firewallAllowCmd)
	firewallCmd.AddCommand(firewallRemoveCmd)
	rootCmd.AddCommand(firewallCmd)
}
