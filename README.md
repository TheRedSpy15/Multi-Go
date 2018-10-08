# Multi-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/TheRedSpy15/Multi-Go)](https://goreportcard.com/report/github.com/TheRedSpy15/Multi-Go)
[![codebeat badge](https://codebeat.co/badges/d6180a76-99be-4013-a0c2-0e4bcf0b9655)](https://codebeat.co/projects/github-com-theredspy15-multi-go-master)
[![CodeFactor](https://www.codefactor.io/repository/github/theredspy15/multi-go/badge)](https://www.codefactor.io/repository/github/theredspy15/multi-go)

A command line multi-tool made in Go, and aimed at security experts to make life a little more convenient. It does this by combining a massive array of different tasks, into one program.
### Currently capable of:
- file encryption/decryption
- file hashing
- DOS attack
- email
- scrape website
- password generator
- system info
### Working on (will add more over time):
- secure file deletion
- file compression/decompression
- check if account is breached (HaveIBeenPwned)
- toggle incoming connections
- system vulerability audit (online/offline modes)
- clean temporary files
## How to
### Download:
There isn't an offical executable yet. However, I should have one under 'releases' in the next week.
### Use:
1. Open the terminal
2. paste path to executable
3. follow that with "-t/--Task [task] -r/--Target [target]"
4. Note: the 'target' is optional, depending on the task
### Contribute:
Simply make a pull pull request, I have yet to turn down one.
**NOTE:** I Currently, I am just relying on TODOS in the comments of the code, as a temporary (as in, will change) replacement for 'issues'

**IMPORTANT:** When working on adding a feature, you must follow this pattern!
1. Create the method/function to be called in Tasks.go (with the name "newFeatureTask").
2. Write all your code in there.
3. Break that up into multiple functions and put those in Utils.go.
4. Go back and add a lot of comments.

If the new feature is complete:
1. Add it to the list in listTasks(), in Tasks.go.
2. Add the case to the switch statement in MultiGo.go, so it (your new feature in Tasks.go) can be called.
3. Finished!
## Important
Multi Go is intended to be used on linux. It might run on Windows. Currently it isn't tested, nor supported! I will eventually work on a Windows patch
