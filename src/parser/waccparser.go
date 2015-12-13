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

//line waccparser.y:242

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 122,
	53, 82,
	-2, 79,
	-1, 123,
	53, 83,
	-2, 80,
}

const parserNprod = 90
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 415

var parserAct = []int{

	84, 4, 128, 22, 21, 83, 133, 71, 55, 132,
	169, 7, 24, 47, 64, 65, 66, 67, 68, 69,
	24, 70, 182, 121, 166, 36, 36, 36, 24, 25,
	31, 32, 77, 78, 162, 120, 35, 160, 165, 161,
	183, 166, 168, 24, 56, 96, 57, 58, 59, 34,
	158, 80, 61, 60, 81, 81, 94, 110, 111, 112,
	113, 114, 115, 116, 49, 50, 63, 86, 62, 46,
	93, 164, 92, 91, 73, 52, 54, 48, 51, 184,
	123, 122, 34, 179, 125, 34, 74, 130, 129, 75,
	134, 135, 136, 137, 138, 139, 140, 119, 142, 143,
	144, 145, 146, 147, 148, 149, 150, 151, 152, 153,
	154, 86, 87, 88, 31, 32, 90, 37, 33, 156,
	157, 56, 178, 57, 58, 59, 24, 24, 72, 61,
	60, 82, 163, 34, 76, 131, 36, 34, 44, 95,
	172, 49, 50, 63, 89, 62, 103, 105, 102, 104,
	159, 34, 52, 54, 48, 51, 6, 2, 23, 53,
	43, 171, 123, 122, 173, 134, 125, 176, 24, 129,
	177, 175, 34, 85, 180, 181, 127, 24, 20, 5,
	19, 24, 170, 8, 11, 12, 13, 14, 15, 16,
	17, 79, 3, 1, 18, 0, 0, 0, 0, 31,
	32, 26, 27, 28, 29, 30, 99, 101, 100, 97,
	98, 108, 109, 103, 105, 102, 104, 106, 107, 9,
	0, 26, 27, 28, 29, 30, 99, 101, 100, 174,
	0, 45, 0, 103, 105, 102, 104, 0, 10, 99,
	101, 100, 97, 98, 108, 109, 103, 105, 102, 104,
	106, 107, 126, 26, 27, 28, 29, 30, 0, 0,
	185, 99, 101, 100, 97, 98, 108, 109, 103, 105,
	102, 104, 106, 107, 26, 27, 28, 29, 124, 0,
	0, 0, 155, 99, 101, 100, 97, 98, 108, 109,
	103, 105, 102, 104, 106, 107, 0, 0, 0, 0,
	0, 0, 0, 167, 99, 101, 100, 97, 98, 108,
	109, 103, 105, 102, 104, 106, 107, 118, 0, 0,
	0, 0, 0, 0, 141, 0, 117, 0, 0, 0,
	0, 0, 0, 99, 101, 100, 97, 98, 108, 109,
	103, 105, 102, 104, 106, 107, 99, 101, 100, 97,
	98, 108, 109, 103, 105, 102, 104, 106, 107, 99,
	101, 100, 97, 98, 108, 109, 103, 105, 102, 104,
	106, 107, 99, 101, 100, 97, 98, 108, 0, 103,
	105, 102, 104, 106, 107, 41, 40, 42, 38, 39,
	99, 101, 100, 97, 98, 0, 0, 103, 105, 102,
	104, 106, 107, 44, 99, 101, 100, 97, 98, 0,
	0, 103, 105, 102, 104,
}
var parserPact = []int{

	153, -1000, -1000, 176, 113, -1000, -1000, -26, -1000, 60,
	350, 7, 14, 14, 14, 14, 14, 14, 14, 176,
	69, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	80, 14, 14, -1000, 176, -3, 76, 91, 59, 16,
	15, 13, -1, 86, 14, -1000, 85, 324, -1000, -1000,
	-1000, -1000, -1000, -1000, 85, -1000, 14, 14, 14, 14,
	14, 14, 14, -1000, 324, 324, 324, 324, 311, 298,
	92, -27, -1000, -1000, -1000, -1000, 249, 324, 324, -1000,
	196, 91, -1000, -1000, 324, -1000, -1000, 81, -53, 14,
	14, 14, 14, 14, 14, 14, 269, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	-1000, -1000, -1000, -1000, 191, 191, 226, 176, 176, -1000,
	-2, -8, -1000, -1000, 80, 83, 144, -19, -1000, -28,
	-1000, 14, 17, -17, 324, 324, 324, 324, 324, 324,
	248, -1000, 191, 191, 104, 104, 104, -1000, -1000, -1000,
	-1000, 369, 369, 355, 337, -1000, 26, -10, 249, 176,
	134, 228, -1000, 171, 14, -1000, 14, -1000, 176, -1000,
	66, 78, 176, -1000, 14, -34, 324, 23, -1000, -1000,
	74, 204, -1000, -1000, -1000, -1000,
}
var parserPgo = []int{

	0, 193, 192, 179, 156, 1, 5, 219, 0, 6,
	176, 2, 29, 173, 8, 160, 159, 4, 7, 3,
	158, 23,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 3, 10, 10, 11, 7,
	7, 7, 6, 6, 6, 6, 6, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 13, 9,
	9, 9, 14, 15, 15, 16, 12, 12, 20, 21,
	21, 21, 18, 18, 18, 17, 17, 17, 17, 19,
}
var parserR2 = []int{

	0, 4, 2, 0, 7, 8, 3, 1, 2, 1,
	1, 1, 1, 1, 1, 6, 5, 3, 1, 1,
	4, 3, 4, 4, 4, 4, 4, 2, 2, 2,
	2, 2, 2, 7, 5, 3, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 0, 2, 4, 3, 1, 2, 2, 6, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -5, -3, -4, -18, 7, -7,
	62, 8, 9, 10, 11, 12, 13, 14, 18, 4,
	2, -17, -19, -20, -14, -12, 25, 26, 27, 28,
	29, 23, 24, 5, 59, 62, 53, 57, 38, 39,
	36, 35, 37, -15, 53, -7, 62, -8, 63, 50,
	51, 64, 61, -16, 62, -14, 30, 32, 33, 34,
	39, 38, 54, 52, -8, -8, -8, -8, -8, -8,
	-5, -18, 59, 5, 17, 20, 54, -8, -8, -4,
	54, 57, 55, -6, -8, -13, -12, 21, 22, 53,
	57, 57, 57, 57, 57, 53, -8, 38, 39, 35,
	37, 36, 44, 42, 45, 43, 46, 47, 40, 41,
	-8, -8, -8, -8, -8, -8, -8, 15, 19, 5,
	62, -21, -17, -19, 29, -18, 56, -10, -11, -18,
	-6, 54, 62, -9, -8, -8, -8, -8, -8, -8,
	-8, 55, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, 56, -5, -5, 58, 6,
	56, 58, 62, -8, 54, 55, 58, 55, 16, 20,
	-21, -5, 6, -11, 58, -9, -8, -5, 56, 5,
	-5, -8, 56, 17, 5, 56,
}
var parserDef = []int{

	0, -2, 3, 0, 0, 2, 18, 0, 19, 0,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 82, 83, 84, 10, 11, 85, 86, 87, 88,
	0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 72, 0, 27, 9, 28, 40, 41,
	42, 43, 44, 45, 46, 47, 0, 0, 0, 0,
	0, 0, 0, 75, 29, 30, 31, 32, 0, 0,
	0, 0, 36, 37, 38, 39, 0, 76, 77, 17,
	0, 0, 89, 21, 12, 13, 14, 0, 0, 71,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	48, 49, 50, 51, 52, 53, 0, 0, 0, 35,
	0, 0, -2, -2, 81, 0, 0, 0, 7, 0,
	20, 0, 0, 0, 70, 22, 23, 24, 25, 26,
	0, 74, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 0, 0, 0, 0,
	0, 0, 8, 0, 71, 68, 0, 73, 0, 34,
	0, 0, 0, 6, 0, 0, 69, 0, 78, 4,
	0, 0, 16, 33, 5, 15,
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
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 23:
		//line waccparser.y:142
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 24:
		//line waccparser.y:143
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 25:
		//line waccparser.y:144
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 26:
		//line waccparser.y:145
		{
			parserVAL.stmt = Assignment{Lhs: parserS[parserpt-3].ident, Rhs: Binop{Left: parserS[parserpt-3].ident, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}, Pos: parserS[parserpt-3].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 27:
		//line waccparser.y:147
		{
			parserVAL.stmt = Read{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].assignlhs}
		}
	case 28:
		//line waccparser.y:148
		{
			parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 29:
		//line waccparser.y:149
		{
			parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 30:
		//line waccparser.y:150
		{
			parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 31:
		//line waccparser.y:151
		{
			parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 32:
		//line waccparser.y:152
		{
			parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr}
		}
	case 33:
		//line waccparser.y:153
		{
			parserVAL.stmt = If{Conditional: parserS[parserpt-5].expr, ThenStat: parserS[parserpt-3].stmts, ElseStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-6].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 34:
		//line waccparser.y:154
		{
			parserVAL.stmt = While{Conditional: parserS[parserpt-3].expr, DoStat: parserS[parserpt-1].stmts, Pos: parserS[parserpt-4].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 35:
		//line waccparser.y:155
		{
			parserVAL.stmt = Scope{StatList: parserS[parserpt-1].stmts, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 36:
		//line waccparser.y:156
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 37:
		//line waccparser.y:160
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 38:
		//line waccparser.y:163
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 39:
		//line waccparser.y:167
		{
			parserlex.Error("Syntax error : Invalid statement")
			parserVAL.stmt = nil
		}
	case 40:
		//line waccparser.y:172
		{
			parserVAL.expr = parserS[parserpt-0].integer
		}
	case 41:
		//line waccparser.y:174
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 42:
		//line waccparser.y:175
		{
			parserVAL.expr = parserS[parserpt-0].boolean
		}
	case 43:
		//line waccparser.y:177
		{
			parserVAL.expr = parserS[parserpt-0].character
		}
	case 44:
		//line waccparser.y:178
		{
			parserVAL.expr = parserS[parserpt-0].stringconst
		}
	case 45:
		//line waccparser.y:179
		{
			parserVAL.expr = parserS[parserpt-0].pairliter
		}
	case 46:
		//line waccparser.y:180
		{
			parserVAL.expr = parserS[parserpt-0].ident
		}
	case 47:
		//line waccparser.y:181
		{
			parserVAL.expr = parserS[parserpt-0].arrayelem
		}
	case 48:
		//line waccparser.y:183
		{
			parserVAL.expr = Unop{Unary: NOT, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 49:
		//line waccparser.y:184
		{
			parserVAL.expr = Unop{Unary: LEN, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 50:
		//line waccparser.y:185
		{
			parserVAL.expr = Unop{Unary: ORD, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 51:
		//line waccparser.y:186
		{
			parserVAL.expr = Unop{Unary: CHR, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 52:
		//line waccparser.y:187
		{
			parserVAL.expr = Unop{Unary: SUB, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 53:
		//line waccparser.y:188
		{
			parserVAL.expr = parserS[parserpt-0].expr
		}
	case 54:
		//line waccparser.y:191
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: PLUS, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 55:
		//line waccparser.y:192
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: SUB, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 56:
		//line waccparser.y:193
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MUL, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 57:
		//line waccparser.y:194
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: MOD, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 58:
		//line waccparser.y:195
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: DIV, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 59:
		//line waccparser.y:196
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 60:
		//line waccparser.y:197
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GT, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 61:
		//line waccparser.y:198
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: LTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 62:
		//line waccparser.y:199
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: GTE, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 63:
		//line waccparser.y:200
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: EQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 64:
		//line waccparser.y:201
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: NEQ, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 65:
		//line waccparser.y:202
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: AND, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 66:
		//line waccparser.y:203
		{
			parserVAL.expr = Binop{Left: parserS[parserpt-2].expr, Binary: OR, Right: parserS[parserpt-0].expr, Pos: parserS[parserpt-2].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 67:
		//line waccparser.y:204
		{
			parserVAL.expr = parserS[parserpt-1].expr
		}
	case 68:
		//line waccparser.y:207
		{
			parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs}
		}
	case 69:
		//line waccparser.y:209
		{
			parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)
		}
	case 70:
		//line waccparser.y:210
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}
		}
	case 71:
		//line waccparser.y:211
		{
			parserVAL.exprs = []Evaluation{}
		}
	case 72:
		//line waccparser.y:213
		{
			parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs: parserS[parserpt-0].exprs, Pos: parserS[parserpt-1].pos, FileText: &parserlex.(*Lexer).input}
		}
	case 73:
		//line waccparser.y:215
		{
			parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)
		}
	case 74:
		//line waccparser.y:216
		{
			parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}
		}
	case 75:
		//line waccparser.y:219
		{
			parserVAL.pairliter = PairLiter{}
		}
	case 76:
		//line waccparser.y:222
		{
			parserVAL.pairelem = PairElem{Fsnd: Fst, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 77:
		//line waccparser.y:223
		{
			parserVAL.pairelem = PairElem{Fsnd: Snd, Expr: parserS[parserpt-0].expr, Pos: parserS[parserpt-1].pos}
		}
	case 78:
		//line waccparser.y:225
		{
			parserVAL.typedefinition = PairType{FstType: parserS[parserpt-3].pairelemtype, SndType: parserS[parserpt-1].pairelemtype}
		}
	case 79:
		//line waccparser.y:227
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 80:
		//line waccparser.y:228
		{
			parserVAL.pairelemtype = parserS[parserpt-0].typedefinition
		}
	case 81:
		//line waccparser.y:229
		{
			parserVAL.pairelemtype = Pair
		}
	case 82:
		//line waccparser.y:231
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 83:
		//line waccparser.y:232
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 84:
		//line waccparser.y:233
		{
			parserVAL.typedefinition = parserS[parserpt-0].typedefinition
		}
	case 85:
		//line waccparser.y:235
		{
			parserVAL.typedefinition = Int
		}
	case 86:
		//line waccparser.y:236
		{
			parserVAL.typedefinition = Bool
		}
	case 87:
		//line waccparser.y:237
		{
			parserVAL.typedefinition = Char
		}
	case 88:
		//line waccparser.y:238
		{
			parserVAL.typedefinition = String
		}
	case 89:
		//line waccparser.y:240
		{
			parserVAL.typedefinition = ArrayType{Type: parserS[parserpt-2].typedefinition}
		}
	}
	goto parserstack /* stack new state and value */
}
