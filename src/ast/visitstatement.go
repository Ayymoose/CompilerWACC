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
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Conditional is not boolean expression"))
	}
	thenSymTab := symbolTable.New()
	symbolTable.Children = append(symbolTable.Children, thenSymTab)
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
	symbolTable.Children = append(symbolTable.Children, elseSymTab)
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
			semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+":Return Types do not match: "+exprTyp.typeString()+"::"+returnType.typeString()))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (function *Function) checkFunc(root *Program) errorSlice {
	var semanticErrors []error
	funcSymbolTable := function.SymbolTable
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
	if len(semanticErrors) > 0 {

		return semanticErrors
	}

	return nil
}

func (root *Program) SemanticCheck() errorSlice {
	var semanticErrors []error
	if containsDuplicateFunc(root.FunctionList) {
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(root.Pos)+"Program has function redefinitions"))
	}
	for _, functionProg := range root.FunctionList {
		funcErrs := functionProg.checkFunc(root)
		if funcErrs != nil {
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
			semanticErrors = append(semanticErrors, errors.New("line:"+"Cannot have return statement in main"))
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
		semanticErrors = append(semanticErrors, errors.New("line: "+fmt.Sprint(node.Pos)+" :Variable already declared::"+string(node.Lhs)))
	}
	exprType, errs := node.Rhs.Eval(functionTable, symbolTable)
	if errs != nil {
		semanticErrors = append(semanticErrors, errs)
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
						semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
					}
				default:
					semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
				}
			}
			if pairTypeStruct.SndType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				case ConstType:
					if exprType.(ConstType) != Pair {
						semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
					}
				default:
					semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
				}
			}
		default:
			semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
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
			semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :LHS is not of type Array"+lhsType.typeString()+","+rhsType.typeString()))
		}
	}
	if !typesMatch(lhsType, rhsType) && rhsType != nil && lhsType != nil {
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Assignment types do not match"+lhsType.typeString()+","+rhsType.typeString()))
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
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Cannot read non Char or Int type"))
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
			semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Cannot free non pair type"))
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
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Bad exit expression, must be int"))
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
		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Conditional is not boolean expression"))
//		semanticErrors = append(semanticErrors, errorConditional(node.FileText,node.Pos))
	}
	thenSymTab := symbolTable.New()
	symbolTable.Children = append(symbolTable.Children, thenSymTab)
	for _, thenstat := range node.ThenStat {
		errs := thenstat.visitStatement(functionTable, thenSymTab)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	elseSymTab := symbolTable.New()
	symbolTable.Children = append(symbolTable.Children, elseSymTab)
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
				semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Conditional is not boolean expression"))
//		semanticErrors = append(semanticErrors, errorConditional(node.FileText, node.Pos))
	}
	whileSymTab := symbolTable.New()
	symbolTable.Children = append(symbolTable.Children, whileSymTab)
	for _, dostat := range node.DoStat {
		errs := dostat.visitStatement(functionTable, whileSymTab)
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
	symbolTable.Children = append(symbolTable.Children, newSymTab)
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
