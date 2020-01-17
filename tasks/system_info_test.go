package tasks_test

/*
   Copyright 2020 TheRedSpy15

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"testing"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func CPUTest(t *testing.T) {
	_, err1 := cpu.Counts(false) // get cpu count total
	if err1 != nil {
		t.Fatal(err1)
	}
	_, err2 := cpu.Counts(true) // get cpu logical count
	if err2 != nil {
		t.Fatal(err2)
	}
}

func MemoryTest(t *testing.T) {
	_, err := mem.VirtualMemory() // get virtual memory info object
	if err != nil {
		t.Fatal(err)
	}
}

func SwapTest(t *testing.T) {
	_, err := mem.SwapMemory() // get swap memory object
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
