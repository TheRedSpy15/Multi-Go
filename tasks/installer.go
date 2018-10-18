package tasks

import (
	"fmt"
	"os"
	"strings"

	"github.com/TheRedSpy15/Multi-Go/utils"
)

func Installer(target string) {
	utils.CheckTarget(target)

	if target == "install" {
		if _, err := os.Stat("/bin"); !os.IsNotExist(err) {
			fmt.Println("bin exists")

			appPath, _ := os.Executable()
			srcPath := strings.Replace(appPath, "<nil>", "", 1)
			strings.TrimSpace(srcPath)

			fmt.Println(srcPath)
			err := os.Link(srcPath, "/bin")
			if err != nil {
				panic(err.Error())
			}

			utils.RunCmd("chmod +x $HOME/bin/Multi-Go")
		}
	}
}
