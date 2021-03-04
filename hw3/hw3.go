package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	file, err := os.Create("new.go")
	if err != nil {
		fmt.Println("not created file!")
	}
	defer file.Close()
	fmt.Println(file.Name())
}
