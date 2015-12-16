//line waccparser.y:2
package parser

import __yyfmt__ "fmt"

//line waccparser.y:2
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

//line waccparser.y:289

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 145,
	58, 100,
	-2, 97,
	-1, 146,
	58, 101,
	-2, 98,
}

const parserNprod = 108
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 609

var parserAct = []int{

	174, 173, 7, 26, 73, 36, 170, 150, 124, 11,
	144, 43, 4, 226, 13, 42, 130, 65, 66, 67,
	68, 69, 70, 71, 51, 25, 72, 37, 38, 35,
	37, 38, 128, 129, 37, 38, 61, 61, 88, 89,
	35, 52, 61, 53, 54, 55, 228, 199, 35, 57,
	56, 209, 193, 108, 109, 110, 111, 112, 113, 114,
	208, 45, 46, 59, 132, 58, 35, 29, 61, 179,
	64, 61, 61, 48, 131, 44, 47, 175, 123, 125,
	133, 93, 60, 62, 127, 39, 234, 143, 246, 216,
	239, 146, 148, 43, 230, 151, 154, 155, 156, 157,
	158, 159, 160, 161, 162, 163, 164, 165, 166, 153,
	192, 190, 75, 145, 231, 204, 203, 125, 210, 35,
	171, 201, 127, 176, 177, 172, 137, 41, 41, 76,
	122, 40, 77, 41, 182, 183, 243, 184, 201, 185,
	186, 135, 187, 188, 136, 35, 35, 227, 52, 41,
	53, 54, 55, 125, 41, 116, 57, 56, 127, 41,
	134, 194, 191, 236, 233, 201, 201, 152, 45, 46,
	59, 74, 58, 41, 141, 200, 197, 201, 198, 205,
	48, 50, 44, 47, 30, 31, 32, 33, 34, 41,
	41, 140, 138, 78, 146, 148, 214, 151, 217, 218,
	215, 212, 221, 171, 222, 220, 223, 115, 225, 207,
	116, 119, 117, 139, 115, 168, 145, 229, 237, 202,
	35, 35, 232, 87, 178, 235, 61, 142, 35, 181,
	79, 86, 180, 79, 101, 103, 100, 102, 9, 241,
	244, 195, 219, 196, 35, 90, 6, 245, 97, 99,
	98, 95, 96, 242, 2, 101, 103, 100, 102, 206,
	24, 118, 23, 35, 83, 82, 84, 80, 81, 35,
	12, 15, 16, 17, 18, 19, 20, 21, 86, 86,
	91, 22, 86, 27, 49, 14, 37, 38, 30, 31,
	32, 33, 34, 30, 31, 32, 33, 34, 85, 97,
	99, 98, 24, 126, 23, 28, 101, 103, 100, 102,
	149, 169, 12, 15, 16, 17, 18, 19, 20, 21,
	8, 63, 5, 22, 10, 3, 29, 14, 37, 38,
	30, 31, 32, 33, 34, 30, 31, 32, 33, 34,
	52, 1, 53, 54, 55, 0, 0, 0, 57, 56,
	30, 31, 32, 33, 147, 0, 0, 0, 0, 0,
	45, 46, 59, 0, 58, 0, 92, 213, 29, 0,
	0, 0, 48, 50, 44, 47, 97, 99, 98, 95,
	96, 106, 107, 101, 103, 100, 102, 104, 105, 30,
	31, 32, 33, 34, 0, 0, 0, 0, 0, 0,
	238, 97, 99, 98, 95, 96, 106, 107, 101, 103,
	100, 102, 104, 105, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 94, 97, 99, 98, 95,
	96, 106, 107, 101, 103, 100, 102, 104, 105, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 224,
	97, 99, 98, 95, 96, 106, 107, 101, 103, 100,
	102, 104, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 240, 97, 99, 98, 95, 96, 106, 107, 101,
	103, 100, 102, 104, 105, 0, 0, 0, 0, 0,
	0, 0, 0, 167, 97, 99, 98, 95, 96, 106,
	107, 101, 103, 100, 102, 104, 105, 0, 0, 0,
	0, 0, 0, 0, 211, 97, 99, 98, 95, 96,
	106, 107, 101, 103, 100, 102, 104, 105, 121, 0,
	0, 0, 0, 0, 0, 189, 0, 120, 0, 0,
	0, 0, 0, 0, 97, 99, 98, 95, 96, 106,
	107, 101, 103, 100, 102, 104, 105, 97, 99, 98,
	95, 96, 106, 107, 101, 103, 100, 102, 104, 105,
	97, 99, 98, 95, 96, 106, 107, 101, 103, 100,
	102, 104, 105, 97, 99, 98, 95, 96, 106, 0,
	101, 103, 100, 102, 104, 105, 97, 99, 98, 95,
	96, 0, 0, 101, 103, 100, 102, 104, 105,
}
var parserPact = []int{

	250, -1000, -1000, 240, 258, -1000, 17, 126, -1000, -1000,
	305, 14, -1000, -1000, 15, 2, 113, 113, 113, 113,
	113, 113, 113, 258, 107, -1000, -1000, -1000, 131, 224,
	-1000, -1000, -1000, -1000, 164, -1000, -1000, 113, 113, 238,
	-1000, 300, 13, 361, -1000, -1000, -1000, -1000, -1000, -1000,
	173, -1000, 113, 113, 113, 113, 113, 113, 113, -1000,
	148, 152, 202, -1000, 221, 530, 530, 530, 530, 530,
	517, 504, 125, 10, -1000, -1000, -1000, -1000, 6, 12,
	98, 82, 130, 151, 112, 169, 113, 320, 530, 530,
	263, -1000, 113, 105, -1, 113, 113, 113, 113, 113,
	113, 113, 113, 113, 113, 113, 113, 113, -1000, -1000,
	-1000, -1000, 259, 259, 432, 154, 6, -1000, 113, 9,
	258, 258, -1000, 93, -1000, 530, -1000, -1000, 165, 1,
	223, 220, 113, -1000, 113, -1000, 113, -1000, 113, 113,
	-1000, 113, 113, 475, 48, -1000, -1000, 164, 168, 47,
	-1000, -16, 6, 217, 259, 259, 187, 187, 187, -1000,
	-1000, -1000, -1000, 208, 208, 556, 543, -1000, 232, 115,
	-1000, -21, -1000, 114, 530, 160, 95, 90, 113, 200,
	-8, -17, 58, 530, 530, 530, 530, 530, 454, -1000,
	320, 359, 263, -1000, 25, 258, 258, 231, 263, -1000,
	-1000, 113, 113, 258, -1000, 386, 113, -55, -1000, -1000,
	-1000, -1000, 86, -1000, -22, -1000, 113, 69, 109, 258,
	-1000, 530, 103, 64, 113, 102, 159, -1000, 155, 336,
	-1000, -1000, 85, -1000, -1000, 410, -1000, 113, -1, -1000,
	-1000, 75, 216, -1000, 258, 63, -1000,
}
var parserPgo = []int{

	0, 341, 325, 322, 12, 320, 238, 2, 8, 305,
	0, 1, 311, 6, 310, 7, 5, 303, 24, 298,
	284, 25, 4, 3, 283, 10, 14,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 9, 8,
	8, 8, 8, 8, 8, 8, 8, 7, 7, 7,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 17, 11, 11, 11,
	18, 19, 19, 20, 16, 16, 24, 25, 25, 25,
	22, 22, 22, 21, 21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 3, 1,
	1, 1, 6, 5, 7, 3, 3, 3, 1, 12,
	1, 4, 1, 5, 7, 2, 2, 2, 2, 2,
	2, 7, 7, 5, 3, 2, 2, 2, 2, 3,
	4, 4, 4, 4, 4, 3, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 0,
	2, 4, 3, 1, 2, 2, 6, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	66, -22, 12, -26, 27, 13, 14, 15, 16, 17,
	18, 19, 23, 4, 2, -21, -23, -24, -9, 68,
	30, 31, 32, 33, 34, -18, -16, 28, 29, 68,
	5, 64, -22, -10, 69, 55, 56, 70, 67, -20,
	68, -18, 35, 37, 38, 39, 44, 43, 59, 57,
	68, 58, 68, -9, 68, -10, -10, -10, -10, -10,
	-10, -10, -7, -22, 64, 5, 22, 25, 62, 9,
	43, 44, 41, 40, 42, -19, 58, 59, -10, -10,
	7, -6, 66, 68, 64, 43, 44, 40, 42, 41,
	49, 47, 50, 48, 51, 52, 45, 46, -10, -10,
	-10, -10, -10, -10, -10, 59, 62, 60, 59, 9,
	20, 24, 5, 68, -8, -10, -17, -16, 26, 27,
	10, 68, 58, 68, 62, 43, 62, 44, 62, 62,
	40, 62, 58, -10, -25, -21, -23, 34, -22, -14,
	-15, -22, 62, -26, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, 61, 61, -12,
	-13, -22, -8, -11, -10, 68, -7, -7, 59, 68,
	9, 9, -11, -10, -10, -10, -10, -10, -10, 60,
	63, -4, 63, 68, -8, 24, 11, 61, 63, 68,
	61, 63, 59, 21, 25, -10, 59, 9, 68, 68,
	60, 60, -25, 8, -22, -15, 64, -7, -7, 11,
	-13, -10, -11, -7, 63, -11, 68, 61, 68, -10,
	25, 5, -7, 61, 22, -10, 61, 59, 64, 5,
	61, -11, -26, 61, 24, -7, 25,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 28,
	0, 0, 30, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 100, 101, 102, 0, 15,
	103, 104, 105, 106, 0, 16, 17, 0, 0, 0,
	1, 0, 0, 0, 58, 59, 60, 61, 62, 63,
	64, 65, 0, 0, 0, 0, 0, 0, 0, 93,
	0, 0, 0, 35, 15, 36, 37, 38, 39, 40,
	0, 0, 0, 0, 45, 46, 47, 48, 0, 0,
	0, 0, 0, 0, 0, 90, 0, 0, 94, 95,
	0, 27, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 66, 67,
	68, 69, 70, 71, 0, 0, 0, 107, 89, 0,
	0, 0, 44, 0, 49, 19, 20, 21, 0, 0,
	0, 64, 89, 18, 0, 55, 0, 56, 0, 0,
	57, 0, 0, 0, 0, -2, -2, 99, 0, 9,
	6, 0, 0, 0, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 0, 0,
	13, 0, 31, 0, 88, 0, 0, 0, 0, 0,
	0, 0, 0, 50, 51, 52, 53, 54, 0, 92,
	0, 0, 0, 7, 0, 0, 0, 0, 0, 14,
	33, 0, 89, 0, 43, 0, 89, 0, 25, 26,
	86, 91, 0, 4, 0, 5, 0, 0, 0, 0,
	12, 87, 0, 0, 0, 0, 0, 96, 0, 0,
	42, 10, 0, 34, 41, 0, 23, 89, 0, 11,
	22, 0, 0, 24, 0, 0, 29,
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
		//line waccparser.y:110
		{
			parserlex.(*Lexer).prog = &Program{ClassList: parserS[parserpt-3].classes, FunctionList: parserS[parserpt-2].functions, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 2:
		//line waccparser.y:115
		{
			parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)
		}
	case 3:
		//line waccparser.y:116
		{
			parserVAL.classes = []*Class{}
		}
	case 4:
		//line waccparser.y:119
		{
			if !checkClassIdent(parserS[parserpt-4].ident) {
				parserlex.Error("Invalid class name")
			}
			parserVAL.class = &Class{Ident: ClassType(parserS[parserpt-4].ident), FieldList: parserS[parserpt-2].fields, FunctionList: parserS[parserpt-1].functions}
		}
	case 5:
		//line waccparser.y:125
		{
			parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)
		}
	case 6:
		//line waccparser.y:126
		{
			parserVAL.fields = []Field{parserS[parserpt-0].field}
		}
	case 7:
		//line waccparser.y:128
		{
			parserVAL.field = Field{FieldType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 8:
		//line waccparser.y:130
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 9:
		//line waccparser.y:131
		{
			parserVAL.functions = []*Function{}
		}
	case 10:
		//line waccparser.y:134
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		//line waccparser.y:140
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 12:
		//line waccparser.y:146
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 13:
		//line waccparser.y:147
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 14:
		//line waccparser.y:149
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 15:
		//line waccparser.y:151
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 16:
		//line waccparser.y:152
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 17:
		//line waccparser.y:153
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 18:
		//line waccparser.y:154
		{
			parserVAL.assignlhs = FieldAssign{ObjectName: parserS[parserpt-2].ident, Field: parserS[parserpt-0].ident}
		}
	case 19:
		//line waccparser.y:156
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 20:
		//line waccparser.y:157
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 21:
		//line waccparser.y:158
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 22:
		//line waccparser.y:159
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:160
		{
			parserVAL.assignrhs = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 24:
		//line waccparser.y:161
		{
			parserVAL.assignrhs = CallInstance{Class: parserS[parserpt-5].ident, Func: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs}
		}
	case 25:
		//line waccparser.y:162
		{
			parserVAL.assignrhs = ThisInstance{Func: parserS[parserpt-0].ident}
		}
	case 26:
		//line waccparser.y:163
		{
			parserVAL.assignrhs = Instance{IdentLHS: parserS[parserpt-2].ident, IdentRHS: parserS[parserpt-0].ident}
		}
	case 27:
		//line waccparser.y:165
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 28:
		//line waccparser.y:166
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 29:
		//line waccparser.y:167
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			w := While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			d := Declare{DecType: parserS[parserpt-10].typedefinition, Lhs: parserS[parserpt-9].ident, Rhs: parserS[parserpt-7].assignrhs, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			parserVAL.stmts = []Statement{d, w}
		}
	case 30:
		//line waccparser.y:174
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:175
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 32:
		//line waccparser.y:177
		{
			parserVAL.stmt = parserS[parserpt-0].stmt
		}
	case 33:
		//line waccparser.y:179
		{
			parserVAL.stmt = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 34:
		//line waccparser.y:181
		{
			parserVAL.stmt = CallInstance{Class: parserS[parserpt-5].ident, Func: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs}
		}
	case 35:
		//line waccparser.y:183
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 36:
		//line waccparser.y:184
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 37:
		//line waccparser.y:185
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 38:
		//line waccparser.y:186
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 39:
		//line waccparser.y:187
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 40:
		//line waccparser.y:188
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 41:
		//line waccparser.y:189
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 42:
		//line waccparser.y:192
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			parserVAL.stmt = While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 43:
		//line waccparser.y:196
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		//line waccparser.y:197
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:198
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 46:
		//line waccparser.y:202
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 47:
		//line waccparser.y:205
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 48:
		//line waccparser.y:209
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 49:
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:216
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:217
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:218
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 54:
		//line waccparser.y:219
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 55:
		//line waccparser.y:220
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 56:
		//line waccparser.y:221
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 57:
		//line waccparser.y:222
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 58:
		//line waccparser.y:224
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 59:
		//line waccparser.y:225
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 60:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 61:
		//line waccparser.y:227
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 62:
		//line waccparser.y:228
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 63:
		//line waccparser.y:229
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 64:
		//line waccparser.y:230
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 65:
		//line waccparser.y:231
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 66:
		//line waccparser.y:232
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:233
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:234
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:235
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:236
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:237
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 72:
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		//line waccparser.y:244
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		//line waccparser.y:245
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		//line waccparser.y:246
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 81:
		//line waccparser.y:247
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 82:
		//line waccparser.y:248
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 83:
		//line waccparser.y:249
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 84:
		//line waccparser.y:250
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 85:
		//line waccparser.y:251
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 86:
		//line waccparser.y:254
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 87:
		//line waccparser.y:256
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 88:
		//line waccparser.y:257
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 89:
		//line waccparser.y:258
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 90:
		//line waccparser.y:260
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 91:
		//line waccparser.y:262
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 92:
		//line waccparser.y:263
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 93:
		//line waccparser.y:266
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 94:
		//line waccparser.y:269
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 95:
		//line waccparser.y:270
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 96:
		//line waccparser.y:272
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 97:
		//line waccparser.y:274
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:275
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 99:
		//line waccparser.y:276
		{
			parserVAL.pairelemtype = Pair
		}
	case 100:
		//line waccparser.y:278
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 101:
		//line waccparser.y:279
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 102:
		//line waccparser.y:280
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 103:
		//line waccparser.y:282
		{
			parserVAL.typedefinition = Int
		}
	case 104:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = Bool
		}
	case 105:
		//line waccparser.y:284
		{
			parserVAL.typedefinition = Char
		}
	case 106:
		//line waccparser.y:285
		{
			parserVAL.typedefinition = String
		}
	case 107:
		//line waccparser.y:287
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
