package parse

import (
	"fmt"
	"strconv"
	"unicode"

	"grammar"
)

func lexIdentifier(l *Lexer) stateFn {
	if !(unicode.IsLetter(l.peek()) || l.peek() == '_') {
		l.next()
		return l.errorf("bad identifier syntax: %q", l.input[l.start:l.pos])
	}
	for isAlphaNumeric(l.peek()) || l.peek() == '_' {
		l.next()
	}
	l.emit(grammar.IDENTIFIER)
	return lexInsideProgram
}

func lexChar(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '"':
			fmt.Println(string(l.input[l.pos-2]))
			if l.input[l.pos-2] != '\\' {
				return l.errorf("unescaped char %s", string(l.input[l.pos-1]))
			}
		case '\\':
			if _, ok := grammar.EscapeChars[l.peek()]; !ok {
				return l.errorf("Not an escape character %s", strconv.QuoteRuneToASCII(l.next()))
			} else {
				l.next()
				break
			}
			fallthrough
		case grammar.Eof, '\n':
			return l.errorf("unterminated character constant")
		case '\'':
			break Loop
		}
	}
	l.emit(grammar.CHARLITER)
	return lexInsideProgram
}

func lexString(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != grammar.Eof && r != '\n' {
				break
			}
			fallthrough
		case grammar.Eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			break Loop
		}
	}
	l.emit(grammar.STRINGLITER)
	return lexInsideProgram
}

func lexNumber(l *Lexer) stateFn {
	l.accept("+-")
	digits := "0123456789"
	l.acceptRun(digits)
	l.emit(grammar.NUMBER)
	return lexInsideProgram
}
