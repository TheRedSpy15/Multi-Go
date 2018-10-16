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
	"compress/gzip"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/daviddengcn/go-colortext"
	"github.com/shirou/gopsutil/cpu"
	"gopkg.in/gomail.v2"
)

// BUG: won't work unless ran from non-dialog mode / by using commandline flags
// Takes a file path, and then prints the hash of the file
func hashFile(target string) {
	checkTarget(target)             // make sure target is valid
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	file := readFileIntoByte(target)                          // get bytes of file to hash
	hash := sha1.New()                                        // create sha1 object
	hash.Write(file)                                          // hash file to object
	target = base64.URLEncoding.EncodeToString(hash.Sum(nil)) // encode hash sum into string

	fmt.Println("SHA-1 hash :", target)
}

// TODO: find some way to shrink
// ListTasks - lists all currently working tasks
func listTasks() {
	ct.Foreground(ct.Red, true)
	fmt.Println("Available tasks:")
	time.Sleep(1 * time.Second)

	fmt.Println("\n-- Utility --")
	ct.Foreground(ct.Yellow, false)
	fmt.Println("Scrape -r [URL]")
	fmt.Println("Email")
	fmt.Println("systemInfo")
	fmt.Println("Compress -r [file path]")
	time.Sleep(1 * time.Second)

	ct.Foreground(ct.Red, true)
	fmt.Println("\n-- Security --")
	ct.Foreground(ct.Yellow, false)
	fmt.Println("(sudo) Firewall -r [enable/disable/status]")
	fmt.Println("(sudo) Audit")
	fmt.Println("Hash -r [file path]")
	fmt.Println("encryptFile -r [file path]")
	fmt.Println("decryptFile -r [file path]")
	fmt.Println("pwnAccount -r [email]")
	fmt.Println("generatePassword -r [length]")
	time.Sleep(1 * time.Second)

	ct.Foreground(ct.Red, true)
	fmt.Println("\n-- Pentesting -- ")
	ct.Foreground(ct.Yellow, false)
	fmt.Println("DOS -r [IP:PORT]")
	time.Sleep(1 * time.Second)

	ct.Foreground(ct.Red, true)
	fmt.Println("\n-- Other --")
	ct.Foreground(ct.Yellow, false)
	fmt.Println("About")
}

// TODO: make & add more info functions
// Prints extensive info about system
func systemInfoTask() {
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow
	printCPU()                      // print cpu info
	printMemory()                   // print memory info
	printHost()                     // print host info
}

// TODO: break up into Util functions
// Check if an account has been pwned
func pwnAccount(target string) {
	checkTarget(target) // make sure target is valid

	pwnURL := fmt.Sprintf(`https://haveibeenpwned.com/api/v2/breachedaccount/%v`, target)
	response, err := http.Get(pwnURL) // make response object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error)
	}

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

// Encrypts the target file
func encryptFileTask(target string) {
	checkTarget(target)             // make sure target is valid
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	data := readFileIntoByte(target) // read file bytes
	print("Enter Password: ")
	password := getPassword() // get password securely

	encryptFile(target, data, password) // encrypt file
	fmt.Println("\nFile encrypted!")
}

// BUG: decrypted file is unusable
// NOTE: decrypt file doesn't actually save as unencrypted
// Decrypts the target file
func decryptFileTask(target string) {
	checkTarget(target)             // make sure target is valid
	ct.Foreground(ct.Yellow, false) // set text color to dark yellow

	print("Enter Password: ")
	password := getPassword() // get password securely

	file, err := os.Create(target) // create file object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}
	defer file.Close()                        // makes sure file gets closed
	file.Write(decryptFile(target, password)) // decrypt file
	fmt.Println("\nFile decrypted!")
}

// TODO: run the right command that cleans "thumbs" & the system cache
// Clean cached files
func cleanTask() {
	ct.Foreground(ct.Red, true) // set text color to bright red
	fmt.Println("Not a working feature yet!")
	cmd := exec.Command("rm", "-rf", "~/.thumbs/*") // don't think this is the right command
	cmd.Run()
}

// Prints details about the program
func about() {
	printBanner()

	ct.Foreground(ct.Yellow, false) // set text color to dark yellow
	fmt.Println("Multi Go v0.6.1", "\nBy: TheRedSpy15")
	fmt.Println("GitHub:", "https://github.com/TheRedSpy15")
	fmt.Println("Project Page:", "https://github.com/TheRedSpy15/Multi-Go")
	fmt.Println("\nMulti Go allows IT admins and Cyber Security experts")
	fmt.Println("to conveniently perform all sorts of tasks.")
}

// Scrapes target website
func scapeTask(target string) {
	checkTarget(target)               // make sure target is valid
	collyAddress(target, true, false) // run colly - scraping happens here
}

// Runs multiple checks, and reports found security issues to user
func auditTask() {
	runAuditOffline()
}

// TODO: rework gzip extension adding
// Compresses the target file in gzip format
func compressTask(target string) {
	checkTarget(target) // make sure target is valid

	file, err := os.Create(target) // create file object
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}
	defer file.Close() // make sure file gets closed

	os.Rename(target, target+".gz") // add gzip extension

	w := gzip.NewWriter(file)         // make gzip writer for target file
	w.Write(readFileIntoByte(target)) // write compressed data
	defer w.Close()                   // make sure writer gets closed

	ct.Foreground(ct.Green, true) // set text color to bright green
	fmt.Println("finished!")
}

// NOTE: make sure to check for gzip extension
// Decompresses the target file in gzip format
func decompressTask(target string) {
	ct.Foreground(ct.Red, true) // set text color to bright red
	fmt.Println("Not a working feature yet!")
}

// TODO: add support for more systems - think only works on debian/ubuntu
// Allows the user to enable/disable system firewall
func toggleFirewall(target string) {
	checkTarget(target)                  // make sure target is valid
	fmt.Println(runCmd("ufw", "status")) // run command & print result
}

// Generates a random string for use as a password
func generatePasswordTask(target string) {
	ct.Foreground(ct.Yellow, false)       // set text color to dark yellow
	conversion, _ := strconv.Atoi(target) // convert target (string), to int
	fmt.Println("Password:", randomString(conversion))
}

// TODO: add amplification - such as NTP monlist
// Indefinitely sends data to target
func dosTask(target string) {
	checkTarget(target) // make sure target is valid

	conn, err := net.Dial("udp", target) // setup connection object
	defer conn.Close()                   // make sure to close connection when finished
	if err != nil {
		ct.Foreground(ct.Red, true) // sets text color to bright red
		panic(err.Error)
	} else { // nothing bad happened when connecting to target
		ct.Foreground(ct.Green, true) // ets text color to bright red
		fmt.Println("Checks passed!")
	}

	ct.Foreground(ct.Red, true)                                            // set text color to bright red
	fmt.Println("\nWarning: you are solely responsible for your actions!") // disclaimer
	fmt.Println("ctrl + c to cancel")
	fmt.Println("\n10 seconds until DOS")
	ct.ResetColor() // reset text color to default

	time.Sleep(10 * time.Second) // 10 second delay - give chance to cancel

	threads, err := cpu.Counts(false) // get threads on system to set DOS thread limit
	if err != nil {
		ct.Foreground(ct.Red, true) // set text color to bright red
		panic(err.Error())
	}

	for i := 0; i < threads; i++ { // create DOS threads within limit
		go dos(conn)                   // create thread
		ct.Foreground(ct.Yellow, true) // set text color to dark yellow
		fmt.Println("Thread created!")
	}
}

// BUG: no such host (likely because \n in input)
// TODO: break up into Util functions
// TODO: add more comments
// TODO: find out if attachment works with path, or just name
// Send email
func emailTask() {
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
	password := getPassword()

	// confirmation
	print("Confirm send? (yes/no): ")
	confirm, _ := reader.ReadString('\n')          // get string of user confirm choice
	if strings.TrimRight(confirm, "\n") == "yes" { // yes - confirm send
		// sending
		d := gomail.NewDialer(provider, portCode, from, password)

		if err := d.DialAndSend(e); err != nil {
			ct.Foreground(ct.Red, true) // set text
			panic(err.Error())
		}
	} else { // cancelled
		ct.Foreground(ct.Red, true)
		fmt.Println("Cancelled!")
	}
}
