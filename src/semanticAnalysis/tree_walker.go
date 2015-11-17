package semanticAnalysis

import (
	"grammar"
	ast "abstractSyntaxTree"
)

type FuctionMap struct {
	functionMap map[string][]grammar.Type
}

func NewFuncMap() *FunctionMap {
    return &FunctionMap{functionMap : make(map[string][]grammar.Type)}
}

func (exprNode *ast.ExprNode) evalExpr(SymbolTable *SymbolTable) (grammar.Type, []string) {
	evalElem := exprNode.ExprElem

	switch nodeType := evalElem.(type) {
		case IdentNode:
			if !SymbolTable.isDefined((evalElem.(IdentNode).Ident)) {
				return nil,  []string{"character is not defined"}
			}
			if 	len(SymbolTable.getTypeOfIdent((evalElem.(IdentNode).Ident))) > 1 {
				return nil, []string{"Can't eval ident for fucntion"}
				} else {
					return 	SymbolTable.getTypeOfIdent((evalElem.(IdentNode).Ident))[0], nil
				}
		case IntLiterNode:
			return grammar.BaseInt, nil
		case BoolLiterNode:
			return  grammar.BaseBool, nil
		case CharLiterNode:
			return grammar.BaseChar, nil
		case StringLiterNode:
			return grammar.BaseString, nil
		case ArrayElemNode:
			ident := (evalElem.(ArrayElemNode)).Ident
			if !SymbolTable.isDefined(ident) {
				return nil,  []string{"character is not defined"}
			}
			if	x := SymbolTable.getTypeOfIdent(ident); grammar.ArrayType != SymbolTable.getTypeOfIdent(ident).(type) {
				return nil, []string{"not array type"}
			} else if x.Dimen < len((evalElem.(ArrayElemNode)).Expr) {
						return nil, []string{"array length too small"}
			} else if  x.Dimen > len((evalElem.(ArrayElemNode)).Expr) {
				return grammar.ArrayType{x.Dimen-len((evalElem.(ArrayElemNode)).Expr), x.BaseType}, nil
			} else {
				return x.BaseType, nil
			}
		case PairLiterNode:
//			return evalElem.(PairLiterNode).PairLiter
			return grammar.PairType{}, nil
		case UnaryOpExprNode:
			itemType, errs := (evalElem.(UnaryOpExprNode)).Expr.evalExp(SymbolTable)
			if errs != nil {
				return nil,errs
			}
		switch itemType {
	  	case grammar.BaseBool:
				if (evalElem.(UnaryOpExprNode)).UnaryOper.UnaryOperElem.(type) != NotNode {
					return nil, []string{"Wrong unary operation"}
			 	} else {
			 		return grammar.BaseBool, nil
			 	}
		 	case grammar.BaseInt:
			 	if (evalElem.(UnaryOpExprNode)).UnaryOper.UnaryOperElem.(type) == NegNode {
					return grammar.BaseInt, nil
				} else if (evalElem.(UnaryOpExprNode)).UnaryOper.UnaryOperElem.(type) == ChrNode {
					return grammar.BaseChar, nil
				} else {
					 return nil, []string{"Wrong unary operation"}
				}
			case grammar.BaseChar:
				if (evalElem.(UnaryOpExprNode)).UnaryOper.UnaryOperElem.(type) == OrdNode {
			 		return grammar.BaseInt, nil
		 		} else {
					return nil, []string{"Wrong unary operation"}
		 		}
			case grammar.BaseString, grammar.ArrayType:
				if (evalElem.(UnaryOpExprNode)).UnaryOper.UnaryOperElem.(type) == LenNode {
			 		return grammar.BaseInt, nil
		 		} else {
					return nil, []string{"Wrong unary operation"}
		 		}
				default:
					return nil, []string{"Cannot perform unary type of this type"}
				}
		case EBENode:
			x, errsX := (evalElem.(EBENode))[0].evalExpr(SymbolTable)
			y, errsY := (evalElem.(EBENode))[1].evalExpr(SymbolTable)
			if errsX != nil {
				return nil, errsX
			}
			if errsY != nil {
				return nil, errsY
			}
			if x != y {
				return nil, []string{"Non matching types on binop expr"}
			}
			switch x {
				case grammar.BaseBool:
					switch (evalElem.(EBENode)).BinaryOper.BinaryOperElem {
						case EQNode, NEQNode, AndNode, OrNode:
							return grammar.BaseBool, nil
						default:
							return nil, []string{"Not valid binop on booleans"}
					}
				case grammar.BaseInt:
					switch (evalElem.(EBENode)).BinaryOper.BinaryOperElem {
						case SubNode, AddNode, ModNode, DivNode, MultNode:
							return grammar.BaseInt, nil
						case EQNode, NEQNode, LTENode, LTNode, GTENode, GTNode:
							return grammar.BaseBool, nil
						default:
							return nil, []string{"Not valid binop on booleans"}
					}
				case grammar.BaseChar:
					switch (evalElem.(EBENode)).BinaryOper.BinaryOperElem {
						case EQNode, NEQNode, LTENode, LTNode, GTENode, GTNode:
							return grammar.BaseBool, nil
						default:
							return nil, []string{"Not valid binop on booleans"}
					}
				default:
					switch (evalElem.(EBENode)).BinaryOper.BinaryOperElem {
						case EQNode, NEQNode:
							return grammar.BaseBool, nil
						default:
							return nil, []string{"Not valid binop on booleans"}
					}
			}
		default:
			return nil, []string{"Not an expression"}
	}
}

func (root *ast.ProgramNode) visitProgram() {
	// Create new symbolTable
	st := SymbolTable.New()

	// Before you visit a function need to validate that it is not being redefined
	// Create HashMap of function names and how many times it has been declared
	// Value should always be one
	// If not 1 then function is being redeclared

	// Hashmap of (funcIdent, timesDeclared)
	funcMap := make(map[string]int)
	for _, function := range root.Func {
		funcMap[function.Ident]++
	}

  // Initialise functionMap struct
	fMap := NewFuncMap()


	// Visit all function nodes
	for _, function := range root.Func {
		// Each time pass:
		// Slice of functions
		// Fucntion Hashmap
		// SymbolTable at top level scope
		pass, msg := function.visitFunction(funcMap, st)
		// if visitFuction passes then add function name and parameter types to current symbol table
		// key is function ident (the function name)
		// value is a list of parameter types with first one being the function type
		// Added in order
		if pass {
			st.insertFunctionDeclaration(function, fMap)
		}
	}

	// Visit all statement nodes
	for _, statement := range root.Stat {
		statement.visitStatement(st) // do i need *st
	}
}

func (st *ast.SymbolTable) insertFunctionDeclaration(function *FuncNode, fm *FunctionMap) { // insert function into program symbolTable
	var types []grammar.ItemType

	// Add the function type
	types[0] = function.Type

	// Add all of the function parameter types
	// There can be 0 or more function parameters
	for i, param := range function.ParamList {
		// not starting indexing fro 0 cos that is where function type is located
		types[i+1] = param.Type // grammar.ItemType
	}

	st.insert(function.Ident, types)
	fm.insertIntoFuncMap(function.Ident, types)
}

func (fm *FunctionMap) insertIntoFuncMap(ident string, types []grammar.Type) {
	  fm.functionMap[ident] = types
}

func (f *ast.FuncNode) visitFunction(funcMap map[string]int, st *SymbolTable) (bool, []string) {
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
func (st *ast.SymbolTable) insertFunctionParams(f FuncNode) {
	var typeParam []grammar.ItemType

	for _, param := range f.ParamList {
		typeParam[0] = param.Type
		st.insert(param.Ident, typeParam)
	}
}

func (s *ast.StatementNode) visitStatement(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// DO I NEED TO USE S (STATNODE) HERE RATHER THAN NODE???

	switch statType := s.StatElem.(type) {

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

func (d *ast.DeclarationNode) visitDeclaration(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Get type
	var baseType []grammar.ItemType
	baseType = d.Type // []grammar.ItemType

	// Get identifier
	ident := d.Ident // string

	// Check if ident is already declared in current symbolTable
	// If not recurse on parent symboltables until found (if at all)
	declared := symbolTable.isDefined(ident)
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
/*
func (a *AssignRHSNode) checkAssignRHSType(itemType grammar.ItemType, SymbolTable *SymbolTable) (bool, []string, grammar.ItemType) {
	node := a.AssignRHSElem //Node

	switch node.(type) {
	case ExprNode:
		return (node.(ExprNode)).ExprNodeType(SymbolTable)
	case ArrayLiterNode:
	case PairElemNode:

	}  */

	//	ExprNode
	//	ArrayLiterNode

	// be carefull with dealing with newpairs ASK NANA HOW HE DEALS WITH PAIRS

	//	PairElemNode
	// be carefull with call -> only works on functions so just check return type of function
	/*
			hexpr i
		| harray-liter i
		| ‘newpair’ ‘(’ hexpr i ‘,’ hexpr i ‘)’
		| hpair-elemi
		| ‘call’ hidenti ‘(’ harg-listi? ‘)’


} */
/*
func (exprNode *ExprNode) ExprNodeType(SymbolTable *SymbolTable) (grammar.ItemType, []string) {

	elem := exprNode.ExprElem
	switch elem.(type) {
	case

	}
}   */

func (a *ast.AssignRHSNode) visitAssignRHS() {
	nodeA := a.AssignRHSElem

switch nodeA.(type) {
case ExprNode:
case ArrayLiterNode:
case NewPairNode:
case PairElemNode:


}
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

func (a *ast.AssignmentNode) visitAssignment(symbolTable *SymbolTable) {
	// needs to be declared before you can assign

	// If identifier exists in current symboltable
	// If value is of basetype then update symbol table
	// if types match -> else type error
	// If not of base type then recurse on parent symboltables to find declaration
	// and update identifier in current symbol table if types match
	// If not in scope (nothing is found after traversing parent symbol tables to root)
	// then -> error variable is undeclared

}

func (r *ast.ReadNode) visitRead(symbolTable *SymbolTable) (bool, []string) {
		var errorMsgs []string // An array of error message
	elem := r.AssignLHS.AssignLHSElem

	switch elem.(type) {
		case IdentNode:
			if !symbolTable.isDefined(elem.(IdentNode).Ident) {
					errorMsgs = append(errorMsgs, "Identifier not defined")
			} else if itemType := symbolTable.getTypeOfIdent(key string); !(itemType == grammar.BaseInt || itemType == grammar.BaseChar) {
					errorMsgs = append(errorMsgs, "Type not int or char")
			}
		case ArrayElemNode:
			ident := (elem.(ArrayElemNode)).Ident
			if !SymbolTable.isDefined(ident) {
			errorMsgs = append(errorMsgs, "character is not defined")
			}
			if	x := SymbolTable.getTypeOfIdent(ident); grammar.ArrayType != SymbolTable.getTypeOfIdent(ident).(type) {
				errorMsgs = append(errorMsgs, "not array type")
			} else if x.Dimen < len((evalElem.(ArrayElemNode)).Expr) {
					errorMsgs = append(errorMsgs,"array length too small")
			} else if  x.Dimen > len((evalElem.(ArrayElemNode)).Expr) {
				 		errorMsgs = append(errorMsgs, "cannot read type array")
			} else if !(x.BaseType == grammar.BaseInt || x.BaseType == grammar.BaseChar) {
				errorMsgs = append(errorMsgs, "cannot read type that isn't char or int")
			}
		case PairElemNode:
			itemType, errs :=	(elem.(PairElemNode)).Elem.evalExpr(symbolTable)
			if errs != nil {
				errorMsgs = append(errorMsgs, errs)
				} else if !(itemType == grammar.BaseInt || itemType == grammar.BaseChar) {
					errorMsgs = append(errorMsgs, "cannot read type that isn't char or int")
				}
		default:
					errorMsgs = append(errorMsgs, "Not ident, array-elem or pair-elem")
		}
		if len(errorMsgs) != 0 {
			return false, errorMsgs
		} else {
			return true, nil
		}
	// can only read in a char or int
	// read must be followed by a variable

	// check if variable is declared in current scope then parent scopes
	// if declared and of correct type then store (vairable identifier , value) in symbol table
	// else throw undeclared error
}

func (f *ast.FreeNode) visitFree(symbolTable *SymbolTable) (bool, []string){
	// Free parameter must evaluate to a valid pair <pair(T1, T1)> or a valid array <T[]>

itemType, errs := f.Expr.evalExpr(symbolTable)
		if errs != nil {
			return false, errs
		} else if !(itemType == grammar.ArrayType || itemType == grammar.PairType) {
				return false, []string{"Cannot free type that isn't pair/array"}
		} else {
			return true, nil
		}
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

func (r *ast.ReturnNode) visitReturn(symbolTable *SymbolTable) (bool, []string){

	itemType, errs := r.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	}
		for _, ident := range symbolTable.semanticMap {
			if len(ident) > 0 && itemType == ident[0]{
				return true, nil
			} else {
				return false, []string{"Mismatching return types"}
			}
		}

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

func (e *ast.ExitNode) visitExit(symbolTable *SymbolTable) (bool, []string){

	itemType, errs := e.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else if itemType != grammar.BaseInt {
		return false, []string{"Cannot exit with non int type"}
	} else {
		return true, nil
	}
	// If identifier is of basetype then must be int
	// otherwise error -> cannot exit with non - int value
	// if not of base type then find declaration in current then parent symbolTable scope
	// if found check that type is int
	// if not found undeclared variable error
}

func (p *ast.PrintNode) visitPrint(symbolTable *SymbolTable) {

	itemType, errs := p.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else {
		return true, nil
	}
	// check if the expression of print statement is all of the same type
	// e.e. if print x+y....

	// for each expression
	// check current symbol table then parent symbol tables for the type
	// if undeclared then undeclared variable error
}

func (p *ast.PrintlnNode) visitPrintln(symbolTable *SymbolTable) {

	itemType, errs := p.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else {
		return true, nil
	}
	//Same as above not sure if needed to be implemented
}

func (i *IfNode) visitIf(symbolTable *SymbolTable) (bool, []string) {
	// if <expr> then
	// expr must evaluate to bool
	// check if <expr> variable have been declared in current/parent scope
	//error if not
	// visit statements...
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	baseType, msgs := i.Expr.evalExpr(symbolTable)

	if msgs != nil {
		pass = false
		errorMsgs = append(errorMsgs, "We have an error in if expr")
	}

	firstStat := i.StatList[0].(StatementNode)
	secondStat := i.StatList[1].(StatementNode)

	passFirst, msgs := firstStat.visitStatement(symbolTable)
	passSecond, msgs := secondStat.visitStatement(symbolTable)

	if !passFirst {
		errorMsgs = append(errorMsgs, msgs)
	}

	if !passFirst {
		errorMsgs = append(errorMsgs, msgs)
	}

  pass = passFirst && passSecond
	return pass, errorMsgs
}

func (w *WhileNode) visitWhile(symbolTable *SymbolTable) (bool, []string) {
	//while <expr> do

	// expr must evaluate to bool
	// check if <expr> variable have been declared in current/parent scope
	//error if not
	// visit statements...
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??


	baseType, msgs := w.Expr.evalExpr(symbolTable) //ExprNode
	if msgs != nil {
		pass = false
		errorMsgs = append(errorMsgs, msgs)
	}

	passed, msgs := w.Stat.(StatementNode).visitStatement(symbolTable) //StatementNode
	if !passed {
		pass = false
		errorMsgs = append(errorMsgs, msgs)
	}
	return pass, errorMsgs
}

func (s *ScopeNode) visitScope(symbolTable *SymbolTable) {
	// entering a new scope so create a new symbol table

	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??
	// Create a new symbol table as we enter a new scope
	newScope := symbolTable.New()

	passed, msgs := s.Stat.visitStatement(newScope)
	if !passed {
	  errorMsgs = append(errorMsgs, msgs)
	}
	return pass, errorMsgs
}

//

func (c *CallNode) visitCall(symbolTable *SymbolTable, fMap *FunctionMap) {
	// when calling a function you get value and its type. type must match variable if it is assigned to something
	// when calling a function , the parameter type must  match the original parameter type
	// when calling a function need to have the same number of arguments as orifinal
	//
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Check if ident matches declared function in function Hashmap -> i.e that it exits
  if fMap.functionMap[c.Ident] == nil {
		pass = false
		errorMsgs = append(errorMsgs, "Function with ident "+c.Ident+"is not able to be called as it is undeclared or redefinied")
	}

	// Check if parameter types all match
	if fMap.functionMap[c.Ident] != nil {
		types := fMap.functionMap[c.Ident]
		for _, arg := range c.ArgList {
		  //arg // []ExprNode HOW TO DEAL WITH THIS INDIRECTION ????? DOUBLE FOR LOOP??
			// compareTYPES





		}
	}

	c.Arglist // []Arglistnode






}

/** DONT REALLLY NEED JUST CHECK ON THE TYPES OF THE FUNCTION ARGUMENTS
func (a *ArgListNode) visitArgList(symbolTable *SymbolTable) (bool, []string) {

	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	for _,expr := range a.ExprList {
		passed, msgs := expr.visitExpr(symbolTable)
		pass &= passed
		errorMsgs = append(errorMsgs, msgs)
	}

	return pass, errorMsgs
}
**/
