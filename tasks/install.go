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
	"os"
	"path"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// Install - will install the current executable to the specified target
// TODO: document
func Install(target string) {
	utils.CheckTarget(target)

	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		err = os.MkdirAll(target, os.ModePerm)
	}
	utils.CheckErr(err)

	appPath, err := os.Executable()
	utils.CheckErr(err)
	srcPath := strings.Replace(appPath, "<nil>", "", 1)
	srcPath = strings.TrimSpace(srcPath)

	if srcPath == "" {
		utils.CheckErr(fmt.Errorf("could not get path of current executable"))
	}

	target = path.Join(target, path.Base(srcPath))
	if _, err := os.Stat(target); err == nil {
		fmt.Println("Target already exists. Overwriting...")

		fmt.Println("Removing old file")
		err = os.Remove(target)
		utils.CheckErr(err)
	}

	fmt.Println("Copying to target")
	err = os.Link(srcPath, target)
	utils.CheckErr(err)

	fmt.Println("Setting execution permission")
	utils.RunCmd("chmod", "+x", target)

	fmt.Println("Done")
}
