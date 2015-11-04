package main

type token struct {
	typ    tokenType
	lexeme string
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
	return false, []string{"parseProgram is not implemented"}
}

/* TERMINALS */

/* ---------------------------------------------------------------------------*/
