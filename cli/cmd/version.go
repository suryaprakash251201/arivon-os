package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var arivonVersion = "0.1.0-dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Arivon OS version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Arivon OS %s\n", arivonVersion)
		fmt.Printf("  Go:        %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:   %s/%s\n", runtime.GOOS, runtime.GOARCH)

		if debianVer := getDebianVersion(); debianVer != "" {
			fmt.Printf("  Debian:    %s\n", debianVer)
		}

		if kernelVer := getKernelVersion(); kernelVer != "" {
			fmt.Printf("  Kernel:    %s\n", kernelVer)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getDebianVersion() string {
	data, err := os.ReadFile("/etc/debian_version")
	if err != nil {
		return ""
	}
	return string(data)
}

func getKernelVersion() string {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return ""
	}
	return string(data)
}
