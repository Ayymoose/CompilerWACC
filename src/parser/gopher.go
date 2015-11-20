//line new.y:2
package parser

import __yyfmt__ "fmt"

//line new.y:2
//line new.y:5
type parserSymType struct {
	yys        int
	str        string
	functions  []*Function
	function   *Function
	statements []*Statement
	statement  *Statement
	params     []*Parameter
	param      *Parameter
	exprs      []*Expr
	expr       *Expr
	chars      []*Char
	char       *Char
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
const NEW_PAIR = 57363
const CALL = 57364
const FST = 57365
const SND = 57366
const INT = 57367
const BOOL = 57368
const CHAR = 57369
const STRING = 57370
const PAIR = 57371
const IDENTIFIER = 57372
const NOT = 57373
const NEG = 57374
const LEN = 57375
const ORD = 57376
const CHR = 57377
const MULT = 57378
const DIV = 57379
const MOD = 57380
const ADD = 57381
const SUB = 57382
const AND = 57383
const OR = 57384
const GT = 57385
const GTE = 57386
const LT = 57387
const LTE = 57388
const EQ = 57389
const NEQ = 57390
const INTEGER = 57391
const POSITIVE = 57392
const NEGATIVE = 57393
const TRUE = 57394
const FALSE = 57395
const CHARACTER = 57396
const NULL = 57397

var parserToknames = []string{
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
	"NEW_PAIR",
	"CALL",
	"FST",
	"SND",
	"INT",
	"BOOL",
	"CHAR",
	"STRING",
	"PAIR",
	"IDENTIFIER",
	"NOT",
	"NEG",
	"LEN",
	"ORD",
	"CHR",
	"MULT",
	"DIV",
	"MOD",
	"ADD",
	"SUB",
	"AND",
	"OR",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"EQ",
	"NEQ",
	"INTEGER",
	"POSITIVE",
	"NEGATIVE",
	"TRUE",
	"FALSE",
	"CHARACTER",
	"NULL",
	" >=",
	" <=",
	" >",
	" <",
	" !=",
	" ==",
	" *",
	" /",
	" %",
	" +",
	" -",
	" &&",
	" ||",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 78,
	74, 37,
	-2, 46,
	-1, 79,
	74, 38,
	-2, 47,
	-1, 125,
	57, 0,
	59, 0,
	60, 0,
	61, 0,
	-2, 57,
	-1, 126,
	57, 0,
	59, 0,
	60, 0,
	61, 0,
	-2, 58,
	-1, 127,
	56, 0,
	58, 0,
	-2, 59,
	-1, 128,
	57, 0,
	59, 0,
	60, 0,
	61, 0,
	-2, 60,
	-1, 129,
	56, 0,
	58, 0,
	-2, 61,
	-1, 130,
	57, 0,
	59, 0,
	60, 0,
	61, 0,
	-2, 62,
	-1, 145,
	70, 32,
	-2, 84,
}

const parserNprod = 94
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 410

var parserAct = []int{

	85, 146, 14, 117, 71, 15, 7, 6, 5, 77,
	5, 47, 134, 133, 158, 26, 84, 33, 150, 41,
	64, 65, 66, 67, 68, 69, 48, 38, 70, 26,
	27, 73, 74, 29, 28, 76, 149, 26, 110, 166,
	82, 81, 79, 78, 27, 34, 165, 72, 148, 26,
	105, 106, 27, 114, 36, 37, 87, 89, 30, 31,
	83, 39, 141, 120, 27, 29, 57, 58, 59, 60,
	61, 163, 153, 115, 88, 140, 145, 34, 122, 35,
	75, 119, 111, 29, 109, 62, 63, 52, 53, 37,
	56, 124, 125, 126, 127, 128, 129, 130, 131, 107,
	121, 123, 37, 151, 50, 37, 2, 37, 3, 90,
	1, 54, 55, 32, 51, 136, 137, 97, 88, 49,
	138, 37, 37, 144, 26, 26, 81, 79, 78, 46,
	143, 142, 9, 10, 11, 12, 13, 108, 16, 27,
	27, 45, 30, 31, 44, 152, 157, 119, 159, 29,
	135, 160, 43, 42, 161, 162, 40, 156, 155, 164,
	26, 86, 118, 26, 29, 57, 58, 59, 60, 61,
	9, 10, 11, 12, 80, 27, 116, 8, 27, 4,
	0, 0, 0, 0, 62, 63, 52, 53, 0, 56,
	0, 0, 98, 99, 100, 101, 102, 103, 104, 0,
	0, 0, 0, 50, 0, 0, 0, 0, 0, 0,
	54, 55, 93, 94, 95, 96, 91, 92, 98, 99,
	100, 101, 102, 103, 104, 0, 0, 0, 0, 0,
	72, 139, 0, 0, 0, 0, 0, 0, 93, 94,
	95, 96, 91, 92, 98, 99, 100, 101, 102, 103,
	104, 0, 0, 147, 0, 0, 0, 98, 99, 100,
	101, 102, 103, 104, 93, 94, 95, 96, 91, 92,
	98, 99, 100, 101, 102, 103, 104, 0, 94, 154,
	96, 91, 92, 98, 99, 100, 101, 102, 103, 104,
	93, 94, 95, 96, 91, 92, 98, 99, 100, 101,
	102, 103, 104, 0, 167, 113, 0, 0, 0, 0,
	112, 0, 0, 0, 0, 0, 93, 94, 95, 96,
	91, 92, 98, 99, 100, 101, 102, 103, 104, 0,
	132, 98, 99, 100, 101, 102, 103, 104, 0, 0,
	0, 0, 93, 94, 95, 96, 91, 92, 0, 0,
	0, 93, 94, 95, 96, 91, 92, 98, 99, 100,
	101, 102, 103, 104, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 93, 94, 95,
	96, 91, 92, 25, 0, 0, 0, 17, 18, 19,
	20, 21, 22, 23, 0, 0, 0, 24, 0, 0,
	0, 0, 30, 31, 9, 10, 11, 12, 13, 29,
}
var parserPact = []int{

	102, -1000, 107, 379, 107, 3, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 10, 49, 3, -11, 119, 134, 134,
	134, 134, 134, 134, 134, 379, -27, -1000, -1000, -1000,
	134, 134, -1000, 11, -40, 145, -1000, 379, -12, 35,
	-1000, 321, -1000, -1000, -1000, -1000, -1000, -27, -1000, 134,
	134, 50, -1000, -1000, 30, 30, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 321, 321, 321, 321, 295, 286,
	48, -1000, 134, 321, 321, 107, -1000, -8, -1000, -1000,
	10, -29, -18, 35, -1000, 321, -1000, 9, -1000, 53,
	134, 134, 134, 134, 134, 134, 134, 134, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 321, 260, -1000, -63, -1000,
	-65, 30, 379, 379, -1000, 156, 5, -1000, -9, 3,
	145, -1000, 134, 7, 182, 247, 247, 221, 247, 221,
	247, 321, -1000, -1000, -1000, -1000, 32, 16, -57, -1000,
	97, 107, -1000, 2, 208, 134, -61, 134, 379, -1000,
	-1000, 379, -1000, -1000, 134, 1, -1000, 182, -1000, 321,
	29, 34, 234, -1000, -1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 108, 179, 2, 5, 7, 6, 177, 9, 11,
	176, 3, 162, 16, 138, 0, 1, 4, 26, 34,
	161, 158, 157, 153, 152, 144, 141, 129, 119, 117,
	114, 82, 38, 110,
}
var parserR1 = []int{

	0, 33, 1, 1, 1, 2, 10, 10, 11, 11,
	12, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 14, 14, 14, 13, 13, 13, 13,
	13, 21, 21, 22, 16, 19, 19, 4, 4, 4,
	5, 5, 5, 5, 6, 7, 8, 8, 8, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 28, 28, 28, 28, 28,
	29, 29, 29, 29, 29, 29, 29, 9, 18, 17,
	17, 23, 30, 30, 30, 24, 24, 25, 26, 32,
	32, 31, 20, 27,
}
var parserR2 = []int{

	0, 4, 2, 1, 0, 8, 1, 0, 3, 1,
	2, 4, 3, 2, 2, 2, 2, 2, 2, 7,
	5, 3, 3, 1, 1, 1, 1, 1, 6, 1,
	5, 1, 0, 2, 2, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 3, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	3, 2, 1, 1, 0, 1, 1, 3, 3, 2,
	1, 1, 4, 1,
}
var parserChk = []int{

	-1000, -33, 4, -1, -2, -4, -5, -6, -7, 25,
	26, 27, 28, 29, -3, -4, -14, 8, 9, 10,
	11, 12, 13, 14, 18, 4, -9, -18, -19, 30,
	23, 24, -1, -9, 74, 69, 5, 73, -9, 72,
	-14, -15, -23, -24, -25, -26, -27, -9, -18, -28,
	69, -30, 52, 53, 76, 77, 55, 31, 32, 33,
	34, 35, 50, 51, -15, -15, -15, -15, -15, -15,
	-3, -17, 74, -15, -15, 69, 75, -8, -5, -6,
	29, -4, -3, 72, -13, -15, -20, 21, -19, 22,
	74, 60, 61, 56, 57, 58, 59, -29, 36, 37,
	38, 39, 40, 41, 42, -15, -15, 49, -31, 54,
	-32, -31, 15, 19, 5, -15, -10, -11, -12, -4,
	71, -13, 69, -9, -15, -15, -15, -15, -15, -15,
	-15, -15, 70, 76, 77, -32, -3, -3, -17, 75,
	70, 71, -9, -8, -15, 69, -16, 71, 16, 20,
	75, 6, -11, 70, 71, -21, -22, -15, 75, -15,
	-3, -3, -15, 70, -16, 17, 5, 70,
}
var parserDef = []int{

	0, -2, 4, 0, 3, 0, 37, 38, 39, 40,
	41, 42, 43, 0, 0, 0, 0, 0, 84, 84,
	84, 84, 84, 84, 84, 0, 23, 24, 25, 77,
	84, 84, 2, 0, 0, 0, 1, 0, 0, 84,
	13, 14, 49, 50, 51, 52, 53, 54, 55, 84,
	84, 0, 85, 86, 0, 0, 93, 65, 66, 67,
	68, 69, 82, 83, 15, 16, 17, 18, 0, 0,
	0, 78, 84, 35, 36, 7, 44, 0, -2, -2,
	48, 0, 22, 84, 12, 26, 27, 0, 29, 0,
	84, 84, 84, 84, 84, 84, 84, 84, 70, 71,
	72, 73, 74, 75, 76, 56, 0, 81, 0, 91,
	0, 90, 0, 0, 21, 0, 0, 6, 9, 0,
	0, 11, 84, 0, 0, -2, -2, -2, -2, -2,
	-2, 63, 64, 87, 88, 89, 0, 0, 0, 80,
	0, 0, 10, 0, 0, -2, 0, 84, 0, 20,
	79, 0, 8, 45, 84, 0, 31, 0, 92, 34,
	0, 0, 0, 30, 33, 19, 5, 28,
}
var parserTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 60, 77, 3, 3, 64, 67, 3,
	69, 70, 62, 65, 71, 66, 3, 63, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 73,
	59, 72, 58, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 74, 3, 75, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 68,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55,
}
var parserTok3 = []int{
	8217, 76, 0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var parserDebug = 0

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

const parserFlag = -1000

func parserTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(parserToknames) {
		if parserToknames[c-4] != "" {
			return parserToknames[c-4]
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

func parserlex1(lex parserLexer, lval *parserSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		c = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			c = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		c = parserTok3[i+0]
		if c == char {
			c = parserTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(c), uint(char))
	}
	return c
}

func parserParse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parserchar), parserStatname(parserstate))
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
		parserchar = parserlex1(parserlex, &parserlval)
	}
	parsern += parserchar
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parserchar { /* valid shift */
		parserchar = -1
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
			parserchar = parserlex1(parserlex, &parserlval)
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
			if parsern < 0 || parsern == parserchar {
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
			parserlex.Error("syntax error")
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parserchar))
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
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parserchar))
			}
			if parserchar == parserEofCode {
				goto ret1
			}
			parserchar = -1
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
		//line new.y:97
		{
			yylex.(*lexer).prog = &Program{Functions: parserS[parserpt-2].functions, Statements: parserS[parserpt-1].statement}
			return 0
		}
	case 2:
		//line new.y:102
		{
			parserVAL.functions = append([]Function{parserS[parserpt-1].function}, parserS[parserpt-0].functions...)
		}
	case 3:
		//line new.y:103
		{
			parserVAL.functions = []Function{parserS[parserpt-0].function}
		}
	case 4:
		//line new.y:104
		{
			parserVAL.functions
		}
	case 5:
		//line new.y:108
		{
			parserVAL.function = &Function{Type: parserS[parserpt-7].typeDef, Ident: parserS[parserpt-6].str, Params: parserS[parserpt-4].params, Stat: parserS[parserpt-1].statement}
		}
	case 6:
		//line new.y:111
		{
			parserVAL.params = parserS[parserpt-0].params
		}
	case 7:
		//line new.y:112
		{
			parserVAL.params
		}
	case 8:
		//line new.y:115
		{
			parserVAL.params = append([]Parameter{parserS[parserpt-2].param}, parserS[parserpt-0].params...)
		}
	case 9:
		//line new.y:116
		{
			parserVAL.params = []Parameter{parserS[parserpt-0].param}
		}
	case 10:
		//line new.y:119
		{
			parserVAL.param = Parameter{Type: parserS[parserpt-1].typeDef, Ident: parserS[parserpt-0].str}
		}
	case 11:
		//line new.y:123
		{
			parserVAL.statement = Statement{Type: parserS[parserpt-3].typeDef, Ident: parserS[parserpt-2].str, AssignRHS: parserS[parserpt-0].assign_rhs}
		}
	case 12:
		//line new.y:124
		{
			parserVAL.statement = Statement{AssignLHS: parserS[parserpt-2].assign_lhs, AssignRHS: parserS[parserpt-0].assign_rhs}
		}
	case 13:
		//line new.y:125
		{
			parserVAL.statement = Statement{AssignLHS: parserS[parserpt-0].assign_lhs}
		}
	case 14:
		//line new.y:126
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 15:
		//line new.y:127
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 16:
		//line new.y:128
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 17:
		//line new.y:129
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 18:
		//line new.y:130
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 19:
		//line new.y:131
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-5].expr, StatL: parserS[parserpt-3].statement, StatR: parserS[parserpt-1].statement}
		}
	case 20:
		//line new.y:132
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-3].expr, Stat: parserS[parserpt-1].statement}
		}
	case 21:
		//line new.y:133
		{
			parserVAL.statement = Statement{Stat: parserS[parserpt-1].statement}
		}
	case 22:
		//line new.y:134
		{
			parserVAL.statement = append([]Statement{parserS[parserpt-2].statement}, parserS[parserpt-0].statement...)
		}
	case 23:
		//line new.y:137
		{
			parserVAL.assign_lhs = AssignLHS{Ident: parserS[parserpt-0].str}
		}
	case 24:
		//line new.y:138
		{
			parserVAL.assign_lhs = AssignLHS{ArrayElem: parserS[parserpt-0].array_elem}
		}
	case 25:
		//line new.y:139
		{
			parserVAL.assign_lhs = AssignLHS{PairElem: parserS[parserpt-0].pair_elem}
		}
	case 26:
		//line new.y:142
		{
			parserVAL.assign_rhs = AssignRHS{Expr: parserS[parserpt-0].expr}
		}
	case 27:
		//line new.y:143
		{
			parserVAL.assign_rhs = AssignRHS{ArrayLiter: parserS[parserpt-0].array_liter}
		}
	case 28:
		//line new.y:144
		{
			parserVAL.assign_rhs = AssignRHS{ExprL: parserS[parserpt-3].expr, ExprR: parserS[parserpt-1].expr}
		}
	case 29:
		//line new.y:145
		{
			parserVAL.assign_rhs = AssignRHS{PairElem: parserS[parserpt-0].pair_elem}
		}
	case 30:
		//line new.y:146
		{
			parserVAL.assign_rhs = AssignRHS{Ident: parserS[parserpt-3].str, Args: parserS[parserpt-1].arg}
		}
	case 31:
		//line new.y:149
		{
			parserVAL.arg = parserS[parserpt-0].arg
		}
	case 32:
		//line new.y:150
		{
			parserVAL.arg
		}
	case 33:
		//line new.y:153
		{
			parserVAL.arg = append([]Expr{parserS[parserpt-1].expr}, parserS[parserpt-0].expr...)
		}
	case 34:
		//line new.y:156
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].expr}
		}
	case 35:
		//line new.y:159
		{
			parserVAL.pair_elem = PairElem{Expr: parserS[parserpt-0].expr}
		}
	case 36:
		//line new.y:160
		{
			parserVAL.pair_elem = PairElem{Expr: parserS[parserpt-0].expr}
		}
	case 37:
		//line new.y:163
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 38:
		//line new.y:164
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 39:
		//line new.y:165
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 40:
		//line new.y:168
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 41:
		//line new.y:169
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 42:
		//line new.y:170
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 43:
		//line new.y:171
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 44:
		//line new.y:174
		{
			parserVAL.typeDef = ArrayType{Type: parserS[parserpt-2].typeDef}
		}
	case 45:
		//line new.y:177
		{
			parserVAL.typeDef = PairType{Lpair: parserS[parserpt-3].typeDef, Rpair: parserS[parserpt-1].typeDef}
		}
	case 46:
		//line new.y:180
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].typeDef}
		}
	case 47:
		//line new.y:181
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].typeDef}
		}
	case 48:
		//line new.y:182
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].str}
		}
	case 49:
		//line new.y:185
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].int}
		}
	case 50:
		//line new.y:186
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].bool}
		}
	case 51:
		//line new.y:187
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].rune}
		}
	case 52:
		//line new.y:188
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].str}
		}
	case 53:
		//line new.y:189
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].pair}
		}
	case 54:
		//line new.y:190
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].str}
		}
	case 55:
		//line new.y:191
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].array_elem}
		}
	case 56:
		//line new.y:202
		{
			parserVAL.expr = Expr{UnaryOp: parserS[parserpt-1].uop, Expr: parserS[parserpt-0].expr}
		}
	case 57:
		//line new.y:203
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].not, Rexpr: parserS[parserpt-0].expr}
		}
	case 58:
		//line new.y:204
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].eq, Rexpr: parserS[parserpt-0].expr}
		}
	case 59:
		//line new.y:205
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].gte, Rexpr: parserS[parserpt-0].expr}
		}
	case 60:
		//line new.y:206
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].lte, Rexpr: parserS[parserpt-0].expr}
		}
	case 61:
		//line new.y:207
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].gt, Rexpr: parserS[parserpt-0].expr}
		}
	case 62:
		//line new.y:208
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].lt, Rexpr: parserS[parserpt-0].expr}
		}
	case 63:
		//line new.y:210
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].bop, Rexpr: parserS[parserpt-0].expr}
		}
	case 64:
		//line new.y:211
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-1].expr}
		}
	case 65:
		//line new.y:214
		{
			parserVAL.uop = UnaryOp{Op: parserS[parserpt-0].str}
		}
	case 66:
		//line new.y:215
		{
			parserVAL.uop = UnaryOp{Op: parserS[parserpt-0].str}
		}
	case 67:
		//line new.y:216
		{
			parserVAL.uop = UnaryOp{Op: parserS[parserpt-0].str}
		}
	case 68:
		//line new.y:217
		{
			parserVAL.uop = UnaryOp{Op: parserS[parserpt-0].str}
		}
	case 69:
		//line new.y:218
		{
			parserVAL.uop = UnaryOp{Op: parserS[parserpt-0].str}
		}
	case 70:
		//line new.y:221
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 71:
		//line new.y:222
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 72:
		//line new.y:223
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 73:
		//line new.y:224
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 74:
		//line new.y:225
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 75:
		//line new.y:226
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 76:
		//line new.y:227
		{
			parserVAL.bop = BinaryOp{Op: parserS[parserpt-0].str}
		}
	case 77:
		parserVAL.str = parserS[parserpt-0].str
	case 78:
		//line new.y:232
		{
			parserVAL.array_elem = ArrayElem{Ident: parserS[parserpt-1].str, Exprs: parserS[parserpt-0].expr}
		}
	case 79:
		//line new.y:236
		{
			parserVAL.expr = append([]Expr{parserS[parserpt-2].expr}, parserS[parserpt-1].expr...)
		}
	case 80:
		//line new.y:237
		{
			parserVAL.expr = []Expr{parserS[parserpt-1].expr}
		}
	case 81:
		//line new.y:240
		{
			parserVAL.int = IntLiter{Sign: parserS[parserpt-1].sign, Int: parserS[parserpt-0].str}
		}
	case 82:
		//line new.y:243
		{
			parserVAL.sign = Sign{Sign: parserS[parserpt-0].str}
		}
	case 83:
		//line new.y:244
		{
			parserVAL.sign = Sign{Sign: parserS[parserpt-0].str}
		}
	case 84:
		//line new.y:245
		{
			parserVAL.sign
		}
	case 85:
		//line new.y:248
		{
			parserVAL.bool = Bool{Bool: parserS[parserpt-0].str}
		}
	case 86:
		//line new.y:249
		{
			parserVAL.bool = Bool{Bool: parserS[parserpt-0].str}
		}
	case 87:
		//line new.y:252
		{
			parserVAL.rune = Char{Char: parserS[parserpt-1].chr}
		}
	case 88:
		//line new.y:255
		{
			parserVAL.str = Str{Chars: parserS[parserpt-1].chr}
		}
	case 89:
		//line new.y:258
		{
			parserVAL.chr = append([]Char{parserS[parserpt-1].chr}, parserS[parserpt-0].chr...)
		}
	case 90:
		//line new.y:259
		{
			parserVAL.chr = []Char{parserS[parserpt-0].chr}
		}
	case 91:
		//line new.y:262
		{
			parserVAL.chr = Char{Char: parserS[parserpt-0].str}
		}
	case 92:
		//line new.y:264
		{
			parserVAL.array_liter = append([]Expr{parserS[parserpt-2].expr}, parserS[parserpt-1].expr...)
		}
	case 93:
		//line new.y:267
		{
			parserVAL.pair = PairLiter{PairLit: parserS[parserpt-0].str}
		}
	}
	goto parserstack /* stack new state and value */
}
