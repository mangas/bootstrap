package dep

import (
	"os"

	mageutil "github.com/getcouragenow/bootstrap/ci/dep/util"
	"github.com/magefile/mage/mg"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Dep namespace
type Dep mg.Namespace

// InstallWindows install windows Dependencies.
func (Dep) InstallWindows() {
	mageutil.Windows{}.InstallDependency()

}

// InstallLinux install linux Dependencies.
func (Dep) InstallLinux() {
	mageutil.Linux{}.InstallDependency()
	// fmt.Println("Not yet ready")
}

// InstallMac install Mac Dependencies.
func (Dep) InstallMac() {
	mageutil.Mac{}.InstallDependency()
}
