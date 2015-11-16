package abstractSyntaxTree

import (
	"grammar"
)

// Node interface
// Contains methods to be called on ondividual structs
type Node interface {
	BuildNode() Node
}

type BuildArguments struct {
	Pos          grammar.Position
	Type         grammar.ItemType
	Ident        string
	ChildListOne []Node
	ChildListTwo []Node
}

// Structs for common fields
// Not all structs have all of the common fields
// Hence multiple field structs
type Position struct {
	Pos grammar.Position
}

type PosIdent struct {
	Pos   grammar.Position
	Ident string
}

type PosIdentType struct {
	Pos   grammar.Position
	Ident string
	Type  grammar.ItemType
}

// Root node of AST
type ProgramNode struct {
	Position
	Node
	Func []FunctionNode //rename funcslice
	Stat []StatementNode
}

type FunctionNode struct {
	PosIdentType
	Node
	ParamList []ParameterNode // *ParamListNode // Optional
	StatList  []StatementNode
}

type ParameterNode struct {
	PosIdentType
	Node
}

// TypeIdent
type DeclarationNode struct {
	PosIdentType
	Node
	AssignRHS AssignRHSNode
}

type ArrayElemNode struct {
	PosIdent
	Node
	Expr ExprNode // repeatable
}

type CallNode struct {
	PosIdent
	Node
	ArgList ArgListNode // optional
}

/*
type ParamListNode struct {
	Node
	Pos   grammar.Position
	Param []*ParamNode // one or more
}
*/

type StatementNode struct {
	Position
	Node
	StatElem Node
}

// AssignLHSToRight
type AssignmentNode struct {
	Position
	Node
	AssignLHS AssignLHSNode
	AssignRHS AssignRHSNode
}

type ReadNode struct {
	Position
	Node
	AssignLHS AssignLHSNode
}

type FreeNode struct {
	Position
	Node
	Expr ExprNode
}

type ReturnNode struct {
	Position
	Node
	Expr ExprNode
}

type ExitNode struct {
	Position
	Node
	Expr ExprNode
}

type PrintNode struct {
	Position
	Node
	Expr ExprNode
}

type PrintlnNode struct {
	Position
	Node
	Expr ExprNode
}

type IfNode struct {
	Position
	Node
	Expr ExprNode
	Stat []StatementNode // two stats
}

type WhileNode struct {
	Position
	Node
	Expr ExprNode
	Stat StatementNode
}

//    ‘begin’ <stat> ‘end’
// SPECIAL NEEDS
type ScopeNode struct {
	Position
	Node
	StatElem StatementNode
}

type AssignLHSNode struct {
	Position
	Node
	AssignLHSElem Node
}

/*
type IdentNode struct {
	Node
	Pos grammar.Position
	//	Ident grammar.IDENTIFIER    check this error
	Ident string
}
*/

type PairElemNode struct {
	Position
	Node
	Fst ExprNode
	Snd ExprNode
}

type AssignRHSNode struct {
	Position
	Node
	AssignRHSElem Node
}

type ExprNode struct {
	Position
	Node
	ExprElem Node
}

type ArrayLiterNode struct {
	Position
	Node
	Expr []ExprNode
}

type NewPairNode struct {
	Position
	Node
	Expr ExprNode // need two
}

type ArgListNode struct {
	Position
	Node
	Expr []ExprNode // one or more args list
}

type TypeNode struct {
	Position
	Node
	TypeElem Node
}

type BaseTypeNode struct {
	Position
	Node
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
	Position
	Node
	Type TypeNode
}

type PairTypeNode struct {
	Position
	Node
	PairElemType []PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	Position
	Node
	PairElemTypeElem Node
}

type IntLiterNode struct {
	Position
	Node
	IntSign IntSignNode // optional
	Digit   DigitNode   // repeatable
}

type BoolLiterNode struct {
	Position
	Node
	BoolLiter bool
}

type CharLiterNode struct {
	Position
	Node
	CharLiterElem CharacterNode
}

type StringLiterNode struct {
	Position
	Node
	StrLiter []CharacterNode
}

type CharacterNode struct {
	Position
	Node
	// any ascii characterr
	CharacterElem Node
}

type EscapedCharNode struct {
	Position
	Node
	EscapedCharElem Node
}

type NullNode struct {
	Position
	Node
	NullTerminator grammar.ItemType
}

type BackslashNode struct {
	Position
	Node
	Backslash grammar.ItemType
}

type TabNode struct {
	Position
	Node
	Tab grammar.ItemType
}

type LineFeedNode struct {
	Position
	Node
	LineFeed grammar.ItemType
}

type FormFeedNode struct {
	Position
	Node
	FormFeed grammar.ItemType
}

type CarriageReturnNode struct {
	Position
	Node
	CarriageReturn grammar.ItemType
}

type DoubleQuoteNode struct {
	Position
	Node
	DoubleQuote grammar.ItemType
}

type SingleQuoteNode struct {
	Position
	Node
	SingleQuote grammar.ItemType
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	Position
	Node
	PairLiter grammar.ItemType //should be NULL
}

type UnaryOpExprNode struct {
	Position
	Node
	UnaryOper UnaryOpNode
	Expr      ExprNode
}

type EBENode struct {
	Position
	Node
	Expr       []ExprNode
	BinaryOper BinaryOpNode
	// one either side of expr
}

// ( <expr> ) inside expr handled by parser

type UnaryOpNode struct {
	Position
	Node
	UnaryOperElem Node
}

type NotNode struct {
	Position
	Node
	Not grammar.ItemType
}

type NegNode struct {
	Position
	Node
	Neg grammar.ItemType
}

type LenNode struct {
	Position
	Node
	Len grammar.ItemType
}

type OrdNode struct {
	Position
	Node
	Ord grammar.ItemType
}

type ChrNode struct {
	Position
	Node
	Chr grammar.ItemType
}

type BinaryOpNode struct {
	Position
	Node
	BinaryOperElem *Node
}

type MultNode struct {
	Position
	Node
	Mult grammar.ItemType
}

type DivNode struct {
	Position
	Node
	Div grammar.ItemType
}

type ModNode struct {
	Position
	Node
	Mod grammar.ItemType
}

type AddNode struct {
	Position
	Node
	Add grammar.ItemType
}

type SubNode struct {
	Position
	Node
	Sub grammar.ItemType
}

type GTNode struct {
	Position
	Node
	Gt grammar.ItemType
}

type GTENode struct {
	Position
	Node
	Gte grammar.ItemType
}

type LTNode struct {
	Position
	Node
	Lt grammar.ItemType
}

type LTENode struct {
	Position
	Node
	Lte grammar.ItemType
}

type EQNode struct {
	Position
	Node
	Eq grammar.ItemType
}

type NEQNode struct {
	Position
	Node
	Neq grammar.ItemType
}

type AndNode struct {
	Position
	Node
	And grammar.ItemType
}

type OrNode struct {
	Position
	Node
	Or grammar.ItemType
}

type DigitNode struct {
	Position
	Node
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
	Digit int
}

type IntSignNode struct {
	Position
	Node
	IntSignElem Node
}

type PositiveNode struct {
	Position
	Node
	Add grammar.ItemType
}

type NegativeNode struct {
	Position
	Node
	Sub grammar.ItemType
}

//no need for comment node

func (p *ProgramNode) BuildNode() Node {
	p.Pos = BuildArguments.Pos
	etcccccc
	return p
}

func (p *ProgramNode) BuildNode() Node {
	return p
}

func (f *FunctionNode) BuildNode() Node {
	return f
}

func (p *ParameterNode) BuildNode() Node {
	return p
}

func (s *StatementNode) BuildNode() Node {
	return s
}

func (d *DeclarationNode) BuildNode() Node {
	return d
}

func (a *AssignmentNode) BuildNode() Node {
	return a
}

func (r *ReadNode) BuildNode() Node {
	return r
}

func (f *FreeNode) BuildNode() Node {
	return f
}

func (r *ReturnNode) BuildNode() Node {
	return r
}

func (e *ExitNode) BuildNode() Node {
	return e
}

func (p *PrintNode) BuildNode() Node {
	return p
}

func (p *PrintlnNode) BuildNode() Node {
	return p
}

func (i *IfNode) BuildNode() Node {
	return i
}

func (w *WhileNode) BuildNode() Node {
	return w
}

func (s *ScopeNode) BuildNode() Node {
	return s
}

func (a *AssignLHSNode) BuildNode() Node {
	return a
}

func (a *ArrayElemNode) BuildNode() Node {
	return a
}

func (p *PairElemNode) BuildNode() Node {
	return p
}

func (a *AssignRHSNode) BuildNode() Node {
	return a
}

func (e *ExprNode) BuildNode() Node {
	return e
}

func (a *ArrayLiterNode) BuildNode() Node {
	return a
}

func (n *NewPairNode) BuildNode() Node {
	return n
}

func (c *CallNode) BuildNode() Node {
	return c
}

func (a *ArgListNode) BuildNode() Node {
	return a
}

func (t *TypeNode) BuildNode() Node {
	return t
}

func (b *BaseTypeNode) BuildNode() Node {
	return b
}

func (a *ArrayTypeNode) BuildNode() Node {
	return a
}

func (p *PairTypeNode) BuildNode() Node {
	return p
}

func (p *PairElemTypeNode) BuildNode() Node {
	return p
}

func (i *IntLiterNode) BuildNode() Node {
	return i
}

func (b *BoolLiterNode) BuildNode() Node {
	return b
}

func (c *CharLiterNode) BuildNode() Node {
	return c
}

func (s *StringLiterNode) BuildNode() Node {
	return s
}
func (c *CharacterNode) BuildNode() Node {
	return c
}

func (e *EscapedCharNode) BuildNode() Node {
	return e
}

func (n *NullNode) BuildNode() Node {
	return n
}

func (b *BackslashNode) BuildNode() Node {
	return b
}

func (t *TabNode) BuildNode() Node {
	return t
}

func (l *LineFeedNode) BuildNode() Node {
	return l
}

func (f *FormFeedNode) BuildNode() Node {
	return f
}

func (c *CarriageReturnNode) BuildNode() Node {
	return c
}

func (d *DoubleQuoteNode) BuildNode() Node {
	return d
}

func (s *SingleQuoteNode) BuildNode() Node {
	return s
}

func (p *PairLiterNode) BuildNode() Node {
	return p
}

func (u *UnaryOpExprNode) BuildNode() Node {
	return u
}

func (e *EBENode) BuildNode() Node {
	return e
}

func (u *UnaryOpNode) BuildNode() Node {
	return u
}

func (n *NotNode) BuildNode() Node {
	return n
}

func (n *NegNode) BuildNode() Node {
	return n
}

func (l *LenNode) BuildNode() Node {
	return l
}

func (o *OrdNode) BuildNode() Node {
	return o
}

func (c *ChrNode) BuildNode() Node {
	return c
}

func (b *BinaryOpNode) BuildNode() Node {
	return b
}

func (m *MultNode) BuildNode() Node {
	return m
}

func (d *DivNode) BuildNode() Node {
	return d
}

func (m *ModNode) BuildNode() Node {
	return m
}

func (a *AddNode) BuildNode() Node {
	return a
}

func (s *SubNode) BuildNode() Node {
	return s
}

func (g *GTNode) BuildNode() Node {
	return g
}

func (g *GTENode) BuildNode() Node {
	return g
}

func (l *LTNode) BuildNode() Node {
	return l
}

func (l *LTENode) BuildNode() Node {
	return l
}

func (e *EQNode) BuildNode() Node {
	return e
}

func (n *NEQNode) BuildNode() Node {
	return n
}

func (a *AndNode) BuildNode() Node {
	return a
}

func (o *OrNode) BuildNode() Node {
	return o
}

func (d *DigitNode) BuildNode() Node {
	return d
}

func (i *IntSignNode) BuildNode() Node {
	return i
}

func (p *PositiveNode) BuildNode() Node {
	return p
}

func (n *NegativeNode) BuildNode() Node {
	return n
}
