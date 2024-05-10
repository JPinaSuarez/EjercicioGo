package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num1, num2, num3 int
	numeroAleatorio := rand.Intn(100)
	fmt.Println("Ingresa el primer Número:")
	fmt.Scan(&num1)
	fmt.Println("Ingresa el segundo Número:")
	fmt.Scan(&num2)
	fmt.Println("Ingresa el tercer Número:")
	fmt.Scan(&num3)
	suma := num1 + num2 + num3 + numeroAleatorio
	resta := num1 - num2 - num3 - numeroAleatorio
	multiplicacion := num1 * num2 * num3 * numeroAleatorio
	division := num1 / (num2 + num3 + numeroAleatorio)
	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicación:", multiplicacion)
	fmt.Println("División:", division)
	if suma >= 10 {
		fmt.Println("Pagas el IVA:", suma*(2))
	} else {
		fmt.Println("No pagas el IVA")
	}
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
