package main

import (
	"backend/services/user_services/internal/db"
	"fmt"
	"backend/services/user_services/internal/repository"
)

func main(){

	dbConn := db.Connect(db.LoadConfig())

	// Create repository instance
	repo := repository.NewUserRepository(dbConn)

	user, err := repo.GetUserByEmail("alice@example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("User found:", user)

}
