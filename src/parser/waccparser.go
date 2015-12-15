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

const BEGIN = 57346
const END = 57347
const CLASS = 57348
const IS = 57349
const SKIP = 57350
const READ = 57351
const FREE = 57352
const RETURN = 57353
const EXIT = 57354
const PRINT = 57355
const PRINTLN = 57356
const IF = 57357
const THEN = 57358
const ELSE = 57359
const FI = 57360
const WHILE = 57361
const DO = 57362
const DONE = 57363
const FOR = 57364
const NEWPAIR = 57365
const CALL = 57366
const FST = 57367
const SND = 57368
const INT = 57369
const BOOL = 57370
const CHAR = 57371
const STRING = 57372
const PAIR = 57373
const NOT = 57374
const NEG = 57375
const LEN = 57376
const ORD = 57377
const CHR = 57378
const MUL = 57379
const DIV = 57380
const MOD = 57381
const PLUS = 57382
const SUB = 57383
const AND = 57384
const OR = 57385
const GT = 57386
const GTE = 57387
const LT = 57388
const LTE = 57389
const EQ = 57390
const NEQ = 57391
const POSITIVE = 57392
const NEGATIVE = 57393
const TRUE = 57394
const FALSE = 57395
const NULL = 57396
const OPENSQUARE = 57397
const OPENROUND = 57398
const CLOSESQUARE = 57399
const CLOSEROUND = 57400
const ASSIGNMENT = 57401
const COMMA = 57402
const SEMICOLON = 57403
const ERROR = 57404
const STRINGCONST = 57405
const IDENTIFIER = 57406
const INTEGER = 57407
const CHARACTER = 57408

var parserToknames = []string{
	"BEGIN",
	"END",
	"CLASS",
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
	"FOR",
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
