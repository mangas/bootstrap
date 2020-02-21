package dep

import (
	"os"

	mageutil "github.com/getcouragenow/bootstrap/ci/os/util"
	"github.com/magefile/mage/mg"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// OS namespace
type OS mg.Namespace

// InstallWindows install windows Dependencies.
func (OS) InstallWindows() {
	mageutil.Windows{}.InstallDependency()

}

// InstallLinux install linux Dependencies.
func (OS) InstallLinux() {
	mageutil.Linux{}.InstallDependency()
	// fmt.Println("Not yet ready")
}

// InstallMac install Mac Dependencies.
func (OS) InstallMac() {
	mageutil.Mac{}.InstallDependency()
}
