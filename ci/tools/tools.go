package tools

import (
	"github.com/getcouragenow/bootstrap/ci/tools/gofish"
	"github.com/getcouragenow/bootstrap/ci/tools/hover"
	"github.com/getcouragenow/bootstrap/ci/tools/hugo"
	"github.com/getcouragenow/bootstrap/ci/tools/mage"
	mageselect "github.com/getcouragenow/bootstrap/ci/tools/mage_select"
	"github.com/magefile/mage/mg"
)

// Tools namespace
type Tools mg.Namespace

// IHover install hover.
func (Tools) IHover() {
	hover.Install()
}

// UHover uninstall hover.
func (Tools) UHover() {
	hover.Uninstall()
}

// IHugo install hugo
func (Tools) IHugo() {
	hugo.Install()
}

// UHugo install hugo
func (Tools) UHugo() {
	hugo.Uninstall()
}

// IMage install mage
func (Tools) IMage() {
	mage.Install()
}

// IMageSelect install mage_select
func (Tools) IMageSelect() {
	mageselect.Install()
}

// UMageSelect uninstall mage_select
func (Tools) UMageSelect() {
	mageselect.Uninstall()
}

// IGofish install gofish
func (Tools) IGofish() {
	gofish.Install()
}

// MageSelectUpdate update mage_select
func (Tools) MageSelectUpdate() {
	mageselect.Update()
}
