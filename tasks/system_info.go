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
	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// SystemInfo prints extensive info about system
// TODO: make & add more info functions
func SystemInfo() {
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow
	utils.PrintCPU()                // print cpu info
	utils.PrintMemory()             // print memory info
	utils.PrintHost()               // print host info
}
