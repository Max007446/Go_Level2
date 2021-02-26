package main

import (
	"fmt"
)

func sumrez(a int64, b int64) int64 {
	return a + b
}
func main() {
	var x, y int64
	fmt.Printf("Введите первое слагаемое для суммы \n")
	fmt.Scan(&x)
	fmt.Printf("Введите второе слагаемое для суммы \n")
	fmt.Scan(&y)
	fmt.Println("Сумма = ", sumrez(x, y))
}
