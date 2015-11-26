package parser

import "errors"

// Parse lexes and parses the program returning root of tree
func ParseFile(filename, text string) (*Program, error) {
	l := newLex(filename, text)
	e := parserParse(l)
	if e == 0 && !l.parseError {
		return l.prog, nil
	}
	return nil, errors.New("Compilation halted due to lex and parse errors")
}
