package cmd

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
	if rootCmd.Use != "arivon" {
		t.Errorf("Expected root command Use to be 'arivon', got '%s'", rootCmd.Use)
	}
}

func TestInfoCommand(t *testing.T) {
	if infoCmd.Use != "info" {
		t.Errorf("Expected info command Use to be 'info', got '%s'", infoCmd.Use)
	}
}

func TestStatusCommand(t *testing.T) {
	if statusCmd.Use != "status" {
		t.Errorf("Expected status command Use to be 'status', got '%s'", statusCmd.Use)
	}
}

func TestHealthCommand(t *testing.T) {
	if healthCmd.Use != "health" {
		t.Errorf("Expected health command Use to be 'health', got '%s'", healthCmd.Use)
	}
}

func TestDoctorCommand(t *testing.T) {
	if doctorCmd.Use != "doctor" {
		t.Errorf("Expected doctor command Use to be 'doctor', got '%s'", doctorCmd.Use)
	}
}

func TestVersionDefault(t *testing.T) {
	if arivonVersion == "" {
		t.Error("arivonVersion should not be empty")
	}
}

func TestAllSubcommandsRegistered(t *testing.T) {
	expected := []string{
		"info", "status", "health", "doctor", "update", "logs",
		"services", "security", "firewall", "ssh", "docker",
		"network", "storage", "backup", "tailscale", "benchmark",
		"monitoring", "firstboot",
	}

	for _, name := range expected {
		found := false
		for _, cmd := range rootCmd.Commands() {
			if cmd.Name() == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected subcommand '%s' not found", name)
		}
	}
}
