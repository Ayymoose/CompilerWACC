package lexer

import (
	"fmt"
	"strings"

	grammar "github.com/OliWheeler/wacc_19/src/grammar"
)

type stateFn func(*Lexer) stateFn

type Item struct {
	typ grammar.ItemType
	val string
}

type Lexer struct {
	name  string
	input string
	state stateFn
	start int
	pos   int
	width int
	Items chan Item
}

func (i Item) String() string {
	switch i.typ {
	case grammar.EOF:
		return "EOF"
	case grammar.ERROR:
		return i.val
	}
	//	return fmt.Sprintf("grammar.Item{%v, %q},", grammar.DebugTokens[i.typ], i.val)
	return fmt.Sprintf("%v : %q", grammar.DebugTokens[i.typ], i.val)
}

func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.Items)
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextItem() Item {
	item := <-l.Items
	//l.lastPos = item.pos
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
		Items: make(chan Item),
	}
	go l.run()
	return l
}
