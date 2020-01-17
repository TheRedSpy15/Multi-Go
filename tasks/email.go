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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
	"github.com/daviddengcn/go-colortext"
	"gopkg.in/gomail.v2"
)

// Email sends an email
// BUG: no such host (likely because \n in input)
// TODO: break up into Util functions
// TODO: add more comments
// TODO: find out if attachment works with path, or just name
func Email() {
	reader := bufio.NewReader(os.Stdin) // make reader object
	e := gomail.NewMessage()            // make email object
	ct.Foreground(ct.Yellow, false)     // set text color to dark yellow
	fmt.Println("Prepare email")
	ct.ResetColor() // reset text color to default

	// email setup
	print("From: ")
	from, _ := reader.ReadString('\n') // from
	e.SetHeader("From", from)

	print("To: ")
	to, _ := reader.ReadString('\n') // to
	e.SetHeader("To", to)

	print("Subject: ")
	subject, _ := reader.ReadString('\n') // subject
	e.SetHeader("Subject", subject)

	print("Text: ")
	text, _ := reader.ReadString('\n') // text
	e.SetHeader("text/html", text)

	print("File path (if sending one): ") // attachment
	Path, _ := reader.ReadString('\n')
	if Path != "" {
		e.Attach(Path)
	}

	// authentication
	print("Provider (example: smtp.gmail.com): ") // provider
	provider, _ := reader.ReadString('\n')
	print("Port (example: 587): ") // port
	port, _ := reader.ReadString('\n')
	portCode, _ := strconv.Atoi(port)
	print("Password (leave blank if none): ") // password
	password := utils.GetPassword()

	// confirmation
	print("Confirm send? (yes/no): ")
	confirm, _ := reader.ReadString('\n')          // get string of user confirm choice
	if strings.TrimRight(confirm, "\n") == "yes" { // yes - confirm send
		// sending
		d := gomail.NewDialer(provider, portCode, from, password)

		err := d.DialAndSend(e)
		utils.CheckErr(err)
	} else { // cancelled
		ct.Foreground(ct.Red, true)
		fmt.Println("Cancelled!")
	}
}
