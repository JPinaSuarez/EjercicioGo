package main

import "fmt"

func main() {
	var num1, num2 float64
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num2)
	suma := num1 + num2
	fmt.Println("Suma de los dos números:", suma)

}
