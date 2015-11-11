package abstractSyntaxTree

import (
	grammar "github.com/henrykhadass/wacc_19/src/grammar"
)

// A node is an element in te abstract syntax tree
type Node interface {
	Type() NodeType
	String() string
	Position() Pos
}

// NodeType identifies the type of an abstract syntax tree node
type NodeType int

// Pos represents a byte position in the original input text from which
// this template was parsed
type Pos int

func (p Pos) Position() Pos {
	return p
}

// Type returns itself and provides an easy default implementation
// for embedding in a Node. Embedded in all non-trivial Nodes
func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeSUCKADICk NodeType = iota
	NodeProgram
	NodeFunc
)

type ProgramNode struct {
	NodeType
	Pos
	Func []*FuncNode
	Stat *StatNode
}

type FuncNode struct {
	NodeType
	Pos
	Type      *TypeNode
	Ident     *IdentNode
	ParamList *ParamListNode // Optional
	Stat      *StatNode
}

type ParamListNode struct {
	NodeType
	Pos
	Param *ParamNode
	Param []*ParamNode
}

type ParamNode struct {
	NodeType
	Pos
	Type  *TypeNode
	Ident *IdentNode
}

type StatNode struct {
	NodeType
	Pos
	StatElem *Node
}

type TypeIdentNode struct {
	NodeType
	Pos
	Type      *TypeNode
	Ident     *IdentNode
	AssignRHS *AssignRHSNode
}

type AssignLHSToRightNode struct {
	NodeType
	Pos
	AssignLHS *AssignLHSNode
	AssignRHS *AssignRHSNode
}

type ReadNode struct {
	NodeType
	Pos
	AssignLHS *AssignLHSNode
}

type FreeNode struct {
	NodeType
	Pos
	Expr *ExprNode
}

type ReturnNode struct {
	NodeType
	Pos
	Expr *ExprNode
}

type ExitNode struct {
	NodeType
	Pos
	Expr *ExprNode
}

type PrintNode struct {
	NodeType
	Pos
	Expr *ExprNode
}

type PrintlnNode struct {
	NodeType
	Pos
	Expr *ExprNode
}

type IfNode struct {
	NodeType
	Pos
	Expr *ExprNode
	Stat *StatNode
	Stat *StatNode
}

type WhileNode struct {
	NodeType
	Pos
	Expr *ExprNode
	Stat *StatNode
}

//    ‘begin’ <stat> ‘end’
//     <stat> ‘;’ <stat>
type AssignLHSNode struct {
	NodeType
	Pos
	AssignLHSElem *Node
}

type IdentNode struct {
	NodeType
	Pos
	//	Ident grammar.IDENTIFIER    check this error
}

type ArrayElemNode struct {
	NodeType
	Pos
	Ident *IdentNode
	Expr  *ExprNode // repeatable
}

type PairElemNode struct {
	NodeType
	Pos
	Fst *ExprNode
	Snd *ExprNode
}

type AssignRHSNode struct {
	NodeType
	Pos
	AssignRHSElem *Node
}

type ExprNode struct {
	NodeType
	Pos
	ExprElem *Node
}

type ArrayLiterNode struct {
	NodeType
	Pos
	Expr *ExprNode
	Expr []*ExprNode
}

type NewPairNode struct {
	NodeType
	Pos
	Expr *ExprNode
	Expr *ExprNode
}

type CallNode struct {
	NodeType
	Pos
	Ident   *IdentNode
	ArgList *ArgListNode // optional
}

type ArgListNode struct {
	NodeType
	Pos
	Expr *ExprNode
	Expr []*ExprNode
}

type TypeNode struct {
	NodeType
	Pos
	TypeElem *Node
}

type BaseTypeNode struct {
	NodeType
	Pos
	BaseTypeElem *Node
}

type IntNode struct {
	NodeType
	Pos
	Int int
}

type BoolNode struct {
	NodeType
	Pos
	Bool bool
}

type CharNode struct {
	NodeType
	Pos
	Char rune
}

type StringNode struct {
	NodeType
	Pos
	String string
}

type ArrayTypeNode struct {
	NodeType
	Pos
	Type *TypeNode
}

type PairTypeNode struct {
	NodeType
	Pos
	PairElemType *PairElemTypeNode
	PairElemType *PairElemTypeNode
}

type PairElemTypeNode struct {
	NodeType
	Pos
	PairElemTypeElem *Node
}

type IntLiterNode struct {
	NodeType
	Pos
	IntSign *IntSignNode // optional
	Digit   *DigitNode   // repeatable
}

type BoolLiterNode struct {
	NodeType
	Pos
	BoolLiter bool
}

type CharLiterNode struct {
	NodeType
	Pos
	CharLiterElem *CharacterNode
}

type StrLiterNode struct {
	NodeType
	Pos
	StrLiter []*CharacterNode
}

type CharacterNode struct {
	NodeType
	Pos
	// any ascii characterr
	CharacterElem *Node ///WTF IS they type???  call escaped char here
}

type EscapedCharNode struct {
	NodeType
	Pos
	EscapedCharElem *Node
}

type NullNode struct {
	NodeType
	Pos
	NullElem grammar.NULL_TERMINATOR
}

type BackslashNode struct {
	NodeType
	Pos
	BackslashElem grammar.BACKSLASH
}

type TabNode struct {
	NodeType
	Pos
	TabElem grammar.TAB
}

type LineFeedNode struct {
	NodeType
	Pos
	LineFeedElem grammar.LINE_FEED
}

type FormFeedNode struct {
	NodeType
	Pos
	FormFeedElem grammar.FORM_FEED
}

type CarriageReturnNode struct {
	NodeType
	Pos
	ReturnElem grammar.CARRIAGE_RETURN
}

type DoubleQuoteNode struct {
	NodeType
	Pos
	DoubleQuoteElem grammar.DOUBLE_QUOTE
}

type SingleQuoteNode struct {
	NodeType
	Pos
	SingleQuoteElem grammar.SINGLE_QUOTE
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	NodeType
	Pos
	PairLiter nil
}

type UnaryOperExprNode struct {
	NodeType
	Pos
	UnaryOper *UnaryOperNode
	Expr      *ExprNode
}

type EBENode struct {
	NodeType
	Pos
	Expr       *ExpreNode
	BinaryOper *BinaryOperNode
	Expre      *ExprNode
}

// ( <expr> ) inside expr how to deal with that

type UnaryOperNode struct {
	NodeType
	Pos
	UnaryOperElem *Node
}

type NotNode struct {
	NodeType
	Pos
	NotElem grammar.NOT
}

type NegNode struct {
	NodeType
	Pos
	NegElem grammar.NEG
}

type LenNode struct {
	NodeType
	Pos
	LenElem grammar.LEN
}

type OrdNode struct {
	NodeType
	Pos
	OrdElem grammar.ORD
}

type ChrNode struct {
	NodeType
	Pos
	ChrElem grammar.CHR
}

type BinaryOperNode struct {
	NodeType
	Pos
	BinaryOperElem *Node
}

type MultNode struct {
	NodeType
	Pos
	MultElem grammar.MULT
}

type DivNode struct {
	NodeType
	Pos
	DivElem grammar.DIV
}

type ModNode struct {
	NodeType
	Pos
	ModElem grammar.MOD
}

type AddNode struct {
	NodeType
	Pos
	AddElem grammar.ADD
}

type SubNode struct {
	NodeType
	Pos
	SubElem grammar.SUB
}

type GtNode struct {
	NodeType
	Pos
	GtElem grammar.GT
}

type GteNode struct {
	NodeType
	Pos
	GteElem grammar.GTE
}

type LtNode struct {
	NodeType
	Pos
	LtElem grammar.LT
}

type LteNode struct {
	NodeType
	Pos
	LteElem grammar.LTE
}

type EqNode struct {
	NodeType
	Pos
	EqElem grammar.EQ
}

type NeqNode struct {
	NodeType
	Pos
	NeqElem grammar.NEQ
}

type AndNode struct {
	NodeType
	Pos
	AndElem grammar.AND
}

type OrNode struct {
	NodeType
	Pos
	OrElem grammar.OR
}

type DigitNode struct {
	NodeType
	Pos
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
}

type IntSignNode struct {
	NodeType
	Pos
	IntSignElem *Node
}

type PlusNode struct {
	NodeType
	Pos
	PlusElem grammar.ADD
}

type NegateNode struct {
	NodeType
	Pos
	NegateElem grammar.SUB
}

type CommentNode struct {
	NodeType
	Pos
	//   chars //????
	//   EOL //eol type?????
}
