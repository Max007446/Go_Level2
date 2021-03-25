package parserfile

import (
	"fmt"
)

func parserfile() {
	fmt.Println("Основной поток")
	go func() {
		fmt.Println("one gorutine")
	}()
	go func() {
		fmt.Println("two gorutine")
	}()
	go func() {
		fmt.Println("three gorutine")
	}()
	go func() {
		fmt.Println("four gorutine")
	}()
}
func two() {
	fmt.Println("Function two")
}
func three() {
	fmt.Println("Function three")
}
