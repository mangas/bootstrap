package mageselect

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/sh"
)

const (
	LibName = "mage-select"
	Lib     = "github.com/iwittkau/"
)

var (
	GOPATH  = ""
	GOBIN   = ""
	LibPath = ""
)

func init() {
	GOPATH = os.Getenv("GOPATH")
	GOBIN = os.Getenv("GOBIN")

	LibPath = path.Join(GOPATH, "src", Lib)
}

// Install mage-select.
func MageSelect_Install() {
	if exec.Command("mages").Run() != nil {
		getMageSelect()
	}
	fmt.Println("\nFinished.")
}

// Update mage-select.
func MageSelect_Update() {
	getMageSelect()
	fmt.Println("\n Mage-select was updated.")
}

// Uninstall mage-select.
func MageSelect_Uninstall() {
	sh.RunWith(nil, "rm", "-rf", LibPath)
	sh.RunWith(nil, "rm", path.Join(GOBIN, "mages"))
	fmt.Println("Deleted")
}

func getMageSelect() {
	// get mage select from github
	fmt.Println("\nGet dependencies...")
	if err := sh.RunWith(map[string]string{"GO111MODULE": "off"}, "go", "get", "-u", "-d", Lib+LibName); err != nil {
		return
	}

	// install mage select
	fmt.Println("\nInstall dependencies...")
	magePath := path.Join(LibPath, LibName)
	os.Chdir(magePath)
	sh.RunV("mage", "install")
}
