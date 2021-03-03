package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Предупреждение:", v)
		}
	}()
	_, err = os.Open("max.go")
	if err != nil {
		panic(err)
	}

}
