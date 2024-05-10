package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Ingresa edad:")
	fmt.Scan(&num1)
	if num1 >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
