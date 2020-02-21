package scaffold

import "github.com/magefile/mage/sh"

// Install scaffold
func Install() {
	sh.RunV("gofish", "install", "skaffold")
}
