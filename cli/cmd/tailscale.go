package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var tailscaleCmd = &cobra.Command{
	Use:   "tailscale",
	Short: "Manage Tailscale (optional)",
}

var tailscaleInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Tailscale",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Tailscale...")

		steps := [][]string{
			{"curl", "-fsSL", "https://tailscale.com/install.sh", "-o", "/tmp/tailscale-install.sh"},
			{"sh", "/tmp/tailscale-install.sh"},
			{"rm", "/tmp/tailscale-install.sh"},
		}

		for _, step := range steps {
			c := exec.Command(step[0], step[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				fmt.Printf("Failed at: %s\n", strings.Join(step, " "))
				os.Exit(1)
			}
		}

		fmt.Println("Tailscale installed.")
		fmt.Println("Run: arivon tailscale connect")
	},
}

var tailscaleStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Tailscale status",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("tailscale", "status").Output()
		if err != nil {
			fmt.Println("Tailscale not installed or not running.")
			return
		}
		fmt.Println(string(out))
	},
}

var tailscaleConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to Tailscale network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting to Tailscale...")
		c := exec.Command("tailscale", "up")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to connect.")
			os.Exit(1)
		}
		fmt.Println("Connected.")
	},
}

var tailscaleDisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from Tailscale network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Disconnecting from Tailscale...")
		c := exec.Command("tailscale", "down")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Failed to disconnect.")
			os.Exit(1)
		}
		fmt.Println("Disconnected.")
	},
}

func init() {
	tailscaleCmd.AddCommand(tailscaleInstallCmd)
	tailscaleCmd.AddCommand(tailscaleStatusCmd)
	tailscaleCmd.AddCommand(tailscaleConnectCmd)
	tailscaleCmd.AddCommand(tailscaleDisconnectCmd)
	rootCmd.AddCommand(tailscaleCmd)
}
