package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Manage system updates",
}

var updateCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for available updates",
	Run: func(cmd *cobra.Command, args []string) {
		updates := system.PendingUpdates()
		if updates == 0 {
			fmt.Println("System is up to date.")
		} else {
			fmt.Printf("%d update(s) available.\n", updates)
		}
	},
}

var updateSecurityCmd = &cobra.Command{
	Use:   "security",
	Short: "Install security updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing security updates...")
		upgrade := exec.Command("apt", "upgrade", "-y")
		upgrade.Stdout = os.Stdout
		upgrade.Stderr = os.Stderr
		upgrade.Run()
	},
}

var updateSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Install all updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing all updates...")
		upgrade := exec.Command("apt", "upgrade", "-y")
		upgrade.Stdout = os.Stdout
		upgrade.Stderr = os.Stderr
		upgrade.Run()
	},
}

var updateHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Show update history",
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := exec.Command("apt", "history").Output()
		fmt.Println(string(out))
	},
}

func init() {
	updateCmd.AddCommand(updateCheckCmd)
	updateCmd.AddCommand(updateSecurityCmd)
	updateCmd.AddCommand(updateSystemCmd)
	updateCmd.AddCommand(updateHistoryCmd)
	rootCmd.AddCommand(updateCmd)
}
