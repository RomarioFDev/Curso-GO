package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Se define la variable
	resultado := a + b
	fmt.Println("Suma: ", resultado)
	/**
	* Una vez se define la variable, podemos reutilizarla sin que afecten
	* los resultados anteriores
	* Solo que se tienen que mandar a imprimir junto a la operacion, sino, toma
	* el primer valor
	 */
	resultado = a - b
	fmt.Println("Resta: ", resultado)

	resultado = a * b
	fmt.Println("Multiplicacion: ", resultado)

	resultado = a / b
	fmt.Println("Divicion: ", resultado)

	/**
	* resultado = 3 / 2
	* Si mandamos esto, nunca nos saldra el decimal ya que como al principio
	* quedo inicializado como int solo mostrara valores enteros
	 */
	div := 3.0 / 2.0 // Especificar los decimales
	fmt.Println("Divicion decimal: ", div)

	resultado = a % b
	fmt.Println("Modulo: ", resultado)

}
