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
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// HashFile takes a file path, and then prints the hash of the file
// BUG: won't work unless ran from non-dialog mode / by using commandline flags
func HashFile(target string) {
	utils.CheckTarget(target)       // make sure target is valid
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	file := utils.ReadFileIntoByte(target)                    // get bytes of file to hash
	hash := sha1.New()                                        // create sha1 object
	hash.Write(file)                                          // hash file to object
	target = base64.URLEncoding.EncodeToString(hash.Sum(nil)) // encode hash sum into string

	fmt.Println("SHA-1 hash :", target)
}
