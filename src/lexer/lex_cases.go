package lexer

import (
	"unicode"

	"github.com/OliWheeler/wacc_19/src/grammar"
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
	if !unicode.IsLetter(l.next()) {
		return l.errorf("bad identifier syntax: %q", l.input[l.start:l.pos])
	}
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) {
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
		case '\\':
			if r := l.next(); r != grammar.Eof && r != '\n' {
				break
			}
			fallthrough
		case grammar.Eof, '\n':
			return l.errorf("unterminated character constant")
		case '\'':
			break Loop
		}
	}
	l.emit(grammar.CHARACTER)
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
	l.emit(grammar.STRINGVALUE)
	return lexInsideProgram
}

func lexNumber(l *Lexer) stateFn {
	l.accept("+-")
	digits := "0123456789"
	l.acceptRun(digits)
	if unicode.IsLetter(l.peek()) {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}
	l.emit(grammar.NUMBER)
	return lexInsideProgram
}
