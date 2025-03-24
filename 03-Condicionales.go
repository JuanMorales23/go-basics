package main

import "fmt"

/*
>> Los condicionales en GO son los mismos que en los demás lenguajes
if, else, if else, swtich/case

>> Los operadores de comparación en GO son:
== operador de equivalencia (es igual a)
!= no es igual a
< menor que
> mayor que
<= menor o igal
>= mayor o igual

>> Se pueden hacer comparaciones como en javascript
bandera := number1 == 10 //true o false

# Pero no existen ternarios en GO

También está el switch/case
Seguiría la misma estructura que los demás lenguajes
*/
func condicionales() {
	var edad int = 20
	permiso := true
	if edad <= 18 {
		fmt.Println("No es mayor de edad y no puede ingresar")
	} else if edad <= 18 && permiso {
		fmt.Println("Es menor de edad pero tiene permiso de ingresar")
	} else {
		fmt.Println("Es mayor de edad")
	}

	var estadoCivil string
	print("Difite su estado civil")
	fmt.Scan(&estadoCivil)
	switch estadoCivil {
	case "Casado":
		fmt.Println("Cumple con los requisitos")
	case "Soltero":
		fmt.Println("Cumple con los requisitos")
	case "Unión libre":
		fmt.Println("No cumple con los requisitos")
	case "Divorciado":
		fmt.Println("No cumple con los requisitos, si tiene acta de divorcio")
	default:
		fmt.Println("No valido")
	}
}
