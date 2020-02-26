package util

import (
	"errors"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	macEnv = []string{
		"GO111MODULE=on",
		"GOPATH=$HOME/workspace/go",
		"GOBIN=$GOPATH/bin",
		"GOROOT=/usr/local/opt/go/libexec",
		"FLUTTER_PATH=$HOME/flutter/bin",
		"FLUTTER_ROOT=$HOME/fvm/master",
		"JAVA_HOME=/Library/Java/JavaVirtualMachines/adoptopenjdk-12.0.2.jdk/Contents/Home",
		"ANDROID_SDK=$HOME/Library/Android/sdk",
		"ANDROID_HOME=$ANDROID_SDK",
		"ANDROID_NDK=$ANDROID_SDK/ndk-bundle",
		"ANDROID_PLATFORM_TOOLS=$ANDROID_SDK/platform-tools",
		"ANDROID_TOOLS=$ANDROID_SDK/tools",
		"PATH=$PATH:$FLUTTER_PATH:$GOROOT/bin:$JAVA_HOME/bin:$HOME/.pub-cache/bin",
	}
)

// Mac namespace
type Mac mg.Namespace

// InstallDependency install dep
func (Mac) InstallDependency() {
	if err := sh.Run("xcode-select", "--install"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("/usr/bin/ruby", "-e", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "upgrade"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "install", "git"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("xcode-select", "--install"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "cask", "install visual-studio-code"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "install", "protobuf"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "install", "gcc"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "install", "--HEAD", "libimobiledevice"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "install", "ideviceinstaller", "ios-deploy", "cocoapods"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("pod", "setup"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "cask", "install", "adoptopenjdk"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("brew", "cask", "install", "android-studio"); err != nil {
		color.Red(err.Error())
	}
	if err := sh.Run("cd", "$(HOME)/workspace", "&&", "git", "clone", "-b", "master", "https://github.com/flutter/flutter.git"); err != nil {
		color.Red(err.Error())
	}

	if err := sh.RunV("./flutter/bin/flutter", "--version"); err != nil {
		color.Red(err.Error())
	}
	setEnvVarsMac(".bashrc")
}

func setEnvVarsMac(envFile string) error {
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
