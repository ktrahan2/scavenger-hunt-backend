package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter()
	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(JwtVerify)
	//users
	// secure.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/create-user", newUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/delete-user/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/login", login).Methods("OPTIONS", "POST")
	//hunt_items
	router.HandleFunc("/hunt_items", allItems).Methods("GET")
	router.HandleFunc("/hunt_items/{id}", getItem).Methods("GET")
	// router.HandleFunc("/update-hunt-item", updateHuntItem).Methods("PUT")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("OH NO")
	}

	http.ListenAndServe(":"+port, router)
}
