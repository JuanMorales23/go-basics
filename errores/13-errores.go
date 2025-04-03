package errores

import (
	"errors"
	"fmt"
)

//El manejo de errores en GO se representan a traves de una interfaz nativa llamada error
// que contiene un unico metodo Error() que retorna un string

//Go nos proporciona dos packages que simplifica la creación de errores.
// 1. errors, que nos permite crear un error a partr de una cadena de texto
// Ej: err := errors.New("Error de prueba")
// 2. fmt
// Ej: err := fmt.Errorf("Error de prueba %d", 10)

//El package error nos proporciona funciones utiles como New, Unwrap, Is y As
// New: nos permite crear un nuevo error a partir de un string
// Toma una variable de tipo string y devuelve una variable de tipo error conteniendo dicha cadena de texto como mensaje de error

// Unwrap: nos permite obtener el error original de un error envuelto
// Nos permite desevolver los errores dentro de la cadena err. Toma el último error de la cadena y lo devuelve como un error

// Is: nos permite verificar si un error es de un tipo determinado
// Recibe como argumentos ds variables de tipo error, estas son err y target. Si se encuentra la conicidencia dentro de la cadena
// de errores de er con el error target, devuelve true. En caso contrario devuelve false
// Esta función la utilizamos cuando necesitamos saber si dentro de la cadena de errores, ocurrió uno exactamente igual al que estamos buscando (target)
// Ej: errors.Is(err, os.ErrNotExist)

// As: nos permite obtener el error de un tipo determinado
// Recibe 2 argumentos, uina variable de tipo error y un tarjet (interfaz vacía) que nos permite pasarle un puntero a un tipo error
// y retorna una variable bool
// Lo que hace es comparar dentro de la cadena de errores de err, si el error pasado como primer argunmento es del tipo del segundo argumento
// Ej: errors.As(err, &os.PathError{})

// Manejo de errores
// Una particularidad de GO es que las funciones pueden ser multi-retorno
// aprovechamos esta caracteristica cuando necesitamos trabajar con una función potencial de fallo
// func() (result interface{}, err error) {
//     ....
//}

var ErrDivisionPorCero = errors.New("division por cero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionPorCero
	}
	return a / b, nil
}

func errores() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Resultado:", result)
}
