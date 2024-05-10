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
	time.Sleep(2 * time.Second)
	canal <- mensaje
}

func (r Rectangulo) CalcularArea() float64 {
	return r.base * r.altura
}

func main() {
	canal := make(chan string)
	go imprimirMensaje("Se está iniciando la calculadora...", canal)
	fmt.Println(<-canal)

	fmt.Print("Ingresa la base del rectángulo: ")
	var base float64
	fmt.Scan(&base)
	fmt.Print("Ingresa la altura del rectángulo: ")
	var altura float64
	fmt.Scan(&altura)

	rect := Rectangulo{base: base, altura: altura}
	area := rect.CalcularArea()
	fmt.Println("Área del rectángulo:", area)

	if area < 10 {
		fmt.Println("Rectángulo regular")
	} else {
		fmt.Println("Rectángulo irregular")
	}
}
