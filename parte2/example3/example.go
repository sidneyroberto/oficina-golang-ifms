package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Digite o seu nome: ")
	fmt.Scanln(&name)
	fmt.Printf("Seja bem vindo, %s! Tenha um bom dia!", name)
}
