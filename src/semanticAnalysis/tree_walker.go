package semanticAnalysis

import "grammar"

func (ast *ProgramNode) visitProgram() {
	// Create new symbolTable
	st := SymbolTable.New()

	// Before you visit a function need to validate that it is not being redefined
	// Create HashMap of function names and how many times it has been declared
	// Value should always be one
	// If not 1 then function is being redeclared

	// Hashmap of (funcIdent, timesDeclared)
	funcMap := make(map[string]int)
	for _, function := range ast.Func {
		funcMap[function.PosIdentType.Ident]++
	}

	// Visit all function nodes
	for _, function := range ast.Func {
		// Each time pass:
		// Slice of functions
		// Fucntion Hashmap
		// SymbolTable at top level scope
		pass, msg := function.visitFunction(ast.Func, funcMap, st)
		// if visitFuction passes then add function name and parameter types to current symbol table
		// key is function ident (the function name)
		// value is a list of parameter types with first one being the function type
		// Added in order
		if pass {
			st.insertFunctionDeclaration(function)
		}
	}

	// Visit all statement nodes
	for _, statement := range ast.Stat {
		statement.visitStatement(st) // do i need *st
	}
}

func (st *SymbolTable) insertFunctionDeclaration(function *FuncNode) { // insert function into program symbolTable
	var types []grammar.ItemType

	// Add the function type
	types[0] = function.Type

	// Add all of the function parameter types
	// There can be 0 or more function parameters
	for i, param := range function.ParamList {
		// not starting indexing fro 0 cos that is where function type is located
		types[i+1] = param.PosIdentType.Type // grammar.ItemType
	}

	st.insert(function.PosIdentType.Ident, types)
}

func (f *FuncNode) visitFunction(functions []FuncNode, funcMap map[string]int, st *SymbolTable) (bool, []string) {
	// Functions are only defined at the top level scope
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Check that the function called on this method is not being redefined
	key := funcMap[f.PosIdentType.Ident]
	if key > 1 {
		pass = false
		errorMsgs = append(errorMsgs, "Function"+function.Ident+" is being redefined")
		// we can exit here but we are continuing with evaluation for now
	}

	// Create new symbol table for function scope
	funcSymbolTable := st.New()

	// Extract all function parameters and types and insert types into symbol table
	// TODO:: ASK NANA IF HIS PARAMS ARE STRINGS OR grammar.ITEMTYPES  -> I'm assuming they are strings
	funcSymbolTable.insertFunctionParams(f)

	// Visit all function statements
	for _, statement := range function.StatList {
		pass, msgs := statement.visitStatement(funcSymbolTable)
		if !pass {
			errorMsgs = append(errorMsgs, msgs)
		}
	}
	//Empty errorMsgs if all ok
	return pass, errorMsgs
}

// ALSO DO I NEED POINTERS OR NOT???
func (st *SymbolTable) insertFunctionParams(f FuncNode) {
	var typeParam []grammar.ItemType

	for _, param := range f.ParamList {
		typeParam[0] = param.PosIdentType.Type
		s.insert(param.PosIdentType.Ident, typeParam)
	}
}

func (s *StatNode) visitStatement(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// DO I NEED TO USE S (STATNODE) HERE RATHER THAN NODE???
	switch statType := Node.(type) {
	case DeclarationNode:
		pass, errorMsgs = statType.visitDeclaration(symbolTable)
	case AssignmentNode:
		pass, errorMsgs = statType.visitAssignment(symbolTable)
	case ReadNode:
		pass, errorMsgs = statType.visitRead(symbolTable)
	case FreeNode:
		pass, errorMsgs = statType.visitFree(symbolTable)
	case ReturnNode:
		pass, errorMsgs = statType.visitReturn(symbolTable)
	case ExitNode:
		pass, errorMsgs = statType.visitExit(symbolTable)
	case PrintNode:
		pass, errorMsgs = statType.visitPrint(symbolTable)
	case PrintlnNode:
		pass, errorMsgs = statType.visitPrintln(symbolTable)
	case IfNode:
		pass, errorMsgs = statType.visitIf(symbolTable)
	case WhileNode:
		pass, errorMsgs = statType.visitWhile(symbolTable)
	case ScopeNode: // 'begin' <stat> 'end'
		pass, errorMsgs = statType.visitScope(symbolTable)
	default:
		pass = false
		errorMsgs = append(errorMsgs, statType+"is not a valid statement node")
	}
	return pass, errorMsgs
}

func (d *DeclarationNode) visitDeclaration(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Get type
	var baseType []grammar.ItemType
	baseType = d.PosIdentType.Type // []grammar.ItemType

	// Get identifier
	ident := d.PosIdentType.Ident // string

	// Check if ident is already declared in current symbolTable
	// If not recurse on parent symboltables until found (if at all)
	declared, msgs := symbolTable.lookUp(ident)
	/*	if !declared {
			symbolTable.insert(ident, baseType)
		}
	*/
	if declared {
		pass = false
		errorMsgs = append(errorMsgs, ident+"Has already been declared")
	}

	// Get assignment type
	// IMPORTANT := ORRRR =
	// returns boolean if types match WOULD BE GOOD IF IT RETURNED TYPE
	match, msgs, foundType := d.AssignRHSNode.checkAssignRHSType(baseType)
	if !match {
		pass = false
		errorMsgs = append(errorMsgs, "Expected "+baseType+" found "+foundType)
	}

	// Only add to symbol table after all checks have passed
	if !declared {
		symbolTable.insert(ident, baseType)
	}

	// If not in symbol table then add assignment to symbol table iff
	// if basetype of int / char / bool/ string/ matches type of declaration
	// if not then type error
	// if it is an identifier then look up its definition in current symboltable
	// if declaration of identifier not in current symbol table traverse on parent symbol tables
	// if ident not found then undeclared identifier error
	// ABove statement only occurs during assignment not declaration
	// also cannot assign to a function
	return pass, errorMsgs
}

//func checkType(SymbolTable *SymbolTable)

func checkForDeclaration(symbolTable *SymbolTable) (bool, []string) {

	// Check if declaration exists already in symboltable by checking if ident is a key
	// if already exists redeclaration error
	symbolTable.parent == nill{ // == null only at top level scope i.e. we are coming in from a function branch traversal
	// == null only at top level statement traversal
	// when constructing symbol table i believe parent pointer is initialised to nill but need to double check
	// if not then source of errorMsgs
	}

}

func (a *AssignRHSNode) checkAssignRHSType(baseType string) (bool, []string, grammar.ItemType) {
	node := a.AssignRHSElem //Node

	switch node {
	case ExprNode:

	}
	ExprNode
	ArrayLiterNode
	// be carefull with dealing with newpairs ASK NANA HOW HE DEALS WITH PAIRS
	PairElemNode
	// be carefull with call -> only works on functions so just check return type of function
	/*
			hexpr i
		| harray-liter i
		| ‘newpair’ ‘(’ hexpr i ‘,’ hexpr i ‘)’
		| hpair-elemi
		| ‘call’ hidenti ‘(’ harg-listi? ‘)’
	*/

}

func (a *AssignRHSNode) visitAssignRHS() {

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
	// needs to be declared before you can assign

	// If identifier exists in current symboltable
	// If value is of basetype then update symbol table
	// if types match -> else type error
	// If not of base type then recurse on parent symboltables to find declaration
	// and update identifier in current symbol table if types match
	// If not in scope (nothing is found after traversing parent symbol tables to root)
	// then -> error variable is undeclared

}

func (r *ReadNode) visitRead(symbolTable *SymbolTable) {
	// can only read in a char or int
	// read must be followed by a variable

	// check if variable is declared in current scope then parent scopes
	// if declared and of correct type then store (vairable identifier , value) in symbol table
	// else throw undeclared error
}

func (f *FreeNode) visitFree(symbolTable *SymbolTable) {
	// Free parameter must evaluate to a valid pair <pair(T1, T1)> or a valid array <T[]>

	//begin
	//	   pair(int, char) a = newpair(10, 'a') ;
	//     free a
	//end

	// If identifier is in current scope check if identifier is pair or array type
	// If not then error - > cannot free type of type <int/bool/char> must be of type pair/array
	// If not in current symbol table scope then recurse on parents
	// IF not pair or array type then same error as above
	// if identifier not found after traversal then undeclared variable error
}

func (r *ReturnNode) visitReturn(symbolTable *SymbolTable) {
	// return expression must match return type of functionmin current symbol table scope
	//once return statement is executed then exit function immediatley

	// if identifier if of base type
	// check if baseType and function type match -> if not then return type error
	// if identifier is variable
	// check if in current symbol table scope
	// if so check if types of function and variable match
	//if not in current symbol table
	// recurse on parent symbol tables to find declaration
	// check type match
	// if not declared then return undeclared variable error
}
func (e *ExitNode) visitExit(symbolTable *SymbolTable) {
	// If identifier is of basetype then must be int
	// otherwise error -> cannot exit with non - int value
	// if not of base type then find declaration in current then parent symbolTable scope
	// if found check that type is int
	// if not found undeclared variable error

}

func (p *PrintNode) visitPrint(symbolTable *SymbolTable) {
	// check if the expression of print statement is all of the same type
	// e.e. if print x+y....

	// for each expression
	// check current symbol table then parent symbol tables for the type
	// if undeclared then undeclared variable error
}

func (p *PrintlnNode) visitPrintln(symbolTable *SymbolTable) {
	//Same as above not sure if needed to be implemented
}

func (i *IfNode) visitIf(symbolTable *SymbolTable) {
	// if <expr> then
	// expr must evaluate to bool
	// check if <expr> variable have been declared in current/parent scope
	//error if not
	// visit statements...
}

func (w *WhileNode) visitWhile(symbolTable *SymbolTable) {
	//while <expr> do

	// expr must evaluate to bool
	// check if <expr> variable have been declared in current/parent scope
	//error if not
	// visit statements...
}

func (s *ScopeNode) visitScope(symbolTable *SymbolTable) {
	// entering a new scope so create a new symbol table
}

//

func (c *CallNode) visitCall(symbolTable *SymbolTable) {
	// when calling a function you get value and its type. type must match variable if it is assigned to something
	// when calling a function , the parameter type must  match the original parameter type
	// when calling a function need to have the same number of arguments as orifinal
	//
}
