package main

import "fmt"

func main() {
	// Definir arreglos
	var numeros [5]int
	numeros[0] = 1
	numeros[4] = 5

	fmt.Println(numeros)    // Imprime todo el arreglo
	fmt.Println(numeros[2]) // Imprime solo un indice en especifico

	// Array de nombres con limite de indices
	nombres := [3]string{"Romario", "Juan", "Luis"}
	// Array de colores sin logitud
	colores := [...]string{"Verde", "Rojo", "Azul"}

	fmt.Println(nombres)
	fmt.Println(colores, "Longitud", len(colores))

	// Definir indices aqui importa mucho el valor inicial ya que va en secuencia
	monedas := [...]string{
		0: "Dolar",
		1: "Peso Mexicano",
		2: "Peso Argentino",
		3: "Peso Colombiano",
		4: "Euro",
	}
	fmt.Println(monedas, len(monedas))

	submoneda := monedas[1:]
	fmt.Println(submoneda)
}
