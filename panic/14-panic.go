package panic

import (
	"fmt"
	"os"
)

// Los panics son errores que no se pueden manejar y necesitan tratamientoe especial
// Un panic es una forma de detener la ejecución normal de un programa y señala un error critico e inesperado
// Esto para la ejecución del programa

func panics() {
	//Se llama de la siguiente manera
	//panic("Error inesperado")

	fmt.Println("Abrir un archivo que no existe")
	_, err := os.Open(("archivo.txt"))
	if err != nil {
		panic(err)
	}
	//Por el panic se temrina la ejecución del programa y no sigue con el código
	fmt.Println("Esto no se ejecuta")

	//Index out of bounds Panics
	//Podemos provocar un panic al acceder a un índice fuera de los límites de un slice o array
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[10]) // Esto provocará un panic: runtime error: index out of range

	//para manejar el panic se utilizan funciones defer y recover
	//defer se utiliza para ejecutar una función justo antes de que la función que la contiene retorne
	//es una sentencia que nos permite diferir la ejecución de ciertas funciones y asesugrar que sean ejecutadas antes de la finalización
	//
	//recover se utiliza para recuperar el control de un panic
}
