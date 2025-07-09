package authentication


// backend authentication and database connection
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB sets up the connection to MySQL
func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3333)/web_go"
	
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }

    if err := DB.Ping(); err != nil {
        panic("Database unreachable: " + err.Error())
    }

    fmt.Println("✅ Connected to MySQL")

}