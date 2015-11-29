package armArchitecture

import (
	"fmt"
	"reflect"
)

// Operation of ARM instruction set
type Operation struct {
	operand  Operand
	listVars []interface{} // can be up to at least 3 vars
	// ISUE WILL OCCURE WHEN WE HAVE A MIXTURE OF REGISTERS AND INTS IN THE SLICE
	// CAN't have a SLICE WITH DIFFERENT TYPES
}

// Returns string representation of Operation
func (o Operation) String() string {
	result := "\t" + fmt.Sprint(o.operand) + " "

	list := reflect.ValueOf(o.listVars)
	for i := 0; i < list.Len(); i++ {
		result += fmt.Sprint(list.Index(i)) // HOW TO CONVERT TO STRING WHEN INT // REGISTER
		if i+1 != list.Len() {
			result += ", "
		}
	}
	return result
}
