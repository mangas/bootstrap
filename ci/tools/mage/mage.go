package mage

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/getcouragenow/bootstrap/ci/utils"
)

const (
	libName = "mage"
	lib     = "github.com/magefile/"
)

var (
	goPath    = ""
	goBIN     = ""
	libPath   = ""
	libFSPath = ""
)

func init() {
	p, err := utils.GetPath()
	if err != nil {
		log.Println(err)
		return
	}
	goPath = p["GOPATH"]
	goBIN = p["GOBIN"]
	libPath = path.Join(goPath, "src", lib)
	libFSPath = path.Join(libPath, libName)
}

// Install mage
func Install() {

	err := os.RemoveAll(libPath)
	if err != nil {
		log.Fatal(err)
	}

	os.Getenv("GOBIN")
	err = os.MkdirAll(libPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(libPath)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("git", "clone", "ssh://git@github.com/magefile/mage.git")
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(libFSPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("go", "run", "bootstrap.go")
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
