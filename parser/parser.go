package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	var funcname = "parserfile"
	a, err := parserNew(funcname)
	fmt.Println("Количество вызовов асинхронных функций:", a, err)
}
func parserNew(fName string) (a int32, err error) {
	fset := token.NewFileSet()
	src, err := os.Open("parserfile.go")
	_ = err
	f, err := parser.ParseFile(fset, fName, src, 0)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.GoStmt:
			a++
		}
		return true
	})
	return a, err
}
