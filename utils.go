package main

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
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/daviddengcn/go-colortext"
	"github.com/gocolly/colly"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"golang.org/x/crypto/ssh/terminal"
)

// Util function - check if target is empty, panic if is
func checkTarget(target string) {
	if target == "" { // check if target is blank
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic("target cannot be empty when performing this task!")
	}
}

// TODO: document
// Run a command on the system & print result
func runCmd(command string, arg ...string) string {
	cmd := exec.Command(command)
	for _, arg := range arg {
		cmd.Args = append(cmd.Args, arg)
	}

	var o bytes.Buffer
	cmd.Stdout = &o // asign o to cmd's Stdout

	if err := cmd.Run(); err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}

	return o.String()
}

// Util function - used for getting []byte of file
func readFileIntoByte(filename string) []byte {
	var data []byte                // specify type
	file, err := os.Open(filename) // make file object
	defer file.Close()             // close file on function end
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	} else {
		data, err = ioutil.ReadAll(file) // read all
		if err != nil {
			ct.Foreground(ct.Red, true) // set text color to bright red
			panic(err.Error())
		}
	}
	return data // return file bytes
}

// Util function - securely get password from user
func getPassword() string {
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin)) // run password command, make var with result
	password := string(bytePassword)                             // cast to string var

	return password
}

// Util function - displays banner text
func printBanner() {
	ct.Foreground(ct.Red, true) // set text color to bright red

	println(`
 __  __       _ _   _    ____
|  \/  |_   _| | |_(_)  / ___| ___
| |\/| | | | | | __| | | |  _ / _ \
| |  | | |_| | | |_| | | |_| | (_) |
|_|  |_|\__,_|_|\__|_|  \____|\___/`)
}

// Util function - scrapes a website link
func collyAddress(target string, savePage bool, ip bool) {
	if ip { // check if target is an IP address not URL
		target = "http://" + target + "/" // modify target to be valid address
	}

	c := colly.NewCollector() // make colly object
	c.IgnoreRobotsTxt = true  // ignore RobotsText

	// configuring colly/collector object
	c.OnRequest(func(r *colly.Request) {
		println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { // print error message on error
		ct.Foreground(ct.Red, true) // set text color to bright red
		log.Println("Something went wrong:", err)
		ct.ResetColor() // reset text color to default
	})

	c.OnResponse(func(r *colly.Response) {
		println("Visited", r.Request.URL)
		println("Response:", r.StatusCode)
	})

	c.OnScraped(func(r *colly.Response) { // finished with site
		println("Finished", r.Request.URL)

		if savePage { // check if save is enabled
			err := r.Save(r.FileName()) // saving data

			if err != nil {
				ct.Foreground(ct.Red, true) // set text color to bright red
				panic("Error saving")
			} else { // saved
				ct.Foreground(ct.Green, true) // set text color to bright red
				println("Saved - ", r.FileName())
				ct.ResetColor() // reset text color to default color
			}
		}
	})

	c.Visit(target) // actually using colly/collector object, and visiting target
}

// TODO: not finished yet
// Util function - constantly sends data to a target
func dos(conn net.Conn) {
	p := make([]byte, 2048)

	defer conn.Close() // make sure to close the connection when done

	println("Starting loop")
	for true { // DOS loop
		fmt.Fprintf(conn, "Sup UDP Server, how you doing?")
		_, err := bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}

		println("looped")
	}
}

// TODO: add more checks
// TODO: add wifi encryption check
// TODO: document
// Audits the system without using third party service
func runAuditOffline() {
	ct.Foreground(ct.Red, true)
	problems := make([]string, 1)

	println("-- Beginning Audit --")
	println("This is a major WIP!\n")
	ct.Foreground(ct.Yellow, false)

	// firewall
	if !strings.Contains(runCmd("ufw", "status"), "active") { // disabled / is not active
		problems[0] = "Firewall disabled"
	}
	println("Check 1 complete!")

	ct.Foreground(ct.Red, true)
	fmt.Println("Problems found:", problems)
}

// TODO: rewrite in my own code
// TODO: add more comments
// Util function - returns a random string
/* Original: https://stackoverflow.com/questions/22892120
/how-to-generate-a-random-string-of-a-fixed-length-in-go#31832326 */
func randomString(length int) string {
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // letters to use
		letterIdxBits = 6                                                      // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits                                     // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano()) // create random source

	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// TODO: add more info - atleast usage
// Util function - prints CPU info
func printCPU() {
	cpuCount, _ := cpu.Counts(false)       // get cpu count total
	cpuCountLogical, _ := cpu.Counts(true) // get cpu logical count
	println("\n-- CPU --\n")
	println("CPU Count: (logical)", cpuCountLogical) // cpu count logical
	println("CPU Count:", cpuCount)                  // cpu count total
}

// TODO: get physical memory instead of swap
// TODO: convert values to gigabytes
// Util function - prints info about system memory
func printMemory() {
	mem, err := mem.SwapMemory() // get virtual memory info object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}
	println("\n-- Memory --\n")
	println("Memory Used:", mem.Used)   // used
	println("Memory Free:", mem.Free)   // free
	println("Memory Total:", mem.Total) // total
}

// Util function - prints info about system host
func printHost() {
	hostInfo, err := host.Info() // get host info object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}
	println("\n-- Host --\n")
	println("Kernal Version:", hostInfo.KernelVersion)     // kernal version
	println("Platform:", hostInfo.Platform)                // platform
	println("Platform Family:", hostInfo.PlatformFamily)   // platform family
	println("Platform Version:", hostInfo.PlatformVersion) // platform version
	println("Uptime:", hostInfo.Uptime)                    // uptime
	println("Host Name:", hostInfo.Hostname)               // hostname
	println("Host ID:", hostInfo.HostID)                   // host id
	println("OS:", hostInfo.OS)                            // os
}
