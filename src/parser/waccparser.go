
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
fieldaccess Evaluation
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

//line waccparser.y:301


//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	69, 116,
	-2, 15,
	-1, 46,
	69, 116,
	-2, 71,
}

const parserNprod = 124
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 697

var parserAct = []int{

	210, 204, 179, 27, 37, 26, 171, 25, 7, 4,
	264, 45, 156, 238, 13, 209, 69, 71, 72, 73,
	74, 75, 76, 230, 52, 219, 201, 177, 145, 30,
	139, 77, 38, 131, 31, 32, 33, 34, 35, 68,
	104, 105, 31, 32, 33, 34, 35, 30, 54, 84,
	39, 40, 109, 36, 83, 65, 124, 125, 126, 127,
	128, 129, 130, 36, 41, 78, 278, 80, 30, 269,
	11, 36, 181, 202, 273, 252, 44, 267, 138, 229,
	181, 38, 42, 239, 270, 81, 242, 132, 82, 266,
	184, 255, 36, 242, 227, 247, 155, 157, 242, 39,
	40, 159, 240, 263, 134, 43, 175, 135, 173, 45,
	172, 43, 185, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 197, 183, 43, 79, 43, 31,
	32, 33, 34, 174, 43, 30, 157, 43, 43, 67,
	159, 43, 43, 135, 182, 207, 208, 212, 206, 213,
	170, 214, 215, 167, 216, 217, 254, 211, 242, 36,
	153, 30, 30, 166, 222, 63, 223, 176, 224, 225,
	150, 226, 180, 243, 241, 242, 242, 96, 221, 249,
	88, 165, 62, 157, 144, 36, 36, 159, 228, 236,
	55, 237, 56, 57, 58, 231, 169, 152, 60, 59,
	205, 164, 31, 32, 33, 34, 35, 143, 149, 147,
	48, 49, 64, 142, 61, 233, 234, 87, 168, 151,
	163, 246, 51, 70, 47, 50, 148, 134, 146, 245,
	199, 175, 251, 173, 248, 172, 102, 220, 200, 258,
	181, 253, 88, 260, 256, 141, 86, 88, 259, 85,
	154, 257, 276, 265, 9, 232, 235, 30, 133, 103,
	30, 261, 106, 271, 30, 29, 268, 113, 115, 114,
	111, 112, 122, 123, 117, 119, 116, 118, 120, 121,
	66, 36, 30, 88, 36, 277, 6, 275, 36, 2,
	53, 272, 140, 94, 250, 180, 95, 30, 107, 158,
	178, 30, 203, 205, 8, 5, 36, 3, 160, 1,
	63, 0, 0, 0, 92, 91, 93, 89, 90, 0,
	0, 36, 0, 0, 0, 36, 161, 62, 39, 40,
	0, 0, 95, 0, 0, 55, 0, 56, 57, 58,
	0, 0, 0, 60, 59, 117, 119, 116, 118, 100,
	99, 101, 97, 98, 63, 48, 49, 64, 162, 61,
	0, 0, 0, 0, 0, 0, 0, 51, 70, 47,
	50, 62, 0, 0, 31, 32, 33, 34, 35, 55,
	0, 56, 57, 58, 0, 0, 0, 60, 59, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	49, 64, 0, 61, 0, 0, 23, 0, 22, 0,
	0, 51, 46, 47, 50, 38, 0, 12, 14, 15,
	16, 17, 18, 19, 20, 0, 0, 0, 21, 0,
	0, 0, 24, 39, 40, 31, 32, 33, 34, 35,
	0, 0, 0, 0, 113, 115, 114, 0, 23, 0,
	22, 117, 119, 116, 118, 0, 0, 38, 0, 12,
	14, 15, 16, 17, 18, 19, 20, 0, 0, 0,
	21, 10, 0, 28, 24, 39, 40, 31, 32, 33,
	34, 35, 113, 115, 114, 111, 112, 122, 123, 117,
	119, 116, 118, 120, 121, 0, 0, 0, 0, 0,
	113, 115, 114, 111, 112, 122, 110, 117, 119, 116,
	118, 120, 121, 108, 0, 28, 113, 115, 114, 111,
	112, 122, 123, 117, 119, 116, 118, 120, 121, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 262,
	113, 115, 114, 111, 112, 122, 123, 117, 119, 116,
	118, 120, 121, 0, 0, 0, 0, 0, 0, 0,
	0, 274, 113, 115, 114, 111, 112, 122, 123, 117,
	119, 116, 118, 120, 121, 0, 0, 0, 0, 0,
	0, 0, 0, 198, 113, 115, 114, 111, 112, 122,
	123, 117, 119, 116, 118, 120, 121, 0, 0, 0,
	0, 0, 0, 0, 244, 113, 115, 114, 111, 112,
	122, 123, 117, 119, 116, 118, 120, 121, 137, 0,
	0, 0, 0, 0, 0, 218, 0, 136, 0, 0,
	0, 0, 0, 0, 113, 115, 114, 111, 112, 122,
	123, 117, 119, 116, 118, 120, 121, 113, 115, 114,
	111, 112, 122, 123, 117, 119, 116, 118, 120, 121,
	113, 115, 114, 111, 112, 122, 123, 117, 119, 116,
	118, 120, 121, 113, 115, 114, 111, 112, 0, 0,
	117, 119, 116, 118, 120, 121, 113, 115, 114, 111,
	112, 0, 0, 117, 119, 116, 118,
}
var parserPact = []int{

	285, -1000, -1000, 280, 404, -1000, -5, 77, -1000, -1000,
	343, -14, -1000, -1000, 70, 154, 154, 154, 154, 154,
	154, 154, 404, 62, -15, 190, 187, 158, 273, 114,
	308, -1000, -1000, -1000, -1000, 176, -1000, -1000, 249, 154,
	154, 255, -1000, 446, -17, 441, 237, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 154, 154, 154, 154, 154,
	154, 154, -36, 248, -1000, 44, -1000, 237, -1000, 619,
	237, 619, 619, 619, 619, 606, 593, 73, -39, -1000,
	-1000, -1000, -1000, 232, 185, 152, 146, 123, -41, 165,
	163, 107, 156, 97, 191, 154, 299, 157, 118, 90,
	155, 87, 98, -42, 619, 619, 3, -1000, 154, 81,
	21, 154, 154, 154, 154, 154, 154, 154, 154, 154,
	154, 154, 154, 154, -1000, -1000, -1000, -1000, 403, 403,
	521, 170, 178, -43, 11, 299, 404, 404, -1000, 80,
	154, 154, -1000, -1000, -1000, -1000, 154, -1000, 154, -1000,
	154, 154, -1000, 154, 154, 564, -1000, 619, -1000, -1000,
	-44, 177, 154, 154, -1000, 154, -1000, 154, 154, -1000,
	154, 30, 190, 187, 176, 158, -1000, -1000, 15, -1000,
	-46, -1000, 299, 230, 273, 403, 403, 297, 297, 297,
	-1000, -1000, -1000, -1000, 645, 645, 632, 459, -1000, 154,
	154, -1000, 244, 127, -1000, -56, -1000, 61, 76, 112,
	619, 111, 619, 619, 619, 619, 619, 543, -1000, 169,
	154, 34, 619, 619, 619, 619, 619, 98, 171, 3,
	-1000, 10, 404, 94, 29, 404, 239, 3, -1000, 404,
	-1000, -1000, 154, -1000, -1000, 154, 475, -1000, 41, -1000,
	-59, -1000, 154, 63, -1000, -1000, 72, 404, -1000, 46,
	619, 22, 154, -1000, 167, 226, -1000, -1000, 69, -1000,
	-1000, 499, 21, -1000, -1000, 227, 404, 40, -1000,
}
var parserPgo = []int{

	0, 309, 307, 305, 9, 304, 254, 8, 12, 265,
	0, 15, 302, 1, 300, 2, 4, 299, 48, 293,
	290, 7, 65, 5, 3, 6, 14, 24,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 9, 9,
	27, 8, 8, 8, 8, 8, 7, 7, 7, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 17, 11, 11,
	11, 18, 19, 19, 20, 16, 16, 24, 25, 25,
	25, 25, 25, 22, 22, 22, 22, 21, 21, 21,
	21, 23, 23, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 3,
	3, 1, 1, 1, 5, 6, 3, 1, 12, 1,
	4, 1, 2, 2, 2, 2, 2, 2, 7, 7,
	5, 3, 2, 2, 2, 2, 5, 5, 3, 4,
	4, 4, 4, 4, 3, 3, 3, 4, 4, 4,
	4, 4, 3, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 5, 3, 3, 3, 1,
	0, 2, 4, 3, 1, 2, 2, 6, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	67, -22, 13, -26, 14, 15, 16, 17, 18, 19,
	20, 24, 4, 2, 28, -21, -23, -24, 69, -9,
	-27, 31, 32, 33, 34, 35, -18, -16, 11, 29,
	30, 69, 5, 65, -22, -10, 69, 70, 56, 57,
	71, 68, -27, -20, -18, 36, 38, 39, 40, 45,
	44, 60, 28, 11, 58, 69, -9, 69, -27, -10,
	69, -10, -10, -10, -10, -10, -10, -7, -22, 65,
	5, 23, 26, 69, -27, 59, 59, 59, 10, 44,
	45, 42, 41, 43, -19, 59, 63, 44, 45, 42,
	41, 43, 60, 10, -10, -10, 7, -6, 67, 69,
	65, 44, 45, 41, 43, 42, 50, 48, 51, 49,
	52, 53, 46, 47, -10, -10, -10, -10, -10, -10,
	-10, 69, -27, 10, 60, 63, 21, 25, 5, 69,
	60, 60, 61, 61, 61, 69, 63, 44, 63, 45,
	63, 63, 41, 63, 59, -10, -8, -10, -17, -16,
	9, 27, 59, 63, 44, 63, 45, 63, 63, 41,
	63, -25, -21, -23, 35, -24, 69, 69, -14, -15,
	-22, 69, 63, -26, 69, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, 62, 60,
	60, 69, 62, -12, -13, -22, -8, -7, -7, -11,
	-10, -11, -10, -10, -10, -10, -10, -10, 61, 69,
	60, -11, -10, -10, -10, -10, -10, 64, -4, 64,
	69, -8, 25, -11, -11, 12, 62, 64, 69, 22,
	26, 62, 64, 62, 61, 60, -10, 61, -25, 8,
	-22, -15, 65, -7, 62, 62, -7, 12, -13, -7,
	-10, -11, 64, 62, 69, -10, 26, 5, -7, 23,
	62, -10, 65, 5, 62, -26, 25, -7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 27,
	0, 0, 29, 31, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 114, 115, -2, 0,
	18, 117, 118, 119, 120, 0, 16, 17, 0, 0,
	0, 0, 1, 0, 0, 0, -2, 65, 66, 67,
	68, 69, 70, 72, 73, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 104, 0, 32, 15, 18, 33,
	71, 34, 35, 36, 37, 0, 0, 0, 0, 42,
	43, 44, 45, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 101, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 105, 106, 0, 26, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 74, 75, 76, 77, 78, 79,
	0, 0, 0, 0, 0, 0, 0, 0, 41, 0,
	100, 100, 121, 123, 122, 20, 0, 54, 0, 55,
	0, 0, 56, 0, 0, 0, 48, 21, 22, 23,
	0, 0, 100, 0, 62, 0, 63, 0, 0, 64,
	0, 0, 108, 109, 110, 111, 112, 19, 9, 6,
	0, 116, 0, 0, 15, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 100,
	100, 96, 0, 0, 13, 0, 30, 0, 0, 0,
	99, 0, 49, 50, 51, 52, 53, 0, 103, 0,
	0, 0, 57, 58, 59, 60, 61, 0, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 0, 14, 0,
	40, 46, 0, 47, 102, 100, 0, 97, 0, 4,
	0, 5, 0, 0, 94, 95, 0, 0, 12, 0,
	98, 0, 0, 107, 0, 0, 39, 10, 0, 38,
	24, 0, 0, 11, 25, 0, 0, 0, 28,
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
		//line waccparser.y:110
		{
	                                         parserlex.(*Lexer).prog = &Program{ClassList : parserS[parserpt-3].classes , FunctionList : parserS[parserpt-2].functions , StatList : parserS[parserpt-1].stmts , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
	                                         }
	case 2:
		//line waccparser.y:115
		{ parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)}
	case 3:
		//line waccparser.y:116
		{ parserVAL.classes = []*Class{} }
	case 4:
		//line waccparser.y:119
		{ if !checkClassIdent(parserS[parserpt-4].ident) {
	                                                         	parserlex.Error("Invalid class name")
	                                                     }
	                                                     parserVAL.class = &Class{Ident : ClassType(parserS[parserpt-4].ident), FieldList : parserS[parserpt-2].fields , FunctionList : parserS[parserpt-1].functions}
	                                                   }
	case 5:
		//line waccparser.y:125
		{ parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)}
	case 6:
		//line waccparser.y:126
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
		//line waccparser.y:155
		{ parserVAL.assignlhs = parserS[parserpt-0].fieldaccess}
	case 19:
		//line waccparser.y:156
		{ parserVAL.assignlhs = ThisInstance{parserS[parserpt-0].ident} }
	case 20:
		//line waccparser.y:158
		{parserVAL.fieldaccess = FieldAccess{ &parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-2].ident, parserS[parserpt-0].ident, } }
	case 21:
		//line waccparser.y:160
		{parserVAL.assignrhs = parserS[parserpt-0].expr}
	case 22:
		//line waccparser.y:161
		{parserVAL.assignrhs = parserS[parserpt-0].arrayliter}
	case 23:
		//line waccparser.y:162
		{parserVAL.assignrhs = parserS[parserpt-0].pairelem}
	case 24:
		//line waccparser.y:163
		{ parserVAL.assignrhs = NewObject{Class : ClassType(parserS[parserpt-3].ident) , Init : parserS[parserpt-1].exprs , Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input}}
	case 25:
		//line waccparser.y:164
		{ parserVAL.assignrhs = NewPair{FstExpr : parserS[parserpt-3].expr, SndExpr : parserS[parserpt-1].expr, Pos : parserS[parserpt-5].pos, FileText :&parserlex.(*Lexer).input } }
	case 26:
		//line waccparser.y:166
		{ parserVAL.stmts = append(parserS[parserpt-2].stmts,parserS[parserpt-0].stmt)   }
	case 27:
		//line waccparser.y:167
		{ parserVAL.stmts = []Statement{parserS[parserpt-0].stmt} }
	case 28:
		//line waccparser.y:169
		{
	                                                                                                                 stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                                                                                 w := While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-11].pos, FileText :&parserlex.(*Lexer).input}
	                                                                                                                 d := Declare{DecType : parserS[parserpt-10].typedefinition, Lhs : parserS[parserpt-9].ident, Rhs : parserS[parserpt-7].assignrhs, Pos : parserS[parserpt-11].pos ,FileText :&parserlex.(*Lexer).input }
	                                                                                                                 parserVAL.stmts = []Statement{d,w}
	                                                                                                                }
	case 29:
		//line waccparser.y:177
		{ parserVAL.stmt = Skip{Pos : parserS[parserpt-0].pos ,FileText :&parserlex.(*Lexer).input } }
	case 30:
		//line waccparser.y:178
		{ parserVAL.stmt = Declare{DecType : parserS[parserpt-3].typedefinition, Lhs : parserS[parserpt-2].ident, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input } }
	case 31:
		//line waccparser.y:179
		{ parserVAL.stmt = parserS[parserpt-0].stmt }
	case 32:
		//line waccparser.y:180
		{ parserVAL.stmt = Read{ &parserlex.(*Lexer).input, parserS[parserpt-1].pos , parserS[parserpt-0].assignlhs, } }
	case 33:
		//line waccparser.y:181
		{ parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 34:
		//line waccparser.y:182
		{ parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 35:
		//line waccparser.y:183
		{ parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 36:
		//line waccparser.y:184
		{ parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 37:
		//line waccparser.y:185
		{ parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 38:
		//line waccparser.y:186
		{ parserVAL.stmt = If{Conditional : parserS[parserpt-5].expr, ThenStat : parserS[parserpt-3].stmts, ElseStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input } }
	case 39:
		//line waccparser.y:187
		{
	                                                              stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                              parserVAL.stmt = While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input}
	                                                             }
	case 40:
		//line waccparser.y:191
		{ parserVAL.stmt = While{Conditional : parserS[parserpt-3].expr, DoStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input} }
	case 41:
		//line waccparser.y:192
		{ parserVAL.stmt = Scope{StatList : parserS[parserpt-1].stmts, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input } }
	case 42:
		//line waccparser.y:193
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 43:
		//line waccparser.y:197
		{ parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 44:
		//line waccparser.y:200
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 45:
		//line waccparser.y:204
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 46:
		//line waccparser.y:208
		{ parserVAL.stmt = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 47:
		//line waccparser.y:209
		{ parserVAL.stmt = CallInstance{Class : (parserS[parserpt-3].fieldaccess.(FieldAccess)).ObjectName, Func: (parserS[parserpt-3].fieldaccess.(FieldAccess)).Field, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 48:
		//line waccparser.y:211
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].assignlhs, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 49:
		//line waccparser.y:212
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 50:
		//line waccparser.y:213
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 51:
		//line waccparser.y:214
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 52:
		//line waccparser.y:215
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 53:
		//line waccparser.y:216
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 54:
		//line waccparser.y:217
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 55:
		//line waccparser.y:218
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 56:
		//line waccparser.y:219
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : MUL,  Right : parserS[parserpt-2].ident,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 57:
		//line waccparser.y:221
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 58:
		//line waccparser.y:222
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 59:
		//line waccparser.y:223
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 60:
		//line waccparser.y:224
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 61:
		//line waccparser.y:225
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 62:
		//line waccparser.y:226
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 63:
		//line waccparser.y:227
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 64:
		//line waccparser.y:228
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : MUL,  Right : parserS[parserpt-2].fieldaccess,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 65:
		//line waccparser.y:230
		{ parserVAL.expr =  parserS[parserpt-0].integer }
	case 66:
		//line waccparser.y:231
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 67:
		//line waccparser.y:232
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 68:
		//line waccparser.y:233
		{ parserVAL.expr =  parserS[parserpt-0].character }
	case 69:
		//line waccparser.y:234
		{ parserVAL.expr =  parserS[parserpt-0].stringconst }
	case 70:
		//line waccparser.y:235
		{ parserVAL.expr = parserS[parserpt-0].fieldaccess }
	case 71:
		//line waccparser.y:236
		{ parserVAL.expr = parserS[parserpt-0].ident}
	case 72:
		//line waccparser.y:237
		{ parserVAL.expr =  parserS[parserpt-0].pairliter }
	case 73:
		//line waccparser.y:238
		{ parserVAL.expr =  parserS[parserpt-0].arrayelem }
	case 74:
		//line waccparser.y:239
		{ parserVAL.expr = Unop{Unary : NOT, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 75:
		//line waccparser.y:240
		{ parserVAL.expr = Unop{Unary : LEN, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 76:
		//line waccparser.y:241
		{ parserVAL.expr = Unop{Unary : ORD, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 77:
		//line waccparser.y:242
		{ parserVAL.expr = Unop{Unary : CHR, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 78:
		//line waccparser.y:243
		{ parserVAL.expr = Unop{Unary : SUB, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 79:
		//line waccparser.y:244
		{ parserVAL.expr = parserS[parserpt-0].expr }
	case 80:
		//line waccparser.y:245
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 81:
		//line waccparser.y:246
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : SUB,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 82:
		//line waccparser.y:247
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 83:
		//line waccparser.y:248
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 84:
		//line waccparser.y:249
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 85:
		//line waccparser.y:250
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 86:
		//line waccparser.y:251
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 87:
		//line waccparser.y:252
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 88:
		//line waccparser.y:253
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 89:
		//line waccparser.y:254
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : EQ,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 90:
		//line waccparser.y:255
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : NEQ,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 91:
		//line waccparser.y:256
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : AND,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 92:
		//line waccparser.y:257
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : OR,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 93:
		//line waccparser.y:258
		{ parserVAL.expr = parserS[parserpt-1].expr }
	case 94:
		//line waccparser.y:259
		{ parserVAL.expr = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 95:
		//line waccparser.y:260
		{ parserVAL.expr = CallInstance{Class : (parserS[parserpt-3].fieldaccess.(FieldAccess)).ObjectName, Func: (parserS[parserpt-3].fieldaccess.(FieldAccess)).Field, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 96:
		//line waccparser.y:261
		{ parserVAL.expr = ThisInstance{parserS[parserpt-0].ident} }
	case 97:
		//line waccparser.y:263
		{ parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs } }
	case 98:
		//line waccparser.y:265
		{parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)}
	case 99:
		//line waccparser.y:266
		{parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}}
	case 100:
		//line waccparser.y:267
		{parserVAL.exprs = []Evaluation{}}
	case 101:
		//line waccparser.y:269
		{parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs : parserS[parserpt-0].exprs, Pos : parserS[parserpt-1].pos,FileText :&parserlex.(*Lexer).input  } }
	case 102:
		//line waccparser.y:271
		{parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)}
	case 103:
		//line waccparser.y:272
		{parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}}
	case 104:
		//line waccparser.y:274
		{ parserVAL.pairliter =  PairLiter{} }
	case 105:
		//line waccparser.y:276
		{ parserVAL.pairelem = PairElem{Fsnd: Fst, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 106:
		//line waccparser.y:277
		{ parserVAL.pairelem = PairElem{Fsnd: Snd, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 107:
		//line waccparser.y:279
		{ parserVAL.typedefinition = PairType{FstType : parserS[parserpt-3].pairelemtype, SndType : parserS[parserpt-1].pairelemtype} }
	case 108:
		//line waccparser.y:281
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 109:
		//line waccparser.y:282
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 110:
		//line waccparser.y:283
		{ parserVAL.pairelemtype = Pair}
	case 111:
		//line waccparser.y:284
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition}
	case 112:
		//line waccparser.y:285
		{ parserVAL.pairelemtype = ClassType(parserS[parserpt-0].ident)}
	case 113:
		//line waccparser.y:287
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 114:
		//line waccparser.y:288
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 115:
		//line waccparser.y:289
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 116:
		//line waccparser.y:290
		{ parserVAL.typedefinition = ClassType(parserS[parserpt-0].ident) }
	case 117:
		//line waccparser.y:292
		{ parserVAL.typedefinition =  Int }
	case 118:
		//line waccparser.y:293
		{ parserVAL.typedefinition =  Bool }
	case 119:
		//line waccparser.y:294
		{ parserVAL.typedefinition =  Char }
	case 120:
		//line waccparser.y:295
		{ parserVAL.typedefinition =  String }
	case 121:
		//line waccparser.y:297
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 122:
		//line waccparser.y:298
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 123:
		//line waccparser.y:299
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	}
	goto parserstack /* stack new state and value */
}
