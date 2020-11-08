package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//SelectedItem is the selected items schema
type SelectedItem struct {
	gorm.Model
	HuntListID uint
	HuntItemID uint
}

//allHuntLists selects * from hunt_lists
func allSelectedItems(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	setupResponse(&w, r)

	var selectedItems []SelectedItem
	var selectedItem SelectedItem

	db.Find(&selectedItems)
	selectedItems = append(selectedItems, selectedItem)

	json.NewEncoder(w).Encode(selectedItems)
}

//getSelectedItems selects a selecteditem by ID
func getSelectedItem(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var selectedItem SelectedItem

	db.Find(&selectedItem, key)

	json.NewEncoder(w).Encode(selectedItem)
}

//newSelectedItem creates a new hunt list
func newSelectedItem(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var selectedItem SelectedItem
		json.Unmarshal(reqBody, &selectedItem)
		selectedItem = SelectedItem{
			HuntListID: selectedItem.HuntListID,
			HuntItemID: selectedItem.HuntItemID,
		}

		db.Create(&selectedItem)

		json.NewEncoder(w).Encode(&selectedItem)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
