package codeGeneration

import . "github.com/nanaasiedu/wacc_19/src/parser"
import . "github.com/nanaasiedu/wacc_19/src/backend/fileWriter"

func (p Program) CGvisitProgram(instrs *ARMList) {
	var assemblyString String = ""

	// CHECK FOR MSGS IN TRAVERSAL

	// traverse all functions
	for _, function := range p.Functionlist {
		function.CGvisitFunction(instrs)
	}

	// .text
	appendAssembly(assemblyString, ".text", 0, 2)

	// .global main
	appendAssembly(assemblyString, ".global main", 0, 1)

	//main:
	appendAssembly(assemblyString, "main:", 0, 1)

	// push {lr} to save the caller address
	appendAssembly(assemblyString, "push {lr}", 1, 1)

	// traverse all statements by switching on statement type
	for _, stat := range p.StatList {

		CGevalStat(stat, instrs)

	}

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(assemblyString, "pop {pc}", 1, 1)

	// .ltorg
	appendAssembly(assemblyString, ".ltorg", 1, 1)
}

func CGevalStat(stat interface{}, instrs *ARMList) {
	switch statType := stat.(type) {
	case Declare:
		stat.(Declare).CGvisitDeclareStat(instrs)
	case Assignment:
		stat.(Assignment).CGvisitAssignmentStat(instrs)
	case Read:
		stat.(Read).CGvisitReadStat(instrs)
	case Free:
		stat.(Free).CGvisitFreeStat(instrs)
	case Return:
		stat.(Return).CGvisitReturnStat(instrs)
	case Exit:
		stat.(Exit).CGvisitExitStat(instrs)
	case Print:
		stat.(Print).CGvisitPrintStat(instrs)
	case Println:
		stat.(Println).CGvisitPrintlnStat(instrs)
	case If:
		stat.(If).CGvisitIfStat(instrs)
	case While:
		stat.(While).CGvisitWhileStat(instrs)
	case Scope:
		stat.(Scope).CGvisitScopeStat(instrs)
	default:
		""
	}
}
func (f Function) CGvisitFunction(instrs *ARMList) {
	funcName := "f_" + f.Ident
	// PUSH {lr}
	if f.ParameterTypes != nil {
		for _, param := range f.ParameterTypes {
			//
		}
	}
	// .ltorg
}

// statements

func (d Declare) CGvisitDeclareStat(instrs *ARMList) {
}

func (a Assignment) CGvisitAssignmentStat(instrs *ARMList) {
}

func (r Read) CGvisitReadStat(instrs *ARMList) {
}

func (f Free) CGvisitFreeStat(instrs *ARMList) {

}

func (r Return) CGvisitReturnStat(instrs *ARMList) {

}

func (e Exit) CGvisitExitStat(instrs *ARMList) {

}

func (p Print) CGvisitPrintStat(instrs *ARMList) {

}

func (p Println) CGvisitPrintlnStat(instrs *ARMList) {

}

func (i If) CGvisitIfStat(instrs *ARMList) {

}

func (w While) CGvisitWhileStat(instrs *ARMList) {

}

func (s Scope) CGvisitScopeStat(instrs *ARMList) {

}

// EXPRESSIONS

func (u Unop) CGvisitUnopExpr(instrs *ARMList) {

}

func (b Binop) CGvisitBinopExpr(instrs *ARMList) {

}

func appendAssembly(String *str, String code, int numTabsint, int numNewLines) {
	const default_num_tabs = 1

	for i := 0; i < numTabs+default_num_tabs; i++ {
		*str += "\t"
	}

	*str += code

	for i := 0; i < numNewLines; i++ {
		*str += "\n"
	}
}
