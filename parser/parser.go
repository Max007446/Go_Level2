package main

import (
	"fmt"
	"os"
)

func main() {
	var funcname = "parserfile"
	file, err := os.Open("parserfile.go")
	_ = err
	parserF(file, funcname)
	defer file.close()
}
func parserF(file *os.File, funcF string) {

	fmt.Println(file, funcF)
}
