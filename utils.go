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

// TODO: finish
// TODO: document
// BUG: exit status 1
// Util function - run system audit
func runAuditOnline() {
	const script = `#!/usr/bin/env python
	# -*- coding: utf-8 -*-
	__author__ = 'videns'
	import inspect
	import pkgutil
	import json
	import os
	try:
		import urllib.request as urllib2
	except ImportError:
		import urllib2
	import scanModules


	VULNERS_LINKS = {'pkgChecker':'https://vulners.com/api/v3/audit/audit/',
					 'bulletin':'https://vulners.com/api/v3/search/id/'}

	VULNERS_ASCII = r"""
				 _
	__   ___   _| |_ __   ___ _ __ ___
	\ \ / / | | | | '_ \ / _ \ '__/ __|
	 \ V /| |_| | | | | |  __/ |  \__ \
	  \_/  \__,_|_|_| |_|\___|_|  |___/

	"""


	class scannerEngine():
		def __init__(self):
			self.osInstanceClasses = self.getInstanceClasses()

		def getInstanceClasses(self):
			self.detectors = None
			members = set()
			for modPath, modName, isPkg in pkgutil.iter_modules(scanModules.__path__):
				#find all classed inherited from scanner.osDetect.ScannerInterface in all files
				members = members.union(inspect.getmembers(__import__('%s.%s' % ('scanModules',modName), fromlist=['scanModules']),
											 lambda member:inspect.isclass(member)
														   and issubclass(member, scanModules.osDetect.ScannerInterface)
														   and member.__module__ == '%s.%s' % ('scanModules',modName)
														   and member != scanModules.osDetect.ScannerInterface))
			return members

		def getInstance(self,sshPrefix):
			inited = [instance[1](sshPrefix) for instance in self.osInstanceClasses]
			if not inited:
				raise Exception("No OS Detection classes found")
			osInstance = max(inited, key=lambda x:x.osDetectionWeight)
			if osInstance.osDetectionWeight:
				return osInstance

		def sendVulnRequest(self, url, payload):
			req = urllib2.Request(url)
			req.add_header('Content-Type', 'application/json')
			req.add_header('User-Agent', 'vulners-scanner-v0.1')
			response = urllib2.urlopen(req, json.dumps(payload).encode('utf-8'))
			responseData = response.read()
			if isinstance(responseData, bytes):
				responseData = responseData.decode('utf8')
			responseData = json.loads(responseData)
			return responseData

		def auditSystem(self, sshPrefix, systemInfo=None):
			instance = self.getInstance(sshPrefix)
			installedPackages = instance.getPkg()
			print("="*42)
			if systemInfo:
				print("Host info - %s" % systemInfo)
			print("OS Name - %s, OS Version - %s" % (instance.osFamily, instance.osVersion))
			print("Total found packages: %s" % len(installedPackages))
			if not installedPackages:
				return instance
			# Get vulnerability information
			payload = {'os':instance.osFamily,
					   'version':instance.osVersion,
					   'package':installedPackages}
			url = VULNERS_LINKS.get('pkgChecker')
			response = self.sendVulnRequest(url, payload)
			resultCode = response.get("result")
			if resultCode != "OK":
				print("Error - %s" % response.get('data').get('error'))
			else:
				vulnsFound = response.get('data').get('vulnerabilities')
				if not vulnsFound:
					print("No vulnerabilities found")
				else:
					payload = {'id':vulnsFound}
					allVulnsInfo = self.sendVulnRequest(VULNERS_LINKS['bulletin'], payload)
					vulnInfoFound = allVulnsInfo['result'] == 'OK'
					print("Vulnerable packages:")
					for package in response['data']['packages']:
						print(" "*4 + package)
						packageVulns = []
						for vulns in response['data']['packages'][package]:
							if vulnInfoFound:
								vulnInfo = "{id} - '{title}', cvss.score - {score}".format(id=vulns,
																						   title=allVulnsInfo['data']['documents'][vulns]['title'],
																						   score=allVulnsInfo['data']['documents'][vulns]['cvss']['score'])
								packageVulns.append((vulnInfo,allVulnsInfo['data']['documents'][vulns]['cvss']['score']))
							else:
								packageVulns.append((vulns,0))
						packageVulns = sorted(packageVulns, key=lambda x:x[1])
						packageVulns = [" "*8 + x[0] for x in packageVulns]
						print("\n".join(packageVulns))

			return instance

		def scan(self, checkDocker = False):
			#scan host machine
			hostInstance = self.auditSystem(sshPrefix=None,systemInfo="Host machine")
			#scan dockers
			if checkDocker:
				containers = hostInstance.sshCommand("docker ps")
				if containers:
					containers = containers.splitlines()[1:]
					dockers = [(line.split()[0], line.split()[1]) for line in containers]
					for (dockerID, dockerImage) in dockers:
						sshPrefix = "docker exec %s" % dockerID
						self.auditSystem(sshPrefix, "docker container \"%s\"" % dockerImage)


	if __name__ == "__main__":
		print('\n'.join(VULNERS_ASCII.splitlines()))
		scannerInstance = scannerEngine()
		scannerInstance.scan(checkDocker=False)`

	cmd := exec.Command("python")
	r := strings.NewReader(script)
	var o bytes.Buffer

	cmd.Stdin = r
	cmd.Stdout = &o

	if err := cmd.Run(); err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}

	println(o.String())
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
