.PHONY: baseproto proto-go
	
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

baseproto:
	protoc -I baseproto/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I . \
		--go_out=":baseproto/" \
		--validate_out="lang=go:baseproto/" \
		--msgnames_out="baseproto/baseproto/" \
		baseproto/*.proto

proto-go: baseproto
	protoc -I config/ \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		-I . \
		--go_out=":config/" \
		--validate_out="lang=go:config/" \
		--msgnames_out="config/" \
		config/*.proto

