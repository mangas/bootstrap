package servertools

import (
	"fmt"

	"github.com/getcouragenow/bootstrap/ci/tools/gofish"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// STools namespace
type STools mg.Namespace

// IMinikube install minikube
func (STools) IMinikube() {
	verbose("installing minikube....")
	mg.Deps(gofish.Install)
	sh.RunV("gofish", "install", "minikube")
}

// UMinikube uninstall minikube
func (STools) UMinikube() {
	sh.RunV("gofish", "uninstall", "minikube")
}

// IKubectl install kubectl
func (STools) IKubectl() {
	sh.RunV("gofish", "install", "kubectl")
}

// UKubectl uninstall kubectl
func (STools) UKubectl() {
	sh.RunV("gofish", "uninstall", "kubectl")
}

// Iskaffold install skaffold
func (STools) Iskaffold() {
	sh.RunV("gofish", "install", "skaffold")
}

// Uskaffold uninstall skaffold
func (STools) Uskaffold() {
	sh.RunV("gofish", "uninstall", "skaffold")
}

func verbose(msg string) {
	if mg.Verbose() {
		fmt.Println(msg)
	}
}
