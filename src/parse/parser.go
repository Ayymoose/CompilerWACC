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

// A struct that contains arguments to a parser's parsePattern function call
type patternArgs struct {
	expArgs   []grammar.ItemType
	checks    []parseType
	typs      []patternType
	segErrors []string
}

// Used to define regex iteration patterns E.g. a* a+ a a?
type patternType int

const (
	ONCE     patternType = iota // a
	OPTIONAL                    // a?
	ZEROMORE                    // a*
	ONEMORE                     // a+
	EXPECT                      // E.g. expect(BEGIN)
)

// Specification of a parse function that attempts to parse certain Syntax
// E.g. parseFunc for <func> // parseStat for <stat>
// Returns a bool : true iff the parse was succesful;
//         a []string: a list of error messages
type parseType func() (bool, []string)

/* PARSER --------------------------------------------------------------------*/

// The paser struct will be used as the parser of the stream of tokens given to
// it.
type parser struct {
	tokens []Token // Stream of tokens from the lexer
	curr   int     // Index of current token
	save   []int   // Array of indexs to save points in the token stream
	//                Used for back tracking
	currTok Token // the current token

}

// Basic parser constructer that sets the current token to the first token in the
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
// I.e. When we are at the end of the stream
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

// Back tracks the current token back to the most recent save point
func (p *parser) backTrack() {
	if len(p.save) <= 0 {
		return
	}

	p.curr = p.save[len(p.save)-1]
	p.currTok = p.tokens[p.curr]
}

// Pushes the position of currTok to the end of "save"
// Hence creating a new save point that can be back tracked to using backTrack()
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

// Adds a list of errorMsgs to another list of errorMsgs
func (p *parser) addErrors(errors1 *[]string, errors2 []string) {
	for _, cError := range errors2 {
		if cError != "" {
			*errors1 = append(*errors1, p.makeErrorMsg(cError))
		}
	}
}

/* BNF parse functions -------------------------------------------------------*/

// Initiates parse Operation
func (p *parser) Parse() (bool, []string) {
	var pass, errors = p.parseProgram()

	return pass, errors
}

/* NON-TERMINALS */
// WARNING : Do not use parseOptions to recursively call the same functions
// E.g. Do not have an option begins with parseStat inside of parseStat

// All non terminal parse functions have match the parseType (Above)

func (p *parser) parseProgram() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <program> def

	// Program := 'begin' <func>* <stat> 'end'

	//The tokens that are required
	expected := []grammar.ItemType{grammar.BEGIN, grammar.END}
	//The parsing functions to use
	parseTypes := []parseType{p.parseFunc, p.parseStat}
	//Regex of pattern types
	patternTypes := []patternType{EXPECT, ZEROMORE, ONCE, EXPECT}
	//Error messages
	segmentErrors := []string{"All programs must start with 'begin'", "", "",
		"All programs must terminate with 'end'"}

	//Parse the pattern according to the input
	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseFunc() (bool, []string) {
	var pass = false       // True iff the tokens match a <program> def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseFunction is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseParamList() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseParamList is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseParam() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseParam is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseStat() (bool, []string) {
	var pass = false      // True iff the tokens match a <program> def
	var multiStat = false // True iff the the statement contains multiple
	//                       statements I.e. <stat> ';' <stat>
	var errorMsgs []string     // An array of error messages
	var errorMsgsTemp []string // Error messages place holder

	p.addErrors(&errorMsgs, []string{"parseStat is not implemented"})

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// 'skip' option
	expected = []grammar.ItemType{grammar.SKIP}
	parseTypes = []parseType{}
	patternTypes = []patternType{EXPECT}
	segmentErrors = []string{""}

	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'free' <expr> option
	expected = []grammar.ItemType{grammar.FREE}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "A variable must follow 'free'"}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgsTemp = p.parseOptions(op2, op1)

	errorMsgs = append(errorMsgs, errorMsgsTemp...)

	// <stat> ; <stat> option

	// Check for a ';'
	multiStat, _ = p.parsePattern([]grammar.ItemType{grammar.SEMICOLON},
		[]parseType{},
		[]patternType{EXPECT},
		[]string{""})

	if multiStat {
		pass, errorMsgsTemp = p.parseStat()
	}

	errorMsgs = append(errorMsgs, errorMsgsTemp...)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseAssignLHS() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseAssignLHS is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseAssignRHS() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseAssignRHS is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseArgList() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseArgList is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parsePairElem() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parsePairElem is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseType() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseType is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseBaseType() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <base-type> def

	// base-type := 'int' | 'bool' | 'char' | 'string'

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}

	// 'int'
	expected = []grammar.ItemType{grammar.INT}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'bool'
	expected = []grammar.ItemType{grammar.BOOL}
	patternTypes = []patternType{EXPECT}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'char'
	expected = []grammar.ItemType{grammar.CHAR}
	patternTypes = []patternType{EXPECT}

	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'string'
	expected = []grammar.ItemType{grammar.STRING}
	patternTypes = []patternType{EXPECT}

	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayType() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseArrayType is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parsePairType() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parsePairType is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseElemType() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseElemType is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseExpr() (bool, []string) {
	var pass = false       // True iff the tokens match a (expr) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseExpr is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseUnaryOp() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a ⟨unary-oper⟩ def

	// ⟨unary-oper⟩ ::= ‘!’ | ‘-’ | ‘len’ | ‘ord’ | ‘chr’
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// ‘!’
	expected = []grammar.ItemType{grammar.NOT}
	patternTypes = []patternType{EXPECT}

	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘-’
	expected = []grammar.ItemType{grammar.NEG}
	patternTypes = []patternType{EXPECT}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘len’
	expected = []grammar.ItemType{grammar.LEN}
	patternTypes = []patternType{EXPECT}

	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘ord’
	expected = []grammar.ItemType{grammar.ORD}
	patternTypes = []patternType{EXPECT}

	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘chr’
	expected = []grammar.ItemType{grammar.CHR}
	patternTypes = []patternType{EXPECT}

	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4, op5)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseBinaryOp() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a ⟨binary-oper⟩ def

	// ⟨binary-oper⟩ ::= ‘*’ | ‘/’ | ‘%’ | ‘+’ | ‘-’ | ‘>’ | ‘>=’ | ‘<’ | ‘<=’ | ‘==’ | ‘!=’ | ‘&&’ | ‘||’
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// ‘*’
	expected = []grammar.ItemType{grammar.MULT}
	patternTypes = []patternType{EXPECT}

	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘/’
	expected = []grammar.ItemType{grammar.DIV}
	patternTypes = []patternType{EXPECT}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘%’
	expected = []grammar.ItemType{grammar.MOD}
	patternTypes = []patternType{EXPECT}

	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘+’
	expected = []grammar.ItemType{grammar.ADD}
	patternTypes = []patternType{EXPECT}

	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘-’
	expected = []grammar.ItemType{grammar.SUB}
	patternTypes = []patternType{EXPECT}

	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘>’
	expected = []grammar.ItemType{grammar.GT}
	patternTypes = []patternType{EXPECT}

	op6 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘>=’
	expected = []grammar.ItemType{grammar.GTE}
	patternTypes = []patternType{EXPECT}

	op7 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘<’
	expected = []grammar.ItemType{grammar.ST}
	patternTypes = []patternType{EXPECT}

	op8 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘<=’
	expected = []grammar.ItemType{grammar.STE}
	patternTypes = []patternType{EXPECT}

	op9 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘==’
	expected = []grammar.ItemType{grammar.EQ}
	patternTypes = []patternType{EXPECT}

	op10 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘!=’
	expected = []grammar.ItemType{grammar.NEQ}
	patternTypes = []patternType{EXPECT}

	op11 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘&&’
	expected = []grammar.ItemType{grammar.AND}
	patternTypes = []patternType{EXPECT}

	op12 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘||’
	expected = []grammar.ItemType{grammar.OR}
	patternTypes = []patternType{EXPECT}

	op13 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4, op5, op6, op7, op8, op9, op10, op11, op12, op13)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseIdent() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseIdent is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseArrayElem() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseArrayElem is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseIntLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseIntLiteral is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseDigit() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseDigit is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseIntSign() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <int-sign> def

	// int-sign := '+' | '-'
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}

	// '+' option
	expected = []grammar.ItemType{grammar.ADD}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// '-' option
	expected = []grammar.ItemType{grammar.SUB}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseBoolLiteral() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <bool-liter> def

	// bool-liter := 'true' | 'false'

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}

	// 'true'
	expected = []grammar.ItemType{grammar.TRUE}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	//'false'
	expected = []grammar.ItemType{grammar.FALSE}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op2, op1)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseCharLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseCharLiteral is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseStrLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseStrLiteral is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseCharacter() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseCharacter is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parseEscapedCharacter() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <base-type> def

	// <escaped-char> ::= ‘0’ | ‘b’ | ‘t’ | ‘n’ | ‘f’ | ‘r’ | ‘"’ | ‘’’ | ‘\’

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}

	// ‘0’
	expected = []grammar.ItemType{grammar.NULL_TERMINATOR}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘b’
	expected = []grammar.ItemType{grammar.BACKSPACE}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘t’
	expected = []grammar.ItemType{grammar.TAB}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘n’
	expected = []grammar.ItemType{grammar.LINE_FEED}
	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘f’
	expected = []grammar.ItemType{grammar.FORM_FEED}
	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘r’
	expected = []grammar.ItemType{grammar.CARRIAGE_RETURN}
	op6 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘"’
	expected = []grammar.ItemType{grammar.DOUBLE_QUOTE}
	op7 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘’’
	expected = []grammar.ItemType{grammar.SINGLE_QUOTE}
	op8 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘\’
	expected = []grammar.ItemType{grammar.BACKSLASH}
	op9 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4, op5, op6, op7, op8, op9)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseArrayLiteral is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

func (p *parser) parsePairLiteral() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a ⟨pair-liter⟩ def

	// <pair-liter> ::= ‘null’
	expected := []grammar.ItemType{grammar.NULL}
	parseTypes := []parseType{}
	patternTypes := []patternType{EXPECT}
	segmentErrors := []string{}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseComment() (bool, []string) {
	var pass = false       // True iff the tokens match a (comment) def
	var errorMsgs []string // An array of error messages

	p.addErrors(&errorMsgs, []string{"parseComment is not implemented"})

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

/* PARSE HELPERS */

// Attempts to parse one pattern based on parseCheck
// I.e. <Check>
// This function is obligated to accept one parseCheck
// returns true iff the parseCheck was accepted
// Error obtained are appended to errorMsgs
func (p *parser) parseOne(parseCheck parseType, errorMsgs *[]string) bool {
	var errorMsgTemp []string
	var match = false

	match, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)

	return match
}

// Returns true iff currTok matches the expectedType.
func (p *parser) expectToken(expectedType grammar.ItemType) bool {
	if !p.expect(expectedType) {
		return false
	}

	p.advance()

	return true
}

// Attempts to parse one pattern based on parseCheck
// I.e. <Check>?
// This function has no obligation to parse anything, but it may still add
// possible error messages
// Error obtained are appended to errorMsgs
func (p *parser) parseOptional(parseCheck parseType, errorMsgs *[]string) {
	var errorMsgTemp []string

	_, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)
}

// Attempts to parse zero or patterns based on parseCheck
// I.e. <Check>*
// This function has no obligation to parse anything, but it may still add
// possible error messages
// Error obtained are appended to errorMsgs
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
// Error obtained are appended to errorMsgs
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
// expArgs: All expectToken() checks used within the pattern IN ORDER of use
// segments: All checks (except expectToken checks)
//           used with in the patter IN ORDER
// typs: Used to indicate the number of times each check is processed (regex)
//       Must match order withing expArgs and segments
// segmentErrors: An array of error messages added to errorMsgTemp when its
//                respected segment/expectArgs (depends on ORDER) fails the check
// PRE: len(expArgs) + len(segments) == len(typs)
func (p *parser) parsePattern(expArgs []grammar.ItemType, segments []parseType, typs []patternType, segmentErrors []string) (bool, []string) {
	defer p.removeSave()
	var errorMsgTemp []string

	p.saveToken()

	for i, typ := range typs {
		switch typ {
		case EXPECT:
			if !p.expectToken(expArgs[0]) {
				p.addErrors(&errorMsgTemp, []string{segmentErrors[i]})
				return false, errorMsgTemp
			}

			if len(expArgs) > 1 {
				expArgs = expArgs[1:] // Pop front of the list
			}
			break

		case ONCE:
			if !p.parseOne(segments[0], &errorMsgTemp) {
				p.addErrors(&errorMsgTemp, []string{segmentErrors[i]})
				return false, errorMsgTemp
			}

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}
			break

		case OPTIONAL:
			p.parseOptional(segments[0], &errorMsgTemp)

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}
			break

		case ZEROMORE:
			p.parseZeroOrMore(segments[0], &errorMsgTemp)

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}
			break

		case ONEMORE:
			if !p.parseOneOrMore(segments[0], &errorMsgTemp) {
				p.addErrors(&errorMsgTemp, []string{segmentErrors[i]})
				return false, errorMsgTemp
			}

			if len(segments) > 1 {
				segments = segments[1:] // Pop front of the list
			}
			break

		}

		p.reSave()
	}

	return true, []string{}
}

// Attempts to parse at least one pattern (from a list of possible parsePattern
// arguments) from a list of options
// NOTE: the first option to match will be used, So the order of args is very
//       important. Hence make sure to avoid any overlapping options by
//       rearranging options (or by hard coding the options manually)
// args: A list of arguments to the function parsePattern
func (p *parser) parseOptions(args ...patternArgs) (bool, []string) {
	var pass = false           // Pattern check bool place holder
	var errorMsgs []string     // Error messages returned by each pattern check
	var errorMsgsTemp []string // Error message place holder

	defer p.removeSave()

	p.saveToken()

	for _, patternArg := range args {
		pass, errorMsgsTemp = p.parsePattern(patternArg.expArgs, patternArg.checks, patternArg.typs, patternArg.segErrors)

		errorMsgs = append(errorMsgs, errorMsgsTemp...)

		if pass {
			return true, []string{}
		}

		p.backTrack()

	}

	return pass, errorMsgs
}

/* ---------------------------------------------------------------------------*/

// Is this needed anymore?
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
