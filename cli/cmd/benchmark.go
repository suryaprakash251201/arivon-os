package cmd

import (
	"fmt"
	"time"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Run system benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ARIVON PERFORMANCE")
		fmt.Println()

		start := time.Now()
		_ = system.Memory()
		_ = system.CPUCount()
		_ = system.KernelVersion()
		_ = system.Uptime()
		_ = system.LoadAverage()
		elapsed := time.Since(start)

		fmt.Printf("Info gathering: %v\n", elapsed)
		fmt.Printf("CPU:           %d cores\n", system.CPUCount())
		fmt.Printf("RAM:           %.1f GB\n", float64(system.Memory().Total)/1e9)
		fmt.Printf("Kernel:        %s\n", system.KernelVersion())
		fmt.Printf("Uptime:        %s\n", system.Uptime())
		fmt.Printf("Load:          %s\n", system.LoadAverage())
	},
}

func init() {
	rootCmd.AddCommand(benchmarkCmd)
}
