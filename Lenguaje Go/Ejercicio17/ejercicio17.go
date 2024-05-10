package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numeros)
	//Acceder a elementos individuales
	fmt.Println("Elemento 3:", numeros[2])
	//Agregar elementos al slide
	numeros = append(numeros, 6, 7)
	fmt.Println("Slice después de agregar elementos:", numeros)
	//Cortar un slice
	sliceCortado := numeros[1:4]
	fmt.Println("Slice cortado:", sliceCortado)
}
