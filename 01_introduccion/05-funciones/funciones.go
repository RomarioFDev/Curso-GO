package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

/**
* No es necesario definir las dos variables si son del mismo tipo
* Pero si hay que definir el tipo de retorno
 */
func sumar(a, b int) int {
	return a + b
}

func main() {
	nombre := "Romario"
	saludar(nombre)

	resultado := sumar(10, 5)
	fmt.Println("La suma es:", resultado)
}
