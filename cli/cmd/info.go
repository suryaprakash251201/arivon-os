package cmd

import (
	"fmt"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Arivon OS identity information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ARIVON OS")
		fmt.Println()
		fmt.Printf("  Version       %s\n", arivonVersion)
		fmt.Printf("  Upstream      Debian %s\n", system.DebianVersion())
		fmt.Printf("  Kernel        %s\n", system.KernelVersion())
		fmt.Printf("  Architecture  %s\n", system.Architecture())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
