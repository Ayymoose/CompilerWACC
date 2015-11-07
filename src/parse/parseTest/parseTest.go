package main

import (
	"fmt"

	grammar "github.com/nanaasiedu/wacc_19/src/grammar"
	. "github.com/nanaasiedu/wacc_19/src/parse"
)

func main() {
	p := ConstructParser(testStream6())
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

// PARSE MAY FAIL IF PARSE FUNCTION IS INCOMPLETE (An error warning will display if
// this is the case)
func testStream1() []Token {
	/* abc
	   nman */
	// ERROR: Expected begin

	t1 := Token{grammar.IDENTIFIER, "abc", 0, 0}
	t2 := Token{grammar.IDENTIFIER, "nman", 1, 5}
	tokenStream := []Token{t1, t2}

	return tokenStream
}

func testStream2() []Token {
	/* begin
	     nman
		 end */
	// ERROR: Invalid <func>/<stat>
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.IDENTIFIER, "nman", 1, 2}
	t3 := Token{grammar.END, grammar.Token_strings[grammar.END], 2, 0}
	tokenStream := []Token{t1, t2, t3}

	return tokenStream
}

func testStream3() []Token {
	/* begin
		 skip
	 end */
	// SUCCESS
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.END, grammar.Token_strings[grammar.END], 2, 0}
	tokenStream := []Token{t1, t2, t3}

	return tokenStream
}

func testStream4() []Token {
	/* begin
		 skip ; skip
	 end */
	// SUCCESS
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 1, 7}
	t4 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 2, 2}
	t5 := Token{grammar.END, grammar.Token_strings[grammar.END], 3, 0}
	tokenStream := []Token{t1, t2, t3, t4, t5}

	return tokenStream
}

func testStream5() []Token {
	/* begin
		 skip;
		 free
	 end */
	// ERROR: Invalid <stat> // <expr> expected after free
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 2, 2}
	t4 := Token{grammar.FREE, grammar.Token_strings[grammar.FREE], 3, 2}
	t5 := Token{grammar.END, grammar.Token_strings[grammar.END], 4, 0}
	tokenStream := []Token{t1, t2, t3, t4, t5}

	return tokenStream
}

func testStream6() []Token {
	/* begin
		 skip ; skip; skip; skip; skip
	 end */
	// SUCCESS
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 2, 2}
	t4 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 3, 2}
	t5 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 4, 2}
	t6 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 5, 2}
	t7 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 6, 1}
	t8 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 7, 2}
	t9 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 8, 2}
	t10 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 9, 2}
	t11 := Token{grammar.END, grammar.Token_strings[grammar.END], 10, 0}
	tokenStream := []Token{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11}

	return tokenStream
}
