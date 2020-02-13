package gsheet

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// mage:import
//_ "go.zenithar.org/spotigraph/.mage"

type Gsheet mg.Namespace

var platforms = []string{"windows/amd64", "linux/amd64", "darwin/amd64"}

// Build googlesheet tool.
func (Gsheet) Build() {
	err := os.Chdir("./tool/googlesheet")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = sh.Run("go", "build", ".", "&&", "mv", "./googlesheet", "../../")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Chdir("../../")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = sh.Run("ls")
}

// // Run googlesheet Tests.
// func (Gsheet) RunTests() {
// 	sh.RunV("go", "test", "-v")
// }

// // Clean googlesheet project.
// func (Gsheet) Clean() {
// 	sh.RunV("go", "clean", "-cache")
// 	sh.RunV("go", "clean", "-modcache")
// 	sh.RunV("go", "clean", "-testcache")
// 	sh.RunV("rm", Name)
// }

// // Generate dump data.
// func (Gsheet) RunDataDump() {
// 	getData("datadump")
// }

// // Generate language data.
// func (Gsheet) RunLangDump() {
// 	getData("lang")
// }

// // Generate Hugo Content.
// func (Gsheet) RunHugo() {
// 	getData("hugo")
// }

// func getData(option string) {
// 	fmt.Println("Generating Data...")
// 	sh.Run(Name, "-option="+option)
// 	fmt.Println("Done.")
// }
