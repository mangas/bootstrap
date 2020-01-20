package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rotisserie/eris"
)

func checkExeExists(exe string) (err error) {
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

func runEnvironTest(envValue string) (err error) {
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

func main() {
	checkExeExists("ls")
	checkExeExists("ls2")

	runEnvironTest("")
	runEnvironTest("test value")
}
