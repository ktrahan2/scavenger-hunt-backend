package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

//User schema
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetUsers selects * from users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	setupResponse(&w, r)

	var users []User
	var user User

	db.Table("users").Find(&users)
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

//GetUser selects user by id
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	// setupResponse(&w, r)
// 	fmt.Fprintf(w, "get user endpoint hit")
// }

//NewUser creates a new user
func newUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		db.Create(&User{Username: username, Password: password, Email: email})
	default:
		http.Error(w, http.StatusText(405), 405)
	}
	//should return json to show how it works
	fmt.Fprintf(w, "New User successful")
}

//DeleteUser removes a user by id
// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// }

// //UpdateUser updates a user by id
// func UpdateUser(w http.ResponseWriter, r *http.Request) {

// }
