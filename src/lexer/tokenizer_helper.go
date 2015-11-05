package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func lexInsideProgram(l *Lexer) stateFn {
	/*	if strings.HasPrefix(l.input[l.pos:], token_keyword_strings[END]) {
		return lexEnd
	}  */
	for t, s := range token_keyword_strings {
		if strings.HasPrefix(l.input[l.pos:], s) {
			l.width = len(s)
			l.pos += l.width
			if isAlphaNumeric(l.next()) {
				l.backup()
				return lexIdentifier
			}
			l.emit(t)
			return lexInsideProgram
		}
	}
	for t, s := range token_strings {
		if strings.HasPrefix(l.input[l.pos:], s) {
			l.width = len(s)
			l.pos += l.width
			l.emit(t)
			return lexInsideProgram
		}
	}
	switch r := l.next(); {
	case unicode.IsSpace(r):
		l.ignore()
	case r == '\'':
		return lexChar
	case r == '"':
		return lexString
	case '0' <= r && r <= '9':
		l.backup()
		return lexNumber
	case isAlphaNumeric(r):
		l.backup()
		return lexIdentifier
	case r == eof:
		return nil
	}
	return lexInsideProgram
}

func (l *Lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return char
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) backup() {
	l.pos -= l.width
}

func (l *Lexer) peek() rune {
	char := l.next()
	l.backup()
	return char
}

func (l *Lexer) accept(vaild string) bool {
	if strings.IndexRune(vaild, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *Lexer) emit(t ItemType) {
	l.Items <- Item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.Items <- Item{ERROR,
		fmt.Sprintf(format, args...),
	}
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
