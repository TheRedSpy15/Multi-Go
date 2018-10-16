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
	"github.com/daviddengcn/go-colortext"
)

// Decompress will decompress the target file in gzip format
// NOTE: make sure to check for gzip extension
func Decompress(target string) {
	ct.Foreground(ct.Red, true) // set text color to bright red
	fmt.Println("Not a working feature yet!")
}
