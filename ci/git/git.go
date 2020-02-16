package git

import (
	"github.com/magefile/mage/sh"
)

// SetUpstream set upstream
func SetUpstream() {
	sh.RunV("git", "remote", "add", "upstream", "")
}
