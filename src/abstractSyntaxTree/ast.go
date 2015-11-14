package abstractSyntaxTree

import (
	grammar "github.com/henrykhadass/wacc_19/src/grammar"
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
	getNodeType()
}

// Root node of AST
type ProgramNode struct {
	Node
	Pos  grammar.Position
	Func []*FuncNode
	Stat []*StatNode
}

type FuncNode struct {
	Node
	Pos       grammar.Position
	Type      grammar.ItemType // add the type
	Ident     string           // changed from *IdentNode
	ParamList []*ParamNode     // *ParamListNode // Optional
	StatList  []*StatNode
}

/*
type ParamListNode struct {
	Node
	Pos   grammar.Position
	Param []*ParamNode // one or more
}
*/

type ParamNode struct {
	Node
	Pos   grammar.Position
	Type  grammar.ItemType //string
	Ident string           // changed from  *IdentNode
}

type StatNode struct {
	Node
	Pos      grammar.Position
	StatElem *Node
}

// TypeIdent
type DeclarationNode struct {
	Node
	Pos       grammar.Position
	Type      grammar.ItemType //string // changed from *TypeNode
	Ident     string           // changed from *IdentNode
	AssignRHS *AssignRHSNode
}

// AssignLHSToRight
type AssignmentNode struct {
	Node
	Pos       grammar.Position
	AssignLHS *AssignLHSNode
	AssignRHS *AssignRHSNode
}

type ReadNode struct {
	Node
	Pos       grammar.Position
	AssignLHS *AssignLHSNode
}

type FreeNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
}

type ReturnNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
}

type ExitNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
}

type PrintNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
}

type PrintlnNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
}

type IfNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
	Stat []*StatNode // two stats
}

type WhileNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode
	Stat *StatNode
}

//    ‘begin’ <stat> ‘end’
// SPECIAL NEEDS
type ScopeNode struct {
	Node
	StatElem *StatNode
}

type AssignLHSNode struct {
	Node
	Pos           grammar.Position
	AssignLHSElem *Node
}

/*
type IdentNode struct {
	Node
	Pos grammar.Position
	//	Ident grammar.IDENTIFIER    check this error
	Ident string
}
*/

type ArrayElemNode struct {
	Node
	Pos   grammar.Position
	Ident string    // *IdentNode
	Expr  *ExprNode // repeatable
}

type PairElemNode struct {
	Node
	Pos grammar.Position
	Fst *ExprNode
	Snd *ExprNode
}

type AssignRHSNode struct {
	Node
	Pos           grammar.Position
	AssignRHSElem *Node
}

type ExprNode struct {
	Node
	Pos      grammar.Position
	ExprElem *Node
}

type ArrayLiterNode struct {
	Node
	Pos  grammar.Position
	Expr []*ExprNode
}

type NewPairNode struct {
	Node
	Pos  grammar.Position
	Expr *ExprNode // need two
}

type CallNode struct {
	Node
	Pos     grammar.Position
	Ident   string       // *IdentNode
	ArgList *ArgListNode // optional
}

type ArgListNode struct {
	Node
	Pos  grammar.Position
	Expr []*ExprNode // one or more args list
}

type TypeNode struct {
	Node
	Pos      grammar.Position
	TypeElem *Node
}

type BaseTypeNode struct {
	Node
	Pos          grammar.Position
	BaseTypeElem grammar.ItemType //*Node
}

/*
type IntNode struct {
	Node
	Pos grammar.Position
	Int int
}

type BoolNode struct {
	Node
	Pos  grammar.Position
	Bool bool
}

type CharNode struct {
	Node
	Pos  grammar.Position
	Char rune
}

type StringNode struct {
	Node
	Pos    grammar.Position
	String string
}
*/
type ArrayTypeNode struct {
	Node
	Pos  grammar.Position
	Type *TypeNode
}

type PairTypeNode struct {
	Node
	Pos          grammar.Position
	PairElemType []*PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	Node
	Pos              grammar.Position
	PairElemTypeElem *Node
}

type IntLiterNode struct {
	Node
	Pos     grammar.Position
	IntSign *IntSignNode // optional
	Digit   *DigitNode   // repeatable
}

type BoolLiterNode struct {
	Node
	Pos       grammar.Position
	BoolLiter bool
}

type CharLiterNode struct {
	Node
	Pos           grammar.Position
	CharLiterElem *CharacterNode
}

type StrLiterNode struct {
	Node
	Pos      grammar.Position
	StrLiter []*CharacterNode
}

type CharacterNode struct {
	Node
	Pos grammar.Position
	// any ascii characterr
	CharacterElem *Node
}

type EscapedCharNode struct {
	Node
	Pos             grammar.Position
	EscapedCharElem *Node
}

type NullNode struct {
	Node
	Pos            grammar.Position
	NullTerminator grammar.ItemType
}

type BackslashNode struct {
	Node
	Pos       grammar.Position
	Backslash grammar.ItemType
}

type TabNode struct {
	Node
	Pos grammar.Position
	Tab grammar.ItemType
}

type LineFeedNode struct {
	Node
	Pos      grammar.Position
	LineFeed grammar.ItemType
}

type FormFeedNode struct {
	Node
	Pos      grammar.Position
	FormFeed grammar.ItemType
}

type CarriageReturnNode struct {
	Node
	Pos            grammar.Position
	CarriageReturn grammar.ItemType
}

type DoubleQuoteNode struct {
	Node
	Pos         grammar.Position
	DoubleQuote grammar.ItemType
}

type SingleQuoteNode struct {
	Node
	Pos         grammar.Position
	SingleQuote grammar.ItemType
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	Node
	Pos       grammar.Position
	PairLiter grammar.ItemType //should be NULL
}

type UnaryOperExprNode struct {
	Node
	Pos       grammar.Position
	UnaryOper *UnaryOperNode
	Expr      *ExprNode
}

type EBENode struct {
	Node
	Pos        grammar.Position
	Expr       *ExprNode
	BinaryOper *BinaryOperNode
	Expre      *ExprNode
}

// ( <expr> ) inside expr handled by parser

type UnaryOperNode struct {
	Node
	Pos           grammar.Position
	UnaryOperElem *Node
}

type NotNode struct {
	Node
	Pos grammar.Position
	Not grammar.ItemType
}

type NegNode struct {
	Node
	Pos grammar.Position
	Neg grammar.ItemType
}

type LenNode struct {
	Node
	Pos grammar.Position
	Len grammar.ItemType
}

type OrdNode struct {
	Node
	Pos grammar.Position
	Ord grammar.ItemType
}

type ChrNode struct {
	Node
	Pos grammar.Position
	Chr grammar.ItemType
}

type BinaryOperNode struct {
	Node
	Pos            grammar.Position
	BinaryOperElem *Node
}

type MultNode struct {
	Node
	Pos  grammar.Position
	Mult grammar.ItemType
}

type DivNode struct {
	Node
	Pos grammar.Position
	Div grammar.ItemType
}

type ModNode struct {
	Node
	Pos grammar.Position
	Mod grammar.ItemType
}

type AddNode struct {
	Node
	Pos grammar.Position
	Add grammar.ItemType
}

type SubNode struct {
	Node
	Pos grammar.Position
	Sub grammar.ItemType
}

type GtNode struct {
	Node
	Pos grammar.Position
	Gt  grammar.ItemType
}

type GteNode struct {
	Node
	Pos grammar.Position
	Gte grammar.ItemType
}

type LtNode struct {
	Node
	Pos grammar.Position
	Lt  grammar.ItemType
}

type LteNode struct {
	Node
	Pos grammar.Position
	Lte grammar.ItemType
}

type EqNode struct {
	Node
	Pos grammar.Position
	Eq  grammar.ItemType
}

type NeqNode struct {
	Node
	Pos grammar.Position
	Neq grammar.ItemType
}

type AndNode struct {
	Node
	Pos grammar.Position
	And grammar.ItemType
}

type OrNode struct {
	Node
	Pos grammar.Position
	Or  grammar.ItemType
}

type DigitNode struct {
	Node
	Pos grammar.Position
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
	Digit int
}

type IntSignNode struct {
	Node
	Pos         grammar.Position
	IntSignElem *Node
}

type PlusNode struct {
	Node
	Pos grammar.Position
	Add grammar.ItemType
}

type NegateNode struct {
	Node
	Pos grammar.Position
	Sub grammar.ItemType
}

//no need for comment node
