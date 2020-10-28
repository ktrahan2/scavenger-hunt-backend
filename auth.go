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
	ID    uint   `json:"id"`
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
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		var receivedUser User
		var user User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &receivedUser)
		db.Where("username = ?", receivedUser.Username).First(&user)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(receivedUser.Password))
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong please try again")
		} else {
			//then give them a token
			validToken, err := generateJWT()

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			fmt.Println(validToken)
			Response := JWTTOKEN{
				validToken,
				user.ID,
			}
			//need to send this token and user.id in response
			json.NewEncoder(w).Encode(Response)
		}
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}

//use this function to check for token in header.
//like creating a custom hunt list.

// func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) router.HandleFunc {
// 	return router.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] != nil {
// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return mySigningKey, nil
// 			})
// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}

// 	})
// }
