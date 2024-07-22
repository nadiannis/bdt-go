package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	marshal()
	fmt.Println()
	unmarshal()
	fmt.Println()
	encode()
	fmt.Println()
	decode()
}

func marshal() {
	fmt.Println("marshal...")

	user := User{
		Name:  "john",
		Email: "john@gmail.com",
		Age:   42,
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("user: %+v\n", user)
	fmt.Println("json bytes: ", jsonBytes)
	fmt.Println("json string: ", string(jsonBytes))
}

func unmarshal() {
	fmt.Println("unmarshal...")

	jsonString := `{"name":"john","email":"john@gmail.com","age":42}`
	jsonBytes := []byte(jsonString)

	var user User

	err := json.Unmarshal(jsonBytes, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("json string: ", jsonString)
	fmt.Println("json bytes: ", jsonBytes)
	fmt.Printf("user: %+v\n", user)
}

func encode() {
	fmt.Println("encode...")

	user := User{
		Name:  "john",
		Email: "john@gmail.com",
		Age:   42,
	}

	fmt.Printf("user: %+v\n", user)

	err := json.NewEncoder(os.Stdout).Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func decode() {
	fmt.Println("decode...")

	file, err := os.Open("user.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User

	err = json.NewDecoder(file).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("user: %+v\n", user)
}
