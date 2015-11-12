package main

//   "github.com/OliWheeler/wacc_19/src/lexer"
import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wacc_19/src/parse"
)

func main() {
	var tokens []parse.Token
	//	b, err := ioutil.ReadFile("wacc_examples/valid/variables/capCharDeclaration.wacc")
	b, err := ioutil.ReadFile("wacc_examples/valid/function/nested_functions/fibonacciFullRec.wacc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	lex := parse.Lex("Something", s)
	for item := range lex.Items {
		tokens = append(tokens, item)
	}

	fmt.Println(tokens)

}
