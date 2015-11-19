package parse

import (
	"fmt"

	"github.com/wacc_19/src/grammar"
	//	"grammar"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

type Token struct {
	Typ    grammar.ItemType // The type of this item.
	Lexeme string           // The value of this item.
	Pos    int              // The starting position, in bytes, of this item in the input string.
}

type Lexer struct {
	name       string
	input      string
	state      stateFn
	start      int
	pos        int
	width      int
	lastPos    int // position of most recent item returned by nextItem
	Items      chan Token
	prog       *Program // The parsed program
	lastItem   item     // The last item emitted
	parseError bool
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
	l.lastItem = item
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
func Lex(name string, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		Items: make(chan Token),
	}
	go l.run()
	return l
}

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

// lexInsideAction scans the elements inside action delimiters.
func lexInsideProgram(l *Lexer) stateFn {
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
			s := grammar.Token_strings[key]
			l.width = len(s)
			l.pos += l.width
			switch key {
			case grammar.DOUBLE_QUOTE:
				return lexString
			case grammar.SINGLE_QUOTE:
				return lexChar
			case grammar.NULL_TERMINATOR, grammar.BACKSPACE, grammar.TAB, grammar.LINE_FEED, grammar.FORM_FEED, grammar.CARRIAGE_RETURN:
				l.ignore()
				return lexInsideProgram
			case grammar.COMMENT_LINE:
				for l.next() != '\n' {
					l.ignore()
				}
				l.backup()
				return lexInsideProgram
			}
			l.emit(grammar.ItemType(key))
			return lexInsideProgram
		}
	}
	switch r := l.next(); {
	case unicode.IsSpace(r):
		l.ignore()
		return lexInsideProgram
		/*	case r == '\'':
				return lexChar
			case r == '"':
				return lexString  */
	case '0' <= r && r <= '9':
		l.backup()
		return lexNumber
	case isAlphaNumeric(r) || r == '_':
		l.backup()
		return lexIdentifier
	case r == grammar.Eof:
		return nil
	}
	return l.errorf("Item Not in WACC lanuage: %s", l.input[l.start:l.start+l.width])
}

// next returns the next rune in the input.
func (l *Lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return grammar.Eof
	}
	char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return char
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	char := l.next()
	l.backup()
	return char
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(vaild string) bool {
	if strings.IndexRune(vaild, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// emit passes an item back to the client.
func (l *Lexer) emit(t grammar.ItemType) {
	l.Items <- Token{Typ: t, Lexeme: l.input[l.start:l.pos], Pos: l.start}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	line, col := l.currLocation()
	fmt.Printf("At %d:%d :: ", line, col)
	fmt.Printf(format, args)
	l.Items <- Token{Typ: grammar.ERROR, Lexeme: fmt.Sprintf(format, args...), Pos: l.start}
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

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

// lexInsideAction scans the elements inside action delimiters.
func lexInsideProgram(l *Lexer) stateFn {
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
			s := grammar.Token_strings[key]
			l.width = len(s)
			l.pos += l.width
			switch key {
			case grammar.DOUBLE_QUOTE:
				return lexString
			case grammar.SINGLE_QUOTE:
				return lexChar
			case grammar.NULL_TERMINATOR, grammar.BACKSPACE, grammar.TAB, grammar.LINE_FEED, grammar.FORM_FEED, grammar.CARRIAGE_RETURN:
				l.ignore()
				return lexInsideProgram
			case grammar.COMMENT_LINE:
				for l.next() != '\n' {
					l.ignore()
				}
				l.backup()
				return lexInsideProgram
			}
			l.emit(grammar.ItemType(key))
			return lexInsideProgram
		}
	}
	switch r := l.next(); {
	case unicode.IsSpace(r):
		l.ignore()
		return lexInsideProgram
	case '0' <= r && r <= '9':
		l.backup()
		return lexNumber
	case isAlphaNumeric(r) || r == '_':
		l.backup()
		return lexIdentifier
	case r == grammar.Eof:
		return nil
	}
	return l.errorf("Item Not in WACC lanuage: %s", l.input[l.start:l.start+l.width])
}

// next returns the next rune in the input.
func (l *Lexer) next() (char rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return grammar.Eof
	}
	char, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return char
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	char := l.next()
	l.backup()
	return char
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(vaild string) bool {
	if strings.IndexRune(vaild, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// emit passes an item back to the client.
func (l *Lexer) emit(t grammar.ItemType) {
	l.Items <- Token{Typ: t, Lexeme: l.input[l.start:l.pos], Pos: l.start}
	l.start = l.pos
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	line, col := l.currLocation()
	fmt.Printf("At %d:%d :: ", line, col)
	// fmt.Printf(format, args)
	l.Items <- Token{Typ: grammar.ERROR, Lexeme: fmt.Sprintf(format, args...), Pos: l.start}
	fmt.Println(fmt.Sprintf(format, args...))
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// Error is used by the yacc-generated parser to signal errors.
func (l *Lexer) Error(e string) {
	l.parseError = true
	line, col := l.currLocation()
	fmt.Printf("%q, %d:%d %s at or near %s\n", l.name, line, col, e, printItem(l.lastItem, true))
}

// Lex is used by the yacc-generated parser to fetch the next Lexeme.
func (l *Lexer) Lex(lval *yySymType) int {
	token := l.NextItem()
	*lval = yySymType{str: i.Typ, line: l.lineNumber()}
	return i.typ
}
