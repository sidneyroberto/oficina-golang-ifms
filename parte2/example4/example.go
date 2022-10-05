package main

import (
	"fmt"
	"math"
)

func main() {
	var height float64
	var weight float64

	fmt.Println("Digite a sua altura em cm: ")
	fmt.Scanf("%f", &height)
	fmt.Println("Digite o seu peso em kg: ")
	fmt.Scanf("%f", &weight)

	imc := weight / (math.Pow(height, 2))

	fmt.Printf("Seu IMC: %f\n", imc)
}
