package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	m "github.com/getcouragenow/bootstrap/ci/mage_select"
)

const (
	libName = "mage"
	lib     = "github.com/magefile/"
)

var (
	goPath    = ""
	goBIN     = ""
	libPath   = ""
	libFSPath = ""
)

func init() {
	goPath = os.Getenv("GOPATH")
	goBIN = os.Getenv("GOBIN")
	libPath = path.Join(goPath, "src", lib)
	libFSPath = path.Join(libPath, libName)
}

// func checkExeExists(exe string) (err error) {
// 	path, err := exec.LookPath(exe)
// 	if err != nil {
// 		//fmt.Printf("didn't find '%s' executable\n", exe)
// 		//return err
// 		return eris.Wrapf(err, "didn't find '%s' executable\n", exe)
// 	}
// 	fmt.Printf("'%s' executable is '%s'\n", exe, path)
// 	return nil
// }

// const (
// 	envName = "MY_TEST_ENV_VARIABLE"
// )

// func runEnvironTest(envValue string) (err error) {
// 	cmd := exec.Command("go", "run", "05-print-env-helper.go")
// 	if envValue != "" {
// 		newEnv := append(os.Environ(), fmt.Sprintf("%s=%s", envName, envValue))
// 		cmd.Env = newEnv
// 	}
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 		return err
// 	}
// 	fmt.Printf("%s", out)
// 	return nil
// }

// installMage Install mage
func installMage() {

	err := os.RemoveAll(libPath)
	if err != nil {
		log.Fatal(err)
	}

	os.Getenv("GOBIN")
	err = os.MkdirAll(libPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(libPath)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("git", "clone", "ssh://git@github.com/magefile/mage.git")
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(libFSPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getwd())
	cmd = exec.Command("go", "run", "bootstrap.go")
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func checkDeps() {
	_, err := os.Stat(goBIN + "/mage")
	if err != nil {
		fmt.Println("start installing Mage")
		installMage()
		fmt.Println("Mage installed")
	}

	_, err = os.Stat(goBIN + "/mages")
	if err != nil {
		fmt.Println("start installing Mage-select")
		m.MageSelect_Install()
		fmt.Println("Mage-select installed")
	}
}
func main() {
	checkDeps()
}
