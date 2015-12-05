package codeGeneration

import (
	. "ast"
	//. "backend/armarchitecture"
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
	ADDR_SIZE   = 4
)

// Print format strings
const (
	INT_FORMAT     = "%d\\0"
	STRING_FORMAT  = "%.*s\\0"
	NEWLINE_MSG    = "\\0"
	TRUE_MSG       = "true\\0"
	FALSE_MSG      = "false\\0"
	POINTER_FORMAT = "%p\\0"
)

// FUnction global variable
var functionList []*Function // IS IT A POINTER OR NOT??
var paramMap map[Param]int

// VISIT FUNCTIONS -------------------------------------------------------------

func (cg CodeGenerator) cgVisitProgram(node *Program) {
	// Set properties of global scope
	cg.globalStack.size = GetScopeVarSize(node.StatList)
	cg.globalStack.currP = cg.globalStack.size

	// ASSIGN functions to global variable
	// WE WIll only traverse them if they are called
	functionList = node.FunctionList

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

	// ldr r0, =0 to return 0 as the main return
	appendAssembly(cg.instrs, "LDR r0, =0", 1, 1)

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.instrs, "POP {pc}", 1, 1)

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

// ONLY VISIT FUNCTION IF IT IS CALLED
// IE WE ONLY PUSH ONTO STACK FUNC VARIABLES WHEN A FUNCTION IS CALLED
// but
// WE EXECUTE WHAT IS INSIDE THE FUNCTION REGARDLESS OF WHETHER IT IS CALLED OR NOT
func (cg CodeGenerator) cgVisitCallStat(ident Ident, paramList []Param) {
	for _, function := range functionList {
		if function.Ident == ident {
			// sub sp, sp, #n to create variable space
			appendAssembly(cg.instrs, "SUB sp, sp, #4", 1, 1)

			for _, param := range paramList {
				cg.cgVisitParameter(param, 0) // NED SOME KIND OF MAP HERE FROM IDENT STRING TO IDENT OFFSET INT
				// NEED SOMEHOW TO ACCUMULATE GLOBABL OFFSET
			}

			appendAssembly(cg.instrs, "BL f_"+string(function.Ident), 1, 1)

			if !(paramList == nil) {
				appendAssembly(cg.instrs, "ADD sp, sp, #"+"offset", 1, 1) // ADD LOGIC WHICH GETS GLOBAL OFFSET
			}
			appendAssembly(cg.instrs, "MOV r4, r0", 1, 1)
			appendAssembly(cg.instrs, "STR r4, [sp]", 1, 1)

			cg.cgVisitFunction(*function)

			/*
				if node.ParameterTypes != nil {
					for _, param := range node.ParameterTypes {
						cg.cgVisitParameter(param,0)
					}
				}*/

		}
	}
}

func (cg CodeGenerator) cgVisitFunction(node Function) {
	// f_funcName:
	appendAssembly(cg.funcInstrs, "f_"+string(node.Ident)+":", 0, 1)

	// push {lr} to save the caller address
	appendAssembly(cg.funcInstrs, "PUSH {lr}", 1, 1)

	// traverse all statements by switching on statement type
	// BUT NEED TO KNOW THAT WE NEED TO ADD THIS TO DUNCTION MESSGES??
	// FLAGG??
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	appendAssembly(cg.funcInstrs, "POP {pc}", 1, 1) // TEST harness uses double POP don't think we need it
	appendAssembly(cg.funcInstrs, ".ltorg", 1, 2)
}

// VISIT STATEMENT -------------------------------------------------------------
func (cg CodeGenerator) cgVisitParameter(node Param, offset int) {
	// node.Ident
	/*switch node.ParamType.(type) {
	case Boolean:
		boolean := node.ParamType.(Boolean)
		if bool(boolean) == true { // IS THIS CORRECT ????
			appendAssembly(cg.instrs, "MOV r4, #1", 1, 1)
		} else {
			appendAssembly(cg.instrs, "MOV r4, #0", 1, 1)
		}
		appendAssembly(cg.instrs, "STRB r4, [sp, #"+cg.subCurrP(BOOL_SIZE)+"]!", 1, 1)
		paramMap[boolean] = BOOL_SIZE
	case Character:
		char := node.ParamType.(Character)
		appendAssembly(cg.instrs, "MOV r4, #"+string(char), 1, 1)
		appendAssembly(cg.instrs, "STRB r4, [sp, #"+cg.subCurrP(CHAR_SIZE)+"]!", 1, 1)
		paramMap[char] = CHAR_SIZE
	case Integer:
		integer := node.ParamType.(Integer)
		appendAssembly(cg.instrs, "LDR r4, ="+strconv.Itoa(int(integer)), 1, 1)
		appendAssembly(cg.instrs, "STR r4, [sp, #"+cg.subCurrP(INT_SIZE)+"]!", 1, 1)
		paramMap[integer] = INT_SIZE
	case Str: // OR char[] need to implement
		appendAssembly(cg.instrs, "LDR r4, =msg_"+"0", 1, 1) // NEED TO ADD FUNCTIONALITY WHICH UPDATES THE MESSAGE NUMBERS
		appendAssembly(cg.instrs, "STR r4, [sp, #"+cg.subCurrP(STRING_SIZE)+"]!", 1, 1)
		paramMap["msg_0"] = STRING_SIZE // NEED TO MODIFIY THIS SHIITT

	case ArrayType:

	case PairType: // there is only a pariliteral 'null for this case'
		appendAssembly(cg.instrs, "LDR r4, =0", 1, 1)
		appendAssembly(cg.instrs, "STR r4, [sp, #-4]!", 1, 1)

	}*/
}

//Pushes a pair of elements onto the stack
func (cg CodeGenerator) pushPair(fst Evaluation, snd Evaluation, typeFst Type, typeSnd Type, reg1 string, reg2 string) {
	// Store the address in the free register
	appendAssembly(cg.instrs, "MOV "+reg2+", r0", 1, 1)

	// Load the first element into a register to be stored
	switch typeFst.(type) {
	case PairType:
		fmt.Println("Pair type for fst not implemented")
	case Integer:
		appendAssembly(cg.instrs, "LDR "+reg1+", ="+strconv.Itoa(int(fst.(Integer))), 1, 1)
	case Boolean:
		appendAssembly(cg.instrs, "LDR "+reg1+", ="+boolInt(bool(fst.(Boolean))), 1, 1)
	case Str:
		appendAssembly(cg.instrs, "LDR "+reg1+", "+cg.getMsgLabel(string(fst.(Str))), 1, 1)
	case Character:
	default:
		fmt.Println("Unknown type for pair fst")

	}

	//Allocate memory for the first element
	var fstSize, sndSize = pairTypeSize(typeFst, typeSnd)
	appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(fstSize), 1, 1)
	appendAssembly(cg.instrs, "BL malloc", 1, 1)
	//Store the first element to the newly allocated memory onto the stack
	appendAssembly(cg.instrs, "STR "+reg1+", [r0]", 1, 1)
	//Store the address of allocated memory block of the pair on the stack
	appendAssembly(cg.instrs, "STR r0, ["+reg2+"]", 1, 1)

	//Load the second element into a register to be stored
	switch typeSnd.(type) {

	case PairType:
		fmt.Println("Pair type for snd not implemented")

	case Integer:
		appendAssembly(cg.instrs, "LDR "+reg1+", ="+strconv.Itoa(int(snd.(Integer))), 1, 1)
	case Boolean:
		appendAssembly(cg.instrs, "LDR "+reg1+", ="+boolInt(bool(snd.(Boolean))), 1, 1)
	case Str:
		appendAssembly(cg.instrs, "LDR "+reg1+", "+cg.getMsgLabel(string(snd.(Str))), 1, 1)
	case Character:

	default:
		fmt.Println("Unknown type for pair snd")

	}

	//Allocate memory for the second element
	appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(sndSize), 1, 1)
	appendAssembly(cg.instrs, "BL malloc", 1, 1)
	//Store the second element to the newly allocated memory onto the stack
	appendAssembly(cg.instrs, "STR "+reg1+", [r0]", 1, 1)
	//Store the address of allocated memory block of the pair on the stack
	appendAssembly(cg.instrs, "STR r0, ["+reg2+", #4]", 1, 1)

	//Store the address of the address that contains pointers to the first and second elements
	appendAssembly(cg.instrs, "STR "+reg2+", [sp]", 1, 1)

}

// Global handle function
func (cg CodeGenerator) handle(t Evaluation, srcReg string) {

	switch t.(type) {
	case Boolean:
		cg.handleBool(t, srcReg)
	case Integer:
		cg.handleInt(t, srcReg)
	case Str:
		cg.handleString(t, srcReg)
	case Character:
		cg.handleChar(t, srcReg)
	default:
		fmt.Println("handle() doesn't support type")
	}
}

// Handles the delegation of integer calls (Better name?)
func (cg CodeGenerator) handleInt(t Type, srcReg string) {

	switch t.(type) {
	case PairElem:
		fmt.Println("PairElem not implemented")
	case Call:
		// Henryk please fix this crash
		// cg.cgVisitCallStat(t.(Call).Ident,t.(Call).ParamList)
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
	case Ident:
		// Load the address from the stack to register
		var value, _ = cg.getIdentOffset(t.(Ident))
		appendAssembly(cg.instrs, "LDR "+srcReg+", [sp, #"+strconv.Itoa(value)+"]", 1, 1)
	case ArrayElem:
		fmt.Println("ArrayElem not")
	case Type:
		// Load the value of the declaration to the register
		appendAssembly(cg.instrs, "LDR "+srcReg+", ="+strconv.Itoa(int(t.(Integer))), 1, 1)
	}

}

// Handles the delegation of boolean calls (Better name?)
func (cg CodeGenerator) handleBool(t Type, srcReg string) {
	switch t.(type) {
	case PairElem:
		fmt.Println("PairElem not implemented")
	case Call:
		// Henryk please fix this crash
		// cg.cgVisitCallStat(t.(Call).Ident,t.(Call).ParamList)
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
	case Ident:
		// Load the address from the stack to register
		var value, _ = cg.getIdentOffset(t.(Ident))
		appendAssembly(cg.instrs, "LDR "+srcReg+", [sp, #"+strconv.Itoa(value)+"]", 1, 1)
	case ArrayElem:
		fmt.Println("ArrayElem")
	case Type:
		// Load the bool into a register
		appendAssembly(cg.instrs, "MOV "+srcReg+", #"+boolInt(bool(t.(Boolean))), 1, 1)
	}
}

// Handles the delegation of character calls (Better name?)
func (cg CodeGenerator) handleChar(t Type, srcReg string) {
	switch t.(type) {
	case PairElem:
		fmt.Println("PairElem not implemented")
	case Call:
		// Henryk please fix this crash
		// cg.cgVisitCallStat(t.(Call).Ident,t.(Call).ParamList)
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
	case Ident:
		// Load the address from the stack to register
		var value, _ = cg.getIdentOffset(t.(Ident))
		appendAssembly(cg.instrs, "LDR "+srcReg+", [sp, #"+strconv.Itoa(value)+"]", 1, 1)
	case ArrayElem:
		fmt.Println("ArrayElem not implemeted")
	case Type:
		// Load the character into a register
		appendAssembly(cg.instrs, "MOV "+srcReg+", #"+string(t.(Character)), 1, 1)
	}
}

// Handles the delegation of string calls (Better name?)
func (cg CodeGenerator) handleString(t Evaluation, srcReg string) {
	switch t.(type) {
	case PairElem:
		fmt.Println("PairElem not implemented")
	case Call:
		// cg.cgVisitCallStat(t.(Call).Ident,t.(Call).ParamList)
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
	case Ident:
		// Load the address from the stack to register
		var value, _ = cg.getIdentOffset(t.(Ident))
		appendAssembly(cg.instrs, "LDR "+srcReg+", [sp, #"+strconv.Itoa(value)+"]", 1, 1)
	case ArrayElem:
		fmt.Println("ArrayElem not implemeted")
	case Type:
		// Load the address of the string to the register
		appendAssembly(cg.instrs, "LDR "+srcReg+", "+cg.getMsgLabel(string(t.(Str))), 1, 1)
	}
}

// Puts the array elements onto the stack
func (cg CodeGenerator) pushArrayElements(array []Evaluation, srcReg string, dstReg string, t Type) {

	var arraySize = arraySize(array)
	// Loop through the array pushing it onto the stack
	for i := 0; i < arraySize; i++ {
		// Array of pairs,ints,bools,chars,strings
		var arrayItem = array[i]
		switch t.(type) {
		case ArrayType:
			switch t.(ArrayType).Type {
			case Int:
				cg.handleInt(arrayItem, srcReg)
				appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+INT_SIZE*i)+"]", 1, 1)
			case String:
				cg.handleString(arrayItem, srcReg)
				//appendAssembly(cg.instrs, "LDR "+srcReg+", "+ cg.getMsgLabel(array[i].(string)), 1, 1)
				appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+STRING_SIZE*i)+"]", 1, 1)
			case Bool:
				cg.handleBool(arrayItem, srcReg)
				//appendAssembly(cg.instrs, "MOV "+srcReg+", #"+boolInt(arrayItem.(bool)), 1, 1)
				appendAssembly(cg.instrs, "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+BOOL_SIZE*i)+"]", 1, 1)
			case Char:
				cg.handleChar(arrayItem, srcReg)
				//appendAssembly(cg.instrs, "MOV "+srcReg+", #"+arrayItem.(Character).Value, 1, 1)
				appendAssembly(cg.instrs, "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+CHAR_SIZE*i)+"]", 1, 1)
			default:
				fmt.Println("Array/Pair type not done!")
			}
		}
	}
	// Put the size of the array onto the stack
	appendAssembly(cg.instrs, "LDR "+srcReg+", ="+strconv.Itoa(arraySize), 1, 1)
	// Now store the address of the array onto the stack
	appendAssembly(cg.instrs, "STR "+srcReg+", ["+dstReg+"]", 1, 1)
	appendAssembly(cg.instrs, "STR "+dstReg+", [sp, #"+cg.subCurrP(INT_SIZE)+"]", 1, 1)
}

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {
	var rhs = node.Rhs

	switch node.DecType.(type) {

	case ArrayType:

		switch rhs.(type) {

		case Call:
			//	cgVisitCallStat(node.Rhs.(Call).Ident, node.Rhs.(Call).ParamList)
		case ArrayLiter:
			//Calculate the amount of storage space required for the array
			// = ((arrayLength(array) * sizeOf(arrayType)) + 4 (4-bytes for an address)
			var arraySize = arraySize(rhs.(ArrayLiter).Exprs)
			var arrayStorage = (arraySize * sizeOf(node.DecType)) + ARRAY_SIZE

			appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(arrayStorage), 1, 1)
			//Allocate memory for the array
			appendAssembly(cg.instrs, "BL malloc", 1, 1)
			appendAssembly(cg.instrs, "MOV r4, r0", 1, 1)

			//Start loading each element in the array onto the stack
			cg.pushArrayElements(rhs.(ArrayLiter).Exprs, "r5", "r4", node.DecType)
		default:
			fmt.Println("RHS Type not implemented")
		}
	case PairType:

		// First allocate memory to store two addresses (8-bytes)
		appendAssembly(cg.instrs, "LDR r0, ="+strconv.Itoa(ADDR_SIZE*2), 1, 1)
		appendAssembly(cg.instrs, "BL malloc", 1, 1)

		// Push a pair of elements onto the stack
		cg.pushPair(rhs.(NewPair).FstExpr, rhs.(NewPair).SndExpr, node.DecType.(PairType).FstType, node.DecType.(PairType).SndType, "r5", "r4")

	//cg.handle(node.DecType.(ConstType),"r4")

	case Boolean:
		cg.handleBool(rhs, "r4")
		// Using STRB, store it on the stack
		appendAssembly(cg.instrs, "STRB r4, [sp, #"+cg.subCurrP(BOOL_SIZE)+"]", 1, 1)
	case Character:
		cg.handleChar(rhs, "r4")
		// Using STRB, store it on the stack
		appendAssembly(cg.instrs, "STRB r4, [sp, #"+cg.subCurrP(CHAR_SIZE)+"]", 1, 1)
	case Integer:
		cg.handleInt(rhs, "r4")
		// Store the value of declaration to stack
		appendAssembly(cg.instrs, "STR r4, [sp, #"+cg.subCurrP(INT_SIZE)+"]", 1, 1)
	case Str:
		cg.handleString(rhs, "r4")
		// Store the address onto the stack
		appendAssembly(cg.instrs, "STR r4, [sp, #"+cg.subCurrP(STRING_SIZE)+"]", 1, 1)

	default:
		//	typeOf(node.DecType)
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

	case ArrayElem:
	case Unop:
	case Binop:
	case Call:
	//	cgVisitCallStat(node.Rhs.(Call).Ident, node.Rhs.(Call).ParamList)
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
	// LDR r0, =n : loads return type to r0 argument
	var reg = "r4"
	cg.handleInt(node.Expr, reg)
	appendAssembly(cg.instrs, "MOV r0, "+reg, 1, 1)
	// BL exit : call exit
	appendAssembly(cg.instrs, "BL exit", 1, 1)
}

func (cg CodeGenerator) cgVisitPrintStat(node Print) {
	expr := node.Expr

	switch expr.(type) {

	case Str:
		strValue := expr.(Str)

		// LDR r4, =msg_n : load the string message label
		msgLabel := cg.getMsgLabel(string(strValue))
		appendAssembly(cg.instrs, "LDR r4, "+msgLabel, 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_string
		appendAssembly(cg.instrs, "BL p_print_string", 1, 1)
		// Define relevant print function definition (iff it hasnt been defined)
		cg.cgVisitPrintStatFunc_H("p_print_string")

	case Integer:
		intValue := expr.(Integer)

		// LDR r4, =i : load the value into r4
		appendAssembly(cg.instrs, "LDR r4, ="+strconv.Itoa(int(intValue)), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_int
		appendAssembly(cg.instrs, "BL p_print_int", 1, 1)
		// Define relevant print function definition (iff it hasnt been defined)
		cg.cgVisitPrintStatFunc_H("p_print_int")

	case Character:
		charValue := expr.(Character)

		// MOV r4, #'c' : load the value into r4
		appendAssembly(cg.instrs, "MOV r4, #"+string(charValue), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL putchar
		appendAssembly(cg.instrs, "BL putchar", 1, 1)

	case Boolean:
		boolValue := expr.(Boolean)

		// MOV r4, #e.g.1 : load the value into r4
		appendAssembly(cg.instrs, "MOV r4, #"+boolInt(bool(boolValue)), 1, 1)
		// MOV r0, r4 : prepare parameter for function call
		appendAssembly(cg.instrs, "MOV r0, r4", 1, 1)
		// BL p_print_bool
		appendAssembly(cg.instrs, "BL p_print_bool", 1, 1)
		// Define relevant print function definition (iff it hasnt been defined)
		cg.cgVisitPrintStatFunc_H("p_print_bool")
	default:
		typeOf(expr)

	}
}

// cgVisitPrintStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg CodeGenerator) cgVisitPrintStatFunc_H(funcName string) {
	if cg.AddCheckProgName(funcName) {
		// if the program function has been defined previously
		// a redefinition is unnecessary
		return
	}
	// else define the print function

	// funcLabel:
	appendAssembly(cg.progFuncInstrs, funcName+":", 0, 1)
	// push {lr} to save the caller address
	appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)

	switch funcName {
	case "p_print_int":
		// r1 = int value
		appendAssembly(cg.progFuncInstrs, "MOV r1, r0", 1, 1)
		// r0 = int format string
		appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel(INT_FORMAT), 1, 1)

	case "p_print_bool":
		// Check bool value - 0
		appendAssembly(cg.progFuncInstrs, "CMP r0, #0", 1, 1)
		// If bool = true then r0 = "true"
		appendAssembly(cg.progFuncInstrs, "LDRNE r0, "+cg.getMsgLabel(TRUE_MSG), 1, 1)
		// If bool = false then r0 = "false"
		appendAssembly(cg.progFuncInstrs, "LDREQ r0, "+cg.getMsgLabel(FALSE_MSG), 1, 1)

	case "p_print_string":
		// r1 = string value
		appendAssembly(cg.progFuncInstrs, "LDR r1, [r0]", 1, 1)
		// r2 = r0 + 4
		appendAssembly(cg.progFuncInstrs, "ADD r2, r0, #4", 1, 1)
		// r0 = string format string
		appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel(STRING_FORMAT), 1, 1)

	case "p_print_ln":
		// r0 = new line string
		appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel(NEWLINE_MSG), 1, 1)

	case "p_print_reference":
		// r1 = int value
		appendAssembly(cg.progFuncInstrs, "MOV r1, r0", 1, 1)
		// r0 = pointer format string
		appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel(POINTER_FORMAT), 1, 1)
	}

	//
	appendAssembly(cg.progFuncInstrs, "ADD r0, r0, #4", 1, 1)
	// calls printf or puts
	if funcName == "p_print_ln" {
		appendAssembly(cg.progFuncInstrs, "BL puts", 1, 1)
	} else {
		appendAssembly(cg.progFuncInstrs, "BL printf", 1, 1)
	}
	// Sets fflush argument
	appendAssembly(cg.progFuncInstrs, "MOV r0, #0", 1, 1)
	// calls fflush
	appendAssembly(cg.progFuncInstrs, "BL fflush", 1, 1)

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)

}

func (cg CodeGenerator) cgVisitPrintlnStat(node Println) {
	cg.cgVisitPrintStat(Print{Expr: node.Expr})
	// BL p_print_ln
	appendAssembly(cg.instrs, "BL p_print_ln", 1, 1)
	// Define relevant print function definition (iff it hasnt been defined)
	cg.cgVisitPrintStatFunc_H("p_print_ln")
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
