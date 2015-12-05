package ast

import (
	"errors"
	"fmt"
)

func (root *Program) SemanticCheck() errorSlice {
	var semanticErrors []error
	for _, stat := range root.StatList {
		errs := stat.visitStatement(root.FunctionList, root.SymbolTable)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
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
	exprType, errs := node.Rhs.eval(functionTable, symbolTable)
	if errs != nil {
		semanticErrors = append(semanticErrors, errs)
	}
	if exprType != nil {
		//	fmt.Println(node.DecType.typeString() + "," + exprType.typeString())
	} else {
		fmt.Println("EXPR TYPE IS NIL")
	}
	if exprType != node.DecType {
		fmt.Println("Got in HEEEREEEEE")
		switch node.DecType.(type) {
		case PairType:
			pairTypeStruct := node.DecType.(PairType)
			if pairTypeStruct.FstType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				default:
					semanticErrors = append(semanticErrors, errors.New("Types in declaration do not match:"+node.DecType.typeString()+","+exprType.typeString()))
				}
			}
			if pairTypeStruct.SndType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
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
	lhsType, errl := node.Lhs.eval(functionTable, symbolTable)
	rhsType, errr := node.Rhs.eval(functionTable, symbolTable)
	if errl != nil {
		semanticErrors = append(semanticErrors, errl)
	}
	if errr != nil {
		semanticErrors = append(semanticErrors, errr)
	}
	if lhsType != rhsType {
		semanticErrors = append(semanticErrors, errors.New("Assignment types do not match"))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Read) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.AssignLHS.eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Free) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Return) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.eval(functionTable, symbolTable)
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
	_, err := node.Expr.eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Print) visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.eval(functionTable, symbolTable)
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
	_, err := node.Expr.eval(functionTable, symbolTable)
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
	cond, err := node.Conditional.eval(functionTable, symbolTable)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		semanticErrors = append(semanticErrors, errors.New("Conditional is not boolean expression"))
	}
	for _, thenstat := range node.ThenStat {
		errs := thenstat.visitStatement(functionTable, symbolTable)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	for _, elsestat := range node.ElseStat {
		errs := elsestat.visitStatement(functionTable, symbolTable)
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
	cond, err := node.Conditional.eval(functionTable, symbolTable)
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
