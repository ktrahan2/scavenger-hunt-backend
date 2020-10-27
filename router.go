package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users", newUser).Methods("POST")
	// router.HandleFunc("/user/{username}", GetUser).Methods("GET")
	// router.HandleFunc("/user/{username}/{password}/{email}", DeleteUser).Methods("DELETE")
	// router.HandleFunc("user/id", UpdateUser).Methods("PUT")
	http.ListenAndServe(":7000", router)
}
