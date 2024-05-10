package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	persona := Persona{
		Nombre: "Juan",
		Edad:   30,
	}
	fmt.Println("Nombre:", persona.Nombre)
	fmt.Println("Edad:", persona.Edad)
}
