package main

import (
	"fmt"
	"strings"
)

/*
GO solo maneja el for pero puede usarse de diversas maneras

>> El repetir para de siempre
for i := 0; i <= lim; i++ {
}

>> Se puede hacer también como un forEach cuando no se tiene un límite definido
for k, v := range capitales {
}

>> Para hacer un haga mientras
for { //For infinito
	if(condición para salir){
		break
	}
}
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

	//Podemos recorer mapas de la siguiente manera
	capitales := map[string]string{
		"Colombia":  "Bogotá",
		"Chile":     "Santiago",
		"Venezuela": "Caracas",
	}
	for k, v := range capitales {
		fmt.Println("La capital de " + k + " es: " + v)
	}

	//Para hacer un haga mientras (Do while)
	var fruta string
	for {
		fmt.Println("Ingrese una fruta: ")
		fmt.Scan(&fruta)
		fruta := strings.ToLower(fruta)
		if fruta == "naranja" {
			fmt.Println("No cumple")
			break
		}
	}
}
