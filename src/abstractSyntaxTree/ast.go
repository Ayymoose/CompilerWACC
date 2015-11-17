package abstractSyntaxTree

import (
	"grammar"
)

// Node interface
type Node interface {
	BuildNode(buildArguments BuildArguments) Node
}

// Struct containing information with
// which to construct individual nodes
type BuildArguments struct {
	Pos          int
	Type         grammar.Type
	Ident        string
	StrVal       string
	IntVal       int
	ChildListOne []Node
	ChildListTwo []Node
	BoolVal      bool
	DigitVal     int
}

// Structs for common fields
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
	Type  grammar.Type
}

// Utility Struct
type DataNode struct {
	Position
	IntVal int
	StrVal string
	Nodes  []Node
}

type NullChildNode struct {
}

// Root node of AST
type ProgramNode struct {
	Position
	FuncList []FunctionNode
	StatList []StatementNode
}

type FunctionNode struct {
	PosIdentType
	ParamList []ParameterNode
	StatList  []StatementNode
}

type ParameterNode struct {
	PosIdentType
}

type StatementNode struct {
	Position
	Stat Node
}

type DeclarationNode struct {
	PosIdentType
	AssignRHS AssignRHSNode
}

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
	StatList []StatementNode
}

type WhileNode struct {
	Position
	Expr     ExprNode
	StatList []StatementNode
}

// ‘begin’ <stat> ‘end’
type ScopeNode struct {
	Position
	StatList []StatementNode
}

type AssignLHSNode struct {
	Position
	AssignLHS Node
}

type IdentNode struct {
	PosIdent
}

type ArrayElemNode struct {
	PosIdent
	ExprList []ExprNode
}

type PairElemNode struct {
	Position
	Fst ExprNode
	Snd ExprNode
}

type AssignRHSNode struct {
	Position
	AssignRHS Node
}

type ExprNode struct {
	Position
	Expr Node
}

type ArrayLiterNode struct {
	Position
	ExprList []ExprNode
}

type NewPairNode struct {
	Position
	ExprList []ExprNode
}

type CallNode struct {
	PosIdent
	ArgList []ArgListNode
}

type ArgListNode struct {
	Position
	ExprList []ExprNode // one or more args list
}

type TypeNode struct {
	Position
	Type Node
}

type BaseTypeNode struct {
	Position
	BaseType grammar.Type //Node
}

/*
type IntNode struct {
	Position
	Int int
}

type BoolNode struct {
	Position
	Bool bool
}

type CharNode struct {
	Position
	Char rune
}

type StringNode struct {
	Position
	String string
}
*/

type ArrayTypeNode struct {
	Position
	Type TypeNode
}

type PairTypeNode struct {
	Position
	PairElemTypeList []PairElemTypeNode // need two for pair
}

type PairElemTypeNode struct {
	Position
	PairElemType Node
}

type IntLiterNode struct {
	Position
	IntSign   IntSignNode // optional
	DigitList []DigitNode // repeatable -> can change to int
}

type BoolLiterNode struct {
	Position
	BoolLiter bool
}

type CharLiterNode struct {
	Position
	CharLiter CharacterNode
}

type CharacterNode struct {
	Position
	// Any ascii character
	Character Node
}

type StrLiterNode struct {
	Position
	StrLiterList []CharacterNode
}

type EscapedCharNode struct {
	Position
	EscapedChar Node
}

type PairLiterNode struct {
	Position
	PairLiter grammar.Type //could be NULLNode
}

type NullNode struct {
	Position
	NullTerminator grammar.Type
}

type UnaryOpExprNode struct {
	Position
	UnaryOp UnaryOpNode
	Expr    ExprNode
}

type BinaryOpExprNode struct {
	Position
	Expr     []ExprNode
	BinaryOp BinaryOpNode
}

// ( <expr> ) inside expr handled by parser

type UnaryOpNode struct {
	Position
	UnaryOpType Node
}

type NotNode struct {
	Position
	Not grammar.Type
}

type NegNode struct {
	Position
	Neg grammar.Type
}

type LenNode struct {
	Position
	Len grammar.Type
}

type OrdNode struct {
	Position
	Ord grammar.Type
}

type ChrNode struct {
	Position
	Chr grammar.Type
}

type BinaryOpNode struct {
	Position
	BinaryOpType Node
}

type MultNode struct {
	Position
	Mult grammar.Type
}

type DivNode struct {
	Position
	Div grammar.Type
}

type ModNode struct {
	Position
	Mod grammar.Type
}

type AddNode struct {
	Position
	Add grammar.Type
}

type SubNode struct {
	Position
	Sub grammar.Type
}

type GTNode struct {
	Position
	GT grammar.Type
}

type GTENode struct {
	Position
	GTE grammar.Type
}

type LTNode struct {
	Position
	LT grammar.Type
}

type LTENode struct {
	Position
	LTE grammar.Type
}

type EQNode struct {
	Position
	EQ grammar.Type
}

type NEQNode struct {
	Position
	NEQ grammar.Type
}

type AndNode struct {
	Position
	And grammar.Type
}

type OrNode struct {
	Position
	Or grammar.Type
}

type DigitNode struct {
	Position
	Digit int
}

type IntSignNode struct {
	Position
	IntSign Node
}

type PositiveNode struct {
	Position
	Positive grammar.Type
}

type NegativeNode struct {
	Position
	Neg grammar.Type
}

// Functions which build individual nodes

func (n NullChildNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

func (d DataNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.IntVal = buildArguments.IntVal
	d.StrVal = buildArguments.StrVal

	for i, node := range buildArguments.ChildListOne { // []Node
		d.Nodes[i] = node.(Node) // Don't think we nned to cast it
	}
	return d
}

func (p ProgramNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		p.FuncList[i] = node.(FunctionNode)
	}

	for j, node := range buildArguments.ChildListTwo { // []Node
		p.StatList[j] = node.(StatementNode)
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

	for i, node := range buildArguments.ChildListTwo { // []Node
		w.StatList[i] = node.(StatementNode)
	}
	return w
}

func (s ScopeNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		s.StatList[i] = node.(StatementNode)
	}
	return s
}

func (a AssignLHSNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.AssignLHS = buildArguments.ChildListOne[0]
	return a
}

func (i IdentNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.Ident = buildArguments.Ident
	return i
}

func (a ArrayElemNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		a.ExprList[i] = node.(ExprNode)
	}
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
	a.AssignRHS = buildArguments.ChildListOne[0].(Node)
	return a
}

func (e ExprNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.Expr = buildArguments.ChildListOne[0].(Node)
	return e
}

func (a ArrayLiterNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		a.ExprList[i] = node.(ExprNode)
	}
	return a
}

func (n NewPairNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		n.ExprList[i] = node.(ExprNode)
	}
	return n
}

func (c CallNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Ident = buildArguments.Ident

	for i, node := range buildArguments.ChildListOne { // []Node
		c.Arglist[i] = node.(ArgListNode)
	}
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
	t.Pos = buildArguments.Pos
	t.Type = buildArguments.ChildListOne[0].(Node)
	return t
}

func (b BaseTypeNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BaseType = buildArguments.Type
	return b
}

/*
Insert base type build functions here if needed
*/

func (a ArrayTypeNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.Type = buildArguments.ChildListOne[0].(TypeNode)
	return a
}

func (p PairTypeNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		p.PairElemTypeList[i] = node.(PairElemTypeNode)
	}
	return p
}

func (p PairElemTypeNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.PairElemTypeElem = buildArguments.ChildListOne[0].(PairElemTypeNode)
	return p
}

func (i IntLiterNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.IntSign = buildArguments.ChildListOne[0].(IntSignNode)
	i.Digit = buildArguments.ChildListTwo[0].(DigitNode)
	return i
}

func (b BoolLiterNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BoolLiter = buildArguments.BoolVal
	return b
}

func (c CharLiterNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.CharLiter = buildArguments.ChildListOne[0].(CharacterNode)
	return c
}

func (s StrLiterNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		s.StrLiterList[i] = node.(CharacterNode)
	}
	return s
}
func (c CharacterNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Character = buildArguments.ChildListOne[0].(CharacterNode)
	return c
}

func (e EscapedCharNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.EscapedChar = buildArguments.ChildListOne[0].(Node)
	return e
}

func (n NullNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.NullTerminator = buildArguments.Type
	return n
}

func (p PairLiterNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.PairLiter = buildArguments.Type
	return p
}

func (u UnaryOpExprNode) BuildNode(buildArguments BuildArguments) Node {
	u.Pos = buildArguments.Pos
	u.UnaryOper = buildArguments.ChildListOne[0].(UnaryOpNode)
	u.Expr = buildArguments.ChildListTwo[0].(ExprNode)
	return u
}

func (b BinaryOpExprNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		b.Expr[i] = node.(ExprNode)
	}
	b.BinaryOp = buildArguments.ChildListTwo[0].(BinaryOpNode)
	return e
}

func (u UnaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	u.Pos = buildArguments.Pos
	u.UnaryOpType = buildArguments.ChildListOne[0].(Node)
	return u
}

func (n NotNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Not = buildArguments.Type
	return n
}

func (n NegNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Neg = buildArguments.Type
	return n
}

func (l LenNode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.Len = buildArguments.Type
	return l
}

func (o OrdNode) BuildNode(buildArguments BuildArguments) Node {
	o.Pos = buildArguments.Pos
	o.Ord = buildArguments.Type
	return o
}

func (c ChrNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Chr = buildArguments.Type
	return c
}

func (b BinaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BinaryOpType = buildArguments.ChildListOne[0].(Node)
	return b
}

func (m MultNode) BuildNode(buildArguments BuildArguments) Node {
	m.Pos = buildArguments.Pos
	m.Mult = buildArguments.Type
	return m
}

func (d DivNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.Div = buildArguments.Type
	return d
}

func (m ModNode) BuildNode(buildArguments BuildArguments) Node {
	m.Pos = buildArguments.Pos
	m.Mod = buildArguments.Type
	return m
}

func (a AddNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.Add = buildArguments.Type
	return a
}

func (s SubNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos
	s.Sub = buildArguments.Type
	return s
}

func (g GTNode) BuildNode(buildArguments BuildArguments) Node {
	g.Pos = buildArguments.Pos
	g.GT = buildArguments.Type
	return g
}

func (g GTENode) BuildNode(buildArguments BuildArguments) Node {
	g.Pos = buildArguments.Pos
	g.GTE = buildArguments.Type
	return g
}

func (l LTNode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.LT = buildArguments.Type
	return l
}

func (l LTENode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.LTE = buildArguments.Type
	return l
}

func (e EQNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.EQ = buildArguments.Type
	return e
}

func (n NEQNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.NEQ = buildArguments.Type
	return n
}

func (a AndNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.And = buildArguments.Type
	return a
}

func (o OrNode) BuildNode(buildArguments BuildArguments) Node {
	o.Pos = buildArguments.Pos
	o.Or = buildArguments.Type
	return o
}

func (d DigitNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.Digit = buildArguments.DigitVal
	return d
}

func (i IntSignNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.IntSign = buildArguments.ChildListOne[0].(Node)
	return i
}

func (p PositiveNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.Positive = buildArguments.Type
	return p
}

func (n NegativeNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Neg = buildArguments.Type
	return n
}
