package hover
//
// import (
// 	"log"
// 	"os"
//
// 	"github.com/getcouragenow/bootstrap/ci/utils"
// 	"github.com/magefile/mage/sh"
// )
//
// var (
// 	goPath       = ""
// 	goBin        = ""
// 	libName      = "hover"
// 	lib          = ""
// 	libBranch    = "master"
// 	libTag       = "v0.37.0"
// 	libFSPATH    = ""
// 	linBinName   = libName
// 	libBinFSPATH = ""
// )
//
// func init() {
// 	p, _ := utils.GetPath()
//
// 	goPath = p["GOPATH"]
// 	goBin = p["GOBIN"]
//
// 	lib = "github.com/go-flutter-desktop/" + libName
// 	libFSPATH = goPath + "/src/" + lib
// 	libBinFSPATH = goPath + "/bin/" + linBinName
// }
//
// // Install hover
// func Install() {
// 	if utils.CheckIfAlreadInstalled(libName) {
// 		return
// 	}
// 	// clone repo
// 	err := utils.GitClone(lib, libName, libFSPATH, "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// build hover
// 	os.Chdir(libFSPATH)
// 	sh.RunV("go", "build", "-o", libBinFSPATH, ".")
// }
//
// // Uninstall hover
// func Uninstall() {
// 	sh.Run("go", "clean", "-i", lib)
// }
