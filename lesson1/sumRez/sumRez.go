//Package sumRez returns summa and b var int64
package sumRez

import (
	"fmt"
	"https://github.com/Max007446/Go_Level2"
)

// sumRez returns summa and b var int64
func sumRez(a int64, b int64) int64 {
	return a + b
}
func main() {
	var x, y int64
	fmt.Printf("Введите первое слагаемое для суммы \n")
	fmt.Scan(&x)
	fmt.Printf("Введите второе слагаемое для суммы \n")
	fmt.Scan(&y)
	fmt.Println("Сумма = ", sumRez(x, y))
}
