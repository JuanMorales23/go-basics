package main

import "fmt"

/*
GO solo maneja el for pero puede usarse de diversas maneras

>> El repetir para de siempre
for i := 0; i <= lim; i++ {
}

>> Se puede hacer también como un forEach cuando no se tiene un límite definido

*/

func bucles() {
	//Misma forma del for en lenguajes como java
	sum := 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("La suma de impares es: ", sum)
}
