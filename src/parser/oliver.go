
//line oliver.y:2
package parser
import __yyfmt__ "fmt"
//line oliver.y:2
		
import . "ast"

//line oliver.y:7
type parserSymType struct{
	yys int
str string
number int
char  rune
functions []*Function
function *Function
stmt  interface{}
stmts []interface{}
assignrhs interface{}
assignlhs interface{}
expr  interface{}
exprs []interface{}
params  []Param
param Param
arglist []interface{}
bracketed []interface{}
pairliter interface{}
arrayliter ArrayLiter
pairelem  PairElem
arrayelem ArrayElem
typedefinition Type
pairelemtype Type
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
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line oliver.y:197


//line yacctab:1
var parserExca = []int{
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

var parserAct = []int{

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
var parserPact = []int{

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
var parserPgo = []int{

	0, 150, 141, 130, 117, 1, 7, 110, 0, 3,
	124, 2, 5, 120, 22, 119, 118, 27, 4, 11,
	115, 19,
}
var parserR1 = []int{

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
var parserR2 = []int{

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
var parserChk = []int{

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
var parserDef = []int{

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
var parserTok1 = []int{

	1,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64,
}
var parserTok3 = []int{
	0,
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
		//line oliver.y:82
		{
	                                          parserlex.(*Lexer).prog = &Program{FunctionList : parserS[parserpt-2].functions , StatList : parserS[parserpt-1].stmts , SymbolTable : &SymbolTable{Table: make(map[string]Type)}}
	                                         }
	case 2:
		//line oliver.y:86
		{ parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)}
	case 3:
		//line oliver.y:87
		{ parserVAL.functions = []*Function{} }
	case 4:
		//line oliver.y:90
		{ if !checkStats(parserS[parserpt-1].stmts) {
	          	parserlex.Error("Missing return statement")
	           }
	             parserVAL.function = &Function{Ident : parserS[parserpt-5].str, ReturnType : parserS[parserpt-6].typedefinition, StatList : parserS[parserpt-1].stmts}
	           }
	case 5:
		//line oliver.y:96
		{ if !checkStats(parserS[parserpt-1].stmts) {
	            	parserlex.Error("Missing return statement")
	            }
	             parserVAL.function = &Function{Ident : parserS[parserpt-6].str, ReturnType : parserS[parserpt-7].typedefinition, StatList : parserS[parserpt-1].stmts, ParameterTypes : parserS[parserpt-4].params}
	           }
	case 6:
		//line oliver.y:102
		{ parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)}
	case 7:
		//line oliver.y:103
		{ parserVAL.params = []Param{ parserS[parserpt-0].param } }
	case 8:
		//line oliver.y:105
		{ parserVAL.param = Param{ParamType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].str} }
	case 9:
		//line oliver.y:107
		{parserVAL.assignlhs = parserS[parserpt-0].str}
	case 10:
		//line oliver.y:108
		{parserVAL.assignlhs = parserS[parserpt-0].arrayelem}
	case 11:
		//line oliver.y:109
		{parserVAL.assignlhs = parserS[parserpt-0].pairelem}
	case 12:
		//line oliver.y:111
		{parserVAL.assignrhs = parserS[parserpt-0].expr}
	case 13:
		//line oliver.y:112
		{parserVAL.assignrhs = parserS[parserpt-0].arrayliter}
	case 14:
		//line oliver.y:113
		{parserVAL.assignrhs = parserS[parserpt-0].pairelem}
	case 15:
		//line oliver.y:114
		{ parserVAL.assignrhs = NewPair{FstExpr : parserS[parserpt-3].expr, SndExpr : parserS[parserpt-1].expr} }
	case 16:
		//line oliver.y:115
		{ parserVAL.assignrhs = Call{Ident : parserS[parserpt-3].str, ExprList : parserS[parserpt-1].exprs} }
	case 17:
		//line oliver.y:117
		{ parserVAL.stmts = append(parserS[parserpt-2].stmts,parserS[parserpt-0].stmt) }
	case 18:
		//line oliver.y:118
		{ parserVAL.stmts = []interface{}{parserS[parserpt-0].stmt} }
	case 19:
		//line oliver.y:121
		{ parserVAL.stmt = parserS[parserpt-0].number }
	case 20:
		//line oliver.y:122
		{ parserVAL.stmt = Declare{DecType : parserS[parserpt-3].typedefinition, Lhs : parserS[parserpt-2].str, Rhs : parserS[parserpt-0].assignrhs} }
	case 21:
		//line oliver.y:123
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].assignlhs, Rhs : parserS[parserpt-0].assignrhs} }
	case 22:
		//line oliver.y:124
		{ parserVAL.stmt = Read{parserS[parserpt-0].assignlhs} }
	case 23:
		//line oliver.y:125
		{ parserVAL.stmt = Free{parserS[parserpt-0].expr} }
	case 24:
		//line oliver.y:126
		{ parserVAL.stmt = Return{parserS[parserpt-0].expr} }
	case 25:
		//line oliver.y:127
		{ parserVAL.stmt = Exit{parserS[parserpt-0].expr} }
	case 26:
		//line oliver.y:128
		{ parserVAL.stmt = Print{parserS[parserpt-0].expr} }
	case 27:
		//line oliver.y:129
		{ parserVAL.stmt = Println{parserS[parserpt-0].expr} }
	case 28:
		//line oliver.y:130
		{ parserVAL.stmt = If{Conditional : parserS[parserpt-5].expr, ThenStat : parserS[parserpt-3].stmts, ElseStat : parserS[parserpt-1].stmts} }
	case 29:
		//line oliver.y:131
		{ parserVAL.stmt = While{Conditional : parserS[parserpt-3].expr, DoStat : parserS[parserpt-1].stmts} }
	case 30:
		//line oliver.y:132
		{ parserVAL.stmt = Scope{StatList : parserS[parserpt-1].stmts, SymbolTable : &SymbolTable{Table: make(map[string]Type)} } }
	case 31:
		//line oliver.y:134
		{ parserVAL.expr =  parserS[parserpt-0].number }
	case 32:
		//line oliver.y:135
		{ parserVAL.expr =  true }
	case 33:
		//line oliver.y:136
		{ parserVAL.expr =  false }
	case 34:
		//line oliver.y:137
		{ parserVAL.expr =  parserS[parserpt-0].char }
	case 35:
		//line oliver.y:138
		{ parserVAL.expr =  parserS[parserpt-0].str }
	case 36:
		//line oliver.y:139
		{ parserVAL.expr =  parserS[parserpt-0].pairliter }
	case 37:
		//line oliver.y:140
		{ parserVAL.expr =  Ident{Name : parserS[parserpt-0].str} }
	case 38:
		//line oliver.y:141
		{ parserVAL.expr =  parserS[parserpt-0].arrayelem }
	case 39:
		//line oliver.y:142
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 40:
		//line oliver.y:143
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 41:
		//line oliver.y:144
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 42:
		//line oliver.y:145
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 43:
		//line oliver.y:146
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 44:
		//line oliver.y:147
		{ parserVAL.expr = Unop{Unary : parserS[parserpt-1].number, Expr : parserS[parserpt-0].expr} }
	case 45:
		//line oliver.y:149
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 46:
		//line oliver.y:150
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 47:
		//line oliver.y:151
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 48:
		//line oliver.y:152
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 49:
		//line oliver.y:153
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 50:
		//line oliver.y:154
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 51:
		//line oliver.y:155
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 52:
		//line oliver.y:156
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 53:
		//line oliver.y:157
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 54:
		//line oliver.y:158
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 55:
		//line oliver.y:159
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 56:
		//line oliver.y:160
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 57:
		//line oliver.y:161
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : parserS[parserpt-1].number, Right : parserS[parserpt-0].expr} }
	case 58:
		//line oliver.y:162
		{ parserVAL.expr = parserS[parserpt-1].expr }
	case 59:
		//line oliver.y:164
		{ parserVAL.arrayliter = ArrayLiter{parserS[parserpt-1].exprs}}
	case 60:
		//line oliver.y:166
		{parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)}
	case 61:
		//line oliver.y:167
		{parserVAL.exprs = []interface{}{parserS[parserpt-0].expr}}
	case 62:
		//line oliver.y:168
		{parserVAL.exprs = []interface{}{}}
	case 63:
		//line oliver.y:170
		{parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].str, Exprs : parserS[parserpt-0].bracketed}}
	case 64:
		//line oliver.y:172
		{parserVAL.bracketed = append(parserS[parserpt-3].bracketed, parserS[parserpt-1].expr)}
	case 65:
		//line oliver.y:173
		{parserVAL.bracketed = []interface{}{parserS[parserpt-1].expr}}
	case 66:
		//line oliver.y:175
		{ parserVAL.pairliter =  NULL }
	case 67:
		//line oliver.y:177
		{ parserVAL.pairelem = PairElem{Fsnd: FST, Expr : parserS[parserpt-0].expr} }
	case 68:
		//line oliver.y:178
		{ parserVAL.pairelem = PairElem{Fsnd: SND, Expr : parserS[parserpt-0].expr} }
	case 69:
		//line oliver.y:180
		{ parserVAL.typedefinition = PairType{FstType : parserS[parserpt-3].pairelemtype, SndType : parserS[parserpt-1].pairelemtype} }
	case 70:
		//line oliver.y:182
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 71:
		//line oliver.y:183
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 72:
		//line oliver.y:184
		{ parserVAL.pairelemtype = Pair }
	case 73:
		//line oliver.y:186
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 74:
		//line oliver.y:187
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 75:
		//line oliver.y:188
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 76:
		//line oliver.y:190
		{ parserVAL.typedefinition =  Int }
	case 77:
		//line oliver.y:191
		{ parserVAL.typedefinition =  Bool }
	case 78:
		//line oliver.y:192
		{ parserVAL.typedefinition =  Char }
	case 79:
		//line oliver.y:193
		{ parserVAL.typedefinition =  String }
	case 80:
		//line oliver.y:195
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	}
	goto parserstack /* stack new state and value */
}
