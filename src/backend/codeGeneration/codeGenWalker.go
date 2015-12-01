package codeGeneration

// need help with this import

func (p Program) CGvisitProgram() {
	// .text
	//.global main
	//main:

	// traverse all functions
	for _, function := range p.functionlist {
		function.CGvisitFunction()
	}

	// traverse all statements by switching on statement type
	for _, stat := range p.statList {
		// write switch statement here
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
			stat.(Scope).CGvisitWhileStat()
		default:
			""
		}
	}

}

func (f Fuction) CGvisitFunction() {
	funcName := "f_" + f.ident
	// PUSH {lr}
	if f.parameterTypes != nil {
		for _, param := range f.parameterTypes {
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
