package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"parser"
)

const SYNTAX_ERROR = 100
const SEMANTIC_ERROR = 200

// wacc_examples\valid\expressions\andExpr.wacc
// wacc_examples\valid\basic\exit\exitBasic.wacc
// wacc_examples\valid\expressions\andExpr.wacc
// wacc_examples\valid\if\if6.wacc
// wacc_examples\valid\while\fibonacciFullIt.wacc
// wacc_examples\valid\function\simple_functions\functionDeclaration.wacc
// wacc_examples\valid\function\simple_functions\functionReturnPair.wacc
// wacc_examples\valid\function\simple_functions\functionManyArguments.wacc
// wacc_examples\invalid\syntaxErr\function\noBodyAfterFuncs.wacc
// wacc_examples\invalid\syntaxErr\function\functionNoReturn.wacc
// wacc_examples/valid/pairs/createPair03.wacc

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
