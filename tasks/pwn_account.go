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
	"io/ioutil"
	"net/http"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// PwnAccount checks to see if an account has been pwned
// TODO: break up into Util functions
func PwnAccount(target string) {
	utils.CheckTarget(target) // make sure target is valid

	pwnURL := fmt.Sprintf(`https://haveibeenpwned.com/api/v2/breachedaccount/%v`, target)
	response, err := http.Get(pwnURL) // make response object
	utils.CheckErr(err)

	defer response.Body.Close()                   // close on function end
	bodyBytes, _ := ioutil.ReadAll(response.Body) // read bytes from response

	if len(bodyBytes) == 0 { // nothing found - all good
		ct.Foreground(ct.Green, true) // set text color to bright green
		fmt.Println("Good news — no pwnage found!")
	} else { // account found in breach
		ct.Foreground(ct.Red, true) // set text color to bright red
		fmt.Println("Oh no — account has been pwned!")
	}
}
