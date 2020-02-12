package main

import (
	"fmt"
	"os"

	u "github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/mage"
	"github.com/magefile/mage/parse"
)

func main() {
	u.GetPath()
	fmt.Println(mage.Magefiles(".", "linux", "amd64", "go", os.Stderr, false))
	fmt.Println(mage.GenerateMainfile("xxx", "./xxx", &parse.PkgInfo{}))
	os.Exit(mage.Main())
}
