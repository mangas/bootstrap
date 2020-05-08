# WARNING: THIS FILE IS MANAGED IN THE 'BOOTSTRAP' REPO AND COPIED TO OTHER REPOSITORIES.
# ONLY EDIT THIS FILE FROM WITHIN THE 'CGN/BOILERPLATE' REPOSITORY:
# 
# TO OPT OUT OF UPDATES, SEE https://github.com/lyft/boilerplate/blob/master/Readme.rst

.PHONY: docker_build
docker_build:
	IMAGE_NAME=$$REPOSITORY ./boilerplate/lyft/docker_build/docker_build.sh

.PHONY: dockerhub_push
dockerhub_push:
	IMAGE_NAME=lyft/$$REPOSITORY REGISTRY=docker.io ./boilerplate/lyft/docker_build/docker_build.sh


### GIT-FORK

#See: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork
## Open the Upstream git repo.
git-upstream-open: 
	open https://github.com/$(UPSTREAM_ORG)/$(REPO_NAME).git 

## Open the Fork git repo ( if using from a fork).
git-fork-open: 
	open https://github.com/$(FORK_ORG)/$(REPO_NAME).git

## Setup a forked repo to track the Upstream git repo.
git-fork-setup: 
	# Pre: you git forked ( via web) and git cloned (via ssh)
	# add upstream repo
	git remote add upstream git://github.com/$(UPSTREAM_ORG)/$(REPO_NAME).git

## Merge on any changes Upstream to your local repo.
git-fork-catchup: 
	# This fetches the branches and their respective commits from the upstream repository.
	git fetch upstream 

	# This brings your fork's master branch into sync with the upstream repository, without losing your local changes.
	git merge upstream/master

	# then in VSCODE just sync to push upwards to your fork.
