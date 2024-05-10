package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir entre cero")

	}
	return dividendo / divisor, nil
}
func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El resultado de la división es:", resultado)
	}
	resultado, err = dividir(8, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El resultado de la división es:", resultado)
	}
}
