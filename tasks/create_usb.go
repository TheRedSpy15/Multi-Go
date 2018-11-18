package tasks

import (
	"fmt"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// CreateUsb downloads a zip folder with a bunch of tools at a target location
func CreateUsb(target string) {
	utils.CheckTarget(target)

	fmt.Println("Downloading package")
	utils.CheckErr(utils.DownloadFile(target, "nil")) // tool repo not create yet
	fmt.Println("Finished!")
}
