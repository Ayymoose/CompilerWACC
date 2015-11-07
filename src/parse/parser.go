package parse

import (
	"fmt"
	"strconv"

	grammar "github.com/nanaasiedu/wacc_19/src/grammar" // CHANGE TO MASTER
)

type Token struct {
	Typ     grammar.ItemType
	Lexeme  string
	LineNum int
	RowNum  int
}

// A struct that contains arguments to a parser.expectToken function call
type expectArgs struct {
	expectedType grammar.ItemType
	errorMsg     string
	errorMsgs    *[]string
}

type patternType int

const (
	ONCE patternType = iota
	OPTIONAL
	ZEROMORE
	ONEMORE
	EXPECT
)

type parseType func() (bool, []string)

/* PARSER --------------------------------------------------------------------*/

// The paser struct will be used as the parser of the stream of tokens given to
// it.
type parser struct {
	tokens []Token // Stream of tokens from the lexer
	curr   int     // Index of current token
	save   []int   // Array of indexs to save points in the token stream
	// Used for back tracking
	currTok Token // the current token

}

// Basic parser constructer that set the current token to the first token in the
// tokenStream
func ConstructParser(tokenStream []Token) *parser {
	return &parser{tokenStream, 0, []int{}, tokenStream[0]}
}

// Prints the string value of currTok
// Useful for debugging
func (p *parser) printTok() {
	fmt.Println(p.tokenPos() + " " + grammar.Token_strings[p.currTok.Typ])
}

// Returns true iff the token stream is finished
// I.e. we are at the end of the stream
func (p *parser) isFinished() bool {
	return p.curr >= len(p.tokens)-1
}

// Advances the current token to the next token in the token stream
func (p *parser) advance() {
	if p.isFinished() {
		return
	}

	p.curr++

	p.currTok = p.tokens[p.curr]
}

// Back track the current token back to the most recent save point
func (p *parser) backTrack() {
	if len(p.save) <= 0 {
		return
	}

	p.curr = p.save[len(p.save)-1]
	p.currTok = p.tokens[p.curr]
}

// Pushs the position of currTok to the end of "save"
func (p *parser) saveToken() {
	p.save = append(p.save, p.curr)
}

// Pops the position of the most recent save point
func (p *parser) removeSave() {
	if len(p.save) <= 0 {
		return
	}

	p.save = p.save[:len(p.save)-1]
}

// replaces the most recent save point with the current position
func (p *parser) reSave() {
	p.removeSave()
	p.saveToken()
}

// Returns true iff the current token has the ItemType typ
func (p *parser) expect(typ grammar.ItemType) bool {
	return p.currTok.Typ == typ
}

// Returns a string with the location of the currTok in the input text
func (p *parser) tokenPos() string {
	return "At " + strconv.Itoa(p.currTok.LineNum) + ":" + strconv.Itoa(p.currTok.RowNum)
}

// Returns the string str formmated as an error message for currTok
func (p *parser) makeErrorMsg(str string) string {
	return p.tokenPos() + " :: " + str
}

/* BNF parse functions -------------------------------------------------------*/

// Initiates parse Operation
func (p *parser) Parse() (bool, []string) {
	var pass, errors = p.parseProgram()

	return pass, errors

}

/* NON-TERMINALS */
func (p *parser) parseProgram() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <program> def

	expected := []expectArgs{expectArgs{grammar.BEGIN, "All programs must start with 'begin'", &[]string{}}, expectArgs{grammar.END, "All programs must terminate with 'end'", &[]string{}}}
	parseTypes := []parseType{p.parseFunc, p.parseStat}
	patternTypes := []patternType{EXPECT, ZEROMORE, ONCE, EXPECT}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}

}

func (p *parser) parseFunc() (bool, []string) {
	var pass = false       // True iff the tokens match a <program> def
	var errorMsgs []string // An array of error messages

	errorMsgs = append(errorMsgs, p.makeErrorMsg("parseFunction is not implemented"))

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseStat() (bool, []string) {
	var pass = false       // True iff the tokens match a <program> def
	var errorMsgs []string // An array of error messages

	/* Option: skip */

	// Checks for "skip"
	if p.expectToken(grammar.SKIP, &errorMsgs, "") {
		return true, errorMsgs
	}

	/* Option: 2 */

	errorMsgs = append(errorMsgs, p.makeErrorMsg("parseStat is not implemented"))

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

/* PARSE HELPERS */

// Attempts to parse one pattern based on parseCheck
// I.e. <Check>
// This function is obligated to accept at one parseCheck
// returns true iff the parseCheck was accepted
func (p *parser) parseOne(parseCheck parseType, errorMsgs *[]string) bool {
	var errorMsgTemp []string
	var match = false

	match, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)

	return match
}

// Returns true iff currTok matches the expectedType. If the check fails then
// An optional error message is appened to errorMsgs.
// If the check is not strict then leave errorMsg blank a.k.a ""
func (p *parser) expectToken(expectedType grammar.ItemType, errorMsgs *[]string, errorMsg string) bool {
	if !p.expect(expectedType) {
		if len(errorMsg) > 0 {
			*errorMsgs = append(*errorMsgs, p.makeErrorMsg(errorMsg))
		}

		return false

	}

	p.advance()

	return true
}

// Attempts to parse one pattern based on parseCheck
// I.e. <Check>?
// This function has no obligation to parse anything, but it may still add
// possible error messages
func (p *parser) parseOptional(parseCheck parseType, errorMsgs *[]string) {
	var errorMsgTemp []string

	_, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)
}

// Attempts to parse zero or patterns based on parseCheck
// I.e. <Check>*
// This function has no obligation to parse anything, but it may still add
// possible error messages
func (p *parser) parseZeroOrMore(parseCheck parseType, errorMsgs *[]string) {
	var errorMsgTemp []string
	var match = false

	match, errorMsgTemp = parseCheck()

	for {
		if !match {
			*errorMsgs = append(*errorMsgs, errorMsgTemp...)
			break

		} else {
			match, errorMsgTemp = parseCheck()
		}

	}
}

// Attempts to parse one or patterns based on parseCheck
// I.e. <Check>+
// This function is obligated to accept at least one parseCheck
// returns true iff the first parseCheck was accepted
func (p *parser) parseOneOrMore(parseCheck parseType, errorMsgs *[]string) bool {
	var errorMsgTemp []string
	var match = false

	match, errorMsgTemp = parseCheck()

	// fail parse iff the first Check failed
	if !match {
		*errorMsgs = append(*errorMsgs, errorMsgTemp...)
		return false
	}

	for {
		if !match {
			*errorMsgs = append(*errorMsgs, errorMsgTemp...)
			break

		} else {
			match, errorMsgTemp = parseCheck()
		}

	}

	return true
}

// Attempts to parse a pattern using a series of pattern match request
// PRE: len(exepected) + len(parse) == len(typs)
func (p *parser) parsePattern(expArgs []expectArgs, segments []parseType, typs []patternType) (bool, []string) {
	defer p.removeSave()
	var errorMsgTemp []string

	p.saveToken()

	for _, typ := range typs {
		switch typ {
		case EXPECT:
			if !p.expectToken(expArgs[0].expectedType, &errorMsgTemp, expArgs[0].errorMsg) {
				return false, errorMsgTemp
			}

			if len(expArgs) > 1 {
				expArgs = expArgs[1:] // Pop front of the list
			}

		case ONCE:
			if !p.parseOne(segments[0], &errorMsgTemp) {
				return false, errorMsgTemp
			}

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}

		case OPTIONAL:
			p.parseOptional(segments[0], &errorMsgTemp)

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}

		case ZEROMORE:
			p.parseZeroOrMore(segments[0], &errorMsgTemp)

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}

		case ONEMORE:
			if !p.parseOneOrMore(segments[0], &errorMsgTemp) {
				return false, errorMsgTemp
			}

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}

		}

		p.reSave()
	}

	return true, []string{}
}

/* ---------------------------------------------------------------------------*/

// Maps error messages with the associated terminal
func TerminalErrorMessages(token grammar.ItemType) string {
	switch {
	//You'll have to manuall change these error messages @Nana
	case token.IsEscapedChar():
		return "Missing" + grammar.Token_strings[token]
	case token.IsDeliminator():
		return "Missing" + grammar.Token_strings[token]
	case token.IsReservedWord():
		return "Missing" + grammar.Token_strings[token]
	case token.IsType():
		return "Missing" + grammar.Token_strings[token]
	case token.IsUnaryOp():
		return "Missing" + grammar.Token_strings[token]
	case token.IsBracketType():
		return "Missing" + grammar.Token_strings[token]
	case token.IsBoolean():
		return "Missing" + grammar.Token_strings[token]
	}
	return "Some error"
	//Can use formatted output to print the string nicely
	//E.g fmt.Sprint("Parse error: %s token",msg) etc..
}
