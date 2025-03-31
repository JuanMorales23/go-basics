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
func operaciones(operacion func, num1 int, num2 int) float32 {
	Prueba()
	return operacion(num1, num2)

}

// Solo visible desde el archivo
func Prueba() {

}
