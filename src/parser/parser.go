package parser

import (
	"errors"

	"ast"
)

// Parse lexes and parses the program returning root of tree
func ParseFile(filename, text string) (*ast.Program, error) {
	l := newLex(filename, text)
	/*	for item := range l.Items {
		fmt.Println(item)
	}*/
	e := parserParse(l)
	if e != 0 {
		return nil, errors.New("Compilation halted due to lex and parse errors")
	}
	/*errors := l.prog.semanticCheck()
	if errors != nil {
		for _, str := range error {
			fmt.Println(str)
		}
		return nil, errors.New("Semantic error")
	} */
	return l.prog, nil
}
