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
	"io"
	"strconv"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// GeneratePassword generated a random string for use as a password
func GeneratePassword(target string, output io.Writer) {
	utils.CheckTarget(target)

	conversion, err := strconv.Atoi(target) // convert target (string), to int

	if err != nil {
		message := fmt.Sprintf("Cannot use \"%s\" as length for password\n", target)
		output.Write([]byte(message))
		return
	}

	message := fmt.Sprintf("Password:%s\n", utils.RandomString(conversion))
	output.Write([]byte(message))
}
