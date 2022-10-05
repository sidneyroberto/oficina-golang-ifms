package main

import (
	"fmt"
	"oficina-golang-ifms/parte5/mymath"
)

func main() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := mymath.Sum(numbers)
	fmt.Printf("Result: %0.f\n", result)

	elements := mymath.ArithmeticProgression(10, 2)
	fmt.Printf("Elements: %v\n", elements)
}
