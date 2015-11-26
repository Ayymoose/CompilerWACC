package main

import (
	. "abstractSyntaxTree"
	"fmt"
	"grammar"
)

func main() {
	b := BoolLiterNode{}
	e := ExprNode{Expr: b}
	a := AssignRHSNode{AssignRHS: e}
	d := DeclarationNode{AssignRHS: a}
	d.Ident = "b"
	d.Type = grammar.BaseBool
	statlist := []StatementNode{StatementNode{Stat: d}}
	//statlist
	p := ProgramNode{StatList: statlist}
	succeeded, errs := p.VisitProgram()
	fmt.Println("Succeeded:", succeeded)
	for _, error := range errs {
		fmt.Println(error)
	}
}
