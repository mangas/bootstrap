LIB_NAME=installer
LIB=github.com/jpillora/$(LIB_NAME)
LIB_FSPATH=$(GOPATH)/src/$(LIB)
LIB_BINFSPATH=$(GOPATH)/bin/$(LIB_NAME)
print:
	@echo $(LIB)
	@echo $(LIB_BINFSPATH)
dep:
	GO111MODULE=off go get -u $(LIB)

code:
	cd $(LIB_FSPATH) && code .
build:
	cd  $(LIB_FSPATH) && go build -o $(LIB_BINFSPATH) .
run:
	$(LIB_NAME)