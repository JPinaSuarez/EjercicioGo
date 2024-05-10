package main

import (
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Open("archivo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer archivo.Close()
	//Resto del código para trabajar con el archivo abierto
	fmt.Println("Archivo abierto correctamente.")

}
