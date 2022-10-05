package main

import "fmt"

func sum(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func subtract(number1, number2 float64) float64 {
	return number1 - number2
}

func main() {
	result := sum(34, 45)
	fmt.Printf("Resultado: %f\n", result)

	result = subtract(34, 45)
	fmt.Printf("Resultado: %f\n", result)
}
