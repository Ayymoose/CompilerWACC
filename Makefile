# Makefile
GOPATH := $(CURDIR)
export GOPATH

all: build

build:
	go build src/compiler.go
