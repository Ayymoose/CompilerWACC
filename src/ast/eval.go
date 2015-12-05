package ast

import (
	"errors"
)

func (value Call) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	for _, function := range functionTable {
		if value.Ident == function.Ident {
			if len(value.ParamList) != len(function.ParameterTypes) {
				return nil, errors.New("Different number of parameters in Call and Function Definition")
			}
			for ind := range value.ParamList {
				exprType, err := value.ParamList[ind].eval(functionTable, symbolTable)
				if err != nil {
					return nil, err
				}
				if exprType != function.ParameterTypes[ind].ParamType {
					return nil, errors.New("Parameters of call and defintion do not match")
				}
			}
			return function.ReturnType, nil
		}
	}
	return nil, errors.New("No such function defined")
}

func (value ArrayLiter) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	var currType Type
	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].eval(functionTable, symbolTable)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.eval(functionTable, symbolTable)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, errors.New("Array has mixed types")
			}
		}
		return ArrayType{Type: currType}, nil
	}
	return nil, errors.New("Array has no elements")
}

func (value NewPair) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	fstTyp, err1 := value.FstExpr.eval(functionTable, symbolTable)
	sndTyp, err2 := value.SndExpr.eval(functionTable, symbolTable)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	return PairType{FstType: fstTyp, SndType: sndTyp}, nil
}

func (value PairElem) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	exprTyp, err := value.Expr.eval(functionTable, symbolTable)
	if err != nil {
		return nil, err
	}
	switch exprTyp.(type) {
	case PairType:
		return exprTyp.(PairType), nil
	}
	return nil, errors.New("Cannot get elem of non pair type")
}

func (value Integer) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Int, nil
}

func (value PairLiter) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return PairType{}, nil
}

func (value Str) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return String, nil
}

func (value Character) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Char, nil
}

func (value Boolean) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Bool, nil
}

func (value ArrayElem) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	var currType Type
	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].eval(functionTable, symbolTable)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.eval(functionTable, symbolTable)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, errors.New("Array has mixed types")
			}
		}
	}
	if currType != Int {
		return nil, errors.New("Array cannot have non int expr")
	}
	if !symbolTable.isDefined(value.Ident) {
		return nil, errors.New("Array not defined, identifier cannot be found")
	}
	arrayTyp := symbolTable.getTypeOfIdent(value.Ident)
	for _ = range value.Exprs {
		switch arrayTyp.(type) {
		case ArrayType:
			arrayTyp = (arrayTyp.(ArrayType)).Type
		default:
			return nil, errors.New("Too many nested indexes in array")
		}
	}
	return arrayTyp, nil
}

func (value Ident) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	if symbolTable.isDefined(value) {
		return symbolTable.getTypeOfIdent(value), nil
	}
	return nil, errors.New("Identifier not declared")
}

func (binop Binop) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	typl, err := binop.Left.eval(functionTable, symbolTable)
	typr, err2 := binop.Right.eval(functionTable, symbolTable)
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	switch binop.Binary {
	case PLUS, SUB, MUL, DIV, MOD:
		if typl != Int {
			return nil, errors.New("Left expr of binary int operation is not an Int")
		}

		if typr != Int {
			return nil, errors.New("Right expr of binary int operation is not an Int")
		}
		return Int, nil
	case AND, OR:
		if typl != Bool {
			return nil, errors.New("Left expr of binary bool operation is not an Bool")
		}
		if typr != Bool {
			return nil, errors.New("Right expr of binary bool operation is not an Bool")
		}
		return Bool, nil
	case LT, LTE, GT, GTE:
		if typl != Int && typl != Char {
			return nil, errors.New("Left expr of binary conditional operation is not an Int or Char")
		}
		if typr != Int && typr != Char {
			return nil, errors.New("Right expr of binary conditional operation is not an Int or Char")
		}
		if typl != typr {
			return nil, errors.New("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	case EQ, NEQ:
		if typl != typr {
			return nil, errors.New("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	default:
		return nil, errors.New("Binary operation not recognised")
	}
}

func (value Unop) eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	typExpr, err := value.Expr.eval(functionTable, symbolTable)
	if err != nil {
		return nil, err
	}
	switch value.Unary {
	case NOT:
		if typExpr != Bool {
			return nil, errors.New("Cannot negate a non Bool expression")
		}
		return Bool, nil
	case SUB:
		if typExpr != Int {
			return nil, errors.New("Cannot sub a non Int expression")
		}
		return Int, nil
	case LEN:
		switch typExpr.(type) {
		case ArrayType:
			return nil, errors.New("Cannot perform len on non Array expression")
		}
		return Int, nil
	case ORD:
		if typExpr != Char {
			return nil, errors.New("Cannot perform ord on non Char expression")
		}
		return Int, nil
	case CHR:
		if typExpr != Int {
			return nil, errors.New("Cannot perform len on non Int expression")
		}
		return Char, nil
	default:
		return nil, errors.New("Unary operation not recognised")
	}
}
