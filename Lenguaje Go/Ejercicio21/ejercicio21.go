package main

import "fmt"

type Rectangulo struct {
	base   float64
	altura float64
}

//Método de cálculo del área para el tipo Rectangulo
func (r Rectangulo) CalcularArea() float64 {
	return r.base * r.altura
}
func main() {
	rect := Rectangulo{base: 10, altura: 5}
	//Llamada al método CalcularArea()
	area := rect.CalcularArea()
	fmt.Println("Área del rectancgulo", area)
}
