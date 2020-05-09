
# This make file uses composition to keep things KISS and easy.
# In the boilerpalte make files dont do any includes, because you will create multi permutations of possibilities.



# git include
include ./boilerplate/core/help.mk
include ./boilerplate/core/os.mk
include ./boilerplate/core/gitr.mk
include ./boilerplate/core/go.mk

# example of how to override as needed
override GO_FSPATH = $(PWD)
override GO_BUILD_OUT_FSPATH = $(GOPATH)/bin/bs

STATIK_DEST = $(PWD)/statiks

.PHONY: help statiks build

## Print all settings
print: ## print

	
	$(MAKE) os-print
	
	$(MAKE) gitr-print

	$(MAKE) go-print
	

## So high
high: ## high
	@echo i wanna get...


build: 
	$(MAKE) go-build

statiks:
	@statik -src=boilerplate -ns bproot -p bproot -dest=$(STATIK_DEST) -f
	@statik -src=boilerplate/core -ns bpcore -p bpcore -dest=$(STATIK_DEST) -f
	@statik -src=boilerplate/lyft -ns bplyft -p bplyft -dest=$(STATIK_DEST) -f
	@statik -src=boilerplate/tool -ns bptool -p bptool -dest=$(STATIK_DEST) -f

