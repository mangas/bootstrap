package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/magefile/mage/sh"
	"github.com/rotisserie/eris"
)

// CheckExeExists check if binary exist
func CheckExeExists(exe string) (err error) {
	path, err := exec.LookPath(exe)
	if err != nil {
		//fmt.Printf("didn't find '%s' executable\n", exe)
		//return err
		return eris.Wrapf(err, "didn't find '%s' executable\n", exe)
	}
	fmt.Printf("'%s' executable is '%s'\n", exe, path)
	return nil
}

const (
	envName = "MY_TEST_ENV_VARIABLE"
)

// RunEnvironTest env test
func RunEnvironTest(envValue string) (err error) {
	cmd := exec.Command("go", "run", "05-print-env-helper.go")
	if envValue != "" {
		newEnv := append(os.Environ(), fmt.Sprintf("%s=%s", envName, envValue))
		cmd.Env = newEnv
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return err
	}
	fmt.Printf("%s", out)
	return nil
}

// FlagEnv flags used in build
func FlagEnv(packageName string) map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return map[string]string{
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
	}
}

// SetGoPath set goapth if not exists
func setGoENV() (gopath string, gobin string) {

	gopath = os.Getenv("GOPATH")
	gobin = os.Getenv("GOBIN")
	if gopath == "" {
		var err error
		gopath, err = sh.Output("go", "env", "GOPATH")

		if err != nil {
			log.Fatal(err)
		}
		gobin = gopath + "/bin"
	}
	return
}

// GetPath get Home, GOPATH, and GOBIN
func GetPath() (map[string]string, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	home := usr.HomeDir

	gopath, gobin := setGoENV()
	out := make(map[string]string, 3)
	out["HOME"] = home
	out["GOPATH"] = gopath
	out["GOBIN"] = gobin
	return out, nil
}

// GitClone clone repo
func GitClone(lib, libName, libFSPath, libBranch string) error {

	err := sh.RunV("mkdir", "-p", libFSPath)
	if err != nil {
		return err
	}

	err = os.Chdir(libFSPath)
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	err = sh.RunV("rm", "-rf", libName)
	if err != nil {
		return err
	}

	err = sh.RunV("git", "clone", "ssh://git@"+lib+".git")
	if err != nil {
		return err
	}

	if libBranch != "" {
		err = sh.RunV("git", "checkout", libBranch)
		if err != nil {
			return err
		}
	}

	return nil
}

// GitPull pull repo
func GitPull(libFSPath, libBranch string) error {

	err := os.Chdir(libFSPath)
	if err != nil {
		return err
	}

	if libBranch != "" {
		err = sh.RunV("git", "checkout", libBranch)
		if err != nil {
			return err
		}
	}
	return nil
}

// GitClean delete git repo
func GitClean(libFSPath string) error {
	return sh.RunV("rm", "-rf", libFSPath)
}

// CheckIfAlreadInstalled check if binary is already installed
func CheckIfAlreadInstalled(libName string) bool {
	err := CheckExeExists(libName)
	if err == nil {
		fmt.Println("%s Already installed", libName)
		return true
	}
	return false
}
