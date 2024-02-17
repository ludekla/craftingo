package main

import (
	"fmt"
	"glox/pkg/ast"
	gr "glox/pkg/grammar"
	tk "glox/pkg/tokens"
)

func main() {
	expr := gr.NewBinary[string](
		gr.NewUnary[string](
			tk.NewToken(tk.MINUS, "-", 0.0, 1),
			gr.NewLiteral[string](123),
		),
		tk.NewToken(tk.STAR, "*", 0.0, 1),
		gr.NewGrouping[string](gr.NewLiteral[string](45.67)),
	)

	ap := ast.AstPrinter{}
	fmt.Println(ap.Print(expr))
}
