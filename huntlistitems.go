package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//HuntItem struct is my table for items
type HuntItem struct {
	gorm.Model
	Name  string `json:"name"`
	Image string `json:"image"`
	Theme string `json:"theme"`
}

//allItems gets all hunt items
func allItems(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	setupResponse(&w, r)

	var items []HuntItem
	var item HuntItem

	db.Table("hunt_items").Find(&items)
	items = append(items, item)

	json.NewEncoder(w).Encode(items)
}

//getItem grabs a hunt item by id
func getItem(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var huntItem HuntItem

	db.Table("hunt_items").Find(&huntItem, key)

	json.NewEncoder(w).Encode(huntItem)
}

func updateHuntItem(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	key := vars["id"]
	var huntItem HuntItem

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedItem HuntItem
	json.Unmarshal(reqBody, &updatedItem)

	db.Table("&hunt_items").Find(&huntItem, key)

	db.Model(&huntItem).Updates(HuntItem{
		Name:  updatedItem.Name,
		Image: updatedItem.Image,
		Theme: updatedItem.Theme,
	})

}

// func seedHuntItems() {
// 	huntItem := HuntItem{
// 		Name:  "squirrel",
// 		Image: "https://www.charlotteobserver.com/latest-news/h553hn/picture235468597/alternates/FREE_1140/squirrelmcclatchy.JPG",
// 		Theme: "nature",
// 	}

// 	huntItem2 := HuntItem{
// 		Name:  "acorn",
// 		Image: "https://media.istockphoto.com/photos/one-acorn-picture-id187333325?k=6&m=187333325&s=612x612&w=0&h=erzhZENyxgwPPKDiegjV6lNwDJFmP6iZUNNLDvmD1DI=",
// 		Theme: "nature",
// 	}
// 	huntItem3 := HuntItem{
// 		Name:  "tree",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTW0BCgi1iKHL7gbpxPO_JbLs58AV5NbZLuYg&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem4 := HuntItem{
// 		Name:  "wildflower",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcT8pb26eUzg6hdJ9VFZi2ANSlNt-8WTUPDCpg&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem5 := HuntItem{
// 		Name:  "pinecone",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSmZvvVA7oSpA8_BEBmSRA2KmRJbzkmDbCgWg&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem6 := HuntItem{
// 		Name:  "butterfly",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcR8Fw4YDrb9c8CcHcItePExFfpZnJ5ZDpcx3Q&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem7 := HuntItem{
// 		Name:  "spider",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRdxV0T_nlCg3EgA46BvyXRkuX4FalBasG0cQ&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem8 := HuntItem{
// 		Name:  "berries",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQA4lk4w_DNj5Mm6bPD4w8KsI_inHCxQ9-i7w&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem9 := HuntItem{
// 		Name:  "chipmunk",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTnl-3cp_TfBT2F25tz8ZihKzZkEee8duB6pw&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem10 := HuntItem{
// 		Name:  "toad",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRuLbci8HinWnGJl-WeUJuVEYYMdkA8bjvl2g&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem11 := HuntItem{
// 		Name:  "mushroom",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSECsYtmDNxg1FhQa7rXCCv018T1M0O9epdFA&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem12 := HuntItem{
// 		Name:  "bird",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcT1aWz0Wij9qMUe260m-ObUkiXJU68hARHLSQ&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem13 := HuntItem{
// 		Name:  "ants",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcT_KxA2ZH6ZOi18r6JRkm5zFvXVGv5lklLcAw&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem14 := HuntItem{
// 		Name:  "spiderweb",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTS1cduh5ZhmZYwaaczOG4mRX2xiHBxCaMFLQ&usqp=CAU",
// 		Theme: "nature",
// 	}
// 	huntItem15 := HuntItem{
// 		Name:  "multi-colored leaf",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcT0u5-Y0oEUBJ0MTFvOCZv2m54aU02vSQagDA&usqp=CAU",
// 		Theme: "nature",
// 	}

// 	christmasHuntItem1 := HuntItem{
// 		Name:  "santa",
// 		Image: "https://s3.amazonaws.com/pas-wordpress-media/wp-content/uploads/2011/12/santa.jpg",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem2 := HuntItem{
// 		Name:  "wreath",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSpJx8mOVv_jEZu77fcVFFMjlXG-Ob5-HaC8A&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem3 := HuntItem{
// 		Name:  "reindeer",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcS4AIuRTy-_bT8nlUjQmrl9VzmA4X6DqrVU8A&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem4 := HuntItem{
// 		Name:  "christmas lights",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcR4zMdWpZ02kP5z5mV8CY9hil8yz88V9c215A&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem5 := HuntItem{
// 		Name:  "christmas tree",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSQJnMAUJKN4O9YG3SFKGO8aaPMKc_WRLW2dw&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem6 := HuntItem{
// 		Name:  "sleigh",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSn0LP4-WacBziLAJnXbI9GrR5nA9TbWv_S6Q&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem7 := HuntItem{
// 		Name:  "candy cane",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcS1mv5l5M5aKiRfnBN-AEdWJnubVGBnVCsOLA&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem8 := HuntItem{
// 		Name:  "snowman",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQVEuwvOIkzZUMRJdPs9pRqu8z64wPsWRcZRg&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem9 := HuntItem{
// 		Name:  "stocking",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRdWPg1Id4OH5gRpcIfOXN_B5xY1dj4XbfXMA&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem10 := HuntItem{
// 		Name:  "train",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSR9zD5Q7jJaLggDo7JHc7yeqi4VpvM8G9STQ&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem11 := HuntItem{
// 		Name:  "snowflake",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRZBzHOVC9ckFNWHIJYB4r6ToB3wDA9PstPZg&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem12 := HuntItem{
// 		Name:  "angel",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRDRROinCOhkXiV2uQNFK3L2Vw3oJDChMhv2A&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem13 := HuntItem{
// 		Name:  "elf",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSMX9zZ5Cw-H1yFvPJsEv6LmkBfjSzlPlWCNw&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem14 := HuntItem{
// 		Name:  "icicles",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRCkeDgAu_EAItv7bWWdQFExm8Ltce6FU2y6Q&usqp=CAU",
// 		Theme: "christmas",
// 	}
// 	christmasHuntItem15 := HuntItem{
// 		Name:  "penguin",
// 		Image: "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRy7MChV2YXf022PSSZcxSGeFujpKgUa8F4_w&usqp=CAU",
// 		Theme: "christmas",
// 	}

// 	db.Create(&huntItem)
// 	db.Create(&huntItem2)
// 	db.Create(&huntItem3)
// 	db.Create(&huntItem4)
// 	db.Create(&huntItem5)
// 	db.Create(&huntItem6)
// 	db.Create(&huntItem7)
// 	db.Create(&huntItem8)
// 	db.Create(&huntItem9)
// 	db.Create(&huntItem10)
// 	db.Create(&huntItem11)
// 	db.Create(&huntItem12)
// 	db.Create(&huntItem13)
// 	db.Create(&huntItem14)
// 	db.Create(&huntItem15)

// 	db.Create(&christmasHuntItem1)
// 	db.Create(&christmasHuntItem2)
// 	db.Create(&christmasHuntItem3)
// 	db.Create(&christmasHuntItem4)
// 	db.Create(&christmasHuntItem5)
// 	db.Create(&christmasHuntItem6)
// 	db.Create(&christmasHuntItem7)
// 	db.Create(&christmasHuntItem8)
// 	db.Create(&christmasHuntItem9)
// 	db.Create(&christmasHuntItem10)
// 	db.Create(&christmasHuntItem11)
// 	db.Create(&christmasHuntItem12)
// 	db.Create(&christmasHuntItem13)
// 	db.Create(&christmasHuntItem14)
// 	db.Create(&christmasHuntItem15)

// }
