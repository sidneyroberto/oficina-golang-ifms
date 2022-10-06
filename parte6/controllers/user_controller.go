package controllers

import (
	"encoding/csv"
	"log"
	"oficina-golang-ifms/parte6/models"
	"os"
)

func ReadUsers(filePath string) []models.User {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for _, line := range csvLines[1:] {
		user := models.User{
			Name:  line[0],
			Email: line[1],
			Phone: line[2],
		}

		users = append(users, user)
	}

	return users
}

func ReadUsers2(filePath string) []map[string]string {
	users := ReadUsers(filePath)

	var mappedUsers []map[string]string
	for _, user := range users {
		u := map[string]string{
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
		}

		mappedUsers = append(mappedUsers, u)
	}

	return mappedUsers
}
