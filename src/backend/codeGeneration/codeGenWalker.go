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
	if cg.globalStack.size > 0 {
		appendAssembly(cg.instrs, "SUB sp, sp, #"+strconv.Itoa(cg.globalStack.size), 1, 1)
	}

	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	// add sp, sp, #n to remove variable space
	if cg.globalStack.size > 0 {
		appendAssembly(cg.instrs, "ADD sp, sp, #"+strconv.Itoa(cg.globalStack.size), 1, 1)
	}

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.instrs, "pop {pc}", 1, 1)

	// .ltorg
	appendAssembly(cg.instrs, ".ltorg", 1, 1)
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
		}
	case ArrayType:

	case PairType:

	}
}

//Puts the array elements onto the stack
func (cg CodeGenerator) pushArrayElements(array []interface{}, srcReg string, dstReg string, t Type) {

	var arraySize = arraySize(array)
	//Loop through the array pushing it onto the stack
	for i := 0; i < arraySize; i++ {
		//Array of pairs,ints,bools,chars,strings
		var arrayItem = array[i]
		switch t.(type) {
		case ArrayType:
			switch t.(ArrayType).Type {
			case Int:
				appendAssembly(cg.instrs, "LDR "+srcReg+", ="+strconv.Itoa(arrayItem.(int)), 1, 1)
				appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+INT_SIZE*i)+"]", 1, 1)
			case String:
				//NEED TO GET CORRECT LABEL
				var label = "msg_"
				appendAssembly(cg.instrs, "LDR "+srcReg+", ="+label, 1, 1)
				appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+STRING_SIZE*i)+"]", 1, 1)
				//fmt.Println(array[i])
			case Bool:
				appendAssembly(cg.instrs, "MOV "+srcReg+", #"+boolToString(arrayItem.(bool)), 1, 1)
				appendAssembly(cg.instrs, "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+BOOL_SIZE*i)+"]", 1, 1)
			case Char:
				//WHY DOES THIS PRINT 0 ????
				fmt.Println(array[i])
				//appendAssembly(cg.instrs, "MOV "+srcReg+", #"+strconv.Itoa(arrayItem.(string)), 1, 1)
				appendAssembly(cg.instrs, "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+CHAR_SIZE*i)+"]", 1, 1)
			default:
				fmt.Println("Array/Pair type not done!")
			}
		}
	}
	//Put the size of the array onto the stack
	appendAssembly(cg.instrs, "LDR "+srcReg+", ="+strconv.Itoa(arraySize), 1, 1)
	//Now store the address of the array onto the stack
	appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+"]", 1, 1)
	appendAssembly(cg.instrs, "STR "+dstReg+", [sp, #"+cg.subCurrP(INT_SIZE)+"]", 1, 1)
}

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {
	//MIGHT NEED TO CHANGE THIS BACK TO SIMPLE IF STATEMENT
	var array = node.Rhs
	switch array.(type) {
	case ArrayLiter:
		//Calculate the amount of storage space required for the array
		// = ((arrayLength(array) * sizeOf(arrayType)) + 4 (4-bytes for an address)
		var arraySize = arraySize(array.(ArrayLiter).Exprs)
		var arrayStorage = (arraySize * sizeOf(node.DecType)) + ARRAY_SIZE

		appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(arrayStorage), 1, 1)
		//Allocate memory for the array
		appendAssembly(cg.instrs, "BL malloc", 1, 1)
		//Move the result back into the register
		appendAssembly(cg.instrs, "MOV r4, r0", 1, 1)
		//Start loading each element in the array onto the stack
		cg.pushArrayElements(array.(ArrayLiter).Exprs, "r5", "r4", node.DecType)
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
			appendAssembly(cg.instrs, "STR r4, [sp, #"+cg.subCurrP(INT_SIZE)+"])", 1, 1)
		case String:
			fmt.Println("String not implemented")

		}
	}

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
	default:

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
	//	label := "p_free_pair"
}

func (cg CodeGenerator) cgVisitReturnStat(node Return) {

}

func (cg CodeGenerator) cgVisitExitStat(node Exit) {

}

func (cg CodeGenerator) cgVisitPrintStat(node Print) {
	expr := node.Expr

	switch expr.(type) {
	case string:
		strValue := expr.(string)

		// LDR r4, =msg_n : load the string message label
		msgLabel := cg.getMsgLabel(strValue)
		appendAssembly(cg.instrs, "LDR r4, "+msgLabel, 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_string
		appendAssembly(cg.instrs, "BL p_print_string", 1, 1)

	case int:
		intValue := expr.(int)

		// LDR r4, =i : load the value into r4
		appendAssembly(cg.instrs, "LDR r4, ="+strconv.Itoa(intValue), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_int
		appendAssembly(cg.instrs, "BL p_print_int", 1, 1)

	case rune:
		charValue := expr.(rune)

		// MOV r4, #'c' : load the value into r4
		appendAssembly(cg.instrs, "MOV r4, #"+string(charValue), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL putchar
		appendAssembly(cg.instrs, "BL putchar", 1, 1)

	case bool:
		boolValue := expr.(bool)

		// MOV r4, #e.g.1 : load the value into r4
		appendAssembly(cg.instrs, "MOV r4, #"+boolInt(boolValue), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_bool
		appendAssembly(cg.instrs, "BL p_print_bool", 1, 1)

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
