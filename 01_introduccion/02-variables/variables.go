package main // Paquete principal

import "fmt" // Importamos el paquete fmt para poder imprimir en consola

func main() {
	// Las variables se declaran de la siguiente forma

	//Declaración de variables mas larga
	var nombre string
	nombre = "Juan"

	var nombre2 string = "Roberto" // Declaración con asignación

	edad := 20 // Declaración corta

	/**
	*Imprimir en consola
	*Podemos observar que ya no nos marca error porque ya se esta usando la variable
	 */
	fmt.Println(nombre, nombre2, edad)

	// Declaración de variables si definir un valor
	var nombre3 string
	var edad2 int
	var estado bool

	fmt.Println(nombre3, edad2, estado)

	// Valores constantes
	const pi = 3.14
	fmt.Println(pi)
}

/**
uint8 - 8-bits (0 a 255)
uint16 - 8-bits (0 a 65535)
uint32 - 32-bits (0 a 4294967295)
uint64 - 64-bits (0 a 18446744073709551615)

int8 - 8-bits (-128 a 127)
int16 - 16-bits (-32768 a 32767)
int32 - 32-bits (-2147483648 a 2147483647)
int64 - 64-bits (-9223372036854775808 a 9223372036854775807)

byte - 8-bits(0 a 255)
rune - 32-bits (-2147483648 a 2147483647)
uint - 32 0 64 bits
int - 32 0 64 bits
*/
