package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Storage management",
}

var storageStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show storage status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("STORAGE STATUS")
		fmt.Println()

		out, _ := exec.Command("df", "-h", "/").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) >= 2 {
			fields := strings.Fields(lines[1])
			if len(fields) >= 5 {
				fmt.Printf("Root filesystem: %s / %s (%s used)\n", fields[2], fields[1], fields[4])
			}
		}

		out, _ = exec.Command("lsblk", "-o", "NAME,SIZE,TYPE,MOUNTPOINT").Output()
		fmt.Println()
		fmt.Println("Block devices:")
		fmt.Println(string(out))
	},
}

var storageDisksCmd = &cobra.Command{
	Use:   "disks",
	Short: "List disks",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("lsblk", "-d", "-o", "NAME,SIZE,MODEL,ROTA")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var storageMountsCmd = &cobra.Command{
	Use:   "mounts",
	Short: "Show mount points",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("findmnt", "-t", "ext4,xfs,btrfs,vfat")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var storageHealthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check filesystem health",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("STORAGE HEALTH")
		fmt.Println()

		out, _ := exec.Command("df", "-h", "/").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) >= 2 {
			fields := strings.Fields(lines[1])
			if len(fields) >= 5 {
				pct := strings.TrimSuffix(fields[4], "%")
				if pct == "90" || pct == "95" || pct == "100" {
					fmt.Printf("  WARNING  Root filesystem %s full\n", fields[4])
				} else {
					fmt.Printf("  PASS     Root filesystem %s used\n", fields[4])
				}
			}
		}

		out, _ = exec.Command("cat", "/proc/mounts").Output()
		for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			if strings.Contains(line, "errors=") && !strings.Contains(line, "errors=remount-ro") {
				fmt.Printf("  WARNING  Filesystem errors: %s\n", line)
			}
		}

		fmt.Println()
		fmt.Println("Storage health check complete.")
	},
}

var storageSmartCmd = &cobra.Command{
	Use:   "smart",
	Short: "Show SMART disk health",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("lsblk", "-d", "-n", "-o", "NAME").Output()
		if err != nil {
			fmt.Println("No disks found.")
			return
		}

		disks := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, disk := range disks {
			disk = strings.TrimSpace(disk)
			if disk == "" {
				continue
			}
			fmt.Printf("Disk: /dev/%s\n", disk)
			smart := exec.Command("smartctl", "-H", "/dev/"+disk)
			smart.Stdout = os.Stdout
			smart.Stderr = os.Stderr
			if err := smart.Run(); err != nil {
				fmt.Printf("  SMART not available for /dev/%s\n", disk)
			}
			fmt.Println()
		}
	},
}

func init() {
	storageCmd.AddCommand(storageStatusCmd)
	storageCmd.AddCommand(storageDisksCmd)
	storageCmd.AddCommand(storageMountsCmd)
	storageCmd.AddCommand(storageHealthCmd)
	storageCmd.AddCommand(storageSmartCmd)
	rootCmd.AddCommand(storageCmd)
}
