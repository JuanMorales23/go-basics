// Todo archivo de go debe tener una función main e importart el paquete principal de go
// No se utiliza el ;
package main

import "fmt"

func Func01() {
	fmt.Println("¡Hello World with GO!")
}

//Para ejecutar el programa hay que hacer un build del mismo
//go build <Nombre programa>.go

//Para correr el programa de una vez en vez de estar haciendo build siempre debemos hacer un run
//go run <nombre programa>.go
