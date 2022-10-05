package main

import (
	"fmt"
	"log"
)

func main() {
	var number1 int64
	var number2 uint64

	fmt.Println("Digite um número inteiro: ")
	n, err := fmt.Scanf("%d", &number1)
	if err != nil {
		log.Fatal("Erro ao ler valor: ", err)
	}
	fmt.Printf("%d valores lidos\n", n)

	fmt.Println("Digite um número positivo inteiro: ")
	_, err = fmt.Scanf("%d", &number2)
	if err != nil {
		log.Fatal("Erro ao ler valor: ", err)
	}

	result := number1 + int64(number2)
	fmt.Printf("%d + %d = %d\n", number1, number2, result)

	result = number1 - int64(number2)
	fmt.Printf("%d - %d = %d\n", number1, number2, result)

	result = number1 * int64(number2)
	fmt.Printf("%d * %d = %d\n", number1, number2, result)

	result = number1 / int64(number2)
	fmt.Printf("%d / %d = %d\n", number1, number2, result)

	result = number1 % int64(number2)
	fmt.Printf("%d %% %d = %d\n", number1, number2, result)
}
