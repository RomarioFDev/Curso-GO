package main

import "fmt"

func main() {

	var nombre string
	fmt.Print("Que usuario quieres => ")
	fmt.Scanln(&nombre)

	// if nombre, edad := "Romario", 22; nombre == "Romario" {
	// 	fmt.Printf("Hola %s con edad %d \n", nombre, edad)
	// } else {
	// 	fmt.Println("Error")
	// }

	usuarios := make(map[string]string)
	usuarios["Romario"] = "romario@gmail.com"
	usuarios["Juan"] = "juan@gmail.com"

	if correo, ok := usuarios[nombre]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("El usuario no existe")
	}
}
