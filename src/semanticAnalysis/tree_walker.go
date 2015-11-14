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

func (f *FuncNode) visitFunction(function *FuncNode, st *SymbolTable) {
	// New scope so create new symbol table
	currentSymbolTable := st.New()

	// Type of function
	baseType := function.Type // string
	// lex : string -> token
	// parser : token -> string
	// semantics : string -> token  conversion done in symbol table insert

	// Identifier of function
	identifier := function.Ident // string

	//SymbolTable hashMap will be (k, v) where k is identifier and v is the type
	// parent symbol table contains (functionName, functionType)  ->  parent scope

	// Add all params to local scope
	// can be 0 or more
	for _, parameter := range function.ParamList {
		currentSymbolTable.insert(parameter.Ident, parameter.Type)
	}

	// Add all statements to local scope  -> well depending on stat ;)
	for _, statement := range function.StatList {
		visitStatement(statement, currentSymbolTable)
	}
}

func (s *StatNode) visitStatement(statement StatNode, symbolTable *SymbolTable) {
	switch statType := Node.(type) {
	case DeclarationNode: // add to symbol table
	statType.visitDeclaration(symbolTable *SymbolTable)
	case AssignmentNode: // update symbol table but also traverse parent scope
  statType.visiAssignment(symbolTable *SymbolTable)
	case ReadNode: // traverse scope
	statType.visitRead(symbolTable *SymbolTable)
	case FreeNode: // traverse scope
	statType.visitFree(symbolTable *SymbolTable)
	case ReturnNode: // traverse scope and compare to return type of function
	statType.visitReturn(symbolTable *SymbolTable)
	case ExitNode: // traverse scoppe if identifier else expr which evaluates to integer
	statType.visitExit(symbolTable *SymbolTable)
	case PrintNode: //ts
	statType.visitPrint(symbolTable *SymbolTable)
	case PrintlnNode: //ts
	statType.visitPrintln(symbolTable *SymbolTable)
	case IfNode: //ts
	statType.visitIf(symbolTable *SymbolTable)
	case WhileNode: //ts
	statType.visitWhile(symbolTable *SymbolTable)
	case ScopeNode: // new scope
	statType.visitScope(symbolTable *SymbolTable)
	default:
		"Wrong node type retard"
	}
}

// DeclarationNode
/*
type DeclarationNode struct {
	NodeId
	Pos       grammar.Position
	Type      string // changed from *TypeNode
	Ident     string // changed from *IdentNode
	AssignRHS *AssignRHSNode
}
*/
func (d *DeclarationNode) visitDeclaration(symbolTable *SymbolTable) {
	// Get type
	baseType := declaration.
}


// check for type mismatch  -> int i = true is Wrong
// variable s are case sensitive -> check if variable name is in scope if not variable is undeclared/not in scope
// redeclaration of variable is not allowed
//

// Assignment to vair	 needs to be in scope
/*
begin
  begin
  int x = 5
  end ;
  x = 10
end
*/
// error is undeclared


func (a *AssignmentNode) visitAssignment(symbolTable *SymbolTable) {
	//
	x int  = 2
}

func (r *ReadNode) visitRead(symbolTable *SymbolTable) {
	//
}
func (f *FreeNode) visitFree(symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
func (d *DeclarationNode) visitDeclarationNode(declaration DeclarationNode, symbolTable *SymbolTable) {
	//
}
