package main // Double check its the correct package

type tokenType int

const (
	tokenType_start tokenType = iota

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
	BACKSPACE
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

	tokenType_end
)

var token_strings = map[tokenType]string{
	IDENTIFIER:      "identifier", // Place holder
	COMMA:           ",",
	SEMICOLON:       ";",
	COMMENT_START:   "",
	COMMENT_END:     "",
	COMMENT_LINE:    "#",
	NESTED_COMMENTS: "false",
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
	BEGIN:           "begin",
	END:             "end",
	IS:              "is",
	SKIP:            "skip",
	READ:            "read",
	FREE:            "free",
	RETURN:          "return",
	EXIT:            "exit",
	PRINT:           "print",
	PRINTLN:         "println",
	IF:              "if",
	THEN:            "then",
	ELSE:            "else",
	FI:              "fi",
	WHILE:           "while",
	DO:              "do",
	DONE:            "done",
	NEWPAIR:         "newpair",
	CALL:            "call",
	FST:             "fst",
	SND:             "snd",
	NULL:            "null",
	INT:             "int",
	BOOL:            "bool",
	CHAR:            "char",
	STRING:          "string",
	PAIR:            "pair",
	NOT:             "!",
	LEN:             "len",
	ORD:             "ord",
	CHR:             "chr",
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
	TRUE:            "true",
	FLASE:           "false",
}

func (tokenTyp tokenType) isEscapedChar() bool {
	return tokenTyp > escaped_char_start && tokenTyp < escaped_char_end
}

func (tokenTyp tokenType) isDeliminator() bool {
	return tokenTyp > deliminators_start && tokenTyp < deliminators_end
}

func (tokenTyp tokenType) isReservedWord() bool {
	return tokenTyp > reserved_word_start && tokenTyp < reserved_word_end
}

func (tokenTyp tokenType) isType() bool {
	return tokenTyp > type_start && tokenTyp < type_end
}

func (tokenTyp tokenType) isUnaryOp() bool {
	return tokenTyp > unary_op_start && tokenTyp < unary_op_end
}

func (tokenTyp tokenType) isBinaryOp() bool {
	return tokenTyp > binary_op_start && tokenTyp < binary_op_end
}

func (tokenTyp tokenType) isBracketType() bool {
	return tokenTyp > bracket_type_start && tokenTyp < bracket_type_end
}

func (tokenTyp tokenType) isBoolean() bool {
	return tokenTyp > boolean_start && tokenTyp < boolean_end
}

func lookUp(str string) tokenType {
	for t, s := range tokenType_strings {
		if s == str {
			return t
		}
	}
	return IDENTIFIER
}
