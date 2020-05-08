# install kubernetes toiling for dev
# https://minikube.sigs.k8s.io/docs/start/

GO_OS 		:= $(shell go env GOOS)
GO_ARCH		:= $(shell go env GOARCH)

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


print: ## print

	@echo
	@echo GO_OS: $(GO_OS)
	@echo GO_OS: $(GO_ARCH)
	@echo


os-dep: ##os-dep
	@echo Installing OS required tools
	
ifeq ($(GO_OS), linux)
	@echo Linux detected
endif

ifeq ($(GO_OS), darwin)
	@echo Mac detected
	brew install minikube
	brew install hyperkit
	# set hyper as default driver
	minikube config set vm-driver hyperkit
	# reset minikube to use the hyper driver
	minikube delete
	minikube start
endif

ifeq ($(GO_OS), windows)
	@echo Windows detected
endif

install-skaffold:
	# https://github.com/GoogleContainerTools/skaffold
	