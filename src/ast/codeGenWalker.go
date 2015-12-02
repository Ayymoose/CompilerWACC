package ast

import (
	. "backend/filewriter"
	"fmt"
	"strconv"
)

const (
	INT_SIZE    = 4
	ARRAY_SIZE  = 4
	BOOL_SIZE   = 1
	CHAR_SIZE   = 1
	STRING_SIZE = 4
	PAIR_SIZE   = 4
)

type CodeGenerator struct {
	root        *Program
	instrs      *ARMList
	msgInstrs   ARMList
	symTable    SymbolTable
	globalStack *scopeData
	currStack   *scopeData
}

func ConstructCodeGenerator(cRoot *Program, cInstrs *ARMList, cSymTable SymbolTable) CodeGenerator {
	cg := CodeGenerator{root: cRoot, instrs: cInstrs, msgInstrs: ARMList{},
		symTable: cSymTable, globalStack: &scopeData{}}
	cg.currStack = cg.globalStack
	return cg
}

type scopeData struct {
	currP       int // the current position of the pointer to the stack
	size        int // size of the stack space in bytes
	parentScope *scopeData
}

// Decreases current pointer to the stack by n
// Returns new currP as a string
func (cg CodeGenerator) subCurrP(n int) string {
	cg.currStack.currP = cg.currStack.currP - n
	return strconv.Itoa(cg.currStack.currP)
}

// Using the ARMList pointer provided in the constructor,
// this function will fill the slice with an array of assembly instructions
func (cg CodeGenerator) GenerateCode() {
	cg.cgVisitProgram(cg.root)
}

func (cg CodeGenerator) cgVisitProgram(node *Program) {
	// Set properities of global scope
	cg.globalStack.size = GetScopeVarSize(node.StatList)
	cg.globalStack.currP = cg.globalStack.size

	// traverse all functions
	for _, function := range node.Functionlist {
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
	/*funcName := "f_" + node.Ident
	fmt.Println(funcName)
	// PUSH {lr}
	if f.ParameterTypes != nil {
		for _, param := range f.ParameterTypes {
			fmt.Println(param)
			//
		}
	}
	// .ltorg
	*/
}

// statements
/*
func getValue(node Declare) int {
	return node.Rhs
}*/

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {

	switch node.Type.(type) {
	case ConstType:
		switch node.Type.(ConstType) {
		case Bool:

		case Char:

		case Int:

			//fmt.Println("Type", node.Rhs)

			// Load the value of the declaration to R4
			appendAssembly(cg.instrs, "LDR R4, "+strconv.Itoa(), 1, 1)

			// Store the value of declaration to stack
			appendAssembly(cg.instrs,
				"STR R4, [sp, #"+cg.subCurrP(INT_SIZE)+"])", 1, 1)
		case String:

		}

	}

	// Store the value of the declaration to the stack

}

func (cg CodeGenerator) cgVisitAssignmentStat(node Assignment) {
}

func (cg CodeGenerator) cgVisitReadStat(node Read) {
}

func (cg CodeGenerator) cgVisitFreeStat(node Free) {

}

func (cg CodeGenerator) cgVisitReturnStat(node Return) {

}

func (cg CodeGenerator) cgVisitExitStat(node Exit) {

}

func (cg CodeGenerator) cgVisitPrintStat(node Print) {

}

func (cg CodeGenerator) cgVisitPrintlnStat(node Println) {

}

func (cg CodeGenerator) cgVisitIfStat(node If) {

}

func (cg CodeGenerator) cgVisitWhileStat(node While) {

}

func (cg CodeGenerator) cgVisitScopeStat(node Scope) {

}

// EXPRESSIONS

func (cg CodeGenerator) cgVisitUnopExpr(node Unop) {

}

func (cg CodeGenerator) cgVisitBinopExpr(node Binop) {

}

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
