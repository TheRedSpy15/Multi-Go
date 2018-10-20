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
	"os"
	"strings"

	"github.com/daviddengcn/go-colortext"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// Installer - will either add the ability to easily call Multi-Go, or remove it
// TODO: not working yet
// TODO: make less error prone
func Installer(target string) {
	ct.Foreground(ct.Red, true)
	fmt.Println("This is a major WIP!") // warning
	ct.ResetColor()

	utils.CheckTarget(target)

	if target == "install" { // install
		if _, err := os.Stat("/bin"); !os.IsNotExist(err) { // bin already exists
			utils.CheckErr(err)
			fmt.Println("bin exists")

			// get path of current Multi-Go program to copy
			fmt.Println("Getting application path")
			appPath, _ := os.Executable()
			srcPath := strings.Replace(appPath, "<nil>", "", 1)
			strings.TrimSpace(srcPath)

			// copy to bin
			fmt.Println("Copying to bin")
			err := os.Link(srcPath, "/bin")
			utils.CheckErr(err)

			// execution permisssions
			fmt.Println("Getting execution permission")
			utils.RunCmd("chmod +x $HOME/bin/Multi-Go")

			fmt.Println("Done")
		} // need to create bin
	} // need to add uninstall option
}
