package main

import "fmt"

type Animal interface {
	Sonido() string
}
type Perro struct{}

func (p Perro) Sonido() string {
	return "Guau,guau"
}

type Gato struct{}

func (g Gato) Sonido() string {
	return "Miau,miau"
}

func main() {
	perro := Perro{}
	gato := Gato{}
	hacerSonido(perro)
	hacerSonido(gato)
}

func hacerSonido(a Animal) {
	fmt.Println(a.Sonido())
}
