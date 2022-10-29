package main

import "fmt"

func igv() {
	var precioVenta float64
	const porcentaje = 0.18

	fmt.Println("Ingrese el valor de la venta")
	fmt.Print("=> ")
	fmt.Scanln(&precioVenta)

	iGV := precioVenta * porcentaje
	venta := iGV + precioVenta

	fmt.Println("IGV:", iGV)
	fmt.Printf("Precio venta: %g \n", venta)

}

func main() {
	igv()
}
