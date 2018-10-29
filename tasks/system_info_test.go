package tasks_test

import (
	"testing"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func CPUTest(t *testing.T) {
	_, err1 := cpu.Counts(false)       // get cpu count total
	if err1 != nil {
		t.Fatal(err1)
	}
	_, err2 := cpu.Counts(true) // get cpu logical count
	if err2 != nil {
		t.Fatal(err2)
	}
}

func MemoryTest(t *testing.T) {
	_, err := mem.SwapMemory() // get virtual memory info object
	if err != nil {
		t.Fatal(err)
	}
}

func HostTest(t *testing.T) {
	_, err := host.Info() // get host info object
	if err != nil {
		t.Fatal(err)
	}
}
