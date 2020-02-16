package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/mage"
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
	goPath = os.Getenv("GOPATH")
	goBIN = os.Getenv("GOBIN")
	libPath = path.Join(goPath, "src", lib)
	libFSPath = path.Join(libPath, libName)
}

func main() {
	os.Exit(mage.Main())
}

func installMage() {

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
	fmt.Println(os.Getwd())
	cmd = exec.Command("go", "run", "bootstrap.go")
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

// func checkDeps() {
// 	_, err := os.Stat(goBIN + "/mage")
// 	if err != nil {
// 		fmt.Println("start installing Mage")
// 		installMage()
// 		fmt.Println("Mage installed")
// 	}

// 	_, err = os.Stat(goBIN + "/mages")
// 	if err != nil {
// 		fmt.Println("start installing Mage-select")
// 		m.MageSelect_Install()
// 		fmt.Println("Mage-select installed")
// 	}
// }
