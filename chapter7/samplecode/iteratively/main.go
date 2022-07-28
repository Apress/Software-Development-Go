package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Imports:")
	for _, i := range f.Imports {
		log.Println(" ", i.Path.Value)
	}

	log.Println("Functions:")
	for _, f := range f.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		log.Println(" ", fn.Name.Name)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			log.Println(fmt.Sprintf("return statement found in line %d:", fset.Position(ret.Pos()).Line))
			return true
		}

		return true
	})
}
