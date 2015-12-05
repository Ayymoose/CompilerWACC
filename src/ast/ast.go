package ast

type errorSlice []error

func (e errorSlice) Error() string {
	var result string
	for _, errs := range e {
		result += errs.Error() + "\n"
	}
	return result
}

type ConstType int
type FSND int

type Type interface {
}

type Evaluation interface {
	eval(functionTable []*Function, symbolTable *SymbolTable) (Type, error)
}

type Statement interface {
	visitStatement(functionTable []*Function, symbolTable *SymbolTable) errorSlice
}

type Ident string

type Integer int
type Character string
type Str string
type Boolean bool

// ArrayType struct
type ArrayType struct {
	Type Type
}

// PairType struct
type PairType struct {
	FstType Type
	SndType Type
}
type PairLiter struct {
}

type Skip struct {
}

const (
	Int ConstType = iota
	Bool
	Char
	String
)

const (
	Fst FSND = iota
	Snd
)

type Function struct {
	Ident          Ident
	ReturnType     Type
	ParameterTypes []Param
	StatList       []Statement
}

type Param struct {
	Ident     Ident
	ParamType Type
}

// Program ..
type Program struct {
	FunctionList []*Function
	StatList     []Statement
	SymbolTable  *SymbolTable
}

// Binop struct
type Binop struct {
	Binary int
	Left   Evaluation
	Right  Evaluation
}

// Unop struct
type Unop struct {
	Unary int
	Expr  Evaluation
}

type NewPair struct {
	FstExpr Evaluation
	SndExpr Evaluation
}

// Declare struct
type Declare struct {
	DecType Type
	Lhs     Ident
	Rhs     Evaluation
}

// Assignment struct
type Assignment struct {
	Lhs Evaluation
	Rhs Evaluation
}

// If struct
type If struct {
	Conditional Evaluation
	ThenStat    []Statement
	ElseStat    []Statement
}

// While struct
type While struct {
	Conditional Evaluation
	DoStat      []Statement
}

type Scope struct {
	SymbolTable *SymbolTable
	StatList    []Statement
}

// Read struct
type Read struct {
	AssignLHS Evaluation // should be an assignLHS
}

// Free struct
type Free struct {
	Expr Evaluation
}

// Return struct
type Return struct {
	Expr Evaluation
}

// Exit struct
type Exit struct {
	Expr Evaluation
}

// Print struct
type Print struct {
	Expr Evaluation
}

// Println struct
type Println struct {
	Expr Evaluation
}

type Call struct {
	Ident     Ident
	ParamList []Evaluation
}

/*
type Ident struct {
	Name string
}
*/
type PairElem struct {
	Fsnd FSND
	Expr Evaluation
}

type ArrayLiter struct {
	Exprs []Evaluation
}

type ArrayElem struct {
	Ident Ident
	Exprs []Evaluation
}
