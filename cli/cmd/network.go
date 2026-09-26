package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network management",
}

var networkStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show network status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NETWORK STATUS")
		fmt.Println()

		out, _ := exec.Command("ip", "-4", "addr", "show").Output()
		fmt.Println("Interfaces:")
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, line := range lines {
			if strings.Contains(line, "inet ") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					fmt.Printf("  %s\n", strings.TrimSpace(line))
				}
			}
		}
		fmt.Println()

		out, _ = exec.Command("ip", "route", "show").Output()
		fmt.Println("Routes:")
		fmt.Println(string(out))

		out, _ = exec.Command("cat", "/etc/resolv.conf").Output()
		fmt.Println("DNS:")
		for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			if strings.HasPrefix(line, "nameserver") {
				fmt.Printf("  %s\n", strings.TrimPrefix(line, "nameserver "))
			}
		}
	},
}

var networkInterfacesCmd = &cobra.Command{
	Use:   "interfaces",
	Short: "List network interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("ip", "-o", "link", "show")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var networkRoutesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Show routing table",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("ip", "route", "show")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var networkDNSCmd = &cobra.Command{
	Use:   "dns",
	Short: "Show DNS configuration",
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := exec.Command("cat", "/etc/resolv.conf").Output()
		fmt.Print(string(out))
	},
}

var networkTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run network diagnostics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NETWORK DIAGNOSTICS")
		fmt.Println()

		fmt.Println("[Ping gateway]")
		out, _ := exec.Command("ip", "route").Output()
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) > 0 {
			fields := strings.Fields(lines[0])
			if len(fields) >= 3 {
				gateway := fields[2]
				fmt.Printf("Gateway: %s\n", gateway)
				ping := exec.Command("ping", "-c", "3", "-W", "2", gateway)
				ping.Stdout = os.Stdout
				ping.Stderr = os.Stderr
				if err := ping.Run(); err != nil {
					fmt.Println("Gateway unreachable.")
				}
			}
		}
		fmt.Println()

		fmt.Println("[DNS lookup]")
		dns := exec.Command("nslookup", "google.com")
		dns.Stdout = os.Stdout
		dns.Stderr = os.Stderr
		if err := dns.Run(); err != nil {
			fmt.Println("DNS lookup failed.")
		}
		fmt.Println()

		fmt.Println("[Socket inspection]")
		sockets := exec.Command("ss", "-tlnp")
		sockets.Stdout = os.Stdout
		sockets.Stderr = os.Stderr
		sockets.Run()
	},
}

func init() {
	networkCmd.AddCommand(networkStatusCmd)
	networkCmd.AddCommand(networkInterfacesCmd)
	networkCmd.AddCommand(networkRoutesCmd)
	networkCmd.AddCommand(networkDNSCmd)
	networkCmd.AddCommand(networkTestCmd)
	rootCmd.AddCommand(networkCmd)
}
