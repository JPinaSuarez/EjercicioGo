package main

import "fmt"

func sumar(a, b int) int {
	return a + b
}
func main() {
	numero1 := 10
	numero2 := 5
	resultado := sumar(numero1, numero2)
	fmt.Println("La suma es:", resultado)
}
