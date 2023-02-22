package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB // Database connection pool.
//var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("host")
	DbPort := 5433
	DbUser := os.Getenv("user")
	DbPassword := os.Getenv("password")
	DbName := os.Getenv("dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)

	Db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		//fmt.Println("Cannot connect to database ", "postgres")
		log.Fatal("unable to use data source name", err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("We are connected to the database ", "postgres")
	//defer Db.Close()

	//DB.SetConnMaxLifetime(0)
	//DB.SetMaxIdleConns(3)
	//DB.SetMaxOpenConns(3)

}
