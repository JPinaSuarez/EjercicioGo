package main

import (
	"fmt"
	"time"
)

type Rectangulo struct {
	base   float64
	altura float64
}

func imprimirMensaje(mensaje string, canal chan string) {
	time.Sleep(5 * time.Second)
	canal <- mensaje
}

func (r Rectangulo) CalcularArea() float64 {
	return r.base * r.altura
}
func main() {
	canal := make(chan string)
	go imprimirMensaje("Se esta iniciando la calculadora...", canal)
	fmt.Println(<-canal)
	var num1, num2 float64
	fmt.Print("Ingresa la base del triangulo:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa la altura del triangulo:")
	fmt.Scan(&num2)
	Perimetro := float64(4) * (num1)
	rect := Rectangulo{base: num1, altura: num2}
	area := rect.CalcularArea()
	fmt.Println("Area del rectangulo:", area)
	fmt.Println("Perimetro del triangulo:", Perimetro)
	if area < 10 {
		fmt.Println("Triangulo regular")
	} else {
		fmt.Println("Triangulo irregular")
	}
}
