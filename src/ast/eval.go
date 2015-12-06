package ast

import (
	"errors"
	"fmt"
)

func (value Call) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	for _, function := range functionTable {
		if value.Ident == function.Ident {
			if len(value.ParamList) != len(function.ParameterTypes) {
				return nil, errors.New("Different number of parameters in Call and Function Definition")
			}
			for ind := range value.ParamList {
				exprType, err := value.ParamList[ind].Eval(functionTable, symbolTable)
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

func (value ArrayLiter) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	//	fmt.Println("ARRRAU LITER", value.Exprs)
	var currType Type
	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].Eval(functionTable, symbolTable)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.Eval(functionTable, symbolTable)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, errors.New("Array has mixed types")
			}
		}
		//		fmt.Println("DID WE GET HERE?", currType.typeString())
		return ArrayType{Type: currType}, nil
	}
	return nil, nil
}

func (value NewPair) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	fstTyp, err1 := value.FstExpr.Eval(functionTable, symbolTable)
	sndTyp, err2 := value.SndExpr.Eval(functionTable, symbolTable)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	return PairType{FstType: fstTyp, SndType: sndTyp}, nil
}

func (value PairElem) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	fstsnd := value.Fsnd
	exprTyp, err := value.Expr.Eval(functionTable, symbolTable)
	if err != nil {
		return nil, err
	}
	switch exprTyp.(type) {
	case PairType:
		switch fstsnd {
		case Fst:
			return exprTyp.(PairType).FstType, nil
		case Snd:
			return exprTyp.(PairType).SndType, nil
		}
	}
	return nil, errors.New("Cannot get elem of non pair type")
}

func (value Integer) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Int, nil
}

func (value PairLiter) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Pair, nil
}

func (value Str) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return String, nil
}

func (value Character) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Char, nil
}

func (value Boolean) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	return Bool, nil
}

func (value ArrayElem) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	var currType Type

	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].Eval(functionTable, symbolTable)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.Eval(functionTable, symbolTable)
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
		case ConstType:
			if arrayTyp.(ConstType) == String && len(value.Exprs) == 1 {
				return Char, nil
			}
		default:
			return nil, errors.New("Too many nested indexes in array")
		}
	}
	return arrayTyp, nil
}

func (value Ident) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	//	fmt.Println(symbolTable, "::::", value)
	if symbolTable.isDefined(value) {
		return symbolTable.getTypeOfIdent(value), nil
	}
	//fmt.Println("Supposed to GET HERE IDENT EVAL")
	return nil, errors.New("Identifier not declared")
}

func (binop Binop) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	typl, err := binop.Left.Eval(functionTable, symbolTable)
	typr, err2 := binop.Right.Eval(functionTable, symbolTable)
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
			return nil, errors.New("Right expr of binary bool operation is not an Bool: " + typr.typeString() + " " + fmt.Sprint(binop.Right))
		}
		return Bool, nil
	case LT, LTE, GT, GTE:
		if typl != Int && typl != Char {
			return nil, errors.New("Left expr of binary conditional operation is not an Int or Char")
		}
		if typr != Int && typr != Char {
			return nil, errors.New("Right expr of binary conditional operation is not an Int or Char")
		}
		if !typesMatch(typl, typr) {
			return nil, errors.New("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	case EQ, NEQ:
		if !typesMatch(typl, typr) {
			return nil, errors.New("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	default:
		return nil, errors.New("Binary operation not recognised")
	}
}

func (value Unop) Eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error) {
	typExpr, err := value.Expr.Eval(functionTable, symbolTable)
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
			return Int, nil
		}
		return nil, errors.New("Cannot perform len on non Array expression")
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

func typesMatch(type1 Type, type2 Type) bool {
	switch type1.(type) {
	case PairType:
		if type2 == Pair {
			return true
		}
	}
	switch type2.(type) {
	case PairType:
		if type1 == Pair {
			return true
		}
	}
	return type1 == type2
}
