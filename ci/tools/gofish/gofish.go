package gofish

import (
	"log"
	"os"

	"github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/sh"
)

var (
	goPath       = ""
	goBin        = ""
	libName      = "gofish"
	lib          = ""
	libBranch    = "master"
	libTag       = "latest"
	libFSPATH    = ""
	libBinFSPATH = ""
)

func init() {
	p, _ := utils.GetPath()

	goPath = p["GOPATH"]
	goBin = p["GOBIN"]

	lib = "github.com/fishworks/" + libName
	libFSPATH = goPath + "/src/" + lib
	libBinFSPATH = goPath + "/bin/" + libName
}

// Install gofish
func Install() {
	if utils.CheckIfAlreadInstalled(libName) {
		return
	}
	// clone repo
	err := utils.GitClone(lib, libName, libFSPATH, "")
	if err != nil {
		log.Fatal(err)
	}

	// build gofish
	os.Chdir(libFSPATH)
	sh.RunV("./scripts/install.sh")
}

// Uninstall gofish
func Uninstall() {
	sh.Run("go", "clean", "-i", lib)
}
