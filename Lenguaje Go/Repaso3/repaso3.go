package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num2)
	suma := num1 + num2
	fmt.Println("La suma de los dos números es:", suma)
	if suma%2 == 0 {
		fmt.Println("El número es par.")
	} else {
		fmt.Println("El número es impar.")
	}

}
