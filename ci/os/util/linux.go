package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/rotisserie/eris"
)

var (
	golangFile    = "go1.13.4.linux-amd64.tar.gz"
	protocVersion = "3.10.1"
	goPath        = os.Getenv("GOPATH")
	protocFile    = "protoc-" + protocVersion + "-linux-x86_64.zip"
	linuxEnv      = []string{
		"GO111MODULE=on",
		"GOPATH=$HOME/go",
		"GOBIN=$HOME/go/bin",
		"GOROOT=/usr/local/go",
		"FLUTTER_PATH=$HOME/flutter/bin",
		"JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64",
		"ANDROID_SDK=$HOME/Library/Android/sdk",
		"ANDROID_HOME=$ANDROID_SDK",
		"ANDROID_NDK=$ANDROID_SDK/ndk-bundle",
		"ANDROID_PLATFORM_TOOLS=$ANDROID_SDK/platform-tools",
		"ANDROID_TOOLS=$ANDROID_SDK/tools",
		"PATH=$PATH:$FLUTTER_PATH:$GOROOT/bin:$GOBIN:$JAVA_HOME/bin",
	}
)

// Linux namespace
type Linux mg.Namespace

func execute(cmd string, args ...string) error {

	c := exec.Command(cmd, args...)
	if err := c.Run(); err != nil {
		return err
	}
	fmt.Println(cmd, args)
	return nil
}

func checkExeExists(exe string) (err error) {
	_, err = exec.LookPath(exe)
	if err != nil {
		return eris.Wrapf(err, "didn't find '%s' executable\n", exe)
	}
	return nil
}

// InstallDependency install linux dependencies
func (Linux) InstallDependency() {
	// deps
	sh.RunV("sudo", "apt-get", "upgrade")
	sh.RunV("sudo", "apt-get", "install", "curl")
	sh.RunV("sudo", "apt-get", "install", "snap")
	sh.RunV("sudo", "apt-get", "install", "unzip")
	sh.RunV("sudo", "apt-get", "install", "git")
	sh.RunV("sudo", "apt-get", "install", "openssh-server")

	// install gcc
	sh.RunV("sudo", "apt-get", "install", "gcc")

	// Golang
	golangFile := "go1.13.4.linux-amd64.tar.gz"
	sh.RunV("curl", "https://dl.google.com/go/"+golangFile, "-o", golangFile)
	sh.RunV("tar", "-C", "/usr/local/", "-xzf", golangFile)
	sh.RunV("rm", "-f", golangFile)

	// install protoc
	sh.RunV("curl", "https://github.com/protocolbuffers/protobuf/releases/download/v"+protocVersion+"/protoc-"+protocVersion+"-linux-x86_64.zip", "-L", "-o", "protoc-"+protocVersion+"-linux-x86_64.zip")
	sh.RunV("unzip", "-o", "protoc-"+protocVersion+"-linux-x86_64.zip", "-d", goPath)
	sh.RunV("sudo", "rm", "-f", protocFile)

	// install jdk
	sh.RunV("sudo", "apt-get", "install", "openjdk-8-jdk")

	// Flutter github
	sh.RunV("git", "clone", "-b", "https://github.com/flutter/flutter.git")
	// ./flutter/bin/flutter --version

	setEnvVarsLinux(".bashrc")
}

func setEnvVarsLinux(envFile string) error {
	f, err := os.OpenFile(path.Join(os.Getenv("HOME"), envFile), os.O_CREATE|os.O_WRONLY, 0660)

	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 2)
	if err != nil {
		return err
	}

	for _, env := range linuxEnv {
		_, err = f.Write([]byte("\nexport " + env + "\n"))
		if err != nil {
			return errors.New("Error to export env variable: " + err.Error())
		}
	}
	return nil
}
