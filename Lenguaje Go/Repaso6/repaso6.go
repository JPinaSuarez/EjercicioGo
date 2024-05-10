package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
}
type Rectangulo struct {
	Base   float64
	Altura float64
}
type Circulo struct {
	Radio float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}
func main() {
	rectangulo := Rectangulo{Base: 5, Altura: 10}
	circulo := Circulo{Radio: 3}
	mostrarArea(rectangulo)
	mostrarArea(circulo)
}
func mostrarArea(figura Figura) {
	fmt.Printf("El área de la figura es: %.2f\n", figura.Area())

}
