package main

import (
	"fmt"

"grammar"
"parse"
)

func main() {
	p := ConstructParser(testStream3())
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
	/* begin begin begin
		 skip;

		 begin
		 		X = fst Y
		 end

	 end end end */
	// SUCCESS
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 1, 0}
	t3 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 2, 0}

	t4 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 3, 2}
	ta := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 3, 6}

	tb := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 4, 0}

	tc := Token{grammar.IDENTIFIER, "X", 5, 2}
	td := Token{grammar.EQ, "=", 5, 3}
	te := Token{grammar.FST, "fst", 5, 4}
	tf := Token{grammar.IDENTIFIER, "Y", 5, 7}

	tg := Token{grammar.END, grammar.Token_strings[grammar.END], 6, 0}

	t5 := Token{grammar.END, grammar.Token_strings[grammar.END], 7, 0}
	t6 := Token{grammar.END, grammar.Token_strings[grammar.END], 8, 0}
	t7 := Token{grammar.END, grammar.Token_strings[grammar.END], 9, 0}
	tokenStream := []Token{t1, t2, t3, t4, ta, tb, tc, td, te, tf, tg, t5, t6, t7}

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
		 free nman
	 end */
	// SUCCESS
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 2, 2}
	t4 := Token{grammar.FREE, grammar.Token_strings[grammar.FREE], 3, 2}
	t5 := Token{grammar.IDENTIFIER, "nman", 4, 2}
	t6 := Token{grammar.END, grammar.Token_strings[grammar.END], 5, 0}
	tokenStream := []Token{t1, t2, t3, t4, t5, t6}

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

func testStream7() []Token {
	/* begin
		 skip;
		 free true;
		 print 0
	 end */
	// EXPECTED expr after print
	t1 := Token{grammar.BEGIN, grammar.Token_strings[grammar.BEGIN], 0, 0}
	t2 := Token{grammar.SKIP, grammar.Token_strings[grammar.SKIP], 1, 2}
	t3 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 2, 2}
	t4 := Token{grammar.FREE, grammar.Token_strings[grammar.FREE], 3, 2}
	t5 := Token{grammar.TRUE, "true", 4, 2}
	t6 := Token{grammar.SEMICOLON, grammar.Token_strings[grammar.SEMICOLON], 5, 2}
	t7 := Token{grammar.PRINT, grammar.Token_strings[grammar.PRINT], 6, 2}
	//t8 := Token{grammar.DIGIT, "0", 7, 2}
	t9 := Token{grammar.END, grammar.Token_strings[grammar.END], 8, 0}
	tokenStream := []Token{t1, t2, t3, t4, t5, t6, t7, t9}

	return tokenStream
}
