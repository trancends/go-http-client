package main

import (
	"bytes"
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

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"password"`
	Password string `json:"password"`
}

var baseURL = "https://reqres.in/api"

func main() {
	fmt.Println("Post: ")
	httpNewRequestPost()
	fmt.Println()
	fmt.Println("Delete: ")
	httpNewRequestDelete()
	fmt.Println()
	fmt.Println("Put: ")
	httpNewRequestPut()
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

func httpNewRequestPut() {
	req, err := http.NewRequest("PUT", "https://reqres.in/api/users/1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(bodyBytes))
}

func httpNewRequestDelete() {
	req, err := http.NewRequest("DELETE", "https://reqres.in/api/users/1", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(bodyBytes))
}

func httpNewRequestPost() {
	// param := url.Values{}
	// param.Set("username", "beni")
	// param.Set("email", "beni@mail.com")
	// param.Set("password", "1234")
	// payload := bytes.NewBufferString(param.Encode())
	// jsonStr := []byte(`{"username":"beni","email":"beni@mail.com","password":"1234"}`)
	//
	// payload := bytes.NewBuffer(jsonStr)
	values := map[string]string{"username": "beni", "email": "beni@mail.com", "password": "1234"}

	jsonValue, _ := json.Marshal(values)
	// req, err := http.NewRequest("POST", baseURL+"/register", jsonValue)
	// if err != nil {
	// 	panic(err)
	// }
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// client := http.Client{}
	resp, err := http.Post(baseURL+"/register", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyByte))
	fmt.Println(resp.StatusCode)
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
	fmt.Println("Paging:")
	fmt.Println("Page:", response.Page)
	fmt.Println("PerPage:", response.PerPage)
	fmt.Println("Total:", response.Total)
	fmt.Println("TotalPages:", response.TotalPages)
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
