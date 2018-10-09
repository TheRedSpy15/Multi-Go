package main

// Project TODOS
// TODO: finish email task
// TODO: finish audit task
// TODO: add 'bleach -r [file path]' task
// TODO: add 'compress -r [file path]' task
// TODO: add 'uncompress -r [file path]' task
// TODO: add 'toggleIncoming -r [allow/block]' (inbound connections) task
// TODO: add 'systemInfo' task
// TODO: add 'auditOffline' (also add "run offline?" to 'audit', when no internet) task
// TODO: add 'cleanTemp' task

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
	"os"

	"github.com/akamensky/argparse"
	"github.com/daviddengcn/go-colortext"
)

func main() {
	parser := argparse.NewParser("SecureMultiTool", "Runs multiple security orientated tasks")

	// TODO: use native go flags
	// TODO: if no arguments are given, add dialog that asks user what they want to do (after above todo)
	// Create flags
	t := parser.String("t", "Task", &argparse.Options{Required: true, Help: "Task to run"})
	r := parser.String("r", "Target", &argparse.Options{Required: false, Help: "Target to run task on"})

	// Error handling
	err := parser.Parse(os.Args)
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error)
	}

	// Determine task
	switch *t {
	case "Hash":
		println("\nRunning task:", *t, "\nTarget:", *r)
		hashFile(*r)
	case "pwnAccount":
		println("\nRunning task:", *t, "\nTarget:", *r)
		pwnAccount(*r)
	case "encryptFile":
		println("\nRunning task:", *t, "\nTarget:", *r)
		encryptFileTask(*r)
	case "decryptFile":
		println("\nRunning task:", *t, "\nTarget:", *r)
		decryptFileTask(*r)
	case "Scrape":
		println("\nRunning task:", *t, "\nTarget:", *r)
		scapeTask(*r)
	case "DOS":
		println("\nRunning task:", *t, "\nTarget:", *r)
		dosTask(*r)
	case "Audit":
		println("\nRunning task:", *t, "\nTarget:", *r)
		auditTask(*r)
	case "generatePassword":
		generatePasswordTask()
	case "systemInfo":
		systemInfoTask()
	case "Email":
		emailTask()
	case "About":
		about()
	case "List":
		listTasks()
	default: // invalid
		ct.Foreground(ct.Red, true)
		println("Invalid task - ", *t)
		ct.Foreground(ct.Yellow, false)
		println("Use '--help' or '-t List'")
	}
}
