package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mudo!"

	fmt.Print(hola)
	fmt.Println(mundo)

	nombre := "Romario"
	edad := 22

	/**
	* Printf nos ayuda a formatear la informacion que recibe
	* %s recibe datos string
	* %d recibe datos int
	* %v se utiliza cuando no sabemos que tipo de dato recibe
	* %e, %f y %g son para float
	 */
	fmt.Printf("Hola %s tu edad es %d, te digo %s %s \n", nombre, edad, hola, mundo)

	fmt.Printf("Hola %v tu edad es %v, te digo %v %v \n", nombre, edad, hola, mundo)

	// Guardar el formateo en una variable utilizando el metodo sprintf y retorna la informacion formateada
	mensaje := fmt.Sprintf("Hola %s tu edad es %d", nombre, edad)
	fmt.Println(mensaje)

	/**
	* %T nos siver para saber que tipo de datos es cada variable
	 */
	fmt.Printf("nombre => %T \n", nombre)

	/**
	* utilizo Scanln para guardar el dato que se toma en consola y lo guardamos en la variable
	* que le asignamos pero se utiliza & + nombreVariable
	 */
	fmt.Println("Ingrese el otro nombre")
	fmt.Print("=> ")
	fmt.Scanln(&nombre)
	fmt.Println("Nombre de consola es:", nombre)

}
