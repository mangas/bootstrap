package gofish
//
// import (
// 	"log"
// 	"os"
// 	"runtime"
//
// 	"github.com/fatih/color"
// 	"github.com/getcouragenow/bootstrap/ci/os/util"
// 	"github.com/getcouragenow/bootstrap/ci/utils"
// 	"github.com/magefile/mage/sh"
// )
//
// var (
// 	goPath       = ""
// 	goBin        = ""
// 	libName      = "gofish"
// 	lib          = "github.com/fishworks/" + libName
// 	libBranch    = "master"
// 	libTag       = "latest"
// 	libFSPATH    = ""
// 	libBinFSPATH = ""
// )
//
// func init() {
// 	p, _ := utils.GetPath()
//
// 	goPath = p["GOPATH"]
// 	goBin = p["GOBIN"]
//
// 	libFSPATH = goPath + "/src/" + lib
// 	libBinFSPATH = goPath + "/bin/" + libName
// }
//
// // Install gofish
// func Install() {
// 	if utils.CheckIfAlreadInstalled(libName) {
// 		return
// 	}
//
// 	if runtime.GOOS == "windows" {
// 		posh := util.New()
// 		var err error
// 		_, _, err = posh.Execute("Set-ExecutionPolicy Bypass -Scope Process -Force")
// 		if err != nil {
// 			color.Red(err.Error())
// 		}
//
// 		_, _, err = posh.Execute("iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/fishworks/gofish/master/scripts/install.ps1'))")
// 		if err != nil {
// 			color.Red(err.Error())
// 		}
// 	} else {
// 		// clone repo
// 		err := utils.GitClone(lib, libName, libFSPATH, "")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Chdir(libFSPATH)
// 		// build gofish
// 		sh.RunV("./scripts/install.sh")
// 	}
// }
//
// // Uninstall gofish
// func Uninstall() {
// 	sh.Run("go", "clean", "-i", lib)
// }
