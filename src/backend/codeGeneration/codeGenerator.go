package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"fmt"
	"strconv"
)

// CodeGenerator: Using a constructor that provides an address of the place holder
// for the list of assembly instructions for the program (And the AST with the symbol table),
// the code generator can use the "GenerateCode" function to fill the assembly list
// instrs with the assembly produced by traversing the tree
type CodeGenerator struct {
	root              *Program          // Root of the AST
	instrs            *ARMList          // List of assembly instructions for the program
	msgInstrs         *ARMList          // List of assembly instructions to create msg labels
	symTable          *SymbolTable      // Used to map variable identifiers to their types
	funcSymTable      *SymbolTable      // Used to map function variable indentifier to ther types
	funcInstrs        *ARMList          // List of assembly instructions that define functions and their labels
	progFuncInstrs    *ARMList          // List of assembly instructions that define program generated functions e.g. p_print_string
	progFuncNames     *[]string         // List of program defined function names. Used to avoid program redefinitions
	globalStack       *scopeData        // Stack data for the global scope
	currStack         *scopeData        // Stack data for the current scope
	msgMap            map[string]string // Maps string values to their msg labels
	currentLabelIndex int               // The index of the current generic label (used in control flow functions)
}

// Constructor for the code generator.
func ConstructCodeGenerator(cRoot *Program, cInstrs *ARMList, cSymTable *SymbolTable) CodeGenerator {
	cg := CodeGenerator{root: cRoot, instrs: cInstrs, msgInstrs: new(ARMList),
		funcInstrs: new(ARMList), progFuncInstrs: new(ARMList), progFuncNames: new([]string),
		symTable: cSymTable, globalStack: &scopeData{}}

	// The program starts off with the current scope as the global scope
	cg.currStack = cg.globalStack

	cg.msgMap = make(map[string]string)
	return cg
}

// Evaluates the evaluation using the code generator
func (cg CodeGenerator) eval(e Evaluation) Type {
	eType, _ := e.Eval(cg.root.FunctionList, cg.symTable)
	return eType
}

// Provides information about the stack in relation to a specific scope
type scopeData struct {
	currP       int            // the current position of the pointer to the stack
	size        int            // size of the variable stack scope space in bytes
	parentScope *scopeData     // Address of the parent scope
	isFunc      bool           // true iff the scope date is used for a function scope
	paramMap    *map[Param]int // Map of function parameters to their offset from the start of the function
	// only used if isFunc is true
}

// Creates new scope data for a new scope.
func (cg CodeGenerator) setNewScope(varSpaceSize int) {
	newScope := new(scopeData)
	newScope.currP = varSpaceSize
	newScope.size = varSpaceSize
	newScope.parentScope = cg.currStack
	newScope.isFunc = cg.currStack.isFunc

	if newScope.isFunc {
		newScope.paramMap = cg.currStack.paramMap
	}

	cg.currStack = newScope
	cg.symTable = cg.symTable.GetFrontChild()
}

// Creates new scope data for a new function scope. Sets isFunc to true which
// set the code generator into function mode (So statements evaluate for functions not main)
func (cg CodeGenerator) setNewFuncScope(varSpaceSize int, paramMap *map[Param]int) {
	newScope := new(scopeData)
	newScope.currP = varSpaceSize
	newScope.size = varSpaceSize
	newScope.parentScope = cg.currStack
	newScope.isFunc = true
	newScope.paramMap = paramMap

	cg.currStack = newScope

	//TODO: CODE TO SET CHILD SYMBOL TABLE
}

// Removes current scope and replaces it with the parent scope
func (cg CodeGenerator) removeCurrScope() {
	cg.currStack = cg.currStack.parentScope
	cg.symTable = cg.symTable.Parent
	cg.symTable.RemoveChild()
}

// Returns cg.funcInstrs iff the current scope is a function scope. cg.instrs otherwise
func (cg CodeGenerator) currInstrs() *ARMList {
	if cg.currStack.isFunc {
		return cg.funcInstrs
	} else {
		return cg.instrs
	}
}

// Returns cg.funcSymbolTable iff the current scope is a function scope. cg.symbolTable otherwise
func (cg CodeGenerator) currSymTable() *SymbolTable {
	if cg.currStack.isFunc {
		return cg.funcSymTable
	} else {
		return cg.symTable
	}
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
	cg.buildFullInstr()
}

// Using all the code generators ARMList, instrs is modified to include
// all ARMList instructions in the correct order
func (cg CodeGenerator) buildFullInstr() {
	*cg.instrs = append(*cg.funcInstrs, (*cg.instrs)...)
	*cg.instrs = append(*cg.msgInstrs, (*cg.instrs)...)
	*cg.instrs = append(*cg.instrs, *cg.progFuncInstrs...)
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
	addMsgLabel(cg.msgInstrs, cg.msgMap[strValue], strValue)

	return "=" + cg.msgMap[strValue]
}

// Adds the function name to cg.progFuncNames iff it isnt already in the list
// Returns true iff funcName is already in the list
func (cg CodeGenerator) AddCheckProgName(progName string) bool {
	for _, s := range *cg.progFuncNames {
		if s == progName {
			// if progName has already been defined return true
			return true
		}
	}

	// else add progName to the list
	*cg.progFuncNames = append(*cg.progFuncNames, progName)

	return false
}

// Using symbol tables, a offset to the sp is returned so the ident value can
// be executed
func (cg CodeGenerator) getIdentOffset(ident Ident) (int, Type) {
	return cg.findIdentOffset(ident, cg.currSymTable(), cg.currStack, 0)
}

// Checks if the ident is in the given symbol table. If not the parents are searched
// The function assumes an offset will be found eventually (semantically correct)
func (cg CodeGenerator) findIdentOffset(ident Ident, symTable *SymbolTable,
	scope *scopeData, accOffset int) (int, Type) {
	if symTable == nil {
		fmt.Println("ERROR: incorrect symbol table")
		return 0, Int
	}

	if !symTable.IsOffsetDefined(ident) {
		return cg.findIdentOffset(ident, symTable.Parent, scope.parentScope, accOffset+scope.currP)
	}

	return symTable.GetOffset(string(ident)) + accOffset, symTable.GetTypeOfIdent(ident)

}

func (cg CodeGenerator) getNewLabel() string {
	newLabel := "L" + strconv.Itoa(cg.currentLabelIndex)
	cg.currentLabelIndex++
	return newLabel
}
