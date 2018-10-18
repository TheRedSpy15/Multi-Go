package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

// Installer - will either add the ability to easily call Multi-Go, or remove it
// TODO: not working yet
func Installer(target string) {
	utils.CheckTarget(target)

	if target == "install" {
		if _, err := os.Stat("/bin"); !os.IsNotExist(err) {
			utils.CheckErr(err)
			fmt.Println("bin exists")

			appPath, _ := os.Executable()
			srcPath := strings.Replace(appPath, "<nil>", "", 1)
			strings.TrimSpace(srcPath)

			fmt.Println(srcPath)
			err := os.Link(srcPath, "/bin/Multi-Go")
			utils.CheckErr(err)

			utils.RunCmd("chmod +x $HOME/bin/Multi-Go")
		}
	}
}
