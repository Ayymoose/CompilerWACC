package main

type tokenType int

type token struct {
	typ    tokenType
	lexeme string
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
func (p *parser) parseProgram() (bool, string) {
	return false, ""
}

// Maps error messages with the associated terminal
func (p *parser) terminalErrorMessages(token tokenType) (bool, []string) {
	return false, []string{""}
}

/* TERMINALS */

func (p *parser) parseBegin() (bool, string) {
	return false, ""
}

func (p *parser) parseEnd() (bool, string) {
	return false, ""
}

func (p *parser) parseIs() (bool, string) {
	return false, ""
}

func (p *parser) parseSkip() (bool, string) {
	return false, ""
}

func (p *parser) parseRead() (bool, string) {
	return false, ""
}

func (p *parser) parseFree() (bool, string) {
	return false, ""
}

func (p *parser) parseReturn() (bool, string) {
	return false, ""
}

func (p *parser) parseExit() (bool, string) {
	return false, ""
}

func (p *parser) parsePrint() (bool, string) {
	return false, ""
}

func (p *parser) parsePrintln() (bool, string) {
	return false, ""
}

func (p *parser) parseIf() (bool, string) {
	return false, ""
}

func (p *parser) parseThen() (bool, string) {
	return false, ""
}

func (p *parser) parseElse() (bool, string) {
	return false, ""
}

func (p *parser) parseFi() (bool, string) {
	return false, ""
}

func (p *parser) parseWhile() (bool, string) {
	return false, ""
}

func (p *parser) parseDo() (bool, string) {
	return false, ""
}
func (p *parser) parseDone() (bool, string) {
	return false, ""
}

func (p *parser) parseStat() (bool, string) {
	return false, ""
}

func (p *parser) parseNewpair() (bool, string) {
	return false, ""
}

func (p *parser) parseFst() (bool, string) {
	return false, ""
}

func (p *parser) parseSnd() (bool, string) {
	return false, ""
}

func (p *parser) parseInt() (bool, string) {
	return false, ""
}

func (p *parser) parseBool() (bool, string) {
	return false, ""
}

func (p *parser) parseChar() (bool, string) {
	return false, ""
}

func (p *parser) parseString() (bool, string) {
	return false, ""
}

func (p *parser) parsePair() (bool, string) {
	return false, ""
}

func (p *parser) parseTrue() (bool, string) {
	return false, ""
}

func (p *parser) parseFalse() (bool, string) {
	return false, ""
}

func (p *parser) parseNull() (bool, string) {
	return false, ""
}

// '#' is a comment
func (p *parser) parseComment() (bool, string) {
	return false, ""
}

func (p *parser) parseLBracket() (bool, string) {
	return false, ""
}

func (p *parser) parseRBracket() (bool, string) {
	return false, ""
}

func (p *parser) parseComma() (bool, string) {
	return false, ""
}

func (p *parser) parseEquals() (bool, string) {
	return false, ""
}

func (p *parser) parseSemiColon() (bool, string) {
	return false, ""
}

func (p *parser) parseLSquareBracket() (bool, string) {
	return false, ""
}

func (p *parser) parseRSquareBracket() (bool, string) {
	return false, ""
}

func (p *parser) parseExclamation() (bool, string) {
	return false, ""
}

func (p *parser) parseMinus() (bool, string) {
	return false, ""
}

func (p *parser) parseLen() (bool, string) {
	return false, ""
}

func (p *parser) parseOrd() (bool, string) {
	return false, ""
}

func (p *parser) parseChr() (bool, string) {
	return false, ""
}

/* ---------------------------------------------------------------------------*/
