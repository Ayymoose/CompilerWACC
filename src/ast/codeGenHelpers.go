package ast

import (
	. "backend/filewriter"
	"fmt"
)

//Size in bytes for all the variables in the current scope
func GetScopeVarSize(statList []interface{}) int {
	var scopeSize = 0

	for _, stat := range statList {
		switch stat.(type) {

		case Declare:
			var t = stat.(Declare).Type

			switch t.(type) {
			case PairType:
				//Address of pair on the stack is 4 bytes
				scopeSize += PAIR_SIZE
			case ArrayType:
				var e = stat.(Declare).Rhs.(ArrayLiter)
				var sizeOf = 0

				switch t.(ArrayType).Type {
				case Int:
					sizeOf = INT_SIZE
				case String:
					sizeOf = STRING_SIZE
				case Bool:
					sizeOf = BOOL_SIZE
				case Char:
					sizeOf = CHAR_SIZE
				default:
					fmt.Println("No type found for ArrayType")
				}
				//The size would be the equal to
				//(Number of elements + 1) * sizeOf(element)
				scopeSize += (len(e.Exprs) + 1) * sizeOf

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
// numTabs  /t will be added before the string and
// numNewLines /n will be added after the string
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
