package main

import (
    "fmt"
	// "github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
    "backend/db_connection"
	// "net/http"
	// "time"
)

// var jwtKey = []byte("secret123")

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main(){
    // dbconnection.InitMySQL()

    uri := "mongodb://localhost:27017"

	err := dbconnection.ConnectMongoDB(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
}