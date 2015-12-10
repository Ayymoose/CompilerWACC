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
	INT_SIZE     = 4
	ARRAY_SIZE   = 4
	BOOL_SIZE    = 1
	CHAR_SIZE    = 1
	STRING_SIZE  = 4
	PAIR_SIZE    = 4
	ADDRESS_SIZE = 4
	// Maximum offset of stack pointer
	STACK_SIZE_MAX = 1024
)

// Print format strings
const (
	INT_FORMAT     = "\"%d\\0\""
	STRING_FORMAT  = "\"%.*s\\0\""
	NEWLINE_MSG    = "\"\\0\""
	TRUE_MSG       = "\"true\\0\""
	FALSE_MSG      = "\"false\\0\""
	READ_INT       = "\"%d\\0\""
	READ_CHAR      = "\"%c\\0\""
	POINTER_FORMAT = "\"%p\\0\""
)

// error messages
const (
	NULL_REFERENCE       = "\"NullReferenceError: dereference a null reference\\n\\0\""
	ARRAY_INDEX_NEGATIVE = "\"ArrayIndexOutOfBoundsError: negative index\\n\\0\""
	ARRAY_INDEX_LARGE    = "\"ArrayIndexOutOfBoundsError: index too large\\n\\0\""
	OVERFLOW             = "\"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\\n\""
	DIVIDE_BY_ZERO       = "\"DivideByZeroError: divide or modulo by zero\\n\\0\""
)

// Function global variable
var functionList []*Function

var DEBUG = false

// HELPER FUNCTIONS
// cgVisitReadStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitReadStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// Define the read function
		appendAssembly(cg.progFuncInstrs, funcName+":", 0, 1)
		appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)
		appendAssembly(cg.progFuncInstrs, "MOV r1, r0", 1, 1)
		switch funcName {
		case "p_read_char":
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", READ_CHAR), 1, 1)
		case "p_read_int":
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", READ_INT), 1, 1)
		}
		appendAssembly(cg.progFuncInstrs, "ADD r0, r0, #4", 1, 1)
		appendAssembly(cg.progFuncInstrs, "BL scanf", 1, 1)
		appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)
	}
}

// Null pointer dereference check
func (cg *CodeGenerator) dereferenceNullPointer() {
	if !cg.AddCheckProgName("p_check_null_pointer") {
		appendAssembly(cg.progFuncInstrs, "p_check_null_pointer"+":", 0, 1)
		appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)
		appendAssembly(cg.progFuncInstrs, "CMP r0, #0", 1, 1)
		appendAssembly(cg.progFuncInstrs, "LDREQ r0, "+cg.getMsgLabel("", NULL_REFERENCE), 1, 1)
		appendAssembly(cg.progFuncInstrs, "BLEQ p_throw_runtime_error", 1, 1)
		appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)
	}
	cg.throwRunTimeError()
}

// Run time error check
func (cg *CodeGenerator) throwRunTimeError() {
	if !cg.AddCheckProgName("p_throw_runtime_error") {
		appendAssembly(cg.progFuncInstrs, "p_throw_runtime_error"+":", 0, 1)
		appendAssembly(cg.progFuncInstrs, "BL p_print_string", 1, 1)
		appendAssembly(cg.progFuncInstrs, "MOV r0, #-1", 1, 1)
		appendAssembly(cg.progFuncInstrs, "BL exit", 1, 1)
		cg.cgVisitPrintStatFuncHelper("p_print_string")
	}
}

// Check array bounds
func (cg *CodeGenerator) checkArrayBounds() {
	if !cg.AddCheckProgName("p_check_array_bounds") {
		appendAssembly(cg.progFuncInstrs, "p_check_array_bounds"+":", 0, 1)
		appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)
		appendAssembly(cg.progFuncInstrs, "CMP r0, #0", 1, 1)
		appendAssembly(cg.progFuncInstrs, "LDRLT r0, "+cg.getMsgLabel("", ARRAY_INDEX_NEGATIVE), 1, 1)
		appendAssembly(cg.progFuncInstrs, "BLLT p_throw_runtime_error", 1, 1)
		appendAssembly(cg.progFuncInstrs, "LDR r1, [r1]", 1, 1)
		appendAssembly(cg.progFuncInstrs, "CMP r0, r1", 1, 1)
		appendAssembly(cg.progFuncInstrs, "LDRCS r0, "+cg.getMsgLabel("", ARRAY_INDEX_LARGE), 1, 1)
		appendAssembly(cg.progFuncInstrs, "BLCS p_throw_runtime_error", 1, 1)
		appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)
	}
	cg.throwRunTimeError()
}

// Checks the bounds of an array
func (cg *CodeGenerator) arrayCheckBounds(array []Evaluation, reg1 string, reg2 string) {

	// Load the first index
	cg.evalRHS(array[0], reg2)

	// Set a register to point to the array
	appendAssembly(cg.currInstrs(), "LDR "+reg1+", ["+reg1+"]", 1, 1)

	// reg1 = Address of the array
	// reg2 = Index

	appendAssembly(cg.currInstrs(), "MOV r0, "+reg2, 1, 1)
	appendAssembly(cg.currInstrs(), "MOV r1, "+reg1, 1, 1)
	appendAssembly(cg.currInstrs(), "BL p_check_array_bounds", 1, 1)

	// Get offset of
	appendAssembly(cg.currInstrs(), "ADD "+reg1+", "+reg1+", #4", 1, 1)
	appendAssembly(cg.currInstrs(), "ADD "+reg1+", "+reg1+", "+reg2+", LSL #2", 1, 1)
	appendAssembly(cg.currInstrs(), "LDR "+reg1+", ["+reg1+"]", 1, 1)

	//Load the second index if there is one
	if len(array) > 1 {
		cg.evalRHS(array[1], reg2)

		appendAssembly(cg.currInstrs(), "MOV r0, "+reg2, 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r1, "+reg1, 1, 1)
		appendAssembly(cg.currInstrs(), "BL p_check_array_bounds", 1, 1)

		// Get offset of
		appendAssembly(cg.currInstrs(), "ADD "+reg1+", "+reg1+", #4", 1, 1)
		appendAssembly(cg.currInstrs(), "ADD "+reg1+", "+reg1+", "+reg2+", LSL #2", 1, 1)
		appendAssembly(cg.currInstrs(), "LDR "+reg1+", ["+reg1+"]", 1, 1)

	}

	appendAssembly(cg.currInstrs(), "MOV r0, "+reg1, 1, 1)

}

// cgVisitFreeStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitFreeStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// Define the read function
		appendAssembly(cg.progFuncInstrs, funcName+":", 0, 1)
		switch funcName {
		case "p_free_pair":
			appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)
			appendAssembly(cg.progFuncInstrs, "CMP r0, #0", 1, 1)
			appendAssembly(cg.progFuncInstrs, "LDREQ r0, "+cg.getMsgLabel("", NULL_REFERENCE), 1, 1)
			appendAssembly(cg.progFuncInstrs, "BEQ p_throw_runtime_error", 1, 1)
			appendAssembly(cg.progFuncInstrs, "PUSH {r0}", 1, 1)
			appendAssembly(cg.progFuncInstrs, "LDR r0, [r0]", 1, 1)
			appendAssembly(cg.progFuncInstrs, "BL free", 1, 1)
			appendAssembly(cg.progFuncInstrs, "LDR r0, [sp]", 1, 1)
			appendAssembly(cg.progFuncInstrs, "LDR r0, [r0, #4]", 1, 1)
			appendAssembly(cg.progFuncInstrs, "BL free", 1, 1)
			appendAssembly(cg.progFuncInstrs, "POP {r0}", 1, 1)
			appendAssembly(cg.progFuncInstrs, "BL free", 1, 1)
			appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)
		}
		cg.throwRunTimeError()
		cg.cgVisitPrintStatFuncHelper("p_print_string")
	}
}

// cgVisitPrintStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitPrintStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// funcLabel:
		appendAssembly(cg.progFuncInstrs, funcName+":", 0, 1)
		// push {lr} to save the caller address
		appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)

		switch funcName {
		case "p_print_int":
			// r1 = int value
			appendAssembly(cg.progFuncInstrs, "MOV r1, r0", 1, 1)
			// r0 = int format string
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", INT_FORMAT), 1, 1)

		case "p_print_bool":
			// Check bool value - 0
			appendAssembly(cg.progFuncInstrs, "CMP r0, #0", 1, 1)
			// If bool = true then r0 = "true"
			appendAssembly(cg.progFuncInstrs, "LDRNE r0, "+cg.getMsgLabel("", TRUE_MSG), 1, 1)
			// If bool = false then r0 = "false"
			appendAssembly(cg.progFuncInstrs, "LDREQ r0, "+cg.getMsgLabel("", FALSE_MSG), 1, 1)

		case "p_print_string":
			// r1 = string value
			appendAssembly(cg.progFuncInstrs, "LDR r1, [r0]", 1, 1)
			// r2 = r0 + 4
			appendAssembly(cg.progFuncInstrs, "ADD r2, r0, #4", 1, 1)
			// r0 = string format string
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", STRING_FORMAT), 1, 1)

		case "p_print_ln":
			// r0 = new line string
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", NEWLINE_MSG), 1, 1)

		case "p_print_reference":
			// r1 = int value
			appendAssembly(cg.progFuncInstrs, "MOV r1, r0", 1, 1)
			// r0 = pointer format string
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", POINTER_FORMAT), 1, 1)
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
}

// Because the maximum amount we can add or subtract the stack pointer by is 1024
// These helper functions allocate and deallocate space on the stack for us

func (cg *CodeGenerator) createStackSpace(stackSize int) {
	if stackSize > STACK_SIZE_MAX {
		appendAssembly(cg.currInstrs(), "SUB sp, sp, #"+strconv.Itoa(STACK_SIZE_MAX), 1, 1)
		cg.createStackSpace(stackSize - STACK_SIZE_MAX)
	} else {
		appendAssembly(cg.currInstrs(), "SUB sp, sp, #"+strconv.Itoa(stackSize), 1, 1)
	}
}

// This cleans the stack
func (cg *CodeGenerator) removeStackSpace(stackSize int) {
	if stackSize > STACK_SIZE_MAX {
		appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(STACK_SIZE_MAX), 1, 1)
		cg.removeStackSpace(stackSize - STACK_SIZE_MAX)
	} else {
		appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(stackSize), 1, 1)
	}
}

func (cg *CodeGenerator) debug(message string) {
	if DEBUG == true {
		appendAssembly(cg.currInstrs(), "# "+message, 0, 1)
	}
}

// EVALUATION FUNCTIONS

// Evalutes the RHS of an expression
func (cg *CodeGenerator) evalRHS(t Evaluation, srcReg string) {

	//	cg.debug("----- DEBUG - evalRHS() start -----")

	switch t.(type) {
	// Literals
	case Integer:
		appendAssembly(cg.currInstrs(), "LDR "+srcReg+", ="+strconv.Itoa(int(t.(Integer))), 1, 1)
	case Boolean:
		appendAssembly(cg.currInstrs(), "MOV "+srcReg+", #"+boolInt(bool(t.(Boolean))), 1, 1)
	case Character:
		appendAssembly(cg.currInstrs(), "MOV "+srcReg+", #"+string(t.(Character)), 1, 1)
	case Str:

		//HACK
		if srcReg == "r0" {
			srcReg = "r4"

		}
		appendAssembly(cg.currInstrs(), "LDR "+srcReg+", "+cg.getMsgLabel("", string(t.(Str))), 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)

	case PairLiter:
		//PAIR-LITER NOT DONE
		appendAssembly(cg.currInstrs(), "LDR "+srcReg+", =0", 1, 1)
	case Ident:
		var offset, resType = cg.getIdentOffset(t.(Ident))

		switch resType.(type) {
		case ArrayType:

			//HACK
			if srcReg == "r0" {
				srcReg = "r4"
			}

			appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)

		case PairType:
			appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
		case ConstType:

			switch resType.(ConstType) {
			case Bool, Char:
				appendAssembly(cg.currInstrs(), "LDRSB "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			case Int, String:
				appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)

				//MIGHT NEED TO MOVE OUT
				//appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
			}
		default:
			appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
		}

	case ArrayElem:
		cg.evalArrayElem(t, srcReg, "r5")
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
		appendAssembly(cg.currInstrs(), "MOV "+srcReg+", r0", 1, 1)
	case PairElem:
		cg.evalPairElem(t.(PairElem), srcReg)
	case Call:
		// TODO UNCOMMENT This when you are sure that its not causing the infinite loop
		cg.cgVisitCallStat(t.(Call).Ident, t.(Call).ParamList, srcReg)
	default:
		fmt.Println("ERROR: Expression can not be evaluated")
	}

	//	cg.debug("----- DEBUG - evalRHS() end -----")
}

// Evalute a pair element
func (cg *CodeGenerator) evalPairElem(t PairElem, srcReg string) {

	//Load the address of the pair from the stack
	var offset, _ = cg.getIdentOffset(t.Expr.(Ident))
	appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
	//Check for null pointer deference
	appendAssembly(cg.currInstrs(), "MOV r0, "+srcReg, 1, 1)
	appendAssembly(cg.currInstrs(), "BL p_check_null_pointer", 1, 1)
	cg.dereferenceNullPointer()
	//cg.throwRunTimeError()
	cg.cgVisitPrintStatFuncHelper("p_print_string")

	//Depending on fst or snd , load the address
	switch t.Fsnd {
	case Fst:
		appendAssembly(cg.currInstrs(), "LDR "+srcReg+", ["+srcReg+"]", 1, 1)
	case Snd:
		appendAssembly(cg.currInstrs(), "LDR "+srcReg+", ["+srcReg+", #4]", 1, 1)
	}

	//Double deference
	appendAssembly(cg.currInstrs(), "LDR "+srcReg+", ["+srcReg+"]", 1, 1)

	//if it's a pair load the address or else store on the next available space
	/*switch t.Fsnd {
	case Fst:
	case Snd:

	default:
		//Store on the next available space on the stack
		appendAssembly(cg.currInstrs(), "STR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
	}*/

}

// Evalutes a pair of elements onto the stack
func (cg *CodeGenerator) evalNewPair(fst Evaluation, snd Evaluation, reg1 string, reg2 string) {

	// First allocate memory to store two addresses (8-bytes)
	appendAssembly(cg.currInstrs(), "LDR r0, ="+strconv.Itoa(ADDRESS_SIZE*2), 1, 1)
	appendAssembly(cg.currInstrs(), "BL malloc", 1, 1)

	// Store the address in the free register
	appendAssembly(cg.currInstrs(), "MOV "+reg2+", r0", 1, 1)

	//Get pair sizes
	var typeFst, typeSnd = cg.eval(fst), cg.eval(snd)
	var fstSize, sndSize = sizeOf(typeFst), sizeOf(typeSnd)

	// Load the first element into a register to be stored
	cg.evalRHS(fst, reg1)

	//Allocate memory for the first element
	appendAssembly(cg.currInstrs(), "LDR r0, ="+strconv.Itoa(fstSize), 1, 1)
	appendAssembly(cg.currInstrs(), "BL malloc", 1, 1)

	//Store the first element in the register
	appendAssembly(cg.currInstrs(), "STR "+reg1+", [r0]", 1, 1)

	switch typeFst.(type) {
	case ConstType:
		switch typeFst.(ConstType) {
		case Bool, Char:
			//Store the first element to the newly allocated memory onto the stack
			appendAssembly(cg.currInstrs(), "STRB "+reg1+", [r0]", 1, 1)
		case Int, String:
			//Store the address of allocated memory block of the pair on the stack
			appendAssembly(cg.currInstrs(), "STR r0, ["+reg2+"]", 1, 1)
		}
	default:
		appendAssembly(cg.currInstrs(), "STR "+reg1+", [r0]", 1, 1)
	}

	//Load the second element into a register to be stored
	cg.evalRHS(snd, reg1)

	//Allocate memory for the second element
	appendAssembly(cg.currInstrs(), "LDR r0, ="+strconv.Itoa(sndSize), 1, 1)
	appendAssembly(cg.currInstrs(), "BL malloc", 1, 1)

	switch typeSnd.(type) {
	case ConstType:
		switch typeSnd.(ConstType) {
		case Bool, Char:
			//Store the second element to the newly allocated memory onto the stack
			appendAssembly(cg.currInstrs(), "STRB "+reg1+", [r0]", 1, 1)
		case Int, String:
			//Store the second element to the newly allocated memory onto the stack
			appendAssembly(cg.currInstrs(), "STR "+reg1+", [r0]", 1, 1)
		}
	default:
		appendAssembly(cg.currInstrs(), "STR "+reg1+", [r0]", 1, 1)
	}

	//Store the address of allocated memory block of the pair on the stack
	appendAssembly(cg.currInstrs(), "STR r0, ["+reg2+", #4]", 1, 1)

}

// Evaluates array literals
func (cg *CodeGenerator) evalArrayLiter(typeNode Type, rhs Evaluation, srcReg string, dstReg string) {

	switch rhs.(type) {
	case ArrayLiter:
		//Calculate the amount of storage space required for the array
		// = ((arrayLength(array) * sizeOf(arrayType)) + ADDRESS_SIZE
		var arrayStorage = (len(rhs.(ArrayLiter).Exprs) * sizeOf(typeNode)) + ADDRESS_SIZE

		//Allocate memory for the array
		appendAssembly(cg.currInstrs(), "LDR r0, ="+strconv.Itoa(arrayStorage), 1, 1)
		appendAssembly(cg.currInstrs(), "BL malloc", 1, 1)

		appendAssembly(cg.currInstrs(), "MOV "+dstReg+", r0", 1, 1)

		//Start loading each element in the array onto the stack
		cg.evalArray(rhs.(ArrayLiter).Exprs, srcReg, dstReg, typeNode)
	default:
		fmt.Println("RHS Type not implemented")
	}

}

// Evalutes an array (where t is the type of the array)
func (cg *CodeGenerator) evalArray(array []Evaluation, srcReg string, dstReg string, t Type) {

	var arraySize = len(array)
	// Loop through the array pushing it onto the stack
	for i := 0; i < arraySize; i++ {
		// Array of ,ints,bools,chars,strings
		switch t.(type) {
		case ArrayType:
			cg.evalRHS(array[i], srcReg)
			switch t.(ArrayType).Type {
			case Int:
				appendAssembly(cg.currInstrs(), "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+INT_SIZE*i)+"]", 1, 1)
			case String:
				appendAssembly(cg.currInstrs(), "STR "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+STRING_SIZE*i)+"]", 1, 1)
			case Bool:
				appendAssembly(cg.currInstrs(), "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+BOOL_SIZE*i)+"]", 1, 1)
			case Char:
				appendAssembly(cg.currInstrs(), "STRB "+srcReg+", ["+dstReg+", #"+strconv.Itoa(ARRAY_SIZE+CHAR_SIZE*i)+"]", 1, 1)
			default:
				//Must be an array type
				//Loop through the array
				var offset, _ = cg.getIdentOffset(array[i].(Ident))
				switch t.(ArrayType).Type.(type) {
				case ArrayType:
					appendAssembly(cg.currInstrs(), "LDR "+srcReg+", [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
					appendAssembly(cg.currInstrs(), "STR "+srcReg+", [r4, #"+strconv.Itoa(ARRAY_SIZE+sizeOf(t)*i)+"]", 1, 1)
				}
			}
		}
	}
	// Put the size of the array onto the stack
	appendAssembly(cg.currInstrs(), "LDR "+srcReg+", ="+strconv.Itoa(arraySize), 1, 1)
	// Now store the address of the array onto the stack
	appendAssembly(cg.currInstrs(), "STR "+srcReg+", ["+dstReg+"]", 1, 1)
	appendAssembly(cg.currInstrs(), "STR "+dstReg+", [sp, #"+cg.subCurrP(INT_SIZE)+"]", 1, 1)
}

// Evalutes array elements
func (cg *CodeGenerator) evalArrayElem(t Evaluation, reg1 string, reg2 string) {

	// Define error labels and formatting
	cg.getMsgLabel("", ARRAY_INDEX_NEGATIVE)
	cg.getMsgLabel("", ARRAY_INDEX_LARGE)
	cg.getMsgLabel("", STRING_FORMAT)

	// HACK
	if reg1 == "r0" {
		reg1 = "r4"
	}

	cg.debug("----- DEBUG - evalArrayElem() start -----")

	// Store the address at the next space in the stack
	var offset, _ = cg.getIdentOffset(t.(ArrayElem).Ident)
	appendAssembly(cg.currInstrs(), "ADD "+reg1+", sp, #"+strconv.Itoa(offset), 1, 1)

	//Check the array
	var array = t.(ArrayElem).Exprs

	//Check the bounds of the array
	cg.arrayCheckBounds(array, reg1, reg2)

	// Add message labels
	cg.checkArrayBounds()
	cg.cgVisitPrintStatFuncHelper("p_print_string")

	cg.debug("----- DEBUG - evalArrayElem() end -----")
}

// Evalutes a ord
func (cg *CodeGenerator) evalOrd(node Unop) {
	switch node.Expr.(type) {
	case Ident:
		//If it's an ident
		var offset, _ = cg.getIdentOffset(node.Expr.(Ident))
		appendAssembly(cg.currInstrs(), "LDRSB r0, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
	case ArrayElem:
		fmt.Println("ArrayElem not done for ord")
	case Character:
		fmt.Println("Char not done for ord")
	default:
		fmt.Println("Oh no")
	}
}

// VISIT FUNCTIONS -------------------------------------------------------------

// Visit Program
func (cg *CodeGenerator) cgVisitProgram(node *Program) {
	// Set properties of global scope
	cg.currStack.size = GetScopeVarSize(node.StatList)
	cg.currStack.currP = cg.currStack.size
	cg.currStack.isFunc = false

	// ASSIGN functions to global variable
	// WE WIll only traverse them if they are called
	functionList = node.FunctionList

	// .text
	appendAssembly(cg.currInstrs(), ".text", 0, 2)

	// .global main
	appendAssembly(cg.funcInstrs, ".global main", 0, 1)

	// main:
	appendAssembly(cg.currInstrs(), "main:", 0, 1)

	// Push {lr} to save the caller address
	appendAssembly(cg.currInstrs(), "PUSH {lr}", 1, 1)

	// sub sp, sp, #n to create variable space
	if cg.currStack.size > 0 {
		cg.createStackSpace(cg.globalStack.size)
	}

	// Traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	// add sp, sp, #n to remove variable space
	if cg.currStack.size > 0 {
		cg.removeStackSpace(cg.globalStack.size)
	}

	// ldr r0, =0 to return 0 as the main return
	appendAssembly(cg.currInstrs(), "LDR r0, =0", 1, 1)

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.currInstrs(), "POP {pc}", 1, 1)

	// .ltorg
	appendAssembly(cg.currInstrs(), ".ltorg", 1, 1)
}

// Evaluate a statement
func (cg *CodeGenerator) cgEvalStat(stat interface{}) {
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
		// Something went wrong
	}
}

// Visit Declare node
func (cg *CodeGenerator) cgVisitDeclareStat(node Declare) {
	rhs := node.Rhs

	switch node.DecType.(type) {

	case ConstType:

		cg.evalRHS(rhs, "r4")

		switch node.DecType.(ConstType) {
		case Bool:
			// Using STRB, store it on the stack
			appendAssembly(cg.currInstrs(), "STRB r4, [sp, #"+cg.subCurrP(BOOL_SIZE)+"]", 1, 1)
		case Char:
			// Using STRB, store it on the stack
			appendAssembly(cg.currInstrs(), "STRB r4, [sp, #"+cg.subCurrP(CHAR_SIZE)+"]", 1, 1)
		case Int:
			// Store the value of declaration to stack
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(INT_SIZE)+"]", 1, 1)
		case String:
			// Store the address onto the stack
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(STRING_SIZE)+"]", 1, 1)
		case Pair:
			fmt.Println("Pair not implemented")
		}

	case PairType:
		switch rhs.(type) {
		case NewPair:
			cg.evalNewPair(rhs.(NewPair).FstExpr, rhs.(NewPair).SndExpr, "r5", "r4")
			//Store the address of the pair on the stack
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(PAIR_SIZE)+"]", 1, 1)
		case Ident:
			//TODO: UNFINISHED
			cg.evalRHS(rhs.(Ident), "r4")
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(ADDRESS_SIZE)+"]", 1, 1)
		case PairLiter:
			//Can only be Null
			appendAssembly(cg.currInstrs(), "LDR r4, =0", 1, 1)
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(ADDRESS_SIZE)+"]", 1, 1)
		case Call:
			cg.evalRHS(rhs.(Call), "r4")
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+cg.subCurrP(sizeOf(cg.eval(rhs.(Call))))+"]", 1, 1)
		case PairElem:
			var offset, _ = cg.getIdentOffset(rhs.(ArrayElem).Ident)
			cg.evalPairElem(rhs.(PairElem), "r4")
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
		default:
			appendAssembly(cg.currInstrs(), "Unknown type 1", 1, 1)
		}

	case ArrayType:
		// Evalute an array
		cg.evalArrayLiter(node.DecType, rhs, "r5", "r4")
	default:
		appendAssembly(cg.currInstrs(), "Unknown type 2", 1, 1)
	}

	// Saves Idents offset in the symbol tables offset map
	cg.currSymTable().InsertOffset(string(node.Lhs), cg.currStack.currP)
}

// Visit Assignment node
func (cg *CodeGenerator) cgVisitAssignmentStat(node Assignment) {

	//	cg.debug("----- DEBUG - cgVisitAssignmentStat start -----")

	var rhs = node.Rhs
	var lhs = node.Lhs

	//Evaluate the rhs first
	cg.evalRHS(rhs, "r4")

	// lhs can be
	// IDENT , ARRAY-ELEM , PAIR-ELEM
	switch lhs.(type) {

	case Ident:

		//	cg.debug("----- DEBUG - cgVisitAssignmentStat - Ident start -----")

		ident := lhs.(Ident)
		typeIdent := cg.eval(ident)

		switch typeIdent.(type) {
		case PairType:

			switch rhs.(type) {
			case PairLiter:
				//Store null in the pair
				var offset, _ = cg.getIdentOffset(lhs.(Ident))
				appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			case Ident:
				//Complete
			case NewPair:
				var offset, _ = cg.getIdentOffset(lhs.(Ident))
				cg.evalNewPair(rhs.(NewPair).FstExpr, rhs.(NewPair).SndExpr, "r5", "r4")
				appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			}

		case ConstType:
			offset, _ := cg.getIdentOffset(ident)
			switch typeIdent.(ConstType) {
			case Bool, Char:
				// Using STRB, store it on the stack
				appendAssembly(cg.currInstrs(), "STRB r4, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			case Int, String:
				// Store the value of declaration to stack
				appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
			}
		}

	//	cg.debug("----- DEBUG - cgVisitAssignmentStat - Ident end -----")

	case ArrayElem:

		//	cg.debug("----- DEBUG - cgVisitAssignmentStat - ArrayElem start -----")

		var offset, _ = cg.getIdentOffset(lhs.(ArrayElem).Ident)

		//Have a register point to the start of the array
		appendAssembly(cg.currInstrs(), "ADD r5, sp, #"+strconv.Itoa(offset), 1, 1)

		//Load the index
		cg.evalRHS(lhs.(ArrayElem).Exprs[0], "r6")
		appendAssembly(cg.currInstrs(), "LDR r5, [r5]", 1, 1)

		//r6 = Index
		//r5 = Address of array
		appendAssembly(cg.currInstrs(), "MOV r0, r6", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r1, r5", 1, 1)

		//Branch
		appendAssembly(cg.currInstrs(), "BL p_check_array_bounds", 1, 1)
		cg.checkArrayBounds()

		//What does this instruction do?
		appendAssembly(cg.currInstrs(), "ADD r5, r5, #4", 1, 1)
		//Point to the element to be changed
		appendAssembly(cg.currInstrs(), "ADD r5, r5, r6, LSL #2", 1, 1)

		//Get the type of the RHS
		switch rhs.(type) {
		case Boolean, Character:
			appendAssembly(cg.currInstrs(), "STRB r4, [r5]", 1, 1)
		case Integer, Str:
			appendAssembly(cg.currInstrs(), "STR r4, [r5]", 1, 1)
		case Ident:
			appendAssembly(cg.currInstrs(), "STR r4, [r5]", 1, 1)
		default:
			appendAssembly(cg.currInstrs(), "Type not added", 1, 1)
		}

	//	cg.debug("----- DEBUG - cgVisitAssignmentStat - ArrayElem end -----")

	case PairElem:

		//	cg.debug("----- DEBUG - cgVisitAssignmentStat - PairElem start -----")

		// Load the address of the pair into a register
		var offset, _ = cg.getIdentOffset(lhs.(PairElem).Expr.(Ident))
		appendAssembly(cg.currInstrs(), "LDR r5, [sp, #"+strconv.Itoa(offset)+"]", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r5", 1, 1)

		// Jump
		appendAssembly(cg.currInstrs(), "BL p_check_null_pointer", 1, 1)
		cg.dereferenceNullPointer()
		// Check if it's the fst or snd
		switch lhs.(PairElem).Fsnd {
		case Fst:
			appendAssembly(cg.currInstrs(), "LDR r5, [r5]", 1, 1)
		case Snd:
			appendAssembly(cg.currInstrs(), "LDR r5, [r5, #4]", 1, 1)
		}
		// Store the value into the pair
		appendAssembly(cg.currInstrs(), "STR r4, [r5]", 1, 1)

	//	cg.debug("----- DEBUG - cgVisitAssignmentStat - PairElem end -----")

	default:
		fmt.Println("its neither")
	}

	//	cg.debug("----- DEBUG - cgVisitAssignmentStat end -----")
}

// Visit Read node
func (cg *CodeGenerator) cgVisitReadStat(node Read) {

	switch node.AssignLHS.(type) {
	case Ident:
		constType := cg.eval(node.AssignLHS.(Ident))
		offset, _ := cg.getIdentOffset(node.AssignLHS.(Ident))
		appendAssembly(cg.currInstrs(), "ADD r0, sp, #"+strconv.Itoa(offset), 1, 1)
		switch constType {
		case Char:
			appendAssembly(cg.currInstrs(), "BL p_read_char", 1, 1)
			cg.cgVisitReadStatFuncHelper("p_read_char")
		case Int:
			appendAssembly(cg.currInstrs(), "BL p_read_int", 1, 1)
			cg.cgVisitReadStatFuncHelper("p_read_int")
		}
	case ArrayElem:
		//Complete
	case PairElem:
		switch node.AssignLHS.(PairElem).Fsnd {
		case Fst:
			appendAssembly(cg.currInstrs(), "LDR r4, [sp]", 1, 1)
			appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
			appendAssembly(cg.currInstrs(), "BL p_check_null_pointer", 1, 1)
			cg.dereferenceNullPointer()
			appendAssembly(cg.currInstrs(), "LDR r4, [r4]", 1, 1)
			appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		case Snd:
			appendAssembly(cg.currInstrs(), "LDR r4, [sp]", 1, 1)
			appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
			appendAssembly(cg.currInstrs(), "BL p_check_null_pointer", 1, 1)
			cg.dereferenceNullPointer()
			appendAssembly(cg.currInstrs(), "LDR r4, [r4, #4]", 1, 1)
			appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		}
		switch cg.eval(node.AssignLHS.(PairElem)) {
		case Char:
			appendAssembly(cg.currInstrs(), "BL p_read_char", 1, 1)
			cg.cgVisitReadStatFuncHelper("p_read_char")
		case Int:
			appendAssembly(cg.currInstrs(), "BL p_read_int", 1, 1)
			cg.cgVisitReadStatFuncHelper("p_read_int")
		}
	}
}

// Visit Free node
func (cg *CodeGenerator) cgVisitFreeStat(node Free) {
	//testoffset, _ := cg.getIdentOffset(node.Expr.(PairElem).Expr.(Ident))
	// HACK
	//pointer, _ := strconv.Atoi(cg.subCurrP(ADDRESS_SIZE))
	//val := strconv.Itoa(8 + pointer)
	//	appendAssembly(cg.currInstrs(), "LDR r4, [sp, #"+val+"]", 1, 1)
	// HACK
	appendAssembly(cg.currInstrs(), "LDR r4, [sp]", 1, 1)
	appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	appendAssembly(cg.currInstrs(), "BL p_free_pair", 1, 1)
	cg.cgVisitFreeStatFuncHelper("p_free_pair")
}

// Visit Return node
func (cg *CodeGenerator) cgVisitReturnStat(node Return) {
	cg.evalRHS(node.Expr, "r0")
}

// Visit Exit node
func (cg *CodeGenerator) cgVisitExitStat(node Exit) {
	var reg = "r4"
	cg.evalRHS(node.Expr, reg)
	appendAssembly(cg.currInstrs(), "MOV r0, "+reg, 1, 1)
	appendAssembly(cg.currInstrs(), "BL exit", 1, 1)
}

// Visit Print node
func (cg *CodeGenerator) cgVisitPrintStat(node Print) {

	//cg.debug("----- DEBUG - printStat() start -----")

	// Get value of expr into dstReg
	cg.evalRHS(node.Expr, "r0")
	exprType := cg.eval(node.Expr)

	switch exprType.(type) {
	case ConstType:
		switch exprType.(ConstType) {
		case String:
			// BL p_print_string
			appendAssembly(cg.currInstrs(), "BL p_print_string", 1, 1)
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_string")

		case Int:
			// BL p_print_int
			appendAssembly(cg.currInstrs(), "BL p_print_int", 1, 1)
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_int")

		case Char:
			// BL putchar
			appendAssembly(cg.currInstrs(), "BL putchar", 1, 1)

		case Bool:
			// BL p_print_bool
			appendAssembly(cg.currInstrs(), "BL p_print_bool", 1, 1)
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_bool")
		case Pair:
			// BL p_print_reference
			appendAssembly(cg.currInstrs(), "BL p_print_reference", 1, 1)
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_reference")
		}
	case PairType, ArrayType:

		// BL p_print_reference
		appendAssembly(cg.currInstrs(), "BL p_print_reference", 1, 1)
		// Define relevant print function definition (iff it hasnt been defined)
		cg.cgVisitPrintStatFuncHelper("p_print_reference")

	default:
		appendAssembly(cg.currInstrs(), "Error: type not implemented", 1, 1)
		//	typeOf(expr)
	}

	//cg.debug("----- DEBUG - printStat() end -----")
}

// Visit Println node
func (cg *CodeGenerator) cgVisitPrintlnStat(node Println) {
	cg.cgVisitPrintStat(Print{Expr: node.Expr})
	// BL p_print_ln
	appendAssembly(cg.currInstrs(), "BL p_print_ln", 1, 1)
	// Define relevant print function definition (iff it hasnt been defined)
	cg.cgVisitPrintStatFuncHelper("p_print_ln")

}

// Visit If node
func (cg *CodeGenerator) cgVisitIfStat(node If) {
	fstLabel, sndLabel := cg.getNewLabel(), cg.getNewLabel()
	cg.evalRHS(node.Conditional, "r0")
	appendAssembly(cg.currInstrs(), "CMP r0, #0", 1, 1)
	appendAssembly(cg.currInstrs(), "BEQ "+fstLabel, 1, 1)

	cg.cgVisitScopeStat(Scope{StatList: node.ThenStat})

	appendAssembly(cg.currInstrs(), "B "+sndLabel, 1, 1)
	appendAssembly(cg.currInstrs(), fstLabel+":", 0, 1)

	cg.cgVisitScopeStat(Scope{StatList: node.ElseStat})

	appendAssembly(cg.currInstrs(), sndLabel+":", 0, 1)

}

// Visit While node
func (cg *CodeGenerator) cgVisitWhileStat(node While) {
	fstLabel, sndLabel := cg.getNewLabel(), cg.getNewLabel()

	appendAssembly(cg.currInstrs(), "B "+fstLabel, 1, 1)
	appendAssembly(cg.currInstrs(), sndLabel+":", 0, 1)

	cg.cgVisitScopeStat(Scope{StatList: node.DoStat})

	appendAssembly(cg.currInstrs(), fstLabel+":", 0, 1)
	cg.evalRHS(node.Conditional, "r0")
	appendAssembly(cg.currInstrs(), "CMP r0, #1", 1, 1)
	appendAssembly(cg.currInstrs(), "BEQ "+sndLabel, 1, 1)
}

// Visit Scope node
func (cg *CodeGenerator) cgVisitScopeStat(node Scope) {
	// Amount of bytes on the stack the scope takes up for variables
	varSpaceSize := GetScopeVarSize(node.StatList)
	cg.setNewScope(varSpaceSize)
	// sub sp, sp, #n to create variable space
	if varSpaceSize > 0 {
		appendAssembly(cg.currInstrs(), "SUB sp, sp, #"+strconv.Itoa(varSpaceSize), 1, 1)
	}
	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	// add sp, sp, #n to remove variable space
	if varSpaceSize > 0 {
		appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(varSpaceSize), 1, 1)
	}

	cg.removeCurrScope()

}

// Create a mpa of params to offsets
// make sure visitFunction can take a pointer to this map

// ONLY VISIT FUNCTION IF IT IS CALLED
// IE WE ONLY PUSH ONTO STACK FUNC VARIABLES WHEN A FUNCTION IS CALLED
// but
// WE EXECUTE WHAT IS INSIDE THE FUNCTION REGARDLESS OF WHETHER IT IS CALLED OR NOT
func (cg *CodeGenerator) cgVisitCallStat(ident Ident, paramList []Evaluation, srcReg string) {
	for _, function := range functionList {
		if function.Ident == ident {

			for _, param := range paramList {
				cg.cgVisitParameter(param)
			}

			appendAssembly(cg.currInstrs(), "BL f_"+string(function.Ident), 1, 1)

			offset := cg.cgGetParamSize(paramList)
			appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(offset), 1, 1)

			appendAssembly(cg.currInstrs(), "MOV r4, r0", 1, 1)
			// Remove store line moe to declare
			///TODO
			//appendAssembly(cg.currInstrs(), "STR r4, [sp]", 1, 1)
			cg.cgVisitFunction(*function)
		}
	}
}

func (cg *CodeGenerator) cgGetParamSize(paramList []Evaluation) int {
	totalCount := 0
	for _, param := range paramList {
		totalCount += sizeOf(cg.eval(param))
	}
	return totalCount
}

func (cg *CodeGenerator) cgVisitFunction(node Function) {
	if !cg.isFunctionDefined(node.Ident) {
		cg.addFunctionDef(node.Ident)
	} else {
		return
	}
	varSpaceSize := GetScopeVarSize(node.StatList)
	cg.setNewFuncScope(varSpaceSize, &node.ParameterTypes, node.SymbolTable)

	// f_funcName:
	appendAssembly(cg.currInstrs(), "f_"+string(node.Ident)+":", 0, 1)

	// push {lr} to save the caller address
	appendAssembly(cg.currInstrs(), "PUSH {lr}", 1, 1)

	if varSpaceSize > 0 {
		appendAssembly(cg.currInstrs(), "SUB sp, sp, #"+strconv.Itoa(varSpaceSize), 1, 1)
	}

	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	if varSpaceSize > 0 {
		appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(varSpaceSize), 1, 1)
	}

	// TEST harness uses double POP don't think we need it
	// You're right , we don't need it!
	appendAssembly(cg.currInstrs(), "POP {pc}", 1, 1)
	appendAssembly(cg.currInstrs(), ".ltorg", 1, 2)

	cg.removeFuncScope()
}

// VISIT STATEMENT -------------------------------------------------------------

func (cg *CodeGenerator) cgVisitParameter(node Evaluation) {
	resType := cg.eval(node)
	switch resType {
	case Bool:
		cg.evalRHS(node, "r4")
		appendAssembly(cg.currInstrs(), "STRB r4, [sp, #-"+strconv.Itoa(BOOL_SIZE)+"]!", 1, 1)
	case Char:
		cg.evalRHS(node, "r4")
		appendAssembly(cg.currInstrs(), "STRB r4, [sp, #-"+strconv.Itoa(CHAR_SIZE)+"]!", 1, 1)
	case Int:
		cg.evalRHS(node, "r4")
		appendAssembly(cg.currInstrs(), "STR r4, [sp, #-"+strconv.Itoa(INT_SIZE)+"]!", 1, 1)
	case String:
		cg.evalRHS(node, "r4")
		appendAssembly(cg.currInstrs(), "STR r4, [sp, #"+strconv.Itoa(STRING_SIZE)+"]!", 1, 1)
	case Pair:
		cg.evalRHS(node, "r4")
		appendAssembly(cg.currInstrs(), "STR r4, [sp, #-"+strconv.Itoa(PAIR_SIZE)+"]!", 1, 1)
	default:
		switch resType.(type) {
		case PairType:
			cg.evalRHS(node, "r4")
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #-"+strconv.Itoa(PAIR_SIZE)+"]!", 1, 1)
		case ArrayType:
			cg.evalRHS(node, "r4")
			appendAssembly(cg.currInstrs(), "STR r4, [sp, #-"+strconv.Itoa(ARRAY_SIZE)+"]!", 1, 1)
		}
	}
}

// VISIT EXPRESSIONS -----------------------------------------------------------

// Visit Unop node
func (cg *CodeGenerator) cgVisitUnopExpr(node Unop) {
	switch node.Unary {
	case SUB:
		cg.evalRHS(node.Expr, "r4")
		//Negate the result in the register
		appendAssembly(cg.currInstrs(), "RSBS r4, r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		appendAssembly(cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1)
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case LEN:
		cg.evalRHS(node.Expr, "r4")
		//Assume the RHS is always an array
		//So the RHS type is an Ident
		//Load the length of the array into the register
		appendAssembly(cg.currInstrs(), "LDR r4, [r4]", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case ORD:
		cg.evalOrd(node)
	case CHR:
		cg.evalRHS(node.Expr, "r4")
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case NOT:
		cg.evalRHS(node.Expr, "r4")
		appendAssembly(cg.currInstrs(), "EOR r4, r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	default:
		fmt.Println("oh no")
		fmt.Println(node.Unary)
	}

}

// Visit Binop node
func (cg *CodeGenerator) cgVisitBinopExpr(node Binop) {
	cg.evalRHS(node.Left, "r4")
	appendAssembly(cg.currInstrs(), "PUSH {r4}", 1, 1)
	cg.addExtraOffset(4)
	cg.evalRHS(node.Right, "r4")
	appendAssembly(cg.currInstrs(), "MOV r5, r4", 1, 1)
	appendAssembly(cg.currInstrs(), "POP {r4}", 1, 1)
	cg.subExtraOffset(4)
	switch node.Binary {
	case PLUS:
		appendAssembly(cg.currInstrs(), "ADDS r4, r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case SUB:
		appendAssembly(cg.currInstrs(), "SUBS r4, r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		appendAssembly(cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1)
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case MUL:
		appendAssembly(cg.currInstrs(), "SMULL r4, r5, r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "CMP r5, r4, ASR #31", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		appendAssembly(cg.currInstrs(), "BLNE p_throw_overflow_error", 1, 1)
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case DIV:
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r1, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "BL p_check_divide_by_zero", 1, 1)
		appendAssembly(cg.currInstrs(), "BL __aeabi_idiv", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r4, r0", 1, 1)
		cg.cgVisitBinopExprHelper("p_check_divide_by_zero")
	case MOD:
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r1, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "BL p_check_divide_by_zero", 1, 1)
		appendAssembly(cg.currInstrs(), "BL __aeabi_idivmod", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r1", 1, 1)
		cg.cgVisitBinopExprHelper("p_check_divide_by_zero")
	case AND:
		appendAssembly(cg.currInstrs(), "AND r4, r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case OR:
		appendAssembly(cg.currInstrs(), "ORR r4, r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case LT:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVLT r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVGE r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case GT:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVGT r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVLE r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case LTE:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVLE r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVGT r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case GTE:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVGE r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVLT r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case EQ:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVEQ r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVNE r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	case NEQ:
		appendAssembly(cg.currInstrs(), "CMP r4, r5", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVNE r4, #1", 1, 1)
		appendAssembly(cg.currInstrs(), "MOVEQ r4, #0", 1, 1)
		appendAssembly(cg.currInstrs(), "MOV r0, r4", 1, 1)
	}
}

// cgVisitBinopExpr helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitBinopExprHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// funcLabel:
		appendAssembly(cg.progFuncInstrs, funcName+":", 0, 1)
		switch funcName {
		case "p_throw_overflow_error":
			appendAssembly(cg.progFuncInstrs, "LDR r0, "+cg.getMsgLabel("", OVERFLOW), 1, 1)
			appendAssembly(cg.progFuncInstrs, "BL p_throw_runtime_error", 1, 1)
		case "p_throw_runtime_error":
			//
		case "p_check_divide_by_zero":
			appendAssembly(cg.progFuncInstrs, "PUSH {lr}", 1, 1)
			appendAssembly(cg.progFuncInstrs, "CMP r1, #0", 1, 1)
			appendAssembly(cg.progFuncInstrs, "LDREQ r0, "+cg.getMsgLabel("", DIVIDE_BY_ZERO), 1, 1)
			appendAssembly(cg.progFuncInstrs, "BLEQ p_throw_runtime_error", 1, 1)
			appendAssembly(cg.progFuncInstrs, "POP {pc}", 1, 1)
		default:
		}
		cg.throwRunTimeError()
	}
}

func (cg *CodeGenerator) cgCreateMsgs(instrs *ARMList) map[string]string {
	return nil
}
