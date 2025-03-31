package tests

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, string) {
	if b == 0 {
		return 0, "No se puede dividir por cero"
	}
	return a / b, ""
}
