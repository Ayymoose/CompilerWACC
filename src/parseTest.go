package main

import "fmt"

func main() {
	p := constructParser(testStream1())
	pass, errorMsgs := p.parse()

	if !pass {
		fmt.Println("Syntax error ------ ")

		for _, errorMsg := range errorMsgs {
			fmt.Printf("%s\n", errorMsg)
		}

	}

}

func testStream1() []token {
	t1 := token{BEGIN, token_strings[BEGIN], 0, 0}
	t2 := token{IDENTIFIER, "nman", 0, 5}
	tokenStream := []token{t1, t2}

	return tokenStream
}
