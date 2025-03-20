package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//  _ initialization only
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) Initialize(){
	db.connect()   
	defer db.close()
}

func (db *Database) connect(){
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Successfully connected to the database!")

	db.DB = connection
}


func (db * Database) close(){
	if db.DB !=nil{
		err := db.DB.Close()

		if err !=nil{
			log.Fatal("Error closing database connection:", err)
		}else{
			log.Println("Database connection closed.")
		}
	}
}