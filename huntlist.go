package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//HuntList is the list created randomly and saved to a user
type HuntList struct {
	gorm.Model
	Title   string `json:"title"`
	OwnerID uint   `json:"ownerid"`
	Users   []User
}

//allHuntLists selects * from hunt_lists
func allHuntLists(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	setupResponse(&w, r)

	var huntLists []HuntList
	var huntlist HuntList

	db.Preload("Users").Find(&huntLists)
	huntLists = append(huntLists, huntlist)

	json.NewEncoder(w).Encode(huntLists)
}

//getHuntList selects a huntlist by ID
func getHuntList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var huntList HuntList

	db.Table("hunt_lists").Find(&huntList, key)

	json.NewEncoder(w).Encode(huntList)
}

//newHuntList creates a new hunt list
func newHuntList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var huntList HuntList
		json.Unmarshal(reqBody, &huntList)
		huntList = HuntList{
			Title:   huntList.Title,
			OwnerID: huntList.OwnerID,
			Users:   huntList.Users,
		}

		db.Create(&huntList)

		json.NewEncoder(w).Encode(&huntList.ID)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
