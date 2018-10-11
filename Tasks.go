package main

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
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/daviddengcn/go-colortext"
	"github.com/jordan-wright/email"
)

// TODO: add sha-256 hash
// Takes a file path, and then prints the hash of the file
func hashFile(target string) {
	checkTarget(target)
	file := readFileIntoByte(target)                          // get bytes of file to hash
	hash := sha1.New()                                        // create sha1 object
	hash.Write(file)                                          // hash file to object
	target = base64.URLEncoding.EncodeToString(hash.Sum(nil)) // encode hash sum into string

	fmt.Println("SHA-1 hash :", target)
}

// ListTasks - lists all currently working tasks
func listTasks() {
	ct.Foreground(ct.Yellow, false)
	println("Available tasks:")
	println("Hash -r [file path]")
	println("encryptFile -r [file path]")
	println("decryptFile -r [file path]")
	println("Scrape - [URL]")
	println("DOS - [IP/URL]")
	println("Email")
	println("generatePassword")
	println("systemInfo")
	println("pwnAccount - [Email]")

	println("About") // keep at bottom of print statements
}

// TODO: make & add 'printDisk'
// Prints extensive info about system
func systemInfoTask() {
	ct.Foreground(ct.Yellow, false)
	printCPU()
	printMemory()
	printHost()
}

// Check if an account has been pwned
func pwnAccount(target string) {
	checkTarget(target) // make sure target is valid

	pwnURL := fmt.Sprintf(`https://haveibeenpwned.com/api/v2/breachedaccount/%v`, target)
	response, err := http.Get(pwnURL)
	if err != nil {
		panic(err.Error)
	}

	defer response.Body.Close()                   // close on function end
	bodyBytes, _ := ioutil.ReadAll(response.Body) // read bytes from response

	if len(bodyBytes) == 0 { // nothing found - all good
		ct.Foreground(ct.Green, true)
		println("Good news — no pwnage found!")
	} else {
		ct.Foreground(ct.Red, true)
		println("Oh no — account has been pwned!")
	}
}

// Encrypts the target file
func encryptFileTask(target string) {
	checkTarget(target) // make target is valid

	data := readFileIntoByte(target) // read file bytes
	print("Enter Password: ")
	password := getPassword() // get password securely

	encryptFile(target, data, password) // encrypt file
	println("\nFile encrypted!")
}

// BUG: decrypted file is unusable
// NOTE: decrypt file doesn't actually save as unencrypted
// Decrypts the target file
func decryptFileTask(target string) {
	checkTarget(target) // make target is valid

	print("Enter Password: ")
	password := getPassword() // get password securely

	file, err := os.Create(target) // create file object
	if err != nil {
		ct.Foreground(ct.Red, true)
		panic(err.Error())
	}
	defer file.Close()                        // close file on function end
	file.Write(decryptFile(target, password)) // decrypt file
	println("\nFile decrypted!")
}

// Prints details about the program
func about() {
	printBanner()
	println("Multi Go v1.0.0", "\nBy: TheRedSpy15")
	println("GitHub:", "https://github.com/TheRedSpy15")
	println("Project Page:", "https://github.com/TheRedSpy15/Multi-Go")
	println("\nMulti Go allows IT admins and Cyber Security experts")
	println("to conveniently perform all sorts of tasks.")
}

// Scrapes target website
func scapeTask(target string) {
	checkTarget(target)               // make target is valid
	collyAddress(target, true, false) // run colly
}

// TODO: use project path to find file
// BUG: exit status 1
// Runs linuxScanner.py to audit system vulnerabilities
func auditTask() {
	ct.Foreground(ct.Yellow, false)
	runAudit() // run audit
}

// TODO: use set length
// Generates a random string for use as a password
func generatePasswordTask() {
	ct.Foreground(ct.Yellow, false)
	println("Password:", randomString())
}

// TODO: add amplification
// TODO: more testing
// Indefinitely runs colly on an address
func dosTask(target string) {
	checkTarget(target) // make target is valid
	ct.Foreground(ct.Red, true)
	println("\nWarning: you are solely responsible for your actions!") // disclaimer
	println("ctrl + c to cancel")
	println("\n10 seconds until DOS")
	ct.ResetColor()

	time.Sleep(10 * time.Second) // 10 second delay - give chance to cancel

	for true { // DOS loop
		collyAddress(target, false, true)
	}
}

// BUG: mail: missing word in phrase: mail: invalid string
// TODO: use native go email
// TODO: break up into Util functions
// TODO: find out if attachment works with path, or just name
// Send email
func emailTask() {
	reader := bufio.NewReader(os.Stdin) // make reader object
	e := email.NewEmail()
	ct.Foreground(ct.Yellow, false)
	println("Prepare email")
	ct.ResetColor()

	// email setup
	print("From: ")
	e.From, _ = reader.ReadString('\n') // from

	print("To: ")
	To, _ := reader.ReadString('\n') // to
	e.To = []string{To}

	print("Bcc (leave blank if none): ") // bcc
	Bcc, _ := reader.ReadString('\n')
	e.Bcc = []string{Bcc}

	print("Cc (leave blank if none): ") // cc
	Cc, _ := reader.ReadString('\n')
	e.To = []string{Cc}

	print("Subject: ")
	e.Subject, _ = reader.ReadString('\n') // subject

	print("Text: ")
	Text, _ := reader.ReadString('\n') // text
	e.Text = []byte(Text)

	print("File path (if sending one): ") // attachment
	Path, _ := reader.ReadString('\n')
	if Path != "" {
		e.AttachFile(Path)
	}

	// authentication
	print("Provider (example: smtp.gmail.com): ") // provider
	provider, _ := reader.ReadString('\n')
	print("Port (example: 587): ") // port
	port, _ := reader.ReadString('\n')
	print("Password (leave blank if none): ") // password
	password := getPassword()

	// confirmation
	print("Confirm send? (yes/no): ")
	confirm, _ := reader.ReadString('\n')
	if strings.TrimRight(confirm, "\n") == "yes" {
		// sending
		err := e.Send(provider+":"+port, smtp.PlainAuth("", e.From, password, provider))
		if err != nil {
			ct.Foreground(ct.Red, true)
			println("error sending email -", err.Error())
		}
	} else { // cancelled
		ct.Foreground(ct.Red, true)
		println("Cancelled!")
	}
}
