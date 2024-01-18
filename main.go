package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Avatar    string
}

type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []User `json:"data"`
}

func main() {
	var response Response
	fmt.Println("Hello, World")
	resp, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(response)
}
