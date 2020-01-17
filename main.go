package main

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

// Project TODOS
// TODO tone down comments & make them more meaningful
// TODO improve 'Scrape'
// TODO finish email task
// TODO finish decompress (and review compress)
// TODO add 'tshark -r [file path]' task
// TODO add network scanner
// TODO add wifi password cracker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	ct "github.com/daviddengcn/go-colortext"

	"github.com/TheRedSpy15/Multi-Go/tasks"
	"github.com/TheRedSpy15/Multi-Go/utils"
)

func main() {
	dialogMode := false

	parser := argparse.NewParser("Multi-Go", "Runs multiple security orientated tasks")

	// Create flags
	t := parser.String("t", "Task", &argparse.Options{Required: false, Help: "Task to run"})
	r := parser.String("r", "Target", &argparse.Options{Required: false, Help: "Target to run task on"})

	err := parser.Parse(os.Args) // parse arguments
	utils.CheckErr(err)

	reader := bufio.NewReader(os.Stdin) // make reader object

	if *t == "" { // enter dialog mode
		dialogMode = true
		utils.PrintBanner()
		tasks.List()
	} else {
		ct.Foreground(ct.Yellow, false)
	}

	//Only continue execution in dialog mode
	for contExec := true; contExec; contExec = dialogMode {
		if dialogMode {
			fmt.Print("\nEnter task to run: ")
			choice, _ := reader.ReadString('\n')     // get choice
			choice = strings.TrimRight(choice, "\n") // trim choice so it can be check against properly

			if strings.Contains(choice, "-r") { // check for optional target
				inputs := strings.Split(choice, " -r ") // separate task & target
				*t = inputs[0]
				*r = inputs[1]
			} else { // no optional target
				*t = choice
			}
		}

		taskMap := map[string]func(string){
			"Hash":             tasks.HashFile,
			"pwnAccount":       tasks.PwnAccount,
			"encryptFile":      tasks.EncryptFile,
			"decryptFile":      tasks.DecryptFile,
			"Scrape":           tasks.Scrape,
			"DOS":              func(s string) { tasks.Dos(s, nil) },
			"compress":         tasks.Compress,
			"decompress":       tasks.Decompress,
			"Firewall":         tasks.ToggleFirewall,
			"generatePassword": func(s string) { tasks.GeneratePassword(s, os.Stdout) },
			"Install":          tasks.Install,
			"Bleach":           tasks.Bleach,
			"cyberNews":        func(s string) { tasks.News() },
			"systemInfo":       func(s string) { tasks.SystemInfo() },
			"Clean":            func(s string) { tasks.Clean() },
			"Email":            func(s string) { tasks.Email() },
			"Audit":            func(s string) { tasks.Audit() },
			"About":            func(s string) { tasks.About() },
			"List":             func(s string) { tasks.List() },
			"Exit":             func(s string) { os.Exit(0) },
		}

		task, ok := taskMap[*t]

		if ok {
			fmt.Println("\nRunning task:", *t, "\nTarget:", *r)
			task(*r)
		} else {
			ct.Foreground(ct.Red, true)
			fmt.Println("Invalid task -", *t)
			ct.Foreground(ct.Yellow, false)
			fmt.Println("Use '--help' or '-t List'")
		}
	}
}
