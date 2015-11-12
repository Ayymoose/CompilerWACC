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
	NodeId
	Pos  grammar.Position
	Func []*FuncNode
	Stat []*StatNode
}

type FuncNode struct {
	NodeId
	Pos       grammar.Position
	Type      *TypeNode
	Ident     *IdentNode
	ParamList *ParamListNode // Optional
	Stat      *StatNode
}

type ParamListNode struct {
	NodeId
	Pos   grammar.Position
	Param []*ParamNode // one or more
}

type ParamNode struct {
	NodeId
	Pos   grammar.Position
	Type  *TypeNode
	Ident *IdentNode
}

type StatNode struct {
	NodeId
	Pos      grammar.Position
	StatElem *Node
}

// TypeIdent
type DeclarationNode struct {
	NodeId
	Pos       grammar.Position
	Type      *TypeNode
	Ident     *IdentNode
	AssignRHS *AssignRHSNode
}

// AssignLHSToRight
type AssignmentNode struct {
	NodeId
	Pos       grammar.Position
	AssignLHS *AssignLHSNode
	AssignRHS *AssignRHSNode
}

type ReadNode struct {
	NodeId
	Pos       grammar.Position
	AssignLHS *AssignLHSNode
}

type FreeNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
}

type ReturnNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
}

type ExitNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
}

type PrintNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
}

type PrintlnNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
}

type IfNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode
	Stat []*StatNode // two stats
}

type WhileNode struct {
	NodeId
	Pos  grammar.Position
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
	Pos           grammar.Position
	AssignLHSElem *Node
}

type IdentNode struct {
	NodeId
	Pos grammar.Position
	//	Ident grammar.IDENTIFIER    check this error
	Ident string
}

type ArrayElemNode struct {
	NodeId
	Pos   grammar.Position
	Ident *IdentNode
	Expr  *ExprNode // repeatable
}

type PairElemNode struct {
	NodeId
	Pos grammar.Position
	Fst *ExprNode
	Snd *ExprNode
}

type AssignRHSNode struct {
	NodeId
	Pos           grammar.Position
	AssignRHSElem *Node
}

type ExprNode struct {
	NodeId
	Pos      grammar.Position
	ExprElem *Node
}

type ArrayLiterNode struct {
	NodeId
	Pos  grammar.Position
	Expr []*ExprNode
}

type NewPairNode struct {
	NodeId
	Pos  grammar.Position
	Expr *ExprNode // need two
}

type CallNode struct {
	NodeId
	Pos     grammar.Position
	Ident   *IdentNode
	ArgList *ArgListNode // optional
}

type ArgListNode struct {
	NodeId
	Pos  grammar.Position
	Expr []*ExprNode // one or more args list
}

type TypeNode struct {
	NodeId
	Pos      grammar.Position
	TypeElem *Node
}

type BaseTypeNode struct {
	NodeId
	Pos          grammar.Position
	BaseTypeElem *Node
}

type IntNode struct {
	NodeId
	Pos grammar.Position
	Int int
}

type BoolNode struct {
	NodeId
	Pos  grammar.Position
	Bool bool
}

type CharNode struct {
	NodeId
	Pos  grammar.Position
	Char rune
}

type StringNode struct {
	NodeId
	Pos    grammar.Position
	String string
}

type ArrayTypeNode struct {
	NodeId
	Pos  grammar.Position
	Type *TypeNode
}

type PairTypeNode struct {
	NodeId
	Pos          grammar.Position
	PairElemType []*PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	NodeId
	Pos              grammar.Position
	PairElemTypeElem *Node
}

type IntLiterNode struct {
	NodeId
	Pos     grammar.Position
	IntSign *IntSignNode // optional
	Digit   *DigitNode   // repeatable
}

type BoolLiterNode struct {
	NodeId
	Pos       grammar.Position
	BoolLiter bool
}

type CharLiterNode struct {
	NodeId
	Pos           grammar.Position
	CharLiterElem *CharacterNode
}

type StrLiterNode struct {
	NodeId
	Pos      grammar.Position
	StrLiter []*CharacterNode
}

type CharacterNode struct {
	NodeId
	Pos grammar.Position
	// any ascii characterr
	CharacterElem *Node
}

type EscapedCharNode struct {
	NodeId
	Pos             grammar.Position
	EscapedCharElem *Node
}

type NullNode struct {
	NodeId
	Pos            grammar.Position
	NullTerminator grammar.ItemType
}

type BackslashNode struct {
	NodeId
	Pos       grammar.Position
	Backslash grammar.ItemType
}

type TabNode struct {
	NodeId
	Pos grammar.Position
	Tab grammar.ItemType
}

type LineFeedNode struct {
	NodeId
	Pos      grammar.Position
	LineFeed grammar.ItemType
}

type FormFeedNode struct {
	NodeId
	Pos      grammar.Position
	FormFeed grammar.ItemType
}

type CarriageReturnNode struct {
	NodeId
	Pos            grammar.Position
	CarriageReturn grammar.ItemType
}

type DoubleQuoteNode struct {
	NodeId
	Pos         grammar.Position
	DoubleQuote grammar.ItemType
}

type SingleQuoteNode struct {
	NodeId
	Pos         grammar.Position
	SingleQuote grammar.ItemType
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	NodeId
	Pos       grammar.Position
	PairLiter grammar.ItemType //should be NULL
}

type UnaryOperExprNode struct {
	NodeId
	Pos       grammar.Position
	UnaryOper *UnaryOperNode
	Expr      *ExprNode
}

type EBENode struct {
	NodeId
	Pos        grammar.Position
	Expr       *ExprNode
	BinaryOper *BinaryOperNode
	Expre      *ExprNode
}

// ( <expr> ) inside expr handled by parser

type UnaryOperNode struct {
	NodeId
	Pos           grammar.Position
	UnaryOperElem *Node
}

type NotNode struct {
	NodeId
	Pos grammar.Position
	Not grammar.ItemType
}

type NegNode struct {
	NodeId
	Pos grammar.Position
	Neg grammar.ItemType
}

type LenNode struct {
	NodeId
	Pos grammar.Position
	Len grammar.ItemType
}

type OrdNode struct {
	NodeId
	Pos grammar.Position
	Ord grammar.ItemType
}

type ChrNode struct {
	NodeId
	Pos grammar.Position
	Chr grammar.ItemType
}

type BinaryOperNode struct {
	NodeId
	Pos            grammar.Position
	BinaryOperElem *Node
}

type MultNode struct {
	NodeId
	Pos  grammar.Position
	Mult grammar.ItemType
}

type DivNode struct {
	NodeId
	Pos grammar.Position
	Div grammar.ItemType
}

type ModNode struct {
	NodeId
	Pos grammar.Position
	Mod grammar.ItemType
}

type AddNode struct {
	NodeId
	Pos grammar.Position
	Add grammar.ItemType
}

type SubNode struct {
	NodeId
	Pos grammar.Position
	Sub grammar.ItemType
}

type GtNode struct {
	NodeId
	Pos grammar.Position
	Gt  grammar.ItemType
}

type GteNode struct {
	NodeId
	Pos grammar.Position
	Gte grammar.ItemType
}

type LtNode struct {
	NodeId
	Pos grammar.Position
	Lt  grammar.ItemType
}

type LteNode struct {
	NodeId
	Pos grammar.Position
	Lte grammar.ItemType
}

type EqNode struct {
	NodeId
	Pos grammar.Position
	Eq  grammar.ItemType
}

type NeqNode struct {
	NodeId
	Pos grammar.Position
	Neq grammar.ItemType
}

type AndNode struct {
	NodeId
	Pos grammar.Position
	And grammar.ItemType
}

type OrNode struct {
	NodeId
	Pos grammar.Position
	Or  grammar.ItemType
}

type DigitNode struct {
	NodeId
	Pos grammar.Position
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
	Digit int
}

type IntSignNode struct {
	NodeId
	Pos         grammar.Position
	IntSignElem *Node
}

type PlusNode struct {
	NodeId
	Pos grammar.Position
	Add grammar.ItemType
}

type NegateNode struct {
	NodeId
	Pos grammar.Position
	Sub grammar.ItemType
}

//no need for comment node
