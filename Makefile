export GOPATH := $(PWD)
export GOBIN := $(PWD)/bin

PACKAGES := $(shell env GOPATH=$(GOPATH) go list ./... | grep -v "home")

all: test check

get:
	go get -v $(PACKAGES)

install:
	go install -v $(PACKAGES)

build:
	go build -v -o bin/talkserial-linux64 src/*.go

release:
	env GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -v -o bin/talkserial-arm7 src/*.go

release5:
	env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -v -o bin/talkserial-arm5 src/*.go
