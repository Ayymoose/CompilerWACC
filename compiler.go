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
	b, err := ioutil.ReadFile("wacc_examples/valid/variables/capCharDeclaration.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/valid/function/nested_functions/fibonacciFullRec.wacc")
	//	b, err := ioutil.ReadFile("wacc_examples/invalid/syntaxErr/expressions/missingOperand1.wacc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	lexer := parse.Lex("Something", s)
	for item := range lexer.Items {
		fmt.Println(lexer.TokenLocation(item))
		tokens = append(tokens, item)
	}
	fmt.Println(tokens)
	fmt.Println("\n------ Completed Lexing ------\n")
	parser := lexer.ConstructParser(tokens)
	passed, errs := parser.Parse()
	fmt.Println("Parsing Successful:", passed)
	fmt.Println("Parsing errors:", errs)
	fmt.Println("\n------ Completed Parsing ------\n")

}
