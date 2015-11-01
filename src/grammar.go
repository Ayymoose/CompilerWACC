package main // Double check its the correct package

type Token int

const (
  token_start Token  = iota

  COMMA
  SEMICOLON
  COMMENT_START
  COMMENT_END
  COMMENT_LINE
  NESTED_COMMENTS
  SPACE  // do we need this???
  UNDERSCORE


  escaped_char_start  // NOT SURE IF I SHOULD DEFINE IT AS A BACKSLASH FOLLOWED BY THE ESCAPED
  NULL_TERMINATOR     // CHARACTER (I.E. RECURSIVLY LIKE IN THE SPEC) OR HAVE EACH ESCAPED CHARACTER
  BACKSPACE           // INCLUDE ITS BACKSLASH
  TAB
  LINE_FEED
  FORM_FEED
  CARRIAGE_RETURN
  DOUBLE_QUOTE
  SINGLE_QUOTE
  BACKSLASH         // NOT TOO NEAT BACKSLASH BACKSLASH BASICALLY
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

  INT
  BOOL
  CHAR
  STRING


  PAIR


  reserved_word_end


  unary_op_start
  NOT
  LEN
  ORD
  CHR
  NEG // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
  unary_op_end

  binary_op_start
  MULTIPLY
  DIVIDE
  MOD
  ADD
  MINUS // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct
  GT
  GTE
  ST
  STE
  EQUAL // BETTER NAMES??
  NEQUAL // BETTER NAMES?
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

var token_strings = map[Token]string{
  COMMA : ",",
  SEMICOLON : ";",
  COMMENT_START : "",
  COMMENT_END : "",
  COMMENT_LINE : "#",
  NESTED_COMMENTS : false,
  SPACE : " ", // remove if not needed
  UNDERSCORE : "_",
  NULL_TERMINATOR : "\0",
  BACKSPACE : "\b",
  TAB : "\t",
  LINE_FEED : "\n",
  FORM_FEED : "\f",
  CARRIAGE_RETURN : "\r",
  DOUBLE_QUOTE : "\"",
  SINGLE_QUOTE : "\'",
  BACKSLASH : "\\",
  BEGIN : "begin",
  END : "end",
  IS : "is",
  SKIP : "skip",
  READ : "read",
  FREE : "free",
  RETURN : "return",
  EXIT : "exit",
  PRINT : "print",
  PRINTLN : "println",
  IF : "if",
  THEN : "then",
  ELSE : "else",
  FI : "fi",
  WHILE : "while",
  DO : "do",
  DONE : "done",
  NEWPAIR : "newpair",
  CALL : "call",
  FST : "fst",
  SND : "snd",
  NULL : "null",
  INT : "int",
  BOOL : "bool",
  CHAR : "char",
  STRING : "string",
  PAIR : "pair",
  NOT : "!",
  LEN : "len",
  ORD : "ord"
  CHR : "chr",
  NEG : "- ",     // THIS IS FOR '-' WHICH ALSO EXISTS IN BINARY_OP ASWELL. NOT SURE OF THE BEST NAME
  MULTIPLY : "*",
  DIVIDE : "/",
  MOD : "%",
  ADD : "+",
  MINUS : "-",  // '-' ALSO SAME SYMBOL -> NOT SURE IF ENTIRELY correct
  GT : ">",
  GTE : ">=",
  ST : "<",
  STE : "<=",
  EQUAL : "==",  // BETTER NAMES??
  NEQUAL : "!=", // BETTER NAMES?
  AND : "&&",
  OR : "||",
  OPEN_SQUARE : "[",
  OPEN_ROUND : "(",
  CLOSE_SQUARE : "]",
  CLOSE_ROUND : ")",
  TRUE : "true",
  FLASE : "false",
}

func isEscapedChar(token Token) bool {
  return token > escaped_char_start && token < escaped_char_end
}

// etc...
