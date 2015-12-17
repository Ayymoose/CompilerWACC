
//line waccparser.y:2
package parser
import __yyfmt__ "fmt"
//line waccparser.y:2

import (
. "ast"
)


//line waccparser.y:10
type parserSymType struct{
	yys int
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



var parserToknames = []string{
	"BEGIN",
	"END",
	"CLASS",
	"OPEN",
	"CLOSE",
	"DOT",
	"THIS",
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
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line waccparser.y:293


//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 139,
	58, 96,
	-2, 93,
	-1, 140,
	58, 97,
	-2, 94,
}

const parserNprod = 104
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 653

var parserAct = []int{

	122, 7, 165, 145, 25, 61, 4, 121, 24, 138,
	35, 43, 37, 38, 36, 36, 64, 65, 66, 67,
	68, 69, 70, 13, 71, 74, 224, 219, 51, 213,
	216, 191, 202, 34, 214, 72, 192, 183, 87, 88,
	11, 195, 75, 34, 196, 76, 42, 188, 181, 189,
	115, 34, 36, 107, 108, 109, 110, 111, 112, 113,
	29, 30, 31, 32, 33, 41, 50, 147, 41, 131,
	34, 28, 41, 39, 41, 41, 119, 114, 60, 40,
	115, 63, 114, 135, 73, 137, 41, 130, 124, 28,
	140, 163, 43, 41, 139, 149, 150, 151, 152, 153,
	154, 155, 156, 157, 158, 159, 160, 161, 28, 92,
	134, 129, 132, 77, 210, 116, 85, 148, 170, 168,
	169, 142, 34, 167, 61, 146, 124, 136, 173, 174,
	128, 175, 133, 176, 177, 41, 178, 179, 41, 120,
	84, 81, 80, 82, 78, 79, 34, 34, 222, 205,
	166, 182, 9, 143, 186, 185, 199, 187, 124, 84,
	28, 100, 102, 99, 101, 96, 98, 97, 94, 95,
	105, 193, 100, 102, 99, 101, 103, 104, 29, 30,
	31, 32, 33, 171, 28, 28, 140, 201, 203, 204,
	139, 198, 206, 207, 90, 27, 86, 209, 29, 30,
	31, 32, 33, 212, 89, 6, 2, 215, 26, 217,
	62, 49, 83, 184, 123, 34, 34, 142, 200, 146,
	34, 96, 98, 97, 223, 166, 144, 164, 100, 102,
	99, 101, 172, 190, 34, 8, 5, 3, 194, 1,
	0, 0, 221, 0, 0, 0, 126, 34, 0, 0,
	0, 34, 0, 28, 28, 0, 0, 0, 28, 0,
	0, 0, 125, 0, 37, 38, 0, 211, 0, 0,
	0, 52, 28, 53, 54, 55, 0, 0, 0, 57,
	56, 0, 0, 0, 0, 28, 0, 0, 0, 28,
	0, 45, 46, 59, 127, 58, 29, 30, 31, 32,
	141, 0, 0, 48, 36, 44, 47, 29, 30, 31,
	32, 33, 52, 0, 53, 54, 55, 0, 0, 0,
	57, 56, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 45, 46, 59, 0, 58, 0, 23, 0,
	22, 0, 0, 0, 48, 36, 44, 47, 12, 14,
	15, 16, 17, 18, 19, 20, 0, 0, 0, 21,
	0, 23, 0, 22, 37, 38, 29, 30, 31, 32,
	33, 12, 14, 15, 16, 17, 18, 19, 20, 0,
	0, 0, 21, 0, 0, 0, 0, 37, 38, 29,
	30, 31, 32, 33, 0, 0, 0, 52, 0, 53,
	54, 55, 10, 0, 36, 57, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 45, 46, 59,
	0, 58, 0, 0, 0, 91, 0, 36, 0, 48,
	36, 44, 47, 96, 98, 97, 94, 95, 105, 106,
	100, 102, 99, 101, 103, 104, 96, 98, 97, 94,
	95, 0, 0, 100, 102, 99, 101, 218, 96, 98,
	97, 94, 95, 105, 106, 100, 102, 99, 101, 103,
	104, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 93, 96, 98, 97, 94, 95, 105, 106,
	100, 102, 99, 101, 103, 104, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 208, 96, 98, 97,
	94, 95, 105, 106, 100, 102, 99, 101, 103, 104,
	0, 0, 0, 0, 0, 0, 0, 0, 220, 96,
	98, 97, 94, 95, 105, 106, 100, 102, 99, 101,
	103, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	162, 96, 98, 97, 94, 95, 105, 106, 100, 102,
	99, 101, 103, 104, 0, 0, 0, 0, 0, 0,
	0, 197, 96, 98, 97, 94, 95, 105, 106, 100,
	102, 99, 101, 103, 104, 118, 0, 0, 0, 0,
	0, 0, 180, 0, 117, 0, 0, 0, 0, 0,
	0, 96, 98, 97, 94, 95, 105, 106, 100, 102,
	99, 101, 103, 104, 96, 98, 97, 94, 95, 105,
	106, 100, 102, 99, 101, 103, 104, 96, 98, 97,
	94, 95, 105, 106, 100, 102, 99, 101, 103, 104,
	96, 98, 97, 94, 95, 0, 0, 100, 102, 99,
	101, 103, 104,
}
var parserPact = []int{

	202, -1000, -1000, 199, 336, -1000, -54, 74, -1000, -1000,
	277, -53, -1000, -1000, -16, 362, 362, 362, 362, 362,
	362, 362, 336, 20, -1000, -1000, -1000, 51, 101, -1000,
	-1000, -1000, -1000, 57, -1000, -1000, 187, 362, 362, 197,
	-1000, 359, -53, 418, -1000, -1000, -1000, -1000, -1000, -1000,
	82, -1000, 362, 362, 362, 362, 362, 362, 362, -1000,
	18, 55, -1000, 82, 587, 587, 587, 587, 587, 574,
	561, 71, -53, -1000, -1000, -1000, -1000, 236, 68, 25,
	50, 70, 21, 69, 362, 266, -54, 587, 587, 168,
	-1000, 362, 5, -16, 362, 362, 362, 362, 362, 362,
	362, 362, 362, 362, 362, 362, 362, -1000, -1000, -1000,
	-1000, 181, 181, 489, 30, 236, -1000, 336, 336, -1000,
	-12, -1000, 587, -1000, -1000, 59, 174, 362, 362, -1000,
	362, -1000, 362, 362, -1000, 362, 362, 532, -15, -1000,
	-1000, 57, 66, -1000, -26, -1000, -53, 236, 130, 181,
	181, 114, 114, 114, -1000, -1000, -1000, -1000, 406, 406,
	600, 125, -1000, 146, -14, -1000, -53, -1000, 10, 11,
	362, -54, -19, 587, 587, 587, 587, 587, 587, 511,
	-1000, 266, 148, 168, -1000, -32, 336, 336, 138, 168,
	-1000, 336, -1000, 443, -1000, -1000, 362, -1000, 53, -1000,
	-53, -1000, 362, 4, 29, 336, -1000, 8, 362, 587,
	-1000, 23, 393, -1000, -1000, 22, -1000, 467, -16, -1000,
	-1000, 124, 336, 1, -1000,
}
var parserPgo = []int{

	0, 239, 237, 236, 6, 235, 152, 1, 7, 195,
	0, 232, 227, 2, 226, 3, 10, 214, 28, 212,
	211, 8, 35, 4, 208, 9, 23, 66,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 7, 27, 27, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 17, 11, 11, 11, 18, 19, 19, 20,
	16, 16, 24, 25, 25, 25, 22, 22, 22, 21,
	21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 3, 3, 1, 12, 1, 3, 1, 4,
	1, 2, 2, 2, 2, 2, 2, 7, 7, 5,
	3, 2, 2, 2, 2, 3, 4, 4, 4, 4,
	4, 3, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 0, 2, 4, 3, 1,
	2, 2, 6, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	66, -22, 12, -26, 13, 14, 15, 16, 17, 18,
	19, 23, 4, 2, -21, -23, -24, -9, -27, 30,
	31, 32, 33, 34, -18, -16, 68, 28, 29, -27,
	5, 64, -22, -10, 69, 55, 56, 70, 67, -20,
	-27, -18, 35, 37, 38, 39, 44, 43, 59, 57,
	-27, 58, -9, -27, -10, -10, -10, -10, -10, -10,
	-10, -7, -22, 64, 5, 22, 25, 62, 43, 44,
	41, 40, 42, -19, 58, 59, 9, -10, -10, 7,
	-6, 66, -27, 64, 43, 44, 40, 42, 41, 49,
	47, 50, 48, 51, 52, 45, 46, -10, -10, -10,
	-10, -10, -10, -10, 59, 62, 60, 20, 24, 5,
	-27, -8, -10, -17, -16, 26, 10, 58, 62, 43,
	62, 44, 62, 62, 40, 62, 58, -10, -25, -21,
	-23, 34, -22, -27, -14, -15, -22, 62, -26, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, 61, 61, -12, -13, -22, -8, -7, -7,
	59, 9, -11, -10, -10, -10, -10, -10, -10, -10,
	60, 63, -4, 63, -27, -8, 24, 11, 61, 63,
	-27, 21, 25, -10, -27, 60, 63, 60, -25, 8,
	-22, -15, 64, -7, -7, 11, -13, -7, 63, -10,
	61, -27, -10, 25, 5, -7, 22, -10, 64, 5,
	61, -26, 24, -7, 25,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 24,
	0, 0, 28, 30, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 97, 98, 0, 15, 99,
	100, 101, 102, 0, 16, 17, 26, 0, 0, 0,
	1, 0, 0, 0, 54, 55, 56, 57, 58, 59,
	60, 61, 0, 0, 0, 0, 0, 0, 0, 89,
	0, 0, 31, 15, 32, 33, 34, 35, 36, 0,
	0, 0, 0, 41, 42, 43, 44, 0, 0, 0,
	0, 0, 0, 86, 0, 0, 0, 90, 91, 0,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 62, 63, 64,
	65, 66, 67, 0, 0, 0, 103, 0, 0, 40,
	0, 45, 18, 19, 20, 0, 0, 85, 0, 51,
	0, 52, 0, 0, 53, 0, 0, 0, 0, -2,
	-2, 95, 0, 27, 9, 6, 0, 0, 0, 68,
	69, 70, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 0, 0, 13, 0, 29, 0, 0,
	0, 0, 0, 84, 46, 47, 48, 49, 50, 0,
	88, 0, 0, 0, 7, 0, 0, 0, 0, 0,
	14, 0, 39, 0, 22, 82, 0, 87, 0, 4,
	0, 5, 0, 0, 0, 0, 12, 0, 0, 83,
	92, 0, 0, 38, 10, 0, 37, 0, 0, 11,
	21, 0, 0, 0, 25,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70,
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
		//line waccparser.y:111
		{
	                                         parserlex.(*Lexer).prog = &Program{ClassList : parserS[parserpt-3].classes , FunctionList : parserS[parserpt-2].functions , StatList : parserS[parserpt-1].stmts , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
	                                         }
	case 2:
		//line waccparser.y:116
		{ parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)}
	case 3:
		//line waccparser.y:117
		{ parserVAL.classes = []*Class{} }
	case 4:
		//line waccparser.y:120
		{ if !checkClassIdent(parserS[parserpt-4].ident) {
	                                                         	parserlex.Error("Invalid class name")
	                                                     }
	                                                     parserVAL.class = &Class{Ident : ClassType(parserS[parserpt-4].ident), FieldList : parserS[parserpt-2].fields , FunctionList : parserS[parserpt-1].functions}
	                                                   }
	case 5:
		//line waccparser.y:126
		{ parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)}
	case 6:
		//line waccparser.y:127
		{ parserVAL.fields = []Field{ parserS[parserpt-0].field } }
	case 7:
		//line waccparser.y:129
		{ parserVAL.field = Field{FieldType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 8:
		//line waccparser.y:131
		{ parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)}
	case 9:
		//line waccparser.y:132
		{ parserVAL.functions = []*Function{} }
	case 10:
		//line waccparser.y:135
		{ if !checkStats(parserS[parserpt-1].stmts) {
	          	parserlex.Error("Missing return statement")
	           }
	             parserVAL.function = &Function{Ident : parserS[parserpt-5].ident, ReturnType : parserS[parserpt-6].typedefinition, StatList : parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 11:
		//line waccparser.y:141
		{ if !checkStats(parserS[parserpt-1].stmts) {
	            	parserlex.Error("Missing return statement")
	            }
	             parserVAL.function = &Function{Ident : parserS[parserpt-6].ident, ReturnType : parserS[parserpt-7].typedefinition, StatList : parserS[parserpt-1].stmts, ParameterTypes : parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 12:
		//line waccparser.y:147
		{ parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)}
	case 13:
		//line waccparser.y:148
		{ parserVAL.params = []Param{ parserS[parserpt-0].param } }
	case 14:
		//line waccparser.y:150
		{ parserVAL.param = Param{ParamType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 15:
		//line waccparser.y:152
		{parserVAL.assignlhs = parserS[parserpt-0].ident}
	case 16:
		//line waccparser.y:153
		{parserVAL.assignlhs = parserS[parserpt-0].arrayelem}
	case 17:
		//line waccparser.y:154
		{parserVAL.assignlhs = parserS[parserpt-0].pairelem}
	case 18:
		//line waccparser.y:157
		{parserVAL.assignrhs = parserS[parserpt-0].expr}
	case 19:
		//line waccparser.y:158
		{parserVAL.assignrhs = parserS[parserpt-0].arrayliter}
	case 20:
		//line waccparser.y:159
		{parserVAL.assignrhs = parserS[parserpt-0].pairelem}
	case 21:
		//line waccparser.y:160
		{ parserVAL.assignrhs = NewPair{FstExpr : parserS[parserpt-3].expr, SndExpr : parserS[parserpt-1].expr, Pos : parserS[parserpt-5].pos, FileText :&parserlex.(*Lexer).input } }
	case 22:
		//line waccparser.y:163
		{ parserVAL.assignrhs = ThisInstance{parserS[parserpt-0].ident} }
	case 23:
		//line waccparser.y:166
		{ parserVAL.stmts = append(parserS[parserpt-2].stmts,parserS[parserpt-0].stmt)   }
	case 24:
		//line waccparser.y:167
		{ parserVAL.stmts = []Statement{parserS[parserpt-0].stmt} }
	case 25:
		//line waccparser.y:168
		{
	                                                                                                                 stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                                                                                  w := While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-11].pos, FileText :&parserlex.(*Lexer).input}
	                                                                                                                  d := Declare{DecType : parserS[parserpt-10].typedefinition, Lhs : parserS[parserpt-9].ident, Rhs : parserS[parserpt-7].assignrhs, Pos : parserS[parserpt-11].pos ,FileText :&parserlex.(*Lexer).input }
	                                                                                                                  parserVAL.stmts = []Statement{d,w}
	                                                                                                                }
	case 26:
		//line waccparser.y:175
		{ parserVAL.ident = parserS[parserpt-0].ident}
	case 27:
		//line waccparser.y:176
		{ parserVAL.ident = Ident(string(parserS[parserpt-2].ident)+"."+string(parserS[parserpt-0].ident))}
	case 28:
		//line waccparser.y:178
		{ parserVAL.stmt = Skip{Pos : parserS[parserpt-0].pos ,FileText :&parserlex.(*Lexer).input } }
	case 29:
		//line waccparser.y:179
		{ parserVAL.stmt = Declare{DecType : parserS[parserpt-3].typedefinition, Lhs : parserS[parserpt-2].ident, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input } }
	case 30:
		//line waccparser.y:181
		{ parserVAL.stmt = parserS[parserpt-0].stmt }
	case 31:
		//line waccparser.y:187
		{ parserVAL.stmt = Read{ &parserlex.(*Lexer).input, parserS[parserpt-1].pos , parserS[parserpt-0].assignlhs, } }
	case 32:
		//line waccparser.y:188
		{ parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 33:
		//line waccparser.y:189
		{ parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 34:
		//line waccparser.y:190
		{ parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 35:
		//line waccparser.y:191
		{ parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 36:
		//line waccparser.y:192
		{ parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 37:
		//line waccparser.y:193
		{ parserVAL.stmt = If{Conditional : parserS[parserpt-5].expr, ThenStat : parserS[parserpt-3].stmts, ElseStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input } }
	case 38:
		//line waccparser.y:196
		{
	                                                                                  stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                                                  parserVAL.stmt = While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input}
	                                                                                  }
	case 39:
		//line waccparser.y:200
		{ parserVAL.stmt = While{Conditional : parserS[parserpt-3].expr, DoStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input} }
	case 40:
		//line waccparser.y:201
		{ parserVAL.stmt = Scope{StatList : parserS[parserpt-1].stmts, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input } }
	case 41:
		//line waccparser.y:202
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 42:
		//line waccparser.y:206
		{ parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 43:
		//line waccparser.y:209
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 44:
		//line waccparser.y:213
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 45:
		//line waccparser.y:218
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].assignlhs, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 46:
		//line waccparser.y:219
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 47:
		//line waccparser.y:220
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 48:
		//line waccparser.y:221
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 49:
		//line waccparser.y:222
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 50:
		//line waccparser.y:223
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 51:
		//line waccparser.y:224
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 52:
		//line waccparser.y:225
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 53:
		//line waccparser.y:226
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : MUL,  Right : parserS[parserpt-2].ident,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 54:
		//line waccparser.y:228
		{ parserVAL.expr =  parserS[parserpt-0].integer }
	case 55:
		//line waccparser.y:229
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 56:
		//line waccparser.y:230
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 57:
		//line waccparser.y:231
		{ parserVAL.expr =  parserS[parserpt-0].character }
	case 58:
		//line waccparser.y:232
		{ parserVAL.expr =  parserS[parserpt-0].stringconst }
	case 59:
		//line waccparser.y:233
		{ parserVAL.expr =  parserS[parserpt-0].pairliter }
	case 60:
		//line waccparser.y:234
		{ parserVAL.expr =  parserS[parserpt-0].ident }
	case 61:
		//line waccparser.y:235
		{ parserVAL.expr =  parserS[parserpt-0].arrayelem }
	case 62:
		//line waccparser.y:236
		{ parserVAL.expr = Unop{Unary : NOT, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 63:
		//line waccparser.y:237
		{ parserVAL.expr = Unop{Unary : LEN, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 64:
		//line waccparser.y:238
		{ parserVAL.expr = Unop{Unary : ORD, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 65:
		//line waccparser.y:239
		{ parserVAL.expr = Unop{Unary : CHR, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 66:
		//line waccparser.y:240
		{ parserVAL.expr = Unop{Unary : SUB, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 67:
		//line waccparser.y:241
		{ parserVAL.expr = parserS[parserpt-0].expr }
	case 68:
		//line waccparser.y:242
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 69:
		//line waccparser.y:243
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : SUB,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 70:
		//line waccparser.y:244
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 71:
		//line waccparser.y:245
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 72:
		//line waccparser.y:246
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 73:
		//line waccparser.y:247
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 74:
		//line waccparser.y:248
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 75:
		//line waccparser.y:249
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 76:
		//line waccparser.y:250
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 77:
		//line waccparser.y:251
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : EQ,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 78:
		//line waccparser.y:252
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : NEQ,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 79:
		//line waccparser.y:253
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : AND,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 80:
		//line waccparser.y:254
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : OR,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 81:
		//line waccparser.y:255
		{ parserVAL.expr = parserS[parserpt-1].expr }
	case 82:
		//line waccparser.y:258
		{ parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs } }
	case 83:
		//line waccparser.y:260
		{parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)}
	case 84:
		//line waccparser.y:261
		{parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}}
	case 85:
		//line waccparser.y:262
		{parserVAL.exprs = []Evaluation{}}
	case 86:
		//line waccparser.y:264
		{parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs : parserS[parserpt-0].exprs, Pos : parserS[parserpt-1].pos,FileText :&parserlex.(*Lexer).input  } }
	case 87:
		//line waccparser.y:266
		{parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)}
	case 88:
		//line waccparser.y:267
		{parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}}
	case 89:
		//line waccparser.y:270
		{ parserVAL.pairliter =  PairLiter{} }
	case 90:
		//line waccparser.y:273
		{ parserVAL.pairelem = PairElem{Fsnd: Fst, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 91:
		//line waccparser.y:274
		{ parserVAL.pairelem = PairElem{Fsnd: Snd, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 92:
		//line waccparser.y:276
		{ parserVAL.typedefinition = PairType{FstType : parserS[parserpt-3].pairelemtype, SndType : parserS[parserpt-1].pairelemtype} }
	case 93:
		//line waccparser.y:278
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 94:
		//line waccparser.y:279
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 95:
		//line waccparser.y:280
		{ parserVAL.pairelemtype = Pair}
	case 96:
		//line waccparser.y:282
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 97:
		//line waccparser.y:283
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 98:
		//line waccparser.y:284
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 99:
		//line waccparser.y:286
		{ parserVAL.typedefinition =  Int }
	case 100:
		//line waccparser.y:287
		{ parserVAL.typedefinition =  Bool }
	case 101:
		//line waccparser.y:288
		{ parserVAL.typedefinition =  Char }
	case 102:
		//line waccparser.y:289
		{ parserVAL.typedefinition =  String }
	case 103:
		//line waccparser.y:291
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	}
	goto parserstack /* stack new state and value */
}
