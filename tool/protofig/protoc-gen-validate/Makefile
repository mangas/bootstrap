include ../../../boilerplate/tool/tools.mk

override LIB_NAME=protoc-gen-validate
override LIB_BIN_NAME=protoc-gen-validate
override LIB_ORG=envoyproxy
override LIB=github.com/$(LIB_ORG)/$(LIB_NAME)
override LIB_BRANCH=master
override LIB_TAG=v0.3.0-java
override LIB_FSPATH=$(GOPATH)/src/$(LIB)
override LIB_BIN_FSPATH=$(GOPATH)/bin/$(LIB_BIN_NAME)

## build
build: go-get-dir
	cd $(LIB_FSPATH) && GO111MODULE=off $(MAKE) build

go-get-dir:
	GO111MODULE=off go get -d $(LIB)

## Clean generated binary and source
build-clean:
	@rm -rf $(LIB_BIN_FSPATH)
	@rm -rf $(LIB_FSPATH)

