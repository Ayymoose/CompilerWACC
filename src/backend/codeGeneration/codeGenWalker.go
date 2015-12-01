package codeGeneration

// need help with this import
import . "github.com/henrykhadass/wacc_19/src/parser"

func (p Program) CGvisitProgram() {
	// .text
	//.global main
	//main:

	// traverse all functions
	for _, function := range p.Functionlist {
		function.CGvisitFunction()
	}

	// traverse all statements by switching on statement type
	for _, stat := range p.StatList {
		// write switch statement here
		CGevalStat(stat)

	}
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
