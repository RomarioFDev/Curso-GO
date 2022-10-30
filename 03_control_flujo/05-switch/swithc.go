package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)
	// Forma uno de trabajar con switch
	/* 	switch {
	   	case vocal == "a":
	   		fmt.Println(vocal, "es vocal")
	   	default:
	   		fmt.Println("No existe")
	   	} */

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "es una vocal")
	default:
		fmt.Println(vocal, "es una consonante")
	}
}
