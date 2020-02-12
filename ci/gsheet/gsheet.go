package gsheet

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// mage:import
//_ "go.zenithar.org/spotigraph/.mage"

type Gsheet mg.Namespace

const (
	Name = "googlesheet"
)

// Build googlesheet tool.
func (Gsheet) Build() {
	sh.Run("go", "build", "-o", Name, ".")
}

// Run googlesheet Tests.
func (Gsheet) RunTests() {
	sh.RunV("go", "test", "-v")
}

// Clean googlesheet project.
func (Gsheet) Clean() {
	sh.RunV("go", "clean", "-cache")
	sh.RunV("go", "clean", "-modcache")
	sh.RunV("go", "clean", "-testcache")
	sh.RunV("rm", Name)
}

// Generate dump data.
func (Gsheet) RunDataDump() {
	getData("datadump")
}

// Generate language data.
func (Gsheet) RunLangDump() {
	getData("lang")
}

// Generate Hugo Content.
func (Gsheet) RunHugo() {
	getData("hugo")
}

func getData(option string) {
	fmt.Println("Generating Data...")
	sh.Run(Name, "-option="+option)
	fmt.Println("Done.")
}
