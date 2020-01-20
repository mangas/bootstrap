LIB_NAME=sync-ssh-keys
LIB=github.com/samber/$(LIB_NAME)
LIB_BRANCH=master
LIB_FSPATH=$(GOPATH)/src/$(LIB)

LIB_BIN_NAME=$(LIB_NAME)
LIB_BIN_FSPATH=$(GOPATH)/bin/$(LIB_BIN_NAME)

SRC=main.go output.go
BIN=sync-ssh-keys
DOCKER_IMAGE=samber/sync-ssh-keys
VERSION=$(shell cat VERSION)
BUILD_ID := $(shell git rev-parse --short HEAD)

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD_ID)"



print:
	@echo ssh keys sync lib

	@echo
	@echo LIB_NAME: $(LIB_NAME)
	@echo LIB: $(LIB)
	@echo LIB_BRANCH: $(LIB_BRANCH)
	@echo LIB_FSPATH: $(LIB_FSPATH)
	@echo

	@echo
	@echo LIB_BIN_NAME: $(LIB_BIN_NAME)
	@echo LIB_BIN_FSPATH: $(LIB_BIN_FSPATH)
	@echo


git-clone:
	mkdir -p $(LIB_FSPATH)
	cd $(LIB_FSPATH) && cd .. && rm -rf $(LIB_NAME) && git clone ssh://git@$(LIB).git
	cd $(LIB_FSPATH) && git checkout $(LIB_BRANCH)
code:
	code $(LIB_FSPATH)

draw:
	cd $(LIB_FSPATH) && 


## GO

all: build

dep:
	cd $(LIB_FSPATH) && GO111MODULE=on go mod download

build:
	cd $(LIB_FSPATH) && CGO_ENABLED=0 GO111MODULE=on go build -v $(LDFLAGS) -o $(LIB_BIN_FSPATH) $(SRC)

run-build:
	$(LIB_BIN_FSPATH) --github-username winwisely99

run-dev:
	cd $(LIB_FSPATH) && go run -v $(LDFLAGS) ${SRC} --github-username samber