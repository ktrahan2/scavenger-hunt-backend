package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTTOKEN is the structure of a Token
type JWTTOKEN struct {
	Token string `json:"token"`
}

func generateJWT() (string, error) {
	mySigningKey := []byte(os.Getenv("SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	validToken, err := generateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println(validToken)
	//need to send this token in response
}

// func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) router.HandleFunc {
// 	return router.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {

// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}

// 	})
// }
