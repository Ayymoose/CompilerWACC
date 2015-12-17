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

//line waccparser.y:290

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 110
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 689

var parserAct = []int{

	187, 7, 181, 27, 186, 158, 36, 134, 26, 37,
	150, 44, 4, 123, 13, 220, 67, 68, 69, 70,
	71, 72, 73, 244, 74, 75, 52, 233, 25, 239,
	11, 35, 38, 39, 236, 234, 43, 128, 199, 94,
	95, 35, 197, 30, 31, 32, 33, 153, 141, 35,
	209, 223, 230, 212, 114, 115, 116, 117, 118, 119,
	120, 214, 42, 41, 212, 77, 42, 140, 222, 35,
	212, 211, 37, 212, 133, 51, 42, 210, 125, 160,
	29, 154, 40, 78, 147, 135, 79, 64, 144, 42,
	66, 137, 149, 42, 84, 42, 155, 42, 29, 44,
	80, 152, 162, 163, 164, 165, 166, 167, 168, 169,
	170, 171, 172, 173, 174, 161, 42, 146, 29, 99,
	143, 151, 159, 42, 132, 76, 135, 35, 184, 185,
	131, 124, 137, 183, 125, 206, 121, 207, 142, 145,
	124, 190, 92, 191, 189, 192, 193, 188, 194, 195,
	182, 129, 178, 35, 35, 30, 31, 32, 33, 34,
	176, 135, 88, 87, 89, 85, 86, 137, 201, 156,
	198, 130, 83, 82, 81, 148, 29, 91, 242, 9,
	91, 203, 202, 204, 28, 225, 179, 205, 122, 213,
	107, 109, 106, 108, 93, 96, 6, 2, 177, 65,
	50, 155, 29, 29, 221, 219, 152, 224, 216, 90,
	226, 227, 136, 228, 30, 31, 32, 33, 34, 157,
	180, 232, 97, 8, 218, 159, 151, 235, 5, 35,
	237, 3, 35, 182, 1, 200, 35, 0, 103, 105,
	104, 101, 102, 112, 243, 107, 109, 106, 108, 110,
	111, 0, 35, 241, 103, 105, 104, 0, 208, 0,
	0, 107, 109, 106, 108, 35, 0, 0, 0, 35,
	0, 0, 0, 62, 0, 61, 0, 0, 29, 0,
	0, 29, 0, 0, 0, 29, 0, 0, 0, 0,
	0, 0, 60, 0, 231, 30, 31, 32, 33, 34,
	53, 29, 54, 55, 56, 0, 0, 0, 58, 57,
	62, 0, 61, 0, 29, 0, 0, 0, 29, 0,
	46, 47, 63, 0, 59, 217, 0, 0, 138, 60,
	38, 39, 49, 37, 45, 48, 0, 53, 0, 54,
	55, 56, 0, 0, 0, 58, 57, 0, 30, 31,
	32, 33, 34, 0, 0, 0, 0, 46, 47, 63,
	139, 59, 23, 0, 22, 0, 0, 0, 0, 49,
	37, 45, 48, 12, 14, 15, 16, 17, 18, 19,
	20, 0, 0, 0, 21, 23, 0, 22, 24, 38,
	39, 30, 31, 32, 33, 34, 12, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 62, 21, 61, 0,
	0, 24, 38, 39, 30, 31, 32, 33, 34, 0,
	0, 0, 0, 0, 0, 60, 0, 10, 0, 37,
	0, 0, 0, 53, 0, 54, 55, 56, 0, 0,
	0, 58, 57, 0, 0, 0, 0, 0, 0, 0,
	98, 0, 37, 46, 47, 63, 0, 59, 0, 0,
	0, 0, 0, 0, 0, 49, 37, 45, 48, 103,
	105, 104, 101, 102, 112, 113, 107, 109, 106, 108,
	110, 111, 103, 105, 104, 101, 102, 0, 0, 107,
	109, 106, 108, 238, 103, 105, 104, 101, 102, 112,
	113, 107, 109, 106, 108, 110, 111, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 100, 103,
	105, 104, 101, 102, 112, 113, 107, 109, 106, 108,
	110, 111, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 229, 103, 105, 104, 101, 102, 112, 113,
	107, 109, 106, 108, 110, 111, 0, 0, 0, 0,
	0, 0, 0, 0, 240, 103, 105, 104, 101, 102,
	112, 113, 107, 109, 106, 108, 110, 111, 0, 0,
	0, 0, 0, 0, 0, 0, 175, 103, 105, 104,
	101, 102, 112, 113, 107, 109, 106, 108, 110, 111,
	0, 0, 0, 0, 0, 0, 0, 215, 103, 105,
	104, 101, 102, 112, 113, 107, 109, 106, 108, 110,
	111, 127, 0, 0, 0, 0, 0, 0, 196, 0,
	126, 0, 0, 0, 0, 0, 0, 103, 105, 104,
	101, 102, 112, 113, 107, 109, 106, 108, 110, 111,
	103, 105, 104, 101, 102, 112, 113, 107, 109, 106,
	108, 110, 111, 103, 105, 104, 101, 102, 112, 113,
	107, 109, 106, 108, 110, 111, 103, 105, 104, 101,
	102, 0, 0, 107, 109, 106, 108, 110, 111,
}
var parserPact = []int{

	193, -1000, -1000, 190, 360, -1000, -60, 58, -1000, -1000,
	264, -60, -1000, -1000, 3, 397, 397, 397, 397, 397,
	397, 397, 360, 60, -60, 115, 114, 113, 31, 121,
	-1000, -1000, -1000, -1000, 82, -1000, -1000, 184, 397, 397,
	188, -1000, 383, -60, 453, -1000, -1000, -1000, -1000, -1000,
	-1000, 118, -1000, 397, 397, 397, 397, 397, 397, 397,
	-60, 178, -56, -1000, 71, -1000, 118, 622, 622, 622,
	622, 622, 609, 596, 32, -60, -1000, -1000, -1000, -1000,
	111, 69, 63, 13, 301, 4, 75, 25, 76, 21,
	116, 397, 12, -60, 622, 622, 183, -1000, 397, 16,
	3, 397, 397, 397, 397, 397, 397, 397, 397, 397,
	397, 397, 397, 397, -1000, -1000, -1000, -1000, 213, 213,
	524, 100, -60, 92, 124, 301, 360, 360, -1000, 15,
	397, -1000, -1000, -1000, -1000, 622, -1000, -1000, 87, 397,
	397, -1000, 397, -1000, 397, 397, -1000, 397, 397, 567,
	-22, 115, 114, 82, -1000, 113, -1000, -26, -1000, -60,
	301, 157, 213, 213, 142, 142, 142, -1000, -1000, -1000,
	-1000, 441, 441, 635, 197, -1000, 397, -1000, 397, 175,
	73, -1000, -60, -1000, 28, 51, 9, 622, 397, 0,
	622, 622, 622, 622, 622, 546, -1000, 12, 317, 183,
	-1000, -50, 360, 6, -11, 360, 173, 183, -1000, 360,
	-1000, -1000, 397, 478, -1000, -1000, -10, -1000, -60, -1000,
	397, 1, -1000, -1000, 30, 360, -1000, 11, 622, 397,
	-1000, 80, 428, -1000, -1000, 24, -1000, 502, 3, -1000,
	-1000, 153, 360, -3, -1000,
}
var parserPgo = []int{

	0, 234, 231, 228, 12, 223, 179, 1, 7, 184,
	0, 4, 220, 2, 219, 5, 6, 212, 26, 209,
	200, 28, 25, 8, 3, 10, 14, 75,
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
	10, 10, 10, 10, 10, 17, 11, 11, 11, 18,
	19, 19, 20, 16, 16, 24, 25, 25, 25, 25,
	22, 22, 22, 21, 21, 21, 21, 23, 23, 23,
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
	3, 3, 5, 3, 5, 3, 3, 1, 0, 2,
	4, 3, 1, 2, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 3, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	67, -22, 13, -26, 14, 15, 16, 17, 18, 19,
	20, 24, 4, 2, 28, -21, -23, -24, -9, -27,
	31, 32, 33, 34, 35, -18, -16, 69, 29, 30,
	-27, 5, 65, -22, -10, 70, 56, 57, 71, 68,
	-20, -27, -18, 36, 38, 39, 40, 45, 44, 60,
	28, 11, 9, 58, -27, -9, -27, -10, -10, -10,
	-10, -10, -10, -10, -7, -22, 65, 5, 23, 26,
	-27, 59, 59, 59, 63, 44, 45, 42, 41, 43,
	-19, 59, 60, 10, -10, -10, 7, -6, 67, -27,
	65, 44, 45, 41, 43, 42, 50, 48, 51, 49,
	52, 53, 46, 47, -10, -10, -10, -10, -10, -10,
	-10, -27, 10, 69, 60, 63, 21, 25, 5, -27,
	60, 61, 61, 61, -8, -10, -17, -16, 27, 59,
	63, 44, 63, 45, 63, 63, 41, 63, 59, -10,
	-25, -21, -23, 35, 69, -24, -27, -14, -15, -22,
	63, -26, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, 62, 60, -27, 60, 62,
	-12, -13, -22, -8, -7, -7, -11, -10, 60, -11,
	-10, -10, -10, -10, -10, -10, 61, 64, -4, 64,
	-27, -8, 25, -11, -11, 12, 62, 64, -27, 22,
	26, 62, 64, -10, 61, 61, -25, 8, -22, -15,
	65, -7, 62, 62, -7, 12, -13, -7, -10, 64,
	62, -27, -10, 26, 5, -7, 23, -10, 65, 5,
	62, -26, 25, -7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 23,
	0, 0, 27, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 100, 101, 102, 0, 15,
	103, 104, 105, 106, 0, 16, 17, 25, 0, 0,
	0, 1, 0, 0, 0, 54, 55, 56, 57, 58,
	59, 60, 61, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 92, 0, 30, 15, 31, 32, 33,
	34, 35, 0, 0, 0, 0, 40, 41, 42, 43,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	89, 0, 0, 0, 93, 94, 0, 22, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 63, 64, 65, 66, 67,
	0, 0, 0, 0, 0, 0, 0, 0, 39, 0,
	88, 107, 109, 108, 45, 18, 19, 20, 0, 88,
	0, 51, 0, 52, 0, 0, 53, 0, 0, 0,
	0, 96, 97, 98, 99, 0, 26, 9, 6, 0,
	0, 0, 68, 69, 70, 71, 72, 73, 74, 75,
	76, 77, 78, 79, 80, 81, 88, 83, 88, 0,
	0, 13, 0, 28, 0, 0, 0, 87, 0, 0,
	46, 47, 48, 49, 50, 0, 91, 0, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 14, 0,
	38, 44, 0, 0, 85, 90, 0, 4, 0, 5,
	0, 0, 82, 84, 0, 0, 12, 0, 86, 0,
	95, 0, 0, 37, 10, 0, 36, 0, 0, 11,
	21, 0, 0, 0, 24,
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
		//line waccparser.y:248
		{
			parserVAL.expr = ThisInstance{parserS[parserpt-0].ident}
		}
	case 84:
		//line waccparser.y:249
		{
			parserVAL.expr = NewObject{Class: ClassType(parserS[parserpt-3].ident), Init: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 85:
		//line waccparser.y:252
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 86:
		//line waccparser.y:254
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 87:
		//line waccparser.y:255
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 88:
		//line waccparser.y:256
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 89:
		//line waccparser.y:258
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 90:
		//line waccparser.y:260
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 91:
		//line waccparser.y:261
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 92:
		//line waccparser.y:264
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 93:
		//line waccparser.y:267
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 94:
		//line waccparser.y:268
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 95:
		//line waccparser.y:270
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 96:
		//line waccparser.y:272
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 97:
		//line waccparser.y:273
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:274
		{
			parserVAL.pairelemtype = Pair
		}
	case 99:
		//line waccparser.y:275
		{
			parserVAL.pairelemtype = ClassType(parserS[parserpt-0].ident)
		}
	case 100:
		//line waccparser.y:277
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
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
		//line waccparser.y:281
		{
			parserVAL.typedefinition = Int
		}
	case 104:
		//line waccparser.y:282
		{
			parserVAL.typedefinition = Bool
		}
	case 105:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = Char
		}
	case 106:
		//line waccparser.y:284
		{
			parserVAL.typedefinition = String
		}
	case 107:
		//line waccparser.y:286
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	case 108:
		//line waccparser.y:287
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	case 109:
		//line waccparser.y:288
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
