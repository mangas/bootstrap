
include ../Makefile

SAMPLE_NAME=microservices
SAMPLE_FSPATH=$(LIB_FSPATH)/examples/$(SAMPLE_NAME)

ex-print: ## ex-print
	@echo
	@echo SAMPLE_NAME: $(SAMPLE_NAME)
	@echo SAMPLE_FSPATH: $(SAMPLE_FSPATH)
	@echo
ex-code: ## ex-code
	code $(SAMPLE_FSPATH)

ex-run: ## ex-run
	cd $(SAMPLE_FSPATH) && skaffold dev

ex-test: ## ex-test
	curl $(minikube service leeroy-web --url)





