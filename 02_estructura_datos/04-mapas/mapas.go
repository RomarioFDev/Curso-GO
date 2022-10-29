package main

import (
	"fmt"
)

func main() {

	/**
	* Esta es una forma mas entendible de representar el objeto
	* pero no se encuentra inicializado asi que mandaria error en caso
	* de querer agregar otro elemento
	 */
	var dias map[int]string = map[int]string{
		1: "Lunes",
		2: "Martes",
		3: "Miercoles",
		4: "Jueves",
		5: "Viernes",
	}
	fmt.Println(dias)
	// Metodo para eliminar un registro del mapa
	delete(dias, 1)

	fmt.Println(dias)
	//---------------------------------------------------------

	// Creacion de un mapa con una funcion Make()
	diaSem := make(map[int]string)

	// Agregamos datos al objeto
	diaSem[1] = "Lunes"
	diaSem[2] = "Martes"
	diaSem[3] = "Miercoles"
	diaSem[4] = "Jueves"
	diaSem[5] = "Viernes"

	fmt.Println(diaSem)

	// Modificar un elemento
	diaSem[4] = "JUEVES"

	fmt.Println(diaSem[4])

	// Eliminar un elemento
	delete(diaSem, 4)

	fmt.Println(diaSem, len(diaSem))

	// Inicializar un mapa con un array de elemetos
	estudiantes := make(map[string][]int)

	estudiantes["Romario"] = []int{1, 2, 3}
	estudiantes["Juean"] = []int{11, 12, 13}
	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Romario"])
}
