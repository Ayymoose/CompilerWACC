package main

import (
	"fmt"

	grammar "github.com/OliWheeler/wacc_19/src/grammar"
)

func main() {
	p := constructParser(testStream2())
	pass, errorMsgs := p.parse()

	if !pass {
		fmt.Println("Syntax error ------ ")

		for _, errorMsg := range errorMsgs {
			fmt.Printf("%s\n", errorMsg)
		}

	} else {
		fmt.Println("Parse was successful!!!")
	}

}

func testStream1() []token {
	t1 := token{grammar.IDENTIFIER, "abc", 0, 0}
	t2 := token{grammar.IDENTIFIER, "nman", 1, 5}
	tokenStream := []token{t1, t2}

	return tokenStream
}

func testStream2() []token {
	t1 := token{grammar.BEGIN, grammar.token_strings[grammar.BEGIN], 0, 0}
	t2 := token{grammar.IDENTIFIER, "nman", 1, 2}
	t3 := token{grammar.END, grammar.token_strings[grammar.END], 2, 0}
	tokenStream := []token{t1, t2, t3}

	return tokenStream
}

func testStream3() []token {
	t1 := token{grammar.BEGIN, grammar.token_strings[grammar.BEGIN], 0, 0}
	t2 := token{grammar.SKIP, grammar.token_strings[grammar.SKIP], 1, 2}
	t3 := token{grammar.END, grammar.token_strings[grammar.END], 2, 0}
	tokenStream := []token{t1, t2, t3}

	return tokenStream
}
