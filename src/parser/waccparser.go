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

//line waccparser.y:271

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 130,
	56, 91,
	-2, 88,
	-1, 131,
	56, 92,
	-2, 89,
}

const parserNprod = 99
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 451

var parserAct = []int{

	89, 144, 75, 7, 135, 129, 4, 10, 59, 139,
	28, 36, 88, 27, 34, 35, 51, 68, 69, 70,
	71, 72, 73, 27, 25, 40, 74, 143, 40, 171,
	40, 27, 40, 40, 196, 81, 82, 176, 24, 172,
	201, 128, 39, 169, 86, 202, 182, 27, 183, 104,
	102, 99, 91, 50, 29, 30, 31, 32, 33, 41,
	195, 118, 119, 120, 121, 122, 123, 124, 107, 109,
	108, 105, 106, 116, 117, 111, 113, 110, 112, 114,
	115, 179, 38, 133, 180, 137, 136, 38, 140, 38,
	200, 191, 180, 87, 85, 145, 146, 91, 147, 141,
	148, 149, 38, 150, 151, 131, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 130,
	174, 98, 175, 96, 178, 85, 197, 127, 86, 167,
	168, 37, 101, 80, 27, 27, 23, 142, 22, 97,
	40, 170, 95, 177, 11, 14, 15, 16, 17, 18,
	19, 20, 48, 77, 100, 21, 103, 9, 189, 173,
	34, 35, 29, 30, 31, 32, 33, 83, 78, 6,
	2, 79, 133, 186, 136, 184, 187, 188, 140, 145,
	192, 193, 27, 38, 38, 190, 194, 26, 38, 57,
	47, 27, 199, 198, 131, 90, 84, 134, 27, 13,
	92, 93, 34, 35, 111, 113, 110, 112, 130, 60,
	76, 61, 62, 63, 138, 8, 5, 65, 64, 45,
	44, 46, 42, 43, 29, 30, 31, 32, 33, 53,
	54, 67, 94, 66, 3, 185, 12, 48, 1, 0,
	56, 58, 0, 52, 55, 60, 0, 61, 62, 63,
	0, 49, 0, 65, 64, 29, 30, 31, 32, 33,
	29, 30, 31, 32, 132, 53, 54, 67, 0, 66,
	0, 0, 0, 0, 0, 0, 56, 58, 0, 52,
	55, 107, 109, 108, 105, 106, 116, 117, 111, 113,
	110, 112, 114, 115, 0, 0, 0, 0, 0, 0,
	0, 0, 203, 107, 109, 108, 105, 106, 116, 117,
	111, 113, 110, 112, 114, 115, 0, 0, 0, 0,
	0, 0, 0, 0, 166, 107, 109, 108, 105, 106,
	116, 117, 111, 113, 110, 112, 114, 115, 0, 0,
	0, 0, 0, 0, 0, 181, 107, 109, 108, 105,
	106, 116, 117, 111, 113, 110, 112, 114, 115, 126,
	0, 0, 107, 109, 108, 0, 152, 0, 125, 111,
	113, 110, 112, 0, 0, 107, 109, 108, 105, 106,
	116, 117, 111, 113, 110, 112, 114, 115, 107, 109,
	108, 105, 106, 116, 117, 111, 113, 110, 112, 114,
	115, 107, 109, 108, 105, 106, 116, 117, 111, 113,
	110, 112, 114, 115, 107, 109, 108, 105, 106, 116,
	0, 111, 113, 110, 112, 114, 115, 107, 109, 108,
	105, 106, 0, 0, 111, 113, 110, 112, 114, 115,
	107, 109, 108, 105, 106, 0, 0, 111, 113, 110,
	112,
}
var parserPact = []int{

	166, -1000, -1000, 163, 134, -1000, -55, 126, -1000, -1000,
	-23, -1000, -1, 181, -12, 212, 212, 212, 212, 212,
	212, 212, 134, 148, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 76, 212, 212, 160, -1000, 134, 68,
	35, 176, 82, 79, -9, 94, -10, 100, 212, -1000,
	96, 363, -1000, -1000, -1000, -1000, -1000, -1000, 96, -1000,
	212, 212, 212, 212, 212, 212, 212, -1000, 363, 363,
	363, 363, 350, 337, 122, -24, -1000, -1000, -1000, -1000,
	232, 363, 363, 196, -1000, 26, 176, -1000, -1000, 363,
	-1000, -1000, 80, -38, 212, 212, -1000, 212, -1000, 212,
	212, -1000, 212, 212, 308, 212, 212, 212, 212, 212,
	212, 212, 212, 212, 212, 212, 212, 212, -1000, -1000,
	-1000, -1000, 324, 324, 265, 134, 134, -1000, -16, -18,
	-1000, -1000, 76, 84, -32, -1000, -26, 150, 61, -1000,
	-28, -1000, 212, 67, 23, 363, 363, 363, 363, 363,
	363, 287, -1000, 324, 324, 159, 159, 159, -1000, -1000,
	-1000, -1000, 402, 402, 389, 376, -1000, 27, 25, 232,
	227, 196, -1000, 134, 149, 196, -1000, 30, 212, -1000,
	212, -1000, 134, -1000, 1, -1000, -31, -1000, 121, 134,
	-1000, 212, 31, 363, 20, -1000, 37, -1000, 40, 243,
	-1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 238, 234, 216, 6, 215, 157, 3, 12, 236,
	0, 1, 214, 9, 197, 4, 10, 195, 8, 190,
	189, 38, 2, 24, 187, 5,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 8, 8,
	8, 8, 8, 7, 7, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 17, 11, 11,
	11, 18, 19, 19, 20, 16, 16, 24, 25, 25,
	25, 22, 22, 22, 21, 21, 21, 21, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 1,
	1, 6, 5, 3, 1, 1, 4, 3, 4, 4,
	4, 4, 4, 3, 3, 3, 2, 2, 2, 2,
	2, 2, 7, 5, 3, 2, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	0, 2, 4, 3, 1, 2, 2, 6, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	-22, 10, -9, 65, 11, 12, 13, 14, 15, 16,
	17, 21, 4, 2, -21, -23, -24, -18, -16, 28,
	29, 30, 31, 32, 26, 27, 66, 5, 62, 65,
	56, 60, 41, 42, 39, 38, 40, -19, 56, -9,
	65, -10, 67, 53, 54, 68, 64, -20, 65, -18,
	33, 35, 36, 37, 42, 41, 57, 55, -10, -10,
	-10, -10, -10, -10, -7, -22, 62, 5, 20, 23,
	57, -10, -10, 7, -6, 57, 60, 58, -8, -10,
	-17, -16, 24, 25, 56, 60, 41, 60, 42, 60,
	60, 38, 60, 56, -10, 41, 42, 38, 40, 39,
	47, 45, 48, 46, 49, 50, 43, 44, -10, -10,
	-10, -10, -10, -10, -10, 18, 22, 5, 65, -25,
	-21, -23, 32, -22, -14, -15, -22, 59, -12, -13,
	-22, -8, 57, 65, -11, -10, -10, -10, -10, -10,
	-10, -10, 58, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, 59, -7, -7, 61,
	-4, 61, 65, 9, 59, 61, 65, -10, 57, 58,
	61, 58, 19, 23, -25, 8, -22, -15, -7, 9,
	-13, 61, -11, -10, -7, 59, 65, 5, -7, -10,
	59, 20, 5, 59,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 24,
	0, 25, 0, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 91, 92, 93, 16, 17, 94,
	95, 96, 97, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 81, 0, 36,
	15, 37, 49, 50, 51, 52, 53, 54, 55, 56,
	0, 0, 0, 0, 0, 0, 0, 84, 38, 39,
	40, 41, 0, 0, 0, 0, 45, 46, 47, 48,
	0, 85, 86, 0, 23, 0, 0, 98, 27, 18,
	19, 20, 0, 0, 80, 0, 33, 0, 34, 0,
	0, 35, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 57, 58,
	59, 60, 61, 62, 0, 0, 0, 44, 0, 0,
	-2, -2, 90, 0, 9, 6, 0, 0, 0, 13,
	0, 26, 0, 0, 0, 79, 28, 29, 30, 31,
	32, 0, 83, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 0, 0, 0,
	0, 0, 7, 0, 0, 0, 14, 0, 80, 77,
	0, 82, 0, 43, 0, 4, 0, 5, 0, 0,
	12, 0, 0, 78, 0, 87, 0, 10, 0, 0,
	22, 42, 11, 21,
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
		//line waccparser.y:176
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 37:
		//line waccparser.y:177
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 38:
		//line waccparser.y:178
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 39:
		//line waccparser.y:179
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 40:
		//line waccparser.y:180
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 41:
		//line waccparser.y:181
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 42:
		//line waccparser.y:182
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 43:
		//line waccparser.y:183
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		//line waccparser.y:184
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:185
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 46:
		//line waccparser.y:189
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 47:
		//line waccparser.y:192
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 48:
		//line waccparser.y:196
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 49:
		//line waccparser.y:201
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 50:
		//line waccparser.y:203
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 51:
		//line waccparser.y:204
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 52:
		//line waccparser.y:206
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 53:
		//line waccparser.y:207
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 54:
		//line waccparser.y:208
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 55:
		//line waccparser.y:209
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 56:
		//line waccparser.y:210
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 57:
		//line waccparser.y:212
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 58:
		//line waccparser.y:213
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 59:
		//line waccparser.y:214
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 60:
		//line waccparser.y:215
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		//line waccparser.y:216
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		//line waccparser.y:217
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 63:
		//line waccparser.y:220
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		//line waccparser.y:221
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:222
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:223
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:224
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:225
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:226
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:227
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:228
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:229
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:230
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:231
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:232
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 76:
		//line waccparser.y:233
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 77:
		//line waccparser.y:236
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 78:
		//line waccparser.y:238
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 79:
		//line waccparser.y:239
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 80:
		//line waccparser.y:240
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 81:
		//line waccparser.y:242
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 82:
		//line waccparser.y:244
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 83:
		//line waccparser.y:245
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 84:
		//line waccparser.y:248
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 85:
		//line waccparser.y:251
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 86:
		//line waccparser.y:252
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 87:
		//line waccparser.y:254
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 88:
		//line waccparser.y:256
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 89:
		//line waccparser.y:257
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 90:
		//line waccparser.y:258
		{
			parserVAL.pairelemtype = Pair
		}
	case 91:
		//line waccparser.y:260
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
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
		//line waccparser.y:264
		{
			parserVAL.typedefinition = Int
		}
	case 95:
		//line waccparser.y:265
		{
			parserVAL.typedefinition = Bool
		}
	case 96:
		//line waccparser.y:266
		{
			parserVAL.typedefinition = Char
		}
	case 97:
		//line waccparser.y:267
		{
			parserVAL.typedefinition = String
		}
	case 98:
		//line waccparser.y:269
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
