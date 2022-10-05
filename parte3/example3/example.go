package main

import "fmt"

func main() {
	n := 0
	for n < 1 {
		fmt.Println("Digite a quantidade de elementos da PA: ")
		fmt.Scanf("%d", &n)
	}

	difference := 0
	for difference < 1 {
		fmt.Println("Digite a razão da PA: ")
		fmt.Scanf("%d", &difference)
	}

	element := 1
	count := 1
	for count <= n {
		end := ", "
		if count == n {
			end = ""
		}

		fmt.Printf("%d%s", element, end)
		count++
		element += difference
	}

	fmt.Println()
}
