package lexer

import "strings"

type stateFn func(*lexer) stateFn

type item struct {
	typ itemType
	val string
}

type lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan item
}

const (
	itemError = iota
	itemEOF
	itemBegin
	itemEnd
	itemInt
	itemElse       // else keyword
	itemIdentifier // identifier
	itemIf         // if keyword
	itemString     // quoted string (includes quotes)
	itemText       // plain text
)

const (
	eof   = -1
	begin = "begin"
	end   = "end"
)

func (i item) String() string {
	// TODO:
	return ""
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func lexText(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], begin) {
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexBegin
		}
		if l.next() == eof {
			break
		}
	}
	//Correclty reached eof
	if l.pos > l.start {
		l.emit(itemText)
	}
	l.emit(itemEOF)
	return nil
}

func lex(name, input string) (*lexer, chan item) {

	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l, l.items
}
