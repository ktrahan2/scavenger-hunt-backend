# Scavenger Hunt API

Scavenger Hunt API helps the react native application, On The Hunt, communicate with the database. 

# Table Of Contents 
- [Description](https://github.com/ktrahan2/spacey-bois-backend/tree/main#description)
- [Example Code](https://github.com/ktrahan2/spacey-bois-backend/tree/main#example-code)
- [Technology Used](https://github.com/ktrahan2/spacey-bois-backend/tree/main#technology-used)
- [Setting up for the Application](https://github.com/ktrahan2/spacey-bois-backend/tree/main#setting-up-for-the-application)
- [Main Features](https://github.com/ktrahan2/spacey-bois-backend/tree/main#main-features)
- [Features in Progress](https://github.com/ktrahan2/spacey-bois-backend/tree/main#features-in-progress)
- [Contact Information](https://github.com/ktrahan2/spacey-bois-backend/tree/main#contact-information)
- [Link to Frontend Repo](https://github.com/ktrahan2/spacey-bois-backend/tree/main#link-to-frontend-repo)

## Description

Scavenger Hunt API was built with Golang using gorm and mux, and a PSQL database. There are 5 tables within the API creating two many to many relationships.
Each table comes with handlers to deal with the various request coming from On The Hunt. Auth was built into the app using JWT tokens and bcrypt. 

## Example Code 
This handler takes in a Hunt List ID and several Item ids and creates a row for each one in the Selected Items table. 
```
  func newSelectedItem(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)

		var incomingItems IncomingItems
		json.Unmarshal(reqBody, &incomingItems)
		items := incomingItems.HuntItemIDs
		var selectedItems []SelectedItem

		for i := 0; i < len(items); i++ {
			var selectedItem SelectedItem

			selectedItem = SelectedItem{
				HuntListID: incomingItems.HuntListID,
				HuntItemID: incomingItems.HuntItemIDs[i],
			}

			db.Create(&selectedItem)

			selectedItems = append(selectedItems, selectedItem)
		}

		json.NewEncoder(w).Encode(&selectedItems)

	default:
		http.Error(w, http.StatusText(405), 405)
	}
```
This function generates a JWT token to be sent to a user after that successfully login or signup.
```
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
```

## Technology Used

- Go
- Gorm
- Mux
- PSQL

## Main Features

- Full CRUD actions created for data manipulation
- Router built with mux
- Auth with bcrypt and JWT toke

## Features in Progress

- Add validations to incoming entries to ensure a clean database. 

## Contact Information

[Kyle Trahan](https://www.linkedin.com/in/kyle-trahan-8384678b/)

## Link to Frontend Repo
https://github.com/ktrahan2/scavenger-hunt-frontend



