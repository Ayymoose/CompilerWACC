package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"fmt"
	"strconv"
)

var zeroInt Integer = 0
var zeroCharater Character = "0"
var zeroString Str = "0"
var zeroBool Boolean = false

// Contains the size in bytes of all print format strings
var mapPrintFormatToSize = map[string]int{
	INT_FORMAT:     3,
	STRING_FORMAT:  5,
	NEWLINE_MSG:    1,
	TRUE_MSG:       5,
	FALSE_MSG:      6,
	POINTER_FORMAT: 3,
}

//Size in bytes for all the variables in the current scope
func GetScopeVarSize(statList []Statement) int {
	var scopeSize = 0
	for _, stat := range statList {
		switch stat.(type) {
		case Declare:
			scopeSize += sizeOf(stat.(Declare).DecType)
		}
	}
	return scopeSize
}

// Calcuates the size of a type
func sizeOf(t Type) int {
	var size = 0
	switch t.(type) {
  case ConstType:
		switch t.(ConstType) {
		case Int:
			size = INT_SIZE
		case Bool:
			size = BOOL_SIZE
		case Char:
			size = CHAR_SIZE
		case String:
			size = STRING_SIZE
		}
	case PairType:
		size = PAIR_SIZE //Recurse on pair Type too?
	case Integer:
		size = INT_SIZE
	case Str:
		size = STRING_SIZE
	case Boolean:
		size = BOOL_SIZE
	case Character:
		size = CHAR_SIZE
	case ArrayType:
    size = sizeOf(t.(ArrayType).Type)
	default:
		fmt.Println("sizeOf(t) t is an unknown type")
		typeOf(t)
	}

	return size
}

// Small function to print the type (remove later)
func typeOf(t Type) {
	switch t.(type) {
	case ConstType:
		fmt.Println("ConstType")
	case FSND:
		fmt.Println("FSND")
	case nil:
		fmt.Println("nil")
	case Function:
		fmt.Println("Function")
	case Param:
		fmt.Println("Param")
	case ArrayType:
		fmt.Println("ArrayType")
	case PairType:
		fmt.Println("PairType")
	case Program:
		fmt.Println("Program")
	case Binop:
		fmt.Println("Binop")
	case Unop:
		fmt.Println("Unop")
	case NewPair:
		fmt.Println("Newpair")
	case Declare:
		fmt.Println("Declare")
	case Assignment:
		fmt.Println("Assignment")
	case If:
		fmt.Println("If")
	case While:
		fmt.Println("While")
	case Scope:
		fmt.Println("Scope")
	case Read:
		fmt.Println("Read")
	case Free:
		fmt.Println("Free")
	case Return:
		fmt.Println("Return")
	case Exit:
		fmt.Println("Exit")
	case Print:
		fmt.Println("Print")
	case Println:
		fmt.Println("Println")
	case Call:
		fmt.Println("Call")
	case Ident:
		fmt.Println("Ident")
	case PairElem:
		fmt.Println("PairElem")
	case ArrayLiter:
		fmt.Println("ArrayLiter")
	case ArrayElem:
		fmt.Println("ArrayElem")
	case Evaluation:
		fmt.Println("Evaluation")
	case Statement:
		fmt.Println("Statement")
	case Integer:
		fmt.Println("Integer")
	case Str:
		fmt.Println("Str")
	case Boolean:
		fmt.Println("Bool")
	case Character:
		fmt.Println("Character")
	case PairLiter:
		fmt.Println("PairLiter")
	case Skip:
		fmt.Println("Skip")
	case int:
		fmt.Println("int")
	case bool:
	  fmt.Println("bool")
	case string:
	  fmt.Println("string")
	case rune:
	  fmt.Println("rune")
	default:
		fmt.Println("Unknown")
	}
}

// Adds the string code to the list of instructions instrs.
// numTabs  \t will be added before the string and
// numNewLines \n will be added after the string
func appendAssembly(instrs *ARMList, code string, numTabs int, numNewLines int) {
	const default_num_tabs = 1
	var str string = ""

	for i := 0; i < numTabs+default_num_tabs; i++ {
		str += "\t"
	}

	str += code

	for i := 0; i < numNewLines; i++ {
		str += "\n"
	}

	*instrs = append(*instrs, str)
}

// Adds a msg label definition for the strValue using the msg label to the
// list of assembly instructions msgInstrs
func addMsgLabel(msgInstrs *ARMList, label string, strValue string) {
	if len(*msgInstrs) == 0 {
		appendAssembly(msgInstrs, ".data", 0, 2)
	}

	appendAssembly(msgInstrs, label+":", 0, 1)

	// size of strValue in bytes
	wordSize := calculateWordSize(strValue)

	appendAssembly(msgInstrs, ".word "+strconv.Itoa(wordSize), 1, 1)
	appendAssembly(msgInstrs, ".ascii "+strValue, 1, 2)
}

// Calculates the size of strValue in bytes
func calculateWordSize(strValue string) int {
	size, contained := mapPrintFormatToSize[strValue]

	// if strValue is a format string then
	if contained {
		return size
	} else {
		quotation := 2
		var escapedChars int
		backSlashSeen := false

		for _, c := range strValue {
			if c == '\\' && !backSlashSeen {
				escapedChars++
				backSlashSeen = true
			} else {
				backSlashSeen = false
			}
		}

		return len(strValue) - escapedChars - quotation
	}
}

// Returns "1" iff b = true. "0" otherwise
func boolInt(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}
