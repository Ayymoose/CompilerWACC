package abstractSyntaxTree

import (
	grammar "github.com/henrykhadass/wacc_19/src/grammar"
)

// Node interface
type Node interface {
	BuildNode(buildArguments BuildArguments) Node
}

type BuildArguments struct {
	Pos          int
	Type         grammar.ItemType
	Ident        string
	ChildListOne []Node
	ChildListTwo []Node
}

// Structs for common fields
// Not all structs have all of the common fields
// Hence multiple field structs
type Position struct {
	Pos int
}

type PosIdent struct {
	Pos   int
	Ident string
}

type PosIdentType struct {
	Pos   int
	Ident string
	Type  grammar.ItemType
}

// Root node of AST
type ProgramNode struct {
	Position
	Func []FunctionNode //rename funcslice
	Stat []StatementNode
}

type FunctionNode struct {
	PosIdentType
	ParamList []ParameterNode // *ParamListNode // Optional
	StatList  []StatementNode
}

type ParameterNode struct {
	PosIdentType
}

// TypeIdent
type DeclarationNode struct {
	PosIdentType
	AssignRHS AssignRHSNode
}

type ArrayElemNode struct {
	PosIdent
	Expr ExprNode // repeatable
}

type CallNode struct {
	PosIdent
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
	StatElem Node
}

// AssignLHSToRight
type AssignmentNode struct {
	Position
	AssignLHS AssignLHSNode
	AssignRHS AssignRHSNode
}

type ReadNode struct {
	Position
	AssignLHS AssignLHSNode
}

type FreeNode struct {
	Position
	Expr ExprNode
}

type ReturnNode struct {
	Position
	Expr ExprNode
}

type ExitNode struct {
	Position
	Expr ExprNode
}

type PrintNode struct {
	Position
	Expr ExprNode
}

type PrintlnNode struct {
	Position
	Expr ExprNode
}

type IfNode struct {
	Position
	Expr ExprNode
	Stat []StatementNode // two stats
}

type WhileNode struct {
	Position
	Expr ExprNode
	Stat StatementNode
}

//    ‘begin’ <stat> ‘end’
// SPECIAL NEEDS
type ScopeNode struct {
	Position
	StatElem StatementNode
}

type AssignLHSNode struct {
	Position
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
	Fst ExprNode
	Snd ExprNode
}

type AssignRHSNode struct {
	Position
	AssignRHSElem Node
}

type ExprNode struct {
	Position
	ExprElem Node
}

type ArrayLiterNode struct {
	Position
	Expr []ExprNode
}

type NewPairNode struct {
	Position
	Expr ExprNode // need two
}

type ArgListNode struct {
	Position
	Expr []ExprNode // one or more args list
}

type TypeNode struct {
	Position
	TypeElem Node
}

type BaseTypeNode struct {
	Position
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
	Type TypeNode
}

type PairTypeNode struct {
	Position
	PairElemType []PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	Position
	PairElemTypeElem Node
}

type IntLiterNode struct {
	Position
	IntSign IntSignNode // optional
	Digit   DigitNode   // repeatable
}

type BoolLiterNode struct {
	Position
	BoolLiter bool
}

type CharLiterNode struct {
	Position
	CharLiterElem CharacterNode
}

type StringLiterNode struct {
	Position
	StrLiter []CharacterNode
}

type CharacterNode struct {
	Position
	// any ascii characterr
	CharacterElem Node
}

type EscapedCharNode struct {
	Position
	EscapedCharElem Node
}

type NullNode struct {
	Position
	NullTerminator grammar.ItemType
}

type BackslashNode struct {
	Position
	Backslash grammar.ItemType
}

type TabNode struct {
	Position
	Tab grammar.ItemType
}

type LineFeedNode struct {
	Position
	LineFeed grammar.ItemType
}

type FormFeedNode struct {
	Position
	FormFeed grammar.ItemType
}

type CarriageReturnNode struct {
	Position
	CarriageReturn grammar.ItemType
}

type DoubleQuoteNode struct {
	Position
	DoubleQuote grammar.ItemType
}

type SingleQuoteNode struct {
	Position
	SingleQuote grammar.ItemType
}

/// NEED TO IMPLEMENT '\'

type PairLiterNode struct {
	Position
	PairLiter grammar.ItemType //should be NULL
}

type UnaryOpExprNode struct {
	Position
	UnaryOper UnaryOpNode
	Expr      ExprNode
}

type EBENode struct {
	Position
	Expr       []ExprNode
	BinaryOper BinaryOpNode
	// one either side of expr
}

// ( <expr> ) inside expr handled by parser

type UnaryOpNode struct {
	Position
	UnaryOperElem Node
}

type NotNode struct {
	Position
	Not grammar.ItemType
}

type NegNode struct {
	Position
	Neg grammar.ItemType
}

type LenNode struct {
	Position
	Len grammar.ItemType
}

type OrdNode struct {
	Position
	Ord grammar.ItemType
}

type ChrNode struct {
	Position
	Chr grammar.ItemType
}

type BinaryOpNode struct {
	Position
	BinaryOperElem *Node
}

type MultNode struct {
	Position
	Mult grammar.ItemType
}

type DivNode struct {
	Position
	Div grammar.ItemType
}

type ModNode struct {
	Position
	Mod grammar.ItemType
}

type AddNode struct {
	Position
	Add grammar.ItemType
}

type SubNode struct {
	Position
	Sub grammar.ItemType
}

type GTNode struct {
	Position
	Gt grammar.ItemType
}

type GTENode struct {
	Position
	Gte grammar.ItemType
}

type LTNode struct {
	Position
	Lt grammar.ItemType
}

type LTENode struct {
	Position
	Lte grammar.ItemType
}

type EQNode struct {
	Position
	Eq grammar.ItemType
}

type NEQNode struct {
	Position
	Neq grammar.ItemType
}

type AndNode struct {
	Position
	And grammar.ItemType
}

type OrNode struct {
	Position
	Or grammar.ItemType
}

type DigitNode struct {
	Position
	//	DigitElem //need to include digit in grammar then do grammar.DIGIT
	Digit int
}

type IntSignNode struct {
	Position
	IntSignElem Node
}

type PositiveNode struct {
	Position
	Add grammar.ItemType
}

type NegativeNode struct {
	Position
	Sub grammar.ItemType
}

//no need for comment node

/*
func (p *ProgramNode) BuildNode() Node {
	p.Pos = BuildArguments.Pos
	etcccccc
	return p
}
*/

func (p *ProgramNode) BuildNode(buildArguments BuildArguments) Node {
	/*
		type ProgramNode struct {
			Position
			Func []FunctionNode //rename funcslice
			Stat []StatementNode
		}*/
	p.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		p.Func[i] = node.(FunctionNode)
	}

	for i, node := range buildArguments.ChildListTwo { // []Node
		p.Stat[i] = node.(StatementNode)
	}

	return p
}

func (f FunctionNode) BuildNode(buildArguments BuildArguments) Node {
	return f
}

func (p ParameterNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (s StatementNode) BuildNode(buildArguments BuildArguments) Node {
	return s
}

func (d DeclarationNode) BuildNode(buildArguments BuildArguments) Node {
	return d
}

func (a AssignmentNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (r ReadNode) BuildNode(buildArguments BuildArguments) Node {
	return r
}

func (f FreeNode) BuildNode(buildArguments BuildArguments) Node {
	return f
}

func (r ReturnNode) BuildNode(buildArguments BuildArguments) Node {
	return r
}

func (e ExitNode) BuildNode(buildArguments BuildArguments) Node {
	return e
}

func (p PrintNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (p PrintlnNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (i IfNode) BuildNode(buildArguments BuildArguments) Node {
	return i
}

func (w WhileNode) BuildNode(buildArguments BuildArguments) Node {
	return w
}

func (s ScopeNode) BuildNode(buildArguments BuildArguments) Node {
	return s
}

func (a AssignLHSNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (a ArrayElemNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (p PairElemNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (a AssignRHSNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (e ExprNode) BuildNode(buildArguments BuildArguments) Node {
	return e
}

func (a ArrayLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (n NewPairNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (c CallNode) BuildNode(buildArguments BuildArguments) Node {
	return c
}

func (a ArgListNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (t TypeNode) BuildNode(buildArguments BuildArguments) Node {
	return t
}

func (b BaseTypeNode) BuildNode(buildArguments BuildArguments) Node {
	return b
}

func (a ArrayTypeNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (p PairTypeNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (p PairElemTypeNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (i IntLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return i
}

func (b BoolLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return b
}

func (c CharLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return c
}

func (s StringLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return s
}
func (c CharacterNode) BuildNode(buildArguments BuildArguments) Node {
	return c
}

func (e EscapedCharNode) BuildNode(buildArguments BuildArguments) Node {
	return e
}

func (n NullNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (b BackslashNode) BuildNode(buildArguments BuildArguments) Node {
	return b
}

func (t TabNode) BuildNode(buildArguments BuildArguments) Node {
	return t
}

func (l LineFeedNode) BuildNode(buildArguments BuildArguments) Node {
	return l
}

func (f FormFeedNode) BuildNode(buildArguments BuildArguments) Node {
	return f
}

func (c CarriageReturnNode) BuildNode(buildArguments BuildArguments) Node {
	return c
}

func (d DoubleQuoteNode) BuildNode(buildArguments BuildArguments) Node {
	return d
}

func (s SingleQuoteNode) BuildNode(buildArguments BuildArguments) Node {
	return s
}

func (p PairLiterNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (u UnaryOpExprNode) BuildNode(buildArguments BuildArguments) Node {
	return u
}

func (e EBENode) BuildNode(buildArguments BuildArguments) Node {
	return e
}

func (u UnaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	return u
}

func (n NotNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (n NegNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (l LenNode) BuildNode(buildArguments BuildArguments) Node {
	return l
}

func (o OrdNode) BuildNode(buildArguments BuildArguments) Node {
	return o
}

func (c ChrNode) BuildNode(buildArguments BuildArguments) Node {
	return c
}

func (b BinaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	return b
}

func (m MultNode) BuildNode(buildArguments BuildArguments) Node {
	return m
}

func (d DivNode) BuildNode(buildArguments BuildArguments) Node {
	return d
}

func (m ModNode) BuildNode(buildArguments BuildArguments) Node {
	return m
}

func (a AddNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (s SubNode) BuildNode(buildArguments BuildArguments) Node {
	return s
}

func (g GTNode) BuildNode(buildArguments BuildArguments) Node {
	return g
}

func (g GTENode) BuildNode(buildArguments BuildArguments) Node {
	return g
}

func (l LTNode) BuildNode(buildArguments BuildArguments) Node {
	return l
}

func (l LTENode) BuildNode(buildArguments BuildArguments) Node {
	return l
}

func (e EQNode) BuildNode(buildArguments BuildArguments) Node {
	return e
}

func (n NEQNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (a AndNode) BuildNode(buildArguments BuildArguments) Node {
	return a
}

func (o OrNode) BuildNode(buildArguments BuildArguments) Node {
	return o
}

func (d DigitNode) BuildNode(buildArguments BuildArguments) Node {
	return d
}

func (i IntSignNode) BuildNode(buildArguments BuildArguments) Node {
	return i
}

func (p PositiveNode) BuildNode(buildArguments BuildArguments) Node {
	return p
}

func (n NegativeNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}
