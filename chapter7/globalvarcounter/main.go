package main

import (
	"fmt"
	"github.com/karrick/godirwalk"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func fix(dir string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			log.Println("working on file", fileName)
			gv := func(f string) int {
				pf, err := parser.ParseFile(fset, f, nil, 0)
				if err != nil {
					fmt.Println(err)
					return 0
				}

				for _, n := range pf.Decls {
					if fn, ok := n.(*ast.GenDecl); ok {
						if fn.Tok == token.VAR {
							return len(fn.Specs)
						}
					}
				}
				return 0
			}(fileName)

			ast.Inspect(file, func(n ast.Node) bool {
				if fn, ok := n.(*ast.FuncDecl); ok {
					log.Println("fn: ", fn.Name.Name)

					if fn.Name.Obj != nil {
						if fDeclaration, ok := fn.Name.Obj.Decl.(*ast.FuncDecl); ok {
							if len(fDeclaration.Type.Params.List) > 0 {
								// how many parameters are there
								log.Println("Len param  = ", len(fDeclaration.Type.Params.List))

								for _, k := range fDeclaration.Type.Params.List {
									if len(k.Names) > 0 {
										log.Println(". parameter name = ", k.Names[0])

										// is it function type ?
										fType, ok := k.Type.(*ast.FuncType)
										if fType != nil && ok {
											log.Println(".... parameter type (function) = ", fType.Params.List, " with ", len(fType.Params.List), " parameters ")
										}

										// is it a selector type ?
										fSelectorExpr, ok := k.Type.(*ast.SelectorExpr)
										if fSelectorExpr != nil && ok {
											fIdent, ok := fSelectorExpr.X.(*ast.Ident)
											if ok {
												log.Println(".... parameter type = ", fIdent.Name)
											}
										}

										// ..or just a nomal type ?
										if fIdent, ok := k.Type.(*ast.Ident); ok {
											if fIdent != nil && fIdent.Obj != nil && ok {
												log.Println(".... parameter type = ", fIdent.Obj.Name)
											} else {
												log.Println(".... parameter type = ", fIdent.Name)
											}
										}
									}
								}
							}
						}
					}
				}
				return true
			})

			log.Println("Total global variable found : ", gv)
		}
	}
}

func main() {
	traverse("/home/nanik/go/src/github.com/securego/gosec")
}

func traverse(mainDir string) {
	godirwalk.Walk(mainDir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if de.IsDir() {
				log.Println(fmt.Sprintf("%s", osPathname))
				fix(osPathname)
			}
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
}
