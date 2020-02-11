package bootstrap

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/getcouragenow/bootstrap/ci/utils"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
	"github.com/pkg/errors"
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

// Install hugo
func Hugo_Install() error {
	// mg.Deps(installHugo)
	return installHugo()

	// err = sh.Run("hugo")
	// if err != nil {
	// 	return err
	// }

	// return nil
}

// Uninstall hugo
func Hugo_Uninstall() error {
	mg.Deps(uninstallHugo)

	err := sh.Run("hugo")
	if err == nil {
		return nil
	}

	return err
}

// clone caddy
func Caddy_clone() error {

	fmt.Println("cloning hugo...")
	dir := filepath.Join(tempDir, "caddy")
	if err := sh.Run("git", "clone", "-q", "https://github.com/mholt/caddy", dir); err != nil {
		return errors.Wrap(err, "failed to clone caddy")
	}
	return nil

}

func installHugo() error {
	// TO go bin
	verbose("installing hugo...")

	// TODO - try using to git clone so it matches makes TAGs, etc and we dont get messed up with go mod fun.
	sh.RunWith(utils.FlagEnv(packageName), goexe, "install", "github.com/gohugoio/hugo", "-tags", "extended")

	//sh.Run("go", "get", "github.com/gohugoio/hugo")
	//sh.RunWith(flagEnv(), goexe, "build", "-tags", "extended", "github.com/gohugoio/hugo")
	//sh.Run("go", "install", "github.com/gohugoio/hugo", "-tags", "extended")
	sh.RunWith(utils.FlagEnv(packageName), goexe, "install", "github.com/gohugoio/hugo", "-tags", "extended")
	//sh.Run("which", "hugo")
	return nil
}

func uninstallHugo() {
	verbose("uninstalling hugo...")
	// go clean -i -n github.com/motemen/gor

	// cleans out bin and go mods
	//sh.Run("go", "clean", "-i", "-n", "github.com/gohugoio/hugo")
	// cleans out bin
	sh.Run("go", "clean", "-i", "github.com/gohugoio/hugo")
	sh.Run("which", "hugo")

}

func hugoPostPreview(post string) error {
	verbose(fmt.Sprintf("creating preview for post %s", post))

	err := sh.Run("hugo-post-preview", "-filename", previewPath(post), "-post", post, "-timeout", "3s")
	if err != nil {
		return err
	}

	return nil
}

func changedPosts() ([]string, error) {
	var posts []string

	files, err := ioutil.ReadDir(postsDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		file := f.Name()
		changed, err := target.Path(previewPath(file), postPath(file))
		if err != nil {
			return nil, err
		}

		if changed {
			posts = append(posts, file)
		}
	}

	return posts, nil
}

func previewPath(post string) string { return fmt.Sprintf("%s/p_%s.png", previewsDir, post) }

func postPath(post string) string { return fmt.Sprintf("%s/%s/index.html", postsDir, post) }

func verbose(msg string) {
	if mg.Verbose() {
		fmt.Println(msg)
	}
}
