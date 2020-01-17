package tasks

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
	"fmt"

	"github.com/TheRedSpy15/Multi-Go/utils"

	"github.com/daviddengcn/go-colortext"
)

// Clean clears cached files
// TODO run "apt-get autoremove"
// TODO run "apt-get dist-upgrade"
// TODO run "apt-get update && apt-get upgrade"
// TODO clean temporary file locations
func Clean() {
	utils.CheckSudo()

	ct.Foreground(ct.Red, true)
	fmt.Println("This is an EXPERIMENTAL feature!")
	ct.Foreground(ct.Yellow, false)

	// autoclean
	utils.RunCmd("apt-get", "autoclean")
	fmt.Println("Phase 1 complete")

	// check
	utils.RunCmd("apt-get", "check")
	fmt.Println("Phase 2 complete")
}
