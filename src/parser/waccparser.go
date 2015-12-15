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
	class          Class
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

//line waccparser.y:264

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 131,
	56, 90,
	-2, 87,
	-1, 132,
	56, 91,
	-2, 88,
}

const parserNprod = 98
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 452

var parserAct = []int{

	90, 146, 75, 6, 137, 24, 9, 141, 59, 36,
	27, 145, 26, 203, 23, 51, 68, 69, 70, 71,
	72, 73, 26, 3, 130, 74, 181, 40, 35, 182,
	26, 40, 89, 77, 81, 82, 198, 40, 9, 83,
	178, 40, 33, 34, 26, 40, 174, 26, 78, 105,
	129, 79, 92, 173, 39, 38, 202, 184, 182, 171,
	87, 119, 120, 121, 122, 123, 124, 125, 108, 110,
	109, 106, 107, 117, 118, 112, 114, 111, 113, 115,
	116, 50, 204, 134, 199, 185, 132, 138, 197, 142,
	76, 193, 135, 103, 100, 131, 147, 148, 92, 149,
	38, 150, 151, 41, 152, 153, 128, 155, 156, 157,
	158, 159, 160, 161, 162, 163, 164, 165, 166, 167,
	143, 102, 99, 176, 38, 177, 86, 97, 88, 87,
	169, 170, 37, 86, 40, 26, 26, 191, 180, 38,
	98, 38, 22, 101, 21, 179, 96, 80, 144, 38,
	10, 13, 14, 15, 16, 17, 18, 19, 48, 104,
	172, 20, 175, 38, 84, 5, 33, 34, 28, 29,
	30, 31, 32, 11, 134, 188, 138, 132, 189, 190,
	142, 147, 194, 195, 26, 192, 131, 49, 196, 38,
	2, 25, 57, 26, 201, 200, 186, 8, 47, 91,
	26, 93, 94, 33, 34, 12, 112, 114, 111, 113,
	60, 136, 61, 62, 63, 140, 7, 4, 65, 64,
	28, 29, 30, 31, 32, 45, 44, 46, 42, 43,
	53, 54, 67, 95, 66, 1, 85, 187, 0, 0,
	0, 56, 58, 48, 52, 55, 60, 0, 61, 62,
	63, 139, 0, 0, 65, 64, 0, 28, 29, 30,
	31, 32, 0, 0, 0, 0, 53, 54, 67, 0,
	66, 28, 29, 30, 31, 32, 0, 56, 58, 0,
	52, 55, 108, 110, 109, 106, 107, 117, 118, 112,
	114, 111, 113, 115, 116, 28, 29, 30, 31, 133,
	0, 0, 0, 205, 108, 110, 109, 106, 107, 117,
	118, 112, 114, 111, 113, 115, 116, 0, 0, 0,
	0, 0, 0, 0, 0, 168, 108, 110, 109, 106,
	107, 117, 118, 112, 114, 111, 113, 115, 116, 0,
	0, 0, 0, 0, 0, 0, 183, 108, 110, 109,
	106, 107, 117, 118, 112, 114, 111, 113, 115, 116,
	127, 0, 0, 108, 110, 109, 0, 154, 0, 126,
	112, 114, 111, 113, 0, 0, 108, 110, 109, 106,
	107, 117, 118, 112, 114, 111, 113, 115, 116, 108,
	110, 109, 106, 107, 117, 118, 112, 114, 111, 113,
	115, 116, 108, 110, 109, 106, 107, 117, 118, 112,
	114, 111, 113, 115, 116, 108, 110, 109, 106, 107,
	117, 0, 112, 114, 111, 113, 115, 116, 108, 110,
	109, 106, 107, 0, 0, 112, 114, 111, 113, 115,
	116, 108, 110, 109, 106, 107, 0, 0, 112, 114,
	111, 113,
}
var parserPact = []int{

	186, -1000, 159, 140, -1000, -57, 127, -1000, -1000, -11,
	-1000, 43, 187, 16, 213, 213, 213, 213, 213, 213,
	213, 140, 28, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 90, 213, 213, 140, 157, -1000, 140, 69,
	70, 177, 86, 80, 34, 83, 33, 103, 213, -1000,
	102, 364, -1000, -1000, -1000, -1000, -1000, -1000, 102, -1000,
	213, 213, 213, 213, 213, 213, 213, -1000, 364, 364,
	364, 364, 351, 338, 101, -15, -1000, -1000, -1000, -1000,
	267, 364, 364, 87, 243, -1000, 192, 177, -1000, -1000,
	364, -1000, -1000, 91, -54, 213, 213, -1000, 213, -1000,
	213, 213, -1000, 213, 213, 309, 213, 213, 213, 213,
	213, 213, 213, 213, 213, 213, 213, 213, 213, -1000,
	-1000, -1000, -1000, 325, 325, 266, 140, 140, -1000, 0,
	-2, -1000, -1000, 90, 78, -1000, -8, -1000, -19, 153,
	64, -1000, -25, -1000, 213, 81, -32, 364, 364, 364,
	364, 364, 364, 288, -1000, 325, 325, 161, 161, 161,
	-1000, -1000, -1000, -1000, 403, 403, 390, 377, -1000, 38,
	62, 267, 229, 243, -1000, 140, 128, 243, -1000, 30,
	213, -1000, 213, -1000, 140, -1000, 29, -1000, -29, -1000,
	79, 140, -1000, 213, -3, 364, -7, -1000, 76, -1000,
	77, 244, -1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 235, 217, 23, 216, 197, 3, 32, 173, 0,
	1, 215, 7, 211, 4, 10, 199, 8, 198, 192,
	14, 2, 5, 191, 24,
}
var parserR1 = []int{

	0, 1, 1, 2, 13, 13, 14, 3, 3, 4,
	4, 11, 11, 12, 8, 8, 8, 7, 7, 7,
	7, 7, 6, 6, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 16, 10, 10, 10,
	17, 18, 18, 19, 15, 15, 23, 24, 24, 24,
	21, 21, 21, 20, 20, 20, 20, 22,
}
var parserR2 = []int{

	0, 4, 5, 6, 3, 1, 2, 2, 0, 7,
	8, 3, 1, 2, 1, 1, 1, 1, 1, 1,
	6, 5, 3, 1, 1, 4, 3, 4, 4, 4,
	4, 4, 3, 3, 3, 2, 2, 2, 2, 2,
	2, 7, 5, 3, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 0,
	2, 4, 3, 1, 2, 2, 6, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -3, -2, 6, -6, -4, -5, -21,
	10, -8, 65, 11, 12, 13, 14, 15, 16, 17,
	21, 4, 2, -20, -22, -23, -17, -15, 28, 29,
	30, 31, 32, 26, 27, -3, 66, 5, 62, 65,
	56, 60, 41, 42, 39, 38, 40, -18, 56, -8,
	65, -9, 67, 53, 54, 68, 64, -19, 65, -17,
	33, 35, 36, 37, 42, 41, 57, 55, -9, -9,
	-9, -9, -9, -9, -6, -21, 62, 5, 20, 23,
	57, -9, -9, -6, 7, -5, 57, 60, 58, -7,
	-9, -16, -15, 24, 25, 56, 60, 41, 60, 42,
	60, 60, 38, 60, 56, -9, 41, 42, 38, 40,
	39, 47, 45, 48, 46, 49, 50, 43, 44, -9,
	-9, -9, -9, -9, -9, -9, 18, 22, 5, 65,
	-24, -20, -22, 32, -21, 5, -13, -14, -21, 59,
	-11, -12, -21, -7, 57, 65, -10, -9, -9, -9,
	-9, -9, -9, -9, 58, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, 59, -6,
	-6, 61, -3, 61, 65, 9, 59, 61, 65, -9,
	57, 58, 61, 58, 19, 23, -24, 8, -21, -14,
	-6, 9, -12, 61, -10, -9, -6, 59, 65, 5,
	-6, -9, 59, 20, 5, 59,
}
var parserDef = []int{

	0, -2, 8, 0, 8, 0, 0, 7, 23, 0,
	24, 0, 14, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 90, 91, 92, 15, 16, 93, 94,
	95, 96, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 80, 0, 35,
	14, 36, 48, 49, 50, 51, 52, 53, 54, 55,
	0, 0, 0, 0, 0, 0, 0, 83, 37, 38,
	39, 40, 0, 0, 0, 0, 44, 45, 46, 47,
	0, 84, 85, 0, 0, 22, 0, 0, 97, 26,
	17, 18, 19, 0, 0, 79, 0, 32, 0, 33,
	0, 0, 34, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 56,
	57, 58, 59, 60, 61, 0, 0, 0, 43, 0,
	0, -2, -2, 89, 0, 2, 8, 5, 0, 0,
	0, 12, 0, 25, 0, 0, 0, 78, 27, 28,
	29, 30, 31, 0, 82, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 0,
	0, 0, 0, 0, 6, 0, 0, 0, 13, 0,
	79, 76, 0, 81, 0, 42, 0, 3, 0, 4,
	0, 0, 11, 0, 0, 77, 0, 86, 0, 9,
	0, 0, 21, 41, 10, 20,
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
		//line waccparser.y:105
		{
			parserlex.(*Lexer).prog = &Program{FunctionList: parserS[parserpt-2].functions, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 2:
		//line waccparser.y:108
		{
			parserlex.(*Lexer).prog = &Program{Class: parserS[parserpt-3].class, FunctionList: parserS[parserpt-2].functions, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 3:
		//line waccparser.y:112
		{
			parserVAL.class = Class{Ident: parserS[parserpt-4].ident, FieldList: parserS[parserpt-2].fields, FunctionList: parserS[parserpt-1].functions}
		}
	case 4:
		//line waccparser.y:114
		{
			parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)
		}
	case 5:
		//line waccparser.y:115
		{
			parserVAL.fields = []Field{parserS[parserpt-0].field}
		}
	case 6:
		//line waccparser.y:117
		{
			parserVAL.field = Field{FieldType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 7:
		//line waccparser.y:119
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 8:
		//line waccparser.y:120
		{
			parserVAL.functions = []*Function{}
		}
	case 9:
		//line waccparser.y:123
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 10:
		//line waccparser.y:129
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 11:
		//line waccparser.y:135
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 12:
		//line waccparser.y:136
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 13:
		//line waccparser.y:138
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 14:
		//line waccparser.y:140
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 15:
		//line waccparser.y:141
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 16:
		//line waccparser.y:142
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 17:
		//line waccparser.y:144
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 18:
		//line waccparser.y:145
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 19:
		//line waccparser.y:146
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 20:
		//line waccparser.y:147
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 21:
		//line waccparser.y:148
		{
			parserVAL.assignrhs = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:150
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 23:
		//line waccparser.y:151
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 24:
		//line waccparser.y:154
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 25:
		//line waccparser.y:155
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 26:
		//line waccparser.y:156
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 27:
		//line waccparser.y:158
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 28:
		//line waccparser.y:159
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 29:
		//line waccparser.y:160
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 30:
		//line waccparser.y:161
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:162
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 32:
		//line waccparser.y:164
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: PLUS, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 33:
		//line waccparser.y:165
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: SUB, Right: Integer(1), Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 34:
		//line waccparser.y:166
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].ident, Rhs: Binop{Left: parserS[parserpt-2].ident, Binary: MUL, Right: parserS[parserpt-2].ident, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 35:
		//line waccparser.y:169
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 36:
		//line waccparser.y:170
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 37:
		//line waccparser.y:171
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 38:
		//line waccparser.y:172
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 39:
		//line waccparser.y:173
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 40:
		//line waccparser.y:174
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 41:
		//line waccparser.y:175
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 42:
		//line waccparser.y:176
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 43:
		//line waccparser.y:177
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 44:
		//line waccparser.y:178
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 45:
		//line waccparser.y:182
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 46:
		//line waccparser.y:185
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 47:
		//line waccparser.y:189
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 48:
		//line waccparser.y:194
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 49:
		//line waccparser.y:196
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 50:
		//line waccparser.y:197
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 51:
		//line waccparser.y:199
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 52:
		//line waccparser.y:200
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 53:
		//line waccparser.y:201
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 54:
		//line waccparser.y:202
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 55:
		//line waccparser.y:203
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 56:
		//line waccparser.y:205
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 57:
		//line waccparser.y:206
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 58:
		//line waccparser.y:207
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 59:
		//line waccparser.y:208
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 60:
		//line waccparser.y:209
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		//line waccparser.y:210
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 62:
		//line waccparser.y:213
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:214
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		//line waccparser.y:215
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:216
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:217
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:218
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 68:
		//line waccparser.y:219
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:220
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 70:
		//line waccparser.y:221
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 71:
		//line waccparser.y:222
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 72:
		//line waccparser.y:223
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:224
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 74:
		//line waccparser.y:225
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 75:
		//line waccparser.y:226
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 76:
		//line waccparser.y:229
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 77:
		//line waccparser.y:231
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 78:
		//line waccparser.y:232
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 79:
		//line waccparser.y:233
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 80:
		//line waccparser.y:235
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 81:
		//line waccparser.y:237
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 82:
		//line waccparser.y:238
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 83:
		//line waccparser.y:241
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 84:
		//line waccparser.y:244
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 85:
		//line waccparser.y:245
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 86:
		//line waccparser.y:247
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 87:
		//line waccparser.y:249
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 88:
		//line waccparser.y:250
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 89:
		//line waccparser.y:251
		{
			parserVAL.pairelemtype = Pair
		}
	case 90:
		//line waccparser.y:253
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 91:
		//line waccparser.y:254
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 92:
		//line waccparser.y:255
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 93:
		//line waccparser.y:257
		{
			parserVAL.typedefinition = Int
		}
	case 94:
		//line waccparser.y:258
		{
			parserVAL.typedefinition = Bool
		}
	case 95:
		//line waccparser.y:259
		{
			parserVAL.typedefinition = Char
		}
	case 96:
		//line waccparser.y:260
		{
			parserVAL.typedefinition = String
		}
	case 97:
		//line waccparser.y:262
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
