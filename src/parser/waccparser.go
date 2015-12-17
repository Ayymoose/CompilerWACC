
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

//line waccparser.y:284


//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 47,
	69, 102,
	-2, 25,
}

const parserNprod = 110
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 575

var parserAct = []int{

	200, 7, 199, 34, 118, 25, 148, 68, 20, 80,
	9, 15, 16, 17, 18, 19, 66, 52, 53, 235,
	243, 13, 237, 60, 79, 12, 156, 236, 11, 219,
	226, 217, 50, 51, 152, 24, 21, 72, 92, 93,
	94, 95, 96, 97, 98, 50, 99, 49, 54, 14,
	210, 64, 101, 113, 114, 62, 168, 71, 61, 22,
	70, 22, 229, 228, 221, 221, 10, 214, 22, 22,
	102, 104, 224, 103, 22, 221, 23, 165, 167, 22,
	4, 120, 139, 140, 141, 142, 143, 144, 145, 33,
	22, 123, 124, 169, 22, 164, 22, 67, 220, 172,
	221, 173, 163, 121, 166, 125, 115, 155, 157, 107,
	59, 58, 100, 146, 57, 203, 157, 195, 241, 193,
	153, 162, 119, 29, 55, 56, 157, 179, 180, 181,
	182, 183, 184, 185, 186, 187, 188, 189, 190, 191,
	177, 159, 170, 64, 174, 175, 122, 62, 154, 159,
	61, 196, 178, 198, 197, 201, 176, 28, 27, 159,
	50, 26, 218, 205, 204, 206, 147, 207, 208, 48,
	209, 212, 194, 211, 106, 49, 171, 56, 213, 15,
	16, 17, 18, 19, 30, 15, 16, 17, 18, 19,
	67, 6, 128, 130, 129, 2, 215, 78, 216, 132,
	134, 131, 133, 105, 223, 69, 90, 158, 89, 65,
	117, 225, 31, 8, 227, 5, 116, 14, 3, 230,
	231, 1, 232, 14, 160, 88, 52, 53, 132, 134,
	131, 133, 0, 81, 238, 82, 83, 84, 239, 0,
	119, 86, 85, 242, 50, 111, 110, 112, 108, 109,
	0, 0, 0, 74, 75, 91, 161, 87, 45, 49,
	44, 0, 0, 0, 0, 77, 71, 73, 76, 32,
	35, 36, 37, 38, 39, 40, 41, 0, 0, 90,
	43, 89, 0, 0, 46, 52, 53, 15, 16, 17,
	18, 19, 0, 15, 16, 17, 18, 19, 88, 15,
	16, 17, 18, 63, 0, 0, 81, 0, 82, 83,
	84, 0, 0, 0, 86, 85, 0, 0, 0, 0,
	0, 0, 0, 42, 0, 47, 74, 75, 91, 9,
	87, 14, 0, 0, 0, 0, 0, 0, 77, 71,
	73, 76, 128, 130, 129, 126, 127, 137, 138, 132,
	134, 131, 133, 135, 136, 128, 130, 129, 126, 127,
	0, 0, 132, 134, 131, 133, 234, 128, 130, 129,
	126, 127, 137, 138, 132, 134, 131, 133, 135, 136,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 150, 128, 130, 129, 126, 127, 137, 138, 132,
	134, 131, 133, 135, 136, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 233, 128, 130, 129, 126,
	127, 137, 138, 132, 134, 131, 133, 135, 136, 0,
	0, 0, 0, 0, 0, 0, 0, 240, 128, 130,
	129, 126, 127, 137, 138, 132, 134, 131, 133, 135,
	136, 0, 0, 0, 0, 0, 0, 0, 0, 192,
	128, 130, 129, 126, 127, 137, 138, 132, 134, 131,
	133, 135, 136, 0, 0, 0, 0, 0, 0, 0,
	222, 128, 130, 129, 126, 127, 137, 138, 132, 134,
	131, 133, 135, 136, 151, 0, 0, 0, 0, 0,
	0, 202, 0, 149, 0, 0, 0, 0, 0, 0,
	128, 130, 129, 126, 127, 137, 138, 132, 134, 131,
	133, 135, 136, 128, 130, 129, 126, 127, 137, 138,
	132, 134, 131, 133, 135, 136, 128, 130, 129, 126,
	127, 137, 138, 132, 134, 131, 133, 135, 136, 128,
	130, 129, 126, 127, 137, 0, 132, 134, 131, 133,
	135, 136, 128, 130, 129, 126, 127, 0, 0, 132,
	134, 131, 133, 135, 136,
}
var parserPact = []int{

	191, -1000, -1000, 185, 262, -1000, -61, 31, -1000, -20,
	-64, 102, 99, 98, -1000, -1000, -1000, -1000, -1000, 63,
	177, -1000, 256, -64, 64, 167, 53, 50, 49, 268,
	-20, -1000, -1000, -62, -1000, -12, 270, 270, 270, 270,
	270, 270, 270, 270, -57, 47, -64, 115, 46, 204,
	-1000, -1000, 270, 270, 43, 154, -64, -1000, -1000, -1000,
	39, 102, 99, 63, 98, 27, -1000, -64, 42, -1000,
	-1000, 115, 495, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 270, 270, 270, 270, 270, 270, 270, -64, 156,
	-63, -1000, 495, 495, 495, 495, 482, 326, 469, 29,
	-1000, -1000, -1000, -1000, 60, 89, 270, 197, 58, 32,
	41, 15, 30, 495, 495, 197, 164, 37, -1000, -64,
	-1000, 268, 148, -20, -1000, 197, 270, 270, 270, 270,
	270, 270, 270, 270, 270, 270, 270, 270, 270, -1000,
	-1000, -1000, -1000, 151, 151, 397, 59, -64, 57, -57,
	-12, -57, -1000, 270, 270, 440, -1000, 495, -1000, -1000,
	55, 270, 270, -1000, 270, -1000, 270, 270, -1000, 270,
	-15, -57, 159, -20, -1000, 5, -1000, -1000, -1000, 151,
	151, 180, 180, 180, -1000, -1000, -1000, -1000, 314, 314,
	521, 508, -1000, 270, -1000, 270, 9, 137, 3, 36,
	495, 419, -1000, 270, 11, 495, 495, 495, 495, 495,
	270, 25, -57, -1000, -1000, 1, 0, -57, -57, -1000,
	-1000, 270, -1000, 351, -1000, 301, -1000, 14, -1000, -1000,
	4, -4, 495, 270, -12, -1000, -1000, -1000, 375, 93,
	-1000, -57, -6, -1000,
}
var parserPgo = []int{

	0, 221, 218, 215, 80, 213, 212, 1, 26, 169,
	0, 2, 210, 4, 209, 16, 33, 207, 9, 203,
	197, 28, 66, 25, 21, 23, 3, 24,
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
	19, 19, 20, 16, 16, 24, 25, 25, 25, 22,
	22, 22, 22, 21, 21, 21, 21, 23, 23, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 3, 0, 12, 1, 3, 1, 4, 1,
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

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, 67,
	-22, -21, -23, -24, 69, 31, 32, 33, 34, 35,
	69, 5, 65, -22, -27, 69, 59, 59, 59, 60,
	7, -6, 13, -22, -26, 14, 15, 16, 17, 18,
	19, 20, 67, 24, 4, 2, 28, 69, -9, -27,
	-18, -16, 29, 30, -27, 60, 10, 61, 61, 61,
	-25, -21, -23, 35, -24, -14, -15, -22, 69, -9,
	-27, 69, -10, 70, 56, 57, 71, 68, -20, -27,
	-18, 36, 38, 39, 40, 45, 44, 60, 28, 11,
	9, 58, -10, -10, -10, -10, -10, -10, -10, -7,
	65, 5, 23, 26, -27, -19, 59, 63, 44, 45,
	42, 41, 43, -10, -10, 63, 62, -12, -13, -22,
	-27, 64, -4, 64, -27, 63, 44, 45, 41, 43,
	42, 50, 48, 51, 49, 52, 53, 46, 47, -10,
	-10, -10, -10, -10, -10, -10, -27, 10, 69, 21,
	65, 25, 5, 60, 59, -10, -8, -10, -17, -16,
	27, 59, 63, 44, 63, 45, 63, 63, 41, 63,
	-8, 12, 62, 64, -27, -25, 8, -15, -8, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, 62, 60, -27, 60, -7, -26, -7, -11,
	-10, -10, 61, 60, -11, -10, -10, -10, -10, -10,
	65, -7, 12, -13, 62, -11, -11, 22, 25, 26,
	62, 64, 61, -10, 61, -10, 5, -7, 62, 62,
	-7, -7, -10, 64, 65, 5, 23, 26, -10, -26,
	62, 25, -7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, 23, 2, 0, 0, 8, 0,
	0, 99, 100, 101, 102, 103, 104, 105, 106, 0,
	0, 1, 0, 0, 0, 25, 0, 0, 0, 0,
	0, 22, 27, 0, 29, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 0, 0, -2, 0, 15,
	16, 17, 0, 0, 0, 0, 0, 107, 109, 108,
	0, 96, 97, 98, 0, 9, 6, 0, 0, 30,
	15, 25, 31, 54, 55, 56, 57, 58, 59, 60,
	61, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 92, 32, 33, 34, 35, 0, 0, 0, 0,
	40, 41, 42, 43, 0, 89, 0, 0, 0, 0,
	0, 0, 0, 93, 94, 0, 0, 0, 13, 0,
	26, 0, 0, 0, 7, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	63, 64, 65, 66, 67, 0, 0, 0, 0, 23,
	0, 23, 39, 88, 0, 0, 45, 18, 19, 20,
	0, 88, 0, 51, 0, 52, 0, 0, 53, 0,
	0, 23, 0, 0, 14, 0, 4, 5, 28, 68,
	69, 70, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 80, 81, 88, 83, 88, 0, 0, 0, 0,
	87, 0, 91, 0, 0, 46, 47, 48, 49, 50,
	0, 0, 23, 12, 95, 0, 0, 23, 23, 38,
	44, 0, 90, 0, 85, 0, 10, 0, 82, 84,
	0, 0, 86, 0, 0, 11, 36, 37, 0, 0,
	21, 23, 0, 24,
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
		//line waccparser.y:108
		{
	                                         parserlex.(*Lexer).prog = &Program{ClassList : parserS[parserpt-3].classes , FunctionList : parserS[parserpt-2].functions , StatList : parserS[parserpt-1].stmts , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
	                                         }
	case 2:
		//line waccparser.y:113
		{ parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)}
	case 3:
		//line waccparser.y:114
		{ parserVAL.classes = []*Class{} }
	case 4:
		//line waccparser.y:117
		{ if !checkClassIdent(parserS[parserpt-4].ident) {
	                                                         	parserlex.Error("Invalid class name")
	                                                     }
	                                                     parserVAL.class = &Class{Ident : ClassType(parserS[parserpt-4].ident), FieldList : parserS[parserpt-2].fields , FunctionList : parserS[parserpt-1].functions}
	                                                   }
	case 5:
		//line waccparser.y:123
		{ parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)}
	case 6:
		//line waccparser.y:124
		{ parserVAL.fields = []Field{parserS[parserpt-0].field}   }
	case 7:
		//line waccparser.y:127
		{ parserVAL.field = Field{FieldType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 8:
		//line waccparser.y:129
		{ parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)}
	case 9:
		//line waccparser.y:130
		{ parserVAL.functions = []*Function{} }
	case 10:
		//line waccparser.y:133
		{ if !checkStats(parserS[parserpt-1].stmts) {
	          	parserlex.Error("Missing return statement")
	           }
	             parserVAL.function = &Function{Ident : parserS[parserpt-5].ident, ReturnType : parserS[parserpt-6].typedefinition, StatList : parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 11:
		//line waccparser.y:139
		{ if !checkStats(parserS[parserpt-1].stmts) {
	            	parserlex.Error("Missing return statement")
	            }
	             parserVAL.function = &Function{Ident : parserS[parserpt-6].ident, ReturnType : parserS[parserpt-7].typedefinition, StatList : parserS[parserpt-1].stmts, ParameterTypes : parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 12:
		//line waccparser.y:145
		{ parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)}
	case 13:
		//line waccparser.y:146
		{ parserVAL.params = []Param{ parserS[parserpt-0].param } }
	case 14:
		//line waccparser.y:149
		{ parserVAL.param = Param{ParamType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 15:
		//line waccparser.y:150
		{parserVAL.assignlhs = parserS[parserpt-0].ident}
	case 16:
		//line waccparser.y:151
		{parserVAL.assignlhs = parserS[parserpt-0].arrayelem}
	case 17:
		//line waccparser.y:152
		{parserVAL.assignlhs = parserS[parserpt-0].pairelem}
	case 18:
		//line waccparser.y:154
		{ parserVAL.assignrhs = parserS[parserpt-0].expr }
	case 19:
		//line waccparser.y:155
		{ parserVAL.assignrhs = parserS[parserpt-0].arrayliter }
	case 20:
		//line waccparser.y:156
		{ parserVAL.assignrhs = parserS[parserpt-0].pairelem }
	case 21:
		//line waccparser.y:157
		{ parserVAL.assignrhs = NewPair{FstExpr : parserS[parserpt-3].expr, SndExpr : parserS[parserpt-1].expr, Pos : parserS[parserpt-5].pos, FileText :&parserlex.(*Lexer).input } }
	case 22:
		//line waccparser.y:159
		{ parserVAL.stmts = append(parserS[parserpt-2].stmts,parserS[parserpt-0].stmt) }
	case 23:
		//line waccparser.y:160
		{ parserVAL.stmts = []Statement{} }
	case 24:
		//line waccparser.y:161
		{
	                                                                                                            stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                                                                            w := While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-11].pos, FileText :&parserlex.(*Lexer).input}
	                                                                                                            d := Declare{DecType : parserS[parserpt-10].typedefinition, Lhs : parserS[parserpt-9].ident, Rhs : parserS[parserpt-7].assignrhs, Pos : parserS[parserpt-11].pos ,FileText :&parserlex.(*Lexer).input }
	                                                                                                            parserVAL.stmts = []Statement{d,w}
	                                                                                                           }
	case 25:
		//line waccparser.y:168
		{ parserVAL.ident = parserS[parserpt-0].ident}
	case 26:
		//line waccparser.y:169
		{ parserVAL.ident = Ident(string(parserS[parserpt-2].ident)+"."+string(parserS[parserpt-0].ident))}
	case 27:
		//line waccparser.y:171
		{ parserVAL.stmt = Skip{Pos : parserS[parserpt-0].pos ,FileText :&parserlex.(*Lexer).input } }
	case 28:
		//line waccparser.y:172
		{ parserVAL.stmt = Declare{DecType : parserS[parserpt-3].typedefinition, Lhs : parserS[parserpt-2].ident, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input } }
	case 29:
		//line waccparser.y:173
		{ parserVAL.stmt = parserS[parserpt-0].stmt }
	case 30:
		//line waccparser.y:174
		{ parserVAL.stmt = Read{ &parserlex.(*Lexer).input, parserS[parserpt-1].pos , parserS[parserpt-0].assignlhs, } }
	case 31:
		//line waccparser.y:175
		{ parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 32:
		//line waccparser.y:176
		{ parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 33:
		//line waccparser.y:177
		{ parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 34:
		//line waccparser.y:178
		{ parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 35:
		//line waccparser.y:179
		{ parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 36:
		//line waccparser.y:180
		{ parserVAL.stmt = If{Conditional : parserS[parserpt-5].expr, ThenStat : parserS[parserpt-3].stmts, ElseStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input } }
	case 37:
		//line waccparser.y:181
		{
	                                                              stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                              parserVAL.stmt = While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input}
	                                                             }
	case 38:
		//line waccparser.y:185
		{ parserVAL.stmt = While{Conditional : parserS[parserpt-3].expr, DoStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input} }
	case 39:
		//line waccparser.y:186
		{ parserVAL.stmt = Scope{StatList : parserS[parserpt-1].stmts, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input } }
	case 40:
		//line waccparser.y:187
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 41:
		//line waccparser.y:191
		{ parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 42:
		//line waccparser.y:194
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 43:
		//line waccparser.y:198
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 44:
		//line waccparser.y:202
		{ parserVAL.stmt = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 45:
		//line waccparser.y:205
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].assignlhs, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 46:
		//line waccparser.y:206
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 47:
		//line waccparser.y:207
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 48:
		//line waccparser.y:208
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 49:
		//line waccparser.y:209
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 50:
		//line waccparser.y:209
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 51:
		//line waccparser.y:210
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 52:
		//line waccparser.y:211
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 53:
		//line waccparser.y:212
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : MUL,  Right : parserS[parserpt-2].ident,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 54:
		//line waccparser.y:214
		{ parserVAL.expr =  parserS[parserpt-0].integer }
	case 55:
		//line waccparser.y:215
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 56:
		//line waccparser.y:216
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 57:
		//line waccparser.y:217
		{ parserVAL.expr =  parserS[parserpt-0].character }
	case 58:
		//line waccparser.y:218
		{ parserVAL.expr =  parserS[parserpt-0].stringconst }
	case 59:
		//line waccparser.y:219
		{ parserVAL.expr =  parserS[parserpt-0].pairliter }
	case 60:
		//line waccparser.y:220
		{ parserVAL.expr =  parserS[parserpt-0].ident }
	case 61:
		//line waccparser.y:221
		{ parserVAL.expr =  parserS[parserpt-0].arrayelem }
	case 62:
		//line waccparser.y:222
		{ parserVAL.expr = Unop{Unary : NOT, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 63:
		//line waccparser.y:223
		{ parserVAL.expr = Unop{Unary : LEN, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 64:
		//line waccparser.y:224
		{ parserVAL.expr = Unop{Unary : ORD, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 65:
		//line waccparser.y:225
		{ parserVAL.expr = Unop{Unary : CHR, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 66:
		//line waccparser.y:226
		{ parserVAL.expr = Unop{Unary : SUB, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 67:
		//line waccparser.y:227
		{ parserVAL.expr = parserS[parserpt-0].expr }
	case 68:
		//line waccparser.y:228
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 69:
		//line waccparser.y:229
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : SUB,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 70:
		//line waccparser.y:230
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 71:
		//line waccparser.y:231
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 72:
		//line waccparser.y:232
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 73:
		//line waccparser.y:233
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 74:
		//line waccparser.y:234
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 75:
		//line waccparser.y:235
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 76:
		//line waccparser.y:236
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 77:
		//line waccparser.y:237
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : EQ,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 78:
		//line waccparser.y:238
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : NEQ,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 79:
		//line waccparser.y:239
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : AND,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 80:
		//line waccparser.y:240
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : OR,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 81:
		//line waccparser.y:241
		{ parserVAL.expr = parserS[parserpt-1].expr }
	case 82:
		//line waccparser.y:242
		{ parserVAL.expr = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 83:
		//line waccparser.y:243
		{ parserVAL.expr = ThisInstance{parserS[parserpt-0].ident} }
	case 84:
		//line waccparser.y:244
		{ parserVAL.expr = NewObject{Class : ClassType(parserS[parserpt-3].ident) ,Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input}}
	case 85:
		//line waccparser.y:246
		{ parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs } }
	case 86:
		//line waccparser.y:248
		{parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)}
	case 87:
		//line waccparser.y:249
		{parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}}
	case 88:
		//line waccparser.y:250
		{parserVAL.exprs = []Evaluation{}}
	case 89:
		//line waccparser.y:252
		{parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs : parserS[parserpt-0].exprs, Pos : parserS[parserpt-1].pos,FileText :&parserlex.(*Lexer).input  } }
	case 90:
		//line waccparser.y:254
		{parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)}
	case 91:
		//line waccparser.y:255
		{parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}}
	case 92:
		//line waccparser.y:258
		{ parserVAL.pairliter =  PairLiter{} }
	case 93:
		//line waccparser.y:261
		{ parserVAL.pairelem = PairElem{Fsnd: Fst, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 94:
		//line waccparser.y:262
		{ parserVAL.pairelem = PairElem{Fsnd: Snd, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 95:
		//line waccparser.y:264
		{ parserVAL.typedefinition = PairType{FstType : parserS[parserpt-3].pairelemtype, SndType : parserS[parserpt-1].pairelemtype} }
	case 96:
		//line waccparser.y:266
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 97:
		//line waccparser.y:267
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 98:
		//line waccparser.y:268
		{ parserVAL.pairelemtype = Pair}
	case 99:
		//line waccparser.y:270
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition          }
	case 100:
		//line waccparser.y:271
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition          }
	case 101:
		//line waccparser.y:272
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition          }
	case 102:
		//line waccparser.y:273
		{ parserVAL.typedefinition = ClassType(parserS[parserpt-0].ident)}
	case 103:
		//line waccparser.y:275
		{ parserVAL.typedefinition =  Int    }
	case 104:
		//line waccparser.y:276
		{ parserVAL.typedefinition =  Bool   }
	case 105:
		//line waccparser.y:277
		{ parserVAL.typedefinition =  Char   }
	case 106:
		//line waccparser.y:278
		{ parserVAL.typedefinition =  String }
	case 107:
		//line waccparser.y:280
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 108:
		//line waccparser.y:281
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 109:
		//line waccparser.y:282
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	}
	goto parserstack /* stack new state and value */
}
