package cloud

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	ClourProjectID = "winwisely-cloudrun-googlesheet"
)

type Cloud mg.Namespace

// Auth Cloud
func (Cloud) AUTH() {
	sh.RunV("gcloud", "auth", "login")
	sh.RunV("gcloud", "config", "set", "project", ClourProjectID)
	sh.RunV("gcloud", "config", "set", "run/region", "europe-west1")
}

// Open Cloud
func (Cloud) Open() {
	baseCloudURL := "https://console.cloud.google.com/"
	// run
	sh.RunV("open", baseCloudURL+"run?project="+ClourProjectID)
	// container
	sh.RunV("open", baseCloudURL+"cloud-build/builds?project="+ClourProjectID)
	// store
	sh.RunV("open", baseCloudURL+"storage/browser?project="+ClourProjectID)
}

// Build Cloud
func (Cloud) Build() {
	sh.RunV("gcloud", "build", "submit", "--tag", "gcr.io/"+ClourProjectID+"/identicon-generator")
}

// Deploy Cloud
func (Cloud) Deploy() {
	sh.RunV("gcloud", "beta", "deploy", "--image", "gcr.io/"+ClourProjectID+"/identicon-generator", "--platform", "managed")
}
