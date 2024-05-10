package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Juan"] = 25
	m["María"] = 30
	fmt.Println("Mapa:", m)
	fmt.Println("Edad de Juan:", m["Juan"])
	delete(m, "Juan")
	fmt.Println("Mapa después de eliminar a Juan:", m)
}
