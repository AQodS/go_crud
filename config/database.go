package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	godotenv.Load()
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", user+":"+password+"@/go_products?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}
