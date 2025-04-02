package main

import "fmt"

/*
>> Las funciones puede retornar o no
la forma de declarar los tipos de los parametros es la siguiente
luego de los paréntesis se debe declarar el tipo a retornar

	func nombreFunc(var tipo, var tipo, etc... ) tipoARetornar {
		return...
	}

>> Hay varios tipos de funciones
puede ser sin parámetros
func nombreFunc(){}

con parámetros
func nombreFunc(param1 tipo1, param2 tipo2, etc...){}

con retorno

	func nombreFunc(){
		return "Hola"
	}

sin retorno

	func nombreFunc(param1 tipo1, param2 tipo2){
		fmt.Println("Hola " + param1 + param2)
	}

>> La capitalización de la primera letra en el nombrado de las funciones es importante ya que
si la función empieza con mayúscula es una función pública y si empieza con minúscula es privada
de esta forma podemos controlar la visibilidad de las funciones
*/

func sumar(a int, b int) int {
	return a + b
}

// Se puede aplicar las funciones anónimas en GO
var resta = func(a int, b int) int {
	return a - b
}

func funciones() {
	fmt.Println("La suma es: ", sumar(5, 6))
	fmt.Println("La resta de ", resta(5, 4))
}

// Las funcione se pueden pasar también como parámetros
func operaciones(operacion func(...int) int, num1 int, num2 int) float32 {
	resultado := float32(operacion(num1, num2))
	return resultado
}

// func realizarOperacion() {
// 	operaciones := operaciones(multiplicarfunc(8), 5, 6)
// }

// Solo visible desde el archivo
func prueba() {
	fmt.Println("Hola desde la función prueba")
}

// Solo visible desde el paquete
func PruebaPublica() {
	fmt.Println("Hola desde la función prueba publica")
}

// >> Los métodos son una función con un argumento recepetor especial.
// El arguneto receptor es un tipo de dato que se le pasa a la función
// Con la ayuda de ese argugento receptor podemos acceder a los atributos de la estructura
// Puede ser de tipo estructura o de tipo puntero

// Ejemplo de método con argumento receptor
type Circle struct {
	Radius float64
}

// Para definir un método se utiliza la siguiente sintaxis
// lo que está entre parentesis es el argumento receptor
// Esto ata el método al tipo y podemos acceder a los campos de esa estructura
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
