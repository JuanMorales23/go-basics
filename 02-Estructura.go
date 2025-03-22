/*
Tipos de datos y gestión de variables, es estricto en los tipos de datos pero nos permite inferir facilmente el tipo

>> Tipos de datos básicos:
Numericos (int, float32 y float64)
Cadenas de texto (string)
Booleanos (bool)

>> Tipos de datos compuestos:
array: representa una colección de elementos del mismo tipo [Tamaño fijo]
slice: es similar a un array, pero su tamaño puede ser modificado dinámicamente
map: representa una colección de pares clave-valor
struct: representa un conjunto de campos con diferentes tipos de datos

>> Tipos de datos de referencia:
pointer: representa la dirección de memoria de una variable.
function: representa una función.
interface: define un conjunto de métodos que una estructura debe implementar.

Para crear variables se puede hacer de 3 formas

1. Con la palabra var y de forma explicita
var variable1 string = "Hola que tal"

2. Con la palabra var y de forma implicita
var variable2 = "Mi nombre es"

3. Sin la palabra var y por asignación usando :=
var variable3 := 27


Para el caso de los números tenemos por ejemplo los int la capacidad de especificar la cantidad de bits que tendra
int8
int16
int32, etc...

>> Arreglos
var ListaFrutas = [4] string {"Pera", "Naranja"} //Sin valor
listaPaises := [3] int

>> Slice (Son como arreglos pero sin limite definido de datos), puede ser modificado dinámicamente
var numeros []int = []int {datos}
se utilizan los métodos para agregar datos, por ejemplo el append
numeros = append(numeros, 6)
*/

package main

import "fmt"

func func02() {
	var saludo string = "Hola que tal "
	var nombre = "mi nombre es Juan"
	apellido := "Morales"

	fmt.Println(saludo, nombre, apellido)

	var anioActual int16 = 2025
	var anioReducido = 24
	nacimiento := 1997

	fmt.Println("Variables numerica => anio = ", anioActual, ", anio reducido = ", anioReducido, ", fecha nacimiento = ", nacimiento)

	var ListaFrutas = [5]string{"Pera", "Manzana", "Durazno"}
	listaPaises := [4]string{}
	listaPaises[0] = "Colombia"
	listaPaises[1] = "Ecuador"

	fmt.Println(ListaFrutas)
	fmt.Println(listaPaises)

	var numeros []int = []int{1, 2, 3, 4, 5}
	//Podemos crear slices desde un rango dado, por ejemplo desde la posición 2 para crear un nuevo arreglo o slice
	//Tener en cuenta que el rango incluye el inicial pero no el final
	//Rango definido [n:m]
	//Rango no definido [1:]
	//Rango inverso [:5]
	numeros2 := numeros[2:5]

}
