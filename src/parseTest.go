package main

import "fmt"

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
	t1 := token{IDENTIFIER, "abc", 0, 0}
	t2 := token{IDENTIFIER, "nman", 1, 5}
	tokenStream := []token{t1, t2}

	return tokenStream
}

func testStream2() []token {
	t1 := token{BEGIN, token_strings[BEGIN], 0, 0}
	t2 := token{IDENTIFIER, "nman", 1, 2}
	t3 := token{END, token_strings[END], 2, 0}
	tokenStream := []token{t1, t2, t3}

	return tokenStream
}

func testStream3() []token {
	t1 := token{BEGIN, token_strings[BEGIN], 0, 0}
	t2 := token{SKIP, token_strings[SKIP], 1, 2}
	t3 := token{END, token_strings[END], 2, 0}
	tokenStream := []token{t1, t2, t3}

	return tokenStream
}
