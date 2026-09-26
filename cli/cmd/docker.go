package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage Docker",
}

var dockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Docker Engine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Docker Engine...")

		steps := [][]string{
			{"apt", "update"},
			{"apt", "install", "-y", "ca-certificates", "curl", "gnupg"},
			{"install", "-m", "0755", "-d", "/etc/apt/keyrings"},
			{"curl", "-fsSL", "https://download.docker.com/linux/debian/gpg", "-o", "/etc/apt/keyrings/docker.asc"},
			{"chmod", "a+r", "/etc/apt/keyrings/docker.asc"},
		}

		for _, step := range steps {
			c := exec.Command(step[0], step[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				fmt.Printf("Failed at: %s\n", strings.Join(step, " "))
				os.Exit(1)
			}
		}

		arch := "amd64"
		if _, err := exec.Command("uname", "-m").Output(); err == nil {
			out, _ := exec.Command("uname", "-m").Output()
			if strings.TrimSpace(string(out)) == "aarch64" {
				arch = "arm64"
			}
		}

		repoCmd := fmt.Sprintf("echo \"deb [arch=%s signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian $(. /etc/os-release && echo $VERSION_CODENAME) stable\" > /etc/apt/sources.list.d/docker.list", arch)
		c := exec.Command("bash", "-c", repoCmd)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()

		installSteps := [][]string{
			{"apt", "update"},
			{"apt", "install", "-y", "docker-ce", "docker-ce-cli", "containerd.io", "docker-buildx-plugin", "docker-compose-plugin"},
		}

		for _, step := range installSteps {
			c := exec.Command(step[0], step[1:]...)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				fmt.Printf("Failed at: %s\n", strings.Join(step, " "))
				os.Exit(1)
			}
		}

		exec.Command("systemctl", "enable", "docker").Run()
		exec.Command("systemctl", "start", "docker").Run()

		fmt.Println("Docker installed successfully.")
	},
}

var dockerStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Docker status",
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := exec.Command("systemctl", "is-active", "docker").Output()
		status := strings.TrimSpace(string(out))
		if status == "active" {
			fmt.Println("Docker: Running")
		} else {
			fmt.Println("Docker: Stopped")
		}

		out, err := exec.Command("docker", "version", "--format", "Server: {{.Server.Version}}").Output()
		if err == nil {
			fmt.Printf("Version: %s\n", strings.TrimSpace(string(out)))
		}

		out, err = exec.Command("docker", "info", "--format", "Containers: {{.Containers}}").Output()
		if err == nil {
			fmt.Printf("%s\n", strings.TrimSpace(string(out)))
		}
	},
}

var dockerPsCmd = &cobra.Command{
	Use:   "ps",
	Short: "List Docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker", "ps", "--format", "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t{{.Ports}}")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var dockerImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List Docker images",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker", "images", "--format", "table {{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.Size}}")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var dockerVolumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "List Docker volumes",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker", "volume", "ls")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var dockerNetworksCmd = &cobra.Command{
	Use:   "networks",
	Short: "List Docker networks",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("docker", "network", "ls")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

var dockerAuditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit Docker security",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DOCKER SECURITY AUDIT")
		fmt.Println()

		out, err := exec.Command("docker", "ps", "--format", "{{.ID}} {{.Names}} {{.Image}}").Output()
		if err != nil {
			fmt.Println("Docker not available.")
			return
		}

		containers := strings.TrimSpace(string(out))
		if containers == "" {
			fmt.Println("No running containers.")
			return
		}

		lines := strings.Split(containers, "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			id := fields[0]
			name := fields[1]

			fmt.Printf("Container: %s\n", name)

			out, _ := exec.Command("docker", "inspect", "--format", "{{.HostConfig.Privileged}}", id).Output()
			if strings.TrimSpace(string(out)) == "true" {
				fmt.Println("  WARNING  Privileged container")
			}

			out, _ = exec.Command("docker", "inspect", "--format", "{{.HostConfig.NetworkMode}}", id).Output()
			if strings.TrimSpace(string(out)) == "host" {
				fmt.Println("  WARNING  Host networking")
			}

			out, _ = exec.Command("docker", "inspect", "--format", "{{.HostConfig.PidMode}}", id).Output()
			if strings.TrimSpace(string(out)) == "host" {
				fmt.Println("  WARNING  Host PID namespace")
			}

			out, _ = exec.Command("docker", "inspect", "--format", "{{json .HostConfig.Binds}}", id).Output()
			if strings.Contains(string(out), "/var/run/docker.sock") {
				fmt.Println("  WARNING  Docker socket mounted")
			}

			out, _ = exec.Command("docker", "inspect", "--format", "{{json .HostConfig.PortBindings}}", id).Output()
			portBindings := strings.TrimSpace(string(out))
			if portBindings != "null" && portBindings != "{}" {
				fmt.Printf("  INFO     Ports exposed: %s\n", portBindings)
			}

			fmt.Println()
		}
	},
}

var dockerCleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up unused Docker resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This will remove unused Docker resources.")
		fmt.Print("Continue? [y/N]: ")

		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			fmt.Println("Cancelled.")
			return
		}

		fmt.Println("Removing unused containers...")
		exec.Command("docker", "container", "prune", "-f").Run()

		fmt.Println("Removing unused images...")
		exec.Command("docker", "image", "prune", "-f").Run()

		fmt.Println("Removing unused volumes...")
		exec.Command("docker", "volume", "prune", "-f").Run()

		fmt.Println("Removing unused networks...")
		exec.Command("docker", "network", "prune", "-f").Run()

		fmt.Println("Cleanup complete.")
	},
}

func init() {
	dockerCmd.AddCommand(dockerInstallCmd)
	dockerCmd.AddCommand(dockerStatusCmd)
	dockerCmd.AddCommand(dockerPsCmd)
	dockerCmd.AddCommand(dockerImagesCmd)
	dockerCmd.AddCommand(dockerVolumesCmd)
	dockerCmd.AddCommand(dockerNetworksCmd)
	dockerCmd.AddCommand(dockerAuditCmd)
	dockerCmd.AddCommand(dockerCleanupCmd)
	rootCmd.AddCommand(dockerCmd)
}
