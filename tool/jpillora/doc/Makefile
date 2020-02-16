# https://github.com/cugu/godoc2readme

LIB_NAME=godoc2readme
LIB=github.com/cugu/$(LIB_NAME)
LIB_FSPATH=$(GOPATH)/src/$(LIB)
LIB_BINFSPATH=$(GOPATH)/bin/$(LIB_NAME)
print:
	@echo $(LIB)
	@echo $(LIB_BINFSPATH)
dep:
	GO111MODULE=off go get -u $(LIB)

	# extra
	GO111MODULE=off go get -u github.com/go-zeromq/zmq4

	GO111MODULE=off go get -u github.com/go-zeromq/zyre



build:
	cd  $(LIB_FSPATH) && go build -o $(LIB_BINFSPATH) .
run:
	$(LIB_NAME)