package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	//Arrange (Given) -- Preparación de los datos
	expectedResult := 5
	num1 := 2
	num2 := 3

	// Act (When) -- Afirmación o ejecución de la función
	result := Add(num1, num2)

	//Assert (Then)
	if expectedResult != result {
		t.Errorf("Se esperaba %d, pero se obtuvo %d", expectedResult, result)
	}
}

func TestSubtract(t *testing.T) {
	//Arrange (Given) -- Preparación de los datos
	expectedResult := 1
	num1 := 5
	num2 := 4

	// Act (When) -- Afirmación o ejecución de la función
	result := Subtract(num1, num2)

	//Assert (Then)
	if expectedResult != result {
		t.Errorf("Se esperaba %d, pero se obtuvo %d", expectedResult, result)
	}
}

func TestMultiply(t *testing.T) {
	//Arrange (Given) -- Preparación de los datos
	expectedResult := 1
	num1 := 3
	num2 := 5

	// Act (When) -- Afirmación o ejecución de la función
	result := Multiply(num1, num2)

	//Assert (Then)
	require.Equal(t, expectedResult, result, "Esperado %d, pero se obtuvo %d", expectedResult, result)
}
