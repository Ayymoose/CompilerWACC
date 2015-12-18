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
	"NEW",
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

//line waccparser.y:292

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	5, 25,
	65, 25,
	-2, 0,
	-1, 22,
	5, 25,
	65, 25,
	-2, 0,
	-1, 28,
	69, 104,
	-2, 27,
	-1, 125,
	22, 25,
	65, 25,
	-2, 0,
	-1, 126,
	26, 25,
	65, 25,
	-2, 0,
	-1, 203,
	26, 25,
	65, 25,
	-2, 0,
	-1, 205,
	5, 25,
	65, 25,
	-2, 0,
	-1, 209,
	23, 25,
	65, 25,
	-2, 0,
	-1, 225,
	5, 25,
	65, 25,
	-2, 0,
	-1, 244,
	26, 25,
	65, 25,
	-2, 0,
}

const parserNprod = 112
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 698

var parserAct = []int{

	187, 7, 186, 158, 181, 37, 134, 27, 4, 151,
	41, 45, 188, 241, 235, 221, 67, 68, 69, 70,
	71, 72, 73, 13, 74, 77, 26, 25, 38, 39,
	200, 31, 32, 33, 34, 35, 75, 246, 198, 94,
	95, 11, 215, 78, 234, 212, 79, 44, 147, 209,
	31, 32, 33, 34, 35, 114, 115, 116, 117, 118,
	119, 120, 179, 238, 123, 212, 127, 124, 41, 160,
	146, 144, 52, 43, 43, 53, 43, 30, 124, 40,
	36, 237, 132, 43, 64, 76, 135, 66, 160, 143,
	36, 137, 43, 150, 42, 30, 210, 80, 36, 45,
	161, 156, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 175, 148, 30, 99, 145, 36,
	153, 152, 142, 43, 162, 135, 43, 184, 185, 85,
	137, 183, 131, 159, 121, 43, 31, 32, 33, 34,
	154, 141, 191, 190, 192, 231, 193, 194, 128, 195,
	196, 130, 123, 223, 43, 212, 213, 133, 218, 211,
	182, 212, 135, 206, 93, 207, 199, 137, 202, 189,
	177, 129, 83, 30, 155, 82, 36, 81, 149, 92,
	204, 31, 32, 33, 34, 35, 107, 109, 106, 108,
	214, 244, 203, 225, 9, 178, 205, 122, 30, 30,
	29, 36, 36, 84, 220, 222, 156, 224, 217, 96,
	6, 227, 226, 228, 2, 65, 229, 51, 91, 160,
	136, 157, 233, 180, 8, 153, 152, 236, 5, 3,
	1, 239, 201, 103, 105, 104, 219, 159, 97, 0,
	107, 109, 106, 108, 182, 0, 245, 0, 0, 103,
	105, 104, 101, 102, 112, 208, 107, 109, 106, 108,
	110, 111, 0, 0, 243, 89, 88, 90, 86, 87,
	0, 0, 0, 0, 0, 138, 30, 62, 30, 36,
	0, 36, 30, 92, 0, 36, 0, 0, 0, 0,
	0, 0, 232, 139, 61, 38, 39, 0, 30, 0,
	0, 36, 54, 0, 55, 56, 57, 0, 0, 0,
	59, 58, 0, 30, 0, 0, 36, 30, 0, 0,
	36, 62, 47, 48, 63, 140, 60, 0, 0, 0,
	0, 0, 0, 0, 50, 41, 46, 49, 61, 0,
	0, 31, 32, 33, 34, 35, 54, 0, 55, 56,
	57, 0, 0, 0, 59, 58, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 47, 48, 63, 0,
	60, 23, 0, 22, 0, 0, 0, 0, 50, 28,
	46, 49, 12, 14, 15, 16, 17, 18, 19, 20,
	0, 0, 0, 21, 23, 0, 22, 24, 38, 39,
	31, 32, 33, 34, 35, 12, 14, 15, 16, 17,
	18, 19, 20, 0, 0, 0, 21, 62, 0, 0,
	24, 38, 39, 31, 32, 33, 34, 35, 0, 0,
	0, 0, 0, 0, 61, 0, 10, 0, 28, 0,
	0, 0, 54, 0, 55, 56, 57, 0, 0, 0,
	59, 58, 0, 0, 0, 0, 0, 0, 0, 98,
	0, 28, 47, 48, 63, 0, 60, 0, 0, 0,
	0, 0, 0, 0, 50, 41, 46, 49, 103, 105,
	104, 101, 102, 112, 113, 107, 109, 106, 108, 110,
	111, 103, 105, 104, 101, 102, 0, 0, 107, 109,
	106, 108, 240, 103, 105, 104, 101, 102, 112, 113,
	107, 109, 106, 108, 110, 111, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 100, 103, 105,
	104, 101, 102, 112, 113, 107, 109, 106, 108, 110,
	111, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 230, 103, 105, 104, 101, 102, 112, 113, 107,
	109, 106, 108, 110, 111, 0, 0, 0, 0, 0,
	0, 0, 0, 242, 103, 105, 104, 101, 102, 112,
	113, 107, 109, 106, 108, 110, 111, 0, 0, 0,
	0, 0, 0, 0, 0, 176, 103, 105, 104, 101,
	102, 112, 113, 107, 109, 106, 108, 110, 111, 0,
	0, 0, 0, 0, 0, 0, 216, 103, 105, 104,
	101, 102, 112, 113, 107, 109, 106, 108, 110, 111,
	126, 0, 0, 0, 0, 0, 0, 197, 0, 125,
	0, 0, 0, 0, 0, 0, 103, 105, 104, 101,
	102, 112, 113, 107, 109, 106, 108, 110, 111, 103,
	105, 104, 101, 102, 112, 113, 107, 109, 106, 108,
	110, 111, 103, 105, 104, 101, 102, 112, 113, 107,
	109, 106, 108, 110, 111, 103, 105, 104, 101, 102,
	0, 0, 107, 109, 106, 108, 110, 111,
}
var parserPact = []int{

	210, -1000, -1000, 204, 369, -1000, -59, 89, -1000, -1000,
	310, -59, -1000, -1000, -1, 406, 406, 406, 406, 406,
	406, 406, 369, 20, -59, 118, 116, 113, 193, 66,
	224, -1000, -1000, -1000, -1000, 104, -1000, -1000, 406, 406,
	202, 193, -1000, 392, -59, 462, -1000, -1000, -1000, -1000,
	-1000, -1000, 120, -1000, 406, 406, 406, 406, 406, 406,
	406, -59, 187, -1000, 4, -1000, 120, 631, 631, 631,
	631, 631, 618, 605, 61, -59, -1000, -1000, -1000, -1000,
	111, 90, 71, 21, -59, 266, 78, 26, 55, 7,
	52, 119, 406, 105, 631, 631, 19, -1000, 406, 37,
	-1, 406, 406, 406, 406, 406, 406, 406, 406, 406,
	406, 406, 406, 406, -1000, -1000, -1000, -1000, 192, 192,
	533, 110, -59, 0, 266, 369, 369, -1000, 15, 406,
	-1000, -1000, -1000, -1000, -1000, 631, -1000, -1000, -57, 109,
	406, 406, -1000, 406, -1000, 406, 406, -1000, 406, 406,
	576, -26, 118, 116, 104, -1000, 113, -34, -1000, -59,
	-1000, 266, 167, 192, 192, 138, 138, 138, -1000, -1000,
	-1000, -1000, 450, 450, 644, 208, -1000, 406, -1000, 184,
	101, -1000, -59, -1000, 27, 70, 97, 631, 96, 406,
	-19, 631, 631, 631, 631, 631, 555, -1000, 105, 150,
	19, -1000, -50, 369, 91, 369, 181, 19, -1000, 369,
	-1000, -1000, 406, 406, 487, -1000, -1000, 83, -1000, -59,
	-1000, 406, 18, -1000, 9, 369, -1000, 58, 631, 1,
	406, -1000, 92, 437, -1000, -1000, 8, -1000, -1000, 511,
	-1, -1000, -1000, 166, 369, 11, -1000,
}
var parserPgo = []int{

	0, 230, 229, 228, 8, 224, 194, 1, 6, 200,
	0, 2, 223, 4, 221, 3, 5, 220, 75, 218,
	217, 27, 36, 26, 7, 9, 23, 72,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 7, 7, 27, 27, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 17, 11, 11, 11,
	18, 19, 19, 20, 16, 16, 24, 25, 25, 25,
	25, 22, 22, 22, 22, 21, 21, 21, 21, 23,
	23, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 5, 6, 3, 1, 0, 12, 1, 3, 1,
	4, 1, 2, 2, 2, 2, 2, 2, 7, 7,
	5, 3, 2, 2, 2, 2, 5, 3, 4, 4,
	4, 4, 4, 3, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 3, 3, 1, 0,
	2, 4, 3, 1, 2, 2, 6, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	3, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	67, -22, 13, -26, 14, 15, 16, 17, 18, 19,
	20, 24, 4, 2, 28, -21, -23, -24, 69, -9,
	-27, 31, 32, 33, 34, 35, -18, -16, 29, 30,
	-27, 69, 5, 65, -22, -10, 70, 56, 57, 71,
	68, -20, -27, -18, 36, 38, 39, 40, 45, 44,
	60, 28, 11, 58, -27, -9, -27, -10, -10, -10,
	-10, -10, -10, -10, -7, -22, 65, 5, 23, 26,
	-27, 59, 59, 59, 10, 63, 44, 45, 42, 41,
	43, -19, 59, 60, -10, -10, 7, -6, 67, -27,
	65, 44, 45, 41, 43, 42, 50, 48, 51, 49,
	52, 53, 46, 47, -10, -10, -10, -10, -10, -10,
	-10, -27, 10, 60, 63, 21, 25, 5, -27, 60,
	61, 61, 61, -27, -8, -10, -17, -16, 9, 27,
	59, 63, 44, 63, 45, 63, 63, 41, 63, 59,
	-10, -25, -21, -23, 35, 69, -24, -14, -15, -22,
	69, 63, -26, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, 62, 60, -27, 62,
	-12, -13, -22, -8, -7, -7, -11, -10, 69, 60,
	-11, -10, -10, -10, -10, -10, -10, 61, 64, -4,
	64, -27, -8, 25, -11, 12, 62, 64, -27, 22,
	26, 62, 64, 60, -10, 61, 61, -25, 8, -22,
	-15, 65, -7, 62, -7, 12, -13, -7, -10, -11,
	64, 62, -27, -10, 26, 5, -7, 23, 62, -10,
	65, 5, 62, -26, 25, -7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, -2, 2, 0, 0, 8, 24,
	0, 0, 29, 31, 0, 0, 0, 0, 0, 0,
	0, 0, -2, 0, 0, 101, 102, 103, -2, 0,
	15, 105, 106, 107, 108, 0, 16, 17, 0, 0,
	0, 27, 1, 0, 0, 0, 56, 57, 58, 59,
	60, 61, 62, 63, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 0, 32, 15, 33, 34, 35,
	36, 37, 0, 0, 0, 0, 42, 43, 44, 45,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 90, 0, 0, 94, 95, 0, 23, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 64, 65, 66, 67, 68, 69,
	0, 0, 0, 0, 0, -2, -2, 41, 0, 89,
	109, 111, 110, 28, 47, 18, 19, 20, 0, 0,
	89, 0, 53, 0, 54, 0, 0, 55, 0, 0,
	0, 0, 97, 98, 99, 100, 0, 9, 6, 0,
	104, 0, 0, 70, 71, 72, 73, 74, 75, 76,
	77, 78, 79, 80, 81, 82, 83, 89, 85, 0,
	0, 13, 0, 30, 0, 0, 0, 88, 0, 0,
	0, 48, 49, 50, 51, 52, 0, 92, 0, 0,
	0, 7, 0, -2, 0, -2, 0, 0, 14, -2,
	40, 46, 0, 89, 0, 86, 91, 0, 4, 0,
	5, 0, 0, 84, 0, -2, 12, 0, 87, 0,
	0, 96, 0, 0, 39, 10, 0, 38, 21, 0,
	0, 11, 22, 0, -2, 0, 26,
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
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
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
			parserlex.(*Lexer).prog = &Program{ClassList: parserS[parserpt-3].classes, FunctionList: parserS[parserpt-2].functions, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 2:
		//line waccparser.y:116
		{
			parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)
		}
	case 3:
		//line waccparser.y:117
		{
			parserVAL.classes = []*Class{}
		}
	case 4:
		//line waccparser.y:120
		{
			if !checkClassIdent(parserS[parserpt-4].ident) {
				parserlex.Error("Invalid class name")
			}
			parserVAL.class = &Class{Ident: ClassType(parserS[parserpt-4].ident), FieldList: parserS[parserpt-2].fields, FunctionList: parserS[parserpt-1].functions}
		}
	case 5:
		//line waccparser.y:126
		{
			parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)
		}
	case 6:
		//line waccparser.y:127
		{
			parserVAL.fields = []Field{parserS[parserpt-0].field}
		}
	case 7:
		//line waccparser.y:130
		{
			parserVAL.field = Field{FieldType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 8:
		//line waccparser.y:132
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 9:
		//line waccparser.y:133
		{
			parserVAL.functions = []*Function{}
		}
	case 10:
		//line waccparser.y:136
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		//line waccparser.y:142
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 12:
		//line waccparser.y:148
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 13:
		//line waccparser.y:149
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 14:
		//line waccparser.y:151
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 15:
		//line waccparser.y:153
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 16:
		//line waccparser.y:154
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 17:
		//line waccparser.y:155
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 18:
		//line waccparser.y:157
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 19:
		//line waccparser.y:158
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 20:
		//line waccparser.y:159
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 21:
		//line waccparser.y:160
		{
			parserVAL.assignrhs = NewObject{Class: ClassType(parserS[parserpt-3].ident), Init: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:161
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:163
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 24:
		//line waccparser.y:164
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 25:
		//line waccparser.y:165
		{
			parserVAL.stmts = []Statement{}
		}
	case 26:
		//line waccparser.y:166
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			w := While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			d := Declare{DecType: parserS[parserpt-10].typedefinition, Lhs: parserS[parserpt-9].ident, Rhs: parserS[parserpt-7].assignrhs, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			parserVAL.stmts = []Statement{d, w}
		}
	case 27:
		//line waccparser.y:173
		{
			parserVAL.ident = parserS[parserpt-0].ident
		}
	case 28:
		//line waccparser.y:174
		{
			parserVAL.ident = Ident(string(parserS[parserpt-2].ident) + "." + string(parserS[parserpt-0].ident))
		}
	case 29:
		//line waccparser.y:176
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 30:
		//line waccparser.y:177
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:178
		{
			parserVAL.stmt = parserS[parserpt-0].stmt
		}
	case 32:
		//line waccparser.y:179
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 33:
		//line waccparser.y:180
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 34:
		//line waccparser.y:181
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 35:
		//line waccparser.y:182
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 36:
		//line waccparser.y:183
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 37:
		//line waccparser.y:184
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 38:
		//line waccparser.y:185
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 39:
		//line waccparser.y:186
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			parserVAL.stmt = While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 40:
		//line waccparser.y:190
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 41:
		//line waccparser.y:191
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 42:
		//line waccparser.y:192
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 43:
		//line waccparser.y:196
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 44:
		//line waccparser.y:199
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 45:
		//line waccparser.y:203
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 46:
		//line waccparser.y:207
		{
			parserVAL.stmt = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		//line waccparser.y:211
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		//line waccparser.y:213
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:216
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:217
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:218
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 54:
		//line waccparser.y:219
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 55:
		//line waccparser.y:220
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 56:
		//line waccparser.y:222
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 57:
		//line waccparser.y:223
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 58:
		//line waccparser.y:224
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 59:
		//line waccparser.y:225
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 60:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 61:
		//line waccparser.y:227
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 62:
		//line waccparser.y:228
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 63:
		//line waccparser.y:229
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 64:
		//line waccparser.y:230
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:231
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:232
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:233
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:234
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:235
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 70:
		//line waccparser.y:236
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		//line waccparser.y:244
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		//line waccparser.y:245
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		//line waccparser.y:246
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 81:
		//line waccparser.y:247
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 82:
		//line waccparser.y:248
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 83:
		//line waccparser.y:249
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 84:
		//line waccparser.y:250
		{
			parserVAL.expr = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 85:
		//line waccparser.y:251
		{
			parserVAL.expr = ThisInstance{parserS[parserpt-0].ident}
		}
	case 86:
		//line waccparser.y:253
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 87:
		//line waccparser.y:255
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 88:
		//line waccparser.y:256
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 89:
		//line waccparser.y:257
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 90:
		//line waccparser.y:259
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 91:
		//line waccparser.y:261
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 92:
		//line waccparser.y:262
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 93:
		//line waccparser.y:265
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 94:
		//line waccparser.y:268
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 95:
		//line waccparser.y:269
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 96:
		//line waccparser.y:271
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 97:
		//line waccparser.y:273
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:274
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 99:
		//line waccparser.y:275
		{
			parserVAL.pairelemtype = Pair
		}
	case 100:
		//line waccparser.y:276
		{
			parserVAL.pairelemtype = ClassType(parserS[parserpt-0].ident)
		}
	case 101:
		//line waccparser.y:278
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 102:
		//line waccparser.y:279
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 103:
		//line waccparser.y:280
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 104:
		//line waccparser.y:281
		{
			parserVAL.typedefinition = parserS[parserpt-0].ident
		}
	case 105:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = Int
		}
	case 106:
		//line waccparser.y:284
		{
			parserVAL.typedefinition = Bool
		}
	case 107:
		//line waccparser.y:285
		{
			parserVAL.typedefinition = Char
		}
	case 108:
		//line waccparser.y:286
		{
			parserVAL.typedefinition = String
		}
	case 109:
		//line waccparser.y:288
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	case 110:
		//line waccparser.y:289
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	case 111:
		//line waccparser.y:290
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
