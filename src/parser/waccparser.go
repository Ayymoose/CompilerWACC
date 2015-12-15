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

//line waccparser.y:282

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 139,
	56, 95,
	-2, 92,
	-1, 140,
	56, 96,
	-2, 93,
}

const parserNprod = 103
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 511

var parserAct = [...]int{

	122, 144, 12, 7, 148, 26, 25, 55, 138, 4,
	151, 172, 35, 44, 39, 43, 47, 64, 65, 66,
	67, 68, 35, 71, 70, 213, 43, 72, 229, 35,
	43, 35, 43, 43, 36, 208, 188, 121, 87, 88,
	184, 117, 120, 116, 37, 38, 226, 216, 191, 35,
	37, 38, 30, 31, 32, 33, 34, 108, 109, 110,
	111, 112, 113, 114, 221, 73, 197, 41, 214, 190,
	10, 97, 99, 98, 95, 96, 106, 107, 101, 103,
	100, 102, 104, 105, 29, 41, 137, 69, 92, 41,
	29, 41, 140, 139, 210, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 220,
	43, 190, 183, 124, 119, 212, 181, 75, 169, 167,
	42, 41, 170, 35, 168, 41, 35, 124, 152, 174,
	150, 175, 76, 176, 177, 77, 178, 179, 173, 97,
	99, 98, 95, 96, 106, 107, 101, 103, 100, 102,
	104, 105, 142, 182, 24, 145, 23, 149, 37, 38,
	224, 194, 11, 14, 15, 16, 17, 18, 19, 20,
	40, 41, 195, 22, 74, 135, 131, 13, 37, 38,
	30, 31, 32, 33, 34, 202, 129, 140, 139, 203,
	199, 206, 205, 35, 130, 207, 209, 152, 46, 35,
	41, 35, 189, 124, 190, 128, 192, 211, 215, 217,
	134, 219, 35, 186, 91, 187, 21, 92, 29, 132,
	78, 93, 223, 91, 196, 225, 35, 41, 86, 43,
	35, 228, 133, 171, 94, 35, 30, 31, 32, 33,
	34, 136, 85, 125, 126, 37, 38, 142, 201, 145,
	227, 9, 56, 149, 57, 58, 59, 218, 204, 185,
	61, 60, 82, 81, 83, 79, 80, 146, 101, 103,
	100, 102, 49, 50, 63, 127, 62, 6, 200, 28,
	85, 89, 2, 27, 52, 54, 48, 51, 56, 53,
	57, 58, 59, 90, 45, 84, 61, 60, 30, 31,
	32, 33, 34, 30, 31, 32, 33, 34, 49, 50,
	63, 123, 62, 30, 31, 32, 33, 141, 143, 147,
	52, 54, 48, 51, 97, 99, 98, 95, 96, 106,
	107, 101, 103, 100, 102, 104, 105, 97, 99, 98,
	95, 96, 8, 5, 101, 103, 100, 102, 222, 97,
	99, 98, 95, 96, 106, 107, 101, 103, 100, 102,
	104, 105, 97, 99, 98, 3, 1, 0, 0, 101,
	103, 100, 102, 193, 97, 99, 98, 95, 96, 106,
	107, 101, 103, 100, 102, 104, 105, 0, 0, 0,
	0, 0, 0, 0, 0, 166, 97, 99, 98, 95,
	96, 106, 107, 101, 103, 100, 102, 104, 105, 0,
	0, 0, 0, 0, 0, 0, 198, 97, 99, 98,
	95, 96, 106, 107, 101, 103, 100, 102, 104, 105,
	118, 0, 0, 0, 0, 0, 0, 180, 0, 115,
	0, 0, 0, 0, 0, 0, 97, 99, 98, 95,
	96, 106, 107, 101, 103, 100, 102, 104, 105, 97,
	99, 98, 95, 96, 106, 107, 101, 103, 100, 102,
	104, 105, 97, 99, 98, 95, 96, 106, 107, 101,
	103, 100, 102, 104, 105, 97, 99, 98, 95, 96,
	106, 0, 101, 103, 100, 102, 104, 105, 97, 99,
	98, 95, 96, 0, 0, 101, 103, 100, 102, 104,
	105,
}
var parserPact = [...]int{

	278, -1000, -1000, 271, 152, -1000, -52, 165, -1000, -1000,
	54, -1000, -1000, -53, 132, 255, 255, 255, 255, 255,
	255, 24, 255, 152, 112, -1000, -1000, -1000, 160, 224,
	-1000, -1000, -1000, -1000, 171, -1000, -1000, 255, 255, 274,
	-1000, 152, 157, 163, 177, -1000, 186, 434, -1000, -1000,
	-1000, -1000, -1000, -1000, 186, -1000, 255, 255, 255, 255,
	255, 255, 255, -1000, 434, 434, 434, 434, 421, -23,
	-21, 408, 109, -24, -1000, -1000, -1000, -1000, 219, 145,
	134, 159, 172, 115, 185, 255, 285, 434, 434, 275,
	-1000, 208, 219, -1000, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, -1000, -1000,
	-1000, -1000, 324, 324, 336, 152, 64, 255, 152, -1000,
	28, -1000, 434, -1000, -1000, 176, -55, 255, 255, -1000,
	255, -1000, 255, 255, -1000, 255, 255, 379, 55, -1000,
	-1000, 171, 173, 51, -1000, -26, 250, 154, -1000, -30,
	-1000, 143, 434, 324, 324, 223, 223, 223, -1000, -1000,
	-1000, -1000, 299, 299, 460, 447, -1000, 29, 219, 311,
	138, 255, 167, 8, 434, 434, 434, 434, 434, 358,
	-1000, 285, 270, 275, -1000, 152, 249, 275, -1000, -1000,
	255, 152, -27, 18, -1000, 33, 255, -1000, -1000, 56,
	-1000, -41, -1000, 63, 152, -1000, 434, 27, 255, 235,
	255, 50, -1000, 166, -1000, 59, -1000, 286, 152, 101,
	-1000, -1000, 18, 23, -1000, 228, -1000, 152, 5, -1000,
}
var parserPgo = [...]int{

	0, 366, 365, 343, 9, 342, 251, 3, 37, 279,
	0, 10, 319, 4, 318, 1, 34, 311, 7, 295,
	289, 6, 65, 5, 283, 8, 2,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 17, 11, 11, 11, 18, 19, 19, 20, 16,
	16, 24, 25, 25, 25, 22, 22, 22, 21, 21,
	21, 21, 23,
}
var parserR2 = [...]int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 3, 1, 1, 4, 1, 5, 2,
	2, 2, 2, 2, 2, 7, 12, 9, 5, 3,
	2, 2, 2, 2, 3, 4, 4, 4, 4, 4,
	3, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 0, 2, 4, 3, 1, 2,
	2, 6, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3,
}
var parserChk = [...]int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	-22, 10, -26, 25, 11, 12, 13, 14, 15, 16,
	17, 64, 21, 4, 2, -21, -23, -24, -9, 66,
	28, 29, 30, 31, 32, -18, -16, 26, 27, 66,
	5, 62, 66, 56, 66, -9, 66, -10, 67, 53,
	54, 68, 65, -20, 66, -18, 33, 35, 36, 37,
	42, 41, 57, 55, -10, -10, -10, -10, -10, -22,
	-26, -10, -7, -22, 62, 5, 20, 23, 60, 41,
	42, 39, 38, 40, -19, 56, 57, -10, -10, 7,
	-6, 57, 60, 58, 57, 41, 42, 38, 40, 39,
	47, 45, 48, 46, 49, 50, 43, 44, -10, -10,
	-10, -10, -10, -10, -10, 18, 66, 62, 22, 5,
	66, -8, -10, -17, -16, 24, 25, 56, 60, 41,
	60, 42, 60, 60, 38, 60, 56, -10, -25, -21,
	-23, 32, -22, -14, -15, -22, 59, -12, -13, -22,
	-8, -11, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, 59, -7, 60, -10,
	-7, 57, 66, -11, -10, -10, -10, -10, -10, -10,
	58, 61, -4, 61, 66, 9, 59, 61, 66, 59,
	61, 19, -8, 62, 23, -10, 57, 58, 58, -25,
	8, -22, -15, -7, 9, -13, -10, -7, 62, -26,
	61, -11, 59, 66, 5, -7, 20, -10, 22, -10,
	59, 5, 62, -7, 59, -26, 23, 22, -7, 23,
}
var parserDef = [...]int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 24,
	0, 25, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 95, 96, 97, 0, 15,
	98, 99, 100, 101, 0, 16, 17, 0, 0, 0,
	1, 0, 0, 0, 0, 29, 15, 30, 53, 54,
	55, 56, 57, 58, 59, 60, 0, 0, 0, 0,
	0, 0, 0, 88, 31, 32, 33, 34, 0, 0,
	0, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	0, 0, 0, 0, 85, 0, 0, 89, 90, 0,
	23, 0, 0, 102, 84, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 61, 62,
	63, 64, 65, 66, 0, 0, 0, 0, 0, 39,
	0, 44, 18, 19, 20, 0, 0, 84, 0, 50,
	0, 51, 0, 0, 52, 0, 0, 0, 0, -2,
	-2, 94, 0, 9, 6, 0, 0, 0, 13, 0,
	26, 0, 83, 67, 68, 69, 70, 71, 72, 73,
	74, 75, 76, 77, 78, 79, 80, 0, 0, 0,
	0, 0, 0, 0, 45, 46, 47, 48, 49, 0,
	87, 0, 0, 0, 7, 0, 0, 0, 14, 28,
	0, 0, 0, 0, 38, 0, 84, 81, 86, 0,
	4, 0, 5, 0, 0, 12, 82, 0, 0, 0,
	0, 0, 91, 0, 10, 0, 35, 0, 0, 0,
	22, 11, 0, 0, 21, 0, 37, 0, 0, 36,
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
		//line waccparser.y:169
		{
			parserVAL.stmt = parserDollar[1].stmt
		}
	case 28:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line waccparser.y:171
		{
			parserVAL.stmt = Call{Ident: parserDollar[2].ident, ParamList: parserDollar[4].exprs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 29:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:173
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].assignlhs}
		}
	case 30:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:174
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 31:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:175
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 32:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:176
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 33:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:177
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 34:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:178
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].expr}
		}
	case 35:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line waccparser.y:179
		{
			parserVAL.stmt = If{Conditional: parserDollar[2].expr, ThenStat: parserDollar[4].stmts, ElseStat: parserDollar[6].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 36:
		parserDollar = parserS[parserpt-12 : parserpt+1]
		//line waccparser.y:181
		{
			stats := append(parserDollar[11].stmts, parserDollar[9].stmt)
			parserVAL.stmt = While{Conditional: parserDollar[7].expr, DoStat: stats, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 37:
		parserDollar = parserS[parserpt-9 : parserpt+1]
		//line waccparser.y:185
		{
			stats := append(parserDollar[8].stmts, parserDollar[6].stmt)
			parserVAL.stmt = While{Conditional: parserDollar[4].expr, DoStat: stats, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 38:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line waccparser.y:189
		{
			parserVAL.stmt = While{Conditional: parserDollar[2].expr, DoStat: parserDollar[4].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 39:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:190
		{
			parserVAL.stmt = Scope{StatList: parserDollar[2].stmts, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
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
		//line waccparser.y:195
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
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:202
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 44:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:207
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].assignlhs, Rhs: parserDollar[3].assignrhs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:208
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: PLUS, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:209
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: SUB, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:210
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: DIV, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:211
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MUL, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:212
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MOD, Right: parserDollar[4].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:213
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: PLUS, Right: Integer(1), Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: SUB, Right: Integer(1), Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserDollar[1].ident, Rhs: Binop{Left: parserDollar[1].ident, Binary: MUL, Right: parserDollar[1].ident, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:217
		{
			parserVAL.expr = parserDollar[1].integer
		}
	case 54:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:218
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 55:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:219
		{
			parserVAL.expr = parserDollar[1].boolean
		}
	case 56:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:220
		{
			parserVAL.expr = parserDollar[1].character
		}
	case 57:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:221
		{
			parserVAL.expr = parserDollar[1].stringconst
		}
	case 58:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:222
		{
			parserVAL.expr = parserDollar[1].pairliter
		}
	case 59:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:223
		{
			parserVAL.expr = parserDollar[1].ident
		}
	case 60:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:224
		{
			parserVAL.expr = parserDollar[1].arrayelem
		}
	case 61:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:225
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:226
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:227
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:228
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:229
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:230
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 67:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:231
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: PLUS, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:232
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: SUB, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:233
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MUL, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:234
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: MOD, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:235
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: DIV, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:236
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GT, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: LTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: GTE, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: EQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: NEQ, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: AND, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserDollar[1].expr, Binary: OR, Right: parserDollar[3].expr, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:244
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 81:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:247
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserDollar[1].pos, parserDollar[2].exprs}
		}
	case 82:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:249
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 83:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:250
		{
			parserVAL.exprs = []Evaluation{parserDollar[1].expr}
		}
	case 84:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line waccparser.y:251
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 85:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:253
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserDollar[1].ident, Exprs: parserDollar[2].exprs, Pos: parserDollar[1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 86:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line waccparser.y:255
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 87:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:256
		{
			parserVAL.exprs = []Evaluation{parserDollar[2].expr}
		}
	case 88:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:259
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 89:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:262
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 90:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line waccparser.y:263
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserDollar[2].expr, Pos: parserDollar[1].pos}
		}
	case 91:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line waccparser.y:265
		{
			parserVAL.typedefinition = PairType{FstType: parserDollar[3].pairelemtype, SndType: parserDollar[5].pairelemtype}
		}
	case 92:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:267
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 93:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:268
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 94:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:269
		{
			parserVAL.pairelemtype = Pair
		}
	case 95:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:271
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 96:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:272
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 97:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:273
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 98:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:275
		{
			parserVAL.typedefinition = Int
		}
	case 99:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:276
		{
			parserVAL.typedefinition = Bool
		}
	case 100:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:277
		{
			parserVAL.typedefinition = Char
		}
	case 101:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line waccparser.y:278
		{
			parserVAL.typedefinition = String
		}
	case 102:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line waccparser.y:280
		{
			parserVAL.typedefinition = ArrayType{Type: parserDollar[1].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
