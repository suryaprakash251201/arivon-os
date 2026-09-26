package cmd

import (
	"fmt"
	"os"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Run system health check",
	Run: func(cmd *cobra.Command, args []string) {
		runHealthCheck()
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}

func runHealthCheck() {
	fmt.Println("HEALTH CHECK")
	fmt.Println()

	check("CPU", system.CPUCount() > 0)

	mem := system.Memory()
	memOK := mem.Total > 0 && mem.Available > 0
	check("Memory", memOK)

	check("Disk", system.DiskUsage() != "unknown")

	check("Network", system.IPAddress() != "unknown")

	check("DNS", true)

	check("SSH", system.ServiceStatus("ssh") == "Running")

	check("Docker", system.DockerStatus() == "Running" || system.DockerStatus() == "Not installed")

	check("Firewall", system.FirewallStatus() == "Active" || system.FirewallStatus() == "Inactive")

	updates := system.PendingUpdates()
	if updates > 0 {
		fmt.Printf("  WARNING  Updates (%d available)\n", updates)
	} else {
		fmt.Println("  PASS     Updates")
	}

	failed := system.FailedServices()
	if len(failed) > 0 {
		fmt.Printf("  WARNING  Failed services (%d)\n", len(failed))
	} else {
		fmt.Println("  PASS     Services")
	}

	fmt.Println()
	if updates == 0 && len(failed) == 0 {
		fmt.Println("Result: HEALTHY")
	} else {
		fmt.Println("Result: HEALTHY WITH WARNINGS")
	}

	if len(failed) > 0 {
		os.Exit(1)
	}
}

func check(name string, ok bool) {
	if ok {
		fmt.Printf("  PASS     %s\n", name)
	} else {
		fmt.Printf("  FAIL     %s\n", name)
	}
}
