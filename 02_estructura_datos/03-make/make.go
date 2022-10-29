package main

import "fmt"

func main() {
	// Funcion make
	numeros := make([]int, 3, 3)
	fmt.Println(numeros, len(numeros), cap(numeros))
	numeros[0] = 2
	numeros[1] = 3
	numeros[2] = 4

	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)
	fmt.Println(numeros, cap(numeros))
}
