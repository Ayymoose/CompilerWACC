# Makefile
GOPATH := $(CURDIR)
export GOPATH

all: wacc_19

wacc_19:
	go build src/compile.go
#	go build src/semantics.go
