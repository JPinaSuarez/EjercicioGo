package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num2)
	fmt.Print("Ingresa el tercer número:")
	fmt.Scan(&num3)
	suma := num1 + num2 + num3
	if suma >= 50 {
		fmt.Println("Pagas Iva", suma*float64(0.16))
	} else {
		fmt.Println("No pagas Iva")
	}

}
