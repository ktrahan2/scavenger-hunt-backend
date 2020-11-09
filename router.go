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
	router.HandleFunc("/create-hunt-list", newHuntList).Methods("POST", "OPTIONS")
	router.HandleFunc("/update-hunt-list/{id}", updateHuntList).Methods("PUT")
	//user_lists
	router.HandleFunc("/user-lists", allUserLists).Methods("GET")
	router.HandleFunc("/user-lists/{id}", getUserList).Methods("GET")
	router.HandleFunc("/create-user-list", newUserList).Methods("POST", "OPTIONS")
	router.HandleFunc("/update-user-list", updateUserList).Methods("PUT")
	//selected_items
	router.HandleFunc("/selected-items", allSelectedItems).Methods("GET")
	router.HandleFunc("/selected-items/{id}", getSelectedItem).Methods("GET")
	router.HandleFunc("/create-selected-item", newSelectedItem).Methods("POST", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	http.ListenAndServe(":"+port, router)
}
