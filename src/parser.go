package main

import "strconv"

type token struct {
	typ     tokenType
	lexeme  string
	lineNum int
	rowNum  int
}

/* PARSER --------------------------------------------------------------------*/

// The paser struct will be used as the parser of the stream of tokens given to
// it.
type parser struct {
	tokens  []token // Stream of tokens from the lexer
	curr    int     // Index of current token
	save    int     // Index of a back-track token
	currTok token   // the current token

}

// Basic parser constructer that set the current token to the first token in the
// tokenStream
func constructParser(tokenStream []token) *parser {
	return &parser{tokenStream, 0, 0, tokenStream[0]}
}

// Returns true iff the token stream is finished
// I.e. we are at the end of the stream
func (p *parser) isFinished() bool {
	return p.curr > len(p.tokens)
}

// Advances the current token to the next token in the token stream
func (p *parser) advance() {
	if p.isFinished() {
		return
	}

	p.curr++

	p.currTok = p.tokens[p.curr]
}

// Back track the current token back to the save token
func (p *parser) backTrack() {
	p.currTok = p.tokens[p.save]
	p.curr = p.save
}

// Saves the position of the current token to the field "save"
func (p *parser) saveToken() {
	p.save = p.curr
}

// Returns true iff the current token has the tokenType typ
func (p *parser) expect(typ tokenType) bool {
	return p.currTok.typ == typ
}

/* BNF parse functions -------------------------------------------------------*/

// Initiates parse Operation
func (p *parser) parse() (bool, []string) {
	var pass, errors = p.parseProgram()

	return pass, errors

}

/* NON-TERMINALS */
func (p *parser) parseProgram() (bool, []string) {

	var pass = false       // True iff the tokens match a <program> def
	var errorMsgs []string // An array of error messages

	p.saveToken()

	if !p.expect(BEGIN) {
		pass = false
		errorMsgs = append(errorMsgs, "At "+strconv.Itoa(p.currTok.lineNum)+":"+strconv.Itoa(p.currTok.rowNum)+" ::All programs must start with a 'begin' keyword")

		return pass, errorMsgs

	}

	p.advance()

	return pass, errorMsgs

}

/* TERMINALS */

func (p *parser) parseBegin() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseEnd() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseIs() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseSkip() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseRead() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseFree() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseReturn() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseExit() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parsePrint() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parsePrintln() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseIf() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseThen() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseElse() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseFi() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseWhile() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseDo() (bool, []string) {
	return false, []string{""}
}
func (p *parser) parseDone() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseStat() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseNewpair() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseFst() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseSnd() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseInt() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseBool() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseChar() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseString() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parsePair() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseTrue() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseFalse() (bool, []string) {
	return false, []string{""}
}

func (p *parser) parseNull() (bool, []string) {
	return false, []string{""}
}

/* ---------------------------------------------------------------------------*/
// Maps error messages with the associated terminal
func terminalErrorMessages(token tokenType) string {
	switch {
	//You'll have to manuall change these error messages @Nana
	case token.isEscapedChar():
		return "Missing" + token_strings[token]
	case token.isDeliminator():
		return "Missing" + token_strings[token]
	case token.isReservedWord():
		return "Missing" + token_strings[token]
	case token.isType():
		return "Missing" + token_strings[token]
	case token.isUnaryOp():
		return "Missing" + token_strings[token]
	case token.isBracketType():
		return "Missing" + token_strings[token]
	case token.isBoolean():
		return "Missing" + token_strings[token]
	}
	return "Some error"
	//Can use formatted output to print the string nicely
	//E.g fmt.Sprint("Parse error: %s token",msg) etc..
}
