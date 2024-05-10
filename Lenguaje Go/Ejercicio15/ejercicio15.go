package main

import "fmt"

func main() {
	edad := 25
	punteroEdad := &edad
	fmt.Println("Valor de edad:", edad)
	fmt.Println("Dirección de memoria:", punteroEdad)
	fmt.Println("Valor a traves del puntero:", *punteroEdad)
	*punteroEdad = 30
	fmt.Println("Nuevo valor de edad:", edad)
}
