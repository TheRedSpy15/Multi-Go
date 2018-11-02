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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// Audit - Runs several security checks, then prints found vulnerabilites
// TODO add current software version checks
// BODY and recommend updating
// TODO add using default DNS check
// BODY and recommend using 1.1.1.1, 9.9.9.9, or some other great DNS
// TODO add antivirus check
// BODY and recommend getting something like Clamtk
// TODO add guest user check
// BODY recommend disabling
// TODO add auto update check
// BODY recommend enabling
// TODO add password policy check
// BODY recommend minimum of 8 char length password, etc
// TODO add empty trash check
// BODY recommend emptying
// TODO add vpn check BODY recommend the user getting one
// BODY recommend nordvpn, PIA, tunnelbear, etc
// TODO add microphone check (recommend disabling)
// BODY recommend preferably a command to disable
// TODO add camera check (recommend disabling or covering)
// BODY recommend preferably a command to disable
// TODO trim white space & brackets in solution & problem arrays
// BODY More & more important as more checks are added
func Audit() {
	utils.CheckSudo()

	ct.Foreground(ct.Red, true)
	problems := make([]string, 4)  // an array to add collection of problems
	solutions := make([]string, 4) // an array to add collection of solutions to problems
	check := 0                     // used to increment value when printing check complete

	// banner
	fmt.Println("-- Beginning Audit --")
	fmt.Println("This is a major WIP!")
	ct.Foreground(ct.Yellow, false)

	// firewall
	if !strings.Contains(utils.RunCmd("ufw", "status"), "active") { // disabled / is not active
		problems[check] = "Firewall disabled"
		solutions[check] = "Enable firewall"
	}
	check++
	fmt.Println("Check", check, "complete!")

	// network connection type
	if strings.Contains(utils.RunCmd("nmcli", "d"), "wifi") { // using wifi
		problems[check] = "Using wifi instead of ethernet"
		solutions[check] = "Switch to ethernet"

		check++
		fmt.Println("Check", check, "complete!")

		// encrypted wifi
		if !strings.Contains(utils.RunCmd("nmcli", "-t", "-f", "active,ssid", "dev", "wifi"), "yes") { // not secure
			problems[check] = "Using insecure wifi"
			solutions[check] = "Use a VPN or switch to more secure wifi"
		}
		check++
		fmt.Println("Check", check, "complete!")
	} else { // skip over encrypt wifi check - not using wifi
		check++
		fmt.Println("Check", check, "complete!")
	}

	// guest account
	// TODO: not finished - can't find file
	if _, err := os.Stat("/etc/lightdm/lightdm.conf"); !os.IsNotExist(err) { // look for proper conf file
		file, err := os.Open("/etc/lightdm/lightdm.conf") // open file
		utils.CheckErr(err)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() { // scan loop
			if !strings.Contains(scanner.Text(), "allow-guest=false") { // look guest disable
				problems[check] = "Guest access is still enabled"
				solutions[check] = `Add "allow-guest=false" to lightdm.conf file`
			}
		}

		check++
		fmt.Println("Check", check, "complete!")
	} else { // file not found
		check++
		ct.Foreground(ct.Red, true)
		fmt.Println("Check", check, "FAILED!")
		ct.Foreground(ct.Yellow, false)
	}

	// bash history
	// TODO: not finished - can't find file
	if _, err := os.Stat("/.bash_history"); !os.IsNotExist(err) {
		file, err := os.Open("/.bash_history") // open file
		utils.CheckErr(err)

		fileStat, _ := file.Stat() // get file info
		fmt.Println(fileStat.Size())
		if fileStat.Size() >= 100 { // check file size
			problems[check] = "Command line history found"
			solutions[check] = `Run: "cat /dev/null > ~/.bash_history && history -c && exit"`
		}

		check++
		fmt.Println("Check", check, "complete!")
	} else { // file not found
		check++
		ct.Foreground(ct.Red, true)
		fmt.Println("Check", check, "FAILED!")
		ct.Foreground(ct.Yellow, false)
	}

	ct.Foreground(ct.Red, true)
	fmt.Println("Problems found:", problems)

	ct.Foreground(ct.Green, true)
	fmt.Println("Solutions:", solutions)
}
