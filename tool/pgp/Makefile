# https://gist.github.com/eliquious/9e96017f47d9bd43cdf9


# git
REPO_NAME=$(notdir $(shell pwd))
UPSTREAM_ORG=getcouragenow
FORK_ORG=$(shell basename $(dir $(abspath $(dir $$PWD))))

# remove the "v" prefix
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


print:
	@echo
	@echo FORK_ORG: $(FORK_ORG)

	@echo UPSTREAM_ORG: $(UPSTREAM_ORG)
	@echo REPO_NAME: $(REPO_NAME)
	@echo

build: ## Build
	go build -o goencrypt main.go


gen: ## Produces keyname.pubkey and keyname.privkey
	./goencrypt keygen keyname