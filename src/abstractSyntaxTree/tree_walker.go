package abstractSyntaxTree

import (
	"fmt"
	"grammar"
)

// This struct holds all the valid functions
// (if any) in a map
type FunctionMap struct {
	functionMap map[string][]grammar.Type
}

// FunctionMap constructor
func NewFuncMap() *FunctionMap {
	return &FunctionMap{functionMap: make(map[string][]grammar.Type)}
}

// Start traversing the tree at the root node
func (root *ProgramNode) VisitProgram() (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??
	symTable := &SymbolTable{parent: nil, semanticMap: make(map[string][]grammar.Type)}
	// Create new symbolTable for global scope
	//symTable := SymbolTable.New()

	// Initialise functionMap struct
	fMap := NewFuncMap()

	// First step in validating functions is to
	// create a frequency hashmap of (funcIdent, timesDeclared)
	funcMap := make(map[string]int)
	for _, function := range root.FuncList {
		funcMap[function.Ident]++
	}

	// Visit all function nodes
	for _, function := range root.FuncList {
		pass, errorMsgs = function.visitFunction(funcMap, symTable)
		if pass {
			// If visitFunction passes insert function declaration into
			// current symbol table as (funcIdent, []Types)
			// The first type in the slice is the function type
			symTable.insertFunctionDeclaration(function, fMap)
		}
	}

	// Visit all statement nodes
	for _, statement := range root.StatList {
		pass, errorMsgs = statement.visitStatement(symTable)
	}
	return pass, errorMsgs
}

func (f FunctionNode) visitFunction(funcMap map[string]int, symTable *SymbolTable) (bool, []string) {
	// Functions are only defined at the top level scope
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Check that the function calling this method is not being redefined
	key := funcMap[f.Ident]
	if key > 1 {
		pass = false
		errorMsgs = append(errorMsgs, "Function "+f.Ident+" is being redefined in top level scope")
	}

	// Create new symbol table for function scope
	funcSymTable := symTable.New()

	// Insert function parameters (ident, type) into current symbol table
	funcSymTable.insertFunctionParams(f)

	// Visit all function statements
	for _, statement := range f.StatList {
		fine, msgs := statement.visitStatement(funcSymTable)
		if !fine {
			pass = false
			errorMsgs = append(errorMsgs, msgs...)
		}
	}
	//Empty errorMsgs if all ok
	return pass, errorMsgs
}

func (st *SymbolTable) insertFunctionDeclaration(function FunctionNode, fm *FunctionMap) {
	var types []grammar.Type

	// Add the function type
	types[0] = function.Type

	// Add all the function parameters. There can be 0 or more
	for i, param := range function.ParamList {
		types[i+1] = param.Type
	}

	// Insert valid function declarations into
	// current symbol table and FunctionMap struct
	st.insert(function.Ident, types)
	fm.insertIntoFuncMap(function.Ident, types)
}

func (fm *FunctionMap) insertIntoFuncMap(ident string, types []grammar.Type) {
	fm.functionMap[ident] = types
}

func (st *SymbolTable) insertFunctionParams(f FunctionNode) {
	var typeParam []grammar.Type

	for _, param := range f.ParamList {
		typeParam[0] = param.Type
		st.insert(param.Ident, typeParam)
	}
}

func (s *StatementNode) visitStatement(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	switch statType := s.Stat.(type) {
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
		errorMsgs = append(errorMsgs, fmt.Sprint(statType)+" is not a valid statement node")
	}
	return pass, errorMsgs
}

func (d *DeclarationNode) visitDeclaration(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	// Get type
	baseType := d.Type

	// Get identifier
	ident := d.Ident

	// Check if Ident is already declared in current symbolTable
	// If not recurse on parent symboltables until found (if at all)
	declared := symbolTable.isDefined(ident)
	if declared {
		pass = false
		errorMsgs = append(errorMsgs,
			"Declaration Error : "+ident+" has already been declared")
	}

	// Evaluate RHS expression
	foundType, errors := d.AssignRHS.evalAssignRHSNode(symbolTable)
	if errors != nil {
		errorMsgs = append(errorMsgs, errors...)
	}
	// Types must match
	if foundType != baseType {
		pass = false
		errorMsgs = append(errorMsgs,
			"Type Error : Expected type "+fmt.Sprint(baseType)+", bur found "+fmt.Sprint(foundType))
	}

	// Only add to symbol table after all checks have passed
	if !declared && (foundType == baseType) {
		symbolTable.insert(ident, []grammar.Type{baseType})
	}
	return pass, errorMsgs
}

func (a *AssignmentNode) visitAssignment(symbolTable *SymbolTable) (bool, []string) {

	var errorMsgs []string // An array of error message
	lhs := a.AssignLHS
	rhs := a.AssignRHS
	evalElem := lhs.AssignLHS
	var lhsType grammar.Type
	//var rhsType grammar.Type

	switch lhs.AssignLHS.(type) {
	case IdentNode:
		if !symbolTable.isDefined(lhs.AssignLHS.(IdentNode).Ident) {
			errorMsgs = append(errorMsgs, "Identifier not defined")
		} else {
			lhsType = symbolTable.getTypeOfIdent((lhs.AssignLHS.(IdentNode)).Ident)[0]
		}
	case ArrayElemNode:
		ident := (lhs.AssignLHS.(ArrayElemNode)).Ident
		if !symbolTable.isDefined(ident) {
			errorMsgs = append(errorMsgs, "character is not defined")
		}
		if x := symbolTable.getTypeOfIdent(ident); x[0].GetMainType() != grammar.TypeArray {
			errorMsgs = append(errorMsgs, "not array type")
		} else if x[0].(grammar.ArrayType).Dimen < len((evalElem.(ArrayElemNode)).ExprList) {
			errorMsgs = append(errorMsgs, "array dimen too small")
		} else if x[0].(grammar.ArrayType).Dimen > len((evalElem.(ArrayElemNode)).ExprList) {
			lhsType = grammar.ArrayType{x[0].(grammar.ArrayType).Dimen - len((evalElem.(ArrayElemNode)).ExprList), x[0].(grammar.ArrayType).BaseType}
		} else {
			lhsType = x[0].(grammar.ArrayType).BaseType
		}
	case PairElemNode:
		itemType, errs := (lhs.AssignLHS.(PairElemNode)).Fst.evalExpr(symbolTable) // THIS ISN't RIGHT
		if errs != nil {
			errorMsgs = append(errorMsgs, errs...)
		} else {
			lhsType = itemType
		}
	default:
		errorMsgs = append(errorMsgs, "LHS not defined correctly")
	}

	typeOfRHS, errs := rhs.evalAssignRHSNode(symbolTable)
	if errs != nil {
		errorMsgs = append(errorMsgs, "Error in RHS of assigment")
	} else if typeOfRHS != lhsType {
		errorMsgs = append(errorMsgs, "Type of RHS and LHS do not match")
	}
	if len(errorMsgs) != 0 {
		return false, errorMsgs
	} else {
		return true, nil
	}
}

func (a *AssignRHSNode) evalAssignRHSNode(symbolTable *SymbolTable) (grammar.Type, []string) {
	rhs := a.AssignRHS

	switch rhs.(type) {
	case ExprNode:
		rightHandType, errs := (rhs.(ExprNode)).evalExpr(symbolTable)
		if errs != nil {
			return nil, errs
		} else {
			return rightHandType, nil
		}
	case ArrayLiterNode:
		typeArray, errs := (rhs.(ArrayLiterNode)).evalArrayLiterNode(symbolTable)
		if errs != nil {
			return nil, errs
		} else {
			return typeArray, nil
		}
	case NewPairNode:
		pairNode := (rhs.(NewPairNode))
		fst := pairNode.ExprList[0]
		//		snd := pairNode.ExprList[1]
		fstType, errsF := fst.evalExpr(symbolTable)
		sndType, errsS := fst.evalExpr(symbolTable)
		if errsF != nil {
			return nil, errsF
		}
		if errsS != nil {
			return nil, errsS
		}
		return grammar.PairType{fstType, sndType}, nil
	case PairElemNode:
		mainType, errs := ((rhs.(PairElemNode)).Fst).evalExpr(symbolTable) // Not right
		if errs != nil {
			return nil, errs
		} else {
			return mainType, nil
		}
	case CallNode:
		returnType := symbolTable.getTypeOfIdent(((rhs.(CallNode)).Ident))[0]
		if returnType == nil {
			return nil, []string{"Function not defined in symbol table"}
		} else {
			return returnType, nil
		}
	default:
		return nil, []string{"Not one of the correct nodes"}
	}
}

func (a ArrayLiterNode) evalArrayLiterNode(symbolTable *SymbolTable) (grammar.Type, []string) {
	expressions := a.ExprList
	fsType, errs := expressions[0].evalExpr(symbolTable)
	if errs != nil {
		return nil, []string{"Error in first expression in arrayliter"}
	}
	for _, expr := range expressions {
		newType, errs := expr.evalExpr(symbolTable)
		if errs != nil {
			return nil, []string{"Error in first expression in arrayliter"}
		} else if fsType != newType {
			return nil, []string{"Cannot make array of different types"}
		}
	}
	return fsType, nil
}

func (r *ReadNode) visitRead(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error message
	elem := r.AssignLHS.AssignLHS

	switch elem.(type) {
	case IdentNode:
		if !symbolTable.isDefined(elem.(IdentNode).Ident) {
			errorMsgs = append(errorMsgs, "Identifier"+elem.(IdentNode).Ident+" is not defined")
		} else if itemType := symbolTable.getTypeOfIdent((elem.(IdentNode)).Ident)[0]; !(itemType == grammar.BaseInt || itemType == grammar.BaseChar) {
			errorMsgs = append(errorMsgs, " type not int or char")
		}
	case ArrayElemNode:
		ident := (elem.(ArrayElemNode)).Ident
		if !symbolTable.isDefined(ident) {
			errorMsgs = append(errorMsgs, "character is not defined")
		}
		if x := symbolTable.getTypeOfIdent(ident); symbolTable.getTypeOfIdent(ident)[0].GetMainType() != grammar.TypeArray {
			errorMsgs = append(errorMsgs, "not array type")
		} else if x[0].(grammar.ArrayType).Dimen < len((elem.(ArrayElemNode)).ExprList) {
			errorMsgs = append(errorMsgs, "array length too small")
		} else if x[0].(grammar.ArrayType).Dimen > len((elem.(ArrayElemNode)).ExprList) {
			errorMsgs = append(errorMsgs, "cannot read type array")
		} else if !(x[0].(grammar.BaseType) == grammar.BaseInt || x[0].(grammar.BaseType) == grammar.BaseChar) {
			errorMsgs = append(errorMsgs, "cannot read type that isn't char or int")
		}
	case PairElemNode:
		itemType, errs := (elem.(PairElemNode)).Fst.evalExpr(symbolTable) // Not right!!!!!!
		if errs != nil {
			errorMsgs = append(errorMsgs, errs...)
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
}

func (f *FreeNode) visitFree(symbolTable *SymbolTable) (bool, []string) {
	// Free parameter must evaluate to a valid pair <pair(T1, T1)> or a valid array <T[]>
	itemType, errs := f.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else if !(itemType.GetMainType() != grammar.TypeArray || itemType.GetMainType() == grammar.TypePair) {
		return false, []string{"Cannot free type that isn't pair/array"}
	} else {
		return true, nil
	}
}

func (r *ReturnNode) visitReturn(symbolTable *SymbolTable) (bool, []string) {

	itemType, errs := r.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	}
	for _, ident := range symbolTable.semanticMap {
		if len(ident) > 0 && itemType == ident[0] {
			return true, nil
		} else if len(ident) > 0 && itemType != ident[0] {
			return false, []string{"Mismatching return types"}
		}
	}
	return false, []string{"No function found in symbol table"}
}

func (e *ExitNode) visitExit(symbolTable *SymbolTable) (bool, []string) {
	itemType, errs := e.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else if itemType != grammar.BaseInt {
		return false, []string{"Cannot exit with non int type"}
	} else {
		return true, nil
	}
}

func (p *PrintNode) visitPrint(symbolTable *SymbolTable) (bool, []string) {
	_, errs := p.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else {
		return true, nil
	}
}

func (p *PrintlnNode) visitPrintln(symbolTable *SymbolTable) (bool, []string) {
	_, errs := p.Expr.evalExpr(symbolTable)
	if errs != nil {
		return false, errs
	} else {
		return true, nil
	}
}

func (i *IfNode) visitIf(symbolTable *SymbolTable) (bool, []string) {
	var pass = true
	var errorMsgs []string
	baseType, msgs := i.Expr.evalExpr(symbolTable)

	if msgs != nil {
		pass = false
		errorMsgs = append(errorMsgs, msgs...)
	}

	if baseType != grammar.BaseBool {
		pass = false
		errorMsgs = append(errorMsgs, "Evaluated type "+fmt.Sprint(baseType)+"not of baseBool type")
	}

	firstStat := i.StatList[0]
	secondStat := i.StatList[1]

	passFirst, msgs := firstStat.visitStatement(symbolTable)
	passSecond, msgs := secondStat.visitStatement(symbolTable)

	if !passFirst {
		pass = false
		errorMsgs = append(errorMsgs, msgs...)
	}

	res := pass && passFirst && passSecond
	return res, errorMsgs
}

func (w *WhileNode) visitWhile(symbolTable *SymbolTable) (bool, []string) {
	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??

	baseType, msgs := w.Expr.evalExpr(symbolTable) //ExprNode
	if msgs != nil {
		pass = false
		errorMsgs = append(errorMsgs, msgs...)
	}

	if baseType != grammar.BaseBool {
		pass = false
		errorMsgs = append(errorMsgs, "Evaluated type "+fmt.Sprint(baseType)+" not of baseBool type")
	}
	for _, statNode := range w.StatList {
		passed, msgs := statNode.visitStatement(symbolTable)
		if !passed {
			pass = false
			errorMsgs = append(errorMsgs, msgs...)
		}
	}
	return pass, errorMsgs
}

func (s *ScopeNode) visitScope(symbolTable *SymbolTable) (bool, []string) {
	// entering a new scope so create a new symbol table

	var errorMsgs []string // An array of error messages
	var pass = true        // source of bugs??
	// Create a new symbol table as we enter a new scope
	newScope := symbolTable.New()
	for _, statNode := range s.StatList {
		passed, msgs := statNode.visitStatement(newScope)
		if !passed {
			errorMsgs = append(errorMsgs, msgs...)
		}
	}
	return pass, errorMsgs
}

func (exprNode ExprNode) evalExpr(symbolTable *SymbolTable) (grammar.Type, []string) {
	evalElem := exprNode.Expr

	switch evalElem.(type) {
	case IdentNode:
		if !symbolTable.isDefined((evalElem.(IdentNode).Ident)) {
			return nil, []string{"character is not defined"}
		}
		if len(symbolTable.getTypeOfIdent((evalElem.(IdentNode).Ident))) > 1 {
			return nil, []string{"Can't eval ident for function"}
		} else {
			return symbolTable.getTypeOfIdent((evalElem.(IdentNode).Ident))[0], nil
		}
	case IntLiterNode:
		return grammar.BaseInt, nil
	case BoolLiterNode:
		return grammar.BaseBool, nil
	case CharLiterNode:
		return grammar.BaseChar, nil
	case StrLiterNode:
		return grammar.BaseString, nil
	case ArrayElemNode:
		ident := (evalElem.(ArrayElemNode)).Ident
		if !symbolTable.isDefined(ident) {
			return nil, []string{"character is not defined"}
		}
		if x := symbolTable.getTypeOfIdent(ident); grammar.TypeArray != symbolTable.getTypeOfIdent(ident)[0].GetMainType() {
			return nil, []string{"not array type"}
		} else if x[0].(grammar.ArrayType).Dimen < len((evalElem.(ArrayElemNode)).ExprList) {
			return nil, []string{"array length too small"}
		} else if x[0].(grammar.ArrayType).Dimen > len((evalElem.(ArrayElemNode)).ExprList) {
			return grammar.ArrayType{x[0].(grammar.ArrayType).Dimen - len((evalElem.(ArrayElemNode)).ExprList), x[0].(grammar.ArrayType).BaseType}, nil
		} else {
			return x[0].(grammar.ArrayType).BaseType, nil
		}
	case PairLiterNode:
		return grammar.PairType{}, nil
	case UnaryOpExprNode:
		itemType, errs := (evalElem.(UnaryOpExprNode)).Expr.evalExpr(symbolTable)
		if errs != nil {
			return nil, errs
		}
		switch itemType.GetMainType() {
		case grammar.TypeBool:
			switch (evalElem.(UnaryOpExprNode)).UnaryOp.UnaryOpType.(type) {
			case NotNode:
				return nil, []string{"Wrong unary operation"}
			default:
				return grammar.BaseBool, nil
			}
		case grammar.TypeInt:
			switch (evalElem.(UnaryOpExprNode)).UnaryOp.UnaryOpType.(type) {
			case NegNode:
				return grammar.BaseInt, nil
			case ChrNode:
				return grammar.BaseChar, nil
			default:
				return nil, []string{"Wrong unary operation"}
			}
		case grammar.TypeChar:
			switch (evalElem.(UnaryOpExprNode)).UnaryOp.UnaryOpType.(type) {
			case OrdNode:
				return grammar.BaseInt, nil
			case LenNode:
				return grammar.BaseInt, nil
			default:
				return nil, []string{"Wrong unary operation"}
			}
		case grammar.TypeString, grammar.TypeArray:
			switch (evalElem.(UnaryOpExprNode)).UnaryOp.UnaryOpType.(type) {
			case LenNode:
				return grammar.BaseInt, nil
			default:
				return nil, []string{"Wrong unary operation"}
			}
		default:
			return nil, []string{"Cannot perform unary type of this type"}
		}
	case BinaryOpExprNode:
		x, errsX := (evalElem.(BinaryOpExprNode)).Expr[0].evalExpr(symbolTable)
		y, errsY := (evalElem.(BinaryOpExprNode)).Expr[1].evalExpr(symbolTable)
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
			switch (evalElem.(BinaryOpExprNode)).BinaryOp.BinaryOpType.(type) {
			case EQNode, NEQNode, AndNode, OrNode:
				return grammar.BaseBool, nil
			default:
				return nil, []string{"Not valid binop on booleans"}
			}
		case grammar.BaseInt:
			switch (evalElem.(BinaryOpExprNode)).BinaryOp.BinaryOpType.(type) {
			case SubNode, AddNode, ModNode, DivNode, MultNode:
				return grammar.BaseInt, nil
			case EQNode, NEQNode, LTENode, LTNode, GTENode, GTNode:
				return grammar.BaseBool, nil
			default:
				return nil, []string{"Not valid binop on booleans"}
			}
		case grammar.BaseChar:
			switch (evalElem.(BinaryOpExprNode)).BinaryOp.BinaryOpType.(type) {
			case EQNode, NEQNode, LTENode, LTNode, GTENode, GTNode:
				return grammar.BaseBool, nil
			default:
				return nil, []string{"Not valid binop on booleans"}
			}
		default:
			switch (evalElem.(BinaryOpExprNode)).BinaryOp.BinaryOpType.(type) {
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
