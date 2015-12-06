package ast

import (
	"errors"
	"fmt"
)

func containsDuplicateFunc(functionTable []*Function) bool {
	freqMap := make(map[Ident]int)
	for _, funcDec := range functionTable {
		freqMap[funcDec.Ident]++
	}
	for _, val := range freqMap {
		if val > 1 {
			return true
		}
	}
	return false
}

func (node If) checkIfReturn(functionTable []*Function, symbolTable *SymbolTable, returnType Type) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		semanticErrors = append(semanticErrors, errors.New("Conditional is not boolean expression"))
	}
	thenSymTab := symbolTable.New()
	for _, thenstat := range node.ThenStat {
		switch thenstat.(type) {
		case Return:
			errs := thenstat.(Return).checkReturnReturn(functionTable, thenSymTab, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		case If:
			errs := thenstat.(If).checkIfReturn(functionTable, thenSymTab, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		default:
			errs := thenstat.visitStatement(functionTable, thenSymTab)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		}
	}
	elseSymTab := symbolTable.New()
	for _, elsestat := range node.ElseStat {
		switch elsestat.(type) {
		case Return:
			errs := elsestat.(Return).checkReturnReturn(functionTable, elseSymTab, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		case If:
			errs := elsestat.(If).checkIfReturn(functionTable, elseSymTab, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		default:
			errs := elsestat.visitStatement(functionTable, elseSymTab)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Return) checkReturnReturn(functionTable []*Function, symbolTable *SymbolTable, returnType Type) errorSlice {
	var semanticErrors []error
	exprTyp, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	} else {
		if !typesMatch(exprTyp, returnType) {
			semanticErrors = append(semanticErrors, errors.New("Return Types do not match: "+exprTyp.typeString()+"::"+returnType.typeString()))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (function *Function) checkFunc(root *Program) errorSlice {
	var semanticErrors []error
	funcSymbolTable := root.SymbolTable.New()
	for _, param := range function.ParameterTypes {
		funcSymbolTable.insert(param.Ident, param.ParamType)
	}
	for ind, stat := range function.StatList {
		if ind == len(function.StatList)-1 {
			switch stat.(type) {
			case If:
				errIf := stat.(If).checkIfReturn(root.FunctionList, funcSymbolTable, function.ReturnType)
				if errIf != nil {
					semanticErrors = append(semanticErrors, errIf)
				}
			case Return:
				errRet := stat.(Return).checkReturnReturn(root.FunctionList, funcSymbolTable, function.ReturnType)
				if errRet != nil {
					semanticErrors = append(semanticErrors, errRet)
				}
			}
		}
		errs := stat.visitStatement(root.FunctionList, funcSymbolTable)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	/*	funretstaterrs := function.checkFuncRetStats(root.FunctionList, funcSymbolTable, function.StatList[len(function.StatList)-1])
		if funretstaterrs != nil {
			//		fmt.Println("GOT HERE funretstaterrs!")
			semanticErrors = append(semanticErrors, funretstaterrs)
		} */
	if len(semanticErrors) > 0 {
		//	fmt.Println("GOT HERE WHATTTTT!")
		return semanticErrors
	}
	//	fmt.Println("GOT HERE CLEAN")
	return nil
}

func (root *Program) SemanticCheck() errorSlice {
	var semanticErrors []error
	if containsDuplicateFunc(root.FunctionList) {
		semanticErrors = append(semanticErrors, errors.New("Program has function redefinitions"))
	}
	for ind, functionProg := range root.FunctionList {
		//	fmt.Println(" At functionProg: ", ind, "/", len(root.FunctionList)-1, "::", functionProg.Ident)
		funcErrs := functionProg.checkFunc(root)
		if funcErrs != nil {
			fmt.Println("Error At Func: ", ind, "/", len(root.FunctionList)-1, "::", functionProg.Ident)
			fmt.Println(funcErrs)
			semanticErrors = append(semanticErrors, funcErrs)
		}
	}
	for _, stat := range root.StatList {
		errs := stat.visitStatement(root.FunctionList, root.SymbolTable)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
		switch stat.(type) {
		case Return:
			semanticErrors = append(semanticErrors, errors.New("Cannot have return statement in main"))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Skip) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	return nil
}

func (node Declare) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	if symbolTable.isDefinedInScope(node.Lhs) {
		semanticErrors = append(semanticErrors, errors.New("Variable already declared::"+string(node.Lhs)))
	}
	exprType, errs := node.Rhs.Eval(functionTable, symbolTable)
	if errs != nil {
		semanticErrors = append(semanticErrors, errs)
	}
	if exprType != nil {
		//	fmt.Println(node.DecType.typeString() + "," + exprType.typeString())
	} else {
		fmt.Println("EXPR TYPE IS NIL")
	}
	if exprType != node.DecType && exprType != nil {
		switch node.DecType.(type) {
		case PairType:
			pairTypeStruct := node.DecType.(PairType)
			if pairTypeStruct.FstType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				case ConstType:
					if exprType.(ConstType) != Pair {
						semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
					}
				default:
					semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
				}
			}
			if pairTypeStruct.SndType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				case ConstType:
					if exprType.(ConstType) != Pair {
						semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
					}
				default:
					semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
				}
			}
		default:
			semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
		}
	}
	symbolTable.insert(node.Lhs, node.DecType)
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Assignment) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	lhsType, errl := node.Lhs.Eval(functionTable, symbolTable)
	rhsType, errr := node.Rhs.Eval(functionTable, symbolTable)
	if errl != nil {
		semanticErrors = append(semanticErrors, errl)
	}
	if errr != nil {
		semanticErrors = append(semanticErrors, errr)
	}
	if rhsType == nil && errl == nil && errr == nil {
		switch lhsType.(type) {
		case ArrayType:
			// Do nothing
		default:
			semanticErrors = append(semanticErrors, errors.New("LHS is not of type Array"+lhsType.typeString()+","+rhsType.typeString()))
		}
	}
	if !typesMatch(lhsType, rhsType) && rhsType != nil && lhsType != nil {
		semanticErrors = append(semanticErrors, errors.New("Assignment types do not match"+lhsType.typeString()+","+rhsType.typeString()))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Read) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.AssignLHS.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	} else if exprTyp != Char && exprTyp != Int {
		semanticErrors = append(semanticErrors, errors.New("Cannot read non Char or Int type"))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Free) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	switch exprTyp.(type) {
	case PairType:
		// Do nothing
	default:
		if exprTyp != Pair {
			semanticErrors = append(semanticErrors, errors.New("Cannot free non pair type"))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Return) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}

	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Exit) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if exprTyp != Int {
		semanticErrors = append(semanticErrors, errors.New("Bad exit expression, must be int"))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Print) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Println) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node If) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		semanticErrors = append(semanticErrors, errors.New("Conditional is not boolean expression"))
	}
	thenSymTab := symbolTable.New()
	for _, thenstat := range node.ThenStat {
		errs := thenstat.visitStatement(functionTable, thenSymTab)
		//	fmt.Println(thenSymTab)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	elseSymTab := symbolTable.New()
	for _, elsestat := range node.ElseStat {
		errs := elsestat.visitStatement(functionTable, elseSymTab)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node While) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		semanticErrors = append(semanticErrors, errors.New("Conditional is not boolean expression"))
	}
	for _, dostat := range node.DoStat {
		errs := dostat.visitStatement(functionTable, symbolTable)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil

}

func (node Scope) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	newSymTab := symbolTable.New()
	for _, stat := range node.StatList {
		errs := stat.visitStatement(functionTable, newSymTab)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (function *Function) checkFuncRetStats(functionTable []*Function, symbolTable *SymbolTable, stat Statement) errorSlice {
	var semanticErrors []error
	switch stat.(type) {
	case Return:
		ExprTyp, err := stat.(Return).Expr.Eval(functionTable, symbolTable)
		if err != nil {
			fmt.Println("WEIRD ERROR TING:", err)
			semanticErrors = append(semanticErrors, err)
		}
		//else {
		//	fmt.Println(stat.(Return).Expr)
		//		fmt.Println(ExprTyp.typeString())
		fmt.Println("GOT HERE, kl: ", stat.(Return).Expr)
		if ExprTyp == nil {
			fmt.Println("ALLOW NIL TING")
		}
		if !typesMatch(ExprTyp, function.ReturnType) {
			semanticErrors = append(semanticErrors, errors.New("Return type does not match:"+ExprTyp.typeString()+"::"+function.ReturnType.typeString()))
		}
		//		}

	case Exit:
		ExprTyp, _ := stat.(Exit).Expr.Eval(functionTable, symbolTable)
		if ExprTyp != Int {
			semanticErrors = append(semanticErrors, errors.New("Invalid Exit statement"))
		}
	case If:
		ifstat := stat.(If)
		thenerr := function.checkFuncRetStats(functionTable, symbolTable, ifstat.ThenStat[len(ifstat.ThenStat)-1])
		elseerr := function.checkFuncRetStats(functionTable, symbolTable, ifstat.ElseStat[len(ifstat.ElseStat)-1])
		if thenerr != nil {
			semanticErrors = append(semanticErrors, thenerr)
		}
		if elseerr != nil {
			semanticErrors = append(semanticErrors, elseerr)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}
