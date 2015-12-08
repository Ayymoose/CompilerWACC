//line oliver.y:1
package parser

import __yyfmt__ "fmt"

//line oliver.y:3
import (
	. "ast"
)

//line oliver.y:10
type parserSymType struct {
	yys         int
	str         string
	stringconst Str
	number      int
	pos         int
	integer     Integer
	ident       Ident
	character   Character
	boolean     Boolean

	functions      []*Function
	function       *Function
	stmt           Statement
	stmts          []Statement
	assignrhs      Evaluation
	assignlhs      Evaluation
	expr           Evaluation
	exprs          []Evaluation
	params         []Param
	param          Param
	bracketed      []Evaluation
	pairliter      Evaluation
	arrayliter     ArrayLiter
	pairelem       PairElem
	arrayelem      ArrayElem
	typedefinition Type
	pairelemtype   Type
}

var parserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BEGIN",
	"END",
	"IS",
	"SKIP",
	"READ",
	"FREE",
	"RETURN",
	"EXIT",
	"PRINT",
	"PRINTLN",
	"IF",
	"THEN",
	"ELSE",
	"FI",
	"WHILE",
	"DO",
	"DONE",
	"NEWPAIR",
	"CALL",
	"FST",
	"SND",
	"INT",
	"BOOL",
	"CHAR",
	"STRING",
	"PAIR",
	"NOT",
	"NEG",
	"LEN",
	"ORD",
	"CHR",
	"MUL",
	"DIV",
	"MOD",
	"PLUS",
	"SUB",
	"AND",
	"OR",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"EQ",
	"NEQ",
	"POSITIVE",
	"NEGATIVE",
	"TRUE",
	"FALSE",
	"NULL",
	"OPENSQUARE",
	"OPENROUND",
	"CLOSESQUARE",
	"CLOSEROUND",
	"ASSIGNMENT",
	"COMMA",
	"SEMICOLON",
	"ERROR",
	"STRINGCONST",
	"IDENTIFIER",
	"INTEGER",
	"CHARACTER",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line oliver.y:223

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 111,
	53, 77,
	-2, 74,
	-1, 112,
	53, 78,
	-2, 75,
}

const parserNprod = 85
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 435

var parserAct = [...]int{

	78, 4, 117, 122, 63, 25, 121, 77, 7, 31,
	32, 21, 39, 56, 57, 58, 59, 60, 61, 110,
	62, 36, 36, 47, 65, 167, 36, 24, 20, 152,
	146, 107, 71, 72, 24, 35, 66, 151, 166, 67,
	150, 142, 24, 80, 168, 144, 163, 145, 23, 97,
	98, 99, 100, 101, 102, 103, 149, 106, 24, 150,
	74, 33, 75, 75, 37, 148, 162, 34, 34, 76,
	109, 70, 120, 36, 156, 114, 108, 69, 64, 118,
	34, 80, 112, 119, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 34, 111,
	34, 81, 82, 31, 32, 6, 138, 139, 143, 140,
	48, 34, 49, 50, 51, 34, 9, 2, 53, 52,
	22, 147, 45, 90, 92, 89, 91, 38, 24, 24,
	41, 42, 55, 83, 54, 26, 27, 28, 29, 30,
	73, 44, 46, 40, 43, 155, 68, 114, 157, 123,
	118, 160, 159, 161, 112, 79, 116, 5, 164, 165,
	3, 1, 154, 0, 0, 0, 115, 24, 0, 0,
	0, 111, 0, 0, 0, 24, 19, 0, 18, 0,
	24, 8, 10, 11, 12, 13, 14, 15, 16, 0,
	0, 0, 17, 0, 0, 0, 0, 31, 32, 26,
	27, 28, 29, 30, 0, 0, 48, 0, 49, 50,
	51, 0, 0, 0, 53, 52, 26, 27, 28, 29,
	30, 26, 27, 28, 29, 113, 41, 42, 55, 0,
	54, 0, 0, 0, 0, 0, 23, 44, 46, 40,
	43, 86, 88, 87, 84, 85, 95, 96, 90, 92,
	89, 91, 93, 94, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 158, 86, 88, 87, 84, 85,
	95, 96, 90, 92, 89, 91, 93, 94, 0, 0,
	0, 0, 0, 0, 0, 0, 169, 86, 88, 87,
	84, 85, 95, 96, 90, 92, 89, 91, 93, 94,
	0, 0, 0, 0, 0, 0, 0, 0, 137, 86,
	88, 87, 84, 85, 95, 96, 90, 92, 89, 91,
	93, 94, 0, 0, 0, 0, 0, 0, 0, 153,
	86, 88, 87, 84, 85, 95, 96, 90, 92, 89,
	91, 93, 94, 105, 0, 0, 86, 88, 87, 0,
	141, 0, 104, 90, 92, 89, 91, 0, 0, 86,
	88, 87, 84, 85, 95, 96, 90, 92, 89, 91,
	93, 94, 86, 88, 87, 84, 85, 95, 96, 90,
	92, 89, 91, 93, 94, 86, 88, 87, 84, 85,
	95, 96, 90, 92, 89, 91, 93, 94, 86, 88,
	87, 84, 85, 95, 0, 90, 92, 89, 91, 93,
	94, 86, 88, 87, 84, 85, 0, 0, 90, 92,
	89, 91, 93, 94, 86, 88, 87, 84, 85, 0,
	0, 90, 92, 89, 91,
}
var parserPact = [...]int{

	113, -1000, -1000, 174, 56, -1000, -1000, -27, -1000, 7,
	-14, 176, 176, 176, 176, 176, 176, 176, 174, 19,
	-1000, -1000, -1000, 24, -1000, -1000, -1000, -1000, -1000, -1000,
	17, 176, 176, -1000, 174, 6, 14, 80, -1000, 350,
	-1000, -1000, -1000, -1000, -1000, -1000, 24, -1000, 176, 176,
	176, 176, 176, 176, 176, -1000, 350, 350, 350, 350,
	337, 324, 52, -31, -1000, -1000, -1000, -1000, 23, 176,
	196, 350, 350, -1000, 110, 80, -1000, -1000, 350, -1000,
	-1000, 18, -56, 176, 176, 176, 176, 176, 176, 176,
	176, 176, 176, 176, 176, 176, 176, -1000, -1000, -1000,
	-1000, 311, 311, 252, 174, 174, -1000, 5, 176, 295,
	-17, -1000, -1000, 17, 20, 102, -11, -1000, -32, -1000,
	176, 11, 1, 350, 311, 311, 81, 81, 81, -1000,
	-1000, -1000, -1000, 389, 389, 376, 363, -1000, 21, 9,
	274, -1000, 196, 174, 68, 191, -1000, 206, 176, -1000,
	176, 174, -1000, -1000, 10, 41, 174, -1000, 176, -18,
	350, 8, -1000, -1000, 39, 230, -1000, -1000, -1000, -1000,
}
var parserPgo = [...]int{

	0, 161, 160, 157, 105, 1, 7, 116, 0, 3,
	156, 2, 5, 155, 23, 146, 122, 28, 4, 11,
	120, 19,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 3, 10, 10, 11, 7,
	7, 7, 6, 6, 6, 6, 6, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 13, 9, 9, 9, 14, 15, 15,
	16, 12, 12, 20, 21, 21, 21, 18, 18, 18,
	17, 17, 17, 17, 19,
}
var parserR2 = [...]int{

	0, 4, 2, 0, 7, 8, 3, 1, 2, 1,
	1, 1, 1, 1, 1, 6, 5, 3, 1, 1,
	4, 3, 2, 2, 2, 2, 2, 2, 7, 5,
	3, 2, 2, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 0, 2, 4, 3,
	1, 2, 2, 6, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3,
}
var parserChk = [...]int{

	-1000, -1, 4, -2, -5, -3, -4, -18, 7, -7,
	8, 9, 10, 11, 12, 13, 14, 18, 4, 2,
	-17, -19, -20, 62, -14, -12, 25, 26, 27, 28,
	29, 23, 24, 5, 59, 62, 53, 57, -7, -8,
	63, 50, 51, 64, 61, -16, 62, -14, 30, 32,
	33, 34, 39, 38, 54, 52, -8, -8, -8, -8,
	-8, -8, -5, -18, 59, 5, 17, 20, -15, 53,
	54, -8, -8, -4, 54, 57, 55, -6, -8, -13,
	-12, 21, 22, 53, 38, 39, 35, 37, 36, 44,
	42, 45, 43, 46, 47, 40, 41, -8, -8, -8,
	-8, -8, -8, -8, 15, 19, 5, 62, 53, -8,
	-21, -17, -19, 29, -18, 56, -10, -11, -18, -6,
	54, 62, -9, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, 56, -5, -5,
	-8, 55, 58, 6, 56, 58, 62, -8, 54, 55,
	58, 16, 20, 55, -21, -5, 6, -11, 58, -9,
	-8, -5, 56, 5, -5, -8, 56, 17, 5, 56,
}
var parserDef = [...]int{

	0, -2, 3, 0, 0, 2, 18, 0, 19, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	77, 78, 79, 9, 10, 11, 80, 81, 82, 83,
	0, 0, 0, 1, 0, 0, 0, 0, 22, 23,
	35, 36, 37, 38, 39, 40, 41, 42, 0, 0,
	0, 0, 0, 0, 0, 70, 24, 25, 26, 27,
	0, 0, 0, 0, 31, 32, 33, 34, 67, 0,
	0, 71, 72, 17, 0, 0, 84, 21, 12, 13,
	14, 0, 0, 66, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 43, 44, 45,
	46, 47, 48, 0, 0, 0, 30, 0, 0, 0,
	0, -2, -2, 76, 0, 0, 0, 7, 0, 20,
	0, 0, 0, 65, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 0, 0,
	0, 69, 0, 0, 0, 0, 8, 0, 66, 63,
	0, 0, 29, 68, 0, 0, 0, 6, 0, 0,
	64, 0, 73, 4, 0, 0, 16, 28, 5, 15,
}
var parserTok1 = [...]int{

	1,
}
var parserTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64,
}
var parserTok3 = [...]int{
	0,
}

var parserErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	parserDebug        = 0
	parserErrorVerbose = false
)

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

type parserParser interface {
	Parse(parserLexer) int
	Lookahead() int
}

type parserParserImpl struct {
	lookahead func() int
}

func (p *parserParserImpl) Lookahead() int {
	return p.lookahead()
}

func parserNewParser() parserParser {
	p := &parserParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const parserFlag = -1000

func parserTokname(c int) string {
	if c >= 1 && c-1 < len(parserToknames) {
		if parserToknames[c-1] != "" {
			return parserToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !parserErrorVerbose {
		return "syntax error"
	}

	for _, e := range parserErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + parserTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := parserPact[state]
	for tok := TOKSTART; tok-1 < len(parserToknames); tok++ {
		if n := base + tok; n >= 0 && n < parserLast && parserChk[parserAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if parserDef[state] == -2 {
		i := 0
		for parserExca[i] != -1 || parserExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; parserExca[i] >= 0; i += 2 {
			tok := parserExca[i]
			if tok < TOKSTART || parserExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if parserExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += parserTokname(tok)
	}
	return res
}

func parserlex1(lex parserLexer, lval *parserSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		token = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			token = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		token = parserTok3[i+0]
		if token == char {
			token = parserTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(token), uint(char))
	}
	return char, token
}

func parserParse(parserlex parserLexer) int {
	return parserNewParser().Parse(parserlex)
}

func (parserrcvr *parserParserImpl) Parse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	var parserDollar []parserSymType
	_ = parserDollar // silence set and not used
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parsertoken := -1 // parserchar translated into internal numbering
	parserrcvr.lookahead = func() int { return parserchar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		parserstate = -1
		parserchar = -1
		parsertoken = -1
	}()
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parsertoken), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar, parsertoken = parserlex1(parserlex, &parserlval)
	}
	parsern += parsertoken
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parsertoken { /* valid shift */
		parserchar = -1
		parsertoken = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar, parsertoken = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parsertoken {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error(parserErrorMessage(parserstate, parsertoken))
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parsertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parsertoken))
			}
			if parsertoken == parserEofCode {
				goto ret1
			}
			parserchar = -1
			parsertoken = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	// parserp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if parserp+1 >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:97
		{
			parserlex.(*Lexer).prog = &Program{FunctionList: parserDollar[2].functions, StatList: parserDollar[3].stmts, SymbolTable: NewInstance()}
		}
	case 2:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:101
		{
			parserVAL.functions = append(parserDollar[1].functions, parserDollar[2].function)
		}
	case 3:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:102
		{
			parserVAL.functions = []*Function{}
		}
	case 4:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:105
		{
			if !checkStats(parserDollar[6].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserDollar[2].ident, ReturnType: parserDollar[1].typedefinition, StatList: parserDollar[6].stmts, SymbolTable: NewInstance()}
		}
	case 5:
		parserDollar = parserS[parserpt-8 : parserpt+1]
		//line oliver.y:111
		{
			if !checkStats(parserDollar[7].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserDollar[2].ident, ReturnType: parserDollar[1].typedefinition, StatList: parserDollar[7].stmts, ParameterTypes: parserDollar[4].params, SymbolTable: NewInstance()}
		}
	case 6:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:117
		{
			parserVAL.params = append(parserDollar[1].params, parserDollar[3].param)
		}
	case 7:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:118
		{
			parserVAL.params = []Param{parserDollar[1].param}
		}
	case 8:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:120
		{
			parserVAL.param = Param{ParamType: parserDollar[1].typedefinition, Ident: parserDollar[2].ident}
		}
	case 9:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:122
		{
			parserVAL.assignlhs = parserDollar[1].ident
		}
	case 10:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:123
		{
			parserVAL.assignlhs = parserDollar[1].arrayelem
		}
	case 11:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:124
		{
			parserVAL.assignlhs = parserDollar[1].pairelem
		}
	case 12:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:126
		{
			parserVAL.assignrhs = parserDollar[1].expr
		}
	case 13:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:127
		{
			parserVAL.assignrhs = parserDollar[1].arrayliter
		}
	case 14:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:128
		{
			parserVAL.assignrhs = parserDollar[1].pairelem
		}
	case 15:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:129
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserDollar[3].expr, SndExpr: parserDollar[5].expr, Pos: parserDollar[1].pos}
		}
	case 16:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:130
		{
			parserVAL.assignrhs = Call{Ident: parserDollar[2].ident, ParamList: parserDollar[4].exprs, Pos: parserDollar[1].pos}
		}
	case 17:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:132
		{
			parserVAL.stmts = append(parserDollar[1].stmts, parserDollar[3].stmt)
		}
	case 18:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:133
		{
			parserVAL.stmts = []Statement{parserDollar[1].stmt}
		}
	case 19:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:136
		{
			parserVAL.stmt = Skip{Pos: parserDollar[1].pos}
		}
	case 20:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:137
		{
			parserVAL.stmt = Declare{DecType: parserDollar[1].typedefinition, Lhs: parserDollar[2].ident, Rhs: parserDollar[4].assignrhs, Pos: parserDollar[1].pos}
		}
	case 21:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:138
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].assignlhs, Rhs: parserDollar[3].assignrhs, Pos: parserDollar[1].pos}
		}
	case 22:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:139
		{
			parserVAL.stmt = Read{parserDollar[1].pos, parserDollar[2].assignlhs}
		}
	case 23:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:140
		{
			parserVAL.stmt = Free{parserDollar[1].pos, parserDollar[2].expr}
		}
	case 24:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:141
		{
			parserVAL.stmt = Return{parserDollar[1].pos, parserDollar[2].expr}
		}
	case 25:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:142
		{
			parserVAL.stmt = Exit{parserDollar[1].pos, parserDollar[2].expr}
		}
	case 26:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:143
		{
			parserVAL.stmt = Print{parserDollar[1].pos, parserDollar[2].expr}
		}
	case 27:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:144
		{
			parserVAL.stmt = Println{parserDollar[1].pos, parserDollar[2].expr}
		}
	case 28:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:145
		{
			parserVAL.stmt = If{Conditional: parserDollar[2].expr, ThenStat: parserDollar[4].stmts, ElseStat: parserDollar[6].stmts, Pos: parserDollar[1].pos}
		}
	case 29:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:146
		{
			parserVAL.stmt = While{Conditional: parserDollar[2].expr, DoStat: parserDollar[4].stmts, Pos: parserDollar[1].pos}
		}
	case 30:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:147
		{
			parserVAL.stmt = Scope{StatList: parserDollar[2].stmts, Pos: parserDollar[1].pos}
		}
	case 31:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:148
		{
			parserVAL.stmt = nil
		}
	case 32:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:149
		{
			parserVAL.stmt = nil
		}
	case 33:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:150
		{
			parserVAL.stmt = nil
		}
	case 34:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:151
		{
			parserVAL.stmt = nil
		}
	case 35:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:153
		{
			parserVAL.expr = parserDollar[1].integer
		}
	case 36:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:155
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 37:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:156
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 38:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:158
		{
			parserVAL.expr = parserDollar[1].character
		}
	case 39:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:159
		{
			parserVAL.expr = parserDollar[1].stringconst
		}
	case 40:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:160
		{
			parserVAL.expr = parserDollar[1].pairliter
		}
	case 41:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:161
		{
			parserVAL.expr = parserDollar[1].ident
		}
	case 42:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:162
		{
			parserVAL.expr = parserDollar[1].arrayelem
		}
	case 43:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:164
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 44:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:165
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 45:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:166
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 46:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:167
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 47:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:168
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 48:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:169
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 49:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:172
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: PLUS, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 50:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:173
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: SUB, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 51:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:174
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MUL, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 52:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:175
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MOD, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 53:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:176
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: DIV, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 54:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:177
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 55:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:178
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 56:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:179
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 57:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:180
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 58:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:181
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: EQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 59:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:182
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: NEQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 60:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:183
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: AND, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 61:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:184
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: OR, Right: parserDollar[3].expr, Pos: parserDollar[1].pos}
		}
	case 62:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:185
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 63:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:188
		{
			parserVAL.arrayliter = ArrayLiter{parserDollar[1].pos, parserDollar[2].exprs}
		}
	case 64:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:190
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 65:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:191
		{
			parserVAL.exprs = []Evaluation{parserDollar[1].expr}
		}
	case 66:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:192
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 67:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:194
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserDollar[1].ident, Exprs: parserDollar[2].exprs, Pos: parserDollar[1].pos}
		}
	case 68:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:196
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 69:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:197
		{
			parserVAL.exprs = []Evaluation{parserDollar[2].expr}
		}
	case 70:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:200
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 71:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:203
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 72:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:204
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 73:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:206
		{
			parserVAL.typedefinition = PairType{FstType: parserDollar[3].pairelemtype, SndType: parserDollar[5].pairelemtype}
		}
	case 74:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:208
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 75:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:209
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 76:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:210
		{
			parserVAL.pairelemtype = Pair
		}
	case 77:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:212
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 78:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:213
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 79:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:214
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 80:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:216
		{
			parserVAL.typedefinition = Int
		}
	case 81:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:217
		{
			parserVAL.typedefinition = Bool
		}
	case 82:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:218
		{
			parserVAL.typedefinition = Char
		}
	case 83:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:219
		{
			parserVAL.typedefinition = String
		}
	case 84:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:221
		{
			parserVAL.typedefinition = ArrayType{Type: parserDollar[1].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
