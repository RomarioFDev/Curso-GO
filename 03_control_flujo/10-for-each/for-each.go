package main

import "fmt"

func main() {
	nombres := []string{"Romario", "Juan", "Pedro"}

	// For normal
	/* 	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	} */

	// ForEach
	/**
	* Si no quiero imprimir el indice y solo quiero enlistar los nombres colocamos
	* un _ al principio del for, pero si queremos que se vean los indices, colocamos
	* una variable llamada indice o como la queramos identificar
	 */
	/* 	for indice, elementos := range nombres {
		fmt.Println(indice, elementos)
	} */
	for _, elementos := range nombres {
		fmt.Println(elementos)
	}
}
