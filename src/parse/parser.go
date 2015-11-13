package parse

//TODO: Change the if statement in the parse functions' return type so that there is only one return
//I.e return pass, errorMsgs
//TODO: FIX ARRAY TYPE PARSE

import (
	"fmt"
	"strconv"

	grammar "github.com/nanaasiedu/wacc_19/src/grammar"
)

func (i Token) String() string {
	switch i.Typ {
	case grammar.EOF:
		return "EOF"
	case grammar.ERROR:
		return i.Lexeme
	}
	//	return fmt.Sprintf("grammar.Item{%v, %q},", grammar.DebugTokens[i.typ], i.val)
	return fmt.Sprintf("%v : %q", grammar.DebugTokens[i.Typ], i.Lexeme)
}

// Used as the currTok onced the oken stream is finished
var TERMINATE_TOKEN grammar.ItemType = -1

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
	lexer  *Lexer
	tokens []Token // Stream of tokens from the lexer
	curr   int     // Index of current token
	save   []int   // Array of indexs to save points in the token stream
	//                Used for back tracking
	currTok Token // the current token

}

// Basic parser constructer that sets the current token to the first token in the
// tokenStream
func (l *Lexer) ConstructParser(tokenStream []Token) *parser {
	return &parser{l, tokenStream, 0, []int{}, tokenStream[0]}
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
		p.currTok = Token{Typ: TERMINATE_TOKEN,
			Pos: p.currTok.Pos,
		}
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
	return "At " + strconv.Itoa(p.lexer.lineNumber(p.currTok)) //p.currTok.LineNum) //+ ":" + strconv.Itoa(p.currTok.RowNum)
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

	// <program> := 'begin' <func>* <stat> 'end'

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
	var pass = false       // True iff the tokens match a <func> def
	var errorMsgs []string // An array of error messages

	// <func> := <type> <ident> '(' <param-list>? ')' 'is' <stat> 'end'

	expected := []grammar.ItemType{grammar.OPEN_ROUND, grammar.CLOSE_ROUND, grammar.IS, grammar.END}
	parseTypes := []parseType{p.parseType, p.parseIdent, p.parseParamList, p.parseStat}
	patternTypes := []patternType{ONCE, ONCE, EXPECT, OPTIONAL, EXPECT, EXPECT, ONCE, EXPECT}
	segmentErrors := []string{"", "", "", "",
		"Expected ')' to close function parameter list",
		"Expected 'is' after function parameter list",
		"Expected statement body for function",
		"All functions must end with 'end'"}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseParamList() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages

	// <param-list> := <param> ( ',' <param> )*

	expected := []grammar.ItemType{}
	parseTypes := []parseType{p.parseParam, p.parseExtraParam}
	patternTypes := []patternType{ONCE, ZEROMORE}
	segmentErrors := []string{"", "", "", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseExtraParam() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-liter) def
	var errorMsgs []string // An array of error messages

	// Helper parse function for parseParamList
	// extram param := ',' <param>

	expected := []grammar.ItemType{grammar.COMMA}
	parseTypes := []parseType{p.parseParam}
	patternTypes := []patternType{EXPECT, ONCE}
	segmentErrors := []string{"", "Expected a parameter after comma"}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseParam() (bool, []string) {
	var pass = false       // True iff the tokens match a (param) def
	var errorMsgs []string // An array of error messages

	// <param> := <type> <ident>
	expected := []grammar.ItemType{}
	parseTypes := []parseType{p.parseType, p.parseIdent}
	patternTypes := []patternType{ONCE, ONCE}
	segmentErrors := []string{"", "An parameter identifier must follow its type"}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseStat() (bool, []string) {
	var pass = false      // True iff the tokens match a <program> def
	var multiStat = false // True iff the the statement contains multiple
	//                       statements I.e. <stat> ';' <stat>
	var errorMsgs []string     // An array of error messages
	var errorMsgsTemp []string // Error messages place holder

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

	// <type> <ident> '=' <assign-rhs> option
	expected = []grammar.ItemType{grammar.EQ}
	parseTypes = []parseType{p.parseType, p.parseIdent, p.parseAssignRHS}
	patternTypes = []patternType{ONCE, ONCE, EXPECT, ONCE}
	segmentErrors = []string{"", "Identifier must follow its type",
		"Expected '=' assignment after identifier",
		"Expected value on rhs of assignment"}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <assign-lhs> '=' <assign-rhs> option
	expected = []grammar.ItemType{grammar.EQ}
	parseTypes = []parseType{p.parseAssignLHS, p.parseAssignRHS}
	patternTypes = []patternType{ONCE, EXPECT, ONCE}
	segmentErrors = []string{"", "Expected '=' assignment after variable", ""}

	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'read' <assign-lhs> option
	expected = []grammar.ItemType{grammar.READ}
	parseTypes = []parseType{p.parseAssignLHS}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "Expected variable after 'read'"}

	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'free' <expr> option
	expected = []grammar.ItemType{grammar.FREE}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "A variable must follow 'free'"}

	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'return' <Expr> option
	expected = []grammar.ItemType{grammar.RETURN}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "Expected expression after 'return'"}

	op6 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'exit' <Expr> option
	expected = []grammar.ItemType{grammar.EXIT}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "Expected expression after 'exit'"}

	op7 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'print' <Expr> option
	expected = []grammar.ItemType{grammar.PRINT}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "Expected expression after 'print'"}

	op8 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'println' <Expr> option
	expected = []grammar.ItemType{grammar.PRINTLN}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}
	segmentErrors = []string{"", "Expected expression after 'print'"}

	op9 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'If' <Expr> 'then' <stat> 'else' <stat> 'fi' option
	expected = []grammar.ItemType{grammar.IF, grammar.THEN, grammar.ELSE, grammar.FI}
	parseTypes = []parseType{p.parseExpr, p.parseStat, p.parseStat}
	patternTypes = []patternType{EXPECT, ONCE, EXPECT, ONCE, EXPECT, ONCE, EXPECT}
	segmentErrors = []string{"", "Expected a boolean after 'if'",
		"Expected 'then' after If-boolean",
		"Expected statement",
		"Expected 'else'",
		"Expected statement",
		"Expected 'fi' to end If-statement"}

	op10 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'while' <Expr> 'do' <stat> 'done' option
	expected = []grammar.ItemType{grammar.WHILE, grammar.DO, grammar.DONE}
	parseTypes = []parseType{p.parseExpr, p.parseStat}
	patternTypes = []patternType{EXPECT, ONCE, EXPECT, ONCE, EXPECT}
	segmentErrors = []string{"", "Expected a boolean after 'while'",
		"Expected 'do' after while-boolean",
		"Expected statement",
		"Expected 'done' to end while-loop"}

	op11 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'begin' <stat> 'end' option
	expected = []grammar.ItemType{grammar.BEGIN, grammar.END}
	parseTypes = []parseType{p.parseStat}
	patternTypes = []patternType{EXPECT, ONCE, EXPECT}
	segmentErrors = []string{"", "Expected statement in scope",
		"Expected 'end' to end scope"}

	op12 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op12, op2, op3, op4, op2, op5, op1, op4, op6, op7, op8,
		op9, op10, op11)

	// <stat> ; <stat> option
	if pass { // If a <stat> has been read so far

		// Check for a ';'
		multiStat, _ = p.parsePattern([]grammar.ItemType{grammar.SEMICOLON},
			[]parseType{},
			[]patternType{EXPECT},
			[]string{""})

		if multiStat {
			pass, errorMsgsTemp = p.parseStat()
		}

	}

	errorMsgs = append(errorMsgs, errorMsgsTemp...)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseAssignLHS() (bool, []string) {
	var pass = false       // True iff the tokens match a (assign-LHS) def
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// All options are parsed onced
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	// <ident> option
	parseTypes = []parseType{p.parseIdent}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <array-elem> option
	parseTypes = []parseType{p.parseArrayElem}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <pair-elem> option
	parseTypes = []parseType{p.parsePairElem}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseAssignRHS() (bool, []string) {
	var pass = false       // True iff the tokens match a (assign-RHS) def
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// <expr> option
	expected = []grammar.ItemType{}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <array-liter> option
	expected = []grammar.ItemType{}
	parseTypes = []parseType{p.parseArrayLiteral}
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'newpair' '(' <expr> ',' <expr> ')' option
	expected = []grammar.ItemType{grammar.NEWPAIR, grammar.OPEN_ROUND, grammar.COMMA, grammar.CLOSE_ROUND}
	parseTypes = []parseType{p.parseExpr, p.parseExpr}
	patternTypes = []patternType{EXPECT, EXPECT, ONCE, EXPECT, ONCE, EXPECT}
	segmentErrors = []string{"", "Expected '(' after newpair",
		"", "", "", ""}

	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <pair-elem>
	expected = []grammar.ItemType{}
	parseTypes = []parseType{p.parsePairElem}
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'call' <ident> '(' <arg-list>? ')' option
	expected = []grammar.ItemType{grammar.CALL, grammar.OPEN_ROUND, grammar.CLOSE_ROUND}
	parseTypes = []parseType{p.parseIdent, p.parseArgList}
	patternTypes = []patternType{EXPECT, ONCE, EXPECT, OPTIONAL, EXPECT}
	segmentErrors = []string{"", "Expected function identifier",
		"Expected '(' after identifer",
		"", "Expected ')' to end argument list"}

	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4, op5)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArgList() (bool, []string) {
	var pass = false       // True iff the tokens match a (arg-list) def
	var errorMsgs []string // An array of error messages

	// <arg-list> := <expr> ( ',' <expr> )*

	expected := []grammar.ItemType{}
	parseTypes := []parseType{p.parseExpr, p.parseExtraArg}
	patternTypes := []patternType{ONCE, ZEROMORE}
	segmentErrors := []string{"", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseExtraArg() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages

	// Helper parse function for parseArgList
	// extram param := ',' <expr>

	expected := []grammar.ItemType{grammar.COMMA}
	parseTypes := []parseType{p.parseExpr}
	patternTypes := []patternType{EXPECT, ONCE}
	segmentErrors := []string{"", "Expected an argument after comma"}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parsePairElem() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-elem) def
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// All options expect a token then an Expr
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE}

	// 'fst' <expr>
	expected = []grammar.ItemType{grammar.FST}
	segmentErrors = []string{"", "Expected an expression after 'fst'"}

	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'snd' <expr>
	expected = []grammar.ItemType{grammar.SND}
	segmentErrors = []string{"", "Expected an expression after 'snd'"}

	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseType() (bool, []string) {
	var pass = false       // True iff the tokens match a <type> def
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// All options are parsed once
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	// <base-type>
	parseTypes = []parseType{p.parseBaseType}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <array-type>
	parseTypes = []parseType{p.parseArrayType}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <pair-type>
	parseTypes = []parseType{p.parsePairType}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
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
	segmentErrors = []string{""}

	// 'int'
	expected = []grammar.ItemType{grammar.INT}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'bool'
	expected = []grammar.ItemType{grammar.BOOL}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'char'
	expected = []grammar.ItemType{grammar.CHAR}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'string'
	expected = []grammar.ItemType{grammar.STRING}
	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayType() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// <array-type> := <type> '[' ']'

	// All options are parsed once
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	// <base-type>
	parseTypes = []parseType{p.parseBaseType}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <pair-type>
	parseTypes = []parseType{p.parsePairType}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2)

	// if <type> != <base-type> or <pair-type> then check if <type> = <array-type>
	if !pass {
		// NOT IMPLEMENTED
	}

	// Check for '[' ']' after <type> is parsed successfully
	expected = []grammar.ItemType{grammar.OPEN_SQUARE, grammar.CLOSE_SQUARE}
	parseTypes = []parseType{}
	patternTypes = []patternType{EXPECT, EXPECT}
	segmentErrors = []string{"", "Expected '[' after type", "Expected ']'"}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parsePairType() (bool, []string) {
	var pass = false       // True iff the tokens match a (pair-type) def
	var errorMsgs []string // An array of error messages

	// <pair-type> := 'pair' '(' <pair-elem-type> ',' <pair-elem-type> ')'

	expected := []grammar.ItemType{grammar.PAIR, grammar.OPEN_ROUND, grammar.COMMA, grammar.CLOSE_ROUND}
	parseTypes := []parseType{p.parsePairElemType, p.parsePairElemType}
	patternTypes := []patternType{EXPECT, EXPECT, ONCE, EXPECT, ONCE, EXPECT}
	segmentErrors := []string{"", "", "", "", "", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parsePairElemType() (bool, []string) {
	var pass = false       // True iff the tokens match a (elem-type) def
	var errorMsgs []string // An array of error messages

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// First two options parse once
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	// <base-type>
	parseTypes = []parseType{p.parseBaseType}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <array-type>
	parseTypes = []parseType{p.parseArrayType}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'pair'
	expected = []grammar.ItemType{grammar.PAIR}
	parseTypes = []parseType{}
	patternTypes = []patternType{EXPECT}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseExpr() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages
	var errorMsgsTemp []string

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// First seven options parse once
	patternTypes = []patternType{ONCE}
	segmentErrors = []string{""}

	// <int-liter>
	parseTypes = []parseType{p.parseIntLiteral}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <bool-liter>
	parseTypes = []parseType{p.parseBoolLiteral}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <char-liter>
	parseTypes = []parseType{p.parseCharLiteral}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <str-liter>
	parseTypes = []parseType{p.parseStrLiteral}
	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <pair-liter>
	parseTypes = []parseType{p.parsePairLiteral}
	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <ident>
	parseTypes = []parseType{p.parseIdent}
	op6 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <array-elem>
	parseTypes = []parseType{p.parseArrayElem}
	op7 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// <unary-oper> <expr>
	parseTypes = []parseType{p.parseUnaryOp, p.parseExpr}
	patternTypes = []patternType{ONCE, ONCE}
	segmentErrors = []string{"", ""}
	op8 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// '(' <expr> ')'
	expected = []grammar.ItemType{grammar.OPEN_ROUND, grammar.CLOSE_ROUND}
	parseTypes = []parseType{p.parseExpr}
	patternTypes = []patternType{EXPECT, ONCE, EXPECT}
	segmentErrors = []string{"", "", ""}
	op9 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op6, op1, op2, op3, op4, op5, op7, op8, op9)

	// <binary-oper> <expr>
	expected = []grammar.ItemType{}
	parseTypes = []parseType{p.parseBinaryOp, p.parseExpr}
	patternTypes = []patternType{ONCE, ONCE}
	segmentErrors = []string{"", ""}

	_, errorMsgsTemp = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)
	errorMsgs = append(errorMsgs, errorMsgsTemp...)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseUnaryOp() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a ⟨unary-oper⟩ def

	// <unary-oper> ::= ‘!’ | ‘-’ | ‘len’ | ‘ord’ | ‘chr’
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// All options EXPECT one token
	patternTypes = []patternType{EXPECT}
	segmentErrors = []string{""}

	// ‘!’
	expected = []grammar.ItemType{grammar.NOT}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘-’
	expected = []grammar.ItemType{grammar.NEG}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘len’
	expected = []grammar.ItemType{grammar.LEN}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘ord’
	expected = []grammar.ItemType{grammar.ORD}
	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘chr’
	expected = []grammar.ItemType{grammar.CHR}
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

	// <binary-oper> ::= ‘*’ | ‘/’ | ‘%’ | ‘+’ | ‘-’ | ‘>’ | ‘>=’ | ‘<’ | ‘<=’ | ‘==’ | ‘!=’ | ‘&&’ | ‘||’
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	//All options EXPECT one token
	patternTypes = []patternType{EXPECT}
	segmentErrors = []string{""}

	// ‘*’
	expected = []grammar.ItemType{grammar.MULT}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘/’
	expected = []grammar.ItemType{grammar.DIV}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘%’
	expected = []grammar.ItemType{grammar.MOD}
	op3 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘+’
	expected = []grammar.ItemType{grammar.ADD}
	op4 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘-’
	expected = []grammar.ItemType{grammar.SUB}
	op5 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘>’
	expected = []grammar.ItemType{grammar.GT}
	op6 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘>=’
	expected = []grammar.ItemType{grammar.GTE}
	op7 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘<’
	expected = []grammar.ItemType{grammar.ST}
	op8 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘<=’
	expected = []grammar.ItemType{grammar.STE}
	op9 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘==’
	expected = []grammar.ItemType{grammar.EQ}
	op10 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘!=’
	expected = []grammar.ItemType{grammar.NEQ}
	op11 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘&&’
	expected = []grammar.ItemType{grammar.AND}
	op12 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// ‘||’
	expected = []grammar.ItemType{grammar.OR}
	op13 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op1, op2, op3, op4, op5, op6, op7, op8, op9, op10, op11, op12, op13)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseIdent() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages

	// <ident> := valid identifer

	expected := []grammar.ItemType{grammar.IDENTIFIER}
	parseTypes := []parseType{}
	patternTypes := []patternType{EXPECT}
	segmentErrors := []string{""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayElem() (bool, []string) {
	var pass = false       // True iff the tokens match a (array-elem) def
	var errorMsgs []string // An array of error messages

	// <array-elem> ::= <ident> (‘[’ <expr> ‘]’)+
	expected := []grammar.ItemType{}
	parseTypes := []parseType{p.parseIdent, p.parseArrayElemHelper}
	patternTypes := []patternType{ONCE, ONEMORE}
	segmentErrors := []string{"", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayElemHelper() (bool, []string) {
	var pass = false       // True iff the tokens match a (array-elem) def
	var errorMsgs []string // An array of error messages

	// Partial implementation of array-elem
	// <array-elem-dimension> ::= ‘[’ <expr> ‘]’
	expected := []grammar.ItemType{grammar.OPEN_SQUARE, grammar.CLOSE_SQUARE}
	parseTypes := []parseType{p.parseExpr}
	patternTypes := []patternType{EXPECT, OPTIONAL, EXPECT}
	segmentErrors := []string{"", "", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseIntLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (int-liter) def
	var errorMsgs []string // An array of error messages

	// <int-liter> ::= <int-sign>? <digit>+
	expected := []grammar.ItemType{grammar.NUMBER}
	parseTypes := []parseType{p.parseIntSign}
	patternTypes := []patternType{OPTIONAL, EXPECT}
	segmentErrors := []string{"", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
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
	segmentErrors = []string{""}

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

	// <bool-liter> := 'true' | 'false'

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}
	segmentErrors = []string{""}

	// 'true' option
	expected = []grammar.ItemType{grammar.TRUE}
	op1 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	// 'false' option
	expected = []grammar.ItemType{grammar.FALSE}
	op2 := patternArgs{expected, parseTypes, patternTypes, segmentErrors}

	pass, errorMsgs = p.parseOptions(op2, op1)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseCharLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (char-liter) def
	var errorMsgs []string // An array of error messages

	// <char-liter> ::= CHARLITER
	expected := []grammar.ItemType{grammar.CHARLITER}
	parseTypes := []parseType{}
	patternTypes := []patternType{EXPECT}
	segmentErrors := []string{""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseStrLiteral() (bool, []string) {
	var pass = false       // True iff the tokens match a (str-liter) def
	var errorMsgs []string // An array of error messages

	// (str-liter) ::= STRINGLITER
	expected := []grammar.ItemType{grammar.STRINGLITER}
	parseTypes := []parseType{}

	patternTypes := []patternType{EXPECT}
	segmentErrors := []string{""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseEscapedCharacter() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a <escaped-char> def

	// <escaped-char> ::= ‘0’ | ‘b’ | ‘t’ | ‘n’ | ‘f’ | ‘r’ | ‘"’ | ‘’’ | ‘\’

	//Place holders
	expected := []grammar.ItemType{}
	parseTypes := []parseType{}
	patternTypes := []patternType{}
	segmentErrors := []string{}

	// Each options EXPECTS one token
	patternTypes = []patternType{EXPECT}
	segmentErrors = []string{""}

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
	var pass = false       // True iff the tokens match a (array-liter) def
	var errorMsgs []string // An array of error messages

	// <array-liter> ::= ‘[’ (expr (‘,’ <expr>)*)? ‘]’
	expected := []grammar.ItemType{grammar.OPEN_SQUARE, grammar.CLOSE_SQUARE}
	parseTypes := []parseType{p.parseArrayExtra}
	patternTypes := []patternType{EXPECT, OPTIONAL, EXPECT}
	segmentErrors := []string{"", "", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseArrayExtra() (bool, []string) {
	var pass = false
	var errorMsgs []string // An array of error messages

	// Partial grammar for array-liter
	// <array-liter> ::= <expr> (‘,’ <expr>)*
	expected := []grammar.ItemType{}
	parseTypes := []parseType{p.parseExpr, p.parseExtraArg}
	patternTypes := []patternType{ONCE, ZEROMORE}
	segmentErrors := []string{"", ""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parsePairLiteral() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = false       // True iff the tokens match a ⟨pair-liter⟩ def

	// <pair-liter> ::= ‘null’
	expected := []grammar.ItemType{grammar.NULL}
	parseTypes := []parseType{}
	patternTypes := []patternType{EXPECT}
	segmentErrors := []string{""}

	pass, errorMsgs = p.parsePattern(expected, parseTypes, patternTypes, segmentErrors)

	if !pass {
		return false, errorMsgs
	}

	return true, []string{}
}

func (p *parser) parseComment() (bool, []string) {
	return true, []string{}
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

	// For debugging
	/*fmt.Println("Error IDEN = ", segmentErrors[0])
	fmt.Println("Len seg + exp = ", len(expArgs)+len(segments), " len typs = ", len(typs), " len errors = ", len(segmentErrors))
	fmt.Print("\n")*/

	p.saveToken()

	for i, typ := range typs {
		if i >= len(segmentErrors) {
			i = 0
		}

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
