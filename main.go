package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	connectToDatabase()
	handleRequest()
}

func connectToDatabase() {

	host := os.Getenv("HOST")
	databaseUsername := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, databaseUsername, password, database)

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Successfully connected!")
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":7000", router)
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	log.Println("hey")
}
