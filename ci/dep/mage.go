package dep

import (
	"os"

	mageutil "github.com/getcouragenow/bootstrap/_mage/util"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// install windows Dependencies.
func Dependencies_InstallWindows() {
	mageutil.Windows{}.InstallDependency()

}

// install Mac Dependencies.
func Dependencies_InstallMac() {
	mageutil.Mac{}.InstallDependency()
}
