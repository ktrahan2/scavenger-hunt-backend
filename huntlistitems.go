package main

import (
	"github.com/jinzhu/gorm"
)

//HuntItem struct is my table for items
type HuntItem struct {
	gorm.Model
	Name  string `json:"name"`
	Image string `json:"image"`
	Theme string `json:"theme"`
}

func seedHuntItems() {
	var huntItems = []HuntItem{
		{
			Name:  "squirrel",
			Image: "",
			Theme: "nature",
		},
		{
			Name:  "acorn",
			Image: "https://media.istockphoto.com/photos/one-acorn-picture-id187333325?k=6&m=187333325&s=612x612&w=0&h=erzhZENyxgwPPKDiegjV6lNwDJFmP6iZUNNLDvmD1DI=",
			Theme: "nature",
		},
	}

	db.Create(&huntItems)

}

// func allItems(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	setupResponse(&w, r)

// 	var items []HuntItem
// 	var item HuntItem
// }
