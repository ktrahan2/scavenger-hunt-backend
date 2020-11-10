package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//UserList scehma
type UserList struct {
	gorm.Model
	HuntListID  uint
	UserID      uint
	CheckedItem pq.StringArray `gorm:"type:text[]"`
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

	db.Find(&userLists)
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

func getSpecificList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	userID := vars["userid"]
	huntListID := vars["huntlistid"]
	var userList []UserList

	db.Where("user_id = ? AND hunt_list_id = ?", userID, huntListID).Find(&userList)

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
			HuntListID:  userList.HuntListID,
			UserID:      userList.UserID,
			CheckedItem: userList.CheckedItem,
		}
		db.Create(&userList)

		json.NewEncoder(w).Encode(&userList)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}

//UpdateUserList updates a userlist by id
func updateUserList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var userList UserList

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedUserList UserList
	json.Unmarshal(reqBody, &updatedUserList)

	db.Find(&userList, key)

	db.Model(&userList).Updates(UserList{
		HuntListID:  updatedUserList.HuntListID,
		UserID:      updatedUserList.UserID,
		CheckedItem: updatedUserList.CheckedItem,
	})

	json.NewEncoder(w).Encode(&userList)
}

// deleteUserList removes a user by id
func deleteUserList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var userList UserList

	db.Table("user_lists").Find(&userList, key)

	db.Delete(&userList)
}
