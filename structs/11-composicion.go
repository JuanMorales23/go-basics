package structs

// La composición es una forma de crear estructuras complejas a partir de otras estructuras más simples.
// Se utiliza para representar relaciones "tiene un" entre diferentes entidades.
// Es similar a la herencia en otros lenguajes de programación, pero en lugar de heredar propiedades y métodos,
// se combinan diferentes estructuras para crear una nueva.
type Motor struct {
	Type       string
	HorsePower int
}

type Chasis struct {
	Material string
}

type Car struct {
	Motor
	Chasis
	Color string
}

// Se conoce también como Embedding structs, ya que se puede utilizar una estructura dentro de otra
func Composicion() {
	c := Car{
		Motor:  Motor{},
		Chasis: Chasis{},
		Color:  "Rojo",
	}

	c.Motor.HorsePower = 200
	c.Motor.Type = "V8"
	c.Chasis.Material = "Aluminio"
}
