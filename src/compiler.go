package main

//   "github.com/OliWheeler/wacc_19/src/lexer"
import (
	"fmt"
	"io/ioutil"
	"os"

	lexer "github.com/OliWheeler/wacc_19/src/lexer"
)

func main() {
	var tokens []lexer.Item
	b, err := ioutil.ReadFile("../../../../../wacc_examples/valid/expressions/boolCalc.wacc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	_ = "breakpoint"
	lex := lexer.Lex("Something", s)
	for item := range lex.Items {
		tokens = append(tokens, item)
	}
	fmt.Println(tokens)
}
