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

//line waccparser.y:285

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 143,
	59, 97,
	-2, 94,
	-1, 144,
	59, 98,
	-2, 95,
}

const parserNprod = 105
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 648

var parserAct = []int{

	127, 170, 7, 149, 175, 4, 142, 126, 26, 37,
	226, 44, 38, 39, 76, 208, 66, 67, 68, 69,
	70, 71, 72, 13, 25, 73, 52, 63, 231, 223,
	221, 35, 77, 36, 123, 78, 197, 37, 220, 90,
	91, 35, 198, 210, 199, 200, 200, 202, 188, 35,
	200, 194, 37, 195, 110, 111, 112, 113, 114, 115,
	116, 30, 31, 32, 33, 34, 135, 42, 51, 35,
	42, 42, 186, 29, 75, 40, 119, 42, 151, 42,
	62, 42, 118, 65, 134, 119, 139, 41, 141, 136,
	42, 29, 168, 79, 42, 44, 138, 144, 153, 154,
	155, 156, 157, 158, 159, 160, 161, 162, 163, 164,
	165, 29, 95, 143, 129, 80, 217, 74, 137, 120,
	152, 118, 11, 35, 173, 174, 176, 172, 43, 117,
	88, 177, 176, 179, 60, 180, 178, 181, 182, 167,
	183, 184, 53, 124, 54, 55, 56, 42, 35, 35,
	58, 57, 133, 129, 187, 125, 63, 140, 147, 190,
	87, 9, 46, 47, 61, 29, 59, 229, 176, 191,
	212, 132, 192, 193, 49, 37, 45, 48, 201, 89,
	84, 83, 85, 81, 82, 129, 103, 105, 102, 104,
	29, 29, 207, 204, 209, 144, 211, 213, 87, 28,
	214, 215, 92, 6, 93, 2, 146, 27, 50, 219,
	150, 143, 86, 128, 64, 222, 148, 224, 35, 189,
	35, 99, 101, 100, 35, 169, 8, 5, 103, 105,
	102, 104, 230, 3, 1, 0, 171, 0, 0, 35,
	196, 0, 0, 0, 205, 0, 0, 0, 0, 228,
	0, 0, 35, 0, 0, 0, 35, 0, 0, 0,
	29, 0, 29, 0, 0, 0, 29, 30, 31, 32,
	33, 34, 0, 0, 0, 218, 0, 0, 0, 60,
	0, 29, 30, 31, 32, 33, 34, 53, 0, 54,
	55, 56, 0, 0, 29, 58, 57, 0, 29, 30,
	31, 32, 33, 34, 146, 206, 150, 46, 47, 61,
	0, 59, 23, 171, 22, 0, 0, 0, 0, 49,
	37, 45, 48, 12, 14, 15, 16, 17, 18, 19,
	20, 0, 0, 0, 21, 0, 0, 0, 24, 38,
	39, 30, 31, 32, 33, 34, 30, 31, 32, 33,
	145, 130, 60, 38, 39, 0, 0, 0, 0, 0,
	53, 0, 54, 55, 56, 0, 0, 0, 58, 57,
	0, 0, 0, 0, 0, 0, 0, 10, 0, 37,
	46, 47, 61, 131, 59, 23, 0, 22, 0, 0,
	0, 0, 49, 37, 45, 48, 12, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 21, 0, 0,
	0, 24, 38, 39, 30, 31, 32, 33, 34, 99,
	101, 100, 97, 98, 108, 109, 103, 105, 102, 104,
	106, 107, 0, 0, 0, 0, 0, 99, 101, 100,
	97, 98, 108, 225, 103, 105, 102, 104, 106, 107,
	94, 0, 37, 99, 101, 100, 97, 98, 108, 109,
	103, 105, 102, 104, 106, 107, 99, 101, 100, 97,
	98, 0, 0, 103, 105, 102, 104, 96, 99, 101,
	100, 97, 98, 108, 109, 103, 105, 102, 104, 106,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 216, 99, 101, 100, 97, 98, 108, 109, 103,
	105, 102, 104, 106, 107, 0, 0, 0, 0, 0,
	0, 0, 0, 227, 99, 101, 100, 97, 98, 108,
	109, 103, 105, 102, 104, 106, 107, 0, 0, 0,
	0, 0, 0, 0, 0, 166, 99, 101, 100, 97,
	98, 108, 109, 103, 105, 102, 104, 106, 107, 0,
	0, 0, 0, 0, 0, 0, 203, 99, 101, 100,
	97, 98, 108, 109, 103, 105, 102, 104, 106, 107,
	122, 0, 0, 0, 0, 0, 0, 185, 0, 121,
	0, 0, 0, 0, 0, 0, 99, 101, 100, 97,
	98, 108, 109, 103, 105, 102, 104, 106, 107, 99,
	101, 100, 97, 98, 108, 109, 103, 105, 102, 104,
	106, 107, 99, 101, 100, 97, 98, 108, 109, 103,
	105, 102, 104, 106, 107, 99, 101, 100, 97, 98,
	0, 0, 103, 105, 102, 104, 106, 107,
}
var parserPact = []int{

	201, -1000, -1000, 197, 310, -1000, -60, 82, -1000, -1000,
	251, -32, -1000, -1000, -17, 106, 106, 106, 106, 106,
	106, 106, 310, 9, -60, -1000, -1000, -1000, 52, 139,
	-1000, -1000, -1000, -1000, 70, -1000, -1000, 169, 106, 106,
	195, -1000, 383, -32, 412, -1000, -1000, -1000, -1000, -1000,
	-1000, 101, -1000, 106, 106, 106, 106, 106, 106, 106,
	-60, -1000, 22, 58, -1000, 101, 581, 581, 581, 581,
	581, 568, 555, 29, -32, -1000, -1000, -1000, -1000, 95,
	324, 108, 21, 26, 55, 23, 98, 106, 315, -60,
	581, 581, 268, -1000, 106, 15, -17, 106, 106, 106,
	106, 106, 106, 106, 106, 106, 106, 106, 106, 106,
	-1000, -1000, -1000, -1000, 180, 180, 483, 79, 30, 324,
	-1000, 310, 310, -1000, 13, 106, -1000, 581, -1000, -1000,
	71, 106, 106, -1000, 106, -1000, 106, 106, -1000, 106,
	106, 526, 8, -1000, -1000, 70, 97, -1000, -16, -1000,
	-32, 324, 144, 180, 180, 138, 138, 138, -1000, -1000,
	-1000, -1000, 425, 425, 594, 396, -1000, 106, 161, -11,
	-1000, -32, -1000, 14, 16, -18, 581, 106, -14, 581,
	581, 581, 581, 581, 505, -1000, 315, 236, 268, -1000,
	-50, 310, -19, 310, 158, 268, -1000, 310, -1000, -1000,
	106, 437, -1000, -1000, 54, -1000, -32, -1000, 106, 12,
	-1000, 25, 310, -1000, 6, 581, 106, -1000, 61, 378,
	-1000, -1000, 5, -1000, 461, -17, -1000, -1000, 142, 310,
	2, -1000,
}
var parserPgo = []int{

	0, 234, 233, 227, 5, 226, 161, 2, 7, 199,
	0, 4, 225, 1, 216, 3, 33, 213, 26, 212,
	208, 24, 117, 8, 207, 6, 23, 68,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 7, 7, 7, 27, 27, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 17, 11, 11, 11, 18, 19, 19,
	20, 16, 16, 24, 25, 25, 25, 22, 22, 22,
	21, 21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 3, 1, 12, 1, 3, 1, 4, 1,
	2, 2, 2, 2, 2, 2, 7, 7, 5, 3,
	2, 2, 2, 2, 5, 3, 4, 4, 4, 4,
	4, 3, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 3, 1, 0, 2, 4, 3,
	1, 2, 2, 6, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	67, -22, 13, -26, 14, 15, 16, 17, 18, 19,
	20, 24, 4, 2, 28, -21, -23, -24, -9, -27,
	31, 32, 33, 34, 35, -18, -16, 69, 29, 30,
	-27, 5, 65, -22, -10, 70, 56, 57, 71, 68,
	-20, -27, -18, 36, 38, 39, 40, 45, 44, 60,
	28, 58, -27, 59, -9, -27, -10, -10, -10, -10,
	-10, -10, -10, -7, -22, 65, 5, 23, 26, -27,
	63, 44, 45, 42, 41, 43, -19, 59, 60, 10,
	-10, -10, 7, -6, 67, -27, 65, 44, 45, 41,
	43, 42, 50, 48, 51, 49, 52, 53, 46, 47,
	-10, -10, -10, -10, -10, -10, -10, -27, 60, 63,
	61, 21, 25, 5, -27, 60, -8, -10, -17, -16,
	27, 59, 63, 44, 63, 45, 63, 63, 41, 63,
	59, -10, -25, -21, -23, 35, -22, -27, -14, -15,
	-22, 63, -26, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, 62, 60, 62, -12,
	-13, -22, -8, -7, -7, -11, -10, 60, -11, -10,
	-10, -10, -10, -10, -10, 61, 64, -4, 64, -27,
	-8, 25, -11, 12, 62, 64, -27, 22, 26, 62,
	64, -10, 61, 61, -25, 8, -22, -15, 65, -7,
	62, -7, 12, -13, -7, -10, 64, 62, -27, -10,
	26, 5, -7, 23, -10, 65, 5, 62, -26, 25,
	-7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 23,
	0, 0, 27, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 97, 98, 99, 0, 15,
	100, 101, 102, 103, 0, 16, 17, 25, 0, 0,
	0, 1, 0, 0, 0, 54, 55, 56, 57, 58,
	59, 60, 61, 0, 0, 0, 0, 0, 0, 0,
	0, 90, 0, 0, 30, 15, 31, 32, 33, 34,
	35, 0, 0, 0, 0, 40, 41, 42, 43, 0,
	0, 0, 0, 0, 0, 0, 87, 0, 0, 0,
	91, 92, 0, 22, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 63, 64, 65, 66, 67, 0, 0, 0, 0,
	104, 0, 0, 39, 0, 86, 45, 18, 19, 20,
	0, 86, 0, 51, 0, 52, 0, 0, 53, 0,
	0, 0, 0, -2, -2, 96, 0, 26, 9, 6,
	0, 0, 0, 68, 69, 70, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 86, 0, 0,
	13, 0, 28, 0, 0, 0, 85, 0, 0, 46,
	47, 48, 49, 50, 0, 89, 0, 0, 0, 7,
	0, 0, 0, 0, 0, 0, 14, 0, 38, 44,
	0, 0, 83, 88, 0, 4, 0, 5, 0, 0,
	82, 0, 0, 12, 0, 84, 0, 93, 0, 0,
	37, 10, 0, 36, 0, 0, 11, 21, 0, 0,
	0, 24,
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
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:162
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 23:
		//line waccparser.y:163
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 24:
		//line waccparser.y:164
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			w := While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			d := Declare{DecType: parserS[parserpt-10].typedefinition, Lhs: parserS[parserpt-9].ident, Rhs: parserS[parserpt-7].assignrhs, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			parserVAL.stmts = []Statement{d, w}
		}
	case 25:
		//line waccparser.y:171
		{
			parserVAL.ident = parserS[parserpt-0].ident
		}
	case 26:
		//line waccparser.y:172
		{
			parserVAL.ident = Ident(string(parserS[parserpt-2].ident) + "." + string(parserS[parserpt-0].ident))
		}
	case 27:
		//line waccparser.y:174
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 28:
		//line waccparser.y:175
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 29:
		//line waccparser.y:176
		{
			parserVAL.stmt = parserS[parserpt-0].stmt
		}
	case 30:
		//line waccparser.y:177
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 31:
		//line waccparser.y:178
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 32:
		//line waccparser.y:179
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 33:
		//line waccparser.y:180
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 34:
		//line waccparser.y:181
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 35:
		//line waccparser.y:182
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 36:
		//line waccparser.y:183
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 37:
		//line waccparser.y:184
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			parserVAL.stmt = While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 38:
		//line waccparser.y:188
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 39:
		//line waccparser.y:189
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 40:
		//line waccparser.y:190
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 41:
		//line waccparser.y:194
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 42:
		//line waccparser.y:197
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 43:
		//line waccparser.y:201
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 44:
		//line waccparser.y:205
		{
			parserVAL.stmt = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:209
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		//line waccparser.y:210
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		//line waccparser.y:211
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		//line waccparser.y:212
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:213
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:216
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:217
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 54:
		//line waccparser.y:219
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 55:
		//line waccparser.y:220
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 56:
		//line waccparser.y:221
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 57:
		//line waccparser.y:222
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 58:
		//line waccparser.y:223
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 59:
		//line waccparser.y:224
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 60:
		//line waccparser.y:225
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 61:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 62:
		//line waccparser.y:227
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:228
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		//line waccparser.y:229
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:230
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:231
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:232
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 68:
		//line waccparser.y:233
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:234
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:235
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:236
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		//line waccparser.y:244
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		//line waccparser.y:245
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 81:
		//line waccparser.y:246
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 82:
		//line waccparser.y:247
		{
			parserVAL.expr = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 83:
		//line waccparser.y:250
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 84:
		//line waccparser.y:252
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 85:
		//line waccparser.y:253
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 86:
		//line waccparser.y:254
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 87:
		//line waccparser.y:256
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 88:
		//line waccparser.y:258
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 89:
		//line waccparser.y:259
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 90:
		//line waccparser.y:262
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 91:
		//line waccparser.y:265
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 92:
		//line waccparser.y:266
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 93:
		//line waccparser.y:268
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 94:
		//line waccparser.y:270
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 95:
		//line waccparser.y:271
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 96:
		//line waccparser.y:272
		{
			parserVAL.pairelemtype = Pair
		}
	case 97:
		//line waccparser.y:274
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:275
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 99:
		//line waccparser.y:276
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 100:
		//line waccparser.y:278
		{
			parserVAL.typedefinition = Int
		}
	case 101:
		//line waccparser.y:279
		{
			parserVAL.typedefinition = Bool
		}
	case 102:
		//line waccparser.y:280
		{
			parserVAL.typedefinition = Char
		}
	case 103:
		//line waccparser.y:281
		{
			parserVAL.typedefinition = String
		}
	case 104:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
