package structs

import "fmt"

// Las estructuras son tipos de datos compuestos que permiten agrupar diferentes tipos de datos en una sola variable.
// Son utilizadas para representar entidades del mundo real y sus atributos.
// Las estructuras son similares a las clases en otros lenguajes de programación, pero no tienen métodos ni herencia.
// Se definen utilizando la palabra clave "type" seguida del nombre de la estructura y los campos que contiene.
//
//	type Name struct {
//		field1 type1
//		field2 type2
//	}
func Structs() {
	//Podemos utilizar las estructuras como un tipo de dato dentro de otra estructura
	type Preferences struct {
		Foods  string
		Movies string
		Series string
		Animes string
		Sports string
	}
	type Persona struct {
		Nombre    string
		Apellido  string
		Edad      int
		Profesion string
		Empleado  bool
		Likes     Preferences
	}
	//Se pueden declarar de la siguiente manera
	p1 := Persona{"Esteban", "Gonzalez", 30, "Desarrollador", false, Preferences{"Pizza", "Star Wars", "Breaking Bad", "Naruto", "Futbol"}}
	p2 := Persona{
		Nombre:    "Juan",
		Apellido:  "Morales",
		Edad:      28,
		Profesion: "Ingeniero",
		Empleado:  true,
		Likes:     Preferences{"Hamburguesa", "Avengers", "Game of Thrones", "One Piece", "Basketball"},
	}
	fmt.Println("Nombre:", p2)
	//También se puede ir llenando los campos de la estructura
	var p3 Persona
	p3.Nombre = "Pedro"
	p3.Apellido = "Lopez"
	p3.Edad = 35
	p3.Profesion = "Arquitecto"
	p3.Empleado = true
	//Se puede agregar también de esta forma o agregando el struct completamente
	// p3.Likes.Foods = "Sushi"
	// p3.Likes.Movies = "Inception"
	p3.Likes = Preferences{"Sushi", "Inception", "Breaking Bad", "Naruto", "Futbol"}
	//Para imprimir el valor de un campo de la estructura se utiliza como si fuera un objetos
	fmt.Println("Nombre:", p1.Nombre)

}
