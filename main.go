package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type ResponseDto struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []User `json:"data"`
}

func main() {
	httpNewRequest()
}

func httpGet() {
	var response ResponseDto
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

func httpNewRequest() {
	var response ResponseDto
	req, err := http.NewRequest("GET", "https://reqres.in/api/users", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 12345")
	client := http.Client{}
	resp, err := client.Do(req)
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
	fmt.Println()

	for _, user := range response.Data {
		fmt.Println("Id", user.ID)
		fmt.Println("Email", user.Email)
		fmt.Println("FirstName", user.FirstName)
		fmt.Println("LastName", user.LastName)
		fmt.Println("Avatar", user.Avatar)
		fmt.Println()
	}
}
