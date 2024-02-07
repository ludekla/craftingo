package grammar

import tk "glox/pkg/tokens"

type Expr interface {
	Any
}

type Binary struct {
	left     Expr
	operator tk.Token
	right    Expr
}

type Grouping struct {
	expression Expr
}

type Literal struct {
	value Any
}

type Unary struct {
	operator   tk.Token
	expression Expr
}
