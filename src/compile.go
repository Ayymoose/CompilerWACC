package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"parser"
)

const SYNTAX_ERROR = 100
const SEMANTIC_ERROR = 200

// "wacc_examples\valid\expressions\andExpr.wacc"

func main() {
	file := os.Args[1] // index 1 is file path
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	root, err := parser.ParseFile(file, s)
	if err != nil {
		os.Exit(SYNTAX_ERROR)
	}
	fmt.Println(root)
}
