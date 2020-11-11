package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User schema
type User struct {
	gorm.Model
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	HuntLists []HuntList `gorm:"many2many:user_lists;"`
}

// GetUsers selects * from users
func allUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	setupResponse(&w, r)

	var users []User
	var user User

	db.Preload("HuntLists.HuntItems").Find(&users)
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

// GetUser selects user by id
func getUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var user User

	db.Preload("HuntLists.HuntItems").Find(&user, key)

	json.NewEncoder(w).Encode(user)
}

//NewUser creates a new user
func newUser(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":

		reqBody, _ := ioutil.ReadAll(r.Body)

		var user User

		json.Unmarshal(reqBody, &user)

		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		user = User{
			Username: user.Username,
			Password: string(hash),
			Email:    user.Email,
		}

		db.Create(&user)

		db.Preload("HuntLists.HuntItems").First(&user)

		validToken, err := generateJWT()

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}

		Response := JWTTOKEN{
			validToken,
			user,
		}
		json.NewEncoder(w).Encode(Response)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}

// DeleteUser removes a user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var user User

	db.Table("users").Find(&user, key)

	db.Delete(&user)
}

//UpdateUser updates a user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var user User

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updateduser User
	json.Unmarshal(reqBody, &updateduser)

	db.Find(&user, key)

	db.Model(&user).Updates(User{
		Username:  updateduser.Username,
		Password:  updateduser.Password,
		Email:     updateduser.Email,
		HuntLists: updateduser.HuntLists,
	})
}
