package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/arivon/arivon-os/cli/internal/system"
	"github.com/spf13/cobra"
)

var statusJSON bool

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Arivon OS system status",
	Run: func(cmd *cobra.Command, args []string) {
		if statusJSON {
			printStatusJSON()
		} else {
			printStatus()
		}
	},
}

func init() {
	statusCmd.Flags().BoolVar(&statusJSON, "json", false, "Output in JSON format")
	rootCmd.AddCommand(statusCmd)
}

func printStatus() {
	mem := system.Memory()
	memUsed := mem.Total - mem.Available
	swapUsed := mem.SwapTotal - mem.SwapFree

	fmt.Println("ARIVON OS")
	fmt.Println()
	fmt.Println("System")
	fmt.Printf("  Version       %s\n", arivonVersion)
	fmt.Printf("  Kernel        %s\n", system.KernelVersion())
	fmt.Printf("  Architecture  %s\n", system.Architecture())
	fmt.Printf("  Uptime        %s\n", system.Uptime())
	fmt.Println()
	fmt.Println("Resources")
	fmt.Printf("  CPU           %d cores\n", system.CPUCount())
	fmt.Printf("  RAM           %.1f / %.1f GB\n", float64(memUsed)/1e9, float64(mem.Total)/1e9)
	fmt.Printf("  Swap          %.0f MB\n", float64(swapUsed)/1e6)
	fmt.Printf("  Disk          %s\n", system.DiskUsage())
	fmt.Printf("  Load          %s\n", system.LoadAverage())
	fmt.Println()
	fmt.Println("Services")
	fmt.Printf("  SSH           %s\n", system.ServiceStatus("ssh"))
	fmt.Printf("  Docker        %s\n", system.DockerStatus())
	fmt.Printf("  Firewall      %s\n", system.FirewallStatus())
	fmt.Println()
	fmt.Println("Security")
	updates := system.PendingUpdates()
	if updates > 0 {
		fmt.Printf("  Updates       %d available\n", updates)
	} else {
		fmt.Println("  Updates       None")
	}
	fmt.Println("  Reboot        Not required")
	fmt.Println()
	fmt.Println("Overall")

	healthy := updates == 0 && len(system.FailedServices()) == 0
	if healthy {
		fmt.Println("  Healthy")
	} else {
		fmt.Println("  Degraded")
	}
}

type statusOutput struct {
	System struct {
		Version      string `json:"version"`
		Kernel       string `json:"kernel"`
		Architecture string `json:"architecture"`
		Uptime       string `json:"uptime"`
	} `json:"system"`
	Resources struct {
		CPU      int    `json:"cpu_cores"`
		RAMUsed  string `json:"ram_used"`
		RAMTotal string `json:"ram_total"`
		SwapUsed string `json:"swap_used"`
		Disk     string `json:"disk"`
		Load     string `json:"load"`
	} `json:"resources"`
	Services struct {
		SSH      string `json:"ssh"`
		Docker   string `json:"docker"`
		Firewall string `json:"firewall"`
	} `json:"services"`
	Security struct {
		Updates int  `json:"updates_available"`
		Reboot  bool `json:"reboot_required"`
	} `json:"security"`
	Overall string `json:"overall"`
}

func printStatusJSON() {
	mem := system.Memory()
	memUsed := mem.Total - mem.Available
	swapUsed := mem.SwapTotal - mem.SwapFree

	var s statusOutput
	s.System.Version = arivonVersion
	s.System.Kernel = system.KernelVersion()
	s.System.Architecture = system.Architecture()
	s.System.Uptime = system.Uptime()
	s.Resources.CPU = system.CPUCount()
	s.Resources.RAMUsed = fmt.Sprintf("%.1f GB", float64(memUsed)/1e9)
	s.Resources.RAMTotal = fmt.Sprintf("%.1f GB", float64(mem.Total)/1e9)
	s.Resources.SwapUsed = fmt.Sprintf("%.0f MB", float64(swapUsed)/1e6)
	s.Resources.Disk = system.DiskUsage()
	s.Resources.Load = system.LoadAverage()
	s.Services.SSH = system.ServiceStatus("ssh")
	s.Services.Docker = system.DockerStatus()
	s.Services.Firewall = system.FirewallStatus()
	s.Security.Updates = system.PendingUpdates()

	healthy := s.Security.Updates == 0 && len(system.FailedServices()) == 0
	if healthy {
		s.Overall = "healthy"
	} else {
		s.Overall = "degraded"
	}

	data, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(data))
}
