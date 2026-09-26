package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00")).
			MarginBottom(1)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA"))

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00"))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))
)

type disk struct {
	name string
	size string
	model string
}

type installConfig struct {
	disk      disk
	profile   string
	hostname  string
	username  string
	sshPort   string
	docker    bool
	firewall  bool
}

type model struct {
	step      int
	disks     []disk
	diskIndex int
	config    installConfig
	quitting  bool
	err       error
}

func (m model) Init() tea.Cmd {
	return loadDisks
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.step == 1 && m.diskIndex > 0 {
				m.diskIndex--
			}
		case "down", "j":
			if m.step == 1 && m.diskIndex < len(m.disks)-1 {
				m.diskIndex++
			}
		case "enter":
			if m.step < 5 {
				m.step++
			} else {
				m.quitting = true
				return m, tea.Quit
			}
		case " ":
			if m.step == 3 {
				m.config.docker = !m.config.docker
			}
			if m.step == 4 {
				m.config.firewall = !m.config.firewall
			}
		}
	case disksLoadedMsg:
		m.disks = msg.disks
	case installFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.quitting = true
		}
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	b.WriteString(titleStyle.Render("Arivon OS Installer"))
	b.WriteString("\n\n")

	switch m.step {
	case 0:
		b.WriteString("Select installation profile:\n\n")
		profiles := []string{"minimal", "server", "docker", "network", "full"}
		for i, p := range profiles {
			if i == 0 {
				b.WriteString(selectedStyle.Render("  > " + p))
			} else {
				b.WriteString(dimStyle.Render("    " + p))
			}
			b.WriteString("\n")
		}
		b.WriteString("\n")
		b.WriteString(labelStyle.Render("[Enter] Next  [q] Quit"))

	case 1:
		b.WriteString("Select disk:\n\n")
		if len(m.disks) == 0 {
			b.WriteString("  No disks found.\n")
		}
		for i, d := range m.disks {
			prefix := "  "
			if i == m.diskIndex {
				prefix = selectedStyle.Render("> ")
			}
			b.WriteString(fmt.Sprintf("%s%s  %s  %s\n", prefix, d.name, d.size, d.model))
		}
		b.WriteString("\n")
		b.WriteString(labelStyle.Render("[Up/Down] Select  [Enter] Next  [q] Quit"))

	case 2:
		b.WriteString("Configuration:\n\n")
		b.WriteString(fmt.Sprintf("  Hostname: %s\n", m.config.hostname))
		b.WriteString(fmt.Sprintf("  Username: %s\n", m.config.username))
		b.WriteString(fmt.Sprintf("  SSH Port: %s\n", m.config.sshPort))
		b.WriteString("\n")
		b.WriteString(labelStyle.Render("[Enter] Next  [q] Quit"))

	case 3:
		b.WriteString("Docker:\n\n")
		if m.config.docker {
			b.WriteString(selectedStyle.Render("  [X] Install Docker"))
		} else {
			b.WriteString(dimStyle.Render("  [ ] Install Docker"))
		}
		b.WriteString("\n\n")
		b.WriteString(labelStyle.Render("[Space] Toggle  [Enter] Next  [q] Quit"))

	case 4:
		b.WriteString("Firewall:\n\n")
		if m.config.firewall {
			b.WriteString(selectedStyle.Render("  [X] Enable firewall"))
		} else {
			b.WriteString(dimStyle.Render("  [ ] Enable firewall"))
		}
		b.WriteString("\n\n")
		b.WriteString(labelStyle.Render("[Space] Toggle  [Enter] Install  [q] Quit"))

	case 5:
		b.WriteString("Summary:\n\n")
		b.WriteString(fmt.Sprintf("  Disk:     %s\n", m.disks[m.diskIndex].name))
		b.WriteString(fmt.Sprintf("  Profile:  %s\n", m.config.profile))
		b.WriteString(fmt.Sprintf("  Hostname: %s\n", m.config.hostname))
		b.WriteString(fmt.Sprintf("  User:     %s\n", m.config.username))
		b.WriteString(fmt.Sprintf("  SSH:      Port %s\n", m.config.sshPort))
		b.WriteString(fmt.Sprintf("  Docker:   %v\n", m.config.docker))
		b.WriteString(fmt.Sprintf("  Firewall: %v\n", m.config.firewall))
		b.WriteString("\n")
		b.WriteString(labelStyle.Render("[Enter] Install  [q] Quit"))
	}

	if m.err != nil {
		b.WriteString("\n")
		b.WriteString("Error: " + m.err.Error() + "\n")
	}

	return b.String()
}

type disksLoadedMsg struct {
	disks []disk
}

type installFinishedMsg struct {
	err error
}

func loadDisks() tea.Msg {
	out, err := exec.Command("lsblk", "-d", "-n", "-o", "NAME,SIZE,MODEL").Output()
	if err != nil {
		return disksLoadedMsg{disks: []disk{}}
	}

	var disks []disk
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			model := ""
			if len(fields) >= 3 {
				model = strings.Join(fields[2:], " ")
			}
			disks = append(disks, disk{
				name:  "/dev/" + fields[0],
				size:  fields[1],
				model: model,
			})
		}
	}
	return disksLoadedMsg{disks: disks}
}

func runInstall(m model) tea.Msg {
	d := m.disks[m.diskIndex]

	commands := [][]string{
		{"parted", "-s", d.name, "mklabel", "gpt"},
		{"parted", "-s", d.name, "mkpart", "EFI", "fat32", "1MiB", "512MiB"},
		{"parted", "-s", d.name, "set", "1", "esp", "on"},
		{"parted", "-s", d.name, "mkpart", "root", "ext4", "512MiB", "100%"},
		{"mkfs.fat", "-F32", d.name + "1"},
		{"mkfs.ext4", "-F", d.name + "2"},
	}

	for _, cmd := range commands {
		c := exec.Command(cmd[0], cmd[1:]...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return installFinishedMsg{err: fmt.Errorf("%s failed: %w", cmd[0], err)}
		}
	}

	return installFinishedMsg{}
}

func main() {
	m := model{
		step: 0,
		config: installConfig{
			profile:  "minimal",
			hostname: "arivon",
			username: "admin",
			sshPort:  "22",
		},
	}

	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fm := finalModel.(model)
	if fm.err != nil {
		fmt.Printf("Installation failed: %v\n", fm.err)
		os.Exit(1)
	}

	if !fm.quitting {
		return
	}

	fmt.Println("\nInstallation complete. Reboot to start Arivon OS.")
}
