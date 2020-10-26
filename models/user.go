package models

import "github.com/jinzhu/gorm"

//User Table
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//Users is an exported type that contains the instance of a User
type Users []User
