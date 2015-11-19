
//line grammar.y:2
package parser
import __yyfmt__ "fmt"
//line grammar.y:2
		
//line grammar.y:5
type parserSymType struct {
	yys int
  str  string
  prog *Program
  func *Function
  funcs []*Function
  stat *Statement
  stats []*Statement
  param *Param
  params []*Param
  assignLHS *AssignLHS
  assignRHS *AssignRHS
  args []*Arg
  arg *Arg
  pairElem *PairElem
  type *Type
  base_type *BaseType
  array_type *ArrayType
pair_type *PairType
pair_elem_type *PairElemType
expr *Expr
exprs []*Expr
aexprs []*Expr
unary_op *UnaryOp
binary_op *BinaryOp
ident *Ident
array_elem *ArrayElem
int_liter *IntLiter
int_sign *IntSign
bool_liter *BoolLiter
str_liter *StrLiter
pair_liter *PairLiter
}

