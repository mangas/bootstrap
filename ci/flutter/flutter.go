package flutter

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"

	dep "github.com/getcouragenow/bootstrap/ci/dep"
	h "github.com/getcouragenow/bootstrap/ci/hover"
	mg "github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	LIB_FSPATH    = ""
	SAMPLE_FSPATH = ""
	PROTO_OUTPUT  = ""
	HOME_PATH     = ""

	GOPATH      = ""
	LIB_NAME    = "main"
	LIB         = "github.com/winwisely99/" + LIB_NAME
	LIB_BRANCH  = "master"
	SAMPLE_NAME = "client"

	GO_OS       = ""
	GO_ARCH     = ""
	GIT_VERSION = ""

	AND_KEYSTORE_APP_NAME = "main"
	AND_KEYSTORE_FSPATH   = ""
)

// git rev-parse --show-prefix
// Flutter namespace
type Flutter mg.Namespace

func init() {
	p, err := sh.Output("git", "rev-parse", "--showprefix")
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
		return
	}

	if p != "" {
		fmt.Println(fmt.Errorf("You should be in root of your repository"))
		return
	}

	usr, _ := user.Current()
	home := usr.HomeDir
	pwd, _ := os.Getwd()

	GO_OS = runtime.GOOS
	GO_ARCH = runtime.GOARCH
	GIT_VERSION, _ = sh.Output("git", "describe", "--tags")

	AND_KEYSTORE_FSPATH = path.Join(pwd, "_key", AND_KEYSTORE_APP_NAME)

	if os.Getenv("GOPATH") == "" {
		GOPATH, _ = sh.Output("go", "env", "GOPATH")
		sh.RunV("export", GOPATH)
	}
	if GO_OS == "windows" {
		LIB_FSPATH = "subst \\,/,$(subst C:\\,/c/," + GOPATH + "/src/" + LIB
		SAMPLE_FSPATH = "subst \\,/,$(subst C:\\,/c/," + LIB_FSPATH + "/" + SAMPLE_NAME
		PROTO_OUTPUT = "subst /c/,C:\\,subst \\,/," + LIB_FSPATH
		HOME_PATH = "subst \\,/,subst C:\\,/c/," + home
	} else {
		LIB_FSPATH = GOPATH + "/src/" + LIB
		SAMPLE_FSPATH = LIB_FSPATH + "/" + SAMPLE_NAME
		HOME_PATH = home
	}
}

// InstallForMobile install Flutter mobile dependencies.
func (Flutter) InstallForMobile() {
	install(false)
}

// InstallForDesktop install Flutter mobile and desktop dependencies.
func (Flutter) InstallForDesktop() {
	install(true)
}

func install(desktop bool) {

	switch runtime.GOOS {
	case "linux":
		dep.Dep{}.InstallLinux()
		break
	case "windows":
		dep.Dep{}.InstallWindows()
		break
	case "mac":
		dep.Dep{}.InstallMac()
		break
	}
	// install hover
	if desktop {
		h.Hover{}.Install()
	}
}

// CheckConfig check flutter config
func (Flutter) CheckConfig() {
	sh.RunV("flutter", "config")
	sh.RunV("flutter", "doctor", "-v")
	sh.RunV("flutter", "devices")
}

// CheckDeskConfig check flutter desktop config
func (Flutter) CheckDeskConfig() {
	sh.RunV("hover", "-h")
}
