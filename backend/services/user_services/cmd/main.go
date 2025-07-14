package main

import (
	"backend/services/user_services/internal/db"
	"fmt"
)

func main(){

	dbConn := db.Connect(db.LoadConfig())

	fmt.Println(dbConn)


}
