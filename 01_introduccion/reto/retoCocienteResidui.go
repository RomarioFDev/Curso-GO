package main

import "fmt"

func cociente(numero1, numero2 int) int {
	return numero1 / numero2
}
func residuo(numero1, numero2 int) int {
	return numero1 % numero2
}

func main() {
	var numero1, numero2 int

	fmt.Println("Ingrese el numero 1")
	fmt.Print("=> ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el numero 2")
	fmt.Print("=> ")
	fmt.Scanln(&numero2)

	cociente := cociente(numero1, numero2)
	reciduo := residuo(numero1, numero2)

	fmt.Println("El cociente es:", cociente)
	fmt.Println("El residuo es:", reciduo)
}
