package main

import "fmt"

func main() {
	var consumo, descuento, porcentaje, datoDescuento float32

	// Estructura de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		fmt.Println("Es valido")
		if consumo <= 100 {
			//Descuento del 10%
			datoDescuento = 0.10 * 100
			porcentaje = 0.10
			descuento = consumo * porcentaje
		} else if consumo > 100 && consumo <= 200 {
			//Descuento del 20%
			datoDescuento = 0.20 * 100
			porcentaje = 0.20
			descuento = consumo * porcentaje
		} else if consumo > 200 {
			//Descuento del 30%
			datoDescuento = 0.30 * 100
			porcentaje = 0.30
			descuento = consumo * porcentaje
		}
	} else {
		fmt.Println("No es valido")
	}

	// Operaciones
	montoDescuento := consumo - descuento
	igv := montoDescuento * 0.19
	totalPagar := montoDescuento + igv

	fmt.Println("--------------------- FACTURA DE CONSUMO --------------------")
	fmt.Printf("Descuento que aplica %v%s \n", datoDescuento, "%")
	fmt.Printf("Consumo $%f \n", consumo)
	fmt.Printf("Descuento $%f \n", descuento)
	fmt.Printf("Monto con descuento $%f \n", montoDescuento)
	fmt.Printf("IGV %f%s \n", igv, "%")
	fmt.Printf("Total a pagar $%f \n", totalPagar)

}
