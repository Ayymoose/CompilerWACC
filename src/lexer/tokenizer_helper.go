package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func lexInsideProgram(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], end) {
			return lexEnd
		}
		for t, s := range token_keyword_strings {
			if strings.HasPrefix(l.input[l.pos:], s) {
				l.width = len(s)
				l.pos += l.width
				if unicode.IsLetter(l.next()) {
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
		case r == '+' || r == '-' || '0' <= r && r <= '9':
			l.backup()
			return lexNumber
		}
	}
}

func (l *lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return char
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	char := l.next()
	l.backup()
	return char
}

func (l *lexer) accept(vaild string) bool {
	if strings.IndexRune(vaild, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}
