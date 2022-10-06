package main

import (
	"fmt"
	"log"
	"oficina-golang-ifms/parte6/controllers"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Erro ao tentar pegar o caminho do arquivo users.csv")
	}
	path := filepath.Join(filename, "../data/users.csv")
	//users := controllers.ReadUsers(path)
	users := controllers.ReadUsers2(path)
	for _, user := range users {
		fmt.Printf("User: %v\n", user)
	}
}
