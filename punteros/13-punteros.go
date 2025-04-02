package punteros

import "fmt"

//Para trabajar con punteros utilizamos el operador * y luego el tipo de dato de referencia
//EJ var ptr *int
//El operador & se utiliza para obtener la dirección de memoria de una variable
//EJ var ptr = &var1
//El operador * se utiliza para desreferenciar un puntero, es decir, acceder al valor almacenado en la dirección de memoria
//EJ var1 = *ptr

func Punteros() {
	var number int = 10
	var ptr *int = &number //Declaramos un puntero y le asignamos la dirección de memoria de la variable number
	*ptr = 20              //Desreferenciamos el puntero para asignar un nuevo valor a la variable number
	fmt.Println("La dirección de memoria de ptr es:", ptr)
	fmt.Println("El valor de ptr es:", ptr)     //Imprimimos la dirección de memoria almacenada en el puntero
	fmt.Println("El valor de number es:", *ptr) //Desreferenciamos el puntero para obtener el valor almacenado en la dirección de memoria
}
