package cmd

import (
	"fmt"
	"os"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var firstbootCmd = &cobra.Command{
	Use:    "firstboot",
	Short:  "Display first boot welcome message",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("Welcome to Arivon OS")
		fmt.Println()
		fmt.Printf("  Version    %s\n", arivonVersion)
		fmt.Printf("  Hostname   %s\n", getHostname())
		fmt.Printf("  IP         %s\n", system.IPAddress())
		fmt.Printf("  CPU        %d cores\n", system.CPUCount())
		fmt.Printf("  RAM        %.1f GB\n", float64(system.Memory().Total)/1e9)
		fmt.Printf("  Kernel     %s\n", system.KernelVersion())
		fmt.Println()
		fmt.Println("  SSH:")
		fmt.Printf("  ssh admin@%s\n", system.IPAddress())
		fmt.Println()
		fmt.Println("  Run:")
		fmt.Println("  arivon status")
		fmt.Println("  arivon health")
		fmt.Println("  arivon security audit")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(firstbootCmd)
}

func getHostname() string {
	h, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return h
}
