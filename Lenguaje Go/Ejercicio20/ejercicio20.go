package main

import "fmt"

//Definicion de interfaz
type Animal interface {
	Sonido() string
}

//Implementación para el tipo Perro
type Perro struct{}

func (p Perro) Sonido() string {
	return "Guau guau"
}

//Implementación para el tipo Gato
type Gato struct{}

func (g Gato) Sonido() string {
	return "Miau miau"
}
func main() {
	perro := Perro{}
	gato := Gato{}

	//Llamada a la función que utiliza la interfaz
	hacerSonido(perro)
	hacerSonido(gato)
}

//Función que utiliza la interfaz
func hacerSonido(a Animal) {
	fmt.Println(a.Sonido())
}
