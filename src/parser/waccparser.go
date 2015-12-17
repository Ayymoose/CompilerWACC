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

//line waccparser.y:285

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 140,
	58, 96,
	-2, 93,
	-1, 141,
	58, 97,
	-2, 94,
}

const parserNprod = 104
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 669

var parserAct = []int{

	124, 167, 7, 139, 146, 173, 4, 123, 25, 36,
	35, 43, 37, 38, 221, 203, 65, 66, 67, 68,
	69, 70, 71, 13, 24, 72, 51, 62, 75, 218,
	226, 34, 29, 30, 31, 32, 33, 36, 88, 89,
	196, 34, 215, 197, 194, 76, 184, 117, 77, 34,
	182, 193, 36, 108, 109, 110, 111, 112, 113, 114,
	216, 121, 205, 165, 197, 40, 135, 50, 34, 41,
	148, 41, 28, 41, 39, 190, 212, 191, 136, 61,
	73, 41, 64, 41, 133, 11, 138, 74, 134, 126,
	28, 42, 118, 43, 41, 141, 150, 151, 152, 153,
	154, 155, 156, 157, 158, 159, 160, 161, 162, 28,
	93, 140, 116, 132, 130, 117, 78, 116, 149, 41,
	41, 34, 170, 171, 41, 169, 86, 115, 126, 174,
	175, 131, 176, 129, 177, 178, 172, 179, 180, 164,
	62, 122, 137, 85, 224, 187, 34, 34, 101, 103,
	100, 102, 183, 9, 27, 144, 186, 207, 189, 126,
	90, 87, 28, 6, 2, 174, 26, 143, 49, 63,
	188, 147, 84, 195, 125, 145, 82, 81, 83, 79,
	80, 29, 30, 31, 32, 33, 199, 28, 28, 202,
	204, 141, 206, 208, 85, 91, 209, 168, 211, 29,
	30, 31, 32, 142, 214, 166, 8, 140, 5, 3,
	217, 219, 1, 0, 34, 185, 34, 200, 0, 0,
	34, 97, 99, 98, 95, 96, 106, 225, 101, 103,
	100, 102, 104, 105, 34, 0, 192, 0, 0, 29,
	30, 31, 32, 33, 223, 0, 0, 34, 0, 0,
	0, 34, 0, 0, 0, 28, 0, 28, 0, 0,
	0, 28, 0, 143, 201, 147, 0, 0, 0, 213,
	0, 0, 168, 59, 0, 28, 29, 30, 31, 32,
	33, 52, 0, 53, 54, 55, 0, 0, 28, 57,
	56, 0, 28, 0, 0, 0, 0, 0, 0, 0,
	0, 45, 46, 60, 0, 58, 0, 0, 127, 59,
	37, 38, 0, 48, 36, 44, 47, 52, 0, 53,
	54, 55, 0, 0, 0, 57, 56, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 45, 46, 60,
	128, 58, 0, 23, 0, 22, 0, 0, 0, 48,
	36, 44, 47, 12, 14, 15, 16, 17, 18, 19,
	20, 0, 0, 0, 21, 0, 23, 0, 22, 37,
	38, 29, 30, 31, 32, 33, 12, 14, 15, 16,
	17, 18, 19, 20, 0, 0, 0, 21, 0, 0,
	0, 0, 37, 38, 29, 30, 31, 32, 33, 0,
	0, 0, 0, 0, 0, 59, 0, 10, 0, 36,
	0, 0, 0, 52, 0, 53, 54, 55, 0, 0,
	0, 57, 56, 0, 0, 0, 0, 0, 0, 0,
	92, 0, 36, 45, 46, 60, 0, 58, 0, 0,
	0, 0, 0, 0, 0, 48, 36, 44, 47, 97,
	99, 98, 95, 96, 106, 107, 101, 103, 100, 102,
	104, 105, 97, 99, 98, 95, 96, 0, 0, 101,
	103, 100, 102, 220, 97, 99, 98, 95, 96, 106,
	107, 101, 103, 100, 102, 104, 105, 97, 99, 98,
	0, 0, 0, 0, 101, 103, 100, 102, 94, 97,
	99, 98, 95, 96, 106, 107, 101, 103, 100, 102,
	104, 105, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 210, 97, 99, 98, 95, 96, 106, 107,
	101, 103, 100, 102, 104, 105, 0, 0, 0, 0,
	0, 0, 0, 0, 222, 97, 99, 98, 95, 96,
	106, 107, 101, 103, 100, 102, 104, 105, 0, 0,
	0, 0, 0, 0, 0, 0, 163, 97, 99, 98,
	95, 96, 106, 107, 101, 103, 100, 102, 104, 105,
	0, 0, 0, 0, 0, 0, 0, 198, 97, 99,
	98, 95, 96, 106, 107, 101, 103, 100, 102, 104,
	105, 120, 0, 0, 0, 0, 0, 0, 181, 0,
	119, 0, 0, 0, 0, 0, 0, 97, 99, 98,
	95, 96, 106, 107, 101, 103, 100, 102, 104, 105,
	97, 99, 98, 95, 96, 106, 107, 101, 103, 100,
	102, 104, 105, 97, 99, 98, 95, 96, 106, 107,
	101, 103, 100, 102, 104, 105, 97, 99, 98, 95,
	96, 0, 0, 101, 103, 100, 102, 104, 105,
}
var parserPact = []int{

	160, -1000, -1000, 157, 341, -1000, -59, 60, -1000, -1000,
	246, -31, -1000, -1000, -16, 378, 378, 378, 378, 378,
	378, 378, 341, 23, -1000, -1000, -1000, 54, 136, -1000,
	-1000, -1000, -1000, 67, -1000, -1000, 152, 378, 378, 153,
	-1000, 364, -31, 434, -1000, -1000, -1000, -1000, -1000, -1000,
	85, -1000, 378, 378, 378, 378, 378, 378, 378, -59,
	-1000, 53, 32, -1000, 85, 603, 603, 603, 603, 603,
	590, 577, 56, -31, -1000, -1000, -1000, -1000, 282, 71,
	69, 22, 26, 16, 84, 378, 169, -59, 603, 603,
	151, -1000, 378, 8, -16, 378, 378, 378, 378, 378,
	378, 378, 378, 378, 378, 378, 378, 378, -1000, -1000,
	-1000, -1000, 447, 447, 505, 80, 2, 282, -1000, 341,
	341, -1000, -15, -1000, 603, -1000, -1000, 77, 378, 378,
	-1000, 378, -1000, 378, 378, -1000, 378, 378, 548, -13,
	-1000, -1000, 67, 82, -1000, -17, -1000, -31, 282, 121,
	447, 447, 101, 101, 101, -1000, -1000, -1000, -1000, 422,
	422, 616, 181, -1000, 378, 147, 14, -1000, -31, -1000,
	30, 19, 378, -20, 603, 603, 603, 603, 603, 603,
	527, -1000, 169, 209, 151, -1000, -49, 341, 1, 341,
	146, 151, -1000, 341, -1000, 459, -1000, 378, -1000, 15,
	-1000, -31, -1000, 378, 17, -1000, 55, 341, -1000, 7,
	378, 603, -1000, 58, 409, -1000, -1000, 9, -1000, 483,
	-16, -1000, -1000, 120, 341, 5, -1000,
}
var parserPgo = []int{

	0, 212, 209, 208, 6, 206, 153, 2, 7, 154,
	0, 5, 205, 1, 175, 4, 10, 174, 26, 172,
	168, 24, 80, 8, 166, 3, 23, 67,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 7, 7, 7, 27, 27, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 17, 11, 11, 11, 18, 19, 19, 20,
	16, 16, 24, 25, 25, 25, 22, 22, 22, 21,
	21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 3, 1, 12, 1, 3, 1, 4, 1,
	2, 2, 2, 2, 2, 2, 7, 7, 5, 3,
	2, 2, 2, 2, 3, 4, 4, 4, 4, 4,
	3, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 3, 1, 0, 2, 4, 3, 1,
	2, 2, 6, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	66, -22, 12, -26, 13, 14, 15, 16, 17, 18,
	19, 23, 4, 2, -21, -23, -24, -9, -27, 30,
	31, 32, 33, 34, -18, -16, 68, 28, 29, -27,
	5, 64, -22, -10, 69, 55, 56, 70, 67, -20,
	-27, -18, 35, 37, 38, 39, 44, 43, 59, 27,
	57, -27, 58, -9, -27, -10, -10, -10, -10, -10,
	-10, -10, -7, -22, 64, 5, 22, 25, 62, 43,
	44, 41, 40, 42, -19, 58, 59, 9, -10, -10,
	7, -6, 66, -27, 64, 43, 44, 40, 42, 41,
	49, 47, 50, 48, 51, 52, 45, 46, -10, -10,
	-10, -10, -10, -10, -10, -27, 59, 62, 60, 20,
	24, 5, -27, -8, -10, -17, -16, 26, 58, 62,
	43, 62, 44, 62, 62, 40, 62, 58, -10, -25,
	-21, -23, 34, -22, -27, -14, -15, -22, 62, -26,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, 61, 59, 61, -12, -13, -22, -8,
	-7, -7, 59, -11, -10, -10, -10, -10, -10, -10,
	-10, 60, 63, -4, 63, -27, -8, 24, -11, 11,
	61, 63, -27, 21, 25, -10, 60, 63, 60, -25,
	8, -22, -15, 64, -7, 61, -7, 11, -13, -7,
	63, -10, 61, -27, -10, 25, 5, -7, 22, -10,
	64, 5, 61, -26, 24, -7, 25,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 23,
	0, 0, 27, 29, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 97, 98, 0, 15, 99,
	100, 101, 102, 0, 16, 17, 25, 0, 0, 0,
	1, 0, 0, 0, 53, 54, 55, 56, 57, 58,
	59, 60, 0, 0, 0, 0, 0, 0, 0, 0,
	89, 0, 0, 30, 15, 31, 32, 33, 34, 35,
	0, 0, 0, 0, 40, 41, 42, 43, 0, 0,
	0, 0, 0, 0, 86, 0, 0, 0, 90, 91,
	0, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 61, 62,
	63, 64, 65, 66, 0, 0, 0, 0, 103, 0,
	0, 39, 0, 44, 18, 19, 20, 0, 85, 0,
	50, 0, 51, 0, 0, 52, 0, 0, 0, 0,
	-2, -2, 95, 0, 26, 9, 6, 0, 0, 0,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	77, 78, 79, 80, 85, 0, 0, 13, 0, 28,
	0, 0, 0, 0, 84, 45, 46, 47, 48, 49,
	0, 88, 0, 0, 0, 7, 0, 0, 0, 0,
	0, 0, 14, 0, 38, 0, 82, 0, 87, 0,
	4, 0, 5, 0, 0, 81, 0, 0, 12, 0,
	0, 83, 92, 0, 0, 37, 10, 0, 36, 0,
	0, 11, 21, 0, 0, 0, 24,
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
		//line waccparser.y:129
		{
			parserVAL.field = Field{FieldType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 8:
		//line waccparser.y:131
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 9:
		//line waccparser.y:132
		{
			parserVAL.functions = []*Function{}
		}
	case 10:
		//line waccparser.y:135
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		//line waccparser.y:141
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 12:
		//line waccparser.y:147
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 13:
		//line waccparser.y:148
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 14:
		//line waccparser.y:150
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 15:
		//line waccparser.y:152
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 16:
		//line waccparser.y:153
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 17:
		//line waccparser.y:154
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 18:
		//line waccparser.y:156
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 19:
		//line waccparser.y:157
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 20:
		//line waccparser.y:158
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 21:
		//line waccparser.y:159
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:161
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 23:
		//line waccparser.y:162
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 24:
		//line waccparser.y:163
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			w := While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			d := Declare{DecType: parserS[parserpt-10].typedefinition, Lhs: parserS[parserpt-9].ident, Rhs: parserS[parserpt-7].assignrhs, Pos: parserS[parserpt-11].pos, FileText: &parserlex.(*Lexer).input}
			parserVAL.stmts = []Statement{d, w}
		}
	case 25:
		//line waccparser.y:170
		{
			parserVAL.ident = parserS[parserpt-0].ident
		}
	case 26:
		//line waccparser.y:171
		{
			parserVAL.ident = Ident(string(parserS[parserpt-2].ident) + "." + string(parserS[parserpt-0].ident))
		}
	case 27:
		//line waccparser.y:173
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 28:
		//line waccparser.y:174
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 29:
		//line waccparser.y:176
		{
			parserVAL.stmt = parserS[parserpt-0].stmt
		}
	case 30:
		//line waccparser.y:178
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 31:
		//line waccparser.y:179
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 32:
		//line waccparser.y:180
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 33:
		//line waccparser.y:181
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 34:
		//line waccparser.y:182
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 35:
		//line waccparser.y:183
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 36:
		//line waccparser.y:184
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 37:
		//line waccparser.y:187
		{
			stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
			parserVAL.stmt = While{Conditional: parserS[parserpt-5].expr, DoStat: stats, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 38:
		//line waccparser.y:191
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 39:
		//line waccparser.y:192
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 40:
		//line waccparser.y:193
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 41:
		//line waccparser.y:197
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 42:
		//line waccparser.y:200
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 43:
		//line waccparser.y:204
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 44:
		//line waccparser.y:209
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:210
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		//line waccparser.y:211
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		//line waccparser.y:212
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		//line waccparser.y:213
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:214
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:215
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:216
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:217
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:219
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 54:
		//line waccparser.y:220
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 55:
		//line waccparser.y:221
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 56:
		//line waccparser.y:222
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 57:
		//line waccparser.y:223
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 58:
		//line waccparser.y:224
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 59:
		//line waccparser.y:225
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 60:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 61:
		//line waccparser.y:227
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		//line waccparser.y:228
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:229
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		//line waccparser.y:230
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:231
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:232
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 67:
		//line waccparser.y:233
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:234
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:235
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:236
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:237
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:238
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:239
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:240
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:241
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:242
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:243
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 78:
		//line waccparser.y:244
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 79:
		//line waccparser.y:245
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 80:
		//line waccparser.y:246
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 81:
		//line waccparser.y:247
		{
			parserVAL.expr = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 82:
		//line waccparser.y:250
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 83:
		//line waccparser.y:252
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 84:
		//line waccparser.y:253
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 85:
		//line waccparser.y:254
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 86:
		//line waccparser.y:256
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 87:
		//line waccparser.y:258
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 88:
		//line waccparser.y:259
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 89:
		//line waccparser.y:262
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 90:
		//line waccparser.y:265
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 91:
		//line waccparser.y:266
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 92:
		//line waccparser.y:268
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 93:
		//line waccparser.y:270
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 94:
		//line waccparser.y:271
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 95:
		//line waccparser.y:272
		{
			parserVAL.pairelemtype = Pair
		}
	case 96:
		//line waccparser.y:274
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 97:
		//line waccparser.y:275
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 98:
		//line waccparser.y:276
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 99:
		//line waccparser.y:278
		{
			parserVAL.typedefinition = Int
		}
	case 100:
		//line waccparser.y:279
		{
			parserVAL.typedefinition = Bool
		}
	case 101:
		//line waccparser.y:280
		{
			parserVAL.typedefinition = Char
		}
	case 102:
		//line waccparser.y:281
		{
			parserVAL.typedefinition = String
		}
	case 103:
		//line waccparser.y:283
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
