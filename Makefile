# Makefile
GOPATH := $(CURDIR)
export GOPATH

all: wacc_19

wacc_19:
	go build src/compiler.go
