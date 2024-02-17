package ast

import (
	"fmt"
	gr "glox/pkg/grammar"
)

type AstPrinter struct{}

func (ap AstPrinter) Print(ex gr.Expr[string]) string {
	return ex.Accept(ap)
}

// Implementation of Visitor interface
func (ap AstPrinter) VisitBinary(b gr.Binary[string]) string {
	return ap.parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

func (ap AstPrinter) VisitGrouping(g gr.Grouping[string]) string {
	return ap.parenthesize("group", g.Expression)
}

func (ap AstPrinter) VisitLiteral(l gr.Literal[string]) string {
	return fmt.Sprintf("%v", l.Value)
}

func (ap AstPrinter) VisitUnary(u gr.Unary[string]) string {
	return ap.parenthesize(u.Operator.Lexeme, u.Expression)
}

func (ap AstPrinter) parenthesize(name string, exprs ...gr.Expr[string]) string {
	result := "(" + name
	for _, expr := range exprs {
		result += " " + expr.Accept(ap)
	}
	result += ")"
	return result
}
