package servertools

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// ServerTools namespace
type ServerTools mg.Namespace

// InstallMinikube install minikube
func (ServerTools) InstallMinikube() {
	sh.RunV("gofish", "install", "minikube")
}

// UninstallMinikube uninstall minikube
func (ServerTools) UninstallMinikube() {
	sh.RunV("gofish", "uninstall", "minikube")
}

// InstallKubectl install kubectl
func (ServerTools) InstallKubectl() {
	sh.RunV("gofish", "install", "kubectl")
}

// UninstallKubectl uninstall kubectl
func (ServerTools) UninstallKubectl() {
	sh.RunV("gofish", "uninstall", "kubectl")
}

// Installskaffold install skaffold
func (ServerTools) Installskaffold() {
	sh.RunV("gofish", "install", "skaffold")
}

// Uninstallskaffold uninstall skaffold
func (ServerTools) Uninstallskaffold() {
	sh.RunV("gofish", "uninstall", "skaffold")
}
