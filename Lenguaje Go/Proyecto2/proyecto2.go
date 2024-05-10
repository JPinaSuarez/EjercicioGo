package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contador de Palabras")
	fmt.Println("--------------------")
	//Solicitar al usuario ingresar una frase
	fmt.Println("Ingrese una frase:")
	frase := leerEntrada()
	//Contar el número de palabras en la frase
	numPalabras := contarPalabras(frase)
	//Mostrar el resultado
	fmt.Printf("Número de palabras en la frase: %d\n", numPalabras)
}
func leerEntrada() string {
	var input string
	fmt.Scanln(&input)
	return input
}
func contarPalabras(frase string) int {
	//Eliminar los espacios en blanco adicionales
	fraselimpia := strings.TrimSpace(frase)
	//Dividir la frase en palabras utilizando los espacios en blanco como separadores
	palabras := strings.Split(fraselimpia, " ")
	//Contar el números de palabras
	numPalabras := len(palabras)
	return numPalabras
}
