package grammar

import (
	tk "glox/pkg/tokens"
)

type ReturnType interface{ string | float64 | int }

type Expr[R ReturnType] interface {
	Accept(v Visitor[R]) R
}

type Binary[R ReturnType] struct {
	Left     Expr[R]
	Operator tk.Token
	Right    Expr[R]
}

type Grouping[R ReturnType] struct {
	Expression Expr[R]
}

type Literal[R ReturnType] struct {
	Value any
}

type Unary[R ReturnType] struct {
	Operator   tk.Token
	Expression Expr[R]
}

func NewBinary[R ReturnType](left Expr[R], operator tk.Token, right Expr[R]) Binary[R] {
	return Binary[R]{left, operator, right}
}

func NewGrouping[R ReturnType](expression Expr[R]) Grouping[R] {
	return Grouping[R]{expression}
}

func NewLiteral[R ReturnType](value any) Literal[R] {
	return Literal[R]{value}
}

func NewUnary[R ReturnType](operator tk.Token, expression Expr[R]) Unary[R] {
	return Unary[R]{operator, expression}
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
