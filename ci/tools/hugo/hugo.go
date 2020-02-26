package hugo

import (
	"log"
	"os"

	"github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/sh"
)

var (
	goPath       = ""
	goBin        = ""
	libName      = "hugo"
	lib          = ""
	libBranch    = "master"
	libFSPATH    = ""
	linBinName   = libName
	libBinFSPATH = ""
)

func init() {
	p, _ := utils.GetPath()

	goPath = p["GOPATH"]
	goBin = p["GOBIN"]

	lib = "github.com/gohugoio/" + libName
	libFSPATH = goPath + "/src/" + lib
	libBinFSPATH = goPath + "/bin/" + linBinName
}

// Install hugo
func Install() {
	if utils.CheckIfAlreadInstalled(libName) {
		return
	}
	// clone repo
	err := utils.GitClone(lib, libName, libFSPATH, "")
	if err != nil {
		log.Fatal(err)
	}

	// build hugo
	os.Chdir(libFSPATH)
	sh.RunV("go", "install")
}

// Uninstall hugo
func Uninstall() {
	sh.Run("go", "clean", "-i", lib)
}
