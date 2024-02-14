package grammar

import tk "glox/pkg/tokens"

type ReturnType interface{ string | float64 | int }

type Expr[R ReturnType] interface {
	Accept(v Visitor[R]) R
}

type Binary[R ReturnType] struct {
	left     Expr[R]
	operator tk.Token
	right    Expr[R]
}

type Grouping[R ReturnType] struct {
	expression Expr[R]
}

type Literal[R ReturnType] struct {
	value Expr[R]
}

type Unary[R ReturnType] struct {
	operator   tk.Token
	expression Expr[R]
}

type Visitor[R ReturnType] interface {
	VisitBinary(b Binary[R]) R
	VisitGrouping(g Grouping[R]) R
	VisitLiteral(l Literal[R]) R
	VisitUnary(u Unary[R]) R
}

func (b Binary[R]) Accept(v Visitor[R]) R {
	return v.VisitBinary(b)
}

func (g Grouping[R]) Accept(v Visitor[R]) R {
	return v.VisitGrouping(g)
}

func (l Literal[R]) Accept(v Visitor[R]) R {
	return v.VisitLiteral(l)
}

func (u Unary[R]) Accept(v Visitor[R]) R {
	return v.VisitUnary(u)
}
