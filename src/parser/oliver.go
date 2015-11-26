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

//line oliver.y:189

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 106,
	53, 73,
	-2, 70,
	-1, 107,
	53, 74,
	-2, 71,
}

const parserNprod = 81
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 392

var parserAct = [...]int{

	73, 4, 112, 117, 62, 24, 116, 72, 7, 163,
	158, 20, 38, 55, 56, 57, 58, 59, 60, 105,
	61, 35, 46, 30, 31, 35, 23, 19, 35, 162,
	141, 66, 67, 23, 102, 144, 147, 34, 145, 137,
	101, 23, 75, 161, 139, 145, 140, 70, 92, 93,
	94, 95, 96, 97, 98, 36, 23, 146, 32, 157,
	35, 71, 22, 33, 33, 104, 151, 143, 65, 69,
	109, 33, 70, 115, 113, 33, 75, 107, 114, 118,
	119, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 106, 33, 103, 76, 77, 30, 31,
	33, 133, 134, 64, 135, 47, 138, 48, 49, 50,
	9, 2, 33, 52, 51, 21, 142, 6, 44, 63,
	74, 37, 23, 23, 111, 40, 41, 54, 78, 53,
	5, 25, 26, 27, 28, 29, 43, 45, 39, 42,
	150, 3, 109, 152, 118, 113, 155, 154, 156, 107,
	1, 68, 0, 159, 160, 0, 0, 149, 0, 0,
	0, 23, 110, 0, 0, 106, 0, 0, 0, 23,
	0, 0, 18, 0, 23, 8, 10, 11, 12, 13,
	14, 15, 16, 0, 0, 0, 17, 0, 0, 0,
	0, 30, 31, 25, 26, 27, 28, 29, 0, 0,
	47, 0, 48, 49, 50, 0, 0, 0, 52, 51,
	25, 26, 27, 28, 29, 25, 26, 27, 28, 108,
	40, 41, 54, 0, 53, 0, 0, 0, 0, 0,
	22, 43, 45, 39, 42, 81, 83, 82, 79, 80,
	90, 91, 85, 87, 84, 86, 88, 89, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 153, 81,
	83, 82, 79, 80, 90, 91, 85, 87, 84, 86,
	88, 89, 0, 0, 0, 0, 0, 0, 0, 0,
	164, 81, 83, 82, 79, 80, 90, 91, 85, 87,
	84, 86, 88, 89, 0, 0, 0, 0, 0, 0,
	0, 0, 132, 81, 83, 82, 79, 80, 90, 91,
	85, 87, 84, 86, 88, 89, 0, 0, 0, 0,
	0, 0, 0, 148, 81, 83, 82, 79, 80, 90,
	91, 85, 87, 84, 86, 88, 89, 100, 0, 0,
	0, 0, 0, 0, 136, 0, 99, 0, 0, 0,
	0, 0, 0, 81, 83, 82, 79, 80, 90, 91,
	85, 87, 84, 86, 88, 89, 81, 83, 82, 79,
	80, 90, 91, 85, 87, 84, 86, 88, 89, 81,
	83, 82, 79, 80, 90, 91, 85, 87, 84, 86,
	88, 89,
}
var parserPact = [...]int{

	107, -1000, -1000, 168, 53, -1000, -1000, -25, -1000, -2,
	0, 170, 170, 170, 170, 170, 170, 170, 168, -1000,
	-1000, -1000, 50, -1000, -1000, -1000, -1000, -1000, -1000, 14,
	170, 170, -1000, 168, 15, 6, 75, -1000, 344, -1000,
	-1000, -1000, -1000, -1000, -1000, 50, -1000, 170, 170, 170,
	170, 170, 170, 170, -1000, 344, 344, 344, 344, 331,
	318, 35, -28, 42, 170, 190, 344, 344, -1000, 106,
	75, -1000, -1000, 344, -1000, -1000, 19, -56, 170, 170,
	170, 170, 170, 170, 170, 170, 170, 170, 170, 170,
	170, 170, -1000, -1000, -1000, -1000, -1000, -1000, 246, 168,
	168, -1000, -10, 170, 289, -19, -1000, -1000, 14, 7,
	100, -12, -1000, -32, -1000, 170, 13, -20, 344, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 41, 16, 268, -1000, 190, 168, 60,
	185, -1000, 200, 170, -1000, 170, 168, -1000, -1000, 3,
	5, 168, -1000, 170, -13, 344, 12, -1000, -1000, 4,
	224, -1000, -1000, -1000, -1000,
}
var parserPgo = [...]int{

	0, 150, 141, 130, 117, 1, 7, 110, 0, 3,
	124, 2, 5, 120, 22, 119, 118, 27, 4, 11,
	115, 19,
}
var parserR1 = [...]int{

	0, 1, 2, 2, 3, 3, 10, 10, 11, 7,
	7, 7, 6, 6, 6, 6, 6, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 13,
	9, 9, 9, 14, 15, 15, 16, 12, 12, 20,
	21, 21, 21, 18, 18, 18, 17, 17, 17, 17,
	19,
}
var parserR2 = [...]int{

	0, 4, 2, 0, 7, 8, 3, 1, 2, 1,
	1, 1, 1, 1, 1, 6, 5, 3, 1, 1,
	4, 3, 2, 2, 2, 2, 2, 2, 7, 5,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 0, 2, 4, 3, 1, 2, 2, 6,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3,
}
var parserChk = [...]int{

	-1000, -1, 4, -2, -5, -3, -4, -18, 7, -7,
	8, 9, 10, 11, 12, 13, 14, 18, 4, -17,
	-19, -20, 62, -14, -12, 25, 26, 27, 28, 29,
	23, 24, 5, 59, 62, 53, 57, -7, -8, 63,
	50, 51, 64, 61, -16, 62, -14, 30, 32, 33,
	34, 39, 38, 54, 52, -8, -8, -8, -8, -8,
	-8, -5, -18, -15, 53, 54, -8, -8, -4, 54,
	57, 55, -6, -8, -13, -12, 21, 22, 53, 38,
	39, 35, 37, 36, 44, 42, 45, 43, 46, 47,
	40, 41, -8, -8, -8, -8, -8, -8, -8, 15,
	19, 5, 62, 53, -8, -21, -17, -19, 29, -18,
	56, -10, -11, -18, -6, 54, 62, -9, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, 56, -5, -5, -8, 55, 58, 6, 56,
	58, 62, -8, 54, 55, 58, 16, 20, 55, -21,
	-5, 6, -11, 58, -9, -8, -5, 56, 5, -5,
	-8, 56, 17, 5, 56,
}
var parserDef = [...]int{

	0, -2, 3, 0, 0, 2, 18, 0, 19, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 73,
	74, 75, 9, 10, 11, 76, 77, 78, 79, 0,
	0, 0, 1, 0, 0, 0, 0, 22, 23, 31,
	32, 33, 34, 35, 36, 37, 38, 0, 0, 0,
	0, 0, 0, 0, 66, 24, 25, 26, 27, 0,
	0, 0, 0, 63, 0, 0, 67, 68, 17, 0,
	0, 80, 21, 12, 13, 14, 0, 0, 62, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 40, 41, 42, 43, 44, 0, 0,
	0, 30, 0, 0, 0, 0, -2, -2, 72, 0,
	0, 0, 7, 0, 20, 0, 0, 0, 61, 45,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 0, 0, 0, 65, 0, 0, 0,
	0, 8, 0, 62, 59, 0, 0, 29, 64, 0,
	0, 0, 6, 0, 0, 60, 0, 69, 4, 0,
	0, 16, 28, 5, 15,
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
		//line oliver.y:79
		{
			parserlex.(*Lexer).prog = &Program{functionlist: parserDollar[2].functions, statList: parserDollar[3].stmts, symbolTable: &SymbolTable{Table: make(map[string]Type)}}
		}
	case 2:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:83
		{
			parserVAL.functions = append(parserDollar[1].functions, parserDollar[2].function)
		}
	case 3:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:84
		{
			parserVAL.functions = []*Function{}
		}
	case 4:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:87
		{
			parserVAL.function = &Function{ident: parserDollar[2].str, returnType: parserDollar[1].typedefinition, statlist: parserDollar[6].stmts}
		}
	case 5:
		parserDollar = parserS[parserpt-8 : parserpt+1]
		//line oliver.y:90
		{
			parserVAL.function = &Function{ident: parserDollar[2].str, returnType: parserDollar[1].typedefinition, statlist: parserDollar[7].stmts, parameterTypes: parserDollar[4].params}
		}
	case 6:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:94
		{
			parserVAL.params = append(parserDollar[1].params, parserDollar[3].param)
		}
	case 7:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:95
		{
			parserVAL.params = []Param{parserDollar[1].param}
		}
	case 8:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:97
		{
			parserVAL.param = Param{paramType: parserDollar[1].typedefinition, ident: parserDollar[2].str}
		}
	case 9:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:99
		{
			parserVAL.assignlhs = parserDollar[1].str
		}
	case 10:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:100
		{
			parserVAL.assignlhs = parserDollar[1].arrayelem
		}
	case 11:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:101
		{
			parserVAL.assignlhs = parserDollar[1].pairelem
		}
	case 12:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:103
		{
			parserVAL.assignrhs = parserDollar[1].expr
		}
	case 13:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:104
		{
			parserVAL.assignrhs = parserDollar[1].arrayliter
		}
	case 14:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:105
		{
			parserVAL.assignrhs = parserDollar[1].pairelem
		}
	case 15:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:106
		{
			parserVAL.assignrhs = NewPair{fstExpr: parserDollar[3].expr, sndExpr: parserDollar[5].expr}
		}
	case 16:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:107
		{
			parserVAL.assignrhs = Call{ident: parserDollar[2].str, exprlist: parserDollar[4].exprs}
		}
	case 17:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:109
		{
			parserVAL.stmts = append(parserDollar[1].stmts, parserDollar[3].stmt)
		}
	case 18:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:110
		{
			parserVAL.stmts = []interface{}{parserDollar[1].stmt}
		}
	case 19:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:113
		{
			parserVAL.stmt = parserDollar[1].number
		}
	case 20:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:114
		{
			parserVAL.stmt = Declare{Type: parserDollar[1].typedefinition, lhs: parserDollar[2].str, rhs: parserDollar[4].assignrhs}
		}
	case 21:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:115
		{
			parserVAL.stmt = Assignment{lhs: parserDollar[1].assignlhs, rhs: parserDollar[3].assignrhs}
		}
	case 22:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:116
		{
			parserVAL.stmt = Read{parserDollar[2].assignlhs}
		}
	case 23:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:117
		{
			parserVAL.stmt = Free{parserDollar[2].expr}
		}
	case 24:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:118
		{
			parserVAL.stmt = Return{parserDollar[2].expr}
		}
	case 25:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:119
		{
			parserVAL.stmt = Exit{parserDollar[2].expr}
		}
	case 26:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:120
		{
			parserVAL.stmt = Print{parserDollar[2].expr}
		}
	case 27:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:121
		{
			parserVAL.stmt = Println{parserDollar[2].expr}
		}
	case 28:
		parserDollar = parserS[parserpt-7 : parserpt+1]
		//line oliver.y:122
		{
			parserVAL.stmt = If{conditional: parserDollar[2].expr, thenStat: parserDollar[4].stmts, elseStat: parserDollar[6].stmts}
		}
	case 29:
		parserDollar = parserS[parserpt-5 : parserpt+1]
		//line oliver.y:123
		{
			parserVAL.stmt = While{conditional: parserDollar[2].expr, doStat: parserDollar[4].stmts}
		}
	case 30:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:124
		{
			parserVAL.stmt = Scope{statlist: parserDollar[2].stmts, symbolTable: &SymbolTable{Table: make(map[string]Type)}}
		}
	case 31:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:126
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 32:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:127
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 33:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:128
		{
			parserVAL.expr = parserDollar[1].number
		}
	case 34:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:129
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 35:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:130
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 36:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:131
		{
			parserVAL.expr = parserDollar[1].pairliter
		}
	case 37:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:132
		{
			parserVAL.expr = parserDollar[1].str
		}
	case 38:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:133
		{
			parserVAL.expr = parserDollar[1].arrayelem
		}
	case 39:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:134
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 40:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:135
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 41:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:136
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 42:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:137
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 43:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:138
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 44:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:139
		{
			parserVAL.expr = Unop{unary: parserDollar[1].number, expr: parserDollar[2].expr}
		}
	case 45:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:141
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 46:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:142
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 47:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:143
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 48:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:144
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 49:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:145
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 50:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:146
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 51:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:147
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 52:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:148
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 53:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:149
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 54:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:150
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 55:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:151
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 56:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:152
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 57:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:153
		{
			parserVAL.expr = Binop{left: parserDollar[1].expr, binary: parserDollar[2].number, right: parserDollar[3].expr}
		}
	case 58:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:154
		{
			parserVAL.expr = parserDollar[2].expr
		}
	case 59:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:156
		{
			parserVAL.arrayliter = ArrayLiter{parserDollar[2].exprs}
		}
	case 60:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:158
		{
			parserVAL.exprs = append(parserDollar[1].exprs, parserDollar[3].expr)
		}
	case 61:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:159
		{
			parserVAL.exprs = []interface{}{parserDollar[1].expr}
		}
	case 62:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line oliver.y:160
		{
			parserVAL.exprs = []interface{}{}
		}
	case 63:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:162
		{
			parserVAL.arrayelem = ArrayElem{ident: parserDollar[1].str, exprs: parserDollar[2].bracketed}
		}
	case 64:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line oliver.y:164
		{
			parserVAL.bracketed = append(parserDollar[1].bracketed, parserDollar[3].expr)
		}
	case 65:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:165
		{
			parserVAL.bracketed = []interface{}{parserDollar[2].expr}
		}
	case 66:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:167
		{
			parserVAL.pairliter = NULL
		}
	case 67:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:169
		{
			parserVAL.pairelem = PairElem{fsnd: FST, expr: parserDollar[2].expr}
		}
	case 68:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line oliver.y:170
		{
			parserVAL.pairelem = PairElem{fsnd: SND, expr: parserDollar[2].expr}
		}
	case 69:
		parserDollar = parserS[parserpt-6 : parserpt+1]
		//line oliver.y:172
		{
			parserVAL.typedefinition = PairType{fstType: parserDollar[3].pairelemtype, sndType: parserDollar[5].pairelemtype}
		}
	case 70:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:174
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 71:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:175
		{
			parserVAL.pairelemtype = parserDollar[1].typedefinition
		}
	case 72:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:176
		{
			parserVAL.pairelemtype = Pair
		}
	case 73:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:178
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 74:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:179
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 75:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:180
		{
			parserVAL.typedefinition = parserDollar[1].typedefinition
		}
	case 76:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:182
		{
			parserVAL.typedefinition = Int
		}
	case 77:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:183
		{
			parserVAL.typedefinition = Bool
		}
	case 78:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:184
		{
			parserVAL.typedefinition = Char
		}
	case 79:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line oliver.y:185
		{
			parserVAL.typedefinition = String
		}
	case 80:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line oliver.y:187
		{
			parserVAL.typedefinition = ArrayType{Type: parserDollar[1].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
