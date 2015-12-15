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

//line waccparser.y:288

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 143,
	58, 99,
	-2, 96,
	-1, 144,
	58, 100,
	-2, 97,
}

const parserNprod = 107
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 637

var parserAct = []int{

	172, 168, 7, 148, 171, 4, 26, 224, 36, 123,
	142, 43, 207, 206, 13, 177, 173, 65, 66, 67,
	68, 69, 70, 71, 51, 25, 72, 37, 38, 35,
	37, 38, 61, 61, 61, 61, 62, 61, 87, 88,
	35, 73, 226, 197, 191, 122, 11, 92, 35, 39,
	61, 75, 42, 107, 108, 109, 110, 111, 112, 113,
	60, 237, 229, 115, 214, 150, 35, 29, 76, 208,
	64, 77, 199, 232, 244, 121, 201, 190, 228, 124,
	30, 31, 32, 33, 34, 188, 141, 126, 202, 241,
	133, 199, 43, 144, 139, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 151, 132,
	74, 166, 143, 41, 40, 41, 124, 41, 35, 41,
	41, 41, 174, 175, 126, 170, 136, 41, 146, 135,
	234, 149, 199, 181, 41, 182, 180, 183, 184, 114,
	185, 186, 115, 138, 35, 35, 52, 134, 53, 54,
	55, 124, 78, 189, 57, 56, 169, 225, 205, 126,
	192, 231, 198, 199, 199, 137, 45, 46, 59, 195,
	58, 196, 118, 41, 116, 114, 235, 203, 48, 50,
	44, 47, 200, 82, 81, 83, 79, 80, 86, 176,
	179, 61, 140, 85, 213, 144, 215, 216, 218, 210,
	219, 85, 9, 242, 221, 220, 193, 217, 204, 223,
	194, 28, 178, 89, 143, 227, 6, 2, 35, 35,
	230, 27, 117, 233, 211, 49, 35, 63, 84, 125,
	146, 212, 149, 100, 102, 99, 101, 129, 169, 85,
	239, 147, 35, 167, 90, 243, 30, 31, 32, 33,
	34, 240, 8, 127, 128, 37, 38, 5, 3, 1,
	0, 35, 52, 0, 53, 54, 55, 35, 0, 0,
	57, 56, 30, 31, 32, 33, 34, 30, 31, 32,
	33, 145, 45, 46, 59, 131, 58, 0, 24, 0,
	23, 0, 0, 0, 48, 130, 44, 47, 12, 15,
	16, 17, 18, 19, 20, 21, 0, 0, 0, 22,
	0, 0, 0, 14, 37, 38, 30, 31, 32, 33,
	34, 0, 0, 0, 0, 0, 0, 96, 98, 97,
	24, 0, 23, 0, 100, 102, 99, 101, 0, 0,
	12, 15, 16, 17, 18, 19, 20, 21, 0, 0,
	0, 22, 10, 0, 29, 14, 37, 38, 30, 31,
	32, 33, 34, 30, 31, 32, 33, 34, 52, 0,
	53, 54, 55, 0, 0, 0, 57, 56, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 45, 46,
	59, 0, 58, 0, 91, 0, 29, 0, 0, 0,
	48, 50, 44, 47, 96, 98, 97, 94, 95, 105,
	106, 100, 102, 99, 101, 103, 104, 96, 98, 97,
	94, 95, 0, 0, 100, 102, 99, 101, 236, 96,
	98, 97, 94, 95, 105, 106, 100, 102, 99, 101,
	103, 104, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 93, 96, 98, 97, 94, 95, 105,
	106, 100, 102, 99, 101, 103, 104, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 222, 96, 98,
	97, 94, 95, 105, 106, 100, 102, 99, 101, 103,
	104, 0, 0, 0, 0, 0, 0, 0, 0, 238,
	96, 98, 97, 94, 95, 105, 106, 100, 102, 99,
	101, 103, 104, 0, 0, 0, 0, 0, 0, 0,
	0, 165, 96, 98, 97, 94, 95, 105, 106, 100,
	102, 99, 101, 103, 104, 0, 0, 0, 0, 0,
	0, 0, 209, 96, 98, 97, 94, 95, 105, 106,
	100, 102, 99, 101, 103, 104, 120, 0, 0, 0,
	0, 0, 0, 187, 0, 119, 0, 0, 0, 0,
	0, 0, 96, 98, 97, 94, 95, 105, 106, 100,
	102, 99, 101, 103, 104, 96, 98, 97, 94, 95,
	105, 106, 100, 102, 99, 101, 103, 104, 96, 98,
	97, 94, 95, 105, 106, 100, 102, 99, 101, 103,
	104, 96, 98, 97, 94, 95, 105, 0, 100, 102,
	99, 101, 103, 104, 96, 98, 97, 94, 95, 0,
	0, 100, 102, 99, 101, 103, 104,
}
var parserPact = []int{

	213, -1000, -1000, 210, 286, -1000, -19, 109, -1000, -1000,
	333, -8, -1000, -1000, -32, 2, 111, 111, 111, 111,
	111, 111, 111, 286, 46, -1000, -1000, -1000, 90, 143,
	-1000, -1000, -1000, -1000, 129, -1000, -1000, 111, 111, 206,
	-1000, 328, -21, 389, -1000, -1000, -1000, -1000, -1000, -1000,
	135, -1000, 111, 111, 111, 111, 111, 111, 111, -1000,
	80, 114, 163, -1000, 135, 558, 558, 558, 558, 558,
	545, 532, 70, -23, -1000, -1000, -1000, -1000, 227, 47,
	85, 64, 103, 32, 134, 111, 247, 558, 558, 242,
	-1000, 111, 3, -1, 111, 111, 111, 111, 111, 111,
	111, 111, 111, 111, 111, 111, 111, -1000, -1000, -1000,
	-1000, 287, 287, 460, 50, 227, -1000, 111, -52, 286,
	286, -1000, 1, -1000, 558, -1000, -1000, 130, -53, 203,
	181, 111, 111, -1000, 111, -1000, 111, 111, -1000, 111,
	111, 503, 22, -1000, -1000, 129, 133, 14, -1000, -24,
	227, 182, 287, 287, 186, 186, 186, -1000, -1000, -1000,
	-1000, 377, 377, 584, 571, -1000, 199, 108, -1000, -25,
	-1000, 101, 558, 123, 55, 63, 111, 149, -55, -56,
	9, 558, 558, 558, 558, 558, 482, -1000, 247, 216,
	242, -1000, 0, 286, 286, 196, 242, -1000, -1000, 111,
	111, 286, -1000, 414, 111, -61, -1000, -1000, -1000, -1000,
	96, -1000, -26, -1000, 111, 53, 57, 286, -1000, 558,
	100, 51, 111, 69, 117, -1000, 116, 364, -1000, -1000,
	56, -1000, -1000, 438, -1000, 111, -1, -1000, -1000, 28,
	179, -1000, 286, 49, -1000,
}
var parserPgo = []int{

	0, 259, 258, 257, 5, 252, 202, 2, 9, 211,
	0, 4, 243, 1, 241, 3, 8, 229, 24, 228,
	225, 25, 41, 6, 221, 10, 14,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 8, 8, 8, 7, 7, 7, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 17, 11, 11, 11, 18,
	19, 19, 20, 16, 16, 24, 25, 25, 25, 22,
	22, 22, 21, 21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 7, 3, 3, 3, 1, 12, 1,
	4, 1, 5, 7, 2, 2, 2, 2, 2, 2,
	7, 7, 5, 3, 2, 2, 2, 2, 3, 4,
	4, 4, 4, 4, 3, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 1, 0, 2,
	4, 3, 1, 2, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	66, -22, 12, -26, 27, 13, 14, 15, 16, 17,
	18, 19, 23, 4, 2, -21, -23, -24, -9, 68,
	30, 31, 32, 33, 34, -18, -16, 28, 29, 68,
	5, 64, -22, -10, 69, 55, 56, 70, 67, -20,
	68, -18, 35, 37, 38, 39, 44, 43, 59, 57,
	68, 58, 68, -9, 68, -10, -10, -10, -10, -10,
	-10, -10, -7, -22, 64, 5, 22, 25, 62, 43,
	44, 41, 40, 42, -19, 58, 59, -10, -10, 7,
	-6, 66, 68, 64, 43, 44, 40, 42, 41, 49,
	47, 50, 48, 51, 52, 45, 46, -10, -10, -10,
	-10, -10, -10, -10, 59, 62, 60, 59, 9, 20,
	24, 5, 68, -8, -10, -17, -16, 26, 27, 10,
	68, 58, 62, 43, 62, 44, 62, 62, 40, 62,
	58, -10, -25, -21, -23, 34, -22, -14, -15, -22,
	62, -26, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, 61, 61, -12, -13, -22,
	-8, -11, -10, 68, -7, -7, 59, 68, 9, 9,
	-11, -10, -10, -10, -10, -10, -10, 60, 63, -4,
	63, 68, -8, 24, 11, 61, 63, 68, 61, 63,
	59, 21, 25, -10, 59, 9, 68, 68, 60, 60,
	-25, 8, -22, -15, 64, -7, -7, 11, -13, -10,
	-11, -7, 63, -11, 68, 61, 68, -10, 25, 5,
	-7, 61, 22, -10, 61, 59, 64, 5, 61, -11,
	-26, 61, 24, -7, 25,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 27,
	0, 0, 29, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 99, 100, 101, 0, 15,
	102, 103, 104, 105, 0, 16, 17, 0, 0, 0,
	1, 0, 0, 0, 57, 58, 59, 60, 61, 62,
	63, 64, 0, 0, 0, 0, 0, 0, 0, 92,
	0, 0, 0, 34, 15, 35, 36, 37, 38, 39,
	0, 0, 0, 0, 44, 45, 46, 47, 0, 0,
	0, 0, 0, 0, 89, 0, 0, 93, 94, 0,
	26, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 65, 66, 67,
	68, 69, 70, 0, 0, 0, 106, 88, 0, 0,
	0, 43, 0, 48, 18, 19, 20, 0, 0, 0,
	63, 88, 0, 54, 0, 55, 0, 0, 56, 0,
	0, 0, 0, -2, -2, 98, 0, 9, 6, 0,
	0, 0, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 82, 83, 84, 0, 0, 13, 0,
	30, 0, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 49, 50, 51, 52, 53, 0, 91, 0, 0,
	0, 7, 0, 0, 0, 0, 0, 14, 32, 0,
	88, 0, 42, 0, 88, 0, 24, 25, 85, 90,
	0, 4, 0, 5, 0, 0, 0, 0, 12, 86,
	0, 0, 0, 0, 0, 95, 0, 0, 41, 10,
	0, 33, 40, 0, 22, 88, 0, 11, 21, 0,
	0, 23, 0, 0, 28,
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
		//line waccparser.y:155
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 19:
		//line waccparser.y:156
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 20:
		//line waccparser.y:157
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 21:
		//line waccparser.y:158
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:159
		{
			parserVAL.assignrhs = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:160
		{
			parserVAL.assignrhs = CallInstance{Class: parserS[parserpt-5].ident, Func: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs}
		}
	case 24:
		//line waccparser.y:161
		{
			parserVAL.assignrhs = ThisInstance{Func: parserS[parserpt-0].ident}
		}
	case 25:
		//line waccparser.y:162
		{
			parserVAL.assignrhs = Instance{IdentLHS: parserS[parserpt-2].ident, IdentRHS: parserS[parserpt-0].ident}
		}
	case 26:
		//line waccparser.y:164
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 27:
		//line waccparser.y:165
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 28:
		//line waccparser.y:166
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			w := While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			d := Declare{DecType: parserS[parserpt-10].typedefinition, Lhs: parserS[parserpt-9].ident, Rhs: parserS[parserpt-7].assignrhs, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			parserVAL.stmts = []Statement{d, w}
		}
	case 29:
		//line waccparser.y:173
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 30:
		//line waccparser.y:174
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:176
		{
			parserVAL.stmt = parserS[parserpt-0].stmt
		}
	case 32:
		//line waccparser.y:178
		{
			parserVAL.stmt = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 33:
		//line waccparser.y:180
		{
			parserVAL.stmt = CallInstance{Class: parserS[parserpt-5].ident, Func: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs}
		}
	case 34:
		//line waccparser.y:182
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 35:
		//line waccparser.y:183
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 36:
		//line waccparser.y:184
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 37:
		//line waccparser.y:185
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 38:
		//line waccparser.y:186
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 39:
		//line waccparser.y:187
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 40:
		//line waccparser.y:188
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 41:
		//line waccparser.y:191
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			parserVAL.stmt = While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 42:
		//line waccparser.y:195
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 43:
		//line waccparser.y:196
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		//line waccparser.y:197
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 45:
		//line waccparser.y:201
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 46:
		//line waccparser.y:204
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 47:
		//line waccparser.y:208
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 48:
		//line waccparser.y:213
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:216
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:217
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:218
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 54:
		//line waccparser.y:219
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 55:
		//line waccparser.y:220
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 56:
		//line waccparser.y:221
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 57:
		//line waccparser.y:223
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 58:
		//line waccparser.y:224
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 59:
		//line waccparser.y:225
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 60:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 61:
		//line waccparser.y:227
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 62:
		//line waccparser.y:228
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 63:
		//line waccparser.y:229
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 64:
		//line waccparser.y:230
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 65:
		//line waccparser.y:231
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:232
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:233
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:234
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:235
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:236
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 71:
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		//line waccparser.y:244
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		//line waccparser.y:245
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		//line waccparser.y:246
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 81:
		//line waccparser.y:247
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 82:
		//line waccparser.y:248
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 83:
		//line waccparser.y:249
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 84:
		//line waccparser.y:250
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 85:
		//line waccparser.y:253
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 86:
		//line waccparser.y:255
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 87:
		//line waccparser.y:256
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 88:
		//line waccparser.y:257
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 89:
		//line waccparser.y:259
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 90:
		//line waccparser.y:261
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 91:
		//line waccparser.y:262
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 92:
		//line waccparser.y:265
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 93:
		//line waccparser.y:268
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 94:
		//line waccparser.y:269
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 95:
		//line waccparser.y:271
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 96:
		//line waccparser.y:273
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 97:
		//line waccparser.y:274
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:275
		{
			parserVAL.pairelemtype = Pair
		}
	case 99:
		//line waccparser.y:277
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
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
		//line waccparser.y:281
		{
			parserVAL.typedefinition = Int
		}
	case 103:
		//line waccparser.y:282
		{
			parserVAL.typedefinition = Bool
		}
	case 104:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = Char
		}
	case 105:
		//line waccparser.y:284
		{
			parserVAL.typedefinition = String
		}
	case 106:
		//line waccparser.y:286
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
