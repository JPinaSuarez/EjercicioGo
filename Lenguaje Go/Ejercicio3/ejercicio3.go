package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num2)
	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	division := float64(num1) / float64(num2)
	fmt.Println("Suma", suma)
	fmt.Println("Resta", resta)
	fmt.Println("Multiplicacion", multiplicacion)
	fmt.Println("Division", division)
}
