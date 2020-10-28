package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter()
	secure := router.PathPrefix("/auth").Subrouter()
	secure.Use(JwtVerify)
	//users
	secure.HandleFunc("/users", allUsers).Methods("GET")
	router.HandleFunc("/create-user", newUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/delete-user/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/update-user/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/login", login).Methods("POST")
	http.ListenAndServe(":7000", router)
}
