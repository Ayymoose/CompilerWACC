package main

type tokenType int

type token struct {
	typ    tokenType
	lexeme []string
}

/* PARSER */

// The paser struct will be used as the parser of the stream of tokens given to
// it.
type parser struct {
	tokens  []token // Stream of tokens from the lexer
	curr    int     // Index of current token
	save    int     // Index of a back-track token
	currTok token   // the current token

}

// Returns true iff the token stream is finished
// I.e. we are at the end of the stream
func (p *parser) isFinished() bool {
	return p.curr > len(p.tokens)
}

// Advances the current token to the next token in the token stream
func (p *parser) advance() {
	p.curr++

	if p.isFinished() {
		return
	}

	p.currTok = p.tokens[p.curr]
}

// Initiates parse Operation
func (p *parser) parse() bool {
	return false
}

// Returns true iff the current token has the tokenType typ
func (p *parser) expect(typ tokenType) bool {
	return p.currTok.typ == typ
}

/* BNF parse functions -------------------------------------------------------*/

/* NON-TERMINALS */
func (p *parser) parseProgram() (bool, []string) {
	return false, []string{""}
}

// Maps error messages with the associated terminal
func terminalErrorMessages(token Token) (string, Token) {
	switch {
	//You'll have to manuall change these error messages @Nana
	case token.isEscapedChar():
		return "Missing", token
	case token.isDeliminator():
		return "Missing", token
	case token.isReservedWord():
		return "Missing", token
	case token.isType():
		return "Missing", token
	case token.isUnaryOp():
		return "Missing", token
	case token.isBracketType():
		return "Missing", token
	case token.isBoolean():
		return "Missing", token
	}
	return "Some error", token
	//Can use formatted output to print the string nicely
	//E.g fmt.Sprint("Parse error: %s token",msg) etc..
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
