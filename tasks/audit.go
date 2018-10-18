package tasks

import (
	"fmt"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

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

// TODO: add current software version checks (recommend updating)
// TODO: add using default DNS check (recommend 9.9.9.9 or 1.1.1.1, etc)
// TODO: add antivirus check (recommend getting one)
// TODO: add guest user check (recommend removing)
// TODO: add auto update check (recommend enabling)
// TODO: add password policy check
// TODO: add vpn check (recommend using one)
// TODO: (at a later date) add Fail2Ban checks
// TODO: (at a later date) add ssh setting checks
// TODO: use an iterator or a variable to add problems to the array
// Runs several security checks, then prints found vulnerabilites
func Audit() {
	utils.CheckSudo()

	ct.Foreground(ct.Red, true)   // set text color to bright red
	problems := make([]string, 2) // an array to add collection problems to display

	fmt.Println("-- Beginning Audit --")
	fmt.Println("This is a major WIP!")
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	// firewall
	if !strings.Contains(utils.RunCmd("ufw", "status"), "active") { // disabled / is not active
		problems[0] = "Firewall disabled" // add problem
	}
	fmt.Println("Check 1 complete!")

	// network connection type
	if strings.Contains(utils.RunCmd("nmcli", "d"), "wifi") { // using wifi
		problems[1] = "Using wifi instead of ethernet" // add problem

		fmt.Println("Check 2 complete!")

		// encrypted wifi
		if !strings.Contains(utils.RunCmd("nmcli", "-t", "-f", "active,ssid", "dev", "wifi"), "yes") {
			problems[2] = "Using insecure wifi"
		}
		fmt.Println("Check 3 complete!")
	} else {
		fmt.Println("Check 2 complete!")
	}

	ct.Foreground(ct.Red, true) // set text color to bright red
	fmt.Println("Problems found:", problems)
}
