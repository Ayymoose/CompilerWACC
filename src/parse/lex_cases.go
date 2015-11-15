package parse

import (
	"strconv"
	"unicode"

	"grammar"
)

/*
func lexBegin(l *Lexer) stateFn {
	_ = "breakpoint"
	l.width = len(token_keyword_strings[BEGIN])
	l.pos += l.width
	if unicode.IsLetter(l.next()) {
		l.backup()
		return lexIdentifier
	}
	l.emit(BEGIN)
	return lexInsideProgram
}

func lexEnd(l *Lexer) stateFn {
	l.width = len(token_keyword_strings[END])
	l.pos += l.width
	if unicode.IsLetter(l.next()) {
		l.backup()
		return lexIdentifier
	}
	l.emit(END)
	return lexInsideProgram
}*/

func lexIdentifier(l *Lexer) stateFn {
	if !(unicode.IsLetter(l.peek()) || l.peek() == '_') {
		l.next()
		return l.errorf("bad identifier syntax: %q", l.input[l.start:l.pos])
	}
	for isAlphaNumeric(l.peek()) || l.peek() == '_' {
		l.next()
	}
	//	l.backup()
	l.emit(grammar.IDENTIFIER)
	return lexInsideProgram
}

func lexChar(l *Lexer) stateFn {
Loop:
	for {
		switch l.next() {
		case '"':
			return l.errorf("unescaped char %s", l.input[l.start:l.start+l.width])

		case '\\':
			/*			if r := l.peek(); r != grammar.Eof && r != '\n' {
						fmt.Println(strconv.QuoteRuneToASCII(r))
						fmt.Println("Got to Eof\n\\\\\\\\!!!!!!!!!!!!!!!!")
						l.next()
						break
					}  */
			if _, ok := grammar.EscapeChars[l.peek()]; !ok {
				return l.errorf("Not an escape character %s", strconv.QuoteRuneToASCII(l.next()))
			} else {
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
	/*	if unicode.IsLetter(l.peek()) {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}  */
	l.emit(grammar.NUMBER)
	return lexInsideProgram
}
