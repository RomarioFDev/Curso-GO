package main

import "fmt"

func main() {
	a := 4
	b := 5

	var resultado bool

	// Operadores relacionales de igualdad
	resultado = a == b
	fmt.Printf("%d es igual que %d ?\n=> %t \n", a, b, resultado)

	// Operadores relacionales de desigualdad
	resultado = a != b
	fmt.Printf("%d es diferente que %d ?\n=> %t \n", a, b, resultado)

	// Operadore relacional de mayor que
	resultado = a > b
	fmt.Printf("%d es mayor que %d ?\n=> %t \n", a, b, resultado)

	// Operadore relacional de menor que
	resultado = a < b
	fmt.Printf("%d es menor que %d ?\n=> %t \n", a, b, resultado)

	// Operadore relacional de mayor o igual que
	resultado = a >= b
	fmt.Printf("%d es mayor o igual que %d ?\n=> %t \n", a, b, resultado)

	// Operadore relacional de menor o igual que
	resultado = a <= b
	fmt.Printf("%d es menor o igual que %d ?\n=> %t \n", a, b, resultado)
}
