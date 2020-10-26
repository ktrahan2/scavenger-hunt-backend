package models

import "github.com/jinzhu/gorm"

//User Table
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
