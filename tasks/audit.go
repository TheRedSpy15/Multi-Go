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

// Audit runs multiple checks and reports found security issues to user
// TODO: add more checks
// TODO: add wifi encryption check
// TODO: add something user related checks
// TODO: add current software version checks
// TODO: add using default DNS check (recommend 9.9.9.9 or 1.1.1.1, etc)
// TODO: add antivirus check
// TODO: add guest user check (recommend removing)
// TODO: add auto update check
// TODO: add password policy check
// TODO: (at a later date) add Fail2Ban checks
// TODO: (at a later date) add ssh setting checks
// TODO: document
func Audit() {
	ct.Foreground(ct.Red, true)   // set text color to bright red
	problems := make([]string, 1) // an array to add collection problems to display

	fmt.Println("-- Beginning Audit --")
	fmt.Println("This is a major WIP!")
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	// firewall
	if !strings.Contains(utils.RunCmd("ufw", "status"), "active") { // disabled / is not active
		problems[0] = "Firewall disabled" // add problem
	}
	fmt.Println("Check 1 complete!")

	ct.Foreground(ct.Red, true) // set text color to bright red
	fmt.Println("Problems found:", problems)
}
