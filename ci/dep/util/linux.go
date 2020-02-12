package util

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	GOLANG_FILE    = "go1.13.4.linux-amd64.tar.gz"
	PROTOC_VERSION = "3.10.1"
	GOPATH         = os.Getenv("GOPATH")
	PROTOC_FILE    = ""
)

type Linux mg.Namespace

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
	GOLANG_FILE := "go1.13.4.linux-amd64.tar.gz"
	sh.RunV("curl", "https://dl.google.com/go/"+GOLANG_FILE, "-o", GOLANG_FILE)
	sh.RunV("tar", "-C", "/usr/local/", "-xzf", GOLANG_FILE)
	sh.RunV("rm", "-f", GOLANG_FILE)

	// install protoc
	sh.RunV("curl", "https://github.com/protocolbuffers/protobuf/releases/download/v"+PROTOC_VERSION+"/protoc-"+PROTOC_VERSION+"-linux-x86_64.zip", "-L", "-o", "protoc-"+PROTOC_VERSION+"-linux-x86_64.zip")
	sh.RunV("unzip", "-o", "protoc-"+PROTOC_VERSION+"-linux-x86_64.zip", "-d", GOPATH)
	sh.RunV("sudo", "rm", "-f", PROTOC_FILE)

	// install jdk
	sh.RunV("sudo", "apt-get", "install", "openjdk-8-jdk")

	// Flutter github
	sh.RunV("git", "clone", "-b", "https://github.com/flutter/flutter.git")
	// ./flutter/bin/flutter --version
}
