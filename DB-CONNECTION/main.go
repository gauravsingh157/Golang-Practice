package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)


var DB *sql.DB

func  ConnectDB() {
	var err error

	
	dsn := os.Getenv("DATABASE_URL")

	// Open PostgreSQL connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(" Failed to connect to database:", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal(" Database not responding:", err)
	}

		fmt.Println(" Successfully connected to PostgreSQL database!")
	}
func main(){
	fmt.Println("DataBase Connetion")
	ConnectDB().

}