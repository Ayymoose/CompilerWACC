# Makefile
GOPATH := $(CURDIR)
export GOPATH

all: wacc_19

wacc_19:
	go build src/wacc.go
#	go build src/semantics.go


clean:
	-rm *.output *.yacc.go
