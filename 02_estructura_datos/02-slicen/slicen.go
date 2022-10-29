package main

import "fmt"

func main() {
	// Slicen
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	// Agregar elementos con append
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	// Sub Slicen
	subSlicen := numeros[0:2]
	/**
	* No importa cuando se modifique, el subslice obtendra el valor del
	* arreglo padre
	 */
	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

}
