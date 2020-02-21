package minikube

import (
	"github.com/magefile/mage/sh"
)

// UninstallMinikube uninstall
func UninstallMinikube() {
	sh.RunV("gofish", "install", "minikube")
}

// UninstallKubectl uninstall
func UninstallKubectl() {
	sh.RunV("gofish", "install", "kubectl")
}

// InstallMinikube install
func InstallMinikube() {
	sh.RunV("gofish", "install", "minikube")
}

// InstallKubectl install
func InstallKubectl() {
	sh.RunV("gofish", "install", "kubectl")
}
