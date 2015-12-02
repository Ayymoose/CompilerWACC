package ast

import . "backend/filewriter"

type CodeGenerator struct {
	sp       int
	root     *Program
	instrs   *ARMList
	symTable SymbolTable
}

func ConstructCodeGenerator(cRoot *Program, cInstrs *ARMList, cSymTable SymbolTable) CodeGenerator {
	return CodeGenerator{sp: 0, root: cRoot, instrs: cInstrs, symTable: cSymTable}
}

// Using the ARMList pointer provided in the constructor,
// this function will fill the slice with an array of assembly instructions
func (cg CodeGenerator) GenerateCode() {
	cg.cgVisitProgram(cg.root)
}

func (cg CodeGenerator) cgVisitProgram(node *Program) {
	// CHECK FOR MSGS IN TRAVERSAL

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
	appendAssembly(cg.instrs, "push {lr}", 1, 1)

	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {

		cg.cgEvalStat(stat)

	}

	// pop {pc} to restore the caller address as the next instruction
	appendAssembly(cg.instrs, "pop {pc}", 1, 1)

	// .ltorg
	appendAssembly(cg.instrs, ".ltorg", 1, 1)
}

func getScopeVarSize(statList []interface{}) int {
	return 0
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

func (cg CodeGenerator) cgVisitDeclareStat(node Declare) {

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
