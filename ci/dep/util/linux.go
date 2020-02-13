package util

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	golangFile    = "go1.13.4.linux-amd64.tar.gz"
	protocVersion = "3.10.1"
	goPath        = os.Getenv("GOPATH")
	protocFile    = "protoc-" + protocVersion + "-linux-x86_64.zip"
)

// Linux namespace
type Linux mg.Namespace

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

	setEnvVarsLinux()
}

func setEnvVarsLinux() {
	//Todo env vars
}
