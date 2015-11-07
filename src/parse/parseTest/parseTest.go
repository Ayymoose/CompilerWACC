package main

import (
	"fmt"

	grammar "github.com/nanaasiedu/wacc_19/src/grammar"
	. "github.com/nanaasiedu/wacc_19/src/parse"
)

func main() {
	p := ConstructParser(testStream2())
	pass, errorMsgs := p.Parse()

	if !pass {
		fmt.Println("Syntax error ------ ")

		for _, errorMsg := range errorMsgs {
			fmt.Printf("%s\n", errorMsg)
		}

	} else {
		fmt.Println("Parse was successful!!!")
	}

}

func testStream1() []Token {
	t1 := Token{grammar.IDENTIFIER, "abc", 0, 0}
	t2 := Token{grammar.IDENTIFIER, "nman", 1, 5}
	tokenStream := []Token{t1, t2}

	return tokenStream
}

func testStream2() []Token {
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.IDENTIFIER, "nman", 1, 2}
	t3 := Token{grammar.END, grammar.Token_strings[grammar.END], 2, 0}
	tokenStream := []Token{t1, t2, t3}

	return tokenStream
}

func testStream3() []Token {
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.END, grammar.Token_strings[grammar.END], 2, 0}
	tokenStream := []Token{t1, t2, t3}

	return tokenStream
}
