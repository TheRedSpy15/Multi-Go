package main

// Project TODOS
// TODO: tone down comments
// TODO: add unit testing
// TODO: document parameters
// TODO: improve 'Scrape'
// TODO: finish dos task
// TODO: finish email task
// TODO: finish audit task
// TODO: add 'bleach -r [file path]' task
// TODO: add 'uncompress -r [file path]' task

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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/daviddengcn/go-colortext"
)

func main() {
	parser := argparse.NewParser("SecureMultiTool", "Runs multiple security orientated tasks")

	// Create flags
	t := parser.String("t", "Task", &argparse.Options{Required: false, Help: "Task to run"})
	r := parser.String("r", "Target", &argparse.Options{Required: false, Help: "Target to run task on"})

	err := parser.Parse(os.Args) // parse arguments
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error)
	}

	if *t == "" { // enter dialog mode
		reader := bufio.NewReader(os.Stdin) // make reader object
		printBanner()
		listTasks()

		print("\nEnter task to run: ")
		choice, _ := reader.ReadString('\n')     // get choice
		choice = strings.TrimRight(choice, "\n") // trim choice so it can be check against properly

		if strings.Contains(choice, "-r") { // check for optional target
			inputs := strings.Split(choice, " -r ") // separate task & target
			*t = inputs[0]
			*r = inputs[1]
		} else { // no optional target
			*t = choice
		}

		ct.ResetColor() // reset text color to default
	}

	// Determine task
	switch *t {
	case "Hash":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		hashFile(*r)
	case "pwnAccount":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		pwnAccount(*r)
	case "encryptFile":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		encryptFileTask(*r)
	case "decryptFile":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		decryptFileTask(*r)
	case "Scrape":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		scapeTask(*r)
	case "DOS":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		dosTask(*r)
	case "compress":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		compressTask(*r)
	case "decompress":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		decompressTask(*r)
	case "Firewall":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		toggleFirewall(*r)
	case "generatePassword":
		fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
		generatePasswordTask(*r)
	case "systemInfo":
		systemInfoTask()
	case "Clean":
		cleanTask()
	case "Email":
		emailTask()
	case "Audit":
		auditTask()
	case "About":
		about()
	case "List":
		listTasks()
	default: // invalid
		ct.Foreground(ct.Red, true)
		fmt.Println("Invalid task -", *t)
		ct.Foreground(ct.Yellow, false)
		fmt.Println("Use '--help' or '-t List'")
	}
}
