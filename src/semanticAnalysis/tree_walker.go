package semanticAnalysis

func (ast *ProgramNode) visitProgram() {
	// Create new symbolTable
	st := SymbolTable.New()

	// Visit all function nodes
	for _, function := range ast.FuncNode {
		visitFunction(function, st) // do i need *st??
	}

	// Visit all statement nodes
	for _, statement := range ast.StatNode {
		visitStatement(statement)
	}
}

func (f *FuncNode) visitFunction(st *SymbolTable) {
	// New scope so create new symbol table
	currentSymbolTable := st.New()

	// Type of function
	baseType := f.Type // string
	// lex : string -> token
	// parser : token -> string
	// semantics : string -> token  conversion done in symbol table insert

	// Identifier of function
	identifier := f.Ident // string

	//SymbolTable hashMap will be (k, v) where k is identifier and v is the type
	// parent symbol table contains (functionName, functionType)  ->  parent scope

	// Add all params to local scope
	// can be 0 or more
	for _, parameter := range f.ParamList {
		currentSymbolTable.insert(parameter.Ident, parameter.Type)
	}

	// Add all statements to local scope  -> well depending on stat ;)
	for _, parameter := range f.StatList {
		visitStatement(parameter, currentSymbolTable)
	}
}

func (s *StatNode) visitStatement(param ParamNode, symbolTable *SymbolTable) {
	switch statType := Node.(type) {
	case DeclarationNode: // add to symbol table
	case AssignmentNode: // update symbol table but also traverse parent scope
	case ReadNode: // traverse scope
	case FreeNode: // traverse scope
	case ReturnNode: // traverse scope and compare to return type of function
	case ExitNode: // traverse scoppe if identifier else expr which evaluates to integer
	case PrintNode: //ts
	case PrintlnNode: //ts
	case IfNode: //ts
	case WhileNode: //ts
	case ScopeNode: // new scope
	default:
		"Wrong node type retard"
	}
}
