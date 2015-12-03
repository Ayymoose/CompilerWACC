package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"fmt"
	"strconv"
)

// Contains the size in bytes of all print format strings
var mapPrintFormatToSize = map[string]int{
	INT_FORMAT:    3,
	STRING_FORMAT: 5,
	NEW_LINE:      1,
	TRUE_MSG:      5,
	FALSE_MSG:     6,
}

//Size in bytes for all the variables in the current scope
func GetScopeVarSize(statList []interface{}) int {
	var scopeSize = 0

	for _, stat := range statList {
		switch stat.(type) {

		case Declare:
			var t = stat.(Declare).DecType

			switch t.(type) {
			case PairType:
				//Address of pair on the stack is 4 bytes
				scopeSize += PAIR_SIZE
			case ArrayType:
				//	var e = stat.(Declare).Rhs.(ArrayLiter)

				switch t.(ArrayType).Type {
				case Int:
					scopeSize += INT_SIZE
				case String:
					scopeSize += STRING_SIZE
				case Bool:
					scopeSize += BOOL_SIZE
				case Char:
					scopeSize += CHAR_SIZE
				default:
					fmt.Println("No type found for ArrayType")
				}

			case ConstType:
				switch t.(ConstType) {
				case Int:
					scopeSize += INT_SIZE
				case String:
					scopeSize += STRING_SIZE
				case Bool:
					scopeSize += BOOL_SIZE
				case Char:
					scopeSize += CHAR_SIZE
				default:
					fmt.Println("No type found for ConstType")
				}
			}
			//Anything else is just ignored
		}
	}

	return scopeSize
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
	appendAssembly(msgInstrs, ".ascii \""+strValue+"\"", 1, 2)
}

// Calculates the size of strValue in bytes
func calculateWordSize(strValue string) int {
	size, contained := mapPrintFormatToSize[strValue]

	// if strValue is a format string then
	if contained {
		return size
	} else {
		return len(strValue)
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
