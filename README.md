# Multi-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/TheRedSpy15/Multi-Go)](https://goreportcard.com/report/github.com/TheRedSpy15/Multi-Go)
[![codebeat badge](https://codebeat.co/badges/d6180a76-99be-4013-a0c2-0e4bcf0b9655)](https://codebeat.co/projects/github-com-theredspy15-multi-go-master)
[![CodeFactor](https://www.codefactor.io/repository/github/theredspy15/multi-go/badge)](https://www.codefactor.io/repository/github/theredspy15/multi-go)
[![Build Status](https://travis-ci.com/TheRedSpy15/Multi-Go.svg?branch=master)](https://travis-ci.com/TheRedSpy15/Multi-Go)
[![Build Status](https://semaphoreci.com/api/v1/theredspy15/multi-go/branches/master/shields_badge.svg)](https://semaphoreci.com/theredspy15/multi-go)

A command line multi-tool made in Go, and aimed at security experts to make life a little more convenient. It does this by combining a massive array of different tasks, into one program.

[![asciicast](https://asciinema.org/a/209295.png)](https://asciinema.org/a/209295)

### Currently capable of:
- file hashing
- DOS attack
- password generator
- system info
- check if account is breached (HaveIBeenPwned)
- control firewall
- system security audit
- clean system files
- secure file deletion / bleach
- viewing the headlines in the cyber world (ycombinator rss)

### Plans:
- file compression/decompression (gzip)
- email
- scrape website(s) for information
- file encryption/decryption
- password cracking
- network logging (tshark)
- network scan (ip/port/mac/etc)
- installer for multiple useful pentesting tools

**Will add more to the list over time**

## How to

### Download (pre-compiled):
[Click here to download](https://github.com/TheRedSpy15/Multi-Go/releases/download/0.6.1/MultiGo_0_6_1)

### Use
1. Open the terminal
2. Type `Multi-Go` (don't forget to cd if you didn't install)
3. **OPTIONAL:** follow that with "-t/--Task [task] -r/--Target [target]". Note: the 'target' is optional, depending on the task
4. Hit enter

### Contribute:
Simply make a pull request, I have yet to turn one down.
**NOTE:** Currently, I am just relying on TODOS in the comments of the code, as a temporary (as in, will change) replacement for 'issues'

**IMPORTANT:** When adding a new task, you must follow this pattern!
1. Create a new file in the *tasks* directory and write all of your code there.
3. If you feel any code in your class may be used in other tasks, feel free to put it in `utils.go`.
4. Ensure your code is documented well (running *golint* is helpful).
5. New tasks should have an associated test file (e.g. `mytask_test.go`) in the same folder.

If the new feature is complete:
1. Add the case to the switch statement in `main.go`, so your new task can be called.
2. Finished!

### Build from source (requires Go installed)
```
git clone https://github.com/TheRedSpy15/Multi-Go
cd Multi-Go
go build
```
or
```
go get github.com/TheRedSpy15/Multi-Go
cd go/src/github.com/TheRedSpy15/Multo-Go
```


## Important
Multi Go is intended to be used on linux (mostly Debian & Ubuntu like distros). It might run on Windows. Currently it isn't tested, nor supported! I will eventually work on a Windows patch.
