package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// Installer - will either add the ability to easily call Multi-Go, or remove it
// TODO: not working yet
// TODO: make less error prone
func Installer(target string) {
	utils.CheckTarget(target)

	if target == "install" { // install
		if _, err := os.Stat("/bin"); !os.IsNotExist(err) { // bin already exists
			utils.CheckErr(err)
			fmt.Println("bin exists")

			// get path of current Multi-Go program to copy
			fmt.Println("Getting application path")
			appPath, _ := os.Executable()
			srcPath := strings.Replace(appPath, "<nil>", "", 1)
			strings.TrimSpace(srcPath)

			// copy to bin
			fmt.Println("Copying to bin")
			err := os.Link(srcPath, "/bin")
			utils.CheckErr(err)

			// execution permisssions
			fmt.Println("Getting execution permission")
			utils.RunCmd("chmod +x $HOME/bin/Multi-Go")

			fmt.Println("Done")
		} // need to create bin
	} // need to add uninstall option
}
