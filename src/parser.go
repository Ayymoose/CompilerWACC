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
	defer p.advance()
	return p.currTok.typ == typ
}

// Returns a string with the location of the currTok in the input text
func (p *parser) tokenPos() string {
	return "At " + strconv.Itoa(p.currTok.lineNum) + ":" + strconv.Itoa(p.currTok.rowNum)
}

// Returns the string str formmated as an error message for currTok
func (p *parser) makeErrorMsg(str string) string {
	return p.tokenPos() + " :: " + str
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

	var errorMsg []string // Temporary place holder for an error message
	var match = false     // Temporary place holder for a parse function success bool

	p.saveToken()

	// Check for "begin"
	if !p.expect(BEGIN) {
		p.backTrack()
		errorMsgs = append(errorMsgs, p.makeErrorMsg("All programs must start with a 'begin' keyword"))

		return pass, errorMsgs

	}

	//Check for <Func>*
	match, errorMsg = p.parseFunc()

	for {
		if !match {
			errorMsgs = append(errorMsgs, errorMsg...)
			break

		} else {
			match, errorMsg = p.parseFunc()
		}

	}

	return pass, errorMsgs

}

func (p *parser) parseFunc() (bool, []string) {
	var pass = false       // True iff the tokens match a <program> def
	var errorMsgs []string // An array of error messages

	errorMsgs = append(errorMsgs, p.makeErrorMsg("parseFunction is not implemented"))

	return pass, errorMsgs

}

// Maps error messages with the associated terminal
func (p *parser) terminalErrorMessages(token tokenType) (bool, []string) {
	return false, []string{""}
}

/* TERMINALS */

/* ---------------------------------------------------------------------------*/
