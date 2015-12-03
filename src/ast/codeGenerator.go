package ast

import (
	. "backend/filewriter"
	"strconv"
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
