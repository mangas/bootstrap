package tools

import (
	"github.com/getcouragenow/bootstrap/ci/tools/gofish"
	"github.com/getcouragenow/bootstrap/ci/tools/hover"
	"github.com/getcouragenow/bootstrap/ci/tools/hugo"
	mageselect "github.com/getcouragenow/bootstrap/ci/tools/mage_select"
	"github.com/magefile/mage/mg"
)

// Tools namespace
type Tools mg.Namespace

// HoverInstall install hover.
func (Tools) HoverInstall() {
	hover.Install()
}

// HoverUninstall uninstall hover.
func (Tools) HoverUninstall() {
	hover.Uninstall()
}

// HugoInstall install hugo
func (Tools) HugoInstall() {
	hugo.Install()
}

// HugoUninstall install hugo
func (Tools) HugoUninstall() {
	hugo.Uninstall()
}

// MageInstall install mage
func (Tools) MageInstall() {
	hugo.Install()
}

// MageSelectInstall install mage_select
func (Tools) MageSelectInstall() {
	mageselect.Install()
}

// MageSelectUninstall uninstall mage_select
func (Tools) MageSelectUninstall() {
	mageselect.Uninstall()
}

// GofishInstall install gofish
func (Tools) GofishInstall() {
	gofish.Install()
}

// GofishUninstall uninstall gofish
func (Tools) GofishUninstall() {
	gofish.Uninstall()
}

// MageSelectUpdate uninstall mage_select
func (Tools) MageSelectUpdate() {
	mageselect.Update()
}
