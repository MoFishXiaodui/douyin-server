package server

import (
	"dy/model"
	"errors"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var registeredUsers = make(map[string]bool)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform validation and authentication logic here
	if !isValidUser(user) {
		http.Error(w, "Invalid user credentials", http.StatusUnauthorized)
		return
	}

	// Check if the username is already taken
	if isUsernameTaken(user.Username) {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	// Register the user
	registeredUsers[user.Username] = true

	// Return a success response
	response := map[string]string{
		"message": "User registered successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func isValidUser(user User) bool {
	// Add your own validation logic here
	return true
}

func isUsernameTaken(username string) bool {
	_, exists := registeredUsers[username]
	return exists
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":8080", nil)
}
