package abstractSyntaxTree

import (
	grammar "grammar"
)

// NodeId identifies the type of an abstract syntax tree node
type NodeId int

const (
	NodeProgram NodeId = iota
	NodeFunc
	NodeParamList
	NodeParam
	NodeStat
	NodeDeclaration
	NodeAssignment
	NodeRead
	NodeFree
	NodeReturn
	NodeExit
	NodePrint
	NodePrintln
	NodeIf
	NodeWhile
	NodeAssignLHS
	NodeIdent
	NodeArrayElem
	NodePairElem
	NodeAssignRHS
	NodeExpr
	NodeArrayLiter
	NodeNewPair
	NodeCall
	NodeArgList
	NodeType
	NodeBaseType
	NodeInt
	NodeBool
	NodeChar
	NodeString
	NodeArrayType
	NodePairType
	NodePairElemType
	NodeIntLiter
	NodeBoolLiter
	NodeCharLiter
	NodeStrLiter
	NodeCharacter
	NodeEscapedChar
	NodeNull
	NodeBackslash
	NodeTab
	NodeLineFeed
	NodeFormFeed
	NodeCarriageReturn
	NodeDoubleQuote
	NodeSingleQuote
	NodePairLiter
	NodeUnaryOperExpr
	NodeEBE
	NodeUnaryOper
	NodeNot
	NodeNeg
	NodeLen
	NodeOrd
	NodeChr
	NodeBinaryOper
	NodeMult
	NodeDiv
	NodeMod
	NodeAdd
	NodeSub
	NodeGt
	NodeGte
	NodeLt
	NodeLte
	NodeEq
	NodeNeq
	NodeAnd
	NodeOr
	NodeDigit
	NodeIntSign
	NodePlus
	NodeNegate
	NodeComment
)

type Node interface {
	getNodeType() int
}

// Root node of AST
type ProgramNode struct {
	NodeId
	Pos  int
	Func []*FuncNode
	Stat []*StatNode
}

type FuncNode struct {
	NodeId
	Pos       int
	Type      string       // add the type
	Ident     string       // changed from *IdentNode
	ParamList []*ParamNode // *ParamListNode // Optional
	StatList  []*StatNode
}

func (n FuncNode) getNodeType() int {
	return n.Pos
}

/*
type ParamListNode struct {
	NodeId
	Pos   int
	Param []*ParamNode // one or more
}
*/

type ParamNode struct {
	NodeId
	Pos   int
	Type  string
	Ident string // changed from  *IdentNode
}

type StatNode struct {
	NodeId
	Pos      int
	StatElem *Node
}

// TypeIdent
type DeclarationNode struct {
	NodeId
	Pos       int
	Type      *TypeNode
	Ident     string // changed from *IdentNode
	AssignRHS *AssignRHSNode
}

// AssignLHSToRight
type AssignmentNode struct {
	NodeId
	Pos       int
	AssignLHS *AssignLHSNode
	AssignRHS *AssignRHSNode
}

type ReadNode struct {
	NodeId
	Pos       int
	AssignLHS *AssignLHSNode
}

type FreeNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
}

type ReturnNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
}

type ExitNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
}

type PrintNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
}

type PrintlnNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
}

type IfNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
	Stat []*StatNode // two stats
}

type WhileNode struct {
	NodeId
	Pos  int
	Expr *ExprNode
	Stat *StatNode
}

//    ‘begin’ <stat> ‘end’
// SPECIAL NEEDS
type ScopeNode struct {
	NodeId
	StatElem *StatNode
}

type AssignLHSNode struct {
	NodeId
	Pos           int
	AssignLHSElem *Node
}

/*
type IdentNode struct {
	NodeId
	Pos int
	//	Ident grammar.IDENTIFIER    check this error
	Ident string
}
*/

type ArrayElemNode struct {
	NodeId
	Pos   int
	Ident string    // *IdentNode
	Expr  *ExprNode // repeatable
}

type PairElemNode struct {
	NodeId
	Pos int
	Fst *ExprNode
	Snd *ExprNode
}

type AssignRHSNode struct {
	NodeId
	Pos           int
	AssignRHSElem *Node
}

type ExprNode struct {
	NodeId
	Pos      int
	ExprElem *Node
}

type ArrayLiterNode struct {
	NodeId
	Pos  int
	Expr []*ExprNode
}

type NewPairNode struct {
	NodeId
	Pos  int
	Expr *ExprNode // need two
}

type CallNode struct {
	NodeId
	Pos     int
	Ident   string       // *IdentNode
	ArgList *ArgListNode // optional
}

type ArgListNode struct {
	NodeId
	Pos  int
	Expr []*ExprNode // one or more args list
}

type TypeNode struct {
	NodeId
	Pos      int
	TypeElem *Node
}

type BaseTypeNode struct {
	NodeId
	Pos          int
	BaseTypeElem *Node
}

type IntNode struct {
	NodeId
	Pos int
	Int int
}

type BoolNode struct {
	NodeId
	Pos  int
	Bool bool
}

type CharNode struct {
	NodeId
	Pos  int
	Char rune
}

type StringNode struct {
	NodeId
	Pos    int
	String string
}

type ArrayTypeNode struct {
	NodeId
	Pos  int
	Type *TypeNode
}

type PairTypeNode struct {
	NodeId
	Pos          int
	PairElemType []*PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	NodeId
	Pos              int
	PairElemTypeElem *Node
}

type IntLiterNode struct {
	NodeId
	Pos     int
	IntSign *IntSignNode // optional
	Digit   *DigitNode   // repeatable
}

type BoolLiterNode struct {
	NodeId
	Pos       int
	BoolLiter bool
}

type CharLiterNode struct {
	NodeId
	Pos           int
	CharLiterElem *CharacterNode
}

type StrLiterNode struct {
	NodeId
	Pos      int
	StrLiter []*CharacterNode
}

type CharacterNode struct {
	NodeId
	Pos int
	// any ascii characterr
	CharacterElem *Node
}

type EscapedCharNode struct {
	NodeId
	Pos             int
	EscapedCharElem *Node
}

type NullNode struct {
	NodeId
	Pos            int
	NullTerminator grammar.ItemType
}

type BackslashNode struct {
	NodeId
	Pos       int
	Backslash grammar.ItemType
}

type TabNode struct {
	NodeId
	Pos int
	Tab grammar.ItemType
}

type LineFeedNode struct {
	NodeId
	Pos      int
	LineFeed grammar.ItemType
}

type FormFeedNode struct {
	NodeId
	Pos      int
	FormFeed grammar.ItemType
}

type CarriageReturnNode struct {
	NodeId
	Pos            int
	CarriageReturn grammar.ItemType
}

type DoubleQuoteNode struct {
	NodeId
	Pos         int
	DoubleQuote grammar.ItemType
}

type SingleQuoteNode struct {
	NodeId
	Pos         int
	SingleQuote grammar.ItemType
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	NodeId
	Pos       int
	PairLiter grammar.ItemType //should be NULL
}

type UnaryOperExprNode struct {
	NodeId
	Pos       int
	UnaryOper *UnaryOperNode
	Expr      *ExprNode
}

type EBENode struct {
	NodeId
	Pos        int
	Expr       *ExprNode
	BinaryOper *BinaryOperNode
	Expre      *ExprNode
}

// ( <expr> ) inside expr handled by parser

type UnaryOperNode struct {
	NodeId
	Pos           int
	UnaryOperElem *Node
}

type NotNode struct {
	NodeId
	Pos int
	Not grammar.ItemType
}

type NegNode struct {
	NodeId
	Pos int
	Neg grammar.ItemType
}

type LenNode struct {
	NodeId
	Pos int
	Len grammar.ItemType
}

type OrdNode struct {
	NodeId
	Pos int
	Ord grammar.ItemType
}

type ChrNode struct {
	NodeId
	Pos int
	Chr grammar.ItemType
}

type BinaryOperNode struct {
	NodeId
	Pos            int
	BinaryOperElem *Node
}

type MultNode struct {
	NodeId
	Pos  int
	Mult grammar.ItemType
}

type DivNode struct {
	NodeId
	Pos int
	Div grammar.ItemType
}

type ModNode struct {
	NodeId
	Pos int
	Mod grammar.ItemType
}

type AddNode struct {
	NodeId
	Pos int
	Add grammar.ItemType
}

type SubNode struct {
	NodeId
	Pos int
	Sub grammar.ItemType
}

type GtNode struct {
	NodeId
	Pos int
	Gt  grammar.ItemType
}

type GteNode struct {
	NodeId
	Pos int
	Gte grammar.ItemType
}

type LtNode struct {
	NodeId
	Pos int
	Lt  grammar.ItemType
}

type LteNode struct {
	NodeId
	Pos int
	Lte grammar.ItemType
}

type EqNode struct {
	NodeId
	Pos int
	Eq  grammar.ItemType
}

type NeqNode struct {
	NodeId
	Pos int
	Neq grammar.ItemType
}

type AndNode struct {
	NodeId
	Pos int
	And grammar.ItemType
}

type OrNode struct {
	NodeId
	Pos int
	Or  grammar.ItemType
}

type DigitNode struct {
	NodeId
	Pos int
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
	Digit int
}

type IntSignNode struct {
	NodeId
	Pos         int
	IntSignElem *Node
}

type PlusNode struct {
	NodeId
	Pos int
	Add grammar.ItemType
}

type NegateNode struct {
	NodeId
	Pos int
	Sub grammar.ItemType
}

//no need for comment node
