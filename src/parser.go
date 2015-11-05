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

	p.saveToken()

	// Check for "begin"
	if !p.expectToken(BEGIN, &errorMsgs, "All programs must start with 'begin'") {
		return pass, errorMsgs
	}

	p.saveToken()

	// Check for <Func>*
	p.parseZeroOrMore(p.parseFunc, &errorMsgs)

	p.saveToken()

	// Checks for <Stat>
	if !p.parseOne(p.parseStat, &errorMsgs) {
		return pass, errorMsgs
	}

	p.saveToken()

	// Checks for "end"
	if !p.expectToken(END, &errorMsgs, "All programs must terminate with 'end'") {
		return pass, errorMsgs
	}

	// If all pattern checks were succesful then return true with no error messages
	return true, []string{}

}

/* TERMINALS */
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

	errorMsgs = append(errorMsgs, p.makeErrorMsg("parseStat is not implemented"))

	if !pass {
		p.backTrack()
	}

	return pass, errorMsgs
}

/* PARSE HELPERS */

func (p *parser) parseOne(parseCheck func() (bool, []string), errorMsgs *[]string) bool {
	var errorMsgTemp []string
	var match = false

	match, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)

	return match
}

// Returns true iff currTok matches the expectedType. If the check fails then
// An optional error message is appened to errorMsgs.
// If the check is not strict then leave errorMsg blank a.k.a ""
func (p *parser) expectToken(expectedType tokenType, errorMsgs *[]string, errorMsg string) bool {
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
func (p *parser) parseOptional(parseCheck func() (bool, []string), errorMsgs *[]string) {
	var errorMsgTemp []string

	_, errorMsgTemp = parseCheck()

	*errorMsgs = append(*errorMsgs, errorMsgTemp...)

}

// Attempts to parse zero or patterns based on parseCheck
// I.e. <Check>*
// This function has no obligation to parse anything, but it may still add
// possible error messages
func (p *parser) parseZeroOrMore(parseCheck func() (bool, []string), errorMsgs *[]string) {
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
func (p *parser) parseOneOrMore(parseCheck func() (bool, []string), errorMsgs *[]string) bool {
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
