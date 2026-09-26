package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup management",
}

var backupStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show backup status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BACKUP STATUS")
		fmt.Println()

		out, _ := exec.Command("systemctl", "is-active", "restic-backup.timer").Output()
		status := strings.TrimSpace(string(out))
		if status == "active" {
			fmt.Println("Timer:   Active")
		} else {
			fmt.Println("Timer:   Inactive")
		}

		out, _ = exec.Command("systemctl", "list-timers", "restic-backup.timer", "--no-pager").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, line := range lines {
			if strings.Contains(line, "restic-backup") {
				fmt.Printf("Schedule: %s\n", strings.TrimSpace(line))
			}
		}

		repo := os.Getenv("RESTIC_REPOSITORY")
		if repo == "" {
			repo = "/var/backups/arivon"
		}
		fmt.Printf("Repo:    %s\n", repo)

		out, err := exec.Command("restic", "-r", repo, "snapshots", "--last", "--no-lock").Output()
		if err != nil {
			fmt.Println("Last:    No snapshots found")
		} else {
			snapshots := strings.Split(strings.TrimSpace(string(out)), "\n")
			if len(snapshots) > 1 {
				fmt.Printf("Last:    %s\n", snapshots[1])
			}
		}
	},
}

var backupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a backup",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating backup...")

		repo := os.Getenv("RESTIC_REPOSITORY")
		if repo == "" {
			repo = "/var/backups/arivon"
		}

		c := exec.Command("restic", "-r", repo, "backup", "/etc", "/home", "/var/lib/docker/volumes")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Backup failed.")
			os.Exit(1)
		}

		fmt.Println("Backup complete.")
	},
}

var backupRestoreCmd = &cobra.Command{
	Use:   "restore [snapshot]",
	Short: "Restore from backup",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		snapshot := args[0]
		fmt.Printf("Restoring snapshot: %s\n", snapshot)
		fmt.Println("WARNING: This will overwrite existing files.")
		fmt.Print("Continue? [y/N]: ")

		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			fmt.Println("Cancelled.")
			return
		}

		repo := os.Getenv("RESTIC_REPOSITORY")
		if repo == "" {
			repo = "/var/backups/arivon"
		}

		c := exec.Command("restic", "-r", repo, "restore", snapshot, "--target", "/")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Println("Restore failed.")
			os.Exit(1)
		}

		fmt.Println("Restore complete.")
	},
}

func init() {
	backupCmd.AddCommand(backupStatusCmd)
	backupCmd.AddCommand(backupCreateCmd)
	backupCmd.AddCommand(backupRestoreCmd)
	rootCmd.AddCommand(backupCmd)
}
