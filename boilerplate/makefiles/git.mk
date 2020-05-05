include tools.mk

## Clone current repo
git-clone:
	mkdir -p $(LIB_FSPATH)
	cd $(LIB_FSPATH) && cd .. && rm -rf $(LIB_NAME) && git clone ssh://git@$(LIB).git
## Clone master branch of repo
git-clone-master: git-clone ## git-clone-master
	cd $(LIB_FSPATH) && git checkout $(LIB_BRANCH)
## Clone specific tag of repo
git-clone-tag: git-clone ## git-clone-tag
	cd $(LIB_FSPATH) && git checkout tags/$(LIB_TAG)
	cd $(LIB_FSPATH) && git status
## Pulls local repo from upstream
git-pull:
	cd $(LIB_FSPATH) && git pull
## Remove current local git directory
git-clean:
	rm -rf $(LIB_FSPATH)
