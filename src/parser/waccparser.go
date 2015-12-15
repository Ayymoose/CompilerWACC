//line waccparser.y:1
package parser

import __yyfmt__ "fmt"

//line waccparser.y:3
import (
	. "ast"
)

//line waccparser.y:10
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
	classes        []*Class
	class          *Class
	stmt           Statement
	stmts          []Statement
	assignrhs      Evaluation
	assignlhs      Evaluation
	expr           Evaluation
	exprs          []Evaluation
	params         []Param
	param          Param
	fields         []Field
	field          Field
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
	"CLASS",
	"OPEN",
	"CLOSE",
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
	"FOR",
	"STRINGCONST",
	"IDENTIFIER",
	"INTEGER",
	"CHARACTER",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line waccparser.y:278

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 136,
	56, 94,
	-2, 91,
	-1, 137,
	56, 95,
	-2, 92,
}

const parserNprod = 102
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 527

var parserAct = [...]int{

	119, 7, 12, 168, 145, 141, 25, 4, 135, 36,
	37, 36, 37, 167, 38, 45, 62, 63, 64, 65,
	66, 24, 69, 68, 70, 53, 42, 215, 118, 192,
	34, 35, 42, 42, 42, 42, 208, 85, 86, 34,
	42, 179, 184, 180, 117, 113, 34, 177, 34, 28,
	41, 44, 202, 216, 209, 105, 106, 107, 108, 109,
	110, 111, 71, 114, 90, 211, 34, 10, 94, 96,
	95, 92, 93, 103, 104, 98, 100, 97, 99, 101,
	102, 224, 221, 67, 134, 29, 30, 31, 32, 33,
	188, 137, 217, 148, 149, 150, 151, 152, 153, 154,
	155, 156, 157, 158, 159, 160, 136, 40, 121, 73,
	40, 40, 163, 131, 162, 164, 143, 165, 128, 147,
	40, 40, 121, 116, 74, 169, 170, 75, 171, 40,
	172, 173, 185, 174, 175, 130, 127, 191, 34, 39,
	192, 34, 182, 89, 183, 132, 90, 139, 178, 129,
	142, 76, 146, 94, 96, 95, 92, 93, 103, 104,
	98, 100, 97, 99, 101, 102, 72, 189, 36, 37,
	29, 30, 31, 32, 33, 40, 207, 187, 126, 91,
	40, 89, 190, 198, 137, 197, 194, 201, 200, 84,
	203, 169, 186, 206, 205, 121, 40, 125, 166, 136,
	42, 210, 133, 212, 83, 214, 222, 34, 28, 213,
	199, 34, 181, 34, 87, 218, 98, 100, 97, 99,
	220, 26, 9, 6, 223, 34, 94, 96, 95, 92,
	93, 103, 104, 98, 100, 97, 99, 101, 102, 34,
	139, 196, 142, 34, 2, 27, 146, 219, 34, 122,
	123, 36, 37, 80, 79, 81, 77, 78, 54, 43,
	55, 56, 57, 88, 51, 82, 59, 58, 120, 140,
	144, 83, 29, 30, 31, 32, 33, 8, 47, 48,
	61, 124, 60, 5, 3, 1, 23, 0, 22, 0,
	50, 52, 46, 49, 11, 13, 14, 15, 16, 17,
	18, 19, 0, 0, 0, 21, 0, 0, 0, 195,
	36, 37, 29, 30, 31, 32, 33, 0, 0, 54,
	0, 55, 56, 57, 0, 0, 0, 59, 58, 29,
	30, 31, 32, 33, 29, 30, 31, 32, 138, 47,
	48, 61, 0, 60, 0, 0, 0, 0, 20, 0,
	28, 50, 52, 46, 49, 94, 96, 95, 92, 93,
	103, 104, 98, 100, 97, 99, 101, 102, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 204, 94,
	96, 95, 92, 93, 103, 104, 98, 100, 97, 99,
	101, 102, 0, 0, 0, 0, 0, 0, 0, 0,
	161, 94, 96, 95, 92, 93, 103, 104, 98, 100,
	97, 99, 101, 102, 0, 0, 0, 0, 0, 0,
	0, 193, 94, 96, 95, 92, 93, 103, 104, 98,
	100, 97, 99, 101, 102, 115, 0, 0, 94, 96,
	95, 0, 176, 0, 112, 98, 100, 97, 99, 0,
	0, 94, 96, 95, 92, 93, 103, 104, 98, 100,
	97, 99, 101, 102, 94, 96, 95, 92, 93, 103,
	104, 98, 100, 97, 99, 101, 102, 94, 96, 95,
	92, 93, 103, 104, 98, 100, 97, 99, 101, 102,
	94, 96, 95, 92, 93, 103, 0, 98, 100, 97,
	99, 101, 102, 94, 96, 95, 92, 93, 0, 0,
	98, 100, 97, 99, 101, 102, 94, 96, 95, 92,
	93, 0, 0, 98, 100, 97, 99,
}
var parserPact = [...]int{

	240, -1000, -1000, 217, 284, -1000, -52, 134, -1000, -1000,
	-16, -1000, -1000, -15, 286, 286, 286, 286, 286, 286,
	142, 286, 284, 104, -1000, -1000, -1000, 91, 215, -1000,
	-1000, -1000, -1000, 132, -1000, -1000, 286, 286, 207, -1000,
	284, 86, 121, -1000, 148, 439, -1000, -1000, -1000, -1000,
	-1000, -1000, 148, -1000, 286, 286, 286, 286, 286, 286,
	286, -1000, 439, 439, 439, 439, 426, -21, 1, 413,
	118, -22, -1000, -1000, -1000, -1000, 225, 137, 76, 89,
	75, 85, 146, 286, 306, 439, 439, 244, -1000, 57,
	225, -1000, 286, 286, 286, 286, 286, 286, 286, 286,
	286, 286, 286, 286, 286, -1000, -1000, -1000, -1000, 400,
	400, 341, 284, 52, 286, 284, -1000, 4, -1000, 439,
	-1000, -1000, 141, -53, 286, 286, -1000, 286, -1000, 286,
	286, -1000, 286, 286, 384, -14, -1000, -1000, 132, 144,
	-20, -1000, -23, 203, 83, -1000, -24, -1000, 400, 400,
	171, 171, 171, -1000, -1000, -1000, -1000, 478, 478, 465,
	452, -1000, 113, 225, 115, 67, 286, 125, 79, 439,
	439, 439, 439, 439, 439, 363, -1000, 306, 301, 244,
	-1000, 284, 201, 244, -1000, 284, -10, -17, -1000, 317,
	286, -1000, 286, -1000, 117, -1000, -30, -1000, 49, 284,
	-1000, 45, 286, 187, 286, -32, 439, -1000, 124, -1000,
	48, -1000, 30, 284, 188, -1000, -1000, -17, 59, -1000,
	184, -1000, 284, 58, -1000,
}
var parserPgo = [...]int{

	0, 285, 284, 283, 7, 277, 222, 1, 28, 245,
	0, 3, 270, 4, 269, 5, 31, 268, 25, 265,
	264, 21, 62, 6, 221, 8, 2,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	17, 11, 11, 11, 18, 19, 19, 20, 16, 16,
	24, 25, 25, 25, 22, 22, 22, 21, 21, 21,
	21, 23,
}
var parserR2 = [...]int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 3, 1, 1, 4, 1, 2, 2,
	2, 2, 2, 2, 7, 12, 9, 5, 3, 2,
	2, 2, 2, 3, 4, 4, 4, 4, 4, 3,
	3, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 0, 2, 4, 3, 1, 2, 2,
	6, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3,
}
var parserChk = [...]int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	-22, 10, -26, 11, 12, 13, 14, 15, 16, 17,
	64, 21, 4, 2, -21, -23, -24, -9, 66, 28,
	29, 30, 31, 32, -18, -16, 26, 27, 66, 5,
	62, 66, 56, -9, 66, -10, 67, 53, 54, 68,
	65, -20, 66, -18, 33, 35, 36, 37, 42, 41,
	57, 55, -10, -10, -10, -10, -10, -22, -26, -10,
	-7, -22, 62, 5, 20, 23, 60, 41, 42, 39,
	38, 40, -19, 56, 57, -10, -10, 7, -6, 57,
	60, 58, 41, 42, 38, 40, 39, 47, 45, 48,
	46, 49, 50, 43, 44, -10, -10, -10, -10, -10,
	-10, -10, 18, 66, 62, 22, 5, 66, -8, -10,
	-17, -16, 24, 25, 56, 60, 41, 60, 42, 60,
	60, 38, 60, 56, -10, -25, -21, -23, 32, -22,
	-14, -15, -22, 59, -12, -13, -22, -8, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, 59, -7, 60, -10, -7, 57, 66, -11, -10,
	-10, -10, -10, -10, -10, -10, 58, 61, -4, 61,
	66, 9, 59, 61, 66, 19, -8, 62, 23, -10,
	57, 58, 61, 58, -25, 8, -22, -15, -7, 9,
	-13, -7, 62, -26, 61, -11, -10, 59, 66, 5,
	-7, 20, -10, 22, -10, 59, 5, 62, -7, 59,
	-26, 23, 22, -7, 23,
}
var parserDef = [...]int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 24,
	0, 25, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 94, 95, 96, 0, 15, 97,
	98, 99, 100, 0, 16, 17, 0, 0, 0, 1,
	0, 0, 0, 28, 15, 29, 52, 53, 54, 55,
	56, 57, 58, 59, 0, 0, 0, 0, 0, 0,
	0, 87, 30, 31, 32, 33, 0, 0, 0, 0,
	0, 0, 39, 40, 41, 42, 0, 0, 0, 0,
	0, 0, 84, 0, 0, 88, 89, 0, 23, 0,
	0, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 60, 61, 62, 63, 64,
	65, 0, 0, 0, 0, 0, 38, 0, 43, 18,
	19, 20, 0, 0, 83, 0, 49, 0, 50, 0,
	0, 51, 0, 0, 0, 0, -2, -2, 93, 0,
	9, 6, 0, 0, 0, 13, 0, 26, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 0, 0, 0, 0, 0, 0, 0, 82,
	44, 45, 46, 47, 48, 0, 86, 0, 0, 0,
	7, 0, 0, 0, 14, 0, 0, 0, 37, 0,
	83, 80, 0, 85, 0, 4, 0, 5, 0, 0,
	12, 0, 0, 0, 0, 0, 81, 90, 0, 10,
	0, 34, 0, 0, 0, 22, 11, 0, 0, 21,
	0, 36, 0, 0, 35,
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
	62, 63, 64, 65, 66, 67, 68,
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
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line waccparser.y:111
		{
			parserlex.(*Lexer).prog = &Program{ClassList: parserDollar[2].classes, FunctionList: parserDollar[3].functions, StatList: parserDollar[4].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 2:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:116
		{
			parserVAL.classes = append(parserDollar[1].classes, parserDollar[2].class)
		}
	case 3:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line waccparser.y:117
		{
			parserVAL.classes = []*Class{}
		}
	case 4:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line waccparser.y:120
		{
			if !checkClassIdent(parserDollar[2].ident) {
				parserlex.Error("Invalid class name")
			}
			parserVAL.class = &Class{Ident: ClassType(parserDollar[2].ident), FieldList: parserDollar[4].fields, FunctionList: parserDollar[5].functions}
		}
	case 5:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:126
		{
			parserVAL.fields = append(parserDollar[1].fields, parserDollar[3].field)
		}
	case 6:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:127
		{
			parserVAL.fields = []Field{parserDollar[1].field}
		}
	case 7:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:129
		{
			parserVAL.field = Field{FieldType: parserDollar[1].typedefinition, Ident: parserDollar[2].ident}
		}
	case 8:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:131
		{
			parserVAL.functions = append(parserDollar[1].functions, parserDollar[2].function)
		}
	case 9:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line waccparser.y:132
		{
			parserVAL.functions = []*Function{}
		}
	case 10:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line waccparser.y:135
		{
			if !checkStats(parserDollar[6].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserDollar[2].ident, ReturnType: parserDollar[1].typedefinition, StatList: parserDollar[6].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		parserDollar = parserS[parserpt-8 : parserpt+1]
		//line waccparser.y:141
		{
			if !checkStats(parserDollar[7].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserDollar[2].ident, ReturnType: parserDollar[1].typedefinition, StatList: parserDollar[7].stmts, ParameterTypes: parserDollar[4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 12:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:147
		{
			parserVAL.params = append(parserDollar[1].params, parserDollar[3].param)
		}
	case 13:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:148
		{
			parserVAL.params = []Param{parserDollar[1].param}
		}
	case 14:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:150
		{
			parserVAL.param = Param{ParamType: parserDollar[1].typedefinition, Ident: parserDollar[2].ident}
		}
	case 15:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:152
		{
			parserVAL.assignlhs = parserDollar[1].ident
		}
	case 16:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:153
		{
			parserVAL.assignlhs = parserDollar[1].arrayelem
		}
	case 17:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:154
		{
			parserVAL.assignlhs = parserDollar[1].pairelem
		}
	case 18:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:156
		{
			parserVAL.assignrhs = parserDollar[1].expr
		}
	case 19:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:157
		{
			parserVAL.assignrhs = parserDollar[1].arrayliter
		}
	case 20:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:158
		{
			parserVAL.assignrhs = parserDollar[1].pairelem
		}
	case 21:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line waccparser.y:159
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserDollar[3].expr, SndExpr: parserDollar[5].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line waccparser.y:160
		{
			parserVAL.assignrhs = Call{Ident: parserDollar[2].ident, ParamList: parserDollar[4].exprs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:162
		{
			parserVAL.stmts = append(parserDollar[1].stmts, parserDollar[3].stmt)
		}
	case 24:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:163
		{
			parserVAL.stmts = []Statement{parserDollar[1].stmt}
		}
	case 25:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:166
		{
			parserVAL.stmt = Skip{Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 26:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:167
		{
			parserVAL.stmt = Declare{DecType: parserDollar[1].typedefinition, Lhs: parserDollar[2].ident, Rhs: parserDollar[4].assignrhs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 27:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:168
		{
			parserVAL.stmt = parserDollar[1].stmt
		}
	case 28:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:169
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].assignlhs}
		}
	case 29:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:170
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 30:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:171
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 31:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:172
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 32:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:173
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 33:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:174
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 34:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line waccparser.y:175
		{
			parserVAL.stmt = If{Conditional: parserDollar[2].expr, ThenStat: parserDollar[4].stmts, ElseStat: parserDollar[6].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 35:
		parserDollar = parserS[parserpt-12 : parserpt+1]
		//line waccparser.y:177
		{
			stats := append(parserDollar[11].stmts, parserDollar[9].stmt)
			parserVAL.stmt = While{Conditional: parserDollar[7].expr, DoStat: stats, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 36:
		parserDollar = parserS[parserpt-9 : parserpt+1]
		//line waccparser.y:181
		{
			stats := append(parserDollar[8].stmts, parserDollar[6].stmt)
			parserVAL.stmt = While{Conditional: parserDollar[4].expr, DoStat: stats, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 37:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line waccparser.y:185
		{
			parserVAL.stmt = While{Conditional: parserDollar[2].expr, DoStat: parserDollar[4].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 38:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:186
		{
			parserVAL.stmt = Scope{StatList: parserDollar[2].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 39:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:187
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 40:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:191
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 41:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:194
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 42:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:198
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 43:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:203
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].assignlhs, Rhs: parserDollar[3].assignrhs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:204
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: PLUS, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:205
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: SUB, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:206
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: DIV, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:207
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MUL, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:208
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MOD, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:209
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: PLUS, Right: Integer(1), Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:210
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: SUB, Right: Integer(1), Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:211
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MUL, Right: parserDollar[1].ident, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:213
		{
			parserVAL.expr = parserDollar[1].integer
		}
	case 53:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:214
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 54:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:215
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 55:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:216
		{
			parserVAL.expr = parserDollar[1].character
		}
	case 56:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:217
		{
			parserVAL.expr = parserDollar[1].stringconst
		}
	case 57:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:218
		{
			parserVAL.expr = parserDollar[1].pairliter
		}
	case 58:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:219
		{
			parserVAL.expr = parserDollar[1].ident
		}
	case 59:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:220
		{
			parserVAL.expr = parserDollar[1].arrayelem
		}
	case 60:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:221
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:222
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:223
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:224
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:225
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:226
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 66:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:227
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: PLUS, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:228
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: SUB, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:229
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MUL, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:230
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MOD, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:231
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: DIV, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:232
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:233
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:234
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:235
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:236
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: EQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: NEQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: AND, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: OR, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:240
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 80:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:243
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].exprs}
		}
	case 81:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:245
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 82:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:246
		{
			parserVAL.exprs = []Evaluation{parserDollar[1].expr}
		}
	case 83:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line waccparser.y:247
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 84:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:249
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserDollar[1].ident, Exprs: parserDollar[2].exprs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 85:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:251
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 86:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:252
		{
			parserVAL.exprs = []Evaluation{parserDollar[2].expr}
		}
	case 87:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:255
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 88:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:258
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 89:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:259
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 90:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line waccparser.y:261
		{
			parserVAL.typedefinition = PairType{FstType: parserDollar[3].pairelemtype, SndType: parserDollar[5].pairelemtype}
		}
	case 91:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:263
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 92:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:264
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 93:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:265
		{
			parserVAL.pairelemtype = Pair
		}
	case 94:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:267
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 95:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:268
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 96:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:269
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 97:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:271
		{
			parserVAL.typedefinition = Int
		}
	case 98:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:272
		{
			parserVAL.typedefinition = Bool
		}
	case 99:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:273
		{
			parserVAL.typedefinition = Char
		}
	case 100:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:274
		{
			parserVAL.typedefinition = String
		}
	case 101:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:276
		{
			parserVAL.typedefinition = ArrayType{Type: parserDollar[1].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
