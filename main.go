package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	// "github.com/ktrahan2/scavenger-hunt-backend/v2/models"
)

var db *gorm.DB

var err error

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	InitialMigration()
	handleRequest()
	defer db.Close()
}

// func connectToDatabase() {

// 	host := os.Getenv("DBHOST")
// 	databaseUsername := os.Getenv("USERNAME")
// 	password := os.Getenv("PASSWORD")
// 	database := os.Getenv("DATABASE")
// 	dbport := os.Getenv("DBPORT")

// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, dbport, databaseUsername, password, database)

// 	db, err = gorm.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")

// 	user := models.User{Username: "ktrain", Password: "123", Email: "ktrain@yahoo.com"}
// 	db.Create(&user)

// }

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	http.ListenAndServe(":7000", router)
	log.Println("Listening on 7000")
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
