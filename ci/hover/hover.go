package hover

import (
	"fmt"
	"log"
	"os"

	u "github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Hover namespace
type Hover mg.Namespace

var (
	home         = ""
	goPath       = ""
	goBin        = ""
	libName      = "hover"
	lib          = ""
	libBranch    = "master"
	libTag       = "v0.37.0"
	libFSPATH    = ""
	linBinName   = libName
	libBinFSPATH = ""
)

func init() {
	path, err := u.GetPath()

	if err != nil {
		log.Fatal("Error Could not get go path infos:", err)
	}

	home = path["HOME"]
	goPath = path["GOPATH"]
	goBin = path["GOBIN"]
	lib = "github.com/go-flutter-desktop/" + libName
	libFSPATH = goPath + "/src/" + lib
	libBinFSPATH = goPath + "/bin/" + linBinName

	fmt.Println(libFSPATH)
}

// Install hover.
func (Hover) Install() {

	// clone repo
	err := u.GitClone(lib, libName, libFSPATH, "")
	if err != nil {
		log.Fatal(err)
	}

	// build hover
	os.Chdir(libFSPATH)
	sh.RunV("go", "build", "-o", libBinFSPATH, ".")
	sh.RunV("hover")
}
