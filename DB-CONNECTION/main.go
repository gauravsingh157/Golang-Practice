package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a global database connection variable
var DB *sql.DB

// ConnectDB initializes the PostgreSQL connection
func  ConnectDB() {
	var err error

	// Get database URL from environment variables
	dsn := os.Getenv("DATABASE_URL") // Example: "postgres://user:password@localhost:5432/dbname?sslmode=disable"

	// Open PostgreSQL connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database not responding:", err)
	}

		fmt.Println("✅ Successfully connected to PostgreSQL database!")
	}
func main(){
	fmt.Println("DataBase Connetion")
}