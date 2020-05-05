include ../../boilerplate/makefiles/tools.mk

override LIB_BIN_NAME=protofig
SUFFIX ?= $(shell git config user.name)

PHONY: help

baseproto:
	protoc -I config/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=":config" \
		--validate_out="lang=go:config/" \
		config/baseproto.proto

## Generate protobuf golang code
proto-go: baseproto
	protoc -I config/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=":config/" \
		--validate_out="lang=go:config/" \
		--msgnames_out="config/" \
		--configdefault_out=outfile=$(SUFFIX):"config" \
		config/config.proto

## Test protofig output in output directory using supplied test.json
test-json:
	go run main.go -f ./test.json -o ./output -u $(SUFFIX)

## Show help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
