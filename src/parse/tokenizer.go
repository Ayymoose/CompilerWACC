package parse

import (
	"strings"

	"github.com/nanaasiedu/wacc_19/src/grammar"
)

type stateFn func(*Lexer) stateFn

/*
type Item struct {
	typ grammar.ItemType
	val string
}  */

type Token struct {
	Typ    grammar.ItemType
	Lexeme string
	Pos    int
}

type Lexer struct {
	name    string
	input   string
	state   stateFn
	start   int
	pos     int
	width   int
	lastPos int // position of most recent item returned by nextItem
	Items   chan Token
}

func (l *Lexer) TokenLocation(t Token) (line int, col int) {
	line = 1 + strings.Count(l.input[:t.Pos], "\n")
	col = t.Pos - strings.LastIndex(l.input[:t.Pos], "\n")
	return
}

func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.Items)
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextItem() Token {
	item := <-l.Items
	l.lastPos = item.Pos
	return item
}

func lexText(l *Lexer) stateFn {
	for {
		_ = "breakpoint"
		if strings.HasPrefix(l.input[l.pos:], grammar.Token_keyword_strings[grammar.BEGIN]) {
			/*	if l.pos > l.start {
				l.emit(PLAINTEXT)
			}     */
			l.ignore()
			return lexInsideProgram
		}
		if l.next() == grammar.Eof {
			break
		}
	}
	//Correclty reached eof
	if l.pos > l.start {
		l.emit(grammar.PLAINTEXT)
	}
	l.emit(grammar.EOF)
	return nil
}

func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		Items: make(chan Token),
	}
	go l.run()
	return l
}
