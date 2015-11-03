package main

import "fmt"

func main() {
	p := new(parser)
	fmt.Println("Hello World", p.tokens[0])
}
