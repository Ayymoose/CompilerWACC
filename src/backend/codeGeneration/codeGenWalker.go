package codeGeneration

import . "github.com/nanaasiedu/wacc_19/src/parser"
import . "github.com/nanaasiedu/wacc_19/src/backend/fileWriter"

func (p Program) CGvisitProgram() {
	var assemblyString String = ""

	// CHECK FOR MSGS IN TRAVERSAL

	// traverse all functions
	for _, function := range p.Functionlist {
		function.CGvisitFunction()
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

		CGevalStat(stat)

	}

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(assemblyString, "pop {pc}", 1, 1)

	// .ltorg
	appendAssembly(assemblyString, ".ltorg", 1, 1)
}

func CGevalStat(stat interface{}) {
	switch statType := stat.(type) {
	case Declare:
		stat.(Declare).CGvisitDeclareStat()
	case Assignment:
		stat.(Assignment).CGvisitAssignmentStat()
	case Read:
		stat.(Read).CGvisitReadStat()
	case Free:
		stat.(Free).CGvisitFreeStat()
	case Return:
		stat.(Return).CGvisitReturnStat()
	case Exit:
		stat.(Exit).CGvisitExitStat()
	case Print:
		stat.(Print).CGvisitPrintStat()
	case Println:
		stat.(Println).CGvisitPrintlnStat()
	case If:
		stat.(If).CGvisitIfStat()
	case While:
		stat.(While).CGvisitWhileStat()
	case Scope:
		stat.(Scope).CGvisitScopeStat()
	default:
		""
	}
}
func (f Function) CGvisitFunction() {
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

func (d Declare) CGvisitDeclareStat() {
}

func (a Assignment) CGvisitAssignmentStat() {
}

func (r Read) CGvisitReadStat() {
}

func (f Free) CGvisitFreeStat() {

}

func (r Return) CGvisitReturnStat() {

}

func (e Exit) CGvisitExitStat() {

}

func (p Print) CGvisitPrintStat() {

}

func (p Println) CGvisitPrintlnStat() {

}

func (i If) CGvisitIfStat() {

}

func (w While) CGvisitWhileStat() {

}

func (s Scope) CGvisitScopeStat() {

}

// EXPRESSIONS

func (u Unop) CGvisitUnopExpr() {

}

func (b Binop) CGvisitBinopExpr() {

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
