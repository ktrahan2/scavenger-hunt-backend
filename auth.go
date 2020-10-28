package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

//find the user by username, then compare that password with password sent in
func login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var receivedUser User
		json.Unmarshal(reqBody, &receivedUser)
		var user User
		db.Where("username = ?", receivedUser.Username).First(&user)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(receivedUser.Password))
		if err != nil {
			fmt.Println("Hello")
		}

		//then give them a token
		validToken, err := generateJWT()

		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(validToken)
		//need to send this token in response
	default:
		http.Error(w, http.StatusText(405), 405)
	}

}

//use this function to check for token in header.
//like creating a custom hunt list.

// func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) router.HandleFunc {
// 	return router.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {

// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}

// 	})
// }
