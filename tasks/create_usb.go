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
)

// CreateUsb downloads a zip folder with a bunch of tools at a target location
func CreateUsb(target string) {
	utils.CheckTarget(target)

	fmt.Println("Downloading package")
	utils.CheckErr(utils.DownloadFile(target, "nil")) // tool repo not create yet
	fmt.Println("Finished!")
}
