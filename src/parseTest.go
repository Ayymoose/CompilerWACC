package main

import "fmt"

func main() {
	t1 := token{BEGIN, token_strings[BEGIN]}
	t2 := token{IDENTIFIER, "nman"}
	tokenStream := []token{t1, t2}
	p := constructParser(tokenStream)
	p.advance()

	fmt.Println("Hello World", p.curr)
	fmt.Println("Token Identifer String: ", token_strings[IDENTIFIER], "**")
}
