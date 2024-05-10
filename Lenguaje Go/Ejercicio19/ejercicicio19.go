package main

import (
	"fmt"
	"time"
)

func imprimirMensaje(mensaje string, canal chan string) {
	time.Sleep(2 * time.Second) //Simula una tarea que lleva tiempo
	canal <- mensaje
}
func main() {
	canal := make(chan string)
	go imprimirMensaje("Hola", canal)
	go imprimirMensaje("¡Bienvenido!", canal)
	mensaje1 := <-canal
	mensaje2 := <-canal
	fmt.Println(mensaje1)
	fmt.Println(mensaje2)
}
