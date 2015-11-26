package parser

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
	fst FSND = iota
	snd
)

// ArrayType struct
type ArrayType struct {
	Type Type
}

// PairType struct
type PairType struct {
	fstType Type
	sndType Type
}

// Program ..
type Program struct {
	functionlist []*Function
	statList     []interface{}
	symbolTable  *SymbolTable
}

// Binop struct
type Binop struct {
	binary int
	left   interface{}
	right  interface{}
}

// Unop struct
type Unop struct {
	unary int
	expr  interface{}
}

type NewPair struct {
	fstExpr interface{}
	sndExpr interface{}
}

// Declare struct
type Declare struct {
	Type Type
	lhs  interface{}
	rhs  interface{}
}

// Assignment struct
type Assignment struct {
	lhs interface{}
	rhs interface{}
}

// If struct
type If struct {
	conditional interface{}
	thenStat    []interface{}
	elseStat    []interface{}
}

// While struct
type While struct {
	conditional interface{}
	doStat      []interface{}
}

type Scope struct {
	symbolTable *SymbolTable
	statlist    []interface{}
}

// Read struct
type Read struct {
	assignLHS interface{} // should be an assignLHS
}

// Free struct
type Free struct {
	expr interface{}
}

// Return struct
type Return struct {
	expr interface{}
}

// Exit struct
type Exit struct {
	expr interface{}
}

// Print struct
type Print struct {
	expr interface{}
}

// Println struct
type Println struct {
	expr interface{}
}

type Call struct {
	ident    string
	exprlist []interface{}
}

/*
type Ident struct {
	name string
}
*/

type PairElem struct {
	fsnd FSND
	expr interface{}
}

type ArrayLiter struct {
	exprs []interface{}
}

type ArrayElem struct {
	ident string
	exprs []interface{}
}
