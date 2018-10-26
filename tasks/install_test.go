package tasks

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestInstall(t *testing.T) {
	target := "install"
	cleanup := func() {
		if err := os.RemoveAll(target); err != nil && !os.IsNotExist(err) {
			t.Error(err)
		}
	}

	cleanup()
	t.Run("Creates Target Directory", func(t *testing.T) {
		Install(target)

		i, err := os.Stat(target)
		if err != nil || !i.IsDir() {
			t.Errorf("expected %s to be a directory", target)
		}
		cleanup()
	})

	t.Run("Sets Execution Permission", func(t *testing.T) {
		Install(target)

		files, err := ioutil.ReadDir(target)
		if err != nil {
			t.Error(err)
		} else if files[0].Mode()&0111 == 0 {
			t.Errorf("installed file is not executable")
		}
		cleanup()
	})
}
