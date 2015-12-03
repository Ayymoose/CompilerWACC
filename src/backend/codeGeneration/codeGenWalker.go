package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
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

		cg.cgEvalStat(stat) // NEED A POINTER HERE TO?

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
		case Pair: /// WE NEED THIS RIGHT?
			/// DO WE NEED FSND ASWELL I.E TYPE FSND SWITCH??
		}
	}
}

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {

	switch node.DecType.(type) {
	case ConstType:
		switch node.DecType.(ConstType) {
		case Bool:

		case Char:

		case Int:

			//fmt.Println("Type", node.Rhs)

			// Load the value of the declaration to R4
			//appendAssembly(cg.instrs, "LDR R4, "+strconv.Itoa(), 1, 1)

			// Store the value of declaration to stack
			appendAssembly(cg.instrs,
				"STR R4, [sp, #"+cg.subCurrP(INT_SIZE)+"])", 1, 1)
		case String:

		}

	}

	// Store the value of the declaration to the stack

}

func (cg CodeGenerator) cgVisitAssignmentStat(node Assignment) {
	switch node.Lhs.(type) {
	case ArrayElem:
	case PairElem:
	default: // Ident
	}

	switch node.Rhs.(type) {
	// expr
	case ConstType:
		switch node.Rhs.(ConstType) {
		case Bool:

		case Char:

		case Int:

		case String:

		case Pair:
		}
	case ArrayElem:
	case Unop:
	case Binop:

	default: // Ident   // NEED TO DEAL WITH '(' EXPR ')' CASE AS WELL?
	}
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
