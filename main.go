package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dataBaseConnection()
	handleRequest()
	defer db.Close()
}

//dataBaseConnection connects to database and migrates table
func dataBaseConnection() {

	//use psql info in order to use local database
	// host := os.Getenv("DBHOST")
	// databaseUsername := os.Getenv("USERNAME")
	// password := os.Getenv("PASSWORD")
	// database := os.Getenv("DATABASE")
	// dbport := os.Getenv("DBPORT")

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, dbport, databaseUsername, password, database)

	databaseURL := os.Getenv("DATABASE_URL")
	db, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &HuntItem{})
	// seedHuntItems()

	fmt.Println("Successfully connected!")
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")
}
