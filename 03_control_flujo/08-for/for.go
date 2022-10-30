package main

import "fmt"

func main() {

	// Bucle infinito
	/* 	for {
		fmt.Println("*")
	} */

	// Bucle tipo while
	numeros := 12455
	contador := 0
	for numeros > 0 {
		numeros /= 10
		contador++
	}
	fmt.Println("Cantidad de digitos:", contador)

	// Bucle for
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
