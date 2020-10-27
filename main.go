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
	"github.com/ktrahan2/scavenger-hunt-backend/v2/models"
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

	host := os.Getenv("DBHOST")
	databaseUsername := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DATABASE")
	dbport := os.Getenv("DBPORT")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, dbport, databaseUsername, password, database)

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	db.AutoMigrate(&models.User{})

	user := models.User{Username: "ktrain", Password: "123", Email: "ktrain@yahoo.com"}
	db.Create(&user)

	defer db.Close()
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	http.ListenAndServe(":"+port, nil)
}

//route methods
func getUsers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	log.Println("hey")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("hey")
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
