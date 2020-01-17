package tasks

import (
	"fmt"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

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

// Bleach securely overwrites target file 3 times with Gutmann
func Bleach(target string) {
	utils.CheckTarget(target)

	fmt.Println("Bleaching")
	utils.RunCmd("shred", "-v", "-z", "-n", "3", target) // overwrites 3 times

	ct.Foreground(ct.Green, true)
	fmt.Println("Done")
}
