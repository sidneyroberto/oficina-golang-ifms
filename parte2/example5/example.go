package main

import "fmt"

func main() {
	fmt.Println("Digite o valor da altura do retângulo: ")
	var height uint
	fmt.Scanf("%d", &height)

	fmt.Println("Digite o valor da base do retângulo: ")
	var basis uint
	fmt.Scanf("%d", &basis)

	area := basis * height
	fmt.Printf("Área: %d\n", area)
}
