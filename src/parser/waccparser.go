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
	"CLASSIDENT",
	"INTEGER",
	"CHARACTER",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line waccparser.y:272

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 133,
	56, 92,
	-2, 89,
	-1, 134,
	56, 93,
	-2, 90,
}

const parserNprod = 100
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 455

var parserAct = []int{

	148, 26, 77, 7, 142, 132, 4, 10, 61, 25,
	29, 138, 90, 28, 37, 146, 147, 53, 70, 71,
	72, 73, 74, 75, 28, 41, 41, 76, 50, 41,
	41, 41, 28, 175, 201, 180, 83, 84, 176, 131,
	40, 173, 88, 91, 35, 36, 104, 207, 28, 206,
	106, 202, 187, 93, 101, 30, 31, 32, 33, 34,
	205, 200, 184, 121, 122, 123, 124, 125, 126, 127,
	110, 112, 111, 108, 109, 119, 120, 114, 116, 113,
	115, 117, 118, 52, 134, 136, 140, 188, 139, 91,
	143, 39, 133, 196, 186, 39, 184, 42, 149, 93,
	150, 144, 151, 152, 39, 153, 154, 89, 39, 157,
	158, 159, 160, 161, 162, 163, 164, 165, 166, 167,
	168, 169, 100, 130, 156, 183, 39, 178, 184, 179,
	87, 38, 171, 172, 103, 182, 82, 28, 28, 24,
	99, 23, 98, 145, 174, 107, 181, 11, 15, 16,
	17, 18, 19, 20, 21, 79, 102, 87, 22, 41,
	88, 97, 14, 35, 36, 30, 31, 32, 33, 34,
	80, 49, 105, 81, 9, 134, 136, 191, 139, 189,
	39, 193, 143, 133, 195, 198, 28, 192, 39, 85,
	194, 199, 114, 116, 113, 115, 28, 204, 203, 197,
	177, 6, 13, 28, 94, 95, 35, 36, 190, 2,
	27, 59, 78, 62, 86, 63, 64, 65, 48, 92,
	137, 67, 66, 46, 45, 47, 43, 44, 30, 31,
	32, 33, 34, 55, 56, 69, 96, 68, 12, 141,
	8, 49, 5, 3, 58, 60, 1, 54, 57, 62,
	0, 63, 64, 65, 51, 0, 0, 67, 66, 30,
	31, 32, 33, 34, 30, 31, 32, 33, 135, 55,
	56, 69, 0, 68, 0, 0, 0, 0, 0, 0,
	58, 60, 0, 54, 57, 110, 112, 111, 108, 109,
	119, 120, 114, 116, 113, 115, 117, 118, 0, 0,
	0, 0, 0, 0, 0, 0, 208, 110, 112, 111,
	108, 109, 119, 120, 114, 116, 113, 115, 117, 118,
	0, 0, 0, 0, 0, 0, 0, 0, 170, 110,
	112, 111, 108, 109, 119, 120, 114, 116, 113, 115,
	117, 118, 0, 0, 0, 0, 0, 0, 0, 185,
	110, 112, 111, 108, 109, 119, 120, 114, 116, 113,
	115, 117, 118, 129, 0, 0, 110, 112, 111, 0,
	155, 0, 128, 114, 116, 113, 115, 0, 0, 110,
	112, 111, 108, 109, 119, 120, 114, 116, 113, 115,
	117, 118, 110, 112, 111, 108, 109, 119, 120, 114,
	116, 113, 115, 117, 118, 110, 112, 111, 108, 109,
	119, 120, 114, 116, 113, 115, 117, 118, 110, 112,
	111, 108, 109, 119, 0, 114, 116, 113, 115, 117,
	118, 110, 112, 111, 108, 109, 0, 0, 114, 116,
	113, 115, 117, 118, 110, 112, 111, 108, 109, 0,
	0, 114, 116, 113, 115,
}
var parserPact = []int{

	205, -1000, -1000, 195, 137, -1000, -52, 126, -1000, -1000,
	-25, -1000, 37, 185, -37, 18, 216, 216, 216, 216,
	216, 216, 216, 137, 150, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 79, 216, 216, 182, -1000, 137,
	100, 49, 180, 101, 80, -6, 96, -14, 116, 216,
	88, -1000, 115, 367, -1000, -1000, -1000, -1000, -1000, -1000,
	115, -1000, 216, 216, 216, 216, 216, 216, 216, -1000,
	367, 367, 367, 367, 354, 341, 118, -26, -1000, -1000,
	-1000, -1000, 236, 367, 367, 231, -1000, 27, 180, -1000,
	-1000, 367, -1000, -1000, 86, -50, 216, 216, -1000, 216,
	-1000, 216, 216, -1000, 216, 216, 312, 216, 216, 216,
	216, 216, 216, 216, 216, 216, 216, 216, 216, 216,
	216, -1000, -1000, -1000, -1000, 328, 328, 269, 137, 137,
	-1000, -18, -20, -1000, -1000, 79, 103, -28, -1000, -27,
	191, 68, -1000, -30, -1000, 216, 78, 67, 367, 367,
	367, 367, 367, 367, 291, -1000, 35, 328, 328, 147,
	147, 147, -1000, -1000, -1000, -1000, 406, 406, 393, 380,
	-1000, 33, 64, 236, 200, 231, -1000, 137, 181, 231,
	-1000, 32, 216, -1000, 216, -1000, -1000, 137, -1000, 2,
	-1000, -31, -1000, 46, 137, -1000, 216, 1, 367, 29,
	-1000, 73, -1000, 42, 247, -1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 246, 243, 242, 6, 240, 174, 3, 12, 238,
	0, 16, 239, 4, 220, 11, 10, 219, 8, 218,
	211, 9, 2, 1, 210, 5,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 17, 11,
	11, 11, 18, 19, 19, 20, 16, 16, 24, 25,
	25, 25, 22, 22, 22, 21, 21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 3, 1, 1, 4, 3, 4, 4,
	4, 4, 4, 3, 3, 3, 5, 2, 2, 2,
	2, 2, 2, 7, 5, 3, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 0, 2, 4, 3, 1, 2, 2, 6, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	-22, 10, -9, 65, 25, 11, 12, 13, 14, 15,
	16, 17, 21, 4, 2, -21, -23, -24, -18, -16,
	28, 29, 30, 31, 32, 26, 27, 66, 5, 62,
	65, 56, 60, 41, 42, 39, 38, 40, -19, 56,
	65, -9, 65, -10, 67, 53, 54, 68, 64, -20,
	65, -18, 33, 35, 36, 37, 42, 41, 57, 55,
	-10, -10, -10, -10, -10, -10, -7, -22, 62, 5,
	20, 23, 57, -10, -10, 7, -6, 57, 60, 58,
	-8, -10, -17, -16, 24, 25, 56, 60, 41, 60,
	42, 60, 60, 38, 60, 56, -10, 57, 41, 42,
	38, 40, 39, 47, 45, 48, 46, 49, 50, 43,
	44, -10, -10, -10, -10, -10, -10, -10, 18, 22,
	5, 65, -25, -21, -23, 32, -22, -14, -15, -22,
	59, -12, -13, -22, -8, 57, 65, -11, -10, -10,
	-10, -10, -10, -10, -10, 58, -11, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	59, -7, -7, 61, -4, 61, 65, 9, 59, 61,
	65, -10, 57, 58, 61, 58, 59, 19, 23, -25,
	8, -22, -15, -7, 9, -13, 61, -11, -10, -7,
	59, 65, 5, -7, -10, 59, 20, 5, 59,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 24,
	0, 25, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 92, 93, 94, 16, 17,
	95, 96, 97, 98, 0, 0, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 82, 0,
	0, 37, 15, 38, 50, 51, 52, 53, 54, 55,
	56, 57, 0, 0, 0, 0, 0, 0, 0, 85,
	39, 40, 41, 42, 0, 0, 0, 0, 46, 47,
	48, 49, 0, 86, 87, 0, 23, 0, 0, 99,
	27, 18, 19, 20, 0, 0, 81, 0, 33, 0,
	34, 0, 0, 35, 0, 0, 0, 81, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 58, 59, 60, 61, 62, 63, 0, 0, 0,
	45, 0, 0, -2, -2, 91, 0, 9, 6, 0,
	0, 0, 13, 0, 26, 0, 0, 0, 80, 28,
	29, 30, 31, 32, 0, 84, 0, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	77, 0, 0, 0, 0, 0, 7, 0, 0, 0,
	14, 0, 81, 78, 0, 83, 36, 0, 44, 0,
	4, 0, 5, 0, 0, 12, 0, 0, 79, 0,
	88, 0, 10, 0, 0, 22, 43, 11, 21,
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
	62, 63, 64, 65, 66, 67, 68,
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
			parserVAL.class = &Class{Ident: parserS[parserpt-4].ident, FieldList: parserS[parserpt-2].fields, FunctionList: parserS[parserpt-1].functions}
		}
	case 5:
		//line waccparser.y:121
		{
			parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)
		}
	case 6:
		//line waccparser.y:122
		{
			parserVAL.fields = []Field{parserS[parserpt-0].field}
		}
	case 7:
		//line waccparser.y:124
		{
			parserVAL.field = Field{FieldType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 8:
		//line waccparser.y:126
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 9:
		//line waccparser.y:127
		{
			parserVAL.functions = []*Function{}
		}
	case 10:
		//line waccparser.y:130
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		//line waccparser.y:136
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 12:
		//line waccparser.y:142
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 13:
		//line waccparser.y:143
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 14:
		//line waccparser.y:145
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 15:
		//line waccparser.y:147
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 16:
		//line waccparser.y:148
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 17:
		//line waccparser.y:149
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 18:
		//line waccparser.y:151
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 19:
		//line waccparser.y:152
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 20:
		//line waccparser.y:153
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 21:
		//line waccparser.y:154
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:155
		{
			parserVAL.assignrhs = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:157
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 24:
		//line waccparser.y:158
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 25:
		//line waccparser.y:161
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 26:
		//line waccparser.y:162
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 27:
		//line waccparser.y:163
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 28:
		//line waccparser.y:165
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 29:
		//line waccparser.y:166
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 30:
		//line waccparser.y:167
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:168
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 32:
		//line waccparser.y:169
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 33:
		//line waccparser.y:171
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 34:
		//line waccparser.y:172
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 35:
		//line waccparser.y:173
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 36:
		//line waccparser.y:175
		{
			parserVAL.stmt = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 37:
		//line waccparser.y:177
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 38:
		//line waccparser.y:178
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 39:
		//line waccparser.y:179
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 40:
		//line waccparser.y:180
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 41:
		//line waccparser.y:181
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 42:
		//line waccparser.y:182
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 43:
		//line waccparser.y:183
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		//line waccparser.y:184
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:185
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		//line waccparser.y:186
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 47:
		//line waccparser.y:190
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 48:
		//line waccparser.y:193
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 49:
		//line waccparser.y:197
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 50:
		//line waccparser.y:202
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 51:
		//line waccparser.y:204
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 52:
		//line waccparser.y:205
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 53:
		//line waccparser.y:207
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 54:
		//line waccparser.y:208
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 55:
		//line waccparser.y:209
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 56:
		//line waccparser.y:210
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 57:
		//line waccparser.y:211
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 58:
		//line waccparser.y:213
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 59:
		//line waccparser.y:214
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 60:
		//line waccparser.y:215
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		//line waccparser.y:216
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		//line waccparser.y:217
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:218
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 64:
		//line waccparser.y:221
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:222
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:223
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:224
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:225
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:226
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:227
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:228
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:229
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:230
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:231
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:232
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:233
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 77:
		//line waccparser.y:234
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 78:
		//line waccparser.y:237
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 79:
		//line waccparser.y:239
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 80:
		//line waccparser.y:240
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 81:
		//line waccparser.y:241
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 82:
		//line waccparser.y:243
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 83:
		//line waccparser.y:245
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 84:
		//line waccparser.y:246
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 85:
		//line waccparser.y:249
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 86:
		//line waccparser.y:252
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 87:
		//line waccparser.y:253
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 88:
		//line waccparser.y:255
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 89:
		//line waccparser.y:257
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 90:
		//line waccparser.y:258
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 91:
		//line waccparser.y:259
		{
			parserVAL.pairelemtype = Pair
		}
	case 92:
		//line waccparser.y:261
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 93:
		//line waccparser.y:262
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 94:
		//line waccparser.y:263
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 95:
		//line waccparser.y:265
		{
			parserVAL.typedefinition = Int
		}
	case 96:
		//line waccparser.y:266
		{
			parserVAL.typedefinition = Bool
		}
	case 97:
		//line waccparser.y:267
		{
			parserVAL.typedefinition = Char
		}
	case 98:
		//line waccparser.y:268
		{
			parserVAL.typedefinition = String
		}
	case 99:
		//line waccparser.y:270
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
