package main

import (
	"fmt"
	"strings" // Package donde obtenemos el metodo ToLower
)

func reverse(cadena string) string {

	/**
	* Le quito los espacios a cadena para que no tome ese espacio como un
	* dato mas en el arreglo
	 */
	cadena = strings.Replace(cadena, " ", "", 1)

	/**
	* Split convierte un string en un arreglo y si tiene espacios vacios
	* le indicamos que lo tome como otro elemento y no los meta en un solo
	* espacio
	 */
	arrayCadena := strings.Split(cadena, "")
	// fmt.Println(arrayCadena)

	// Imprimo el arreglo mediante un for y lo itero
	/* 	for i := 0; i < len(arrayCadena); i++ {
		fmt.Println(arrayCadena[i])
	} */
	arrayCadenaInvertida := make([]string, 0)

	for j := len(arrayCadena) - 1; j >= 0; j-- {
		arrayCadenaInvertida = append(arrayCadenaInvertida, arrayCadena[j])
	}
	//fmt.Println(arrayCadenaInvertida)

	// Join une las letras del arreglo y las separa por comas o como queramos
	return strings.Join(arrayCadenaInvertida, ",")
}

func esPalindromo(palabra string) bool {
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra) // Convertir la palabra en minuscula
	palabra = strings.ReplaceAll(palabra, " ", "")
	fmt.Println(palabra)

	/* palabra = strings.Replace(palabra, " ", "", 2) // Metodo para remplazar palabras o letras
	fmt.Println(palabra) */

	// Llamamos al metodo revesa y le pasamos palabra como parametro
	// fmt.Println(reverse(palabra))

	palabraInvertida := reverse(palabra)

	return palabra == palabraInvertida
}

func main() {
	// Detectar si una palabra es palindromo
	//esPalindromo("LUZ AZUL")
	// reverse("Luz Azul")

	if esPalindromo("Luz Azul") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
