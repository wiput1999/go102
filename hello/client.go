package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetUsers() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var users Users

	r := json.NewDecoder(response.Body)

	err = r.Decode(&users)

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}

}

type Users []User

type User struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}
