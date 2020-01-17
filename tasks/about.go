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

// About prints details about the program
func About() {
	utils.PrintBanner()

	ct.Foreground(ct.Yellow, false)
	fmt.Println("Multi Go v0.6.1", "\nBy: TheRedSpy15")
	fmt.Println("GitHub:", "https://github.com/TheRedSpy15")
	fmt.Println("Project Page:", "https://github.com/TheRedSpy15/Multi-Go")
	fmt.Println("\nMulti Go allows IT admins and Cyber Security experts")
	fmt.Println("to conveniently perform all sorts of tasks.")
}
