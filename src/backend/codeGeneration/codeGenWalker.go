package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"fmt"
	"strconv"
)

// CONSTANTS -------------------------------------------------------------------

// Type sizes in bytes
const (
	INT_SIZE    = 4
	ARRAY_SIZE  = 4
	BOOL_SIZE   = 1
	CHAR_SIZE   = 1
	STRING_SIZE = 4
	PAIR_SIZE   = 4
)

// Print format strings
const (
	INT_FORMAT    = "%d\\0"
	STRING_FORMAT = "%.*s\\0"
	NEW_LINE      = "\\0"
	TRUE_MSG      = "true\\0"
	FALSE_MSG     = "false\\0"
)

// VISIT FUNCTIONS -------------------------------------------------------------

func (cg CodeGenerator) cgVisitProgram(node *Program) {
	// Set properities of global scope
	cg.globalStack.size = GetScopeVarSize(node.StatList)
	cg.globalStack.currP = cg.globalStack.size

	// traverse all functions
	for _, function := range node.FunctionList {
		cg.cgVisitFunction(*function)
	}

	// .text
	appendAssembly(cg.instrs, ".text", 0, 2)

	// .global main
	appendAssembly(cg.instrs, ".global main", 0, 1)

	//main:
	appendAssembly(cg.instrs, "main:", 0, 1)

	// push {lr} to save the caller address
	appendAssembly(cg.instrs, "PUSH {lr}", 1, 1)

	// sub sp, sp, #n to create variable space
	appendAssembly(cg.instrs, "SUB sp, sp, #"+strconv.Itoa(cg.globalStack.size), 1, 1)

	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	// add sp, sp, #n to remove variable space
	appendAssembly(cg.instrs, "ADD sp, sp, #"+strconv.Itoa(cg.globalStack.size), 1, 1)

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.instrs, "pop {pc}", 1, 1)

	// .ltorg
	appendAssembly(cg.instrs, ".ltorg", 1, 1)

	// Adds the msg definitions to assembly instructions
	*cg.instrs = append(cg.msgInstrs, *cg.instrs...)
}

func (cg CodeGenerator) cgCreateMsgs(instrs *ARMList) map[string]string {
	return nil
}

func (cg CodeGenerator) cgEvalStat(stat interface{}) {
	switch stat.(type) {
	case Declare:
		cg.cgVisitDeclareStat(stat.(Declare))
	case Assignment:
		cg.cgVisitAssignmentStat(stat.(Assignment))
	case Read:
		cg.cgVisitReadStat(stat.(Read))
	case Free:
		cg.cgVisitFreeStat(stat.(Free))
	case Return:
		cg.cgVisitReturnStat(stat.(Return))
	case Exit:
		cg.cgVisitExitStat(stat.(Exit))
	case Print:
		cg.cgVisitPrintStat(stat.(Print))
	case Println:
		cg.cgVisitPrintlnStat(stat.(Println))
	case If:
		cg.cgVisitIfStat(stat.(If))
	case While:
		cg.cgVisitWhileStat(stat.(While))
	case Scope:
		cg.cgVisitScopeStat(stat.(Scope))
	default:
		//	""
	}
}

func (cg CodeGenerator) cgVisitFunction(node Function) {
	// f_funcName:
	appendAssembly(cg.instrs, "f_"+node.Ident+":", 1, 1)

	// push {lr} to save the caller address
	appendAssembly(cg.instrs, "PUSH {lr}", 1, 1)

	if node.ParameterTypes != nil {
		for _, param := range node.ParameterTypes {
			cg.cgVisitParameter(param)
		}
	}

	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}
}

// VISIT STATEMENT -------------------------------------------------------------
func (cg CodeGenerator) cgVisitParameter(node Param) {
	// node.Ident
	switch node.ParamType.(type) {
	case ConstType:
		switch node.ParamType.(ConstType) {
		case Bool:
		case Char:
		case Int:
		case String:
			//	case Pair:
		}
	case ArrayType:

	case PairType:

	}
}

//Loads the array elements onto the stack
func pushArrayElements(array []interface{}, register string) {

	var arraySize = arraySize(array)
	//Loop through the array pushing it onto the stack
	for i := 0; i < arraySize; i++ {
		//Array of pairs,ints,bools,chars,strings
		var arrayItem = array[i]

		switch arrayItem.(type) {
		case ConstType:

			fmt.Println("int")

			/*	case Int:
					fmt.Println("int")
				case Bool:
					fmt.Println("bool")
				case Char:
					fmt.Println("char")
				case String:
					fmt.Println("string")*/
		default:
			fmt.Println("String/Array/Pair type not implemented!")
		}

		//var arrayItem = array.[i].(int)
		//appendAssembly(cg.instrs, "LDR r5, ="+strconv.Itoa(arrayItem), 1, 1)
	}
}

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {

	//MIGHT NEED TO CHANGE THIS BACK TO SIMPLE IF STATEMENT
	var array = node.Rhs
	switch array.(type) {
	case ArrayLiter:
		//Calculate the amount of storage space required for the array
		// = (arrayLength(array) + 1) * sizeOf(arrayType)
		var arraySize = arraySize(array.(ArrayLiter).Exprs)
		var arrayStorage = (arraySize + 1) * sizeOf(node.DecType)

		appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(arrayStorage), 1, 1)
		//Allocate memory for the array
		appendAssembly(cg.instrs, "BL malloc", 1, 1)
		//Move the result back into the register
		appendAssembly(cg.instrs, "MOV r4, r0", 1, 1)

		//Start loading each element in the array onto the stack
		pushArrayElements(array.(ArrayLiter).Exprs, "r5")

	}

	switch node.DecType.(type) {
	case ConstType:
		switch node.DecType.(ConstType) {
		case Bool:

			//Load the bool into a register
			appendAssembly(cg.instrs, "MOV r4, #"+strconv.Itoa(node.Rhs.(int)), 1, 1)
			//Using STRB, store it on the stack
			appendAssembly(cg.instrs,
				"STRB r4, [sp, #"+cg.subCurrP(BOOL_SIZE)+"])", 1, 1)

		case Char:
			//Load the character into a register
			appendAssembly(cg.instrs, "MOV r4, #"+node.Rhs.(string), 1, 1)
			//Using STRB, store it on the stack
			appendAssembly(cg.instrs,
				"STRB r4, [sp, #"+cg.subCurrP(CHAR_SIZE)+"])", 1, 1)
		case Int:
			// Load the value of the declaration to the register
			appendAssembly(cg.instrs, "LDR r4, "+strconv.Itoa(node.Rhs.(int)), 1, 1)
			//fmt.Println("Type", node.Rhs)

			// Load the value of the declaration to R4
			//appendAssembly(cg.instrs, "LDR R4, "+strconv.Itoa(), 1, 1)

			// Store the value of declaration to stack
			appendAssembly(cg.instrs,
				"STR r4, [sp, #"+cg.subCurrP(INT_SIZE)+"])", 1, 1)
		case String:

		}

	}

	// Store the value of the declaration to the stack

}

//Calcuates the type of the array and returns the size in bytes that the array occupies
func sizeOf(t Type) int {
	var size = 0
	switch t.(type) {
	case ArrayType:
		switch t.(ArrayType).Type {
		case Int:
			size = INT_SIZE
		case String:
			size = STRING_SIZE
		case Bool:
			size = BOOL_SIZE
		case Char:
			size = CHAR_SIZE
		default:
			fmt.Println("No type found!")
		}
	default:
		fmt.Println("No type found!")
	}
	return size
}

//Calcuates the size of an array
func arraySize(array []interface{}) int {
	return len(array)
}

func (cg CodeGenerator) cgVisitAssignmentStat(node Assignment) {
	switch node.Lhs.(type) {
	case ArrayElem:
	case PairElem:
	default: // Ident
	}

	// eval RHS ????
	switch node.Rhs.(type) {
	// expr
	case ConstType:
		switch node.Rhs.(ConstType) {
		case Bool:

		case Char:

		case Int:

		case String:

			//	case Pair:
		}
	case ArrayElem:
	case Unop:
	case Binop:

	default: // Ident   // NEED TO DEAL WITH '(' EXPR ')' CASE AS WELL?
	}
}

func (cg CodeGenerator) cgVisitReadStat(node Read) {
	label := "p_read_"
	offset := 0
	switch node.AssignLHS.(type) {
	case Ident:
	case ArrayElem:
	case PairElem:
	}

	if offset == 4 {
		label += "int"
	} else if offset == 1 {
		label += "char"
	}
	// p_read_int:  /char
	//PUSH {lr}

	// BL scanf
	//POP {pc}
}

func (cg CodeGenerator) cgVisitFreeStat(node Free) {
	//label := "p_free_pair"
}

func (cg CodeGenerator) cgVisitReturnStat(node Return) {

}

func (cg CodeGenerator) cgVisitExitStat(node Exit) {

}

func (cg CodeGenerator) cgVisitPrintStat(node Print) {
	expr := node.Expr

	switch expr.(type) {

	}
}

func (cg CodeGenerator) cgVisitPrintlnStat(node Println) {

}

func (cg CodeGenerator) cgVisitIfStat(node If) {

}

func (cg CodeGenerator) cgVisitWhileStat(node While) {

}

func (cg CodeGenerator) cgVisitScopeStat(node Scope) {

}

// VISIT EXPRESSIONS -----------------------------------------------------------

func (cg CodeGenerator) cgVisitUnopExpr(node Unop) {

}

func (cg CodeGenerator) cgVisitBinopExpr(node Binop) {

}
