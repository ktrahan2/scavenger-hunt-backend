package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//UserList is the struct of the user lists table
type UserList struct {
	gorm.Model
	HuntListID uint `json:"huntlistid"`
	UserID     uint `json:"userid"`
}

//allUserLists is the index for user lists
func allUserLists(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	setupResponse(&w, r)

	var userLists []UserList
	var userList UserList

	db.Table("user_lists").Find(&userLists)
	userLists = append(userLists, userList)

	json.NewEncoder(w).Encode(userLists)
}

//getUserList selects a Userlist by ID
func getUserList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var userList UserList

	db.Table("user_lists").Find(&userList, key)

	json.NewEncoder(w).Encode(userList)
}

//newUserList creates a new users list
func newUserList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var userList UserList
		json.Unmarshal(reqBody, &userList)
		userList = UserList{
			HuntListID: userList.HuntListID,
			UserID:     userList.UserID,
		}
		var user User
		db.Table("users").Find(&user, userList.UserID)
		var huntList HuntList
		db.Table("hunt_lists").Find(&huntList, userList.HuntListID)
		db.Model(&user).Updates(User{
			HuntLists: []HuntList{
				{&huntList},
			},
		})

		db.Create(&userList)
		// db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&User)

		json.NewEncoder(w).Encode(&userList)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
