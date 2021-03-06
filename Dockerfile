# WARNING: THIS FILE IS MANAGED IN THE 'BOILERPLATE' REPO AND COPIED TO OTHER REPOSITORIES.
# ONLY EDIT THIS FILE FROM WITHIN THE 'LYFT/BOILERPLATE' REPOSITORY:
# 
# TO OPT OUT OF UPDATES, SEE https://github.com/lyft/boilerplate/blob/master/Readme.rst

FROM golang:1.13.3-alpine3.10 as builder
RUN apk add git openssh-client make curl

# COPY only the go mod files for efficient caching
COPY go.mod go.sum /go/src/github.com/lyft/bootstrap/
WORKDIR /go/src/github.com/lyft/bootstrap

# Pull dependencies
RUN go mod download

# COPY the rest of the source code
COPY . /go/src/github.com/lyft/bootstrap/

# This 'linux_compile' target should compile binaries to the /artifacts directory
# The main entrypoint should be compiled to /artifacts/bootstrap
RUN make linux_compile

# update the PATH to include the /artifacts directory
ENV PATH="/artifacts:${PATH}"

# This will eventually move to centurylink/ca-certs:latest for minimum possible image size
FROM alpine:3.10
COPY --from=builder /artifacts /bin

RUN apk --update add ca-certificates

CMD ["bootstrap"]
