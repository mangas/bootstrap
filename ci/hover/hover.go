package hover

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

var (
	GOPath    = os.Getenv("GOPATH")
	LibName   = "hover"
	Lib       = "github.com/go-flutter-desktop/" + LibName
	LibBranch = "master"
	LibTag    = "v0.37.0"
	LibFSPATH = GOPath + "/src/" + Lib

	LinBinName   = LibName
	LibBinFSPATH = GOPath + "/bin/" + LinBinName
)

func Hover_Install() {
	cloneHover()
	os.Chdir(LibFSPATH)
	sh.RunV("go", "build", "-o", LibBinFSPATH, ".")
	sh.RunV("hover")
}

func cloneHover() {
	sh.RunV("mkdir", "-p", LibFSPATH)
	fmt.Println(os.Chdir(LibFSPATH))
	fmt.Println(os.Chdir(".."))
	sh.RunV("rm", "-rf", LibName, "&&")
	sh.RunV("git", "clone", "ssh://git@"+Lib+".git")
}
