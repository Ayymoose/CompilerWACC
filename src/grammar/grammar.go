package grammar // Double check its the correct package

type Token int

type ItemType int

func (i ItemType) Len() int {
	return int(i)
}

type Type interface {
	getMainType() AllType
}

type BaseType int

const (
	BaseInt BaseType = iota
	BaseChar
	BaseBool
	BaseString
)

func (baseType BaseType) getMainType() AllType {
	switch baseType {
	case BaseInt:
		return TypeInt
	case BaseChar:
		return TypeChar
	case BaseBool:
		return TypeBool
	case BaseString:
		return TypeString
	}
	return -1
}

type AllType INT

const (
	TypeInt AllType = iota
	TypeChar
	TypeBool
	TypeString
	TypeArray
	TypePair
)

type ArrayType struct {
	Dimen    int
	BaseType BaseType
}
type PairType struct {
	FstType AllType
	SndType AllType
}

func (arrayType ArrayType) getMainType() AllType {
	return TypeArray
}

func (pairType PairType) getMainType() AllType {
	return TypePair
}

const Eof = -1
const (
	token_start ItemType = iota

	IDENTIFIER

	deliminators_start
	COMMA
	SEMICOLON
	COMMENT_START
	COMMENT_END
	COMMENT_LINE
	NESTED_COMMENTS
	//	UNDERSCORE
	deliminators_end

	escaped_char_start
	NULL_TERMINATOR
	TAB
	LINE_FEED
	FORM_FEED
	CARRIAGE_RETURN
	DOUBLE_QUOTE //Uncommented
	SINGLE_QUOTE
	BACKSLASH
	escaped_char_end

	reserved_word_start
	BEGIN
	END
	IS
	SKIP
	READ
	FREE
	RETURN
	EXIT
	PRINTLN
	PRINT

	IF
	THEN
	ELSE
	FI
	WHILE
	DONE
	DO
	NEWPAIR
	CALL
	FST
	SND
	NULL
	PAIR
	reserved_word_end
	BACKSPACE
	type_start
	INT
	BOOL
	CHAR
	STRING
	type_end

	NUMBER
	CHARLITER
	STRINGLITER

	DIGIT

	BOOLEAN
	PLAINTEXT
	EOF
	ERROR

	binary_op_start
	MULT
	DIV
	MOD
	ADD
	SUB // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct

	GTE
	GT

	STE //This too
	ST  //I think this should be LT (Less than)

	EQ
	NEQ
	AND
	OR
	binary_op_end

	unary_op_start
	NOT
	LEN
	ORD
	CHR
	NEG // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	unary_op_end

	bracket_type_start
	OPEN_SQUARE
	OPEN_ROUND
	CLOSE_SQUARE
	CLOSE_ROUND
	ASSIGNMENT
	bracket_type_end

	boolean_start
	TRUE
	FALSE
	boolean_end

	token_end
)

var EscapeChars = map[rune]ItemType{
	'0':  NULL_TERMINATOR,
	'b':  BACKSPACE,
	't':  TAB,
	'n':  LINE_FEED,
	'f':  FORM_FEED,
	'r':  CARRIAGE_RETURN,
	'\\': BACKSLASH,
	'\'': SINGLE_QUOTE,
	'"':  DOUBLE_QUOTE,
}

var DebugTokens = map[ItemType]string{
	token_start: "token_start",

	IDENTIFIER: "IDENTIFIER",

	deliminators_start: "deliminators_start",
	COMMA:              "COMMA",
	SEMICOLON:          "SEMICOLON",
	COMMENT_START:      "COMMENT_START",
	COMMENT_END:        "COMMENT_END",
	COMMENT_LINE:       "COMMENT_LINE",
	NESTED_COMMENTS:    "NESTED_COMMENTS",
	//	UNDERSCORE:         "UNDERSCORE",
	deliminators_end: "deliminators_end",

	escaped_char_start: "escaped_char_start",
	NULL_TERMINATOR:    "escaped_char_start",
	TAB:                "TAB",
	LINE_FEED:          "LINE_FEED",
	FORM_FEED:          "FORM_FEED",
	CARRIAGE_RETURN:    "CARRIAGE_RETURN",
	DOUBLE_QUOTE:       "DOUBLE_QUOTE", //Uncommented these
	SINGLE_QUOTE:       "SINGLE_QUOTE",
	BACKSLASH:          "BACKSLASH",
	escaped_char_end:   "escaped_char_end",

	reserved_word_start: "reserved_word_start",
	BEGIN:               "BEGIN",
	END:                 "END",
	IS:                  "IS",
	SKIP:                "SKIP",
	READ:                "READ",
	FREE:                "FREE",
	RETURN:              "RETURN",
	EXIT:                "EXIT",
	PRINTLN:             "PRINTLN",
	PRINT:               "PRINT",
	IF:                  "IF",
	THEN:                "THEN",
	ELSE:                "ELSE",
	FI:                  "FI",
	WHILE:               "WHILE",
	DONE:                "DONE",
	DO:                  "DO",
	NEWPAIR:             "NEWPAIR",
	CALL:                "CALL",
	FST:                 "FST",
	SND:                 "SND",
	NULL:                "NULL",
	PAIR:                "PAIR",
	reserved_word_end:   "reserved_word_end",
	BACKSPACE:           "BACKSPACE",
	type_start:          "type_start",
	INT:                 "INT",
	BOOL:                "BOOL",
	CHAR:                "CHAR",
	type_end:            "type_end",
	STRING:              "STRING",
	NUMBER:              "NUMBER",
	CHARLITER:           "CHARLITER",
	STRINGLITER:         "STRINGLITER",

	DIGIT:           "DIGIT",
	BOOLEAN:         "BOOLEAN",
	PLAINTEXT:       "PLAINTEXT",
	EOF:             "EOF",
	ERROR:           "ERROR",
	binary_op_start: "binary_op_start",
	MULT:            "MULT",
	DIV:             "DIV",
	MOD:             "MOD",
	ADD:             "ADD",
	SUB:             "SUB", // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct

	GTE:           "GTE",
	GT:            "GT",
	STE:           "STE",
	ST:            "ST",
	NEQ:           "NEQ",
	EQ:            "EQ",
	AND:           "AND",
	OR:            "OR",
	binary_op_end: "binary_op_end",

	unary_op_start: "unary_op_start",
	NOT:            "NOT",
	LEN:            "LEN",
	ORD:            "ORD",
	CHR:            "CHR",
	NEG:            "NEG", // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	unary_op_end:   "unary_op_end",

	bracket_type_start: "bracket_type_start",
	OPEN_SQUARE:        "OPEN_SQUARE",
	OPEN_ROUND:         "OPEN_ROUND",
	CLOSE_SQUARE:       "CLOSE_SQUARE",
	CLOSE_ROUND:        "CLOSE_ROUND",
	ASSIGNMENT:         "ASSIGNMENT",
	bracket_type_end:   "bracket_type_end",

	boolean_start: "boolean_start",
	TRUE:          "TRUE",
	FALSE:         "FALSE",
	boolean_end:   "boolean_end",

	token_end: "token_end",
}

var Token_keyword_strings = map[ItemType]string{
	//	IDENTIFIER: "identifier", // Place holder
	BEGIN:   "begin",
	END:     "end",
	IS:      "is",
	SKIP:    "skip",
	READ:    "read",
	FREE:    "free",
	RETURN:  "return",
	EXIT:    "exit",
	PRINTLN: "println",
	PRINT:   "print",
	IF:      "if",
	THEN:    "then",
	ELSE:    "else",
	FI:      "fi",
	WHILE:   "while",
	DONE:    "done",
	DO:      "do",
	NEWPAIR: "newpair",
	CALL:    "call",
	FST:     "fst",
	SND:     "snd",
	NULL:    "null",
	INT:     "int",
	BOOL:    "bool",
	CHAR:    "char",
	STRING:  "string",
	PAIR:    "pair",
	LEN:     "len",
	ORD:     "ord",
	CHR:     "chr",
	TRUE:    "true",
	FALSE:   "false",
}

var Token_strings = map[ItemType]string{

	COMMA:     ",",
	SEMICOLON: ";",
	//COMMENT_START: "",
	//	COMMENT_END:   "",
	COMMENT_LINE: "#",
	//	NESTED_COMMENTS: false,
	//	UNDERSCORE:      "_",
	NULL_TERMINATOR: "\\0",
	BACKSPACE:       "\b",
	TAB:             "\t",
	LINE_FEED:       "\n",
	FORM_FEED:       "\f",
	CARRIAGE_RETURN: "\r",
	DOUBLE_QUOTE:    "\"", //Uncommented these because I needed to use it @Ayman
	SINGLE_QUOTE:    "'",
	BACKSLASH:       "\\",
	NEG:             "-", // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	MULT:            "*",
	DIV:             "/",
	MOD:             "%",
	ADD:             "+",
	SUB:             "-", // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct

	GTE: ">=",

	STE:          "<=",
	GT:           ">",
	ST:           "<",
	EQ:           "==", // BETTER NAMES??
	NEQ:          "!=", // BETTER NAMES?
	NOT:          "!",
	AND:          "&&",
	OR:           "||",
	OPEN_SQUARE:  "[",
	OPEN_ROUND:   "(",
	CLOSE_SQUARE: "]",
	CLOSE_ROUND:  ")",
	ASSIGNMENT:   "=",
}

func (token ItemType) IsEscapedChar() bool {
	return token > escaped_char_start && token < escaped_char_end
}

func (token ItemType) IsDeliminator() bool {
	return token > deliminators_start && token < deliminators_end
}

func (token ItemType) IsReservedWord() bool {
	return token > reserved_word_start && token < reserved_word_end
}

func (token ItemType) IsType() bool {
	return token > type_start && token < type_end
}

func (token ItemType) IsUnaryOp() bool {
	return token > unary_op_start && token < unary_op_end
}

func (token ItemType) IsBinaryOp() bool {
	return token > binary_op_start && token < binary_op_end
}

func (token ItemType) IsBracketType() bool {
	return token > bracket_type_start && token < bracket_type_end
}

func (token ItemType) IsBoolean() bool {
	return token > boolean_start && token < boolean_end
}

/*
func lookUp(str string) Token {
	for t, s := range token_strings {
		if s == str {
			return t
		}
	}
	return
}
*/
