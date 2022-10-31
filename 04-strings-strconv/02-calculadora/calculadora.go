package main

import (
	"fmt"
	"strconv"
	"strings"
)

func suma(expresion string) int {
	// Separo el string y los guardo en un arreglo
	valores := strings.Split(expresion, "+")
	var resultados int
	// Itero el arreglo en un foreach
	for _, valor := range valores {
		// numero sera igual al valor convertido de los iterandos
		// Y le damos una variable error ya que ahi se guardara la exception
		numero, error := strconv.Atoi(valor) // Se convierten los valores del arreglo en Int
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error tiene que ingresar un numero entero")
		} else {
			resultados += numero // Se sumando todos los valores del arreglo
		}
	}
	return resultados
}
func restar(expresion string) int {
	// Separo el string y los guardo en un arreglo
	valores := strings.Split(expresion, "-")
	var resultados int
	// Itero el arreglo en un foreach
	for _, valor := range valores {
		// numero sera igual al valor convertido de los iterandos
		// Y le damos una variable error ya que ahi se guardara la exception
		numero, error := strconv.Atoi(valor) // Se convierten los valores del arreglo en Int
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error tiene que ingresar un numero entero")
		} else {
			resultados -= numero // Se sumando todos los valores del arreglo
		}
	}
	return resultados
}
func multiplicacion(expresion string) int {
	// Separo el string y los guardo en un arreglo
	valores := strings.Split(expresion, "*")
	var resultados int
	// Itero el arreglo en un foreach
	for _, valor := range valores {
		// numero sera igual al valor convertido de los iterandos
		// Y le damos una variable error ya que ahi se guardara la exception
		numero, error := strconv.Atoi(valor) // Se convierten los valores del arreglo en Int
		if error != nil {
			fmt.Println(error)
			fmt.Println("Error tiene que ingresar un numero entero")
		} else {
			resultados *= numero // Se sumando todos los valores del arreglo
		}
	}
	return resultados
}

func main() {
	var expresion string
	var resultado int
	var menu int
	fmt.Println("================== MENU ==================")
	fmt.Print("1.- Suma\n2.- Resta\n3.- Multiplicacion\n")
	fmt.Print("=> ")
	fmt.Scanln(&menu)

	fmt.Print("Valores\n")
	fmt.Print("=> ")
	fmt.Scanln(&expresion)
	switch menu {
	case 1:
		resultado = suma(expresion)
		fmt.Printf("=> %d\n", resultado)
	case 2:
		resultado = restar(expresion)
		fmt.Printf("=> %d\n", resultado)
	case 3:
		resultado = multiplicacion(expresion)
		fmt.Printf("=> %d\n", resultado)
	}
}
