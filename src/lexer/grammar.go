package lexer // Double check its the correct package

type Token int

type ItemType int

const eof = -1
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
	UNDERSCORE
	deliminators_end

	escaped_char_start
	NULL_TERMINATOR
	TAB
	LINE_FEED
	FORM_FEED
	CARRIAGE_RETURN
	DOUBLE_QUOTE
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
	PRINT
	PRINTLN
	IF
	THEN
	ELSE
	FI
	WHILE
	DO
	DONE
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
	CHARACTER
	STRINGVALUE
	BOOLEAN
	PLAINTEXT
	EOF
	ERROR

	unary_op_start
	NOT
	LEN
	ORD
	CHR
	NEG // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	unary_op_end

	binary_op_start
	MULT
	DIV
	MOD
	ADD
	SUB // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct
	GT
	GTE
	ST
	STE
	EQ
	NEQ
	AND
	OR
	binary_op_end

	bracket_type_start
	OPEN_SQUARE
	OPEN_ROUND
	CLOSE_SQUARE
	CLOSE_ROUND
	bracket_type_end

	boolean_start
	TRUE
	FALSE
	boolean_end

	token_end
)

var debugTokens = map[ItemType]string{
	token_start: "token_start",

	IDENTIFIER: "IDENTIFIER",

	deliminators_start: "deliminators_start",
	COMMA:              "COMMA",
	SEMICOLON:          "SEMICOLON",
	COMMENT_START:      "COMMENT_START",
	COMMENT_END:        "COMMENT_END",
	COMMENT_LINE:       "COMMENT_LINE",
	NESTED_COMMENTS:    "NESTED_COMMENTS",
	UNDERSCORE:         "UNDERSCORE",
	deliminators_end:   "deliminators_end",

	escaped_char_start: "escaped_char_start",
	NULL_TERMINATOR:    "escaped_char_start",
	TAB:                "TAB",
	LINE_FEED:          "LINE_FEED",
	FORM_FEED:          "FORM_FEED",
	CARRIAGE_RETURN:    "CARRIAGE_RETURN",
	DOUBLE_QUOTE:       "DOUBLE_QUOTE",
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
	PRINT:               "PRINT",
	PRINTLN:             "PRINTLN",
	IF:                  "IF",
	THEN:                "THEN",
	ELSE:                "ELSE",
	FI:                  "FI",
	WHILE:               "WHILE",
	DO:                  "DO",
	DONE:                "DONE",
	NEWPAIR:             "NEWPAIR",
	CALL:                "CALL",
	FST:                 "FST",
	SND:                 "SND",
	NULL:                "NULL",
	PAIR:                "PAI",
	reserved_word_end:   "reserved_word_end",
	BACKSPACE:           "BACKSPACE",
	type_start:          "type_start",
	INT:                 "INT",
	BOOL:                "BOOL",
	CHAR:                "CHAR",
	type_end:            "type_end",
	STRING:              "STRING",
	NUMBER:              "NUMBER",
	CHARACTER:           "CHARACTER",
	STRINGVALUE:         "STRINGVALUE",
	BOOLEAN:             "BOOLEAN",
	PLAINTEXT:           "PLAINTEXT",
	EOF:                 "EOF",
	ERROR:               "ERROR",

	unary_op_start: "unary_op_start",
	NOT:            "NOT",
	LEN:            "LEN",
	ORD:            "ORD",
	CHR:            "CHR",
	NEG:            "NEG", // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	unary_op_end:   "unary_op_end",

	binary_op_start: "binary_op_start",
	MULT:            "MULT",
	DIV:             "DIV",
	MOD:             "MOD",
	ADD:             "ADD",
	SUB:             "SUB", // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct
	GT:              "GT",
	GTE:             "GTE",
	ST:              "ST",
	STE:             "STE",
	NEQ:             "NEQ",
	EQ:              "EQ",
	AND:             "AND",
	OR:              "OR",
	binary_op_end:   "binary_op_end",

	bracket_type_start: "bracket_type_start",
	OPEN_SQUARE:        "OPEN_SQUARE",
	OPEN_ROUND:         "OPEN_ROUND",
	CLOSE_SQUARE:       "CLOSE_SQUARE",
	CLOSE_ROUND:        "CLOSE_ROUND",
	bracket_type_end:   "bracket_type_end",

	boolean_start: "boolean_start",
	TRUE:          "TRUE",
	FALSE:         "FALSE",
	boolean_end:   "boolean_end",

	token_end: "token_end",
}

var token_keyword_strings = map[ItemType]string{
	IDENTIFIER: "identifier", // Place holder
	BEGIN:      "begin",
	END:        "end",
	IS:         "is",
	SKIP:       "skip",
	READ:       "read",
	FREE:       "free",
	RETURN:     "return",
	EXIT:       "exit",
	PRINT:      "print",
	PRINTLN:    "println",
	IF:         "if",
	THEN:       "then",
	ELSE:       "else",
	FI:         "fi",
	WHILE:      "while",
	DONE:       "done",
	DO:         "do",
	NEWPAIR:    "newpair",
	CALL:       "call",
	FST:        "fst",
	SND:        "snd",
	NULL:       "null",
	INT:        "int",
	BOOL:       "bool",
	CHAR:       "char",
	STRING:     "string",
	PAIR:       "pair",
	LEN:        "len",
	ORD:        "ord",
	CHR:        "chr",
	TRUE:       "true",
	FALSE:      "false",
}

var token_strings = map[ItemType]string{
	NOT:       "!",
	COMMA:     ",",
	SEMICOLON: ";",
	//COMMENT_START: "",
	//	COMMENT_END:   "",
	COMMENT_LINE: "#",
	//	NESTED_COMMENTS: false,
	UNDERSCORE:      "_",
	NULL_TERMINATOR: "\\0",
	BACKSPACE:       "\b",
	TAB:             "\t",
	LINE_FEED:       "\n",
	FORM_FEED:       "\f",
	CARRIAGE_RETURN: "\r",
	DOUBLE_QUOTE:    "\"",
	SINGLE_QUOTE:    "\\'",
	BACKSLASH:       "\\",
	NEG:             "-", // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
	MULT:            "*",
	DIV:             "/",
	MOD:             "%",
	ADD:             "+",
	SUB:             "-", // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct
	GT:              ">",
	GTE:             ">=",
	ST:              "<",
	STE:             "<=",
	EQ:              "==", // BETTER NAMES??
	NEQ:             "!=", // BETTER NAMES?
	AND:             "&&",
	OR:              "||",
	OPEN_SQUARE:     "[",
	OPEN_ROUND:      "(",
	CLOSE_SQUARE:    "]",
	CLOSE_ROUND:     ")",
}

func (token ItemType) isEscapedChar() bool {
	return token > escaped_char_start && token < escaped_char_end
}

func (token ItemType) isDeliminator() bool {
	return token > deliminators_start && token < deliminators_end
}

func (token ItemType) isReservedWord() bool {
	return token > reserved_word_start && token < reserved_word_end
}

func (token ItemType) isType() bool {
	return token > type_start && token < type_end
}

func (token ItemType) isUnaryOp() bool {
	return token > unary_op_start && token < unary_op_end
}

func (token ItemType) isBinaryOp() bool {
	return token > binary_op_start && token < binary_op_end
}

func (token ItemType) isBracketType() bool {
	return token > bracket_type_start && token < bracket_type_end
}

func (token ItemType) isBoolean() bool {
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
