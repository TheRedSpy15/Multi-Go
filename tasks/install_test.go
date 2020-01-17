package tasks_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	tks "github.com/TheRedSpy15/Multi-Go/tasks"
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
		tks.Install(target)

		i, err := os.Stat(target)
		if err != nil || !i.IsDir() {
			t.Fatalf("expected %s to be a directory", target)
		}
	})
	cleanup()

	t.Run("Sets Execution Permission", func(t *testing.T) {
		tks.Install(target)

		files, err := ioutil.ReadDir(target)
		if err != nil {
			t.Fatal(err)
		} else if files[0].Mode()&0111 == 0 {
			t.Fatalf("installed file is not executable")
		}
		cleanup()

		// check for overwritten file
		tks.Install(target)
		tks.Install(target)

		files, err = ioutil.ReadDir(target)
		if err != nil {
			t.Fatal(err)
		} else if files[0].Mode()&0111 == 0 {
			t.Fatalf("installed file is not executable")
		}
	})
	cleanup()

	t.Run("Overwrites Existing File", func(t *testing.T) {
		// get the name of the file
		tks.Install(target)
		files, err := ioutil.ReadDir(target)
		if err != nil {
			t.Fatal(err)
		}
		name := path.Join(target, files[0].Name())

		// remove the file
		if err = os.Remove(name); err != nil {
			t.Fatal(err)
		}

		// write the file
		err = ioutil.WriteFile(name, []byte("test"), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		tks.Install(target)

		b, err := ioutil.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) == "test" {
			t.Fatal("installed file contents unchanged")
		}
	})
	cleanup()
}
