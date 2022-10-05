package main

import "fmt"

func main() {
	var number1 int64
	var number2 uint64

	fmt.Println("Digite um número inteiro: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Digite um número positivo inteiro: ")
	fmt.Scanf("%d", &number2)

	result := number1 + int64(number2)
	fmt.Printf("%d + %d = %d\n", number1, number2, result)

}
