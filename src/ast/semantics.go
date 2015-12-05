package ast

/*
func (root *Program) semanticCheck() {

}

func (value interface{}) evalRHS(symbolTable *SymbolTable, functionTable []*Function) (Type, error) {
	switch value.(type) {
	case int:
		return Int, nil
	case bool:
		return Bool, nil
	case rune:
		return Char, nil
	case string:
		return String, nil
	case Ident:
		if symbolTable.isDefined(value.Name) {
			return symbolTable.getTypeOfIdent(value.Name), nil
		}
		return nil, error.new("Identifier not declared")
	case ArrayElem:
		return (value.(ArrayElem)).evalArrayElem(symbolTable)
	case Null:
		return Null, nil
	case Binop:
		return (value.(Binop)).evalBinopExpr(symbolTable)
	case Unop:
		return (value.(Unop)).value.evalUnopExpr(symbolTable)
	case ArrayLiter:
	case NewPair:
	case PairElem:
	case Call:

	default:
		return nil, error.new("Type not recognised")
	}

}

func (value Expression) evalExpr(symbolTable *SymbolTable) (Type, error) {
	switch value.(type) {
	case int:
		return Int, nil
	case bool:
		return Bool, nil
	case rune:
		return Char, nil
	case string:
		return String, nil
	case Ident:
		if symbolTable.isDefined(value.Name) {
			return getTypeOfIdent(value.Name), nil
		}
		return nil, error.new("Identifier not declared")
	case ArrayElem:
		return (value.(ArrayElem)).evalArrayElem(symbolTable)
	case Null:
		return Null, nil
	case Binop:
		return (value.(Binop)).evalBinopExpr(symbolTable)
	case Unop:
		return (value.(Unop)).value.evalUnopExpr(symbolTable)
	default:
		return nil, error.new("Type not recognised")
	}
}

func (binop Binop) evalBinopExpr(symbolTable *SymbolTable) (Type, error) {
	typl, err := binop.Left.evalExpr(symbolTable)
	typr, err2 := binop.Right.evalExpr(symbolTable)
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	switch binop.Binary {
	case PLUS, SUB, MUL, DIV, MOD:
		if typl != Int {
			return nil, error.new("Left expr of binary int operation is not an Int")
		}
		if typr != Int {
			return nil, error.new("Right expr of binary int operation is not an Int")
		}
		return Int, nil
	case AND, OR:
		if typl != Bool {
			return nil, error.new("Left expr of binary bool operation is not an Bool")
		}
		if typr != Bool {
			return nil, error.new("Right expr of binary bool operation is not an Bool")
		}
		return Bool, nil
	case LT, LTE, GT, GTE:
		if typl != Int && typl != Char {
			return nil, error.new("Left expr of binary conditional operation is not an Int or Char")
		}
		if typr != Int && typr != Char {
			return nil, error.new("Right expr of binary conditional operation is not an Int or Char")
		}
		if typl != typr {
			return nil, error.new("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	case EQ, NEQ:
		if typl != typr {
			return nil, error.new("Left and right expr of binary conditional operation do not match")
		}
		return Bool, nil
	default:
		return nil, error.new("Binary operation not recognised")
	}
}

func (value Unop) evalUnopExpr(symbolTable *SymbolTable) (Type, error) {
	typExpr, err := value.Expr.evalExpr(symbolTable)
	if err != nil {
		return nil, err
	}
	switch value.Unary {
	case NOT:
		if typExpr != Bool {
			return nil, err("Cannot negate a non Bool expression")
		}
		return Bool, nil
	case SUB:
		if typExpr != Int {
			return nil, err("Cannot sub a non Int expression")
		}
		return Int, nil
	case LEN:
		if typExpr != ArrayType {
			return nil, err("Cannot perform len on non Array expression")
		}
		return Int, nil
	case ORD:
		if typExpr != Char {
			return nil, err("Cannot perform ord on non Char expression")
		}
		return Int, nil
	case CHR:
		if typExpr != Int {
			return nil, err("Cannot perform len on non Int expression")
		}
		return Char, nil
	default:
		return nil, error.new("Unary operation not recognised")
	}
}

func (value ArrayElem) evalArrayElem(symbolTable *SymbolTable) (Type, error) {
	var currType Type
	if len(value.Exprs > 0) {
		currType, err = value.Exprs[0].evalExpr(symbolTable)
		if err != nil {
			return nil, err
		}
		for _, expTypes := range value.Exprs {
			currType2, err2 = value.Exprs[0].evalExpr(symbolTable)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, error.new("Array has mixed types")
			}
		}
	}
	if currType != Int {
		return nil, error.new("Array cannot have non int expr")
	}
	if !symbolTable.isDefined(value.Ident) {
		return nil, error.new("Array not defined, identifier cannot be found")
	}
	arrayTyp := symbolTable.getTypeOfIdent(value.Ident)
	for range value.Exprs {
		switch arrayTyp.(type) {
		case ArrayType:
			arrayTyp = (arrayTyp.(ArrayType)).Type
		default:
			return nil, error.new("Too many nested indexes in array")
		}
	}
	return arrayTyp, nil
}

func (value ArrayLiter) evalArrayLiter(symbolTable *SymbolTable) (Type, error) {
	var currType Type
	if len(value.Exprs > 0) {
		currType, err = value.Exprs[0].evalExpr(symbolTable)
		if err != nil {
			return nil, err
		}
		for _, expTypes := range value.Exprs {
			currType2, err2 = value.Exprs[0].evalExpr(symbolTable)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, error.new("Array has mixed types")
			}
		}
		return ArrayType{Type: currType}, nil
	} else {
		return nil, error.new("Array has no elements")
	}
}

func (value NewPair) evalNewPair(symbolTable *SymbolTable) (Type, error) {
	fstTyp, err1 := value.FstExpr.evalExpr(symbolTable)
	sndTyp, err2 := value.SndExpr.evalExpr(symbolTable)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	return PairType{FstType: fstTyp, SndType: sndTyp}
}

func (value PairElem) evalPairElem(symbolTable *SymbolTable) (Type, error) {
	exprTyp, err := value.Expr.evalExpr(symbolTable)
	if err != nil {
		return nil, err
	}
	return exprTyp, nil
} */

/*
func (value Call) evalCall(symbolTable *SymbolTable, functionTable []*Function) (Type, error) {
	for _, function := range functionTable {
		if value.Ident = {

		}
	}
}

func (functionTable *[]Function) */
