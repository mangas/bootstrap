package flutter

import (
	"runtime"

	"github.com/fatih/color"
	dep "github.com/getcouragenow/bootstrap/ci/dep"
	h "github.com/getcouragenow/bootstrap/ci/hover"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// GoFlutter namespace
type GoFlutter mg.Namespace

// Install Flutter an go dependencies.
func (GoFlutter) Install() {
	switch runtime.GOOS {
	case "linux":
		dep.Dep{}.InstallLinux()
		break
	case "windows":
		dep.Dep{}.InstallWindows()
		break
	case "mac":
		dep.Dep{}.InstallMac()
		break
	}
	// install hover
	h.Hover{}.Install()

	// check
	check()
}

// Check check flutter config
func check() {
	if err := sh.RunV("go"); err != nil {
		color.Red(err.Error())
	}
	color.Green("goleng")

	if err := sh.RunV("hover"); err != nil {
		color.Red(err.Error())
	}
	color.Green("hover")

	if err := sh.RunV("flutter", "doctor", "-v"); err != nil {
		color.Red(err.Error())
	}
	color.Green("flutter")
}
