package lexer // Double check its the correct package

type Token int

type itemType int

const (
	token_start itemType = iota

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
	FLASE
	boolean_end

	token_end
)

var token_keyword_strings = map[itemType]string{
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
	NOT:        "!",
	LEN:        "len",
	ORD:        "ord",
	CHR:        "chr",
	TRUE:       "true",
	FLASE:      "false",
}

var token_strings = map[itemType]string{
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

func (token itemType) isEscapedChar() bool {
	return token > escaped_char_start && token < escaped_char_end
}

func (token itemType) isDeliminator() bool {
	return token > deliminators_start && token < deliminators_end
}

func (token itemType) isReservedWord() bool {
	return token > reserved_word_start && token < reserved_word_end
}

func (token itemType) isType() bool {
	return token > type_start && token < type_end
}

func (token itemType) isUnaryOp() bool {
	return token > unary_op_start && token < unary_op_end
}

func (token itemType) isBinaryOp() bool {
	return token > binary_op_start && token < binary_op_end
}

func (token itemType) isBracketType() bool {
	return token > bracket_type_start && token < bracket_type_end
}

func (token itemType) isBoolean() bool {
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
