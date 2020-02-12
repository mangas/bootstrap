package build

import (
	u "github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var projectName = "Bootstrap"

// Build namespace
type Build mg.Namespace

// Windows build binary for windows(amd64)
func (Build) Windows() {
	build("windows", "amd64")
}

// Mac build binary for mac darwin (amd64)
func (Build) Mac() {
	build("darwin", "amd64")
}

// Linux build binary for linux (amd64)
func (Build) Linux() {
	build("linux", "amd64")
}

func build(goos, goarch string) {
	env := u.FlagEnv(projectName)
	env["GOOS"] = goos
	env["GOARCH"] = goarch
	sh.RunWith(env, "go", "build", ".")
}
