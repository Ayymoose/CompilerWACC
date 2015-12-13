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
	stmt           Statement
	stmts          []Statement
	assignrhs      Evaluation
	assignlhs      Evaluation
	expr           Evaluation
	exprs          []Evaluation
	params         []Param
	param          Param
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
	"INTEGER",
	"CHARACTER",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line waccparser.y:241

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 114,
	53, 78,
	-2, 75,
	-1, 115,
	53, 79,
	-2, 76,
}

const parserNprod = 86
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 425

var parserAct = []int{

	80, 4, 120, 22, 21, 79, 125, 67, 51, 31,
	32, 7, 24, 43, 60, 61, 62, 63, 64, 65,
	24, 66, 36, 113, 124, 25, 36, 36, 24, 171,
	146, 150, 73, 74, 157, 112, 35, 170, 156, 154,
	77, 88, 153, 24, 148, 154, 149, 76, 42, 86,
	77, 37, 166, 102, 103, 104, 105, 106, 107, 108,
	69, 78, 152, 82, 72, 123, 172, 38, 36, 167,
	40, 34, 70, 34, 111, 71, 115, 114, 33, 87,
	117, 34, 40, 122, 121, 160, 126, 127, 128, 2,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, 82, 83, 84, 31, 32, 6, 147,
	23, 144, 145, 52, 68, 53, 54, 55, 24, 24,
	34, 57, 56, 34, 151, 49, 39, 81, 34, 119,
	9, 5, 34, 45, 46, 59, 85, 58, 95, 97,
	94, 96, 41, 75, 48, 50, 44, 47, 3, 159,
	115, 114, 161, 126, 117, 164, 24, 121, 165, 163,
	1, 0, 168, 169, 0, 24, 20, 0, 19, 24,
	158, 8, 11, 12, 13, 14, 15, 16, 17, 0,
	0, 0, 18, 0, 0, 0, 0, 31, 32, 26,
	27, 28, 29, 30, 0, 0, 52, 0, 53, 54,
	55, 0, 0, 0, 57, 56, 26, 27, 28, 29,
	30, 26, 27, 28, 29, 116, 45, 46, 59, 0,
	58, 0, 0, 0, 0, 0, 10, 48, 50, 44,
	47, 91, 93, 92, 89, 90, 100, 101, 95, 97,
	94, 96, 98, 99, 26, 27, 28, 29, 30, 0,
	0, 0, 0, 0, 162, 91, 93, 92, 89, 90,
	100, 101, 95, 97, 94, 96, 98, 99, 0, 0,
	0, 0, 0, 0, 0, 118, 173, 91, 93, 92,
	89, 90, 100, 101, 95, 97, 94, 96, 98, 99,
	0, 0, 0, 0, 0, 0, 0, 0, 143, 91,
	93, 92, 89, 90, 100, 101, 95, 97, 94, 96,
	98, 99, 0, 0, 0, 0, 0, 0, 0, 155,
	91, 93, 92, 89, 90, 100, 101, 95, 97, 94,
	96, 98, 99, 110, 0, 0, 91, 93, 92, 0,
	129, 0, 109, 95, 97, 94, 96, 0, 0, 91,
	93, 92, 89, 90, 100, 101, 95, 97, 94, 96,
	98, 99, 91, 93, 92, 89, 90, 100, 101, 95,
	97, 94, 96, 98, 99, 91, 93, 92, 89, 90,
	100, 101, 95, 97, 94, 96, 98, 99, 91, 93,
	92, 89, 90, 100, 0, 95, 97, 94, 96, 98,
	99, 91, 93, 92, 89, 90, 0, 0, 95, 97,
	94, 96, 98, 99, 91, 93, 92, 89, 90, 0,
	0, 95, 97, 94, 96,
}
var parserPact = []int{

	85, -1000, -1000, 164, 73, -1000, -1000, -26, -1000, -6,
	29, -14, 166, 166, 166, 166, 166, 166, 166, 164,
	55, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	10, 166, 166, -1000, 164, -7, 6, 83, -8, 26,
	166, -1000, 17, 340, -1000, -1000, -1000, -1000, -1000, -1000,
	17, -1000, 166, 166, 166, 166, 166, 166, 166, -1000,
	340, 340, 340, 340, 327, 314, 69, -27, -1000, -1000,
	-1000, -1000, 186, 340, 340, -1000, 219, 83, -1000, -1000,
	340, -1000, -1000, 11, -38, 166, 166, 166, 285, 166,
	166, 166, 166, 166, 166, 166, 166, 166, 166, 166,
	166, 166, -1000, -1000, -1000, -1000, 301, 301, 242, 164,
	164, -1000, -17, -28, -1000, -1000, 10, 15, 103, -12,
	-1000, -31, -1000, 166, 8, -13, 340, 340, 264, -1000,
	301, 301, 96, 96, 96, -1000, -1000, -1000, -1000, 379,
	379, 366, 353, -1000, 22, 14, 186, 164, 79, 181,
	-1000, 196, 166, -1000, 166, -1000, 164, -1000, -4, 64,
	164, -1000, 166, -19, 340, 12, -1000, -1000, 61, 220,
	-1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 160, 148, 131, 108, 1, 5, 130, 0, 6,
	129, 2, 25, 127, 8, 126, 125, 4, 7, 3,
	110, 23,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 3, 10, 10, 11, 7,
	7, 7, 6, 6, 6, 6, 6, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 13, 9, 9, 9, 14, 15,
	15, 16, 12, 12, 20, 21, 21, 21, 18, 18,
	18, 17, 17, 17, 17, 19,
}
var parserR2 = []int{

	0, 4, 2, 0, 7, 8, 3, 1, 2, 1,
	1, 1, 1, 1, 1, 6, 5, 3, 1, 1,
	4, 3, 4, 2, 2, 2, 2, 2, 2, 7,
	5, 3, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 0, 2, 4,
	3, 1, 2, 2, 6, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -5, -3, -4, -18, 7, -7,
	62, 8, 9, 10, 11, 12, 13, 14, 18, 4,
	2, -17, -19, -20, -14, -12, 25, 26, 27, 28,
	29, 23, 24, 5, 59, 62, 53, 57, 38, -15,
	53, -7, 62, -8, 63, 50, 51, 64, 61, -16,
	62, -14, 30, 32, 33, 34, 39, 38, 54, 52,
	-8, -8, -8, -8, -8, -8, -5, -18, 59, 5,
	17, 20, 54, -8, -8, -4, 54, 57, 55, -6,
	-8, -13, -12, 21, 22, 53, 57, 53, -8, 38,
	39, 35, 37, 36, 44, 42, 45, 43, 46, 47,
	40, 41, -8, -8, -8, -8, -8, -8, -8, 15,
	19, 5, 62, -21, -17, -19, 29, -18, 56, -10,
	-11, -18, -6, 54, 62, -9, -8, -8, -8, 55,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, 56, -5, -5, 58, 6, 56, 58,
	62, -8, 54, 55, 58, 55, 16, 20, -21, -5,
	6, -11, 58, -9, -8, -5, 56, 5, -5, -8,
	56, 17, 5, 56,
}
var parserDef = []int{

	0, -2, 3, 0, 0, 2, 18, 0, 19, 0,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 78, 79, 80, 10, 11, 81, 82, 83, 84,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 68,
	0, 23, 9, 24, 36, 37, 38, 39, 40, 41,
	42, 43, 0, 0, 0, 0, 0, 0, 0, 71,
	25, 26, 27, 28, 0, 0, 0, 0, 32, 33,
	34, 35, 0, 72, 73, 17, 0, 0, 85, 21,
	12, 13, 14, 0, 0, 67, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 44, 45, 46, 47, 48, 49, 0, 0,
	0, 31, 0, 0, -2, -2, 77, 0, 0, 0,
	7, 0, 20, 0, 0, 0, 66, 22, 0, 70,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 62, 63, 0, 0, 0, 0, 0, 0,
	8, 0, 67, 64, 0, 69, 0, 30, 0, 0,
	0, 6, 0, 0, 65, 0, 74, 4, 0, 0,
	16, 29, 5, 15,
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
	62, 63, 64,
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
		//line waccparser.y:98
		{
			parserlex.(*Lexer).prog = &Program{FunctionList: parserS[parserpt-2].functions, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 2:
		//line waccparser.y:102
		{
			parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)
		}
	case 3:
		//line waccparser.y:103
		{
			parserVAL.functions = []*Function{}
		}
	case 4:
		//line waccparser.y:106
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-5].ident, ReturnType: parserS[parserpt-6].typedefinition, StatList: parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 5:
		//line waccparser.y:112
		{
			if !checkStats(parserS[parserpt-1].stmts) {
				parserlex.Error("Missing return statement")
			}
			parserVAL.function = &Function{Ident: parserS[parserpt-6].ident, ReturnType: parserS[parserpt-7].typedefinition, StatList: parserS[parserpt-1].stmts, ParameterTypes: parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText: &parserlex.(*Lexer).input}
		}
	case 6:
		//line waccparser.y:118
		{
			parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)
		}
	case 7:
		//line waccparser.y:119
		{
			parserVAL.params = []Param{parserS[parserpt-0].param}
		}
	case 8:
		//line waccparser.y:121
		{
			parserVAL.param = Param{ParamType: parserS[parserpt-1].typedefinition, Ident: parserS[parserpt-0].ident}
		}
	case 9:
		//line waccparser.y:123
		{
			parserVAL.assignlhs = parserS[parserpt-0].ident
		}
	case 10:
		//line waccparser.y:124
		{
			parserVAL.assignlhs = parserS[parserpt-0].arrayelem
		}
	case 11:
		//line waccparser.y:125
		{
			parserVAL.assignlhs = parserS[parserpt-0].pairelem
		}
	case 12:
		//line waccparser.y:127
		{
			parserVAL.assignrhs = parserS[parserpt-0].expr
		}
	case 13:
		//line waccparser.y:128
		{
			parserVAL.assignrhs = parserS[parserpt-0].arrayliter
		}
	case 14:
		//line waccparser.y:129
		{
			parserVAL.assignrhs = parserS[parserpt-0].pairelem
		}
	case 15:
		//line waccparser.y:130
		{
			parserVAL.assignrhs = NewPair{FstExpr: parserS[parserpt-3].expr, SndExpr: parserS[parserpt-1].expr, Pos: parserS[parserpt-5].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 16:
		//line waccparser.y:131
		{
			parserVAL.assignrhs = Call{Ident: parserS[parserpt-3].ident, ParamList: parserS[parserpt-1].exprs, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 17:
		//line waccparser.y:133
		{
			parserVAL.stmts = append(parserS[parserpt-2].stmts, parserS[parserpt-0].stmt)
		}
	case 18:
		//line waccparser.y:134
		{
			parserVAL.stmts = []Statement{parserS[parserpt-0].stmt}
		}
	case 19:
		//line waccparser.y:137
		{
			parserVAL.stmt = Skip{Pos: parserS[parserpt-0].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 20:
		//line waccparser.y:138
		{
			parserVAL.stmt = Declare{DecType: parserS[parserpt-3].typedefinition, Lhs: parserS[parserpt-2].ident, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 21:
		//line waccparser.y:139
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-2].assignlhs, Rhs: parserS[parserpt-0].assignrhs, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 22:
		//line waccparser.y:141
		{
			parserVAL.stmt = Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:146
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 24:
		//line waccparser.y:147
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 25:
		//line waccparser.y:148
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 26:
		//line waccparser.y:149
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 27:
		//line waccparser.y:150
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 28:
		//line waccparser.y:151
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 29:
		//line waccparser.y:152
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 30:
		//line waccparser.y:153
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 31:
		//line waccparser.y:154
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 32:
		//line waccparser.y:155
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 33:
		//line waccparser.y:159
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 34:
		//line waccparser.y:162
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 35:
		//line waccparser.y:166
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 36:
		//line waccparser.y:171
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 37:
		//line waccparser.y:173
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 38:
		//line waccparser.y:174
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 39:
		//line waccparser.y:176
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 40:
		//line waccparser.y:177
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 41:
		//line waccparser.y:178
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 42:
		//line waccparser.y:179
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 43:
		//line waccparser.y:180
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 44:
		//line waccparser.y:182
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 45:
		//line waccparser.y:183
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 46:
		//line waccparser.y:184
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 47:
		//line waccparser.y:185
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 48:
		//line waccparser.y:186
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:187
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 50:
		//line waccparser.y:190
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:191
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:192
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:193
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 54:
		//line waccparser.y:194
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 55:
		//line waccparser.y:195
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 56:
		//line waccparser.y:196
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 57:
		//line waccparser.y:197
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 58:
		//line waccparser.y:198
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 59:
		//line waccparser.y:199
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 60:
		//line waccparser.y:200
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		//line waccparser.y:201
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		//line waccparser.y:202
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:203
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 64:
		//line waccparser.y:206
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 65:
		//line waccparser.y:208
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 66:
		//line waccparser.y:209
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 67:
		//line waccparser.y:210
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 68:
		//line waccparser.y:212
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 69:
		//line waccparser.y:214
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 70:
		//line waccparser.y:215
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 71:
		//line waccparser.y:218
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 72:
		//line waccparser.y:221
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 73:
		//line waccparser.y:222
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 74:
		//line waccparser.y:224
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 75:
		//line waccparser.y:226
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 76:
		//line waccparser.y:227
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 77:
		//line waccparser.y:228
		{
			parserVAL.pairelemtype = Pair
		}
	case 78:
		//line waccparser.y:230
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 79:
		//line waccparser.y:231
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 80:
		//line waccparser.y:232
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 81:
		//line waccparser.y:234
		{
			parserVAL.typedefinition = Int
		}
	case 82:
		//line waccparser.y:235
		{
			parserVAL.typedefinition = Bool
		}
	case 83:
		//line waccparser.y:236
		{
			parserVAL.typedefinition = Char
		}
	case 84:
		//line waccparser.y:237
		{
			parserVAL.typedefinition = String
		}
	case 85:
		//line waccparser.y:239
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
