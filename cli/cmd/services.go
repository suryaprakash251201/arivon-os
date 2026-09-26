package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage system services",
}

var servicesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all services",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("systemctl", "list-units", "--type=service", "--no-pager", "--no-legend")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var servicesStatusCmd = &cobra.Command{
	Use:   "status [service]",
	Short: "Show service status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("systemctl", "status", args[0], "--no-pager")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var servicesFailedCmd = &cobra.Command{
	Use:   "failed",
	Short: "Show failed services",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("systemctl", "--failed", "--no-pager", "--no-legend")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	servicesCmd.AddCommand(servicesListCmd)
	servicesCmd.AddCommand(servicesStatusCmd)
	servicesCmd.AddCommand(servicesFailedCmd)
	rootCmd.AddCommand(servicesCmd)
}
