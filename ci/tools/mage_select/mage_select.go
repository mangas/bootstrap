package mageselect

import (
	"os"
	"os/exec"
	"path"

	"github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/sh"
)

const (
	libName = "mage-select"
	lib     = "github.com/iwittkau/"
)

var (
	goPath  = ""
	goBin   = ""
	libPath = ""
)

func init() {
	p, _ := utils.GetPath()
	goPath = p["GOPATH"]
	goBin = p["GOBIN"]
	libPath = path.Join(goPath, "src", lib)
}

// Install mage-select.
func Install() {
	if exec.Command("mages").Run() != nil {
		getMageSelect()
	}
}

// Update mage-select.
func Update() {
	getMageSelect()
}

// Uninstall mage-select.
func Uninstall() {
	sh.RunWith(nil, "rm", "-rf", libPath)
	sh.RunWith(nil, "rm", path.Join(goBin, "mages"))
}

func getMageSelect() {
	// get mage select from github
	if err := sh.RunWith(map[string]string{"GO111MODULE": "off"}, "go", "get", "-u", "-d", lib+libName); err != nil {
		return
	}

	// install mage select
	magePath := path.Join(libPath, libName)
	os.Chdir(magePath)
	sh.RunV("mage", "install")
}
