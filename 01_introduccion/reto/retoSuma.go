package main

import "fmt"

// Funcion de suma
func suma(numero1, numero2 int) int {
	return numero1 + numero2
}

func main() {
	var numero1 int
	var numero2 int

	fmt.Println("Ingrese el numero 1:")
	fmt.Print("=> ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el numero 2:")
	fmt.Print("=> ")
	fmt.Scanln(&numero2)

	resultado := suma(numero1, numero2)
	fmt.Println("La suma es:", resultado)
}
