package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter()
	//users
	router.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/create-user", newUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/delete-user/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", UpdateUser).Methods("PUT")
	http.ListenAndServe(":7000", router)
}
