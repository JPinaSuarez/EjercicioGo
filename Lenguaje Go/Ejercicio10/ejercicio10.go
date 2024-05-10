package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Ingresa el primer número:")
	fmt.Scan(&num1)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num2)
	fmt.Print("Ingresa el segundo número:")
	fmt.Scan(&num3)
	suma := num1 + num2 + num3
	resta := num1 - num2 - num3
	multiplicacion := num1 * num2 * num3
	division := num1 / (num2 * num3)
	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicación:", multiplicacion)
	fmt.Println("División:", division)
	if suma < 1 {
		fmt.Println("No pagas IVA")
	} else {
		fmt.Println("Pagas IVA", suma*float64(0.16))
	}
	for suma := 10; suma <= 25; suma++ {
		fmt.Println(suma)
	}
}
