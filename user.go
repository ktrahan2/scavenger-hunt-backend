package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

//User Table
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
