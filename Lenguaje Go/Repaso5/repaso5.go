package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}
	fmt.Println("La suma de los números es:", suma)
}
