package ast

type ConstType int
type FSND int

type Type interface {
}

const (
	Int ConstType = iota
	Bool
	Char
	String
	Pair
)

const (
	Fst FSND = iota
	Snd
)

type Function struct {
	Ident          string
	ReturnType     Type
	ParameterTypes []Param
	Statlist       []interface{}
}

type Param struct {
	Ident     string
	ParamType Type
}

// ArrayType struct
type ArrayType struct {
	Type Type
}

// PairType struct
type PairType struct {
	FstType Type
	SndType Type
}

// Program ..
type Program struct {
	Functionlist []*Function
	StatList     []interface{}
	SymbolTable  *SymbolTable
}

// Binop struct
type Binop struct {
	Binary int
	Left   interface{}
	Right  interface{}
}

// Unop struct
type Unop struct {
	Unary int
	Expr  interface{}
}

type NewPair struct {
	FstExpr interface{}
	SndExpr interface{}
}

// Declare struct
type Declare struct {
	Type Type
	Lhs  interface{}
	Rhs  interface{}
}

// Assignment struct
type Assignment struct {
	Lhs interface{}
	Rhs interface{}
}

// If struct
type If struct {
	Conditional interface{}
	ThenStat    []interface{}
	ElseStat    []interface{}
}

// While struct
type While struct {
	Conditional interface{}
	DoStat      []interface{}
}

type Scope struct {
	SymbolTable *SymbolTable
	Statlist    []interface{}
}

// Read struct
type Read struct {
	AssignLHS interface{} // should be an assignLHS
}

// Free struct
type Free struct {
	Expr interface{}
}

// Return struct
type Return struct {
	Expr interface{}
}

// Exit struct
type Exit struct {
	Expr interface{}
}

// Print struct
type Print struct {
	Expr interface{}
}

// Println struct
type Println struct {
	Expr interface{}
}

type Call struct {
	Ident    string
	Exprlist []interface{}
}

/*
type Ident struct {
	name string
}
*/

type PairElem struct {
	Fsnd FSND
	Expr interface{}
}

type ArrayLiter struct {
	Exprs []interface{}
}

type ArrayElem struct {
	Ident string
	Exprs []interface{}
}
