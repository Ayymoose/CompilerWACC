package abstractSyntaxTree

import (
	"grammar"
)

// Node interface
type Node interface {
	BuildNode(buildArguments BuildArguments) Node
}

type BuildArguments struct {
	Pos          int
	Type         grammar.ItemType
	Ident        string
	StrVal       string
	IntVal       int
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
	FuncList []FunctionNode //rename funcslice
	StatList []StatementNode
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
	Expr     ExprNode
	StatList []StatementNode // two stats
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
	Stat StatementNode
}

type AssignLHSNode struct {
	Position
	AssignLHSElem Node
}

type IdentNode struct {
	Position
	Ident string
}

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
	ExprList []ExprNode // one or more args list
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

type DataNode struct {
	Position
	IntVal int
	StrVal string
}

type NullChildNode struct {
}

//no need for comment node

/*
func (p *ProgramNode) BuildNode() Node {
	p.Pos = BuildArguments.Pos
	etcccccc
	return p
}
*/
func (n NullChildNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (d DataNode) BuildNode(buildArguments BuildArguments) Node {
	d.Position.Pos = buildArguments.Pos
	d.IntVal = buildArguments.IntVal
	d.StrVal = buildArguments.StrVal

	return d
}

func (p *ProgramNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		p.FuncList[i] = node.(FunctionNode)
	}

	for i, node := range buildArguments.ChildListTwo { // []Node
		p.StatList[i] = node.(StatementNode)
	}

	return p
}

func (f FunctionNode) BuildNode(buildArguments BuildArguments) Node {

	f.Pos = buildArguments.Pos
	f.Ident = buildArguments.Ident
	f.Type = buildArguments.Type

	for i, node := range buildArguments.ChildListOne { // []Node
		f.ParamList[i] = node.(ParameterNode)
	}

	for i, node := range buildArguments.ChildListTwo { // []Node
		f.StatList[i] = node.(StatementNode)
	}

	return f
}

func (p ParameterNode) BuildNode(buildArguments BuildArguments) Node {

	p.Pos = buildArguments.Pos
	p.Ident = buildArguments.Ident
	p.Type = buildArguments.Type

	return p
}

func (s StatementNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos
	s.StatElem = buildArguments.ChildListOne[0]

	return s
}

func (d DeclarationNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.Ident = buildArguments.Ident
	d.Type = buildArguments.Type
	d.AssignRHS = buildArguments.ChildListOne[0].(AssignRHSNode)

	return d
}

func (a AssignmentNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.AssignLHS = buildArguments.ChildListOne[0].(AssignLHSNode)
	a.AssignRHS = buildArguments.ChildListTwo[0].(AssignRHSNode)
	return a
}

func (r ReadNode) BuildNode(buildArguments BuildArguments) Node {
	r.Pos = buildArguments.Pos
	r.AssignLHS = buildArguments.ChildListOne[0].(AssignLHSNode)

	return r
}

func (f FreeNode) BuildNode(buildArguments BuildArguments) Node {
	f.Pos = buildArguments.Pos
	f.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return f
}

func (r ReturnNode) BuildNode(buildArguments BuildArguments) Node {
	r.Pos = buildArguments.Pos
	r.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return r
}

func (e ExitNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return e
}

func (p PrintNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return p
}

func (p PrintlnNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return p
}

func (i IfNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.Expr = buildArguments.ChildListOne[0].(ExprNode)

	for j, node := range buildArguments.ChildListTwo { // []Node
		i.StatList[j] = node.(StatementNode)
	}

	return i
}

func (w WhileNode) BuildNode(buildArguments BuildArguments) Node {
	w.Pos = buildArguments.Pos
	w.Expr = buildArguments.ChildListOne[0].(ExprNode)
	w.Stat = buildArguments.ChildListTwo[0].(StatementNode)

	return w
}

func (s ScopeNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos
	s.Stat = buildArguments.ChildListOne[0].(StatementNode)

	return s
}

func (a AssignLHSNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.AssignLHSElem = buildArguments.ChildListOne[0]

	return a
}

func (i IdentNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.Ident = buildArguments.Ident

	return i
}

func (a ArrayElemNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return a
}

func (p PairElemNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.Fst = buildArguments.ChildListOne[0].(ExprNode)
	p.Snd = buildArguments.ChildListTwo[0].(ExprNode)

	return p
}

func (a AssignRHSNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.AssignRHSElem = buildArguments.ChildListOne[0].(Node)

	return a
}

func (e ExprNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.ExprElem = buildArguments.ChildListOne[0].(Node)

	return e
}

func (a ArrayLiterNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		a.Expr[i] = node.(ExprNode)
	}

	return a
}

func (n NewPairNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Expr = buildArguments.ChildListOne[0].(ExprNode)

	return n
}

func (c CallNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Ident = buildArguments.Ident
	c.ArgList = buildArguments.ChildListOne[0].(ArgListNode)

	return c
}

func (a ArgListNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		a.ExprList[i] = node.(ExprNode)
	}

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
