dep-nats-jsm:
	#mkdir $(GOPATH)/src/github.com/nats-io
	cd $(GOPATH)/src/github.com/nats-io && git clone git@github.com:nats-io/jetstream.git
	#GO111MODULE=off go get -u github.com/nats-io/jetstream
dep-nats:
	# so that we get jetstream embedded.
	cd $(GOPATH)/src/github.com/nats-io && git clone https://github.com/nats-io/nats-server.git
	cd $(GOPATH)/src/github.com/nats-io/nats-server && git checkout jetstream
	cd $(GOPATH)/src/github.com/nats-io/nats-server && go build
	cd $(GOPATH)/src/github.com/nats-io/nats-server && 	ls -l nats-server

	# https://github.com/ripienaar/piper
	# For a User where they want their state to be on ALL
	# their devices and synchronised.
	# e.g clipboard.

build:
	cd $(GOPATH)/src/github.com/nats-io/jetstream/jsm && go install 

	cd $(GOPATH)/src/github.com/nats-io/nats-server && go install

code:
	cd $(GOPATH)/src/github.com/nats-io && code .

run-server:
	nats-server -js -sd $(PWD)/tmp/test
