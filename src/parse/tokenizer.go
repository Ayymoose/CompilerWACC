package parse

import (
	"strings"

	"grammar"
)

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

type Token struct {
	Typ    grammar.ItemType // The type of this item.
	Lexeme string           // The value of this item.
	Pos    int              // The starting position, in bytes, of this item in the input string.
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

// Returns line number and column number of token in .wacc file
func (l *Lexer) TokenLocation(t Token) (line int, col int) {
	line = 1 + strings.Count(l.input[:t.Pos], "\n")
	col = t.Pos - strings.LastIndex(l.input[:t.Pos], "\n")
	return
}

//Returns line number and column number of current lexing item in .wacc file
func (l *Lexer) currLocation() (line int, col int) {
	line = 1 + strings.Count(l.input[:l.pos], "\n")
	col = l.pos - strings.LastIndex(l.input[:l.pos], "\n")
	return
}

// run runs the state machine for the lexer.
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

// lexText scans until an opening action "BEGIN"
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

// lex creates a new scanner for the input string.
func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		Items: make(chan Token),
	}
	go l.run()
	return l
}
