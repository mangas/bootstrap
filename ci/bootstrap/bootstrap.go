package bootstrap

import (
	"os"

	"github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/sh"
)

// global variables to streamline magefile
var tempDir string
var cwd string
var version string

const (
	postsDir    = "./docs/post"
	previewsDir = "./static"
)

const (
	packageName  = "github.com/getcouragenow/bootstrap"
	noGitLdflags = "-X $PACKAGE/common/bootstrap.buildDate=$BUILD_DATE"
)

var ldflags = "-X $PACKAGE/common/bootstrap.commitHash=$COMMIT_HASH -X $PACKAGE/common/bootstrap.buildDate=$BUILD_DATE"

// allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
var goexe = "go"

// func init() {
// 	if exe := os.Getenv("GOEXE"); exe != "" {
// 		goexe = exe
// 	}

// 	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
// 	// The default is "auto".
// 	os.Setenv("GO111MODULE", "on")
// }

func isCI() bool {
	return os.Getenv("CI") != ""
}

func buildTags() string {
	// To build the extended Hugo SCSS/SASS enabled version, build with
	// HUGO_BUILD_TAGS=extended mage install etc.
	if envtags := os.Getenv("HUGO_BUILD_TAGS"); envtags != "" {
		return envtags
	}
	return "none"

}

// Build bootstrap binary
func Bootstrap() error {
	//return sh.RunWith(flagEnv(), goexe, "build", "-ldflags", ldflags, "-tags", buildTags(), packageName)
	return sh.RunWith(utils.FlagEnv(packageName), goexe, "build", "-ldflags", ldflags, packageName)
}

// Build botstrap binary with race detector enabled
func BootstrapRace() error {
	return sh.RunWith(utils.FlagEnv(packageName), goexe, "build", "-race", "-ldflags", ldflags, "-tags", buildTags(), packageName)
}
