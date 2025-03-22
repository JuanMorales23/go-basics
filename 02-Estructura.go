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

>> Mapas tipo de arreglo que representa una colección de clave valor. Donde cada clave es única y el valor puede ser cualquier
tipo de dato, son conocidos en otros lenguajes como arreglos asociativos, diccionarios, hashmaps, etc...
var edades map[string]int = map[string]int{"Juan": 27, "Sarah", 27 }

>> Struct
Este tipo de datos representa un conjunto de campos con diferentes tipos de datos.
Las estructuras se utilizan para agrupar diferentes tipos de datos en una única variable.
// Declarar una estructura con diferentes campos
type Persona struct {
   nombre string
   edad int
   altura float64
}

var persona1 Persona = Persona{"Juan", 28, 1.75} //Son como las clases en Java, y esta sería la forma de instanciar

Se profundizará en los structs mas adelante...
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

	fmt.Println(numeros)
	fmt.Println(numeros2)

	capitales := map[string]string{
		"Colombia":  "Bogotá",
		"Chile":     "Santiago",
		"Venezuela": "Caracas",
	}
	//Podemos agregar más propiedades y su valor
	capitales["Argentina"] = "Buenos Aires"
	//Podemos quitar también elementos del mapa
	delete(capitales, "Venezuela")

	//Si se le pide un valor no definido va a mostrar un valor vacío
	fmt.Println(capitales)
	fmt.Println("Capital de colombia : ", capitales["Colombia"])
	fmt.Println("Agregando la capital de Argentina: ", capitales)
	fmt.Println("Quitando la capital de Venezuela: ", capitales)

}
