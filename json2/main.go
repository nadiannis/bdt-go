package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	http.HandleFunc("POST /users-inc", incrementAge)
	http.HandleFunc("POST /users-dec", decrementAge)
	http.HandleFunc("POST /users-double", doubleAge)

	fmt.Println("Listening on port", 8080)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

// Use decode-encode
func incrementAge(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Age += 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

// Use unmarshal-marshal
func decrementAge(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Age -= 1

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes = append(jsonBytes, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

// Use decode-marshal
func doubleAge(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Age *= 2

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes = append(jsonBytes, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}
