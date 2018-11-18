package tasks

import (
	"github.com/TheRedSpy15/Multi-Go/utils"
)

// CreateUsb downloads a zip folder with a bunch of tools at a target location
func CreateUsb(target string) {
	utils.CheckTarget(target)

	utils.DownloadFile(target, "nil") // tool repo not create yet
}
