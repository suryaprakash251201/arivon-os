package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "View system logs",
}

var logsDefaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Show default journal logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("journalctl", "--no-pager", "-n", "50")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var logsSSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Show SSH logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("journalctl", "-u", "ssh", "--no-pager", "-n", "50")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var logsDockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Show Docker logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("journalctl", "-u", "docker", "--no-pager", "-n", "50")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var logsKernelCmd = &cobra.Command{
	Use:   "kernel",
	Short: "Show kernel logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("journalctl", "-k", "--no-pager", "-n", "50")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var logsBootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Show boot logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("journalctl", "-b", "--no-pager", "-n", "50")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	logsCmd.AddCommand(logsDefaultCmd)
	logsCmd.AddCommand(logsSSHCmd)
	logsCmd.AddCommand(logsDockerCmd)
	logsCmd.AddCommand(logsKernelCmd)
	logsCmd.AddCommand(logsBootCmd)
	rootCmd.AddCommand(logsCmd)
}
