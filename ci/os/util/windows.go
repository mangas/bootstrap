package util

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	winEnv = []string{
		"GO111MODULE on",
		"scoopApps C:\\Users\\%username%\\scoop\\apps",
		"GOPATH %scoopApps%/go",
		"GOBIN %GOPATH%/bin",
		"%PATH%;%GOBIN%",
	}
)

// Windows namespace
type Windows mg.Namespace

// PowerShell struct
type PowerShell struct {
	powerShell string
}

// New powershell
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

// Execute cmd on shell
func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

// InstallDependency install deps
func (Windows) InstallDependency() {
	var err error
	posh := New()
	//	var stdout,stderr string
	_, _, err = posh.Execute("Set-ExecutionPolicy RemoteSigned -scope CurrentUser")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop update")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install aria2")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop bucket add extras")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install git")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install which")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install make")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install vscode")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install protobuf")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install gcc")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install go")
	if err != nil {
		color.Red(err.Error())
	}

	_, _, err = posh.Execute("scoop install flutter")
	if err != nil {
		color.Red(err.Error())
	}

	setEnvVarsWindows()
}

func setEnvVarsWindows() error {

	for _, env := range winEnv {
		args := strings.Split(env, " ")
		err := sh.RunV("setx", args...)
		if err != nil {
			return errors.New("Error to export env variable: " + err.Error())
		}
	}

	return nil
}

//https://stackoverflow.com/questions/50809752/golang-invoking-powershell-exe-always-returns-ascii-characters
