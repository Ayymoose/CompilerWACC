//line oliver.y:1
package parser

import __yyfmt__ "fmt"

//line oliver.y:3
//line oliver.y:5
type parserSymType struct {
	yys            int
	str            string
	number         int
	functions      []*Function
	function       *Function
	stmt           interface{}
	stmts          []interface{}
	assignrhs      interface{}
	assignlhs      interface{}
	expr           interface{}
	exprs          []interface{}
	params         []Param
	param          Param
	arglist        []interface{}
	bracketed      []interface{}
	pairliter      interface{}
	arrayliter     ArrayLiter
	pairelem       PairElem
	arrayelem      ArrayElem
	typedefinition Type
	pairelemtype   Type
}

const BEGIN = 57346
const END = 57347
const IS = 57348
const SKIP = 57349
const READ = 57350
const FREE = 57351
const RETURN = 57352
const EXIT = 57353
const PRINT = 57354
const PRINTLN = 57355
const IF = 57356
const THEN = 57357
const ELSE = 57358
const FI = 57359
const WHILE = 57360
const DO = 57361
const DONE = 57362
const NEWPAIR = 57363
const CALL = 57364
const FST = 57365
const SND = 57366
const INT = 57367
const BOOL = 57368
const CHAR = 57369
const STRING = 57370
const PAIR = 57371
const NOT = 57372
const NEG = 57373
const LEN = 57374
const ORD = 57375
const CHR = 57376
const MUL = 57377
const DIV = 57378
const MOD = 57379
const PLUS = 57380
const SUB = 57381
const AND = 57382
const OR = 57383
const GT = 57384
const GTE = 57385
const LT = 57386
const LTE = 57387
const EQ = 57388
const NEQ = 57389
const POSITIVE = 57390
const NEGATIVE = 57391
const TRUE = 57392
const FALSE = 57393
const NULL = 57394
const OPENSQUARE = 57395
const OPENROUND = 57396
const CLOSESQUARE = 57397
const CLOSEROUND = 57398
const ASSIGNMENT = 57399
const COMMA = 57400
const SEMICOLON = 57401
const ERROR = 57402
const STRINGCONST = 57403
const IDENTIFIER = 57404
const INTEGER = 57405
const CHARACTER = 57406

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

//line oliver.y:208

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 106,
	53, 89,
	-2, 86,
	-1, 107,
	53, 90,
	-2, 87,
}

const parserNprod = 97
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 614

var parserAct = [...]int{

	73, 4, 24, 154, 72, 20, 117, 150, 116, 105,
	9, 112, 38, 55, 56, 57, 58, 59, 60, 35,
	61, 37, 30, 31, 206, 35, 35, 19, 176, 35,
	147, 66, 67, 62, 141, 102, 46, 7, 34, 75,
	23, 25, 26, 27, 28, 29, 196, 23, 92, 93,
	94, 95, 96, 97, 98, 23, 188, 101, 32, 173,
	137, 22, 146, 33, 187, 104, 145, 71, 193, 33,
	23, 107, 110, 75, 144, 114, 139, 145, 140, 118,
	119, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 106, 69, 177, 143, 70, 33, 109,
	33, 133, 134, 113, 135, 33, 70, 36, 6, 65,
	171, 33, 33, 115, 35, 103, 142, 81, 83, 82,
	79, 80, 90, 91, 85, 87, 84, 86, 88, 89,
	64, 25, 26, 27, 28, 29, 23, 23, 212, 210,
	167, 209, 68, 107, 118, 205, 169, 149, 170, 157,
	168, 165, 166, 174, 175, 25, 26, 27, 28, 108,
	179, 180, 181, 182, 183, 106, 184, 138, 186, 178,
	197, 109, 156, 185, 113, 23, 157, 191, 172, 2,
	75, 21, 194, 23, 157, 44, 63, 74, 111, 151,
	199, 200, 5, 3, 1, 23, 75, 203, 202, 156,
	201, 23, 23, 157, 0, 0, 0, 156, 0, 0,
	23, 0, 207, 208, 0, 157, 157, 211, 0, 0,
	157, 0, 0, 0, 0, 0, 156, 0, 0, 23,
	0, 0, 23, 0, 0, 0, 0, 0, 156, 156,
	0, 23, 23, 156, 0, 0, 23, 76, 77, 30,
	31, 0, 0, 0, 0, 0, 47, 0, 48, 49,
	50, 0, 0, 0, 52, 51, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 40, 41, 54, 78,
	53, 0, 0, 0, 0, 0, 0, 43, 45, 39,
	42, 164, 0, 0, 155, 158, 159, 153, 160, 161,
	162, 152, 0, 0, 0, 163, 0, 0, 0, 0,
	30, 31, 25, 26, 27, 28, 29, 18, 0, 0,
	8, 10, 11, 12, 13, 14, 15, 16, 0, 0,
	0, 17, 0, 0, 0, 0, 30, 31, 25, 26,
	27, 28, 29, 0, 0, 0, 0, 164, 0, 22,
	155, 158, 159, 189, 160, 161, 162, 190, 0, 0,
	0, 163, 0, 0, 0, 0, 30, 31, 25, 26,
	27, 28, 29, 0, 0, 22, 47, 0, 48, 49,
	50, 0, 0, 0, 52, 51, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 40, 41, 54, 0,
	53, 0, 0, 0, 0, 22, 0, 43, 45, 39,
	42, 81, 83, 82, 79, 80, 90, 91, 85, 87,
	84, 86, 88, 89, 0, 0, 0, 0, 0, 0,
	0, 0, 198, 81, 83, 82, 79, 80, 90, 91,
	85, 87, 84, 86, 88, 89, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 81, 83, 82, 79, 80,
	90, 91, 85, 87, 84, 86, 88, 89, 0, 0,
	0, 0, 0, 0, 0, 148, 81, 83, 82, 79,
	80, 90, 91, 85, 87, 84, 86, 88, 89, 204,
	0, 0, 0, 0, 0, 0, 136, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 195, 0, 0, 81,
	83, 82, 79, 80, 90, 91, 85, 87, 84, 86,
	88, 89, 81, 83, 82, 79, 80, 90, 91, 85,
	87, 84, 86, 88, 89, 192, 81, 83, 82, 79,
	80, 90, 91, 85, 87, 84, 86, 88, 89, 0,
	0, 0, 100, 0, 0, 81, 83, 82, 79, 80,
	90, 91, 85, 87, 84, 86, 88, 89, 81, 83,
	82, 79, 80, 90, 91, 85, 87, 84, 86, 88,
	89, 99, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 81, 83, 82, 79, 80, 90, 91, 85, 87,
	84, 86, 88, 89,
}
var parserPact = [...]int{

	175, -1000, -1000, 313, 53, -1000, -1000, -24, -1000, 50,
	-1, 346, 346, 346, 346, 346, 346, 346, 313, -1000,
	-1000, -1000, 77, -1000, -1000, -1000, -1000, -1000, -1000, 55,
	346, 346, -1000, 313, 40, 12, 226, -1000, 501, -1000,
	-1000, -1000, -1000, -1000, -1000, 77, -1000, 346, 346, 346,
	346, 346, 346, 346, -1000, 501, 501, 501, 501, 566,
	533, 52, -27, 62, 346, 130, 501, 501, -1000, 16,
	226, -1000, -1000, 501, -1000, -1000, 59, -54, 346, 346,
	346, 346, 346, 346, 346, 346, 346, 346, 346, 346,
	346, 346, -1000, -1000, -1000, -1000, -1000, -1000, 398, 313,
	313, -1000, 49, 346, 441, 2, -1000, -1000, 55, 61,
	161, 20, -1000, -28, -1000, 346, 42, 19, 501, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 46, 10, 420, -1000, 130, 287, 145,
	106, -1000, 82, 346, -1000, 346, 313, -1000, -1000, 54,
	173, 0, 346, 346, -1000, -1000, -34, 38, -1, 346,
	346, 346, 346, 346, 313, 287, -1000, 346, 8, 501,
	39, -1000, -1000, 343, 520, 501, 11, 226, -1000, 501,
	501, 501, 501, 487, 41, 165, 376, -1000, -1000, 346,
	346, -1000, 287, 226, -1000, 313, -1000, -1000, -1000, 501,
	474, 129, -1000, 4, 287, 287, -1000, 125, 122, 287,
	-1000, 121, -1000,
}
var parserPgo = [...]int{

	0, 194, 193, 192, 108, 3, 1, 189, 7, 4,
	10, 0, 6, 188, 11, 2, 187, 36, 186, 185,
	27, 33, 5, 181, 9,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 3, 8, 8, 8, 8,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	7, 7, 13, 13, 14, 10, 10, 10, 9, 9,
	9, 9, 9, 6, 6, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 16, 12, 12, 12, 17,
	18, 18, 19, 15, 15, 23, 24, 24, 24, 21,
	21, 21, 20, 20, 20, 20, 22,
}
var parserR2 = [...]int{

	0, 4, 2, 0, 7, 8, 4, 9, 7, 2,
	1, 4, 3, 2, 2, 2, 2, 2, 5, 3,
	3, 1, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 3, 1, 1, 4, 3, 2, 2,
	2, 2, 2, 2, 7, 5, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1, 0, 2,
	4, 3, 1, 2, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3,
}
var parserChk = [...]int{

	-1000, -1, 4, -2, -6, -3, -4, -21, 7, -10,
	8, 9, 10, 11, 12, 13, 14, 18, 4, -20,
	-22, -23, 62, -17, -15, 25, 26, 27, 28, 29,
	23, 24, 5, 59, 62, 53, 57, -10, -11, 63,
	50, 51, 64, 61, -19, 62, -17, 30, 32, 33,
	34, 39, 38, 54, 52, -11, -11, -11, -11, -11,
	-11, -6, -21, -18, 53, 54, -11, -11, -4, 54,
	57, 55, -9, -11, -16, -15, 21, 22, 53, 38,
	39, 35, 37, 36, 44, 42, 45, 43, 46, 47,
	40, 41, -11, -11, -11, -11, -11, -11, -11, 15,
	19, 5, 62, 53, -11, -24, -20, -22, 29, -21,
	56, -13, -14, -21, -9, 54, 62, -12, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, 56, -6, -6, -11, 55, 58, 6, 56,
	58, 62, -11, 54, 55, 58, 16, 20, 55, -24,
	-8, -7, 14, 10, -5, 7, -21, -10, 8, 9,
	11, 12, 13, 18, 4, 6, -14, 58, -12, -11,
	-6, 56, 5, 59, -11, -11, 62, 57, -10, -11,
	-11, -11, -11, -11, -6, -8, -11, 56, 17, 10,
	14, -5, 15, 57, -9, 19, 5, 5, 56, -11,
	-11, -8, -9, -6, 15, 16, 20, -8, -8, 16,
	17, -8, 17,
}
var parserDef = [...]int{

	0, -2, 3, 0, 0, 2, 34, 0, 35, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 89,
	90, 91, 25, 26, 27, 92, 93, 94, 95, 0,
	0, 0, 1, 0, 0, 0, 0, 38, 39, 47,
	48, 49, 50, 51, 52, 53, 54, 0, 0, 0,
	0, 0, 0, 0, 82, 40, 41, 42, 43, 0,
	0, 0, 0, 79, 0, 0, 83, 84, 33, 0,
	0, 96, 37, 28, 29, 30, 0, 0, 78, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 55, 56, 57, 58, 59, 60, 0, 0,
	0, 46, 0, 0, 0, 0, -2, -2, 88, 0,
	0, 0, 23, 0, 36, 0, 0, 0, 77, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 0, 0, 0, 81, 0, 0, 0,
	0, 24, 0, 78, 75, 0, 0, 45, 80, 0,
	0, 0, 0, 0, 21, 10, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 22, 0, 0, 76,
	0, 85, 4, 0, 0, 9, 0, 0, 13, 14,
	15, 16, 17, 0, 0, 0, 0, 32, 44, 0,
	0, 20, 0, 0, 12, 0, 19, 5, 31, 6,
	0, 0, 11, 0, 0, 0, 18, 0, 0, 0,
	8, 0, 7,
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
		//line oliver.y:80
		{
			parserlex.(*Lexer).prog = &Program{functionlist: parserDollar[2].functions, statList: parserDollar[3].stmts, symbolTable: &SymbolTable{Table: make(map[string]Type)}}
		}
	case 2:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:84
		{
			parserVAL.functions = append(parserDollar[1].functions, parserDollar[2].function)
		}
	case 3:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:85
		{
			parserVAL.functions = []*Function{}
		}
	case 4:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:88
		{
			parserVAL.function = &Function{ident: parserDollar[2].str, returnType: parserDollar[1].typedefinition, statlist: parserDollar[6].stmts}
		}
	case 5:
		parserDollar = parserS[parserpt-8 : parserpt+1]
		//line oliver.y:91
		{
			parserVAL.function = &Function{ident: parserDollar[2].str, returnType: parserDollar[1].typedefinition, statlist: parserDollar[7].stmts, parameterTypes: parserDollar[4].params}
		}
	case 6:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:95
		{
			parserVAL.stmts = append(parserDollar[1].stmts, Return{parserDollar[4].expr})
		}
	case 7:
		parserDollar = parserS[parserpt-9 : parserpt+1]
		//line oliver.y:96
		{
			parserVAL.stmts = append(parserDollar[1].stmts, If{conditional: parserDollar[4].expr, thenStat: parserDollar[6].stmts, elseStat: parserDollar[8].stmts})
		}
	case 8:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:97
		{
			parserVAL.stmts = []interface{}{If{conditional: parserDollar[2].expr, thenStat: parserDollar[4].stmts, elseStat: parserDollar[6].stmts}}
		}
	case 9:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:98
		{
			parserVAL.stmts = []interface{}{Return{parserDollar[2].expr}}
		}
	case 10:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:100
		{
			parserVAL.stmt = parserDollar[1].number
		}
	case 11:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:101
		{
			parserVAL.stmt = Declare{Type: parserDollar[1].typedefinition, lhs: parserDollar[2].str, rhs: parserDollar[4].assignrhs}
		}
	case 12:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:102
		{
			parserVAL.stmt = Assignment{lhs: parserDollar[1].assignlhs, rhs: parserDollar[3].assignrhs}
		}
	case 13:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:103
		{
			parserVAL.stmt = Read{parserDollar[2].assignlhs}
		}
	case 14:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:104
		{
			parserVAL.stmt = Free{parserDollar[2].expr}
		}
	case 15:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:105
		{
			parserVAL.stmt = Exit{parserDollar[2].expr}
		}
	case 16:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:106
		{
			parserVAL.stmt = Print{parserDollar[2].expr}
		}
	case 17:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:107
		{
			parserVAL.stmt = Println{parserDollar[2].expr}
		}
	case 18:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:108
		{
			parserVAL.stmt = While{conditional: parserDollar[2].expr, doStat: parserDollar[4].stmts}
		}
	case 19:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:109
		{
			parserVAL.stmt = Scope{statlist: parserDollar[2].stmts, symbolTable: &SymbolTable{Table: make(map[string]Type)}}
		}
	case 20:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:111
		{
			parserVAL.stmts = append(parserDollar[1].stmts, parserDollar[3].stmt)
		}
	case 21:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:112
		{
			parserVAL.stmts = []interface{}{parserDollar[1].stmt}
		}
	case 22:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:114
		{
			parserVAL.params = append(parserDollar[1].params, parserDollar[3].param)
		}
	case 23:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:115
		{
			parserVAL.params = []Param{parserDollar[1].param}
		}
	case 24:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:117
		{
			parserVAL.param = Param{paramType: parserDollar[1].typedefinition, ident: parserDollar[2].str}
		}
	case 25:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:119
		{
			parserVAL.assignlhs = parserDollar[1].str
		}
	case 26:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:120
		{
			parserVAL.assignlhs = parserDollar[1].arrayelem
		}
	case 27:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:121
		{
			parserVAL.assignlhs = parserDollar[1].pairelem
		}
	case 28:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:123
		{
			parserVAL.assignrhs = parserDollar[1].expr
		}
	case 29:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:124
		{
			parserVAL.assignrhs = parserDollar[1].arrayliter
		}
	case 30:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:125
		{
			parserVAL.assignrhs = parserDollar[1].pairelem
		}
	case 31:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:126
		{
			parserVAL.assignrhs = NewPair{fstExpr: parserDollar[3].expr, sndExpr: parserDollar[5].expr}
		}
	case 32:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:127
		{
			parserVAL.assignrhs = Call{ident: parserDollar[2].str, exprlist: parserDollar[4].exprs}
		}
	case 33:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:129
		{
			parserVAL.stmts = append(parserDollar[1].stmts, parserDollar[3].stmt)
		}
	case 34:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:130
		{
			parserVAL.stmts = []interface{}{parserDollar[1].stmt}
		}
	case 35:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:132
		{
			parserVAL.stmt = parserDollar[1].number
		}
	case 36:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:133
		{
			parserVAL.stmt = Declare{Type: parserDollar[1].typedefinition, lhs: parserDollar[2].str, rhs: parserDollar[4].assignrhs}
		}
	case 37:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:134
		{
			parserVAL.stmt = Assignment{lhs: parserDollar[1].assignlhs, rhs: parserDollar[3].assignrhs}
		}
	case 38:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:135
		{
			parserVAL.stmt = Read{parserDollar[2].assignlhs}
		}
	case 39:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:136
		{
			parserVAL.stmt = Free{parserDollar[2].expr}
		}
	case 40:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:137
		{
			parserVAL.stmt = Return{parserDollar[2].expr}
		}
	case 41:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:138
		{
			parserVAL.stmt = Exit{parserDollar[2].expr}
		}
	case 42:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:139
		{
			parserVAL.stmt = Print{parserDollar[2].expr}
		}
	case 43:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:140
		{
			parserVAL.stmt = Println{parserDollar[2].expr}
		}
	case 44:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:141
		{
			parserVAL.stmt = If{conditional: parserDollar[2].expr, thenStat: parserDollar[4].stmts, elseStat: parserDollar[6].stmts}
		}
	case 45:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:142
		{
			parserVAL.stmt = While{conditional: parserDollar[2].expr, doStat: parserDollar[4].stmts}
		}
	case 46:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:143
		{
			parserVAL.stmt = Scope{statlist: parserDollar[2].stmts, symbolTable: &SymbolTable{Table: make(map[string]Type)}}
		}
	case 47:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:145
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 48:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:146
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 49:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:147
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 50:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:148
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 51:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:149
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 52:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:150
		{
			parserVAL.expr = parserDollar[1].pairliter
		}
	case 53:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:151
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 54:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:152
		{
			parserVAL.expr = parserDollar[1].arrayelem
		}
	case 55:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:153
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 56:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:154
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 57:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:155
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 58:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:156
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 59:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:157
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 60:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:158
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 61:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:160
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 62:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:161
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 63:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:162
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 64:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:163
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 65:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:164
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 66:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:165
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 67:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:166
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 68:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:167
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 69:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:168
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 70:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:169
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 71:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:170
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 72:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:171
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 73:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:172
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 74:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:173
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 75:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:175
		{
			parserVAL.arrayliter = ArrayLiter{parserDollar[2].exprs}
		}
	case 76:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:177
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 77:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:178
		{
			parserVAL.exprs = []interface{}{parserDollar[1].expr}
		}
	case 78:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:179
		{
			parserVAL.exprs = []interface{}{}
		}
	case 79:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:181
		{
			parserVAL.arrayelem = ArrayElem{ident: parserDollar[1].str, exprs: parserDollar[2].bracketed}
		}
	case 80:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:183
		{
			parserVAL.bracketed = append(parserDollar[1].bracketed, parserDollar[3].expr)
		}
	case 81:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:184
		{
			parserVAL.bracketed = []interface{}{parserDollar[2].expr}
		}
	case 82:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:186
		{
			parserVAL.pairliter = NULL
		}
	case 83:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:188
		{
			parserVAL.pairelem = PairElem{fsnd: FST, expr: parserDollar[2].expr}
		}
	case 84:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:189
		{
			parserVAL.pairelem = PairElem{fsnd: SND, expr: parserDollar[2].expr}
		}
	case 85:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:191
		{
			parserVAL.typedefinition = PairType{fstType: parserDollar[3].pairelemtype, sndType: parserDollar[5].pairelemtype}
		}
	case 86:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:193
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 87:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:194
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 88:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:195
		{
			parserVAL.pairelemtype = Pair
		}
	case 89:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:197
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 90:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:198
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 91:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:199
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 92:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:201
		{
			parserVAL.typedefinition = Int
		}
	case 93:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:202
		{
			parserVAL.typedefinition = Bool
		}
	case 94:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:203
		{
			parserVAL.typedefinition = Char
		}
	case 95:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:204
		{
			parserVAL.typedefinition = String
		}
	case 96:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:206
		{
			parserVAL.typedefinition = ArrayType{Type: parserDollar[1].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
