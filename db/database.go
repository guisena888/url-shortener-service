package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error reading .env file")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	db, errSql := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})

	if errSql != nil {
		fmt.Println("Error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Sucessfully connected to database!")
	}

	db.AutoMigrate(&UrlMap{})
}

type UrlMap struct {
	Hash    string
	LongUrl string
}
