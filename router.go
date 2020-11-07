package main

import (
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
	router.HandleFunc("/hunt-items", allItems).Methods("GET")
	router.HandleFunc("/hunt-items/{id}", getItem).Methods("GET")
	router.HandleFunc("/update-hunt-item/{id}", updateHuntItem).Methods("PUT")
	//hunt_lists
	router.HandleFunc("/hunt-lists", allHuntLists).Methods("GET")
	router.HandleFunc("/hunt-lists/{id}", getHuntList).Methods("GET")
	router.HandleFunc("/create-hunt-list", newList).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	http.ListenAndServe(":"+port, router)
}
