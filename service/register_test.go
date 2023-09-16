package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/register", bytes.NewBufferString(`{"username":"testuser","password":"testpass"}`))

	registerHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response["message"] != "User registered successfully" {
		t.Errorf("Expected message 'User registered successfully', got '%s'", response["message"])
	}
}

func TestIsValidUser(t *testing.T) {
	tests := []struct {
		name     string
		user     User
		expected bool
	}{
		{
			name:     "Valid user",
			user:     User{Username: "testuser", Password: "testpass"},
			expected: true,
		},
		{
			name:     "Invalid password",
			user:     User{Username: "testuser", Password: "wrongpass"},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isValidUser(test.user)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
