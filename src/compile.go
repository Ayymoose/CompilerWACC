package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"grammar"
	"parse"
)

const SYNTAX_ERROR = 100
const SEMANTIC_ERROR = 200

func main() {
	file := os.Args[1] // index 1 is file path
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	var tokens []parse.Token
	lexer := parse.Lex("Something", s)
	for item := range lexer.Items {
		tokens = append(tokens, item)
	}

	for x, token := range tokens {
		if x%2 == 0 {
			fmt.Println()
		}
		fmt.Print(token, ", ")
	}
	if tokens[len(tokens)-1].Typ == grammar.ERROR {
		fmt.Println("#syntax_error#(Lex)")
		os.Exit(SYNTAX_ERROR)
	}
	parser := lexer.ConstructParser(tokens)
	passed, errs := parser.Parse()
	fmt.Println("\n------ Parsing Complete ------\n")
	for _, parseError := range errs {
		fmt.Println(parseError)
	}

	if !passed {
		fmt.Println("#syntax_error#(Parse)")
		os.Exit(SYNTAX_ERROR)
	}

}
