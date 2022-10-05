package main

import "fmt"

func incrementFail(number int) {
	number++
}

func increment(number *int) {
	*number++
}

func main() {
	number := 1
	fmt.Printf("Antes: %d\n", number)
	increment(&number)
	fmt.Printf("Depois: %d\n", number)

	number = 1
	fmt.Printf("Antes: %d\n", number)
	incrementFail(number)
	fmt.Printf("Depois: %d\n", number)
}
