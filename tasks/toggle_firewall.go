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

// ToggleFirewall allows the user to enable/disable system firewall
// TODO: add support for more systems - think only works on debian/ubuntu
func ToggleFirewall(target string) {
	utils.CheckSudo()
	utils.CheckTarget(target)                // make sure target is valid
	fmt.Println(utils.RunCmd("ufw", target)) // run command & print result
}
