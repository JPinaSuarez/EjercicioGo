package main

import "fmt"

func main() {
	var num1, num3 int
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num3)
	suma := num1 + num3
	resta := num1 - num3
	multiplicacion := num1 * num3
	division := float64(num1) / float64(num3)
	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicacion:", multiplicacion)
	fmt.Println("Division:", division)
}
