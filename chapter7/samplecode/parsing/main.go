package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package main
func main() {
	println("Hello!")
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}
