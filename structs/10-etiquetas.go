package structs

func Etiquetas() {
	//Etiquetas
	//Son una forma de agregar información adicional a los campos de una estructura.
	//Se utilizan para proporcionar metadatos sobre los campos, como validaciones, descripciones, etc.
	//Las etiquetas se definen utilizando el símbolo ` y se colocan después del nombre del campo.

	type Persona struct {
		Nombre string `json:"nombre" xml:"nombre" db:"nombre"`
		Edad   int    `json:"edad" xml:"edad" db:"edad"`
	}
}
