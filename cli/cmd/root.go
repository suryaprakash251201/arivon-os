package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "arivon",
	Short: "Arivon OS — Intelligent Infrastructure",
	Long:  "Arivon OS system administration CLI.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
