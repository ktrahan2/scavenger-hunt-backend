package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

//User Table
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

//InitialMigration connects to database and migrates table
func InitialMigration() {
	host := os.Getenv("DBHOST")
	databaseUsername := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	dbport := os.Getenv("DBPORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, databaseUsername, password, database)

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var users []User
	var user User

	db.Table("users").Find(&users)
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}
