LIB_NAME=bootstrap
LIB_ORG=getcouragenow
LIB=github.com/${LIB_ORG}/$(LIB_NAME)
LIB_BRANCH=master
LIB_TAG=master
LIB_FSPATH=$(GOPATH)/src/$(LIB)

LIB_BIN_NAME=bs
LIB_BIN_FSPATH=$(GOPATH)/bin/$(LIB_BIN_NAME)

print: ## print
	@echo
	@echo $(OS)
	@echo LIB_NAME: $(LIB_NAME)
	@echo LIB: $(LIB)
	@echo LIB_BRANCH: $(LIB_BRANCH)
	@echo LIB_BRANCH: $(LIB_TAG)
	@echo LIB_FSPATH: $(LIB_FSPATH)

	@echo
	@echo LIB_BIN_NAME: $(LIB_BIN_NAME)
	@echo LIB_BIN_FSPATH: $(LIB_BIN_FSPATH)
	@echo

git-clone:
	mkdir -p $(LIB_FSPATH)
	cd $(LIB_FSPATH) && cd .. && rm -rf $(LIB_NAME) && git clone ssh://git@$(LIB).git
git-clone-master: git-clone ## git-clone-master
	cd $(LIB_FSPATH) && git checkout $(LIB_BRANCH)
git-clone-tag: git-clone ## git-clone-tag
	cd $(LIB_FSPATH) && git checkout tags/$(LIB_TAG)
	cd $(LIB_FSPATH) && git status
git-pull:
	cd $(LIB_FSPATH) && git pull
git-clean:
	rm -rf $(LIB_FSPATH)

code:
	code $(LIB_FSPATH)

build: ## build
	# builds into GO BIN
	go build -o $(LIB_BIN_FSPATH) .

build-clean: ## build-clean
	rm $(LIB_BIN_FSPATH)

