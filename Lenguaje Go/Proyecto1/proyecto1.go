package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculadora de Suma")
	fmt.Println("-------------------")
	//Solicitar al usuario la lista de numeros
	fmt.Println("Ingresa una lista de números separados por comas:")
	input := leerEntrada()
	//Dividir la entrada números individuales
	numeros := strings.Split(input, ",")
	//Calcular la suma de los números
	suma := calcularSuma(numeros)
	//Mostrar el resultado
	fmt.Printf("La suma de los numeros ingresados es: %d\n", suma)

}
func leerEntrada() string {
	var input string
	fmt.Scanln(&input)
	return input
}
func calcularSuma(numeros []string) int {
	suma := 0
	for _, numStr := range numeros {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Printf("Advertencia: el valor '%s' no es número valido y será omitido.\n", numStr)
			continue
		}
		suma += num
	}
	return suma
}
