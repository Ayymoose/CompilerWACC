package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"grammar"
	"parse"
)

func main() {
	file := os.Args[1]
	b, err := ioutil.ReadFile(file)
	//	b, err := ioutil.ReadFile("wacc_examples/valid/variables/capCharDeclaration.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/function/nested_functions/fibonacciFullRec.wacc")
	//b, err := ioutil.ReadFile("wacc_examples/invalid/syntaxErr/expressions/missingOperand1.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/syntaxErr/basic/unescapedChar.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/oliver.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/basic/skip/comment.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/function/simple_functions/negFunction.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/scope/indentationNotImportant.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/if/whitespace.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/advanced/ticTacToe.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/IO/print/printEscChar.wacc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	var tokens []parse.Token
	lexer := parse.Lex("Something", s)
	for item := range lexer.Items {
		//	fmt.Println(lexer.TokenLocation(item))
		tokens = append(tokens, item)
	}
	//	fmt.Println(tokens)
	//	fmt.Println("\n------ Completed Lexing ------\n")
	for x, token := range tokens {
		if x%2 == 0 {
			fmt.Println()
		}
		fmt.Print(token, ", ")
	}
	if tokens[len(tokens)-1].Typ == grammar.ERROR {
		fmt.Println("#syntax_error#(Lex)")
		os.Exit(100)
	}
	parser := lexer.ConstructParser(tokens)
	passed, errs := parser.Parse()
	//	fmt.Println("Parsing Successful:", passed)
	fmt.Println("\n------ Parsing Complete ------\n")
	for _, parseError := range errs {
		fmt.Println(parseError)
	}

	if !passed {
		fmt.Println("#syntax_error#(Parse)")
		os.Exit(100)
	}

}
