package main

import (
	"errors"
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	return dividendo / divisor, nil
}
func main() {
	/*str := "123g"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Numero Conver:", num)
	*/
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Resultado", result)
}
