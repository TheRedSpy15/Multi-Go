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
	"compress/gzip"
	"fmt"
	"os"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
)

// Compress will compress the target file into gzip format
// TODO: rework gzip extension adding
func Compress(target string) {
	utils.CheckTarget(target) // make sure target is valid

	file, err := os.Create(target) // create file object
	utils.CheckErr(err)
	defer file.Close() // make sure file gets closed

	os.Rename(target, target+".gz") // add gzip extension

	w := gzip.NewWriter(file)               // make gzip writer for target file
	w.Write(utils.ReadFileIntoByte(target)) // write compressed data
	defer w.Close()                         // make sure writer gets closed

	ct.Foreground(ct.Green, true) // set text color to bright green
	fmt.Println("finished!")
}
