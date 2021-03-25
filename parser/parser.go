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
	parserNew(funcname)
}
func parserNew(fName string) {

	fset := token.NewFileSet()
	src, err := os.Open("parserfile.go")
	_ = err
	f, err := parser.ParseFile(fset, fName, src, parser.Mode(token.GO))
	//fmt.Println(f, fName)
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch n.(type) {
		case *ast.GoStmt:
			s = "go"
		}
		if s == "go" {
			fmt.Printf("%s:\t%s\n", s)
		}

		return true
	})
}
