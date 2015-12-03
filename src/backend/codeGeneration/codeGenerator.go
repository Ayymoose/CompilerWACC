package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"strconv"
)

// CodeGenerator: Using a constructor that provides an address of the place holder
// for the list of assembly instructions for the program (And the AST with the symbol table),
// the code generator can use the "GenerateCode" function to fill the assembly list
// instrs with the assembly produced by traversing the tree
type CodeGenerator struct {
	root        *Program          // Root of the AST
	instrs      *ARMList          // List of assembly instructions for the program
	msgInstrs   ARMList           // List of assembly instructions to create msg labels
	symTable    SymbolTable       // Used to map variable identifiers to their values
	funcInstrs  ARMList           // List of assembly instructions that define functions and their labels
	globalStack *scopeData        // Stack data for the global scope
	currStack   *scopeData        // Stack data for the current scope
	msgMap      map[string]string // Maps string values to their msg labels
}

// Constructor for the code generator.
func ConstructCodeGenerator(cRoot *Program, cInstrs *ARMList, cSymTable SymbolTable) CodeGenerator {
	cg := CodeGenerator{root: cRoot, instrs: cInstrs, msgInstrs: ARMList{},
		symTable: cSymTable, globalStack: &scopeData{}}
	// The program starts off with the current scope as the global scope
	cg.currStack = cg.globalStack
	return cg
}

// Provides information about the stack in relation to a specific scope
type scopeData struct {
	currP       int        // the current position of the pointer to the stack
	size        int        // size of the variable stack scope space in bytes
	parentScope *scopeData // Address of the parent scope
}

// Decreases current pointer to the stack by n
// Returns new currP as a string (Does not have to be used)
func (cg CodeGenerator) subCurrP(n int) string {
	cg.currStack.currP = cg.currStack.currP - n
	return strconv.Itoa(cg.currStack.currP)
}

// Using the ARMList pointer provided in the constructor,
// this function will fill the slice with an array of assembly instructions
// based on the provided AST
func (cg CodeGenerator) GenerateCode() {
	cg.cgVisitProgram(cg.root)
}

// Returns a msg label value for the strValue using msgMap
// If strValue is not contained in the map then it will be added to the map
// with a new msg label value (which will be returned)
// e.g. =msg_0
func (cg CodeGenerator) getMsgLabel(strValue string) string {
	msgLabel, contained := cg.msgMap[strValue]

	if contained {
		return "=" + msgLabel
	}

	cg.msgMap[strValue] = "msg_" + strconv.Itoa(len(cg.msgMap))

	addMsgLabel(&cg.msgInstrs, cg.msgMap[strValue], strValue)

	return "=" + cg.msgMap[strValue]
}
