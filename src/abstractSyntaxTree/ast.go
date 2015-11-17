package abstractSyntaxTree

import (
	"grammar"
)

// All AST nodes implement this Node interface
type Node interface {
	BuildNode(buildArguments BuildArguments) Node
}

/* Utility structs -----------------------------------------------------------*/
/* Common field structs ------------------------------------------------------*/
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

/* End of common field structs -----------------------------------------------*/

// Argument struct contains information with which to build individual AST
// nodes. Undeclared fields will default to their uninitialised nil value
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

// Holds information with which to construct further AST nodes
type DataNode struct {
	Position
	IntVal int
	StrVal string
	Nodes  []Node
	Type   grammar.Type
}

/* End of utility structs ----------------------------------------------------*/

/* Structs for WACC BNF ------------------------------------------------------*/

// ‘begin’ <func>* <stat> ‘end’
type ProgramNode struct {
	Position
	FuncList []FunctionNode
	StatList []StatementNode
}

// <type> <ident> ‘(’ <param-list> ‘)’ ‘is’ <stat> ‘end’
type FunctionNode struct {
	PosIdentType
	ParamList []ParameterNode
	StatList  []StatementNode
}

// <type> <ident>
type ParameterNode struct {
	PosIdentType
}

// <stat>
type StatementNode struct {
	Position
	Stat Node
}

// <type> <ident> ‘=’ <assign-rhs>
type DeclarationNode struct {
	PosIdentType
	AssignRHS AssignRHSNode
}

// <assign-lhs> ‘=’ <assign-rhs>
type AssignmentNode struct {
	Position
	AssignLHS AssignLHSNode
	AssignRHS AssignRHSNode
}

// ‘read’ <assign-lhs>
type ReadNode struct {
	Position
	AssignLHS AssignLHSNode
}

// ‘free’ <expr>
type FreeNode struct {
	Position
	Expr ExprNode
}

// ‘return’ <expr>
type ReturnNode struct {
	Position
	Expr ExprNode
}

// ‘exit’ <expr>
type ExitNode struct {
	Position
	Expr ExprNode
}

// ‘print’ <expr>
type PrintNode struct {
	Position
	Expr ExprNode
}

// ‘println’ <expr>
type PrintlnNode struct {
	Position
	Expr ExprNode
}

// ‘if’ <expr> ‘then’ <stat> ‘else’ <stat> ‘fi’
type IfNode struct {
	Position
	Expr     ExprNode
	StatList []StatementNode
}

// ‘while’ <expr> ‘do’ <stat> ‘done’
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

// <assign-lhs>
type AssignLHSNode struct {
	Position
	AssignLHS Node
}

// <ident>
type IdentNode struct {
	PosIdent
}

// <array-elem>
type ArrayElemNode struct {
	PosIdent
	ExprList []ExprNode
}

// ‘fst’ <expr> | ‘snd’ <expr >
type PairElemNode struct {
	Position
	Fst ExprNode
	Snd ExprNode
}

// <assign-rhs>
type AssignRHSNode struct {
	Position
	AssignRHS Node
}

// <expr>
type ExprNode struct {
	Position
	Expr Node
}

// <array-liter>
type ArrayLiterNode struct {
	Position
	ExprList []ExprNode
}

// ‘newpair’ ‘(’ <expr> ‘,’ <expr> ‘)’
type NewPairNode struct {
	Position
	ExprList []ExprNode
}

// <pair-elem>
type CallNode struct {
	PosIdent
	ArgList []ArgListNode
}

// <arg-list>
type ArgListNode struct {
	Position
	ExprList []ExprNode // one or more args list
}

// <type>
type TypeNode struct {
	Position
	Type Node
}

// 'int' / 'bool' / 'char' / 'string'
type BaseTypeNode struct {
	Position
	BaseType grammar.Type
}

// <type> ‘[’ ‘]’
type ArrayTypeNode struct {
	Position
	Type TypeNode
}

// ‘pair’ ‘(’ <pair-elem-type> ‘,’ <pair-elem-type> ‘)’
type PairTypeNode struct {
	Position
	PairElemTypeList []PairElemTypeNode
}

// <pair-elem-type>
type PairElemTypeNode struct {
	Position
	PairElemType Node
}

// <int-sign>? <digit>+
type IntLiterNode struct {
	Position
	IntSign   IntSignNode // optional
	DigitList []DigitNode // repeatable -> can change to int
}

// 'true' | 'false'
type BoolLiterNode struct {
	Position
	BoolLiter bool
}

// ‘’’ <character> ‘’’
type CharLiterNode struct {
	Position
	CharLiter CharacterNode
}

// any-ASCII-character-except-‘\’-‘’’-‘"’ | ‘\’ <escaped-char>
type CharacterNode struct {
	Position
	Character Node
}

// ‘"’ <character>* ‘"’
type StrLiterNode struct {
	Position
	StrLiterList []CharacterNode
}

// ‘0’ | ‘b’ | ‘t’ | ‘n’ | ‘f’ | ‘r’ | ‘"’ | ‘’’ | ‘\’
type EscapedCharNode struct {
	Position
	EscapedChar Node
}

// 'null'
type PairLiterNode struct {
	Position
	PairLiter grammar.Type
}

// '/0'
type NullNode struct {
	Position
	NullTerminator grammar.Type
}

// <unary-oper> <expr>
type UnaryOpExprNode struct {
	Position
	UnaryOp UnaryOpNode
	Expr    ExprNode
}

// <expr> <binary-oper> <expr>
type BinaryOpExprNode struct {
	Position
	Expr     []ExprNode
	BinaryOp BinaryOpNode
}

// '!' | '-' | 'len' | 'ord' | 'chr'
type UnaryOpNode struct {
	Position
	UnaryOpType Node
}

// '!'
type NotNode struct {
	Position
	Not grammar.Type
}

// '-'
type NegNode struct {
	Position
	Neg grammar.Type
}

// 'len'
type LenNode struct {
	Position
	Len grammar.Type
}

// 'ord'
type OrdNode struct {
	Position
	Ord grammar.Type
}

// 'chr'
type ChrNode struct {
	Position
	Chr grammar.Type
}

// '*' | '/' | '%' | '+' | '-' | '>' | '>=' | '<' | '<=' | '==' | '!=' | '&&'
// '||'
type BinaryOpNode struct {
	Position
	BinaryOpType Node
}

// '*'
type MultNode struct {
	Position
	Mult grammar.Type
}

// '/'
type DivNode struct {
	Position
	Div grammar.Type
}

// '%'
type ModNode struct {
	Position
	Mod grammar.Type
}

// '+'
type AddNode struct {
	Position
	Add grammar.Type
}

// '-'
type SubNode struct {
	Position
	Sub grammar.Type
}

// '>'
type GTNode struct {
	Position
	GT grammar.Type
}

// '>='
type GTENode struct {
	Position
	GTE grammar.Type
}

// '<'
type LTNode struct {
	Position
	LT grammar.Type
}

// '<='
type LTENode struct {
	Position
	LTE grammar.Type
}

// '=='
type EQNode struct {
	Position
	EQ grammar.Type
}

// '!='
type NEQNode struct {
	Position
	NEQ grammar.Type
}

// '&&'
type AndNode struct {
	Position
	And grammar.Type
}

// '||'
type OrNode struct {
	Position
	Or grammar.Type
}

// ('0'-'9')
type DigitNode struct {
	Position
	Digit int
}

// '+' | '-'
type IntSignNode struct {
	Position
	IntSign Node
}

// '+'
type PositiveNode struct {
	Position
	Positive grammar.Type
}

// '-'
type NegativeNode struct {
	Position
	Neg grammar.Type
}

// Null node used for error handling in parser
type NullChildNode struct {
}

/* End of structs for WACC BNF -----------------------------------------------*/

/* Interface method called on individual node structs -----------------------*/

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
	s.Stat = buildArguments.ChildListOne[0]
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

//Builds CallNode
func (c CallNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Ident = buildArguments.Ident

	for i, node := range buildArguments.ChildListOne { // []Node
		c.ArgList[i] = node.(ArgListNode)
	}
	return c
}

//Builds ArgListNode
func (a ArgListNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne { // []Node
		a.ExprList[i] = node.(ExprNode)
	}
	return a
}

//Builds TypeNode
func (t TypeNode) BuildNode(buildArguments BuildArguments) Node {
	t.Pos = buildArguments.Pos
	t.Type = buildArguments.ChildListOne[0].(Node)
	return t
}

//Builds BaseTypeNode
func (b BaseTypeNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BaseType = buildArguments.Type
	return b
}

//Builds ArrayTypeNode
func (a ArrayTypeNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.Type = buildArguments.ChildListOne[0].(TypeNode)
	return a
}

//Builds PairTypeNode
func (p PairTypeNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		p.PairElemTypeList[i] = node.(PairElemTypeNode)
	}
	return p
}

//Builds PairElemTypeNode
func (p PairElemTypeNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.PairElemType = buildArguments.ChildListOne[0].(PairElemTypeNode)
	return p
}

//Builds IntLiterNode
func (i IntLiterNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.IntSign = buildArguments.ChildListOne[0].(IntSignNode)
	for j, node := range buildArguments.ChildListTwo {
		i.DigitList[j] = node.(DigitNode)
	}
	return i
}

//Builds BoolLiterNode
func (b BoolLiterNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BoolLiter = buildArguments.BoolVal
	return b
}

//Builds CharLiterNode
func (c CharLiterNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.CharLiter = buildArguments.ChildListOne[0].(CharacterNode)
	return c
}

//Builds StrLiterNode
func (s StrLiterNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		s.StrLiterList[i] = node.(CharacterNode)
	}
	return s
}

//Builds CharacterNode
func (c CharacterNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Character = buildArguments.ChildListOne[0].(CharacterNode)
	return c
}

//Builds EscapedCharNode
func (e EscapedCharNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.EscapedChar = buildArguments.ChildListOne[0].(Node)
	return e
}

//Builds NullNode
func (n NullNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.NullTerminator = buildArguments.Type
	return n
}

//Builds PairLiterNode
func (p PairLiterNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.PairLiter = buildArguments.Type
	return p
}

//Builds UnaryOpExprNode
func (u UnaryOpExprNode) BuildNode(buildArguments BuildArguments) Node {
	u.Pos = buildArguments.Pos
	u.UnaryOp = buildArguments.ChildListOne[0].(UnaryOpNode)
	u.Expr = buildArguments.ChildListTwo[0].(ExprNode)
	return u
}

//Builds BinaryOpExprNode
func (b BinaryOpExprNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos

	for i, node := range buildArguments.ChildListOne {
		b.Expr[i] = node.(ExprNode)
	}
	b.BinaryOp = buildArguments.ChildListTwo[0].(BinaryOpNode)
	return b
}

//Builds UnaryOpNode
func (u UnaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	u.Pos = buildArguments.Pos
	u.UnaryOpType = buildArguments.ChildListOne[0].(Node)
	return u
}

//Builds NotNode
func (n NotNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Not = buildArguments.Type
	return n
}

//Builds NegNode
func (n NegNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Neg = buildArguments.Type
	return n
}

//Builds LenNode
func (l LenNode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.Len = buildArguments.Type
	return l
}

//Builds OrdNode
func (o OrdNode) BuildNode(buildArguments BuildArguments) Node {
	o.Pos = buildArguments.Pos
	o.Ord = buildArguments.Type
	return o
}

//Builds ChrNode
func (c ChrNode) BuildNode(buildArguments BuildArguments) Node {
	c.Pos = buildArguments.Pos
	c.Chr = buildArguments.Type
	return c
}

//Builds BinaryOpNode
func (b BinaryOpNode) BuildNode(buildArguments BuildArguments) Node {
	b.Pos = buildArguments.Pos
	b.BinaryOpType = buildArguments.ChildListOne[0].(Node)
	return b
}

//Builds MultNode
func (m MultNode) BuildNode(buildArguments BuildArguments) Node {
	m.Pos = buildArguments.Pos
	m.Mult = buildArguments.Type
	return m
}

//Builds DivNode
func (d DivNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.Div = buildArguments.Type
	return d
}

//Builds ModNode
func (m ModNode) BuildNode(buildArguments BuildArguments) Node {
	m.Pos = buildArguments.Pos
	m.Mod = buildArguments.Type
	return m
}

//Builds AddNode
func (a AddNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.Add = buildArguments.Type
	return a
}

//Builds SubNode
func (s SubNode) BuildNode(buildArguments BuildArguments) Node {
	s.Pos = buildArguments.Pos
	s.Sub = buildArguments.Type
	return s
}

//Builds GTNode
func (g GTNode) BuildNode(buildArguments BuildArguments) Node {
	g.Pos = buildArguments.Pos
	g.GT = buildArguments.Type
	return g
}

//Builds GTENode
func (g GTENode) BuildNode(buildArguments BuildArguments) Node {
	g.Pos = buildArguments.Pos
	g.GTE = buildArguments.Type
	return g
}

//Builds LTNode
func (l LTNode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.LT = buildArguments.Type
	return l
}

//Builds LTENode
func (l LTENode) BuildNode(buildArguments BuildArguments) Node {
	l.Pos = buildArguments.Pos
	l.LTE = buildArguments.Type
	return l
}

//Builds EQNode
func (e EQNode) BuildNode(buildArguments BuildArguments) Node {
	e.Pos = buildArguments.Pos
	e.EQ = buildArguments.Type
	return e
}

//Builds NEQNode
func (n NEQNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.NEQ = buildArguments.Type
	return n
}

//Builds AndNode
func (a AndNode) BuildNode(buildArguments BuildArguments) Node {
	a.Pos = buildArguments.Pos
	a.And = buildArguments.Type
	return a
}

//Builds OrNode
func (o OrNode) BuildNode(buildArguments BuildArguments) Node {
	o.Pos = buildArguments.Pos
	o.Or = buildArguments.Type
	return o
}

//Builds DigitNode
func (d DigitNode) BuildNode(buildArguments BuildArguments) Node {
	d.Pos = buildArguments.Pos
	d.Digit = buildArguments.DigitVal
	return d
}

//Builds IntSignNode
func (i IntSignNode) BuildNode(buildArguments BuildArguments) Node {
	i.Pos = buildArguments.Pos
	i.IntSign = buildArguments.ChildListOne[0].(Node)
	return i
}

//Builds PositiveNode
func (p PositiveNode) BuildNode(buildArguments BuildArguments) Node {
	p.Pos = buildArguments.Pos
	p.Positive = buildArguments.Type
	return p
}

//Builds NegativeNode
func (n NegativeNode) BuildNode(buildArguments BuildArguments) Node {
	n.Pos = buildArguments.Pos
	n.Neg = buildArguments.Type
	return n
}

// Upcasts Null node to node
func (n NullChildNode) BuildNode(buildArguments BuildArguments) Node {
	return n
}

/* End of interface methods --------------------------------------------------*/
