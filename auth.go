package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//JWTTOKEN is the structure of a Token
type JWTTOKEN struct {
	Token     string `json:"token"`
	FoundUser User   `json:"user"`
}

//Claims is the structure of the claims from a token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Allows a user to login and receive a JWT token
func login(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":

		var receivedUser User
		var foundUser User

		reqBody, _ := ioutil.ReadAll(r.Body)

		json.Unmarshal(reqBody, &receivedUser)

		db.Preload("HuntList.HuntItems").Where("username = ?", receivedUser.Username).First(&foundUser)

		err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(receivedUser.Password))

		if err != nil {
			json.NewEncoder(w).Encode("Unathorized User Information")
		} else {
			validToken, err := generateJWT()

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			Response := JWTTOKEN{
				validToken,
				foundUser,
			}
			json.NewEncoder(w).Encode(Response)
		}
	default:
		http.Error(w, http.StatusText(405), 405)
	}
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

//JwtVerify verifies the JWT token sent in the request header.
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte(os.Getenv("SECRET"))
		var header = r.Header.Get("Token")
		header = strings.TrimSpace(header)
		claims := &Claims{}

		validToken, err := jwt.ParseWithClaims(header, claims, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !validToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
