package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"grammar"
	"parse"
)

func main() {
	var tokens []parse.Token
	//	b, err := ioutil.ReadFile("wacc_examples/valid/variables/capCharDeclaration.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/function/nested_functions/fibonacciFullRec.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/syntaxErr/expressions/missingOperand1.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/syntaxErr/basic/unescapedChar.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/oliver.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/basic/skip/comment.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/function/simple_functions/negFunction.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/scope/indentationNotImportant.wacc")
	b, err := ioutil.ReadFile("wacc_examples/valid/if/whitespace.wacc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	lexer := parse.Lex("Something", s)
	for item := range lexer.Items {
		//	fmt.Println(lexer.TokenLocation(item))
		tokens = append(tokens, item)
	}
	fmt.Println(tokens)
	fmt.Println("\n------ Completed Lexing ------\n")
	if tokens[len(tokens)-1].Typ == grammar.ERROR {
		return
	}
	parser := lexer.ConstructParser(tokens)
	passed, errs := parser.Parse()
	fmt.Println("Parsing Successful:", passed)
	fmt.Println("Parsing errors:", errs)
	fmt.Println("\n------ Completed Parsing ------\n")
}
