package parse

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/nanaasiedu/wacc_19/src/grammar"
)

type ItemTypeSlice []grammar.ItemType

func (i ItemTypeSlice) Len() int {
	return len(i)
}

func (g ItemTypeSlice) Less(i, j int) bool {
	if g[i] < g[j] {
		return true
	}
	return false
}

func (g ItemTypeSlice) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func lexInsideProgram(l *Lexer) stateFn {
	/*	if strings.HasPrefix(l.input[l.pos:], token_keyword_strings[END]) {
		return lexEnd
	}  */

	var keysForKeywords ItemTypeSlice
	for k := range grammar.Token_keyword_strings {
		keysForKeywords = append(keysForKeywords, k)
	}
	sort.Sort(keysForKeywords)
	for _, key := range keysForKeywords {
		if strings.HasPrefix(l.input[l.pos:], grammar.Token_keyword_strings[key]) {
			s := grammar.Token_keyword_strings[key]
			l.width = len(s)
			l.pos += l.width
			if isAlphaNumeric(l.peek()) || l.peek() == '_' {
				//		l.backup()
				return lexIdentifier
			}
			l.emit(grammar.ItemType(key))
			return lexInsideProgram
		}
	}
	var keysForTokens ItemTypeSlice
	for k := range grammar.Token_strings {
		keysForTokens = append(keysForTokens, k)
	}
	sort.Sort(keysForTokens)
	for _, key := range keysForTokens {
		if strings.HasPrefix(l.input[l.pos:], grammar.Token_strings[key]) {
			/*	switch key {
				case grammar.DOUBLE_QUOTE:
					return lexString
				case grammar.SINGLE_QUOTE:
					return lexChar
				}  */
			s := grammar.Token_strings[key]
			l.width = len(s)
			l.pos += l.width
			l.emit(grammar.ItemType(key))
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
	case r == grammar.Eof:
		return nil
	}
	return lexInsideProgram
}

func (l *Lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return grammar.Eof
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

func (l *Lexer) emit(t grammar.ItemType) {
	l.Items <- Token{Typ: t, Lexeme: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.Items <- Token{Typ: grammar.ERROR, Lexeme: fmt.Sprintf(format, args...)}
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
