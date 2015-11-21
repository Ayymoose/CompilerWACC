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
const EqEq = 57398
const NEq = 57399
const AndAnd = 57400
const OrOr = 57401

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
	" -",
	" +",
	" !",
	" /",
	" %",
	" <",
	" >",
	"EqEq",
	"NEq",
	"AndAnd",
	"OrOr",
	" ;",
	" =",
	" [",
	" ]",
	" (",
	" )",
	" ,",
	" *",
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
	-1, 6,
	69, 45,
	-2, 37,
	-1, 8,
	69, 46,
	-2, 39,
	-1, 43,
	69, 45,
	-2, 48,
	-1, 62,
	72, 75,
	-2, 82,
	-1, 142,
	72, 75,
	-2, 82,
	-1, 153,
	72, 32,
	-2, 82,
}

const parserNprod = 93
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 468

var parserAct = []int{

	89, 154, 110, 78, 115, 37, 24, 83, 144, 5,
	143, 5, 122, 86, 23, 171, 151, 141, 121, 58,
	88, 153, 8, 7, 42, 126, 21, 40, 51, 71,
	72, 73, 74, 75, 76, 6, 36, 174, 165, 80,
	81, 41, 57, 79, 46, 44, 36, 85, 18, 77,
	159, 20, 87, 169, 36, 92, 119, 43, 49, 35,
	107, 108, 109, 111, 173, 158, 102, 48, 101, 35,
	16, 114, 157, 116, 97, 98, 112, 35, 96, 95,
	120, 97, 98, 99, 100, 103, 104, 105, 106, 19,
	149, 22, 2, 92, 175, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 125, 46,
	44, 124, 16, 1, 16, 16, 96, 95, 16, 97,
	98, 145, 43, 16, 38, 39, 25, 152, 123, 85,
	150, 19, 146, 147, 3, 63, 127, 36, 36, 17,
	113, 56, 47, 111, 55, 156, 15, 54, 53, 91,
	93, 38, 39, 50, 164, 52, 166, 163, 19, 162,
	35, 35, 170, 168, 160, 90, 172, 84, 19, 36,
	82, 13, 167, 4, 0, 0, 0, 36, 69, 70,
	64, 65, 0, 68, 59, 60, 61, 0, 69, 70,
	64, 65, 35, 68, 59, 60, 61, 94, 0, 62,
	35, 0, 0, 66, 67, 102, 0, 101, 0, 62,
	0, 0, 0, 66, 67, 0, 0, 96, 95, 0,
	97, 98, 99, 100, 103, 104, 105, 106, 102, 0,
	101, 0, 0, 0, 155, 9, 10, 11, 12, 14,
	96, 95, 0, 97, 98, 99, 100, 103, 104, 105,
	106, 102, 0, 101, 0, 0, 0, 161, 9, 10,
	11, 12, 45, 96, 95, 0, 97, 98, 99, 100,
	103, 104, 105, 106, 102, 0, 101, 118, 0, 0,
	142, 0, 0, 0, 0, 0, 96, 95, 0, 97,
	98, 99, 100, 103, 104, 105, 106, 117, 0, 0,
	148, 0, 102, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 95, 0, 97, 98, 99,
	100, 103, 104, 105, 106, 0, 102, 0, 101, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 96, 95,
	0, 97, 98, 99, 100, 103, 104, 105, 106, 102,
	0, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 96, 95, 0, 97, 98, 99, 100, 103, 104,
	105, 106, 102, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 95, 0, 97, 98, 99,
	100, 103, 104, 105, 102, 0, 101, 0, 0, 0,
	0, 0, 0, 102, 0, 101, 96, 95, 0, 97,
	98, 99, 100, 103, 104, 96, 95, 0, 97, 98,
	99, 100, 102, 104, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 95, 0, 97, 98, 99,
	100, 34, 0, 0, 0, 26, 27, 28, 29, 30,
	31, 32, 0, 0, 0, 33, 0, 0, 0, 0,
	38, 39, 9, 10, 11, 12, 14, 19,
}
var parserPact = []int{

	88, -1000, 210, 3, 210, 59, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -18, -45, 86, 437, -1000, -44, -1000,
	-29, 233, -1000, 3, 59, -10, 101, 138, 138, 138,
	138, 138, 138, 138, 437, -26, -1000, -1000, 138, 138,
	210, -1000, -60, -1000, -1000, -45, -1000, -1000, -16, 128,
	-1000, 305, -1000, -1000, -1000, -1000, -1000, -26, -1000, 138,
	138, 138, 138, 27, -1000, -1000, 17, 17, -1000, -1000,
	-1000, 305, 305, 305, 305, 282, 258, 51, -1000, 138,
	305, 305, -54, -1000, -61, 59, 233, 128, -1000, 305,
	-1000, -46, -1000, 59, 138, 138, 138, 138, 138, 138,
	138, 138, 138, 138, 138, 138, 138, 15, 15, -1000,
	-55, 207, -1000, -65, -1000, -68, 17, 437, 437, -1000,
	230, 84, 210, -1000, -56, -1000, 138, -50, 161, 15,
	15, -1000, -1000, 60, 60, 60, 60, 359, 378, 350,
	328, -1000, 138, -1000, -1000, -1000, 56, 45, -19, 437,
	-1000, -1000, 184, 138, -32, 138, -1000, 437, -1000, -26,
	48, 138, -57, -1000, 161, -1000, 305, 47, -33, -1000,
	22, -1000, -1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 134, 173, 142, 14, 6, 35, 23, 22, 24,
	171, 42, 170, 7, 167, 20, 126, 0, 1, 3,
	2, 19, 5, 165, 159, 157, 155, 148, 147, 144,
	141, 135, 73, 4, 113,
}
var parserR1 = []int{

	0, 34, 1, 1, 2, 12, 12, 13, 13, 14,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 3, 16, 16, 16, 15, 15, 15, 15,
	15, 24, 24, 25, 18, 22, 22, 5, 5, 5,
	6, 6, 6, 6, 7, 10, 10, 8, 9, 9,
	9, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 20, 20, 21, 19, 19, 26,
	31, 31, 31, 27, 27, 28, 29, 33, 33, 32,
	23, 30, 11,
}
var parserR2 = []int{

	0, 4, 2, 1, 8, 1, 0, 3, 1, 2,
	4, 3, 2, 2, 2, 2, 2, 2, 7, 5,
	3, 2, 2, 1, 1, 1, 1, 1, 6, 1,
	5, 1, 0, 2, 2, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 3, 1, 1, 6, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 0, 2, 6, 3, 2,
	1, 1, 0, 1, 1, 3, 3, 2, 1, 1,
	4, 1, 1,
}
var parserChk = []int{

	-1000, -34, 4, -1, -2, -5, -6, -7, -8, 25,
	26, 27, 28, -10, 29, -3, 67, -1, -11, 30,
	69, 71, 5, -4, -5, -16, 8, 9, 10, 11,
	12, 13, 14, 18, 4, -11, -21, -22, 23, 24,
	71, 70, -9, -6, -7, 29, -8, -3, -11, 68,
	-16, -17, -26, -27, -28, -29, -30, -11, -21, 56,
	57, 58, 71, -31, 52, 53, 75, 76, 55, 50,
	51, -17, -17, -17, -17, -17, -17, -4, -19, 69,
	-17, -17, -12, -13, -14, -5, 73, 68, -15, -17,
	-23, 21, -22, 22, 69, 57, 56, 59, 60, 61,
	62, 46, 44, 63, 64, 65, 66, -17, -17, -17,
	-20, -17, 49, -32, 54, -33, -32, 15, 19, 5,
	-17, 72, 73, -11, -9, -15, 71, -11, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, 72, 73, 75, 76, -33, -4, -4, 70, 6,
	-13, 72, -17, 71, -18, 73, -20, 16, 20, 69,
	-4, 73, -24, -25, -17, 70, -17, -4, -19, 5,
	-17, 72, -18, 17, 70, 72,
}
var parserDef = []int{

	0, -2, 0, 0, 3, 0, -2, 38, -2, 40,
	41, 42, 43, 0, 0, 0, 0, 2, 0, 92,
	0, 0, 1, 22, 0, 0, 0, 82, 82, 82,
	82, 82, 82, 82, 0, 23, 24, 25, 82, 82,
	6, 44, 0, -2, 49, 50, 46, 21, 0, 82,
	12, 13, 51, 52, 53, 54, 55, 56, 57, 82,
	82, 82, -2, 0, 83, 84, 0, 0, 91, 80,
	81, 14, 15, 16, 17, 0, 0, 0, 76, 82,
	35, 36, 0, 5, 8, 0, 0, 82, 11, 26,
	27, 0, 29, 0, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 58, 59, 60,
	0, 0, 79, 0, 89, 0, 88, 0, 0, 20,
	0, 0, 0, 9, 0, 10, 82, 0, 0, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, -2, 85, 86, 87, 0, 0, 78, 0,
	7, 47, 0, -2, 0, 82, 74, 0, 19, 0,
	0, 82, 0, 31, 0, 90, 34, 0, 0, 4,
	0, 30, 33, 18, 77, 28,
}
var parserTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 76, 3, 3, 60, 3, 3,
	71, 72, 74, 57, 73, 56, 3, 59, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 67,
	61, 68, 62, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 69, 3, 70,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 63, 64, 65, 66,
}
var parserTok3 = []int{
	8217, 75, 0,
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
		//line new.y:98
		{
			yylex.(*lexer).prog = &Program{Functions: parserS[parserpt-2].functions, Statements: parserS[parserpt-1].statements}
			return 0
		}
	case 2:
		//line new.y:103
		{
			parserVAL.functions = append([]Function{parserS[parserpt-1].function}, parserS[parserpt-0].functions...)
		}
	case 3:
		//line new.y:104
		{
			parserVAL.functions = []Function{parserS[parserpt-0].function}
		}
	case 4:
		//line new.y:109
		{
			parserVAL.function = &Function{Type: parserS[parserpt-7].typeDef, Ident: parserS[parserpt-6].str, Params: parserS[parserpt-4].params, Stat: parserS[parserpt-1].statement}
		}
	case 5:
		//line new.y:112
		{
			parserVAL.params = parserS[parserpt-0].params
		}
	case 6:
		//line new.y:113
		{
			parserVAL.params
		}
	case 7:
		//line new.y:116
		{
			parserVAL.params = append([]Parameter{parserS[parserpt-2].param}, parserS[parserpt-0].params...)
		}
	case 8:
		//line new.y:117
		{
			parserVAL.params = []Parameter{parserS[parserpt-0].param}
		}
	case 9:
		//line new.y:120
		{
			parserVAL.param = Parameter{Type: parserS[parserpt-1].typeDef, Ident: parserS[parserpt-0].str}
		}
	case 10:
		//line new.y:124
		{
			parserVAL.statement = Statement{Type: parserS[parserpt-3].typeDef, Ident: parserS[parserpt-2].str, AssignRHS: parserS[parserpt-0].assign_rhs}
		}
	case 11:
		//line new.y:125
		{
			parserVAL.statement = Statement{AssignLHS: parserS[parserpt-2].assign_lhs, AssignRHS: parserS[parserpt-0].assign_rhs}
		}
	case 12:
		//line new.y:126
		{
			parserVAL.statement = Statement{AssignLHS: parserS[parserpt-0].assign_lhs}
		}
	case 13:
		//line new.y:127
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 14:
		//line new.y:128
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 15:
		//line new.y:129
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 16:
		//line new.y:130
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 17:
		//line new.y:131
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-0].expr}
		}
	case 18:
		//line new.y:132
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-5].expr, StatL: parserS[parserpt-3].statement, StatR: parserS[parserpt-1].statement}
		}
	case 19:
		//line new.y:133
		{
			parserVAL.statement = Statement{Expr: parserS[parserpt-3].expr, Stat: parserS[parserpt-1].statement}
		}
	case 20:
		//line new.y:134
		{
			parserVAL.statement = Statement{Stat: parserS[parserpt-1].statement}
		}
	case 21:
		//line new.y:135
		{
			parserVAL.statement = append([]Statement{parserS[parserpt-1].statement}, parserS[parserpt-0].statements...)
		}
	case 22:
		//line new.y:140
		{
			parserVAL.statements = Statement{Stat: parserS[parserpt-1].op}
		}
	case 23:
		//line new.y:144
		{
			parserVAL.assign_lhs = AssignLHS{Ident: parserS[parserpt-0].str}
		}
	case 24:
		//line new.y:145
		{
			parserVAL.assign_lhs = AssignLHS{ArrayElem: parserS[parserpt-0].array_elem}
		}
	case 25:
		//line new.y:146
		{
			parserVAL.assign_lhs = AssignLHS{PairElem: parserS[parserpt-0].pair_elem}
		}
	case 26:
		//line new.y:149
		{
			parserVAL.assign_rhs = AssignRHS{Expr: parserS[parserpt-0].expr}
		}
	case 27:
		//line new.y:150
		{
			parserVAL.assign_rhs = AssignRHS{ArrayLiter: parserS[parserpt-0].array_liter}
		}
	case 28:
		//line new.y:151
		{
			parserVAL.assign_rhs = AssignRHS{ExprL: parserS[parserpt-3].expr, ExprR: parserS[parserpt-1].expr}
		}
	case 29:
		//line new.y:152
		{
			parserVAL.assign_rhs = AssignRHS{PairElem: parserS[parserpt-0].pair_elem}
		}
	case 30:
		//line new.y:153
		{
			parserVAL.assign_rhs = AssignRHS{Ident: parserS[parserpt-3].str, Args: parserS[parserpt-1].arg}
		}
	case 31:
		//line new.y:156
		{
			parserVAL.arg = parserS[parserpt-0].arg
		}
	case 32:
		//line new.y:157
		{
			parserVAL.arg
		}
	case 33:
		//line new.y:160
		{
			parserVAL.arg = append([]Expr{parserS[parserpt-1].expr}, parserS[parserpt-0].expr...)
		}
	case 34:
		//line new.y:163
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].expr}
		}
	case 35:
		//line new.y:166
		{
			parserVAL.pair_elem = PairElem{Expr: parserS[parserpt-0].expr}
		}
	case 36:
		//line new.y:167
		{
			parserVAL.pair_elem = PairElem{Expr: parserS[parserpt-0].expr}
		}
	case 37:
		//line new.y:170
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 38:
		//line new.y:171
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 39:
		//line new.y:172
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 40:
		//line new.y:175
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 41:
		//line new.y:176
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 42:
		//line new.y:177
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 43:
		//line new.y:178
		{
			parserVAL.typeDef = BaseType{Type: parserS[parserpt-0].str}
		}
	case 44:
		//line new.y:181
		{
			parserVAL.typeDef = ArrayType{Type: parserS[parserpt-2].typeDef}
		}
	case 45:
		//line new.y:184
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 46:
		//line new.y:185
		{
			parserVAL.typeDef = TypeDef{Type: parserS[parserpt-0].typeDef}
		}
	case 47:
		//line new.y:188
		{
			parserVAL.typeDef = PairType{Lpair: parserS[parserpt-3].typeDef, Rpair: parserS[parserpt-1].typeDef}
		}
	case 48:
		//line new.y:191
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].typeDef}
		}
	case 49:
		//line new.y:192
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].typeDef}
		}
	case 50:
		//line new.y:193
		{
			parserVAL.typeDef = PairElemType{Type: parserS[parserpt-0].str}
		}
	case 51:
		//line new.y:196
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].int}
		}
	case 52:
		//line new.y:197
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].bool}
		}
	case 53:
		//line new.y:198
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].rune}
		}
	case 54:
		//line new.y:199
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].str}
		}
	case 55:
		//line new.y:200
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].pair}
		}
	case 56:
		//line new.y:201
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].str}
		}
	case 57:
		//line new.y:202
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-0].array_elem}
		}
	case 58:
		//line new.y:203
		{
			parserVAL.expr = Expr{UnaryOp: parserS[parserpt-1].op, Expr: parserS[parserpt-0].expr}
		}
	case 59:
		//line new.y:204
		{
			parserVAL.expr = Expr{UnaryOp: parserS[parserpt-1].op, Expr: parserS[parserpt-0].expr}
		}
	case 60:
		//line new.y:205
		{
			parserVAL.expr = Expr{UnaryOp: parserS[parserpt-1].op, Expr: parserS[parserpt-0].expr}
		}
	case 61:
		//line new.y:209
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 62:
		//line new.y:210
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 63:
		//line new.y:211
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 64:
		//line new.y:212
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 65:
		//line new.y:213
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 66:
		//line new.y:214
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 67:
		//line new.y:215
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].str, Rexpr: parserS[parserpt-0].expr}
		}
	case 68:
		//line new.y:216
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].str, Rexpr: parserS[parserpt-0].expr}
		}
	case 69:
		//line new.y:217
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 70:
		//line new.y:218
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 71:
		//line new.y:219
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 72:
		//line new.y:220
		{
			parserVAL.expr = Expr{Lexpr: parserS[parserpt-2].expr, BinaryOp: parserS[parserpt-1].op, Rexpr: parserS[parserpt-0].expr}
		}
	case 73:
		//line new.y:223
		{
			parserVAL.expr = Expr{Expr: parserS[parserpt-1].expr}
		}
	case 74:
		//line new.y:226
		{
			parserVAL.expr = append([]Expr{parserS[parserpt-2].expr}, parserS[parserpt-0].expr)
		}
	case 75:
		//line new.y:227
		{
			parserVAL.expr
		}
	case 76:
		//line new.y:230
		{
			parserVAL.array_elem = ArrayElem{Ident: parserS[parserpt-1].str, Exprs: parserS[parserpt-0].expr}
		}
	case 77:
		//line new.y:234
		{
			parserVAL.expr = append([]Expr{parserS[parserpt-4].expr}, parserS[parserpt-1].expr...)
		}
	case 78:
		//line new.y:235
		{
			parserVAL.expr = []Expr{parserS[parserpt-1].expr}
		}
	case 79:
		//line new.y:239
		{
			parserVAL.int = IntLiter{Sign: parserS[parserpt-1].sign, Int: parserS[parserpt-0].str}
		}
	case 80:
		//line new.y:242
		{
			parserVAL.sign = Sign{Sign: parserS[parserpt-0].str}
		}
	case 81:
		//line new.y:243
		{
			parserVAL.sign = Sign{Sign: parserS[parserpt-0].str}
		}
	case 82:
		//line new.y:244
		{
			parserVAL.sign
		}
	case 83:
		//line new.y:247
		{
			parserVAL.bool = Bool{Bool: parserS[parserpt-0].str}
		}
	case 84:
		//line new.y:248
		{
			parserVAL.bool = Bool{Bool: parserS[parserpt-0].str}
		}
	case 85:
		//line new.y:251
		{
			parserVAL.rune = Char{Char: parserS[parserpt-1].chr}
		}
	case 86:
		//line new.y:254
		{
			parserVAL.str = Str{Chars: parserS[parserpt-1].chr}
		}
	case 87:
		//line new.y:257
		{
			parserVAL.chr = append([]Char{parserS[parserpt-1].chr}, parserS[parserpt-0].chr...)
		}
	case 88:
		//line new.y:258
		{
			parserVAL.chr = []Char{parserS[parserpt-0].chr}
		}
	case 89:
		//line new.y:261
		{
			parserVAL.chr = Char{Char: parserS[parserpt-0].str}
		}
	case 90:
		//line new.y:263
		{
			parserVAL.array_liter = append([]Expr{parserS[parserpt-2].expr}, parserS[parserpt-1].expr...)
		}
	case 91:
		//line new.y:266
		{
			parserVAL.pair = PairLiter{PairLit: parserS[parserpt-0].str}
		}
	case 92:
		parserVAL.str = parserS[parserpt-0].str
	}
	goto parserstack /* stack new state and value */
}
