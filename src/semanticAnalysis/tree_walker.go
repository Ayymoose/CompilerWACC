package semanticAnalysis

//import (
//  tree "github.com/henrykhadass/wacc_19/src/abstractSyntaxTree"
//)

func (ast *ProgramNode) visitProgram() {
	// Create new symbolTable
	st := New(SymbolTable)

	// Visit all function nodes
	for _, function := range ast.FuncNode {
		visitFunction(function, st)
	}

	// Visit all statement nodes
	for _, statement := range ast.StatNode {
		visitStatement(statement)
	}
}

func (f *FuncNode) visitFunction(st *SymbolTable) {
	// New scope so create new symbol table
	nst := New(st)

	// Type of function
	baseType := f.Type
  pass, errorMsgs = checkBaseType(baseType string)
	if !pass {
		// To be implmented - > need to decide on return type
	}

  // Identifier of function
	identifier := function.Ident
	pass, errorMsgs = checkIdentifier(identifier string)
	if !pass {
	  // To be implemented --> need to decide on a return type
	}
}

func checkIdentifier(identifier string) {
	"^[a-zA-Z0-9_]+$"   // + does not allow empty strings
}

func checkBaseType(baseType string) (bool, string) {
	switch baseType {
	case "int":
		return true, ""
	case "bool":
		return true, ""
	case "char":
		return true, ""
	case "string":
		return true, ""
	default:
		return false, baseType + "is not a WACC base type"
	}
}

/*
func (p *)

func (s *StatNode) visitStatement() {

}
*/
