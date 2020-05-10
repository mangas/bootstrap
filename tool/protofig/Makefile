include ../../boilerplate/tool/tools.mk

override LIB_BIN_NAME=protofig
SUFFIX ?= $(shell git config user.name)

PHONY: help

baseproto:
	protoc -I api/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=":generated/go" \
		--dart_out=":generated/dart" \
		--validate_out="lang=go:generated/go/" \
		api/baseproto.proto

## Generate protobuf golang & dart code
proto: baseproto
	protoc -I api/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=":generated/go/" \
		--dart_out=":generated/dart/" \
		--validate_out="lang=go:generated/go/" \
		--configdef_out=output_format=go,suffix=$(SUFFIX):"generated/go" \
		--configdef_out=output_format=dart,suffix=$(SUFFIX):"generated/dart" \
		--configdef_out=output_format=json,suffix=$(SUFFIX):"generated/" \
		api/config.proto

## Test protofig output in output directory using supplied test.json
test-json:
	mkdir -p ./output
	go run main.go -f ./test.json -o ./output -u $(SUFFIX)
