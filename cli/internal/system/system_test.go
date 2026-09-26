package system

import (
	"runtime"
	"testing"
)

func TestOSRelease(t *testing.T) {
	rel := OSRelease()
	if rel == nil {
		t.Fatal("OSRelease returned nil")
	}
}

func TestKernelVersion(t *testing.T) {
	kv := KernelVersion()
	if kv == "unknown" && runtime.GOOS == "linux" {
		t.Error("KernelVersion returned unknown on Linux")
	}
}

func TestArchitecture(t *testing.T) {
	arch := Architecture()
	if arch == "" {
		t.Error("Architecture returned empty string")
	}
}

func TestMemory(t *testing.T) {
	mem := Memory()
	if runtime.GOOS == "linux" && mem.Total == 0 {
		t.Error("Memory returned zero total on Linux")
	}
}

func TestCPUCount(t *testing.T) {
	count := CPUCount()
	if count == 0 {
		t.Error("CPUCount returned zero")
	}
}

func TestUptime(t *testing.T) {
	uptime := Uptime()
	if uptime == "unknown" && runtime.GOOS == "linux" {
		t.Error("Uptime returned unknown on Linux")
	}
}

func TestLoadAverage(t *testing.T) {
	load := LoadAverage()
	if load == "unknown" && runtime.GOOS == "linux" {
		t.Error("LoadAverage returned unknown on Linux")
	}
}
